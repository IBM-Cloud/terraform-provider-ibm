---
layout: "ibm"
page_title: "IBM : ibm_logs_extension"
description: |-
  Get information about Extension metadata
subcategory: "Cloud Logs"
---

# ibm_logs_extension

Provides a read-only data source to retrieve information about an Extension metadata. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

### Get details of IBMCloudKubernetes extension
```hcl
data "ibm_logs_extension" "logs_extension" {
  instance_id       = ibm_resource_instance.logs_instance.guid
  region            = ibm_resource_instance.logs_instance.location
  logs_extension_id = "IBMCloudKubernetes"
}

```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String) Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.
* `endpoint_type` - (Optional, String) Cloud Logs Instance Endpoint type. Allowed values `public` and `private`.
* `logs_extension_id` - (Required, Forces new resource, String) The unique identifier of the extension.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the Extension metadata.
* `changelog` - (List) The of changelog entries made in each version of the Extension.
  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
Nested schema for **changelog**:
	* `description_md` - (String) The description of the changes made in this version, formatted in Markdown for rich text presentation.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}\\r\\n\\t]+$/`.
	* `version` - (String) The version of the Extension this changelog entry refers to.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
* `deployment` - (List) Deployment details of an Extension scoped by extension ID in the path.
Nested schema for **deployment**:
	* `applications` - (List) Applications that the Extension is deployed for. When this is empty, it is applied to all applications.
	  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `id` - (String) The unique identifier of the extension.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `item_ids` - (List) The list of Extension item IDs to deploy.
	  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `1` item.
	* `subsystems` - (List) Subsystems that the Extension is deployed. When this is empty, it is applied to all subsystems.
	  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `version` - (String) The version of the Extension revision to deploy.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
* `deprecation` - (List) Deprecation details of the Extension.
Nested schema for **deprecation**:
	* `reason` - (String) The reason why the element (e.g., an Extension or a version of it) is being deprecated.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}\\r\\n\\t]+$/`.
	* `replacement_extensions` - (List) The list of Extension IDs that serve as replacements for the deprecated element.
	  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
* `keywords` - (List) The list of keywords to enhance search capabilities on the front-end side.
  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
* `name` - (String) The name of the Extension.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
* `revisions` - (List) The list of all revisions of the Extension, each representing a versioned snapshot of the Extension's functionality and appearance.
  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
Nested schema for **revisions**:
	* `description` - (String) The detailed description of what this revision includes, changes made, and any important information users should be aware of.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `excerpt` - (String) The brief summary or excerpt of the Extension's description for quick reference.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `items` - (List) The Extension items included in this revision.
	  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
	Nested schema for **items**:
		* `description` - (String) The detailed description of the Extension item.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
		* `id` - (String) The ID of the Extension item.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
		* `is_mandatory` - (Boolean) A flag to indicate if the Extension item is mandatory or not. Mandatory items must be specified when deploying the Extension.
		* `name` - (String) The name of the Extension item.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
		* `target_domain` - (String) The domain of the Extension item.
		  * Constraints: Allowable values are: `alert_definition`, `alert`, `enrichment`, `rule_group`, `view`, `dashboard`, `events_to_metrics`.
	* `labels` - (List) The list of labels or tags associated with the Extension for front-end categorization and filtering.
	  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `version` - (String) The version identifier for this revision of the Extension.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.

