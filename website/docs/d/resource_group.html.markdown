---

subcategory: "Resource management"
layout: "ibm"
page_title: "IBM: ibm_resource_group"
description: |-
  Get information about an IBM resource group.
---

# ibm_resource_group
Retrieve information about an existing IBM resource group as a read-only data source. For more information, about resource group, see [managing resource groups](https://cloud.ibm.com/docs/account?topic=account-rgs).

## Example usage
The following example enables you to import the resource group by name.

```terraform
data "ibm_resource_group" "group" {
  name = "test"
}
```

### Example to import the default resource group

```terraform
data "ibm_resource_group" "group" {
  is_default = "true"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `is_default` - (Optional, Bool) Specifies whether you want to import default resource group.  **Note**: Conflicts with the  `name`.
- `name` - (Optional, String) The name of an IBM Cloud resource group. You can retrieve the value by running the `ibmcloud resource groups` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).  **Note**: Conflicts with `is_default`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `account_id` - (String) Account ID.  
- `crn` - (String) The full CRN associated with the resource group.
- `created_at` - (Timestamp) The date when the resource group initially created.
- `id` - (String) The unique identifier of the new resource group.
- `payment_methods_url` - (String) The URL to access the payment methods details that is associated with the resource group.
- `quota_url` - (String) The URL to access the quota details that is associated with the resource group.
- `quota_id` - (String) An alpha-numeric value identifying the quota ID associated with the resource group.
- `resource_linkages` - (String) An array of the resources that is linked to the resource group.
- `state` - (String) The state of the resource group.
- `teams_url` -  (String) The URL to access the team details that is associated with the resource group.
- `updated_at` - (Timestamp) The date when the resource group last updated.

