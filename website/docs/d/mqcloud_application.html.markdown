---
layout: "ibm"
page_title: "IBM : ibm_mqcloud_application"
description: |-
  Get information about mqcloud_application
subcategory: "MQ SaaS"
---

# ibm_mqcloud_application

Provides a read-only data source to retrieve information about a mqcloud_application. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_mqcloud_application" "mqcloud_application" {
	name = ibm_mqcloud_application.mqcloud_application_instance.name
	service_instance_guid = ibm_mqcloud_application.mqcloud_application_instance.service_instance_guid
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `name` - (Optional, String) The name of the application - conforming to MQ rules.
  * Constraints: The maximum length is `12` characters. The minimum length is `1` character.
* `service_instance_guid` - (Required, Forces new resource, String) The GUID that uniquely identifies the MQ SaaS service instance.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the mqcloud_application.
* `applications` - (List) List of applications.
  * Constraints: The maximum length is `50` items. The minimum length is `0` items.
Nested schema for **applications**:
	* `create_api_key_uri` - (String) The URI to create a new apikey for the application.
	* `href` - (String) The URL for this application.
	* `iam_service_id` - (String) The IAM ID of the application.
	  * Constraints: The maximum length is `50` characters. The minimum length is `5` characters.
	* `id` - (String) The ID of the application which was allocated on creation, and can be used for delete calls.
	* `name` - (String) The name of the application - conforming to MQ rules.
	  * Constraints: The maximum length is `12` characters. The minimum length is `1` character.

