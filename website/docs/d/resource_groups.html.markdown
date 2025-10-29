---
subcategory: "Resource management"
layout: "ibm"
page_title: "IBM: ibm_resource_groups"
description: |-
  Get information about IBM resource groups.
---

# ibm_resource_groups
Retrieve information about existing IBM resource groups as a read-only data source. For more information, about resource groups, see [managing resource groups](https://cloud.ibm.com/docs/account?topic=account-rgs).

## Example usage
The following example retrieves all resource groups in the account.

```terraform
data "ibm_resource_groups" "groups" {
}
```

### Example to retrieve resource groups by name

```terraform
data "ibm_resource_groups" "groups" {
  name = "test"
}
```

### Example to retrieve default resource groups

```terraform
data "ibm_resource_groups" "groups" {
  is_default = true
}
```

### Example to include deleted resource groups

```terraform
data "ibm_resource_groups" "groups" {
  include_deleted = true
}
```

### Example to filter resource groups by date

```terraform
data "ibm_resource_groups" "groups" {
  date = "2024-01"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `date` - (Optional, String) The date in YYYY-MM format to filter resource groups. Deleted resource groups are excluded before this month.
- `include_deleted` - (Optional, Bool) Specifies whether to include deleted resource groups in the results.
- `is_default` - (Optional, Bool) Specifies whether to filter for default resource groups. 
- `name` - (Optional, String) The name of an IBM Cloud resource group to filter by. You can retrieve the value by running the `ibmcloud resource groups` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `id` - (String) The account ID used to filter the resource groups.
- `resource_groups` - (List of Objects) A list of resource groups matching the specified filters. Each resource group in the list has the following attributes:
  - `account_id` - (String) The account ID that the resource group belongs to.
  - `crn` - (String) The full CRN (Cloud Resource Name) associated with the resource group.
  - `created_at` - (String) The date and time when the resource group was initially created.
  - `id` - (String) The unique identifier of the resource group.
  - `is_default` - (Bool) Indicates whether this is the default resource group for the account.
  - `name` - (String) The human-readable name of the resource group.
  - `payment_methods_url` - (String) The URL to access the payment methods details associated with the resource group.
  - `quota_id` - (String) An alpha-numeric value identifying the quota ID associated with the resource group.
  - `quota_url` - (String) The URL to access the quota details associated with the resource group.
  - `resource_linkages` - (List of Strings) An array of resource CRNs that are linked to the resource group.
  - `state` - (String) The state of the resource group (e.g., ACTIVE, SUSPENDED).
  - `teams_url` - (String) The URL to access the team details associated with the resource group.
  - `updated_at` - (String) The date and time when the resource group was last updated.

## Usage notes

- When no filters are specified, the data source returns all resource groups in the account.
- The `date` parameter is useful for retrieving historical resource group information and excludes resource groups deleted before the specified month.
- The `include_deleted` parameter allows you to retrieve resource groups that have been soft-deleted from the account.
- The data source returns a list of resource groups, even when filtering by name or default status, as these filters could potentially match multiple resource groups.
