---
layout: "ibm"
page_title: "IBM : ibm_pag_instance"
description: |-
  Create Privileged Access Gateway instance.
subcategory: "Privileged Access Gateway"
---

# ibm_pag_instance
Create, update, or delete the Privileged Access Gateway instance (PAG).

## Example usage

```terraform
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
    create = "1h"
    update = "1h"
    delete = "1h"
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
```



The `ibm_resource_instance` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 1 hour) Used for Creating Instance.
- **update** - (Default 1 hour) Used for Updating Instance.
- **delete** - (Default 1 hour) Used for Deleting Instance.

## Argument reference
Review the argument references that you can specify for your resource. 

- `location` - (Required, Forces new resource, String) Target location or environment to create the PAG instance.
- `parameters_json` - (Required, Forces new resource, String) Parameters to create PAG instance. The value must be a JSON string.

  Nested scheme for `parameters_json`:
  - `cosinstance` - (Required, String) COS instance CRN to use for PAG.
  - `cosbucket` - (Required, String) COS bucket name to use for PAG.
  - `cosendpoint` - (Required, String) COS endpoint to use for PAG.
  - `proxies` - (Required, List of objects)
  
    Nested scheme for `proxies`:
    - `securitygroups` - (Required, Set of strings) Security group(s) ID to use for PAG.
    - `subnet` - (Required, List of Objects) A nested block which requires subnet crn and cidr.
    
      Nested scheme for `subnet`:
      - `crn` - (Required, String) Subnet crn to use for PAG.
      - `cidr` - (Required, String) Subnet cidr to use for PAG.
    - `settings` - (Required, List of Strings) A nested setting block which requires inactivity timeout and system use notification.
    
      Nested scheme for `settings`:
      - `system_use_notification` - (Required, String) Message that is displayed when a user connects to PAG.
      - `inactivity_timeout` - (Required, Number) PAG inactivity timeout value (in minutes).
- `plan` - (Required, String) The name of the plan type supported by service i.e `standard`.
- `name` - (Required, String) A descriptive name used to identify the resource instance.
- `resource_group_id` - (Required, String) The ID of the resource group where you want to create the PAG service. You can retrieve the value from data source `ibm_resource_group`. If not provided creates the service in default resource group.
- `service` - (Required, String) The name of the service i.e `privileged-access-gateway`.
- `pag_vpc_id` - (Required, String) The ID of the VPC to be used for PAG.
- `tags` -  (Optional, Array of Strings) Tags associated with the PAG instance.




## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `account_id` - (String) An alpha-numeric value identifying the account ID.
- `allow_cleanup` - (String) A boolean that dictates if the resource instance should be deleted (cleaned up) during the processing of a region instance delete call.
- `created_at` - (Timestamp) The date when the instance  created.
- `created_by` - (String) The subject who created the instance.
- `guid` - (String) The GUID of the resource instance.
- `id` - (String) The unique identifier of the new resource instance.
- `last_operation` - (String) The status of the last operation requested on the instance.
- `locked` - (String) A boolean that dictates if the resource instance should be deleted (cleaned up) during the processing of a region instance delete call.
- `plan_history` - (String) The plan history of the instance.
- `resource_group_crn` - (String) The long ID (full CRN) of the resource group.
- `resource_id` - (String) The unique ID of the offering. This value is provided by and stored in the global catalog.
- `resource_plan_id` - (String) The unique ID of the plan associated with the offering. This value is provided by and stored in the global catalog.
- `resource_aliases_url` - (String) The relative path to the resource aliases for the instance.
- `resource_bindings_url` - (String) The relative path to the resource bindings for the instance.
- `resource_keys_url` - (String)  The relative path to the resource keys for the instance.
- `status` - (String) The status of resource instance.
- `state` - (String) The current state of the instance. For example, if the instance is deleted, it will return removed.
- `type` - (String) The type of the instance. For example, `service_instance`.
- `update_at` - (Timestamp) The date when the instance last updated.