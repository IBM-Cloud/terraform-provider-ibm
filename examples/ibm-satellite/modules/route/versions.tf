terraform {
  required_providers {
    restapi = {
      source  = "fmontezuma/restapi"
      version = "1.14.1"
    }
    ibm = {
      source = "IBM-Cloud/ibm"
    }
  }
}

provider "restapi" {
  uri   = fileexists("token.log") ? var.cluster_master_url : "test.com"
  debug = true
  headers = {
    Authorization = var.is_endpoint_provision ? format("Bearer %v", chomp(element(tolist(data.local_file.token_file.*.content), 0))) : ""
  }
}