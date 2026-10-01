---
layout: "ibm"
page_title: "IBM : ibm_pha_deployment"
description: |-
  Manages pha_deployment.
subcategory: "HA and DR Automation for IBM® Power® Virtual Server API reference"
---

# ibm_pha_deployment

Create PHA deployments with this resource in PowerHA.

## Example Usage

```hcl
when cluster_type is standard only primary details are allowed
resource "ibm_pha_deployment" "pha_deployment_instance" {
  accept_language = "en-US"
  if_none_match = "abcdef"
  pha_instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
  location_id = "us-south"
  primary_workspace = "saiworkspace1"
  cluster_type = "standard"
  api_key = "xxxx_asnj_cdcdw_csdcdcwcwewdwe_cwsw"
  primary_cluster_nodes = ["8eefautr-4c02-0009-0086-vdvds_cdvcesdvcscv"]
}
```

```hcl
when cluster_type is linked only standby details are also allowed
resource "ibm_pha_deployment" "pha_deployment_instance" {
  accept_language = "en-US"
  if_none_match = "abcdef"
  pha_instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
  location_id = "us-south"
  primary_workspace = "saiworkspace1"
  secondary_location_id = "us-east"
  secondary_workspace = "saiworkspace2"
  cluster_type = "linked"
  api_key = "xxxx_asnj_cdcdw_csdcdcwcwewdwe_cwsw"
  primary_cluster_nodes = ["8eefautr-4c02-0009-0086-vdvds_cdvcesdvcscv"]
  secondary_cluster_nodes = ["8eefautr-4c02-0009-0086-vdvds_cdvcesdvcscv"]
}
```
### Path Parameters

* `pha_instance_id` - (Required, Forces new resource, String) The unique identifier of the powerha service instance.
  * Constraints: The maximum length is `1048` characters. The minimum length is `36` characters. The value must match regular expression `/^.*$/`.

### Query Parameters

* `accept_language` - (Optional, Forces new resource, String) The language in which the response should be returned.
* `accepts_incomplete` - (Optional, Forces new resource, Boolean) Indicates whether the request can be accepted before the operation is complete. 
* `if_none_match` - (Optional, Forces new resource, String) ETag for conditional requests (optional).
  * Constraints: The maximum length is `50` characters. The minimum length is `6` characters. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.

## Argument Reference

You can specify the following arguments for this resource.

* `primary_workspace` - (Required, Forces new resource, String) Primary workspace identifier.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\-]+$/`.
* `secondary_workspace` - (Optional, Forces new resource, String) Secondary workspace identifier.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\-]+$/`.
* `api_key` - (Required, Forces new resource, String) The API key associated with the request.
 * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\-]+$/`.
* `cluster_type` - (Optional, Forces new resource, String) Type of PowerHA cluster being deployed.
  * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^.*$/`.
* `location_id` - (Required, Forces new resource, String) Identifier for the primary deployment   location.
  * Constraints: The maximum length is `32` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-z0-9\-]+$/`.
* `secondary_location_id` - (Optional, Forces new resource, String) Identifier for the secondary deployment location.
  * Constraints: The maximum length is `32` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-z0-9\-]+$/`.
* `configure_type` - (Optional, Forces new resource, String) Configuration type for the deployment.
  * Constraints: The value must be one of `automated` or `existing`.
* `primary_cluster_nodes` - (Optional, Forces new resource, List) List of primary cluster node VM IDs.
  * Constraints: The maximum length is `100` items. The minimum length is `1` item.
  * Each item must be exactly `36` characters and match regular expression `/^[0-9a-fA-F\-]{36}$/`.
* `secondary_cluster_nodes` - (Optional, Forces new resource, List) List of secondary cluster node VM IDs.
  * Constraints: The maximum length is `100` items. The minimum length is `1` item.
  * Each item must be exactly `36` characters and match regular expression `/^[0-9a-fA-F\-]{36}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the pha_deployment.
* `cloud_account_id` - (String) Cloud account identifier.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `connectivity_type` - (String) Type of network connectivity.
  * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `creation_time` - (String) Timestamp expressing creationtime.
  * Constraints: The maximum length is `2048` characters. The minimum length is `20` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `deprovision_time` - (String) Timestamp expressing deprovision time.
  * Constraints: The maximum length is `2048` characters. The minimum length is `20` characters. The value must match regular expression `/^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}Z$/`.
* `guid` - (String) Global unique identifier.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `href` - (String) Resource reference.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `pha_instance_id` - (String) Provision request identifier.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `is_duplicate` - (Boolean) Indicates whether deployment is duplicate.
* `plan_id` - (String) Identifier for the service plan.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `plan_name` - (String) Name of service plan.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._: -]+$/`.
* `powerha_cluster_name` - (String) Name of the PowerHA cluster.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `powerha_cluster_type` - (String) Type of PowerHA cluster.
  * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `powerha_level` - (String) PowerHA version level.
  * Constraints: The maximum length is `20` characters. The minimum length is `5` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.

* `primary_cluster_nodes_details` - (List) List of primary cluste nodes.
  * Constraints: The maximum length is `8` items. The minimum length is `0` items.
