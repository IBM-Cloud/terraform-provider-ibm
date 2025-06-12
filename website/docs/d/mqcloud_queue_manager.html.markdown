---
layout: "ibm"
page_title: "IBM : ibm_mqcloud_queue_manager"
description: |-
  Get information about mqcloud_queue_manager
subcategory: "MQaaS"
---

# ibm_mqcloud_queue_manager

Provides a read-only data source to retrieve information about a mqcloud_queue_manager. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

> **Note:** The MQaaS Terraform provider access is restricted to users of the reserved deployment, reserved capacity, and reserved capacity subscription plans.

## Example Usage

```hcl
data "ibm_mqcloud_queue_manager" "mqcloud_queue_manager" {
	name = ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance.name
	service_instance_guid = ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance.service_instance_guid
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `name` - (Optional, String) A queue manager name conforming to MQ restrictions.
  * Constraints: The maximum length is `48` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9._]*$/`.
* `service_instance_guid` - (Required, String) The GUID that uniquely identifies the MQaaS service instance.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the mqcloud_queue_manager.
* `queue_managers` - (List) List of queue managers.
  * Constraints: The maximum length is `50` items. The minimum length is `0` items.
Nested schema for **queue_managers**:
	* `administrator_api_endpoint_url` - (String) The url through which to access the Admin REST APIs for this queue manager.
	* `available_upgrade_versions_uri` - (String) The uri through which the available versions to upgrade to can be found for this queue manager.
	* `connection_info_uri` - (String) The uri through which the CDDT for this queue manager can be obtained.
	* `date_created` - (String) RFC3339 formatted UTC date for when the queue manager was created.
	* `display_name` - (String) A displayable name for the queue manager - limited only in length.
	  * Constraints: The maximum length is `150` characters.
	* `href` - (String) The URL for this queue manager.
	* `id` - (String) The ID of the queue manager which was allocated on creation, and can be used for delete calls.
	* `location` - (String) The locations in which the queue manager could be deployed.
	  * Constraints: The maximum length is `150` characters. The minimum length is `2` characters. The value must match regular expression `/^([^[:ascii:]]|[a-zA-Z0-9-._: ])+$/`.
	* `name` - (String) A queue manager name conforming to MQ restrictions.
	  * Constraints: The maximum length is `48` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9._]*$/`.
	* `rest_api_endpoint_url` - (String) The url through which to access REST APIs for this queue manager.
	* `size` - (String) The queue manager sizes of deployment available.
	  * Constraints: Allowable values are: `xsmall`, `small`, `medium`, `large`.
	* `status_uri` - (String) A reference uri to get deployment status of the queue manager.
	* `upgrade_available` - (Boolean) Describes whether an upgrade is available for this queue manager.
	* `version` - (String) The MQ version of the queue manager.
	  * Constraints: The maximum length is `15` characters. The minimum length is `7` characters. The value must match regular expression `/^[0-9]+.[0-9]+.[0-9]+_[0-9]+$/`.
	* `web_console_url` - (String) The url through which to access the web console for this queue manager.

