---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_datacenter"
description: |-
  Manages a Datacenter in the Power Virtual Server cloud.
---

# ibm_pi_datacenter

Retrieve information about a Power Systems Datacenter.

## Example Usage

```terraform
data "ibm_pi_datacenter" "datacenter" {
  pi_datacenter_zone= "dal12"
}
```

### Notes

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
  
## Argument Reference

Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Optional, String) The GUID of the service instance associated with an account. Required if private datacenter.
- `pi_datacenter_zone` - (Optional, String) Datacenter zone you want to retrieve. If no value is supplied, the `zone` configured within the IBM provider will be utilized.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `capability_details` - (List) Additional Datacenter Capability Details.

  Nested schema for `capability_details`:
  - `disaster_recovery` - (List) Disaster Recovery Information.

      Nested schema for `disaster_recovery`:
        - `replication_services`- (List) Replication services.

          Nested schema for `replication_services`:
          - `asynchronous_replication` - (List) Asynchronous Replication Target Information.

                Nested schema for `asynchronous_replication`:
                  - `enabled` - (Boolean) Service Enabled.
                  - `target_locations` - (List) List of all replication targets.

                      Nested schema for `target_locations`:
                      - `region` - (String) regionZone of replication site.
                      - `status` - (String) the replication site is `active` or `down`.
          - `synchronous_replication` - (List) Synchronous Replication Target Information.

                  Nested schema for `synchronous_replication`:
                  - `enabled` - (Boolean) Service Enabled.
                  - `target_locations` - (List) List of all replication targets.

                      Nested schema for `target_locations`:
                      - `region` - (String) regionZone of replication site.
                      - `status` - (String) the replication site is `active` or `down`.

  - `supported_systems` - (List) Datacenter System Types Information.

      Nested schema for `supported_systems`:
        - `dedicated` - (List) List of all available dedicated host types.
        - `general` - (List) List of all available host types.

- `pi_datacenter_capabilities` - (Map) Datacenter Capabilities. Capabilities are `true` or `false`.

    Some of `pi_datacenter_capabilities` are:
  - `cloud-connections`, `disaster-recovery-site`, `metrics`,  `power-edge-router`, `power-vpn-connections`

- `pi_datacenter_href` - (String) Datacenter href.
- `pi_datacenter_location` - (Map) Datacenter location.

    Nested schema for `pi_datacenter_location`:
  - `region` - (String) Datacenter location region zone.
  - `type` - (String) Datacenter location region type.
  - `url`- (String) Datacenter location region url.
- `pi_datacenter_status` - (String) Datacenter status, `active`,`maintenance` or `down`.
- `pi_datacenter_type` - (String) Datacenter type, `off-premises` or `on-premises`.