Nested schema for **primary_cluster_nodes_details**:
    * `agent_status` - (String) Status of the PHA agent running on the node.
      * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
    * `cores` - (Float) Number of CPU cores allocated to the node.
      * Constraints: The maximum value is `16`. The minimum value is `1`.
    * `ip_address` - (String) IP address assigned to the virtual machine.
      * Constraints: The maximum length is `1048` characters. The minimum length is `7` characters. The value must match regular expression `/^[0-9.:]+$/`.
    * `memory` - (Integer) Memory allocated to the virtual machine in MB.
      * Constraints: The maximum value is `1048576`. The minimum value is `1`.
    * `pha_level` - (String) PowerHA version level installed on the node.
      * Constraints: The maximum length is `20` characters. The minimum length is `5` characters. The value must match regular expression `/^[0-9.]+$/`.
    * `region` - (String) Region where the virtual machine is deployed.
      * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
    * `ui_server_url` - (String) URL of the PowerHA UI server associated with the node.
      * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
    * `vm_id` - (String) Identifier of the Power virtual server instance.
      * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
    * `vm_name` - (String) Name of the virtual machine.
      * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
    * `vm_status` - (String) Current operational status of the virtual machine.
      * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
    * `workspace_id` - (String) Workspace identifier associated with the node.
      * Constraints: Length must be `36` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `primary_location` - (String) Primary cluster location.
  * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `primary_region_name` - (String) name of the primary workspace region.
  * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `primary_workspace_crn` - (String) CRN of the primary workspace.
  * Constraints: The maximum length is `2048` characters. The minimum length is `20` characters. The value must match regular expression `/^crn:v1:[A-Za-z0-9._:-]+(:[A-Za-z0-9._:-]+)*$/`.
* `primary_workspace_name` - (String) name of the primary workspace.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `provision_end_time` - (String) Time stamp provisioning completed.
  * Constraints: The maximum length is `2048` characters. The minimum length is `20` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `provision_start_time` - (String) Time stamp provisioning started.
  * Constraints: The maximum length is `2048` characters. The minimum length is `20` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `provision_status` - (String) Current provision status.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `region_id` - (String) Deployment region identifier.
  * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `secondary_location` - (String) Secondary cluster location.
  * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.  
* `resource_group` - (String) Name of the resource group.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `resource_group_crn` - (String) CRN of associated resource group.
  * Constraints: The maximum length is `2048` characters. The minimum length is `20` characters. The value must match regular expression `/^[A-Za-z0-9._:\/-]+$/`.
* `resource_instance` - (String) Resource instance identifier.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.

* `secondary_cluster_nodes_details` - (List) List of secondary cluster nodes.
  * Constraints: The maximum length is `8` items. The minimum length is `0` items.
Nested schema for **secondary_cluster_nodes_details**:
    * `agent_status` - (String) Status of the PHA agent running on the node.
      * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
    * `cores` - (Float) Number of CPU cores allocated to the node.
      * Constraints: The maximum value is `16`. The minimum value is `1`.
    * `ip_address` - (String) IP address assigned to the virtual machine.
      * Constraints: The maximum length is `1048` characters. The minimum length is `7` characters. The value must match regular expression `/^[0-9.:]+$/`.
    * `memory` - (Integer) Memory allocated to the virtual machine in MB.
      * Constraints: The maximum value is `1048576`. The minimum value is `1`.
    * `pha_level` - (String) PowerHA version level installed on the node.
      * Constraints: The maximum length is `20` characters. The minimum length is `5` characters. The value must match regular expression `/^[0-9.]+$/`.
    * `region` - (String) Region where the virtual machine is deployed.
      * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
    * `ui_server_url` - (String) URL of the PowerHA UI server associated with the node.
      * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
    * `vm_id` - (String) Identifier of the Power virtual server instance.
      * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
    * `vm_name` - (String) Name of the virtual machine.
      * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
    * `vm_status` - (String) Current operational status of the virtual machine.
      * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
    * `workspace_id` - (String) Workspace identifier associated with the node.
      * Constraints: Length must be `36` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `secondary_workspace_crn` - (String) CRN of the secondary workspace.
  * Constraints: The maximum length is `2048` characters. The minimum length is `20` characters. The value must match regular expression `/^crn:v1:[A-Za-z0-9._:-]+(:[A-Za-z0-9._:-]+)*$/`.
* `service_description` - (String) Description of provisioned service.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:,\\- ]+$/`.
* `service_id` - (String) Identifier for the service.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `service_name` - (String) Name of service.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._: -]+$/`.
* `standby_region_name` - (String) name of the standby worksapce region.
  * Constraints: The maximum length is `1048` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `standby_workspace_name` - (String) name of the standby worksapce.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `user_tags` - (String) User defined tags.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:,-]+$/`.

* `etag` - ETag identifier for pha_deployment.

## Import

You can import the `ibm_pha_deployment` resource by using `id`.
The `id` property can be formed from `pha_instance_id`, and `pha_instance_id` in the following format:

<pre>
&lt;pha_instance_id&gt;/&lt;pha_instance_id&gt;
</pre>
* `pha_instance_id`: A string in the format `8eefautr-4c02-0009-0086-8bd4d8cf61b6`. Unique identifier of the provisioned instance.
* `pha_instance_id`: A string in the format `8eefab28-4c02-4-0017-8bd4d8c`. Provision request identifier.

# Syntax
<pre>
$ terraform import ibm_pha_deployment.pha_deployment &lt;pha_instance_id&gt;/&lt;pha_instance_id&gt;
</pre>
