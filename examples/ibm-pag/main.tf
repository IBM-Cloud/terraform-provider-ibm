data "ibm_resource_group" "pag" {
  name = var.ibm_resource_group_name
}

data "ibm_resource_instance" "pag-cos" {
  name              = var.ibm_cos_instance_name
  resource_group_id = data.ibm_resource_group.pag.id
  service           = "cloud-object-storage"
}

data "ibm_cos_bucket" "pag-cos-bucket" {
  bucket_name          = var.ibm_cos_bucket_name
  resource_instance_id = data.ibm_resource_instance.pag-cos.id
  bucket_type          = var.ibm_cos_bucket_type
  bucket_region        = var.ibm_cos_bucket_region
}

data "ibm_is_vpc" "pag" {
  name = var.ibm_vpc_name
}

data "ibm_is_subnet" "pag_instance_1" {
  name = var.ibm_vpc_subnet_name_instance_1
}

data "ibm_is_subnet" "pag_instance_2" {
  name = var.ibm_vpc_subnet_name_instance_2
}

data "ibm_is_security_group" "pag_instance" {
  name     = each.value
  for_each = var.ibm_vpc_security_groups_instance
}

resource "ibm_pag_instance" "pag" {
  name              = var.ibm_pag_instance_name
  resource_group_id = data.ibm_resource_group.pag.id
  service           = "privileged-access-gateway"
  plan              = var.ibm_pag_service_plan
  location          = var.region
  parameters_json = jsonencode(
    {
      "cosinstance" : data.ibm_resource_instance.pag-cos.crn,
      "cosbucket" : var.ibm_cos_bucket_name,
      "cosendpoint" : data.ibm_cos_bucket.pag-cos-bucket.s3_endpoint_direct
      "proxies" : [
        {
          "name" : "proxy1",
          "securitygroups" : [for sg in data.ibm_is_security_group.pag_instance : sg.id],
          "subnet" : {
            "crn" : data.ibm_is_subnet.pag_instance_1.crn,
            "cidr" : data.ibm_is_subnet.pag_instance_1.ipv4_cidr_block
          }
        },
        {
          "name" : "proxy2",
          "securitygroups" : [for sg in data.ibm_is_security_group.pag_instance : sg.id],
          "subnet" : {
            "crn" : data.ibm_is_subnet.pag_instance_2.crn,
            "cidr" : data.ibm_is_subnet.pag_instance_2.ipv4_cidr_block
          }
        }
       ],
       "settings" : {
        "inactivity_timeout" : var.pag_inactivity_timeout,
        "system_use_notification" : var.system_use_notification
     },
    "vpc_id" : data.ibm_is_vpc.pag.id
    }
  )
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}

resource "ibm_iam_authorization_policy" "pag-cos-iam-policy" {
  source_service_name         = "privileged-access-gateway"
  source_resource_instance_id = ibm_pag_instance.pag.guid
  roles                       = ["Object Writer"]
  resource_attributes {
    name     = "serviceName"
    operator = "stringEquals"
    value    = "cloud-object-storage"
  }

  resource_attributes {
    name     = "accountId"
    operator = "stringEquals"
    value    = data.ibm_resource_group.pag.account_id
  }
  resource_attributes {
    name     = "serviceInstance"
    operator = "stringEquals"
    value    = data.ibm_resource_instance.pag-cos.guid
  }

  resource_attributes {
    name     = "resourceType"
    operator = "stringEquals"
    value    = "bucket"
  }

  resource_attributes {
    name     = "resource"
    operator = "stringEquals"
    value    = var.ibm_cos_bucket_name
  }

}


locals {
  pag_hostnames = [for i in range(var.num_instances) : join(".", ["${ibm_pag_instance.pag.guid}-${i + 1}", "${ibm_pag_instance.pag.location}", "pag", "appdomain", "cloud"])]
}
output "pag-hosts" {
  value = local.pag_hostnames
}