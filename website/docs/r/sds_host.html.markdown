---
layout: "ibm"
page_title: "IBM : ibm_sds_host"
description: |-
  Manages sds_host.
subcategory: "sdsaas"
---

# ibm_sds_host

Create, update, and delete sds_hosts with this resource.

## Example Usage

```hcl
resource "ibm_sds_host" "sds_host_instance" {
  name = "my-host"
  nqn = "nqn.2014-06.org:9345"
  volumes {
		status = "status"
		volume_id = "volume_id"
		volume_name = "volume_name"
		storage_identifiers {
			id = "id"
			namespace_id = 1
			namespace_uuid = "namespace_uuid"
			network_info {
				gateway_ip = "gateway_ip"
				port = 1
			}
		}
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `name` - (Optional, String) The name for this host. The name must not be used by another host.  If unspecified, the name will be a hyphenated list of randomly-selected words.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
* `nqn` - (Required, String) The NQN of the host configured in customer's environment.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
* `volumes` - (Optional, List) The host-to-volume map.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **volumes**:
	* `status` - (Optional, String) The current status of a volume/host mapping attempt.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
	* `storage_identifiers` - (Optional, List) Storage network and ID information associated with a volume/host mapping.
	Nested schema for **storage_identifiers**:
		* `id` - (Optional, String) The storage ID associated with a volume/host mapping.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `namespace_id` - (Optional, Integer) The namespace ID associated with a volume/host mapping.
		* `namespace_uuid` - (Optional, String) The namespace UUID associated with a volume/host mapping.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `network_info` - (Optional, List) The IP and port for volume/host mappings.
		  * Constraints: The maximum length is `200` items. The minimum length is `1` item.
		Nested schema for **network_info**:
			* `gateway_ip` - (Optional, String) Network information for volume/host mappings.
			  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
			* `port` - (Optional, Integer) Network information for volume/host mappings.
	* `volume_id` - (Required, String) The volume ID that needs to be mapped with a host.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
	* `volume_name` - (Required, String) The volume name.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the sds_host.
* `created_at` - (String) The date and time that the host was created.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.


## Import

You can import the `ibm_sds_host` resource by using `id`. The unique identifier for this host.

# Syntax
<pre>
$ terraform import ibm_sds_host.sds_host &lt;id&gt;
</pre>

# Example
```
$ terraform import ibm_sds_host.sds_host 1a6b7274-678d-4dfb-8981-c71dd9d4daa5
```
