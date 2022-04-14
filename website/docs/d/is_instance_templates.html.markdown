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
	- `crn` - (String) The CRN of the instance template.
	- `default_trusted_profile_auto_link` - (Boolean) If set to `true`, the system will create a link to the specified `target` trusted profile during instance creation. Regardless of whether a link is created by the system or manually using the IAM Identity service, it will be automatically deleted when the instance is deleted. Default is true. 
	- `default_trusted_profile_target` - (String) The unique identifier or CRN of the default IAM trusted profile to use for this virtual server instance.
	- `href` - (String) The URL of the instance template.
	- `id` - (String) The ID of the instance template.
	- `image` - (String) The ID of the image to create the template.
	- `keys` - (String) List of SSH key IDs used to allow log in user to the instances.
	- `metadata_service_enabled` - (Boolean) Indicates whether the metadata service endpoint is available to the virtual server instance.
	- `name` - (String) The name of the instance template.
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
	- `primary_network_interfaces` - (List) A nested block describes the primary network interface for the template.

	  Nested scheme for `primary_network_interfaces`:
		- `name` - (String) The name of the interface.
		- `primary_ipv4_address` - (String) The IPv4 address assigned to the primary network interface.
		- `subnet` - (String) The VPC subnet to assign to the interface.
		- `security_groups` - (String) List of security groups of the subnet.
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
	- `vpc` - (String) The VPC ID that the instance templates needs to be created.
	- `zone` - (String) The name of the zone.
