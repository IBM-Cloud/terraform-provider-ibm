---
layout: "ibm"
page_title: "IBM: ibm_org_quota"
sidebar_current: "docs-ibm-datasource-org-quota"
description: |-
  Get information about an IBM Cloud organization quota.
---

# ibm\_org_quota

Import the details of an existing quota for an IBM Cloud org as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_org_quota" "orgquotadata" {
  name = "quotaname"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the quota plan for the IBM Cloud org. You can retrieve the value by running the `ibmcloud cf quotas` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the organization.
* `app_instance_limit` - Defines the total number of app instances that are allowed for the organization.
* `app_tasks_limit` - Defines the total app task limit for the organization.
* `instance_memory_limit` - Defines the total memory usage that is allowed per instance for the organization.
* `memory_limit` - Defines the total memory usage that is allowed for the organization.
* `non_basic_services_allowed` - Defines if non-basic (paid) services are allowed for the organization.
* `total_private_domains` - Defines the total private domain limit for the organization.
* `total_reserved_route_ports` - Defines the number of routes with reserved ports for the organization. 
* `total_routes` - Defines the maximum total routes for the organization.
* `total_service_keys` - Defines the maximum number of service keys for the organization.
* `total_services` - Defines the maximum number of services for organization.
* `trial_db_allowed` - Defines if a trial db is allowed for the organization.
