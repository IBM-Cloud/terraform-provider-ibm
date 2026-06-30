---
layout: "ibm"
page_title: "IBM : ibm_config_aggregator_batch_configurations"
description: |-
  Get information about config_aggregator_batch_configurations
subcategory: "Configuration Aggregator"
---

# ibm_config_aggregator_batch_configurations

Provides a read-only data source to retrieve information about config_aggregator_batch_configurations. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_config_aggregator_batch_configurations" "config_aggregator_batch_configurations" {
  instance_id = "b76b5724-dfe9-45ca-955d-1355b03e9b11"

  config {
    resource_crn = "crn:v1:bluemix:public:iam-am::a/a5dcbaf5d8324661ad99ecafb3e1cb4a::policy:047cf702-998c-461a-8976-bf3361cc8115"
  }

  config {
    resource_crn = "crn:v1:bluemix:public:iam-am::a/a5dcbaf5d8324661ad99ecafb3e1cb4a::policy:047cf702-998c-461a-8976-bf3361cc8116"
    service_name = "iam-access-management"
  }

  config {
    resource_crn = "crn:v1:bluemix:public:iam-am::a/a5dcbaf5d8324661ad99ecafb3e1cb4a::policy:047cf702-998c-461a-8976-bf3361cc8116"
    service_name = "iam-access-management"
    config_type  = ["policy"]
  }

  config {
    resource_crn = "crn:v1:bluemix:public:iam-am::a/a5dcbaf5d8324661ad99ecafb3e1cb4a::policy:047cf702-998c-461a-8976-bf3361cc8116"
    type_id      = "047cf702-998c-461a-8976-bf3361cc8116"
  }
}

```

## Argument Reference

You can specify the following arguments for this data source.

* `configs` - (Required, List) The List of resources requested as part of List Configs Query API.
  * Constraints: The maximum length is `20` items. The minimum length is `0` items.
Nested schema for **configs**:
	* `config_type` - (Optional, List) The list of type for resource configuration that are to be retrieved.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9 ,\\-_]+$/`. The maximum length is `20` items. The minimum length is `0` items.
	* `resource_crn` - (Required, String) The unique CRN of the IBM Cloud resource.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9.\\:\/-]+$/`.
	* `service_name` - (Optional, String) The name of the service to which the resource belongs.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9.\\:\/-]+$/`.
	* `type_id` - (Optional, String) The unique identifier for each IBM resource.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9.\\:\/-]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

---

### `id`
(String) The unique identifier of the config_aggregator_batch_configurations.

---

### `configs`
(List) Array of resource configurations returned by the API.

#### Nested schema for `configs`:

* `about` - (String)  
  JSON string containing metadata about the resource.

* `config` - (String)  
  JSON string containing the configuration details of the resource.

* `config_v2` - (String)  
  JSON string representing version 2 of the configuration (if available). May be empty.

---

### `errors`
(List) List of resources that failed to fetch.

#### Nested schema for `errors`:

* `resource_crn` - (String) The CRN of the resource that failed.
* `message` - (String) The error message returned by the API.
* `error_code` - (String) The error code associated with the failure.

---

### `prev`
(List) The reference to the previous page of entries. If absent, this is the first page.

#### Nested schema for `prev`:

* `href` - (String) The reference to the previous page of entries.
* `start` - (String) The starting token for pagination.

---

## xNotes

- The `configs` attribute contains **JSON-encoded strings**, not structured maps.
- The `errors` attribute is populated when one or more requested resources fail.
- Partial success is supported:
  - Successful responses → `configs`
  - Failed resources → `errors`

