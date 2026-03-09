---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Instance profiles"
description: |-
  Manages IBM Cloud virtual server instance profiles.
---

# ibm_is_instance_profiles
Retrieve information of an existing virtual server instance profiles as a read-only data source. For more information, about virtual server instance profiles, see [instance profiles](https://cloud.ibm.com/docs/vpc?topic=vpc-profiles).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform

data "ibm_is_instance_profiles" "example" {
}

```

## Attribute reference
You can access the following attribute references after your data source is created. 

- `profiles` - (List) List of all server instance profiles in the region.

  Nested scheme for `profiles`:
  - `architecture` - (String) The default Operating System architecture for an instance of the profile.
  - `architecture_type` - (String) The type for this OS architecture.
  - `architecture_values` - (String) The supported OS architecture(s) for an instance with this profile.
  - `availability_class` - (List) The availability class for this OS architecture.
      Nested schema for **availability_class**:
      - `default` - (String) The default availability class for an instance with this profile. Allowable values are:   `spot`, `standard`.
      - `type` - (String) The type for this profile field. Allowable values are: `enum`.
      - `value` - (String) The value for this profile field.
      - `values` - (List) The permitted values for this profile field. Allowable list items are: `spot`, `standard`.
  - `name` - (String) The name of the virtual server instance profile.
  - `family` - (String) The family of the virtual server instance profile.
  - `bandwidth`  - (List) The collection of bandwidth information.

      Nested scheme for `bandwidth`:
      - `default` - (String) The default value for this profile field.
      - `max` - (String) The maximum value for this profile field.
      - `min` - (String) The minimum value for this profile field.
      - `step` - (String) The increment step value for this profile field.
      - `type` - (String) The type for this profile field.
      - `value` - (String) The value for this profile field.
      - `values` - (String) The permitted values for this profile field.

  - `cluster_network_attachment_count` - (List) Nested `cluster_network_attachment_count` blocks have the following structure:

      Nested schema for **cluster_network_attachment_count**:
      - `default` - (Integer)
      - `max` - (Integer) The maximum value for this profile field.
      - `min` - (Integer) The minimum value for this profile field.
      - `step` - (Integer)
      - `type` - (String) The type for this profile field.
      - `values` - (List) The permitted values for this profile field.

  - `confidential_compute_modes` - (List) 

      Nested schema for **confidential_compute_modes**:
      - `default` - (String) The default confidential compute mode for this profile.
      - `type` - (String) The type for this profile field.
      - `values` - (List) The supported confidential compute modes.

  - `disks` - (List) Collection of the instance profile's disks. Nested `disks` blocks has the following structure.

      Nested scheme for `disks`:
      - `quantity` (List) Nested `quantity` blocks has the following structure:
    
          Nested scheme for `quantity`:
          - `default` - (String) The default value for this profile field.
          - `max` - (String) The maximum value for this profile field.
          - `min` - (String) The minimum value for this profile field.
          - `step` - (String) The increment step value for this profile field.
          - `type` - (String) The type for this profile field.
          - `value` - (String) The value for this profile field.
          - `values` - (String) The permitted values for this profile field.

      - `size`  - (List) Nested `size` blocks has the following structure:
      
        Nested scheme for `size`:
        - `default` - (String) The default value for this profile field.
        - `max` - (String) The maximum value for this profile field.
        - `min` - (String) The minimum value for this profile field.
        - `step` - (String) The increment step value for this profile field.
        - `type` - (String) The type for this profile field.
        - `value` - (String) The value for this profile field.
        - `values` - (String) The permitted values for this profile field.
    - `supported_interface_types` - (List) Nested `supported_interface_types` blocks has the following structure:
  
        Nested scheme for `supported_interface_type`:
        - `default` - (String) The disk interface used for attaching the disk. The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally, halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
        - `type` - (String) The type for this profile field.
        - `values` - (String) The supported disk interfaces used for attaching the disk.
    - `status` - (String) The status of the instance profile. Values coule be, `previous`: This instance profile is an older revision, but remains provisionable and usable. `current`: This profile is the latest revision.
  - `gpu_count` - (List) Nested `gpu_count` blocks have the following structure:
    
      Nested scheme for `gpu_count`:
      - `default` - (String) The default value for this profile field.
      - `max` - (String) The maximum value for this profile field.
      - `min` - (String) The minimum value for this profile field.
      - `step` - (String) The increment step value for this profile field.
      - `type` - (String) The type for this profile field.
      - `value` - (String) The value for this profile field.
      - `values` - (String) The permitted values for this profile field.
  - `gpu_manufacturer` - (List) Nested `gpu_manufacturer` blocks have the following structure:
      Nested scheme for `gpu_manufacturer`:
      - `type` - (String) The type for this profile field.
      - `values` - (String) The permitted values for this profile field.
  - `gpu_memory` - (List) Nested `gpu_memory` blocks have the following structure:
      Nested scheme for `gpu_memory`:
      - `default` - (String) The default value for this profile field.
      - `max` - (String) The maximum value for this profile field.
      - `min` - (String) The minimum value for this profile field.
      - `step` - (String) The increment step value for this profile field.
      - `type` - (String) The type for this profile field.
      - `value` - (String) The value for this profile field.
      - `values` - (String) The permitted values for this profile field.
  - `gpu_model` - (List) Nested `gpu_model` blocks have the following structure:
      Nested scheme for `gpu_model`:
      - `type` - (String) The type for this profile field.
      - `values` - (String) The permitted values for this profile field.
  - `reservation_terms` - (List) Nested `reservation_terms` blocks have the following structure:
      Nested scheme for `reservation_terms`:
      - `type` - (String) The type for this profile field.
      - `values` - (String) The supported committed use terms for a reservation using this profile.
  - `total_volume_bandwidth`  Nested `total_volume_bandwidth` blocks have the following structure:
      Nested scheme for `total_volume_bandwidth`:
      - `type` - The type for this profile field.
      - `value` - The value for this profile field.
      - `default` - The default value for this profile field.
      - `max` - The maximum value for this profile field.
      - `min` - The minimum value for this profile field.
      - `step` - The increment step value for this profile field.
      - `values` - The permitted values for this profile field.
  - `href` - (String) The URL for this virtual server instance profile.
  - `memory` - (List) Nested `memory` blocks have the following structure:
    
    Nested scheme for `memory`:
    - `default` - (String) The default value for this profile field.
    - `max` - (String) The maximum value for this profile field.
    - `min` - (String) The minimum value for this profile field.
    - `step` - (String) The increment step value for this profile field.
    - `type` - (String) The type for this profile field.
    - `value` - (String) The value for this profile field.
    - `values` - (String) The permitted values for this profile field.

    - `network_attachment_count` - (List) The number of network attachments supported on an instance with this profile

        Nested scheme for `network_attachment_count`:
        - `max` - (Integer) The maximum number of network attachments supported by an instance using this profile.
        - `min` - (Integer) The minimum number of network attachments supported by an instance using this profile.
        - `type` - (String) The type for this profile field, Ex: range or dependent.
    
    - `network_bandwidth_mode` - (List) The permitted network bandwidth modes for an instance with this profile

        Nested schema for `network_bandwidth_mode`:
        - `default` - (string) The default network bandwidth mode for this profile.
        - `type` - (string) The type for this profile field.
        - `values` - (String) The supported network bandwidth modes for an instance with this profile.
        - `value` - (String) The value for this profile field.


  - `network_interface_count` - (List)
  
      Nested scheme for `network_interface_count`:
      - `max` - (Integer) The maximum number of vNICs supported by an instance using this profile.
      - `min` - (Integer) The minimum number of vNICs supported by an instance using this profile.
      - `type` - (String) The type for this profile field, Ex: range or dependent.
  - `numa_count` - (Integer) The number of NUMA nodes for the Instance Profile.
  - `port_speed` - (List) Nested `port_speed` blocks have the following structure:

      Nested scheme for `port_speed`:
      - `type` - (String) The type for this profile field.
      - `value` - (String) The value for this profile field.
  - `secure_boot_modes` - (List)
  
      Nested schema for **secure_boot_modes**:
      - `default` - (Boolean) The default secure boot mode for this profile.
      - `type` - (String) The type for this profile field.
      - `values` - (List) The supported `enable_secure_boot` values for an instance using   this profile.

  - `supported_cluster_network_profiles` - (List) The cluster network profiles that support this instance profile.

    Nested schema for **supported_cluster_network_profiles**:
      - `href` - (String) The URL for this cluster network profile.
      - `name` - (String) The globally unique name for this cluster network profile.
      - `resource_type` - (String) The resource type.


  - `vcpu_architecture` - (List) Nested `vcpu_architecture` blocks have the following structure:

      Nested scheme for `vcpu_architecture`:
      - `default` - (String) The default VCPU architecture for an instance with this profile.
      - `type` - (String) The type for this profile field.
      - `value` - (String) The VCPU architecture for an instance with this profile.
  - `vcpu_burst_limit` - (List) The permitted value for VCPU burst limit percentage for an instance with this profile.
	    
        Nested schema for **vcpu_burst_limit**:
		- `type` - (String) The type for this profile field.
		- `value` - (Integer) The value for this profile field.
  - `vcpu_count` - (List) Nested `vcpu_count` blocks have the following structure:

      Nested scheme for `vcpu_count`:
      - `default` - (String) The default value for this profile field.
      - `max` - (String) The maximum value for this profile field.
      - `min` - (String) The minimum value for this profile field.
      - `step` - (String) The increment step value for this profile field.
      - `type` - (String) The type for this profile field.
      - `value` - (String) The value for this profile field.
      - `values` - (String) The permitted values for this profile field.
  - `vcpu_manufacturer` - (List) Nested `vcpu_manufacturer` blocks have the following structure:

        Nested scheme for `vcpu_manufacturer`:
        - `default` - (String) The default VCPU manufacturer for an instance with this profile.
        - `type` - (String) The type for this profile field.
        - `value` - (String) The VCPU manufacturer for an instance with this profile.
  - `vcpu_percentage` - (List) The permitted values for VCPU percentage for an instance with this profile.
      
      Nested schema for **vcpu_percentage**:
      - `default` - (Integer) The default value for this profile field.
      - `type` - (String) The type for this profile field. 
      - `values` - (List) The permitted values for this profile field.

  - `volume_bandwidth_qos_modes` - (List) Nested `volume_bandwidth_qos_modes` blocks have the following structure:

      Nested scheme for `volume_bandwidth_qos_modes`:
      - `default` - (String) The default volume bandwidth QoS mode for this profile.
      - `type` - (String) The type for this profile field.
      - `values` - (String) The permitted volume bandwidth QoS modes for an instance using this profile. 