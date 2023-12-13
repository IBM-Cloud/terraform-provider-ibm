---
layout: "ibm"
page_title: "IBM : ibm_mqcloud_application"
description: |-
  Manages mqcloud_application.
subcategory: "MQ on Cloud"
---

# ibm_mqcloud_application

Create, update, and delete mqcloud_applications with this resource.

## Example Usage

```hcl
resource "ibm_mqcloud_application" "mqcloud_application_instance" {
  name = "test-app"
  service_instance_guid = var.service_instance_guid
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `name` - (Required, Forces new resource, String) The name of the application - conforming to MQ rules.
  * Constraints: The maximum length is `12` characters. The minimum length is `1` character.
* `service_instance_guid` - (Required, Forces new resource, String) The GUID that uniquely identifies the MQ on Cloud service instance.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the mqcloud_application.
* `application_id` - (String) The ID of the application which was allocated on creation, and can be used for delete calls.
* `create_api_key_uri` - (String) The URI to create a new apikey for the application.
* `href` - (String) The URL for this application.


## Import

You can import the `ibm_mqcloud_application` resource by using `id`.
The `id` property can be formed from `service_instance_guid`, and `application_id` in the following format:

<pre>
&lt;service_instance_guid&gt;/&lt;application_id&gt;
</pre>
* `service_instance_guid`: A string in the format `a2b4d4bc-dadb-4637-bcec-9b7d1e723af8`. The GUID that uniquely identifies the MQ on Cloud service instance.
* `application_id`: A string. The ID of the application which was allocated on creation, and can be used for delete calls.

# Syntax
<pre>
$ terraform import ibm_mqcloud_application.mqcloud_application &lt;service_instance_guid&gt;/&lt;application_id&gt;
</pre>
