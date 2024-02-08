---
layout: "ibm"
page_title: "IBM : ibm_mqcloud_queue_manager"
description: |-
  Manages mqcloud_queue_manager.
subcategory: "MQ on Cloud"
---

# ibm_mqcloud_queue_manager

Create, update, and delete mqcloud_queue_managers with this resource.

## Example Usage

```hcl
resource "ibm_resource_instance" "mqcloud_instance" {
    name     = "mqcloud-service-name"
    service  = "mqcloud"
    plan     = "default"
    location = "eu-de"
}

resource "ibm_mqcloud_queue_manager" "mqcloud_queue_manager_instance" {
  display_name = "A test queue manager"
  location = "reserved-eu-de-cluster-f884"
  name = "testqm"
  service_instance_guid = ibm_resource_instance.mqcloud_instance.guid
  size = "lite"
  version = "9.3.2_2"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `display_name` - (Optional, String) A displayable name for the queue manager - limited only in length.
  * Constraints: The maximum length is `150` characters.
* `location` - (Required, String) The locations in which the queue manager could be deployed.
  * Constraints: The maximum length is `150` characters. The minimum length is `2` characters. The value must match regular expression `/^([^[:ascii:]]|[a-zA-Z0-9-._: ])+$/`.
* `name` - (Required, String) A queue manager name conforming to MQ restrictions.
  * Constraints: The maximum length is `48` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9._]*$/`.
* `service_instance_guid` - (Required, Forces new resource, String) The GUID that uniquely identifies the MQ on Cloud service instance.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.
* `size` - (Required, String) The queue manager sizes of deployment available. Deployment of lite queue managers for aws_us_east_1 and aws_eu_west_1 locations is not available.
  * Constraints: Allowable values are: `lite`, `xsmall`, `small`, `medium`, `large`.
* `version` - (Optional, String) The MQ version of the queue manager.
  * Constraints: The maximum length is `15` characters. The minimum length is `7` characters. The value must match regular expression `/^[0-9]+.[0-9]+.[0-9]+_[0-9]+$/`.

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
* `service_instance_guid`: A string in the format `a2b4d4bc-dadb-4637-bcec-9b7d1e723af8`. The GUID that uniquely identifies the MQ on Cloud service instance.
* `queue_manager_id`: A string. The ID of the queue manager which was allocated on creation, and can be used for delete calls.

# Syntax
<pre>
$ terraform import ibm_mqcloud_queue_manager.mqcloud_queue_manager &lt;service_instance_guid&gt;/&lt;queue_manager_id&gt;
</pre>
