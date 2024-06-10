---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: instance_templates"
description: |-
  Retrives all the IBM VPC instance templates.
---

# ibm_is_instance_templates
Retrieve information of an existing IBM VPC instance templates. For more information, about VPC instance templates, see [creating an instance template](https://cloud.ibm.com/docs/vpc?topic=vpc-create-instance-template).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
In the following example, you can get information of list of instance templates of VPC Generation-2 infrastructure.

```terraform	
data "ibm_is_instance_templates" "example" {	   
}

```

## Attribute reference
You can access the following attribute references after your data source is created. 

- `templates` - (List of Objects) List of templates.
	- `availability_policy_host_failure` - (String) The availability policy for this virtual server instance. The action to perform if the compute host experiences a failure. 
	- `boot_volume` - (List) A nested block describes the boot volume configuration for the template.

	  Nested scheme for `boot_volume`:
		- `delete_volume_on_instance_delete` - (String) You can configure to delete the boot volume based on instance deletion.
		- `encryption` - (String) The encryption key CRN such as HPCS, Key Protect, etc., is provided to encrypt the boot volume attached.
		- `iops` - (String) The IOPS for the boot volume.
		- `name` - (String) The name of the boot volume.
		- `profile` - (String) The profile for the boot volume configuration.
		- `size` - (String) The boot volume size to configure in giga bytes.
		- `tags` - (String) User Tags associated with the volume. (https://cloud.ibm.com/apidocs/tagging#types-of-tags).

	- `catalog_offering` - (List) The [catalog](https://cloud.ibm.com/docs/account?topic=account-restrict-by-user&interface=ui) offering or offering version to use when provisioning this virtual server instance. If an offering is specified, the latest version of that offering will be used. The specified offering or offering version may be in a different account in the same [enterprise](https://cloud.ibm.com/docs/account?topic=account-what-is-enterprise), subject to IAM policies.

		Nested scheme for `catalog_offering`:

		- `offering_crn` - (String) The CRN for this catalog offering. Identifies a catalog offering by this unique property
		- `version_crn` - (String) The CRN for this version of a catalog offering. Identifies a version of a catalog offering by this unique property
		- `plan_crn` - (String) The CRN for this catalog offering version's billing plan
			
	- `crn` - (String) The CRN of the instance template.
	- `default_trusted_profile_auto_link` - (Boolean) If set to `true`, the system will create a link to the specified `target` trusted profile during instance creation. Regardless of whether a link is created by the system or manually using the IAM Identity service, it will be automatically deleted when the instance is deleted. Default is true. 
	- `default_trusted_profile_target` - (String) The unique identifier or CRN of the default IAM trusted profile to use for this virtual server instance.
	- `href` - (String) The URL of the instance template.
	- `id` - (String) The ID of the instance template.
	- `image` - (String) The ID of the image to create the template.
	- `keys` - (String) List of SSH key IDs used to allow log in user to the instances.
	- `metadata_service_enabled` - (Boolean) Indicates whether the metadata service endpoint is available to the virtual server instance.
	
	~> **NOTE**
	`metadata_service_enabled` is deprecated and will be removed in the future. Refer `metadata_service` instead
	- `metadata_service` - (List) The metadata service configuration. 

       Nested scheme for `metadata_service`:
       - `enabled` - (Boolean) Indicates whether the metadata service endpoint will be available to the virtual server instance.
       - `protocol` - (String) The communication protocol to use for the metadata service endpoint.
       - `response_hop_limit` - (Integer) The hop limit (IP time to live) for IP response packets from the metadata service.
       
	- `name` - (String) The name of the instance template.
	- `network_attachments` - (List) The additional network attachments to create for the virtual server instance.
		Nested schema for **network_attachments**:
		- `name` - (String) The name for this network attachment. Names must be unique within the instance the network attachment resides in. If unspecified, the name will be a hyphenated list of randomly-selected words.
		- `virtual_network_interface` - (List) A virtual network interface for the instance network attachment. This can be specifiedusing an existing virtual network interface, or a prototype object for a new virtualnetwork interface.If an existing virtual network interface is specified, `enable_infrastructure_nat` must be`false`.
			Nested schema for **virtual_network_interface**:
			- `allow_ip_spoofing` - (Boolean) Indicates whether source IP spoofing is allowed on this interface. If `false`, source IP spoofing is prevented on this interface. If `true`, source IP spoofing is allowed on this interface.
			- `auto_delete` - (Boolean) Indicates whether this virtual network interface will be automatically deleted when`target` is deleted.
			- `crn` - (String) The CRN for this virtual network interface.
			- `enable_infrastructure_nat` - (Boolean) If `true`:- The VPC infrastructure performs any needed NAT operations.- `floating_ips` must not have more than one floating IP.If `false`:- Packets are passed unchanged to/from the virtual network interface,  allowing the workload to perform any needed NAT operations.- `allow_ip_spoofing` must be `false`.- If the virtual network interface is attached:  - The target `resource_type` must be `bare_metal_server_network_attachment`.  - The target `interface_type` must not be `hipersocket`.
			- `href` - (String) The URL for this virtual network interface.
			- `id` - (String) The unique identifier for this virtual network interface.
			- `ips` - (List) Additional IP addresses to bind to the virtual network interface. Each item may be either a reserved IP identity, or as a reserved IP prototype object which will be used to create a new reserved IP. All IP addresses must be in the same subnet as the primary IP.If reserved IP identities are provided, the specified reserved IPs must be unbound.If reserved IP prototype objects with addresses are provided, the addresses must be available on the virtual network interface's subnet. For any prototype objects that do not specify an address, an available address on the subnet will be automatically selected and reserved.
				Nested schema for **ips**:
				- `address` - (String) The IP address to reserve, which must not already be reserved on the subnet.If unspecified, an available address on the subnet will automatically be selected.
				- `auto_delete` - (Boolean) Indicates whether this reserved IP member will be automatically deleted when either`target` is deleted, or the reserved IP is unbound.
				- `href` - (String) The URL for this reserved IP.
				- `id` - (String) The unique identifier for this reserved IP.
				- `name` - (String) The name for this reserved IP. The name must not be used by another reserved IP in the subnet. Names starting with `ibm-` are reserved for provider-owned resources, and are not allowed. If unspecified, the name will be a hyphenated list of randomly-selected words.
			- `name` - (String) The name for this virtual network interface. The name must not be used by another virtual network interface in the VPC. If unspecified, the name will be a hyphenated list of randomly-selected words. Names beginning with `ibm-` are reserved for provider-owned resources, and are not allowed.
			- `primary_ip` - (List) The primary IP address to bind to the virtual network interface. May be either areserved IP identity, or a reserved IP prototype object which will be used to create anew reserved IP.If a reserved IP identity is provided, the specified reserved IP must be unbound.If a reserved IP prototype object with an address is provided, the address must beavailable on the virtual network interface's subnet. If no address is specified,an available address on the subnet will be automatically selected and reserved.
				Nested schema for **primary_ip**:
				- `address` - (String) The IP address to reserve, which must not already be reserved on the subnet.If unspecified, an available address on the subnet will automatically be selected.
				- `auto_delete` - (Boolean) Indicates whether this reserved IP member will be automatically deleted when either`target` is deleted, or the reserved IP is unbound.
				- `href` - (String) The URL for this reserved IP.
				- `id` - (String) The unique identifier for this reserved IP.
				- `name` - (String) The name for this reserved IP. The name must not be used by another reserved IP in the subnet. Names starting with `ibm-` are reserved for provider-owned resources, and are not allowed. If unspecified, the name will be a hyphenated list of randomly-selected words.
			- `resource_group` - (List) The resource group to use for this virtual network interface. If unspecified, thevirtual server instance's resource group will be used.
				Nested schema for **resource_group**:
				- `id` - (String) The unique identifier for this resource group.
			- `security_groups` - (List) The security groups to use for this virtual network interface. If unspecified, the default security group of the VPC for the subnet is used.
				Nested schema for **security_groups**:
				- `crn` - (String) The security group's CRN.
				- `href` - (String) The security group's canonical URL.
				- `id` - (String) The unique identifier for this security group.
			- `subnet` - (List) The associated subnet. Required if `primary_ip` does not specify a reserved IP.
				Nested schema for **subnet**:
				- `crn` - (String) The CRN for this subnet.
				- `href` - (String) The URL for this subnet.
				- `id` - (String) The unique identifier for this subnet.

	- `network_interfaces` - (List) A nested block describes the network interfaces for the template.

	  Nested scheme for `network_interfaces`:
		- `name` - (String) The name of the interface.
		- `primary_ipv4_address` - (String) The IPv4 address assigned to the network interface.
		- `subnet` - (String) The VPC subnet to assign to the interface.
		- `security_groups` - (String) List of security groups of  the subnet.
	- `placement_target` - (List) The placement restrictions to use for the virtual server instance.
	  Nested scheme for `placement_target`:
		- `crn` - (String) The unique identifier for this placement target.
		- `href` - (String) The CRN for this placement target.
		- `id` - (String) The URL for this placement target.
	- `profile` - (String) The number of instances created in the instance group.
	- `primary_network_attachment` - (List) The primary network attachment to create for the virtual server instance.
		Nested schema for **primary_network_attachment**:
		- `name` - (String) The name for this network attachment. Names must be unique within the instance the network attachment resides in. If unspecified, the name will be a hyphenated list of randomly-selected words.
		- `virtual_network_interface` - (List) A virtual network interface for the instance network attachment. This can be specifiedusing an existing virtual network interface, or a prototype object for a new virtualnetwork interface.If an existing virtual network interface is specified, `enable_infrastructure_nat` must be`false`.
			Nested schema for **virtual_network_interface**:
			- `allow_ip_spoofing` - (Boolean) Indicates whether source IP spoofing is allowed on this interface. If `false`, source IP spoofing is prevented on this interface. If `true`, source IP spoofing is allowed on this interface.
			- `auto_delete` - (Boolean) Indicates whether this virtual network interface will be automatically deleted when`target` is deleted.
			- `crn` - (String) The CRN for this virtual network interface.
			- `enable_infrastructure_nat` - (Boolean) If `true`:- The VPC infrastructure performs any needed NAT operations.- `floating_ips` must not have more than one floating IP.If `false`:- Packets are passed unchanged to/from the virtual network interface,  allowing the workload to perform any needed NAT operations.- `allow_ip_spoofing` must be `false`.- If the virtual network interface is attached:  - The target `resource_type` must be `bare_metal_server_network_attachment`.  - The target `interface_type` must not be `hipersocket`.
			- `href` - (String) The URL for this virtual network interface.
			- `id` - (String) The unique identifier for this virtual network interface.
			- `ips` - (List) Additional IP addresses to bind to the virtual network interface. Each item may be either a reserved IP identity, or as a reserved IP prototype object which will be used to create a new reserved IP. All IP addresses must be in the same subnet as the primary IP.If reserved IP identities are provided, the specified reserved IPs must be unbound.If reserved IP prototype objects with addresses are provided, the addresses must be available on the virtual network interface's subnet. For any prototype objects that do not specify an address, an available address on the subnet will be automatically selected and reserved.
				Nested schema for **ips**:
				- `address` - (String) The IP address to reserve, which must not already be reserved on the subnet.If unspecified, an available address on the subnet will automatically be selected.
				- `auto_delete` - (Boolean) Indicates whether this reserved IP member will be automatically deleted when either`target` is deleted, or the reserved IP is unbound.
				- `href` - (String) The URL for this reserved IP.
				- `id` - (String) The unique identifier for this reserved IP.
				- `name` - (String) The name for this reserved IP. The name must not be used by another reserved IP in the subnet. Names starting with `ibm-` are reserved for provider-owned resources, and are not allowed. If unspecified, the name will be a hyphenated list of randomly-selected words.
			- `name` - (String) The name for this virtual network interface. The name must not be used by another virtual network interface in the VPC. If unspecified, the name will be a hyphenated list of randomly-selected words. Names beginning with `ibm-` are reserved for provider-owned resources, and are not allowed.
			- `primary_ip` - (List) The primary IP address to bind to the virtual network interface. May be either areserved IP identity, or a reserved IP prototype object which will be used to create anew reserved IP.If a reserved IP identity is provided, the specified reserved IP must be unbound.If a reserved IP prototype object with an address is provided, the address must beavailable on the virtual network interface's subnet. If no address is specified,an available address on the subnet will be automatically selected and reserved.
				Nested schema for **primary_ip**:
				- `address` - (String) The IP address to reserve, which must not already be reserved on the subnet.If unspecified, an available address on the subnet will automatically be selected.
				- `auto_delete` - (Boolean) Indicates whether this reserved IP member will be automatically deleted when either`target` is deleted, or the reserved IP is unbound.
				- `href` - (String) The URL for this reserved IP.
				- `id` - (String) The unique identifier for this reserved IP.
				- `name` - (String) The name for this reserved IP. The name must not be used by another reserved IP in the subnet. Names starting with `ibm-` are reserved for provider-owned resources, and are not allowed. If unspecified, the name will be a hyphenated list of randomly-selected words.
			- `resource_group` - (List) The resource group to use for this virtual network interface. If unspecified, thevirtual server instance's resource group will be used.
				Nested schema for **resource_group**:
				- `id` - (String) The unique identifier for this resource group.
			- `security_groups` - (List) The security groups to use for this virtual network interface. If unspecified, the default security group of the VPC for the subnet is used.
				Nested schema for **security_groups**:
				- `crn` - (String) The security group's CRN.
				- `href` - (String) The security group's canonical URL.
				- `id` - (String) The unique identifier for this security group.
			- `subnet` - (List) The associated subnet. Required if `primary_ip` does not specify a reserved IP.
				Nested schema for **subnet**:
				- `crn` - (String) The CRN for this subnet.
				- `href` - (String) The URL for this subnet.
				- `id` - (String) The unique identifier for this subnet.

	- `primary_network_interfaces` - (List) A nested block describes the primary network interface for the template.

	  Nested scheme for `primary_network_interfaces`:
		- `name` - (String) The name of the interface.
		- `primary_ipv4_address` - (String) The IPv4 address assigned to the primary network interface.
		- `subnet` - (String) The VPC subnet to assign to the interface.
		- `security_groups` - (String) List of security groups of the subnet.
	- `reservation_affinity` - (Optional, List) The reservation affinity for the instance
      Nested scheme for `reservation_affinity`:
	  - `policy` - (Optional, String) The reservation affinity policy to use for this virtual server instance.
	    
		->**policy** 
		&#x2022; disabled: Reservations will not be used
      	</br>&#x2022; manual: Reservations in pool will be available for use
      - `pool` - (string) The unique identifier for this reservation
	- `resource_group` - (String) The resource group ID.
	- `total_volume_bandwidth` - (Integer) The amount of bandwidth (in megabits per second) allocated exclusively to instance storage volumes	
	- `user_data` -  (String) The user data provided for the instance.
	- `volume_attachments` - (List) A nested block describes the storage volume configuration for the template.

	  Nested scheme for `volume_attachments`:
		- `delete_volume_on_instance_delete` - (Bool) You can configure to delete the storage volume to delete based on instance deletion.
		- `name` - (String) The name of the boot volume.
		- `volume` - (String) The storage volume ID created in VPC.
		- `volume_prototype` - (List) A nested block describing prototype for the volume.

		  Nested scheme for `volume_prototype`:
		  - `capacity` - (String) The capacity of the volume in gigabytes. The specified minimum and maximum capacity values for creating or updating volumes can expand in the future.
		  - `encryption_key` - (String) The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.
		  - `iops` - (String) The maximum input/output operations per second (IOPS) for the volume.
		  - `profile` - (String) The global unique name for the volume profile to use for the volume.
		  - `tags` - (String) User Tags associated with the volume. (https://cloud.ibm.com/apidocs/tagging#types-of-tags)
	- `vpc` - (String) The VPC ID that the instance templates needs to be created.
	- `zone` - (String) The name of the zone.
