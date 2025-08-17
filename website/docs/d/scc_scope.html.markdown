---
layout: "ibm"
page_title: "IBM : ibm_scc_scope"
description: |-
  Get information about a Security and Compliance Center scope.
subcategory: "Security and Compliance Center"
---

# ibm_scc_scope

Provides a read-only data source to retrieve information about scc_scope. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

~> NOTE: Security Compliance Center is a regional service. Please specify the IBM Cloud Provider attribute `region` to target another region. Else, exporting the environmental variable IBMCLOUD_SCC_API_ENDPOINT will also override which region is being targeted for all ibm providers(ex. `export IBMCLOUD_SCC_API_ENDPOINT=https://eu-es.compliance.cloud.ibm.com`).

## Example Usage

```hcl
data "ibm_scc_scope" "scc_scope" {
  instance_id = "00000000-1111-2222-3333-444444444444"
  scope_id = "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
}
```
## Argument Reference
You can specify the following arguments for this data source.

* `instance_id` - (Required, Forces new resource, String) The ID of the SCC instance in a particular region.
* `scope_id` - (Required, Forces new resource, String) THe ID of the SCC instance scope in a particular region.

## Attribute Reference

After your data source is created, you can read values from the following attributes.
* `account_id` - (String) The ID of the IBM account associated with the scope.

* `attachment_count` - (Integer) The number of `scc_profile_attachment` using the scope.

* `created_by` - (String) The user who created the scope.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.:,_\\s]*$/`.

* `created_on` - (String) The date when the scope was created.

* `description` - (String) The details of the scope

* `exclusions` - (List) A list of excluded targets from the scope.
Nested schema for **exclusions**:
	* `account_id` - (String) The ID of the account that was excluded.
	* `resource_group_id` - (String) The ID of the resource group that was excluded.
	* `account_group_id` - (String) The ID of the account group in an enterprise that was excluded.

* `id` - (String) The scope ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.

* `name` - (String) The name of the scope.

* `properties` - (List) The properties of the scope. 
Nested schema for **properties**:
	* `account_group_id` - (String) The ID of the account group in an enterprise.
	* `account_id` - (String) The ID of the account.
	* `enterprise_id` - (String) The ID of the enterprise.
	* `ibm_facts_api_instance_id` - (String) The ID of the `scc_provider_type_instance` that is a provider type `ibm_cloud_facts_api`. 
	* `resource_group_id` - (String) The ID of the resource group tied to an account. 

* `updated_by` - (String) The user who updated the control library.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.:,_\\s]*$/`.

* `updated_on` - (String) The date when the control library was updated.
