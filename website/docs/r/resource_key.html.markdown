---

subcategory: "Resource management"
layout: "ibm"
page_title: "IBM : resource_key"
description: |-
  Manages IBM resource key.
---

# ibm_resource_key
Create, update, or delete service credentials for an IAM-enabled service. By default, the `ibm_resource_key` resource creates service credentials that use the public service endpoint of a service. To create service credentials that use the private service endpoint instead, you must explicitly define that by using the `parameter` argument parameter. Note that your service might not support private service endpoints yet. For more information, about resource key, see [adding and viewing credentials](https://cloud.ibm.com/docs/account?topic=account-service_credentials).

## Example usage
The following example enables to create credentials for a resource without a service ID.

```terraform
data "ibm_resource_instance" "resource_instance" {
  name = "myobjectsotrage"
}

resource "ibm_resource_key" "resourceKey" {
  name                 = "myobjectkey"
  role                 = "Viewer"
  resource_instance_id = data.ibm_resource_instance.resource_instance.id

  //User can increase timeouts
  timeouts {
    create = "15m"
    delete = "15m"
  }
}
```

**Note** The current `ibm_resource_key` resource doesn't have support for service_id argument but the service_id can be passed as one of the parameter.

### Example to create by using serviceID 

```terraform
data "ibm_resource_instance" "resource_instance" {
  name = "myobjectsotrage"
}

resource "ibm_iam_service_id" "serviceID" {
  name        = "test"
  description = "New ServiceID"
}

resource "ibm_resource_key" "resourceKey" {
  name                 = "myobjectkey"
  role                 = "Viewer"
  resource_instance_id = data.ibm_resource_instance.resource_instance.id
  parameters = {
    "serviceid_crn" = ibm_iam_service_id.serviceID.crn
  }

  //User can increase timeouts
  timeouts {
    create = "15m"
    delete = "15m"
  }
}
```
### Example to create by using HMAC 

```terraform
data "ibm_resource_group" "group" {
    name ="Default"
}
resource "ibm_resource_instance" "resource_instance" {
  name              = "test-21"
  service           = "cloud-object-storage"
  plan              = "lite"
  location          = "global"
  resource_group_id = data.ibm_resource_group.group.id
  tags              = ["tag1", "tag2"]
  
  //User can increase timeouts
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}
resource "ibm_resource_key" "resourceKey" {
  name                 = "my-cos-bucket-xx-key"
  resource_instance_id = ibm_resource_instance.resource_instance.id
  parameters           = { "HMAC" = true }
  role                 = "Manager"
}

```
### Example to access resource credentials using credentials attribute:

```terraform
resource "ibm_resource_key" "key" {
  name                 = "my-cos-bucket-xx-key"
  resource_instance_id = ibm_resource_instance.resource_instance.id
  role                 = "Manager"
}
output "access_key_id" {
  value = ibm_resource_key.key.credentials["cos_hmac_keys.access_key_id"]
}
output "secret_access_key" {
  value = ibm_resource_key.key.credentials["cos_hmac_keys.secret_access_key"]
}
```

### Example to access resource credentials using credentials_json attribute:

```terraform
resource "ibm_resource_key" "key" {
  name                 = "my-cos-bucket-xx-key"
  resource_instance_id = ibm_resource_instance.resource_instance.id
  role                 = "Manager"
}
locals {
  resource_credentials =jsondecode(ibm_resource_key.key.credentials_json)
}
output "access_key_id" {
  value = local.resource_credentials.cos_hmac_keys.access_key_id
}
output "secret_access_key" {
  value = local.resource_credentials.cos_hmac_keys.secret_access_key
}
```

## Timeouts

The `ibm_resource_key` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for Creating Key.
- **delete** - (Default 10 minutes) Used for Deleting Key.


## Argument reference
Review the argument references that you can specify for your resource. 

- `name` - (Required, Forces new resource, String)  A descriptive name used to identify a resource key.
- `parameters` (Optional, Map) Arbitrary parameters to pass to the resource in JSON format. If you want to create service credentials by using the private service endpoint, include the `service-endpoints =  "private"` parameter.
- `role` - (Optional, Forces new resource, String) The name of the user role. Valid roles are `NONE`,`Writer`, `Reader`, `Manager`, `Administrator`, `Operator`, `Viewer`, and `Editor`. This argument is Optional only during creation of service credentials for Cloud Databases and other non-IAM-enabled services and is Required for all other IAM-enabled services.
- `resource_instance_id` - (Optional, Forces new resource, String) The ID of the resource instance associated with the resource key. **Note** Conflicts with `resource_alias_id`.
- `resource_alias_id` - (Optional, Forces new resource, String, Deprecated) The ID of the resource alias associated with the resource key. **Note** Conflicts with `resource_instance_id`.
- `tags` (Optional, Array of strings) Tags associated with the resource key instance. **Note** Tags are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `account_id` - (String) An alpha-numeric value identifying the account ID.
- `credentials` - (Map) The credentials associated with the key.
- `credentials_json` - (String) The credentials associated with the key in json format.
- `created_at` - (Timestamp) The date when the key was created.
- `created_by` - (String) The subject who created the key.
- `crn` - (String) The full Cloud Resource Name (CRN) associated with the key.
- `deleted_at` - (Timestamp) The date when the key was deleted.
- `deleted_by` - (String) The subject who deleted the key.
- `id` - (String) The unique identifier of the new resource key.
- `status` - (String) The status of the resource key.
- `guid` - (String) A unique internal identifier GUID managed by the resource controller that corresponds to the key.
- `iam_compatible` - (String) Specifies whether the key’s credentials support IAM.
- `resource_group_id` - (String) The short ID of the resource group.
- `source_crn` - (String) The CRN of resource instance or alias associated to the key.
- `state` - (String) The state of the key.
- `resource_instance_url` - (String) The relative path to the resource.
- `updated_at` - (Timestamp) The date when the key was last updated.
- `updated_by` - (String) The subject who updated the key.
- `url` - (String) When you created a new key, a relative URL path is created identifying the location of the key.
- `onetime_credentials` - (Bool) A boolean that dictates if the onetime_credentials is true or false.

## Note
Credentials will be seen as redacted, if the user does not have access equal to or greater than the access of the service credentials. Please refer to the documentation to access credentials - https://cloud.ibm.com/docs/account?topic=account-service_credentials&interface=ui#viewing-credentials-ui.
