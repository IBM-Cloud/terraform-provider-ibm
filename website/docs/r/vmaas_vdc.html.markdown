---
layout: "ibm"
page_title: "IBM : ibm_vmaas_vdc"
description: |-
  Manages vmaas_vdc.
subcategory: "VMware as a Service API"
---

# ibm_vmaas_vdc

Create, update, and delete vmaas_vdcs with this resource.

## Example Usage

```hcl
resource "ibm_vmaas_vdc" "vmaas_vdc_instance" {
  accept_language = "en-us"
  director_site {
		id = "id"
		pvdc {
			id = "pvdc_id"
			provider_type {
				name = "paygo"
			}
		}
		url = "url"
  }
  name = "sampleVDC"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `accept_language` - (Optional, String) Language.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-,;=\\.\\*\\s]{1,256}$/`.
* `cpu` - (Optional, Integer) The vCPU usage limit on the virtual data center (VDC). Supported for VDCs deployed on a multitenant Cloud Director site. This property is applicable when the resource pool type is reserved.
  * Constraints: The maximum value is `2000`. The minimum value is `0`.
* `director_site` - (Required, List) The Cloud Director site in which to deploy the virtual data center (VDC).
Nested schema for **director_site**:
	* `id` - (Required, String) A unique ID for the Cloud Director site.
	* `pvdc` - (Required, List) The resource pool within the Director Site in which to deploy the virtual data center (VDC).
	Nested schema for **pvdc**:
		* `id` - (Required, String) A unique ID for the resource pool.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_-]{1,128}$/`.
		* `provider_type` - (Optional, List) Determines how resources are made available to the virtual data center (VDC). Required for VDCs deployed on a multitenant Cloud Director site.
		Nested schema for **provider_type**:
			* `name` - (Required, String) The name of the resource pool type.
			  * Constraints: Allowable values are: `paygo`, `on_demand`, `reserved`.
	* `url` - (Computed, String) The URL of the VMware Cloud Director tenant portal where this virtual data center (VDC) can be managed.
* `fast_provisioning_enabled` - (Optional, Boolean) Determines whether this virtual data center has fast provisioning enabled or not.
* `name` - (Required, String) A human readable ID for the virtual data center (VDC).
* `ram` - (Optional, Integer) The RAM usage limit on the virtual data center (VDC) in GB (1024^3 bytes). Supported for VDCs deployed on a multitenant Cloud Director site. This property is applicable when the resource pool type is reserved.
  * Constraints: The maximum value is `40960`. The minimum value is `0`.
* `rhel_byol` - (Optional, Boolean) Indicates if the RHEL VMs will be using the license from IBM or the customer will use their own license (BYOL).
* `windows_byol` - (Optional, Boolean) Indicates if the Microsoft Windows VMs will be using the license from IBM or the customer will use their own license (BYOL).

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the vmaas_vdc.
* `crn` - (String) A unique ID for the virtual data center (VDC) in IBM Cloud.
* `deleted_at` - (String) The time that the virtual data center (VDC) is deleted.
* `edges` - (List) The VMware NSX-T networking edges deployed on the virtual data center (VDC). NSX-T edges are used for bridging virtualization networking to the physical public-internet and IBM private networking.
  * Constraints: The maximum length is `128` items. The minimum length is `0` items.
Nested schema for **edges**:
	* `id` - (String) A unique ID for the edge.
	* `public_ips` - (List) The public IP addresses assigned to the edge.
	  * Constraints: The maximum length is `256` items. The minimum length is `1` item.
	* `size` - (String) The size of the edge.The size can be specified only for performance edges. Larger sizes require more capacity from the Cloud Director site in which the virtual data center (VDC) was created to be deployed.
	  * Constraints: Allowable values are: `medium`, `large`, `extra_large`.
	* `status` - (String) Determines the state of the edge.
	  * Constraints: Allowable values are: `creating`, `ready_to_use`, `deleting`, `deleted`.
	* `transit_gateways` - (List) Connected IBM Transit Gateways.
	  * Constraints: The maximum length is `128` items. The minimum length is `0` items.
	Nested schema for **transit_gateways**:
		* `connections` - (List) IBM Transit Gateway connections.
		  * Constraints: The maximum length is `128` items. The minimum length is `1` item.
		Nested schema for **connections**:
			* `base_network_type` - (String) The type of the network that the unbound GRE tunnel is targeting. Only "classic" is supported.
			* `local_bgp_asn` - (Integer) Local network BGP ASN for the connection.
			* `local_gateway_ip` - (String) Local gateway IP address for the connection.
			* `local_tunnel_ip` - (String) Local tunnel IP address for the connection.
			* `name` - (String) The autogenerated name for this connection.
			* `network_account_id` - (String) The ID of the account that owns the connected network.
			* `network_type` - (String) The type of the network that is connected through this connection. Only "unbound_gre_tunnel" is supported.
			* `remote_bgp_asn` - (Integer) Remote network BGP ASN for the connection.
			* `remote_gateway_ip` - (String) Remote gateway IP address for the connection.
			* `remote_tunnel_ip` - (String) Remote tunnel IP address for the connection.
			* `status` - (String) Determines the state of the connection.
			  * Constraints: Allowable values are: `pending`, `creating`, `ready_to_use`, `detached`, `deleting`.
			* `transit_gateway_connection_name` - (String) The user-defined name of the connection created on the IBM Transit Gateway.
			* `zone` - (String) The location of the connection.
		* `id` - (String) A unique ID for an IBM Transit Gateway.
		* `status` - (String) Determines the state of the IBM Transit Gateway based on its connections.
		  * Constraints: Allowable values are: `pending`, `creating`, `ready_to_use`, `deleting`.
	* `type` - (String) The type of edge to be deployed.Efficiency edges allow for multiple VDCs to share some edge resources. Performance edges do not share resources between VDCs.
	  * Constraints: Allowable values are: `performance`, `efficiency`.
	* `version` - (String) The edge version.
* `href` - (String) The URL of this virtual data center (VDC).
* `ordered_at` - (String) The time that the virtual data center (VDC) is ordered.
* `org_name` - (String) The name of the VMware Cloud Director organization that contains this virtual data center (VDC). VMware Cloud Director organizations are used to create strong boundaries between VDCs. There is a complete isolation of user administration, networking, workloads, and VMware Cloud Director catalogs between different Director organizations.
* `provisioned_at` - (String) The time that the virtual data center (VDC) is provisioned and available to use.
* `status` - (String) Determines the state of the virtual data center.
  * Constraints: Allowable values are: `creating`, `ready_to_use`, `modifying`, `deleting`, `deleted`, `failed`.
* `status_reasons` - (List) Information about why the request to create the virtual data center (VDC) cannot be completed.
  * Constraints: The maximum length is `128` items. The minimum length is `0` items.
Nested schema for **status_reasons**:
	* `code` - (String) An error code specific to the error encountered.
	  * Constraints: Allowable values are: `insufficent_cpu`, `insufficent_ram`, `insufficent_cpu_and_ram`.
	* `message` - (String) A message that describes why the error ocurred.
	* `more_info` - (String) A URL that links to a page with more information about this error.
* `type` - (String) Determines whether this virtual data center is in a single-tenant or multitenant Cloud Director site.
  * Constraints: Allowable values are: `single_tenant`, `multitenant`.


## Import

You can import the `ibm_vmaas_vdc` resource by using `id`. A unique ID for the virtual data center (VDC).

# Syntax
<pre>
$ terraform import ibm_vmaas_vdc.vmaas_vdc &lt;id&gt;
</pre>
