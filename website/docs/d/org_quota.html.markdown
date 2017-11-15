---
layout: "ibm"
page_title: "IBM: ibm_org_quota"
sidebar_current: "docs-ibm-datasource-org-quota"
description: |-
  Get information about an IBM Bluemix organization quota.
---

# ibm\_org_quota

Import the details of an existing IBM Bluemix org quota as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_org_quota" "orgquotadata" {
  name = "quotaname"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the Bluemix organization quota. You can retrieve the value by running the `bx cf quotas` command in the [Bluemix CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the organization.
* `app_instance_limit` - Defines the total app instance limit for organization.
* `app_tasks_limit` - Defines the total app task limit for organization.
* `instance_memory_limit` - Defines the  total instance memory limit for organization.
* `memory_limit` - Defines the total memory limit for organization.
* `non_basic_services_allowed` -  Define non basic services are allowed for organization.
* `total_private_domains` - Defines the total private domain limit for organization.
* `total_reserved_route_ports` - Defines the number of reserved route ports for organization. 
* `total_routes` - Defines the total route for organization.
* `total_service_keys` - Defines the total service keys for organization.
* `total_services` - Defines the total services for organization.
* `trial_db_allowed` - Defines trial db are allowed for organization.
