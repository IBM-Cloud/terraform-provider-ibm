data "ibm_pi_tenant" "ds_tenant" {
  pi_cloud_instance_id = var.powerinstanceid
}

data "ibm_pi_volume" "ds_volume" {
  pi_volume_name       = "myvol"
  pi_cloud_instance_id = var.powerinstanceid
}

data "ibm_pi_instance_volumes" "ds_volumes" {
  pi_instance_name     = "mypi"
  pi_cloud_instance_id = var.powerinstanceid
}

data "ibm_pi_key" "ds_instance" {
  pi_key_name          = "mykey"
  pi_cloud_instance_id = var.powerinstanceid
}

data "ibm_pi_instance" "ds_instance" {
  pi_instance_name     = "mypi"
  pi_cloud_instance_id = var.powerinstanceid
}

data "ibm_pi_images" "ds_images" {
	pi_image_name        = "my_pi_images"  
	pi_cloud_instance_id = var.powerinstanceid
}

resource "ibm_pi_volume" "testacc_volume"{
  pi_volume_size       = 20
  pi_volume_name       = "test-volume22"
  pi_volume_type       = "ssd"
  pi_volume_shareable  = true
  pi_cloud_instance_id = var.powerinstanceid    // Get ot by running cmd "ic resource service-instances --long"
}


data "ibm_pi_image" "ds_image" {
  pi_image_name        = "7200-03-03"
  pi_cloud_instance_id = var.powerinstanceid
}


/*resource "ibm_pi_image" "testacc_image" {
  pi_image_name       = "7200-03-02"
  pi_image_id         = data.ibm_pi_image.ds_image.id
  pi_cloud_instance_id = var.powerinstanceid
}*/

resource "ibm_pi_key" "testacc_sshkey2" {
  pi_key_name          =  "mykey22"
  pi_ssh_key           =  var.sshkeyname
  pi_cloud_instance_id =  var.powerinstanceid 
}




