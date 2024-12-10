---
layout: "ibm"
page_title: "IBM : ibm_is_instance_cluster_network_attachment"
description: |-
  Manages InstanceClusterNetworkAttachment.
subcategory: "VPC infrastructure"
---

# ibm_is_instance_cluster_network_attachment

Create, update, and delete InstanceClusterNetworkAttachments with this resource. [About cluster networks](https://cloud.ibm.com/docs/vpc?topic=vpc-about-cluster-network)

## Example Usage

```hcl

resource "ibm_is_instance_action" "is_instance_stop_before" {
	  action = "stop"
	  instance = ibm_is_instance.is_instance.id
}

resource "ibm_is_instance_cluster_network_attachment" "is_instance_cluster_network_attachment_instance" {
	  depends_on = [ibm_is_instance_action.is_instance_stop_before]
    instance_id = ibm_is_instance.is_instance.id
    before {
      id = ibm_is_instance.is_instance.cluster_network_attachments.0.id
    }
    cluster_network_interface {
      name = "my-cluster-network-interface"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
    name = "cna-9"
}
resource "ibm_is_instance_cluster_network_attachment" "is_instance_cluster_network_attachment_instance10" {
	  depends_on = [ibm_is_instance_action.is_instance_stop_before]
    instance_id = ibm_is_instance.is_instance.id
    before {
      id = ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment_instance.instance_cluster_network_attachment_id
    }
    cluster_network_interface {
      name = "my-cluster-network-interface-10"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
    name = "cna-10"
}
resource "ibm_is_instance_cluster_network_attachment" "is_instance_cluster_network_attachment_instance11" {
	  depends_on = [ibm_is_instance_action.is_instance_stop_before]
    instance_id = ibm_is_instance.is_instance.id
    before {
      id = ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment_instance10.instance_cluster_network_attachment_id
}
    cluster_network_interface {
      name = "my-cluster-network-interface-11"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
    name = "cna-11"
}
resource "ibm_is_instance_cluster_network_attachment" "is_instance_cluster_network_attachment_instance12" {
	  depends_on = [ibm_is_instance_action.is_instance_stop_before]
    instance_id = ibm_is_instance.is_instance.id
    before {
      id = ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment_instance11.instance_cluster_network_attachment_id
}
    cluster_network_interface {
      name = "my-cluster-network-interface12"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
    name = "cna-12"
}
resource "ibm_is_instance_cluster_network_attachment" "is_instance_cluster_network_attachment_instance13" {
	  depends_on = [ibm_is_instance_action.is_instance_stop_before]
    instance_id = ibm_is_instance.is_instance.id
    before {
      id = ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment_instance12.instance_cluster_network_attachment_id
    }
    cluster_network_interface {
      name = "my-cluster-network-interface13"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
    name = "cna-13"
}
resource "ibm_is_instance_cluster_network_attachment" "is_instance_cluster_network_attachment_instance14" {
	  depends_on = [ibm_is_instance_action.is_instance_stop_before]
    instance_id = ibm_is_instance.is_instance.id
    before {
      id = ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment_instance13.instance_cluster_network_attachment_id
    }
    cluster_network_interface {
      name = "my-cluster-network-interface14"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
    name = "cna-149"
}
resource "ibm_is_instance_cluster_network_attachment" "is_instance_cluster_network_attachment_instance15" {
	  depends_on = [ibm_is_instance_action.is_instance_stop_before]
    instance_id = ibm_is_instance.is_instance.id
    before {
      id = ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment_instance14.instance_cluster_network_attachment_id
    }
    cluster_network_interface {
      name = "my-cluster-network-interface15"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
    name = "cna-15"
}
resource "ibm_is_instance_cluster_network_attachment" "is_instance_cluster_network_attachment_instance16" {
	  depends_on = [ibm_is_instance_action.is_instance_stop_before]
    instance_id = ibm_is_instance.is_instance.id
    before {
      id = ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment_instance15.instance_cluster_network_attachment_id
    }
    cluster_network_interface {
      name = "my-cluster-network-interface16"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
    name = "cna-16"
}
resource "ibm_is_instance_action" "is_instance_start_after" {
	  # depends_on = [ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment_instance16]
	  action = "start"
	  instance = ibm_is_instance.is_instance.id
}
```

## Argument Reference

You can specify the following arguments for this resource.

  ~> **Note:** 
  **&#x2022;** Instance cluster network attachment creation requires the instance to be in stopped state. Use `ibm_is_instance_action` resource accordingly to stop/start the instance.</br>
  **&#x2022;** Using cluster_network_attachments in `ibm_is_instance` and `ibm_is_instance_cluster_network_attachment` resource together would result in changes shown in both resources alternatively, use either of them or use meta lifecycle argument `ignore_changes` on `ibm_is_instance` resource.</br>


- `before` - (Optional, List) The instance cluster network attachment that is immediately before. If absent, this is thelast instance cluster network attachment.
	Nested schema for **before**:
	- `id` - (Required, String) The unique identifier for this instance cluster network attachment.
- `cluster_network_interface` - (Required, List) The cluster network interface for this instance cluster network attachment.
	
	Nested schema for **cluster_network_interface**:

	- `id` - (Required, String) The unique identifier for this cluster network interface.
	- `name` - (Required, String) The name for this cluster network interface. The name is unique across all interfaces in the cluster network.
	- `primary_ip` - (Required, List) The primary IP for this cluster network interface.
		
		Nested schema for **primary_ip**:
		- `address` - (Required, String) The IP address.If the address is pending allocation, the value will be `0.0.0.0`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
		- `id` - (Required, String) The unique identifier for this cluster network subnet reserved IP.
		- `name` - (Required, String) The name for this cluster network subnet reserved IP. The name is unique across all reserved IPs in a cluster network subnet.
	- `subnet` - (Required, List)
		
		Nested schema for **subnet**:
		- `id` - (Required, String) The unique identifier for this cluster network subnet.

- `instance_id` - (Required, Forces new resource, String) The virtual server instance identifier.
- `name` - (Optional, String) The name for this instance cluster network attachment. The name is unique across all network attachments for the instance.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.


- `id` - The unique identifier of the InstanceClusterNetworkAttachment.
- `before` - (List) The instance cluster network attachment that is immediately before. If absent, this is thelast instance cluster network attachment.
	Nested schema for **before**:
	- `href` - (String) The URL for this instance cluster network attachment.
	- `id` - (String) The unique identifier for this instance cluster network attachment.
	- `name` - (String) The name for this instance cluster network attachment. The name is unique across all network attachments for the instance.
	- `resource_type` - (String) The resource type.
- `cluster_network_interface` - (List) The cluster network interface for this instance cluster network attachment.
	Nested schema for **cluster_network_interface**:
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested schema for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this cluster network interface.
	- `id` - (String) The unique identifier for this cluster network interface.
	- `name` - (String) The name for this cluster network interface. The name is unique across all interfaces in the cluster network.
	- `primary_ip` - (List) The primary IP for this cluster network interface.
		Nested schema for **primary_ip**:
		- `address` - (String) The IP address.If the address is pending allocation, the value will be `0.0.0.0`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.

			Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this cluster network subnet reserved IP.
		- `id` - (String) The unique identifier for this cluster network subnet reserved IP.
		- `name` - (String) The name for this cluster network subnet reserved IP. The name is unique across all reserved IPs in a cluster network subnet.
		- `resource_type` - (String) The resource type.
	- `resource_type` - (String) The resource type.
	- `subnet` - (List)
		Nested schema for **subnet**:
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this cluster network subnet.
		- `id` - (String) The unique identifier for this cluster network subnet.
		- `name` - (String) The name for this cluster network subnet. The name is unique across all cluster network subnets in the cluster network.
		- `resource_type` - (String) The resource type.
- `href` - (String) The URL for this instance cluster network attachment.
- `instance_cluster_network_attachment_id` - (String) The unique identifier for this instance cluster network attachment.
- `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
	Nested schema for **lifecycle_reasons**:
	- `code` - (String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	- `message` - (String) An explanation of the reason for this lifecycle state.
	- `more_info` - (String) Link to documentation about the reason for this lifecycle state.
- `lifecycle_state` - (String) The lifecycle state of the instance cluster network attachment. Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`.
- `name` - (String) The name for this instance cluster network attachment. The name is unique across all network attachments for the instance.
- `resource_type` - (String) The resource type. Allowable values are: `instance_cluster_network_attachment`.


## Import

You can import the `ibm_is_instance_cluster_network_attachment` resource by using `id`.
The `id` property can be formed from `instance_id`, and `instance_cluster_network_attachment_id` in the following format:

<pre>
&lt;instance_id&gt;/&lt;instance_cluster_network_attachment_id&gt;
</pre>
- `instance_id`: A string. The virtual server instance identifier.
- `instance_cluster_network_attachment_id`: A string in the format `0717-fb880975-db45-4459-8548-64e3995ac213`. The unique identifier for this instance cluster network attachment.

# Syntax
<pre>
$ terraform import ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment &lt;instance_id&gt;/&lt;instance_cluster_network_attachment_id&gt;
</pre>
