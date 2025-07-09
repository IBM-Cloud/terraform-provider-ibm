---
layout: "ibm"
page_title: "IBM : ibm_scc_scope_collection"
description: |-
  Get information about a list of Security and Compliance Center scopes.
subcategory: "Security and Compliance Center"
---

# ibm_scc_scope_collection

Provides a read-only data source to retrieve information about scc_scope. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

~> NOTE: Security Compliance Center is a regional service. Please specify the IBM Cloud Provider attribute `region` to target another region. Else, exporting the environmental variable IBMCLOUD_SCC_API_ENDPOINT will also override which region is being targeted for all ibm providers(ex. `export IBMCLOUD_SCC_API_ENDPOINT=https://eu-es.compliance.cloud.ibm.com`).

## Example Usage

```hcl
data "ibm_scc_scope_collection" "scc_scope_collection" {
  instance_id = "00000000-1111-2222-3333-444444444444"
}
```
## Argument Reference
You can specify the following arguments for this data source.

* `instance_id` - (Required, Forces new resource, String) The ID of the SCC instance in a particular region.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `scopes` - (List) A list of scope that accessible to the user. 
Nested schema for **scopes**:
	* `account_id` - (String) The ID of the IBM account associated with the scope.

	* `attachment_count` - (Integer) The number of `scc_profile_attachment` using the scope.

	* `created_by` - (String) The user who created the scope.
	* Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.:,_\\s]*$/`.

	* `created_on` - (String) The date when the scope was created.

	* `description` - (String) The details of the scope.
	
	* `id` - (String) The scope ID.

	* `instance_id` - (String) The ID of the SCC instance tied to the scope.

	* Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.

	* `name` - (String) The name of the scope.

	* `properties` - (List) The properties of the scope. 
	Nested schema for **properties**:
		* `name` - (String) The name of property.
		* `value` - (String) The value of the property in string form.

	* `updated_by` - (String) The user who updated the control library.
	* Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.:,_\\s]*$/`.

	* `updated_on` - (String) The date when the control library was updated.
