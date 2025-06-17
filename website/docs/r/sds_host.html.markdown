---
layout: "ibm"
page_title: "IBM : ibm_sds_host"
description: |-
  Manages sds_host.
subcategory: "Ceph as a Service"
---

# ibm_sds_host

Create, update, and delete sds_hosts with this resource.

## Example Usage

```hcl
resource "ibm_sds_host" "sds_host_instance" {
  name = "my-host"
  nqn = "nqn.2014-06.org:9345"
  volume_mappings {
		status = "mapped"
		storage_identifier {
			subsystem_nqn = "nqn.2014-06.org:1234"
			namespace_id = 1
			namespace_uuid = "namespace_uuid"
			gateways {
				ip_address = "ip_address"
				port = 22
			}
		}
		href = "href"
		id = "1a6b7274-678d-4dfb-8981-c71dd9d4daa5-1a6b7274-678d-4dfb-8981-c71dd9d4da45"
		volume {
			id = "id"
		}
		host {
			id = "id"
			name = "name"
			nqn = "nqn"
		}
		subsystem_nqn = "nqn.2014-06.org:1234"
		namespace {
			id = 1
			uuid = "uuid"
		}
		gateways {
			ip_address = "ip_address"
			port = 22
		}
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `name` - (Optional, String) Unique name of the host.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
* `nqn` - (Required, String) The NQN (NVMe Qualified Name) as configured on the initiator (compute/host) accessing the storage.
  * Constraints: The maximum length is `63` characters. The minimum length is `4` characters. The value must match regular expression `/^nqn\\.\\d{4}-\\d{2}\\.[a-z0-9-]+(?:\\.[a-z0-9-]+)*:[a-zA-Z0-9.\\-:]+$/`.
* `volume_mappings` - (Optional, List) The host-to-volume map.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **volume_mappings**:
	* `gateways` - (Optional, List) List of NVMe gateways.
	  * Constraints: The maximum length is `10` items. The minimum length is `0` items.
	Nested schema for **gateways**:
		* `ip_address` - (Computed, String) Network information for volume/host mappings.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `port` - (Computed, Integer) Network information for volume/host mappings.
		  * Constraints: The maximum value is `65535`. The minimum value is `1`.
	* `host` - (Optional, List) Host mapping schema.
	Nested schema for **host**:
		* `id` - (Required, String) Unique identifer of the host.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `name` - (Required, String) Unique name of the host.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `nqn` - (Required, String) The NQN (NVMe Qualified Name) as configured on the initiator (compute/host) accessing the storage.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
	* `href` - (Required, String) The URL for this resource.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `10` characters.
	* `id` - (Required, String) Unique identifier of the mapping.
	  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `namespace` - (Optional, List) The NVMe namespace properties for a given volume mapping.
	Nested schema for **namespace**:
		* `id` - (Optional, Integer) NVMe namespace ID that can be used to co-relate the discovered devices on host to the corresponding volume.
		  * Constraints: The maximum value is `32`. The minimum value is `1`.
		* `uuid` - (Optional, String) UUID of the NVMe namespace.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `status` - (Required, String) The status of the volume mapping. The enumerated values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
	  * Constraints: Allowable values are: `pending`, `mapped`, `pending_unmapping`, `mapping_failed`. The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `storage_identifier` - (Optional, List) Storage network and ID information associated with a volume/host mapping.
	Nested schema for **storage_identifier**:
		* `gateways` - (Required, List) List of NVMe gateways.
		  * Constraints: The maximum length is `10` items. The minimum length is `0` items.
		Nested schema for **gateways**:
			* `ip_address` - (Computed, String) Network information for volume/host mappings.
			  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
			* `port` - (Computed, Integer) Network information for volume/host mappings.
			  * Constraints: The maximum value is `65535`. The minimum value is `1`.
		* `namespace_id` - (Required, Integer) NVMe namespace ID that can be used to co-relate the discovered devices on host to the corresponding volume.
		  * Constraints: The maximum value is `32`. The minimum value is `1`.
		* `namespace_uuid` - (Required, String) The namespace UUID associated with a volume/host mapping.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `subsystem_nqn` - (Required, String) The NVMe target subsystem NQN (NVMe Qualified Name) that can be used for doing NVMe connect by the initiator.
		  * Constraints: The maximum length is `63` characters. The minimum length is `4` characters. The value must match regular expression `/^nqn\\.\\d{4}-\\d{2}\\.[a-z0-9-]+(?:\\.[a-z0-9-]+)*:[a-zA-Z0-9.\\-:]+$/`.
	* `subsystem_nqn` - (Optional, String) The NVMe target subsystem NQN (NVMe Qualified Name) that can be used for doing NVMe connect by the initiator.
	  * Constraints: The maximum length is `63` characters. The minimum length is `4` characters. The value must match regular expression `/^nqn\\.\\d{4}-\\d{2}\\.[a-z0-9-]+(?:\\.[a-z0-9-]+)*:[a-zA-Z0-9.\\-:]+$/`.
	* `volume` - (Optional, List) The volume reference.
	Nested schema for **volume**:
		* `id` - (Required, String) Unique identifer of the host.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `name` - (Computed, String) Unique name of the host.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the sds_host.
* `created_at` - (String) The date and time when the resource was created.
* `href` - (String) The URL for this resource.
  * Constraints: The maximum length is `1000` characters. The minimum length is `10` characters.


## Import

You can import the `ibm_sds_host` resource by using `id`. Unique identifer of the host.

# Syntax
<pre>
$ terraform import ibm_sds_host.sds_host &lt;id&gt;
</pre>
