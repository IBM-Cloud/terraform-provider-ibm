# 2.4.0 (July 1, 2026)

## Bug Fixes

### Cloud Databases
* Feat 2857:  Implement Gen2 Support for `ibm_database` datasource ([6802](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6802))
* Allow deletion of disabled database instances with 0 members ([6872](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6872))

### IAM
* update error format for a number of cases ([6833](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6833))
* Update error handling for ServiceId ([6839](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6839))

### Kubernetes
* ROKS VNI support related resource and datasources are added ([6875](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6875))

### Power Systems
* Add Deprecated Attribute For SAP Profiles ([6847](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6847))
* Add missing Arg_SourceChecksum to image import ([6849](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6849))
* Allow Remote Restart ([6850](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6850))
* Trusted Profile & MDS ([6851](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6851))
* Fix unit wording ([6876](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6876))
* Enable DHCP Private ([6861](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6861))
* Add vpmem volume update functionality to ibm_pi_instance_vpmem_volume resource ([6870](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6870))


## Enhancements

### Cloud Databases
* Add backup_id support for Gen2 instances ([6869](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6869))
* Feat 2857: Implement Gen2 Support for `ibm_database_connection` datasource ([6808](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6808))

### KMS
* KeyProtect Dedicated Release ([6838](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6838))

### Kubernetes
* Add ibm_container_vpc_bare_metal_worker_reload action ([6865](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6865))

### Power Systems
* Stop skipping first two ip addresses when creating a network in terraform ([6776](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6776))
* Update wording in source-crn option to remove mention of service broker ([6874](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6874))
* Refactor DHCP Resource ([6845](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6845))


## Documentation

### Cloud Logs
* Fix endpoint_type allowed values for Event Notifications integration ([6871](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6871))

### Global Catalog
* issue https://github.ibm.com/ibmcloud/content-catalog/issues/6190 - document cm_validation override values and update acceptance test ([6862](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6862))

### Power Systems
* Remove VSCSI From Doc ([6846](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6846))




# 2.3.0 (June 19, 2026)

## Enhancements

### Catalog Management
* Expose validation.target info in ibm_cm_validation resource and ibm_cm_version datasource ([6830](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6830))


## Documentation

### Transit Gateway
* fix Transit Gateway docs missing opening front matter delimiter ([6863](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6863))


# 2.3.0-beta0 (June 15, 2026)

## Bug Fixes

### Cloud Internet Services
* handle deleted domain gracefully ([6843](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6843))

### IAM
* IAM identity: trusted profiles update ([6793](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6793))
* IAM Identity ServiceId Group support ([6795](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6795))
* IAM Identity: Fix consistency of CR type values ([6829](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6829))

### Kubernetes
* Support network plugin selection at VPC cluster creation ([6784](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6784))
* Support cluster offering selection at VPC cluster creation ([6842](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6842))

### VPC Infrastructure
* Added resource group name and documentation of auto-delete ([6810](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6810))
* use correct error variable in TerraformErrorf calls in vpc service ([6820](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6820))
* fixed the import section for is_* documents ([6814](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6814))
* fixed indentation in is_* resource docs ([6819](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6819))
* remove vpc option for access_control_mode in is_share resource ([6832](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6832))
* updated the vpc ike ipsec tests and documenation ([6856](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6856))


## Enhancements

### Catalog Management
* added private endpoint support for Catalog Management service ([6831](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6831))

### Event Notifications
* resource and data source update to support Email Sandbox and enable view notification payload setting ([6742](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6742))

### General
* Pull in latest ICR SDK version ([6834](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6834))

### ODF
* ODF 4.20 and 4.21 initial support ([6790](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6790))

### Partner Center Sell
* adding metadata.other.location_proxied_by to deployments, default values for cbr, also new parameter for them called event_publishing on the IAM service object ([6844](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6844))

### Power Systems
* Remove specific machine references ([6840](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6840))
* Update Power Go Client to V1.16.0 ([6848](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6848))

### Schematics
* Add suport for Schematics KMS Setting read ([6855](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6855))

### CD Toolchain
* Remove eu-es and jp-osa from deprecation warnings ([6822](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6822))

### VPC Infrastructure
* added multi algo support for vpc ike ipsec policies ([6769](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6769))
* Multivolume vpc snapshot ([6852](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6852))



# 2.2.2 (May 27, 2026)

## Enhancements

### Cloud Databases
* Add support for Valkey in ibm_database ([6786](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6786))

### VPC Infrastructure
* handle UPDATE_PENDING state in security group target deletion ([6809](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6809))

### Cloud Internet Services
* Add batch DNS records ([6749](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6749))

### PlatformNotifications
* remove beta notice ([6801](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6801))

### CDTektonPipeline
* bump CD Go SDK version ([6745](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6745))

