# 2.1.0-beta0 (April 28, 2026)

## Bug Fixes

### Cloud Internet Services
* Allow import of CIS instances with EOM plans ([6728](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6728))

### Power Systems
* Fix image docs to use latest versions of the CLI commands ([6759](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6759))

### Satellite
* in satellite cluster resource, return if getcluster or getworkerpool fails ([6743](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6743))

### VPC Infrastructure
* fix sg, nacl rule protocol migration ([6695](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6695))
* fixed the documents on hyperlink ([6757](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6757))


## Enhancements

### Activity Tracker
* Add atracker app-config target support ([6750](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6750))

### Cloud Object Storage
* Adding the support for objectlock governance mode ([6715](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6715))

### General
* Pulling in the latest networking-go-sdk version ([6760](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6760))

### Power Systems
* Read crash on power volumes datasource ([6754](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6754))

### VPC Infrastructure
* update ibm_container_vpc_cluster patch_version doc ([6752](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6752))
* added instance profile zones availability changes ([6544](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6544))
* support for vpc Image partial availability ([6723](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6723))
* added support for volume jobs, vpc volume migration ([6717](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6717))


# 2.0.2 (April 16, 2026)

## Bug Fixes

### VPC Infrastructure
* fix: set vpn_gateway during import for ibm_is_vpn_gateway_connection ([6732](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6732))

# 2.0.1 (April 14, 2026)

## Bug Fixes

### Cloud Logs
* fix syntax type plan change ([6725](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6725))

### Configuration Aggregator
* Have single source of truth for ICR urls ([6739](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6739))

### Power Systems
* Fix-pi route next hop update ([6741](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6741))

### Transit Gateway
* [TGW] Add documentation for default_prefix_filter ([6734](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6734))


## Enhancements

### Event Notifications
* Support for bounce metrics Data source and metrics, subscription-id filter for metrics ([6594](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6594))

### General
* bump goreleaser/goreleaser-action from 6 to 7 ([6679](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6679))
* bump google.golang.org/grpc from 1.79.2 to 1.79.3 ([6731](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6731))
* bump crazy-max/ghaction-import-gpg from 6.1.0 to 7.0.0 ([6727](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6727))
* bump github.com/go-jose/go-jose/v4 from 4.1.3 to 4.1.4 ([6730](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6730))

### Container Registry
* Pull in latest icr SDK version ([6729](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6729))


# 2.0.0 (March 30, 2026)
This release adds internal support for Terraform's Plugin Framework alongside our existing implementation. All your existing Terraform configurations work exactly as before with zero changes required.
We're bumping to v2.0.0 to signal an important internal architectural enhancement: the provider now supports Terraform Plugin Framework in addition to the existing SDKv2 implementation. This is a major milestone that enables future capabilities while maintaining complete backward compatibility.

## Bug Fixes

### Cloud Internet Services
* handle null timestamps in CIS custom page resource ([6702](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6702))

### Cloud Logs
* add Computed to alert filter_type and incident_settings fields ([6716](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6716))

### General
* Fix the missing read env for provider ([6704](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6704))

### Partner Center Sell
* Add id to composite children plan updatable flag is editable  parnercentersellv1 ([6698](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6698))

### Power Systems
* CCNA Error Msg Refactor ([6711](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6711))
* Capture Error Msg Refactor ([6712](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6712))

### Transit Gateway
* handle create timeout by checking existing connec… ([6703](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6703))

### VPC Infrastructure
* Update size validation range for RFS ([6709](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6709))
* fixed document for is_images datasource ([6689](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6689))


## Enhancements

### General
* Plugin Framework support ([6611](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6611))

### Cloud Databases
* Refactor ibm_database resource and datasource to introduce backend abstraction ([6667](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6667))

### Cloud Logs
* Logs Routing Default Private CSE endpoint and write_status ([6710](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6710))
* Update Views API for PR #164 ([164](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/164))
* Add IBM Cloud Logs Extensions API support ([6700](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6700))

### Code Engine
* Add ibm_code_engine_build_run action with Plugin Framework support ([6611](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6611))

### Configuration Aggregator
* Add IBM Account Management API Support ([6701](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6701))

### Power Systems
* Add externalIP attribute for network_interface ([6705](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6705))
* Add Asaps To SAP Profile/s ([6708](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6708))

### Secrets Manager
* Private path support for Code Engine ([6699](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6699))
* SM event notification datasource should not fail when no registration exists ([6692](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6692))

### Transit Gateway
* support Secrets Manager CRNs in authentication_key ([6713](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6713))

