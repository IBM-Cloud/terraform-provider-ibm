---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : ibm_satellite_link"
description: |-
  Get information about ibm_satellite_link
---

# ibm_satellite_link

Provides a read-only data source for ibm_satellite_link. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_satellite_link" "satellite_link" {
	location = "location_id"
}
```

## Argument reference

The following arguments are supported:

* `location` - (Required, string) The Location ID.

## Attribute reference

In addition to all arguments above, the following attributes are exported:

* `crn` - Service instance associated with this location.
* `created_at` - Timestamp of creation of location.
* `description` - Description of the location.
* `id` - The unique identifier of the ibm_satellite_link.
* `last_change` - Timestamp of latest modification of location.
* `performance` - The last performance data of the Location. Nested `performance` blocks have the following structure:
	* `avg_latency` - Average latency calculated form latency of each Connector between Tunnel Server, unit is ms. -1 means no Connector established Tunnel.
	* `bandwidth` - Average Tatal Bandwidth of last two minutes, unit is Byte/s.
	* `connectors` - The last performance data of the Location read from each Connector. Nested `connectors` blocks have the following structure:
		* `connector` - The name of the connector reported the performance data.
		* `latency` - Latency between Connector and the Tunnel Server it connected.
		* `rx_bw` - Average Transmitted (to Location) Bandwidth of last two minutes read from the Connector, unit is Byte/s.
		* `tx_bw` - Average Transmitted (to Location) Bandwidth of last two minutes read from the Connector, unit is Byte/s.
	* `tunnels` - Tunnels number estbalished from the Location.
	* `health_status` - Tunnels health status based on the Tunnels number established. Down(0)/Critical(1)/Up(>=2).
	* `rx_bandwidth` - Average Receive (to Cloud) Bandwidth of last two minutes, unit is Byte/s.
	* `tx_bandwidth` - Average Transmitted (to Location) Bandwidth of last two minutes, unit is Byte/s.
* `satellite_link_host` - Satellite Link hostname of the location.
* `status` - Enabled/Disabled.	
* `ws_endpoint` - The ws endpoint of the location.		

