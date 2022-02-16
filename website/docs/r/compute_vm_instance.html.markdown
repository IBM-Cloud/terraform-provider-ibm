---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: compute_vm_instance"
description: |-
  Manages IBM Cloud VM instances.
---

# ibm_compute_vm_instance
Create, update, and delete a Virtual Machine (VM) instance. For more information, about IBM Cloud Virtual Machine instance, see [migrating VMDK or VHD images to VPC](https://cloud.ibm.com/docs/cloud-infrastructure?topic=cloud-infrastructure-migrating-images-vpc).

**Note**

- For more information, see the [IBM Cloud Classic Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/services/SoftLayer_Virtual_Guest).
- Update is not supported when the `bulk_vms` parameter is used.

## Example usage
In the following example, you can create a VM instance using a Debian image:

```terraform
resource "ibm_compute_vm_instance" "twc_terraform_sample" {
  hostname                   = "twc-terraform-sample-name"
  domain                     = "bar.example.com"
  os_reference_code          = "DEBIAN_8_64"
  datacenter                 = "wdc01"
  network_speed              = 10
  hourly_billing             = true
  private_network_only       = false
  cores                      = 1
  memory                     = 1024
  disks                      = [25, 10, 20]
  user_metadata              = "{\"value\":\"newvalue\"}"
  dedicated_acct_host_only   = true
  local_disk                 = false
  public_vlan_id             = 1391277
  private_vlan_id            = 7721931
  private_security_group_ids = [576973]
}
```

### In the following example, you can create a VM instance using a block device template:

```terraform
resource "ibm_compute_vm_instance" "terraform-sample-BDTGroup" {
  hostname   = "terraform-sample-blockDeviceTemplateGroup"
  domain     = "bar.example.com"
  datacenter = "ams01"

  //public_network_speed = 10
  hourly_billing = false
  cores          = 1
  memory         = 1024
  local_disk     = false
  image_id       = 12345
  tags = [
    "collectd",
    "mesos-master",
  ]
  public_subnet  = "50.97.46.160/28"
  private_subnet = "10.56.109.128/26"
}
```

### In the following example, you can create a VM instance using a flavor:

```terraform
resource "ibm_compute_vm_instance" "terraform-sample-flavor" {
  hostname             = "terraform-sample-flavor"
  domain               = "bar.example.com"
  os_reference_code    = "DEBIAN_8_64"
  datacenter           = "dal06"
  network_speed        = 10
  hourly_billing       = true
  local_disk           = false
  private_network_only = false
  flavor_key_name      = "B1_2X8X25"
  user_metadata        = "{\\\"value\\\":\\\"newvalue\\\"}"

  // provide disk 3, 4, 5 and so on
  disks = [10, 20, 30]
  tags  = ["collectd"]

  // It should be false
  dedicated_acct_host_only = false
  ipv6_enabled             = true
  secondary_ip_count       = 4
  notes                    = "VM notes"
}
```

### In the following example, you can create multiple vm's

```terraform
resource "ibm_compute_vm_instance" "terraform-bulk-vms" {
  bulk_vms {
    hostname = "vm1"

    domain = "bar.example.com"
  }
  bulk_vms {
    hostname = "vm2"

    domain = "bar.example.com"
  }

  os_reference_code    = "CENTOS_7_64"
  datacenter           = "dal09"
  network_speed        = 100
  hourly_billing       = true
  private_network_only = true
  cores                = 1
  memory               = 1024
  disks                = [25]
  local_disk           = false
}

```

### Example to create a VM instance by using a datacenter_choice. 
This example creates a VM instance by using a datacenter_choice. If VM fails to place order on first datacenter or vlans it retries to place order on subsequent datacenters and vlans untill place order is successful:

```terraform
resource "ibm_compute_vm_instance" "terraform-retry" {
  hostname          = "vmretry"
  domain            = "example.com"
  network_speed     = 100
  hourly_billing    = true
  cores             = 1
  memory            = 1024
  local_disk        = false
  os_reference_code = "DEBIAN_7_64"
  disks             = [25]

  datacenter_choice = [
    {
      datacenter      = "dal09"
      public_vlan_id  = 123245
      private_vlan_id = 123255
    },
    {
      datacenter = "wdc54"
    },
    {
      datacenter      = "dal09"
      public_vlan_id  = 153345
      private_vlan_id = 123255
    },
    {
      datacenter = "dal06"
    },
    {
      datacenter      = "dal09"
      public_vlan_id  = 123245
      private_vlan_id = 123255
    },
    {
      datacenter      = "dal09"
      public_vlan_id  = 1232454
      private_vlan_id = 1234567
    },
  ]

  //User can configure timeouts
  timeouts {
    create = "20m"
    update = "20m"
    delete = "20m"
  }
}


```  

### Example of a quote based ordering

```terraform
resource "ibm_compute_vm_instance" "vm1" {
  # Mandatory fields
  hostname             = "terraformquote"
  domain               = "IBM.cloud"
  quote_id             = "2877000"

  # Optional fields
  os_reference_code    = "DEBIAN_9_64"
  datacenter           = "dal06"
  network_speed        = 100
  hourly_billing       = false
  private_network_only = false
  flavor_key_name      = "B1_2X8X100"
  local_disk           = true
}

```

### In the following example, you can create a VM instance using a Reserved Capacity:

```terraform
resource "ibm_compute_vm_instance" "reservedinstance" {
  hostname          = "terraformreserved"
  domain            = "ibm.com"
  os_reference_code = "DEBIAN_9_64"
  datacenter        = "lon02"
  network_speed     = 10
  hourly_billing    = false
  notes             = "VM notes"
  reserved_capacity_id = "110974"
  local_disk        = false
  reserved_instance_primary_disk = 100
}
```

## Timeouts

The `ibm_is_instance` resource provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 90 minutes) Used to Wait for virtual guest creation.
- **update** - (Default 90 minutes) Used to Wait for upgrade transactions to finish.
- **delete** - (Default 90 minutes) Used to Wait for no active transactions.


