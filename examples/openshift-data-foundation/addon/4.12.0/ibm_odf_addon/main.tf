terraform {
  required_providers {
    ibm = {
      source = "IBM-Cloud/ibm"
      version = "1.56.0"
    }
  }
}

provider "ibm" {
  region             = var.region
  ibmcloud_api_key = var.ibmcloud_api_key
}


resource "ibm_container_addons" "addons" {
    
  manage_all_addons = "false"
  cluster = var.cluster

  addons {

    name    = "openshift-data-foundation"
    version = var.odfVersion
    parameters_json = <<PARAMETERS_JSON
      {
        "odfDeploy":"false"
      } 
        PARAMETERS_JSON 
   
   }

}