---
layout: "ibm"
page_title: "IBM : ibm_code_engine_persistent_data_store"
description: |-
  Manages code_engine_persistent_data_store.
subcategory: "Code Engine"
---

# ibm_code_engine_persistent_data_store

Create, update, and delete code_engine_persistent_data_stores with this resource.

## Example Usage

```hcl
resource "ibm_code_engine_persistent_data_store" "code_engine_persistent_data_store_instance" {
  project_id = ibm_code_engine_project.code_engine_project_instance.project_id
  name = "my-persistent-data-store"
  data {
		bucket_location = "au-syd"
		bucket_name     = "bucket_name"
		secret_name     = "secret_name"
  }
  storage_type = "object_storage"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `data` - (Optional, Forces new resource, List) Data container that allows to specify config parameters and their values as a key-value map. Each key field must consist of alphanumeric characters, `-`, `_` or `.` and must not exceed a max length of 253 characters. Each value field can consists of any character and must not exceed a max length of 1048576 characters.
Nested schema for **data**:
	* `bucket_location` - (Optional, String) Specify the location of the bucket.
	  * Constraints: Allowable values are: `au-syd`, `br-sao`, `ca-mon`, `ca-tor`, `eu-de`, `eu-es`, `eu-gb`, `jp-osa`, `jp-tok`, `us-east`, `us-south`, `ap`, `eu`, `us`, `ams03`, `che01`, `mil01`, `mon01`, `par01`, `sjc04`, `sng01`. The value must match regular expression `/^(au-syd|br-sao|ca-mon|ca-tor|eu-de|eu-es|eu-gb|jp-osa|jp-tok|us-east|us-south|ap|eu|us|ams03|che01|mil01|mon01|par01|sjc04|sng01)$/`.
	* `bucket_name` - (Optional, String) Specify the name of the bucket.
	  * Constraints: The maximum length is `63` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-z0-9]([-a-z0-9]*[a-z0-9])?$/`.
	* `secret_name` - (Optional, String) Specify the name of the HMAC secret.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
* `name` - (Required, Forces new resource, String) The name of the persistent data store.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `storage_type` - (Required, Forces new resource, String) Specify the storage type of the persistent data store.
  * Constraints: Allowable values are: `object_storage`. The value must match regular expression `/^(object_storage)$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the code_engine_persistent_data_store.
* `persistent_data_store_id` - (String) The identifier of the resource.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `created_at` - (String) The timestamp when the resource was created.
* `entity_tag` - (String) The version of the persistent data store, which is used to achieve optimistic locking.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[\\*\\-a-z0-9]+$/`.
* `id` - (String) The identifier of the resource.
* `region` - (String) The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'.

* `etag` - ETag identifier for code_engine_persistent_data_store.

## Import

You can import the `ibm_code_engine_persistent_data_store` resource by using `name`.
The `name` property can be formed from `project_id`, and `name` in the following format:

<pre>
&lt;project_id&gt;/&lt;name&gt;
</pre>
* `project_id`: A string in the format `15314cc3-85b4-4338-903f-c28cdee6d005`. The ID of the project.
* `name`: A string in the format `my-persistent-data-store`. The name of the persistent data store.

# Syntax
<pre>
$ terraform import ibm_code_engine_persistent_data_store.code_engine_persistent_data_store &lt;project_id&gt;/&lt;name&gt;
</pre>
