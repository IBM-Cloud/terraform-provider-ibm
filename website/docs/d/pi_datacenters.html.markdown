---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_datacenters"
description: |-
  Manages Datacenters in the Power Virtual Server cloud.
---

# ibm_pi_datacenters

Retrieve information about Power Systems Datacenters.

## Example usage

The following example retrieves information about Power Systems Datacenters.

```terraform
data "ibm_pi_datacenters" "datacenters" {}
```

## Notes

- Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
- If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  - `region` - `lon`
  - `zone` - `lon04`

Example usage:

  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```

## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `datacenters` - List of Datacenters

  Nested schema for `datacenters`
  - `pi_datacenter_capabilities` - (Map) Datacenter Capabilities.

    Nested schema for `pi_datacenter_capabilities`; are (Bool) `true` or `false`:
    - `cloud-connections` - (Bool) Cloud-connections capability.
    - `disaster-recovery-site` - (Bool) Disaster recovery site.
    - `power-edge-router` - (Bool) Power edge router capability.
    - `vpn-connections`- (Bool) VPN-connections capability.
    - `sysdig-enabled`- (Bool) Sysdig-enabled capability.

  - `pi_datacenter_location` - (Map) Datacenter location.

    Nested schema for `pi_datacenter_location`:
    - `region` - (String) The Datacenter location region zone.
    - `type` - (String) The Datacenter location region type.
    - `url`- (String) The Datacenter location region url.
  - `pi_datacenter_status` - (String) The Datacenter status, `ACTIVE`,`MAINTENANCE` or `DOWN`.
  - `pi_datacenter_type` - (String) The Datacenter type, `Public Cloud` or `Private Cloud`.
