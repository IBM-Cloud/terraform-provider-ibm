---
layout: "ibm"
page_title: "IBM : ibm_sds_volume"
description: |-
  Manages sds_volume.
subcategory: "Ceph as a Service"
---

# ibm_sds_volume

Create, update, and delete sds_volumes with this resource.

## Example Usage

```hcl
resource "ibm_sds_volume" "sds_volume_instance" {
  capacity = 10
  name = "my-volume"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `capacity` - (Required, Integer) The capacity of the volume (in gigabytes).
  * Constraints: The maximum value is `32000`. The minimum value is `1`.
* `name` - (Optional, String) Unique name of the host.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the sds_volume.
* `bandwidth` - (Integer) The maximum bandwidth (in megabits per second) for the volume.
  * Constraints: The maximum value is `8192`. The minimum value is `1`.
* `created_at` - (String) The date and time that the volume was created.
* `href` - (String) The URL for this resource.
  * Constraints: The maximum length is `1000` characters. The minimum length is `10` characters.
* `iops` - (Integer) Iops The maximum I/O operations per second (IOPS) for this volume.
  * Constraints: The maximum value is `96000`. The minimum value is `150`.
* `resource_type` - (String) The resource type of the volume.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
* `status` - (String) The status of the volume resource. The enumerated values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
  * Constraints: Allowable values are: `available`, `pending`, `pending_deletion`, `updating`. The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `status_reasons` - (List) The reasons for the current status (if any).
  * Constraints: The maximum length is `10` items. The minimum length is `0` items.
Nested schema for **status_reasons**:
	* `code` - (String) A snake case string succinctly identifying the status reason.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `message` - (String) An explanation of the status reason.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `10` characters. The value must match regular expression `/^[ -~\\n\\r\\t]*$/`.
	* `more_info` - (String) Link to documentation about this status reason.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters.
* `volume_mappings` - (List) List of volume mappings for this volume.
  * Constraints: The maximum length is `200` items. The minimum length is `0` items.
Nested schema for **volume_mappings**:
	* `gateways` - (List) List of NVMe gateways.
	  * Constraints: The maximum length is `10` items. The minimum length is `0` items.
	Nested schema for **gateways**:
		* `ip_address` - (String) Network information for volume/host mappings.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `port` - (Integer) Network information for volume/host mappings.
		  * Constraints: The maximum value is `65535`. The minimum value is `1`.
	* `host` - (List) Host mapping schema.
	Nested schema for **host**:
		* `id` - (String) Unique identifer of the host.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `name` - (String) Unique name of the host.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `nqn` - (String) The NQN (NVMe Qualified Name) as configured on the initiator (compute/host) accessing the storage.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
	* `href` - (String) The URL for this resource.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `10` characters.
	* `id` - (String) Unique identifier of the mapping.
	  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `namespace` - (List) The NVMe namespace properties for a given volume mapping.
	Nested schema for **namespace**:
		* `id` - (Integer) NVMe namespace ID that can be used to co-relate the discovered devices on host to the corresponding volume.
		  * Constraints: The maximum value is `32`. The minimum value is `1`.
		* `uuid` - (String) UUID of the NVMe namespace.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `status` - (String) The status of the volume mapping. The enumerated values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
	  * Constraints: Allowable values are: `pending`, `mapped`, `pending_unmapping`, `mapping_failed`. The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `storage_identifier` - (List) Storage network and ID information associated with a volume/host mapping.
	Nested schema for **storage_identifier**:
		* `gateways` - (List) List of NVMe gateways.
		  * Constraints: The maximum length is `10` items. The minimum length is `0` items.
		Nested schema for **gateways**:
			* `ip_address` - (String) Network information for volume/host mappings.
			  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
			* `port` - (Integer) Network information for volume/host mappings.
			  * Constraints: The maximum value is `65535`. The minimum value is `1`.
		* `namespace_id` - (Integer) NVMe namespace ID that can be used to co-relate the discovered devices on host to the corresponding volume.
		  * Constraints: The maximum value is `32`. The minimum value is `1`.
		* `namespace_uuid` - (String) The namespace UUID associated with a volume/host mapping.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `subsystem_nqn` - (String) The NVMe target subsystem NQN (NVMe Qualified Name) that can be used for doing NVMe connect by the initiator.
		  * Constraints: The maximum length is `63` characters. The minimum length is `4` characters. The value must match regular expression `/^nqn\\.\\d{4}-\\d{2}\\.[a-z0-9-]+(?:\\.[a-z0-9-]+)*:[a-zA-Z0-9.\\-:]+$/`.
	* `subsystem_nqn` - (String) The NVMe target subsystem NQN (NVMe Qualified Name) that can be used for doing NVMe connect by the initiator.
	  * Constraints: The maximum length is `63` characters. The minimum length is `4` characters. The value must match regular expression `/^nqn\\.\\d{4}-\\d{2}\\.[a-z0-9-]+(?:\\.[a-z0-9-]+)*:[a-zA-Z0-9.\\-:]+$/`.
	* `volume` - (List) The volume reference.
	Nested schema for **volume**:
		* `id` - (String) Unique identifer of the host.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `name` - (String) Unique name of the host.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.


## Import

You can import the `ibm_sds_volume` resource by using `id`. The volume profile id.

# Syntax
<pre>
$ terraform import ibm_sds_volume.sds_volume &lt;id&gt;
</pre>
