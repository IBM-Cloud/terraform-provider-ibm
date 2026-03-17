---
layout: "ibm"
page_title: "IBM : ibm_logs_extension_deployment"
description: |-
  Manages Extension deployment.
subcategory: "Cloud Logs"
---

# ibm_logs_extension_deployment

Create, update, and delete Extension deployments with this resource.

## Example Usage

### Deploy IBMCloudKubernetes extension

```hcl
data "ibm_logs_extensions" "logs_extensions" {
  instance_id       = ibm_resource_instance.logs_instance.guid
  region            = ibm_resource_instance.logs_instance.location
}
resource "ibm_logs_extension_deployment" "logs_extension_deployment_instance" {
  instance_id       = ibm_resource_instance.logs_instance.guid
  region            = ibm_resource_instance.logs_instance.location
  logs_extension_id = "IBMCloudKubernetes" # get extension id from above datasource
  version           = "1.0.0" # value of data.ibm_logs_extension.extension.revisions.0.version
  item_ids = [
    "680b64c9-6cc3-40ca-b1c3-4f6f91470906",
    "72a12a49-a659-484f-a230-0e7ab71600f1",
    "9e6ced78-bea4-4580-b1d0-6f45b1e1ecd7",
    "cd487925-85ea-4c52-b008-e121a470f22c",
  ] # values from [for item in data.ibm_logs_extension.extension.revisions[0].items : item.id]
  applications = ["test-application"]
  subsystems   = ["test-subsystem"]
}
```


## Argument Reference

After your resource is created, you can read values from the listed arguments and the following attributes.



* `instance_id` - (Required, Forces new resource, String)  Cloud Logs Instance GUID.
* `region` - (Optional, Forces new resource, String) Cloud Logs Instance Region.
* `endpoint_type` - (Optional, String) Cloud Logs Instance Endpoint type. Allowed values `public` and `private`.

* `logs_extension_id` - (Required, Forces new resource,String) The unique ID of extension to deploy.
* `item_ids` - (Required, Set) The list of Extension item IDs to deploy.
  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `1` item.
* `version` - (Required, String) The version of the Extension revision to deploy.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
* `applications` - (Optional, Set) Applications that the Extension is deployed for. When this is empty, it is applied to all applications.
  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
* `subsystems` - (Optional, Set) Subsystems that the Extension is deployed. When this is empty, it is applied to all subsystems.
  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `0` items.


## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the ibm_logs_extension_deployment resource.


## Import

You can import the `ibm_logs_extension_deployment` resource by using `id`.  You can import the `ibm_logs_e2m` resource by using `id`. `id` combination of `region`, `instance_id` and `logs_extension_id`.
# Syntax
<pre>
$ terraform import ibm_logs_extension_deployment.logs_extension_deployment < region >/< instance_id >/< logs_extension_id >;
</pre>

# Example
```
$ terraform import ibm_logs_extension_deployment.logs_extension_deployment au-syd/df7f8517-b8be-4b4b-9eb4-617e81487d2e/IBMPostgreSQL

```
