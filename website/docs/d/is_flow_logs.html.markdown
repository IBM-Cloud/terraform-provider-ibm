---
layout: "ibm"
page_title: "IBM : is_flow_logs"
sidebar_current: "docs-ibm-datasources-is-flow-logs"
description: |-
  Manages IBM Cloud Infrastructure Flow Logs.
---

# ibm\_is_flow_logs

Import the details of an existing IBM Cloud Infrastructure flow logs as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_flow_logs" "ds_flow_logs" {
}

```

## Attribute Reference

The following attributes are exported:

* `flow_log_collectors` - List of all flowlogs in the IBM Cloud Infrastructure.
  * `active` - Indicates whether this collector is active.
  * `created_at` - The date and time flow log was created.
  * `crn` - The CRN for this flow log collector.
  * `href` - The URL for this flow log collector.
  * `id` - . The unique identifier for this flow log collector
  * `lifecycle_state` - The lifecycle state of the flow log collector.
  * `name` - Flow Log Collector name.
  * `resource_group` - The resource group of flow log.
  * `storage_bucket` - The Cloud Object Storage bucket name where the collected flows will be logged.
  * `target` - The target id that the flow log collector is to collect flow logs.
  * `vpc` - The VPC this flow log collector is associated with.  



