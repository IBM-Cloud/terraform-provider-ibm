---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM: ibm_org_quota"
description: |-
  Get information about an IBM Cloud organization quota.
---

# `ibm_org_quota`

Retrieve information about a quota for a Cloud Foundry organization. For more information, about organization and usage of quote, see [Updating orgs and spaces](https://cloud.ibm.com/docs/account?topic=account-orgupdates).


## Example usage
The following example retrieves information for an existing quota plan. 


```
data "ibm_org_quota" "orgquotadata" {
  name = "quotaname"
}
```

## Argument reference
Review the input parameters that you can specify for your data source. 

- `name` - (Required, String) The name of the quota plan for the Cloud Foundry organization. You can retrieve the value by running the `ibmcloud cf quotas` command in the IBM Cloud CLI.


## Attribute referenceReview the output parameters that you can access after you retrieved your data source. 

- `app_instance_limit`- (Integer) Defines the total number of app instances that are allowed for the Cloud Foundry organization.
- `app_tasks_limit`- (Integer) Defines the total number of app tasks for a Cloud Foundry organization.
- `id` - (String) The unique identifier of the Cloud Foundry organization.
- `instance_memory_limit`- (Integer) Defines the total amount of memory that an instance can use in a Cloud Foundry organization.
- `memory_limit`- (Integer) Defines the total amount of memory that can be used by the Cloud Foundry organization.
- `non_basic_services_allowed`- (Boolean) Defines if non-basic (paid) services are allowed for the Cloud Foundry organization.
- `total_private_domains`- (Integer) Defines the total number of private domains that can be created for a Cloud Foundry organization.
- `total_reserved_route_ports`- (Integer) Defines the number of routes with reserved ports for the Cloud Foundry organization.
- `total_routes`- (Integer) Defines the maximum number of routes that can be created for a Cloud Foundry organization.
- `total_service_keys`- (Integer) Defines the maximum number of service keys for the Cloud Foundry organization.
- `total_services`- (Integer) Defines the maximum number of services that can be created in a Cloud Foundry organization.
- `trial_db_allowed`- (Boolean) Defines if a trial database is allowed for the Cloud Foundry organization.

