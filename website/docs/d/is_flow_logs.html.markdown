---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_flow_logs"
description: |-
  Manages IBM Cloud Infrastructure Flow Logs.
---

# ibm_is_flow_logs
Retrieve an information of an existing IBM Cloud Infrastructure flow logs as a read-only data source. For more information, about VPC flow log, see [creating a flow log collector](https://cloud.ibm.com/docs/vpc?topic=vpc-ordering-flow-log-collector).


## Example usage

```terraform

data "ibm_is_flow_logs" "ds_flow_logs" {
}

```

## Attribute reference
Review the attribute references that you can access after you retrieve your data source. 

- `flow_log_collectors` - (List) Lists all the flow logs in the IBM Cloud.

  Nested scheme for `flow_log_collectors`:
	- `active` - (String) Indicates whether the collector is active.
	- `created_at` - (Timestamp) The date and time the flow log created.
	- `crn` - (String) The CRN of the flow log collector.
	- `href` - (String) The URL of the flow log collector.
	- `id` - (String) The unique identifier of the flow log collector.
	- `lifecycle_state` - (String) The lifecycle state of the flow log collector.
	- `name` - (String) The flow log collector name.
	- `resource_group` - (String) The resource group of the flow log.
	- `storage_bucket` - (String) The IBM Cloud Object Storage bucket name where the flow logs are logged.
	- `target` - (String) The target ID that the flow log collector collects the flow logs.
	- `vpc` - (String) The VPC of the flow log collector that are associated.