## Argument reference
Review the argument references that you can specify for your resource. 

- `block_storage_ids`- (Optional, Array of Integers) File storage to which this computing instance has access. File storage must be in the same data center as the Bare Metal server. If you use this argument to authorize, access to file storage, then do not use the `allowed_virtual_guest_ids` argument in the `ibm_storage_block` resource in order to prevent the same storage be added twice.
- `bulk_vms`- (Optional, Forces new resource, List) Hostname and domain of the computing instance. The minimum number of VM's to be defined is 2.

  Nested scheme for `bulk_vms`:
	- `domain` - (Required, Forces new resource, String) The domain for the computing instance. If you set this option, do not specify `hostname` and `domain` at the same time.
	- `hostname` - (Required, String) The hostname for the computing instance.
- `cores` - (Optional, Integer) The number of CPU cores that you want to allocate. If you set this option, do not specify `flavor_key_name` at the same time.
- `datacenter` - (Optional, Forces new resource, String) The data center in which you want to provision the instance. **Note** If `dedicated_host_name` or `dedicated_host_id` is provided then the datacenter should be same as the dedicated host datacenter. If `placement_group_name` or `placement_group_id`    is provided then the datacenter should be same as the placement group datacenter.    Conflicts with `datacenter_choice`.
- `datacenter_choice` - (Optional, List of Objects) A nested block to describe datacenter choice options to retry on different data centers and VLANs. Nested `datacenter_choice` blocks must have the following structure:

  Nested scheme for `datacenter_choice`:
  - `datacenter` - (Required, String) The datacenter in which you want to provision the instance.
  - `private_vlan_id` - (Optional, Forces new resource, String) The private VLAN ID for the private network interface of the instance. You can find accepted values in the [VLAN doc](https://cloud.ibm.com/classic/network/vlans). Click the VLAN that you want to use and note the ID number in the browser URL. You can also refer to a VLAN name by using a data source.  **Note** Conflicts with `datacenter`, `private_vlan_id`, `public_vlan_id`, `placement_group_name` and `placement_group_id`.
  - `public_vlan_id` - (Optional, String) The public VLAN ID for the public network interface of the instance. Accepted values are in the [VLAN doc](https://cloud.ibm.com/classic/network/vlans). Click the VLAN that you want to use and note the ID number in the browser URL. You can also refer to a VLAN name by using a data source.
- `dedicated_acct_host_only` - (Optional, Forces new resource, Bool)  Specifies whether the instance must only run on hosts with instances from the same account. The default value is **false**. If VM is provisioned by using `flavorKeyName`, value should be set to **false**.  **Note** Conflicts with `dedicated_host_name`, `dedicated_host_id`, `placement_group_name` and `placement_group_id`.
- `dedicated_host_id` - (Optional, Forces new resource, Integer) Specifies [dedicated host](https://cloud.ibm.com/docs/virtual-servers?topic=virtual-servers-dedicated-virtual-servers) for the instance by its ID. **Note** Conflicts with `dedicated_acct_host_only`, `dedicated_host_name`, `placement_group_name` and `placement_group_id`.
- `dedicated_host_name` - (Optional, Forces new resource, String) Specifies [dedicated host](https://cloud.ibm.com/docs/virtual-servers?topic=virtual-servers-dedicated-virtual-servers) for the instance by its name. **Note** Conflicts with `dedicated_acct_host_only`, `dedicated_host_id`, `placement_group_name` and `placement_group_id`.
- `disks`- (Optional, Array of Integers) The numeric disk sizes in GBs for the instance's block device and disk image settings. The default value is the smallest available capacity for the primary disk. If you specify an image template, the template provides the disk capacity. If you specify the `flavorKeyName`, first disk is provided by the flavor.
- `domain` - (Optional, String) The domain for the computing instance. If you set this option, do not specify `bulk_vms` at the same time.
- `evault` - (Optional, Forces new resource, Integer) Allowed `Evault` in GB per month for monthly based servers.
- `file_storage_ids`- (Optional, Array of Integers) File storage to which this computing instance has access. File storage must be in the same data center as the Bare Metal server. If you use this argument to authorize, access to file storage, then do not use the `allowed_virtual_guest_ids` argument in the `ibm_storage_file` resource in order to prevent the same storage be added twice.
- `flavor_key_name` - (Optional, String) The flavor key name that you want to use to provision the instance. To see available key name, log in to the [IBM Cloud Classic Infrastructure API](https://api.softlayer.com/rest/v3/SoftLayer_Virtual_Guest/getCreateObjectOptions.json), that uses your API key as the password. If you set this option, do not specify `cores` and `memory` at the same time.
- `hourly_billing` - (Optional, Forces new resource, Bool) The billing type for the instance. When set to **true**, the computing instance is billed on hourly usage. Otherwise, the instance is billed monthly. The default value is **true**.
- `hostname` - (Optional, String) The hostname for the computing instance. If you set this option, do not specify `bulk_vms` at the same time.
- `ipv6_enabled` - (Optional, Forces new resource, Bool) The primary public IPv6 address. The default value is **false**.
- `ipv6_static_enabled` - (Optional, Bool) The public static IPv6 address block of `/64`. The default value is **false**.
- `image_id` - (Optional, Forces new resource, Integer) The image template ID that you want to use to provision the computing instance. This is not the global identifier (UUID), but the image template group ID that should point to a valid global identifier. To retrieve the image template ID from the IBM Cloud infrastructure customer portal, navigate to **Devices > Manage > Images**, click the image that you want, and note the ID number in the resulting URL. **Note** Conflicts with `os_reference_code`.
- `local_disk` - (Optional, Forces new resource, Bool) The disk type for the instance. When set to **true**, the disks for the computing instance are provisioned on the host that the instance runs. Otherwise, SAN disks are provisioned. The default value is **true**.
- `memory` - (Optional, Integer) The amount of memory, expressed in megabytes, that you want to allocate. If you set this option, do not specify `flavor_key_name` at the same time.
- `network_speed` - (Optional, Integer) The connection speed (in Mbps) for the instance's network components. The default value is `100`.
- `notes` - (Optional, String)  Descriptive text of up to 1000 characters about the VM instance.
- `os_reference_code` - (Optional, Forces new resource, String) The operating system reference code that is used to provision the computing instance. To see available OS reference codes, log in to the [IBM Cloud Classic Infrastructure API](https://api.softlayer.com/rest/v3/SoftLayer_Virtual_Guest_Block_Device_Template_Group/getVhdImportSoftwareDescriptions.json?objectMask=referenceCode), that uses your API key as the password. **Note** Conflicts with `image_id`.
- `placement_group_id` - (Optional, Forces new resource, Integer) Specifies [placement group](https://cloud.ibm.com/docs/virtual-servers?topic=virtual-servers-dedicated-virtual-servers) for the instance by its ID. **Note** Conflicts with `dedicated_acct_host_only`, `dedicated_host_name`, `dedicated_host_id` and `placement_group_name`.
- `placement_group_name` - (Optional, Forces new resource, String) Specifies [placement group](https://cloud.ibm.com/docs/virtual-servers?topic=virtual-servers-dedicated-virtual-servers) for the instance by its name. **Note** Conflicts with `dedicated_acct_host_only`, `dedicated_host_id`, `dedicated_host_name` and `placement_group_id`-
- `private_network_only` - (Optional, Forces new resource, Bool) When set to **true**, a compute instance has only access to the private network. The default value is **false**.
- `private_security_group_ids`- (Optional, Force new resource, Array of integers) The IDs of security groups to apply on the private interface. This attribute can't be updated. You can use this parameter to add a security group to your virtual server instance when you create it. If you want to add or remove security groups later, you must use the `ibm_network_interface_sg_attachment` resource. If you use this attribute in addition to `ibm_network_interface_sg_attachment` resource you might experience errors. So use one of these consistently for a particular virtual server instance.
- `public_vlan_id` - (Optional, Forces new resource, Integer) The public VLAN ID for the public network interface of the instance. Accepted values are in the [VLAN doc](https://cloud.ibm.com/classic/network/vlans). Click the VLAN that you want and notes the ID number in the browser URL. You can also refer to a VLAN name by using a data source. **Note** Conflicts with `datacenter_choice`.
- `private_vlan_id` - (Optional, Forces new resource, Integer) The private VLAN ID for the private network interface of the instance. You can find accepted values in the [VLAN doc](https://cloud.ibm.com/classic/network/vlans). Click the VLAN that you want to use and note the ID number in the browser URL. You can also refer to a VLAN name by using a data source. **Note** Conflicts with `datacenter_choice`.
- `post_install_script_uri` - (Optional, Forces new resource, String) The URI of the script to be downloaded and executed after installation is complete.
- `quote_id` - (Optional, Forces new resource, String) When you define the `quote_id`, Terraform uses the specification in the quote to create a virtual server. You can find the quote ID in the IBM Cloud portal.
- `public_security_group_ids`- (Optional, Force new resource, Array of Integers) The IDs of security groups to apply on the public interface. This attribute can't be updated. You can use this parameter to add a security group to your virtual server instance when you create it. If you want to add or remove security groups later, you must use the `ibm_network_interface_sg_attachment` resource. If you use this attribute in addition to `ibm_network_interface_sg_attachment` resource, you might experience errors. So use one of these consistently for a particular virtual server instance.
- `public_subnet` - (Optional, Forces new resource, String) The public subnet for the public network interface of the instance. Accepted values are primary public networks. You can find accepted values in the [subnets doc](https://cloud.ibm.com/classic/network/subnets), **Note** You can see the list of public subnets of your account.
- `private_subnet` - (Optional, Forces new resource, String) The private subnet for the private network interface of the instance. Accepted values are primary private networks. You can find accepted values in the [subnets doc](https://cloud.ibm.com/classic/network/subnets), **Note** You can see the list of private subnets of your account.
- `public_bandwidth_limited` - (Optional, Forces new resource, Integer) Allowed public network traffic in GB per month. It can be greater than 0 when the server is a monthly based server. Defaults to the smallest available capacity for the public bandwidth are used.  **Note** Conflicts with `private_network_only` and `public_bandwidth_unlimited`.
- `public_bandwidth_unlimited` - (Optional, Forces new resource, Bool) Allowed unlimited public network traffic in GB per month for a monthly based server. The `network_speed` should be 100 Mbps. Default value is **false**. **Note** Conflicts with `private_network_only` and `public_bandwidth_limited`.
- `reserved_capacity_id` - (Optional, Forces new resource, Integer) The reserved capacity ID to provision the instance.
- `reserved_capacity_name` - (Optional, Forces new resource, String) The reserved capacity name to provision the instance.
- `reserved_instance_primary_disk` - (Optional, Forces new resource, Integer) Size of the main drive.    **Note** We can provision only monthly based servers in a reserved capacity.
- `secondary_ip_count` - (Optional, Forces new resource, Integer) Specifies secondary public IPv4 addresses. Accepted values are `4` and `8`. 
- `ssh_key_ids`- (Optional, Array of integers) The SSH key IDs to install on the computing instance when the instance provisions. **Note** If you don't know the ID(s) for your SSH keys, you can reference your SSH keys by their labels.
- `tags` (Optional, Array of Strings) Tags associated with the VM instance. Permitted characters include: A-Z, 0-9, whitespace, `_` (underscore), `- ` (hyphen), `.` (period), and `:` (colon). All other characters are removed.
- `transient` - (Optional, Forces new resource, Bool) Specifies whether to provision a transient virtual server. The default value is **false**. Transient instances cannot be upgraded or downgraded. Transient instances cannot use local storage. **Note** Conflicts with `dedicated_acct_host_only`, `dedicated_host_id`, `dedicated_host_name`, `cores`, `memory`, `public_bandwidth_limited` and `public_bandwidth_unlimited`.
- `wait_time_minutes` - (Optional, Integer) The duration, expressed in minutes, to wait for the VM instance to become available before declaring it as created. It is also the same amount of time waited for no active transactions before proceeding with an update or deletion. The default value is `90`.
- `wait_time_minutes`- (Deprecated, Integer) Use Timeouts block to wait for the VM instance to become available, or while waiting for non active transactions before proceeding with an update or deletion. The default value is `90`.
- `user_metadata` - (Optional, Forces new resource, String) Arbitrary data to be made available to the computing instance.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the VM instance.
- `ipv4_address` - (String) The public IPv4 address of the VM instance.
- `ip_address_id_private` - (String) The unique identifier for the private IPv4 address that is assigned to the VM instance.
- `ipv4_address_private` - (String) The private IPv4 address of the VM instance.
- `ip_address_id` - (String) The unique identifier for the public IPv4 address that is assigned to the VM instance.
- `ipv6_address` - (String) The public IPv6 address of the VM instance provided when `ipv6_enabled` is set to **true**.
- `ipv6_address_id` - (String) The unique identifier for the public IPv6 address assigned to the VM instance provided when `ipv6_enabled` is set to **true**.
- `private_subnet_id` - (String) The unique identifier of the subnet `ipv4_address_private` belongs to.
- `public_ipv6_subnet` - (String) The public IPv6 subnet provided when `ipv6_enabled` is set to **true**.
- `public_ipv6_subnet_id` - (String) The unique identifier of the subnet `ipv6_address` belongs to.
- `public_subnet_id` - (String) The unique identifier of the subnet `ipv4_address` belongs to.
- `secondary_ip_addresses` - (String) The public secondary IPv4 addresses of the VM instance.
- `public_interface_id` - (String) The ID of the primary public interface.
- `private_interface_id` - (String) The ID of the primary private interface.

## Import

The `ibm_compute_vm_instance` resource can be imported by using instance ID.

**Example**

```
$ terraform import ibm_compute_vm_instance.example 88205074
```