### General
* bump github.com/moby/spdystream from 0.5.0 to 0.5.1 ([6748](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6748))
* bump go.opentelemetry.io/otel from 1.39.0 to 1.41.0 ([6758](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6758))
* Add mumbai region support for metrics-router ([6811](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6811))


# 2.2.1 (May 21, 2026)

## Enhancements

### Schematics
* secure values shouldnt be set in read ([6792](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6792))
* Addressed issue with sch workspace resource update ([6804](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6804))


# 2.2.0 (May 20, 2026)

## Bug Fixes

### VPC Infrastructure
* Support VSI downsize across families along with volume bandwidth ([6794](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6794))



# 2.2.0-beta1 (May 18, 2026)

## Bug Fixes

### Cloud Databases
* Revert "Feat: 2857 Implement Gen2 Support for ibm_database datasource ([6788](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6788))


# 2.2.0-beta0 (May 18, 2026)

## Bug Fixes

### Cloud Internet Services
* add query params for GET Custom List Items ([6761](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6761))


## Enhancements

### Cloud Databases
* Add Gen2 support for IBM Cloud Databases ([6714](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6714))
* 2857 Implement Gen2 Support for ibm_database datasource ([6783](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6783))

### Cloud Logs
* Add support for log data retention tags API ([6753](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6753))

### IAM
* enable cross region ([6782](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6782))
* add missing arguments docs and examples ([6785](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6785))


## Documentation

### General
* Update the readme ([6774](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6774))

### DRAutomation
* updated dra and pha docs ([6780](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6780))



# 2.1.0 (May 5, 2026)

## Bug Fixes

### Backup/Recovery
* add support for service name and protection source refresh ([6755](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6755))

### Cloud Internet Services
* add missing domain settings ([6765](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6765))

### VPC Infrastructure
* added subnets info in is-vpe ([6767](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6767))


## Enhancements

### General
* Update networking-go-sdk to v0.53.4 for logpush jobs fix ([6764](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6764))
* address go vet issues and fmt formatting for Go 1.26.1 ([6768](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6768))
* Add mumbai atracker region ([6766](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6766))
* add go mod tidy, go fmt, and go vet checks to CI workflow ([#6770](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6770)))

### Configuration Aggregator
* Dra pha apis ([6696](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6696))

### VPC Infrastructure
* added support for additional ipv4 protocols ([6751](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6751))


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

### General
* removed redundant checks to fix build on 1.26 ([6682](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6682))

### CIS
* fix mtls empty hostmanes issue ([6669](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6669))

### Global Tagging
* Fix tagging is error ([6681](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6681))


## Enhancements

### CD Tekton Pipeline
* Fix multiline properties ([6680](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6680))

### Cloud Databases
* added a nil check on response in ibm_database ([6660](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6660))

### Cloud Object Storage
* Updating the documentation for COS for cross account access. ([6685](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6685))

### IAM
* Terraform should not plan update-in-place when no resources exist for template assignment ([6216](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6216))

### Power Systems
* Refactor host group update logic ([6677](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6677))

### Toolchain
* Continuous Delivery (CD): Region discontinuation warnings ([6687](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6687))

### DRAutomation

## ENHANCEMENTS
- Added `events` attribute to ibm_pdr_get_events data source

## DEPRECATIONS
- `event` attribute is deprecated, use `events` instead


# 1.88.3 (February 23, 2026)

## Bug Fixes

### Backup/Recovery
* update connection schema ([6645](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6645))


## Enhancements

### General
* updated code owners ([6674](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6674))

# 1.88.2 (February 13, 2026)

## Bug Fixes

### CIS
* Fix instance ruleset rule handling to enable OWASP logging ([6644](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6644))

### Power
* Update unit wording to conform to new standard ([6652](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6652))

### Activity tracker
* remove managed by example ([6659](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6659))

### IAM
* Fix: prevent duplicate user invite API calls causing 409 error ([6665](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6665))

# 1.88.1 (February 10, 2026)

## Bug Fixes

### Cloud Object Storage
* Fix the docs related to COs backup vault policies ([6650](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6650))

### VPC Infrastructure
* fix(instance-group): fixed error on instance group wait ([6655](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6655))

## Enhancements

### IAM
* Add expires_at to service_api_key ([6654](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6654))


# 1.88.0 (February 6, 2026)

## Bug Fixes

### Catalog Management
* fix import of ibm_cm_offering ([6636](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6636))

### Cloud Logs
* change inclusion_filters to optional ([6649](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6649))

### IAM
* Fix policy creation when resourceType is set to 'resource-group' ([6621](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6621))


## Enhancements

### Code Engine
* add support for code engine pds, hmac secrets and trusted profiles ([6610](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6610))

### Schematics
* extend template type validation regex ([6593](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6593))

### Cloud Databases
* Add `async_restore` field for fast PG restore ([6630](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6630))
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

