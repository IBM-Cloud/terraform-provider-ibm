---
layout: "ibm"
page_title: "IBM : ibm_mqcloud_queue_manager_status"
description: |-
  Get information about mqcloud_queue_manager_status
subcategory: "MQ SaaS"
---

# ibm_mqcloud_queue_manager_status

Provides a read-only data source to retrieve information about mqcloud_queue_manager_status. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

> **Note:** The MQaaS Terraform provider access is restricted to users of the reserved deployment, reserved capacity, and reserved capacity subscription plans.

## Example Usage

```hcl
data "ibm_mqcloud_queue_manager_status" "mqcloud_queue_manager_status" {
	queue_manager_id = "b8e1aeda078009cf3db74e90d5d42328"
	service_instance_guid = "a2b4d4bc-dadb-4637-bcec-9b7d1e723af8"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `queue_manager_id` - (Required, Forces new resource, String) The id of the queue manager to retrieve its full details.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[0-9a-fA-F]{32}$/`.
* `service_instance_guid` - (Required, Forces new resource, String) The GUID that uniquely identifies the MQ SaaS service instance.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the mqcloud_queue_manager_status.
* `status` - (String) The deploying and failed states are not queue manager states, they are states which can occur when the request to deploy has been fired, or with that request has failed without producing a queue manager to have any state. The other states map to the queue manager states. State "ending" is either quiesing or ending immediately. State "ended" is either ended normally or endedimmediately. The others map one to one with queue manager states.
  * Constraints: Allowable values are: `initializing`, `deploying`, `starting`, `running`, `stopping`, `stopped`, `status_not_available`, `deleting`, `failed`, `upgrading_version`, `updating_revision`, `initialization_failed`, `restoring_queue_manager`, `restoring_config`, `restore_failed`, `suspended`, `resumable`.

