# --------------------------------------------------
# Provison the service using reosource_instance
# --------------------------------------------------

resource "ibm_resource_instance" "hpcs_instance" {
  count    = (var.provision_instance == true ? 1 : 0)
  name     = var.hpcs_instance_name
  service  = "hs-crypto"
  plan     = var.plan
  location = var.location
  parameters = {
    units = var.units
  }
}
# data source of the hpcs instance
data "ibm_resource_instance" "hpcs_instance" {
  name     = (var.provision_instance == true ? ibm_resource_instance.hpcs_instance.0.name : var.hpcs_instance_name)
  service  = "hs-crypto"
  location = var.location
}

# time_sleep after hpcs service instance was created
resource time_sleep "wait_3_mins" {
  depends_on      = [data.ibm_resource_instance.hpcs_instance]
  create_duration = "3m"
}

# ------------------------------------------------------------------------------------------------------------
# #Intialization of the hpcs crypto service may be use some scripts and null_resource to intialize the service
# -------------------------------------------------------------------------------------------------------------
/*
   Initilialising hpcs instance requires json file that containes all the creditials of admin and master keys as an input..
   This file can be fed as an input to the `hpcs_init` null resource in multiple ways out of which,
   the present template supports file from IBM-Object-Storage Bucket or directly from the local.
   Use null resource blocks accordingly..
*/

# This `download_from_cos` downloads json file cos-bucket..
/*
     Json file should already be fed to cos bucket before accessing it via null resource.
     This block takes cos credentials as input and and dowloads file in the current directory.
     This block can be commented if user doesnt wish to download input from cos-bucket
  */
resource "null_resource" "download_from_cos" {

  provisioner "local-exec" {
    command = <<EOT
    python ./scripts/download_from_cos.py
        EOT
    environment = {
      API_KEY         = var.api_key
      COS_SERVICE_CRN = var.cos_crn
      ENDPOINT        = var.endpoint
      BUCKET          = var.bucket_name
      INPUT_FILE_NAME = var.input_file_name
    }
  }
}

# This `hpcs_init` Initialises HPCS Instance by running all the tke command that are required for initialisation..
/*
     This take the content of the json file as input and perform necessary operations
     The set of CLOUDTKEFILES that are obtained as an input is stored in the `tke_files_path` provided by user as a folder of secrets.
  */
resource "null_resource" "hpcs_init" {
  depends_on = [null_resource.download_from_cos, time_sleep.wait_3_mins] // Dependson can be removed if the download_from_cos null resource is not used..
  provisioner "local-exec" {
    command = <<EOT
    python ./scripts/init.py
        EOT
    environment = {
      CLOUDTKEFILES = var.tke_files_path
      INPUT_FILE    = file(var.input_file_name)
      HPCS_GUID     = data.ibm_resource_instance.hpcs_instance.guid
    }
  }
}
# This `upload_to_cos` uploads CLOUDTKEFILES that are present in `tke_files_path` as a zip file .
/*
     This block takes cos credentials as input and and uploads zip file
     This block can be commented if user doesnt wish to upload files to cos-bucket
     Different cos credentials can also be used to upload files
  */
resource "null_resource" "upload_to_cos" {
  depends_on = [null_resource.hpcs_init]
  provisioner "local-exec" {
    command = <<EOT
    python ./scripts/upload_to_cos.py
        EOT
    environment = {
      API_KEY         = var.api_key
      COS_SERVICE_CRN = var.cos_crn
      ENDPOINT        = var.endpoint
      BUCKET          = var.bucket_name
      CLOUDTKEFILES   = var.tke_files_path
      HPCS_GUID       = data.ibm_resource_instance.hpcs_instance.guid
    }
  }
}
# This `remove_tke_files` removes CLOUDTKEFILES that are present in `tke_files_path` .
/*
     NOTE: This block has to be used only if user wish to delete CLOUDTKEFILES or the input file
  */
resource "null_resource" "remove_tke_files" {
  depends_on = [null_resource.upload_to_cos]
  provisioner "local-exec" {
    command = <<EOT
    python ./scripts/remove_tkefiles.py
        EOT
    environment = {
      CLOUDTKEFILES   = var.tke_files_path
      INPUT_FILE_NAME = var.input_file_name
      HPCS_GUID       = data.ibm_resource_instance.hpcs_instance.guid
    }
  }
}

# --------------------------------
# Cresting Keys for HPCS Instance
# --------------------------------

resource "ibm_kms_key" "key" {
  depends_on   = [null_resource.hpcs_init]
  instance_id  = data.ibm_resource_instance.hpcs_instance.guid
  key_name     = var.key_name
  standard_key = false
  force_delete = true
}
