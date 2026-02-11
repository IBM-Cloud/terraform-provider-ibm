---
layout: "ibm"
page_title: "IBM : ibm_code_engine_persistent_data_store"
description: |-
  Get information about code_engine_persistent_data_store
subcategory: "Code Engine"
---

# ibm_code_engine_persistent_data_store

Provides a read-only data source to retrieve information about a code_engine_persistent_data_store. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_code_engine_persistent_data_store" "code_engine_persistent_data_store" {
	name = ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance.name
	project_id = ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance.project_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `name` - (Required, Forces new resource, String) The name of your persistent data store.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the code_engine_persistent_data_store.
* `persistent_data_store_id` - (String) The identifier of the resource.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `created_at` - (String) The timestamp when the resource was created.
* `data` - (List) Data container that allows to specify config parameters and their values as a key-value map. Each key field must consist of alphanumeric characters, `-`, `_` or `.` and must not exceed a max length of 253 characters. Each value field can consists of any character and must not exceed a max length of 1048576 characters.
Nested schema for **data**:
	* `bucket_location` - (String) Specify the location of the bucket.
	  * Constraints: Allowable values are: `au-syd`, `br-sao`, `ca-mon`, `ca-tor`, `eu-de`, `eu-es`, `eu-gb`, `jp-osa`, `jp-tok`, `us-east`, `us-south`, `ap`, `eu`, `us`, `ams03`, `che01`, `mil01`, `mon01`, `par01`, `sjc04`, `sng01`. The value must match regular expression `/^(au-syd|br-sao|ca-mon|ca-tor|eu-de|eu-es|eu-gb|jp-osa|jp-tok|us-east|us-south|ap|eu|us|ams03|che01|mil01|mon01|par01|sjc04|sng01)$/`.
	* `bucket_name` - (String) Specify the name of the bucket.
	  * Constraints: The maximum length is `63` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-z0-9]([-a-z0-9]*[a-z0-9])?$/`.
	* `secret_name` - (String) Specify the name of the HMAC secret.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
* `entity_tag` - (String) The version of the persistent data store, which is used to achieve optimistic locking.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[\\*\\-a-z0-9]+$/`.
* `region` - (String) The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'.
* `storage_type` - (String) Specify the storage type of the persistent data store.
  * Constraints: Allowable values are: `object_storage`. The value must match regular expression `/^(object_storage)$/`.

