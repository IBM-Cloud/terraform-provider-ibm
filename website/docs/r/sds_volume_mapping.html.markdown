---
layout: "ibm"
page_title: "IBM : ibm_sds_volume_mapping"
description: |-
  Manages sds_volume_mapping.
subcategory: "Ceph as a Service"
---

# ibm_sds_volume_mapping

Create, update, and delete sds_volume_mappings with this resource.

## Example Usage

```hcl
resource "ibm_sds_volume_mapping" "sds_volume_mapping_instance" {
  host_id = ibm_sds_host.sds_host_instance.id
  volume {
		id = "id"
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `host_id` - (Required, Forces new resource, String) A unique host ID.
  * Constraints: The maximum length is `200` characters. The minimum length is `0` characters. The value must match regular expression `/^\\S+$/`.
* `volume` - (Required, List) The volume reference.
Nested schema for **volume**:
	* `id` - (Required, String) Unique identifer of the host.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
	* `name` - (Computed, String) Unique name of the host.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the sds_volume_mapping.
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
* `volume_mapping_id` - (String) Unique identifier of the mapping.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.


## Import

You can import the `ibm_sds_volume_mapping` resource by using `id`.
The `id` property can be formed from `host_id`, and `volume_mapping_id` in the following format:

<pre>
&lt;host_id&gt;/&lt;volume_mapping_id&gt;
</pre>
* `host_id`: A string in the format `r134-69d5c3e2-8229-45f1-89c8-e4dXXb2e126e`. A unique host ID.
* `volume_mapping_id`: A string in the format `1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45`. Unique identifier of the mapping.

# Syntax
<pre>
$ terraform import ibm_sds_volume_mapping.sds_volume_mapping &lt;host_id&gt;/&lt;volume_mapping_id&gt;
</pre>
