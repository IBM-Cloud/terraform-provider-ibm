---
layout: "ibm"
page_title: "IBM : ibm_mqcloud_queue_manager"
description: |-
  Manages mqcloud_queue_manager.
subcategory: "MQ SaaS"
---

# ibm_mqcloud_queue_manager

Create, update, and delete mqcloud_queue_managers with this resource.

> **Note:** The MQaaS Terraform provider access is restricted to users of the reserved deployment, reserved capacity, and reserved capacity subscription plans.

## Example Usage

```hcl
resource "ibm_mqcloud_queue_manager" "mqcloud_queue_manager_instance" {
  display_name = "A test queue manager"
  location = "reserved-eu-de-cluster-f884"
  name = "testqm"
  service_instance_guid = "a2b4d4bc-dadb-4637-bcec-9b7d1e723af8"
  size = "small"
  version = "9.3.2_2"
}
```

## Timeouts

mqcloud_queue_manager provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 15 minutes) Used for creating a mqcloud_queue_manager.
* `update` - (Default 5 minutes) Used for updating a mqcloud_queue_manager.
* `delete` - (Default 5 minutes) Used for deleting a mqcloud_queue_manager.

## Argument Reference

You can specify the following arguments for this resource.

* `display_name` - (Optional, Forces new resource, String) A displayable name for the queue manager - limited only in length.
  * Constraints: The maximum length is `150` characters.
* `location` - (Required, Forces new resource, String) The locations in which the queue manager could be deployed.
  * Constraints: The maximum length is `150` characters. The minimum length is `2` characters. The value must match regular expression `/^([^[:ascii:]]|[a-zA-Z0-9-._: ])+$/`. Details of applicable locations can be found from either the use of the `ibm_mqcloud_queue_manager_options` datasource for the resource instance or can be found using the [IBM API for MQaaS](https://cloud.ibm.com/apidocs/mq-on-cloud) and be set as a variable.
* `name` - (Required, Forces new resource, String) A queue manager name conforming to MQ restrictions.
  * Constraints: The maximum length is `48` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9._]*$/`.
* `service_instance_guid` - (Required, Forces new resource, String) The GUID that uniquely identifies the MQ SaaS service instance.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.
* `size` - (Required, Forces new resource, String) The queue manager sizes of deployment available.
  * Constraints: Allowable values are: `xsmall`, `small`, `medium`, `large`.
* `version` - (Optional, String) The MQ version of the queue manager.
  * Constraints: The maximum length is `15` characters. The minimum length is `7` characters. The value must match regular expression `/^[0-9]+.[0-9]+.[0-9]+_[0-9]+$/`. Details of applicable versions can be found from either the use of the `ibm_mqcloud_queue_manager_options` datasource for the resource instance, can be found using the [IBM API for MQaaS](https://cloud.ibm.com/apidocs/mq-on-cloud) or with the variable not included at all to default to the latest version.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the mqcloud_queue_manager.
* `administrator_api_endpoint_url` - (String) The url through which to access the Admin REST APIs for this queue manager.
* `available_upgrade_versions_uri` - (String) The uri through which the available versions to upgrade to can be found for this queue manager.
* `connection_info_uri` - (String) The uri through which the CDDT for this queue manager can be obtained.
* `date_created` - (String) RFC3339 formatted UTC date for when the queue manager was created.
* `href` - (String) The URL for this queue manager.
* `queue_manager_id` - (String) The ID of the queue manager which was allocated on creation, and can be used for delete calls.
* `rest_api_endpoint_url` - (String) The url through which to access REST APIs for this queue manager.
* `status_uri` - (String) A reference uri to get deployment status of the queue manager.
* `upgrade_available` - (Boolean) Describes whether an upgrade is available for this queue manager.
* `web_console_url` - (String) The url through which to access the web console for this queue manager.


## Import

You can import the `ibm_mqcloud_queue_manager` resource by using `id`.
The `id` property can be formed from `service_instance_guid`, and `queue_manager_id` in the following format:

<pre>
&lt;service_instance_guid&gt;/&lt;queue_manager_id&gt;
</pre>
* `service_instance_guid`: A string in the format `a2b4d4bc-dadb-4637-bcec-9b7d1e723af8`. The GUID that uniquely identifies the MQ SaaS service instance.
* `queue_manager_id`: A string in the format `b8e1aeda078009cf3db74e90d5d42328`. The ID of the queue manager which was allocated on creation, and can be used for delete calls.

# Syntax
<pre>
$ terraform import ibm_mqcloud_queue_manager.mqcloud_queue_manager &lt;service_instance_guid&gt;/&lt;queue_manager_id&gt;
</pre>
