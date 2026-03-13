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

```hcl
resource "ibm_logs_extension_deployment" "logs_extension_deployment_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  applications = ltest_0401
  item_ids = b9a5500c-715e-4ead-9bbe-56fdefffbfcd
  subsystems = ltest_0401
  version = "version"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `applications` - (Optional, List) Applications that the Extension is deployed for. When this is empty, it is applied to all applications.
  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
* `item_ids` - (Required, List) The list of Extension item IDs to deploy.
  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `1` item.
* `subsystems` - (Optional, List) Subsystems that the Extension is deployed. When this is empty, it is applied to all subsystems.
  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
* `version` - (Required, String) The version of the Extension revision to deploy.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the Extension deployment.


## Import

You can import the `ibm_logs_extension_deployment` resource by using `id`. The unique identifier of the extension.

# Syntax
<pre>
$ terraform import ibm_logs_extension_deployment.logs_extension_deployment &lt;id&gt;
</pre>

# Example
```
$ terraform import ibm_logs_extension_deployment.logs_extension_deployment IBMCloudKubernetes
```
