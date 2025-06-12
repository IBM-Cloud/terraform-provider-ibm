---

subcategory: "Resource management"
layout: "ibm"
page_title: "IBM: ibm_resource_key"
description: |-
  Get information about a resource key from IBM Cloud.
---

# ibm_resource_key

Retrieve information about an existing IBM resource key from IBM Cloud as a read-only data source. For more information, about resource key, see [ibmcloud resource service-keys](https://cloud.ibm.com/docs/account?topic=cli-ibmcloud_commands_resource#ibmcloud_resource_service_keys).

## Example usage

```terraform
data "ibm_resource_key" "resourceKeydata" {
  name                  = "myobjectKey"
  resource_instance_id  = ibm_resource_instance.resource.id
}
```
### Example to access resource credentials using credentials attribute:

```terraform
data "ibm_resource_key" "key" {
  name                  = "myobjectKey"
  resource_instance_id  = ibm_resource_instance.resource.id
}
output "access_key_id" {
  value = data.ibm_resource_key.key.credentials["cos_hmac_keys.access_key_id"]
}
output "secret_access_key" {
  value = data.ibm_resource_key.key.credentials["cos_hmac_keys.secret_access_key"]
}
```
### Example to access resource credentials using credentials_json attribute:

```terraform
data "ibm_resource_key" "key" {
  name                  = "myobjectKey"
  resource_instance_id  = ibm_resource_instance.resource.id
}
locals {
  resource_credentials = jsondecode(data.ibm_resource_key.key.credentials_json)
}
output "access_key_id" {
  value = local.resource_credentials.cos_hmac_keys.access_key_id
}
output "secret_access_key" {
  value = local.resource_credentials.cos_hmac_keys.secret_access_key
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `most_recent` - (Optional, Bool) If there are multiple resource keys, you can set this argument to `true` to import only the most recently created key.
- `name` - (Required, String) The name of the resource key. You can retrieve the value by executing the `ibmcloud resource service-keys` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).
- `resource_instance_id` - (Optional, string) The ID of the resource instance that the resource key is associated with. You can retrieve the value by executing the `ibmcloud resource service-instances` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started). **Note**: Conflicts with `resource_alias_id`.
- `resource_alias_id` - (Optional, String, Deprecated) The ID of the resource alias that the resource key is associated with. You can retrieve the value by executing the `ibmcloud resource service-alias` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started). **Note** Conflicts with `resource_instance_id`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `credentials` - (Map) The credentials associated with the key.
- `credentials_json` - (String) The credentials associated with the key in json format.
- `crn` - (String) CRN of resource key.
- `id` - (String) The unique identifier of the resource key.
- `role` - (String) The user role.
- `status` - (String) The status of the resource key.  
- `onetime_credentials` - (Bool) A boolean that dictates if the onetime_credentials is true or false.

## Note
Credentials will be seen as redacted, if the user does not have access equal to or greater than the access of the service credentials. Please refer to the documentation to access credentials - https://cloud.ibm.com/docs/account?topic=account-service_credentials&interface=ui#viewing-credentials-ui.
