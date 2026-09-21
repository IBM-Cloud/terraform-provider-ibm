---

subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_service_id"
description: |-
  Manages IBM IAM service ID.
---

# ibm_iam_service_id

Create, update, and delete an IAM service ID with this resource.  For more information, about IAM role action, see [managing service ID API keys](https://cloud.ibm.com/docs/account?topic=account-serviceidapikeys).

## Example Usage

```hcl
resource "ibm_iam_service_id" "iam_service_id_instance" {
  name        = "test"
  description = "New ServiceID"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `name` - (Required, String) Name of the Service Id. The name is not checked for uniqueness. Therefore multiple names with the same value can exist. Access is done via the UUID of the Service Id.
* `description` - (Optional, String) The optional description of the Service Id. The 'description' property is only available if a description was provided during a create of a Service Id.
* `tags` (Optional, Array of Strings) A list of tags that you want to add to the service ID. **Note** The tags are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the iam_service_id.
* `created_at` - (String) If set contains a date time string of the creation date in ISO format.
* `crn` - (String) Cloud Resource Name of the item. Example Cloud Resource Name: 'crn:v1:bluemix:public:iam-identity:us-south:a/myaccount::serviceid:1234-5678-9012'.
* `iam_id` - (String) Cloud wide identifier for identities of this service ID.
* `locked` - (Boolean) The service ID cannot be changed if set to true.
* `modified_at` - (String) If set contains a date time string of the last modification date in ISO format.
* `version` - (String) The version of the ServiceID object.


## Import

You can import the `ibm_iam_service_id` resource by using `id`. Unique identifier of this Service Id.

# Syntax
<pre>
$ terraform import ibm_iam_service_id.iam_service_id &lt;id&gt;
</pre>
