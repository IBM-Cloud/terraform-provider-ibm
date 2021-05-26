## 1.25.0 (May18, 2021)
FEATURES:
* Support Resource Tag Management
    - **Resources**
        - ibm_resource_tag
    - **DataSources**
        - ibm_resource_tag
* Support VPC dedicated host disk management
    - **Resources**
        - ibm_is_dedicated_host_disk_management
    - **DataSources**
        - ibm_is_dedicated_host_disk
        - ibm_is_dedicated_host_disks
* Support IAM User API Key
    - **Resources**
        - ibm_iam_api_key
    - **DataSources**
        - ibm_iam_api_key
* Support VPC endpoint target gateways
    - **DataSources**
        - ibm_is_endpoint_gateway_targets
* Support VPC security group target management
    - **Resources**
        - ibm_is_security_group_target
    - **DataSources**
        - ibm_is_security_group_target
        - ibm_is_security_group_targets
* Support COS Bucket Object
    - **Resources**
        - ibm_cos_bucket_object
    - **DataSources**
        - ibm_cos_bucket_object
* Support container Registry Retention Policy
    - **Resources**
        - ibm_cr_retention_policy
	

ENHANCEMENTS
* Add the capabilities for offering speeds which provides the bmetered and unmetered
billing options for that offering speed ([#2584](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2584))

* Support for dedicated host/group in instance template and instance ([#2579](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2579))

* location, managed_from and resource_group_id mark these attibutes in ibm_satellite_location either DiffSuppress or throw back an error ([#2567](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2567))

* Add max allowed sessions to account settings ([#2610](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2610))

* Add vcpus and memory to instance profiles data source ([#2492](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2492))

* Support creating IAM policies with operator ([#2533](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2533))

* Mark ibm_container_cluster_config `admin_certificate`, `ca_certificate`, `token` attributes as sensitive ([#2622](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2622))

* Bump up go version: 1.16 ([#2600](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2600))


BUGFIXES
* Fix the deletion of instance group due to load balancer status ([#2547](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2547))

* No way to make a vpc_address_prefix default on vpc when using manual address_prefix_management ([#2282](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2282))

* ibm_is_image data source silently fails if the image is not available ([#2587](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2587))

* ibm_is_image data source does not provide a warning if the image is deprecated ([#2588](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2588))

* update azure script to grow root volume group ([#2621](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2621))

* instance group manager policy synchronization ([#2635](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2635))

* Fix the db task timeout ([#2607](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2607))

* KMS keys created with endpoint_type = "public" regardless of the actual setting ([#2482](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2482))

* Error finding VLAN order: couldn't find resource (21 retries) ([#2613](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2613))

* Terraform import for routingtable fails ([#2580](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2580))


## 1.24.0 (May04, 2021)
FEATURES:
* Support VPC instance disk management
    - **Resources**
        - ibm_is_instance_disk_management
    - **DataSources**
        - ibm_is_instance_disk
        - ibm_is_instance_disks
	
ENHANCEMENTS
* Support resize of VPC instance ([#2448](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2448))
* Support Load balancer Parameter based routing ([#2518](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2516))
* Support horizontal scaling on database with new arguments node_count, node_memory_allocation_mb, node_disk_allocation_mb, node_cpu_allocation_count ([#2313](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2313))
* Support request_metrics_enabled for COS Bucket metric monitoring ([#2530](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2530))
* Support virtual endpoint gateway as target to subnet reserved IP ([#2521](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2521))

BUGFIXES
* Creating ibm_pi_key fails everytime with context deadline exceeded ([#2527](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2527))
* Fix diff on resource key parameters ([#2182](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2182))
* Fails to create PTR records causing Terraform crash ([#2535](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2535))
* Fix crash for VPC instance group manager ([#2554](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2554))
* VPC network ACL rule ICMP does not set type ([#2559](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2559))
* Conflict with exec.image and exec.code/exec.code_path (can't use custom docker images) ([#2556](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2556))

## 1.23.2 (Apr20, 2021)
ENHANCEMENTS
* Add support for COS retention policy ([#1880](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1880))
* Add support for private_address for VPN gateway ([#2282](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2382))
* List all certificates in a certificate manager instance ([#2358](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2358))
* Enhance description for attribute reference ([#2475](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2475))
*  Add support for regional ca-tor COS bucket ([#2483](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2483))
BUGFIXES
* Fix the broken links for classic infrastructure bare metal ([#2481](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2481))
* Fix cis primary certificate crash ([#2490](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2490))
* Fix ibm_satellite_location: cannot specify resource group ([#2499](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2499))
* Fix ibm_satellite_location resource doesn't work correctly to ensure that resource is created / deleted appropriately ([#2497](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2497))
* Fix invalid example for ibm_iam_account_settings ([#2484](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2484))
* Fix the documentiaon for VPC reserved IP ([#2512](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2512))

## 1.23.1 (Apr07, 2021)
ENHANCEMENTS
* Add support to retry the update of patch version ([#2379](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2379))
* Add gateway_connection argument for VPC VPN gateway Connection ([2270](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2270))

BUGFIXES
* Fix the crash for resource key ([#2462](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2462))
* Change the order to place to use billing_order ([#554](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/554))

## 1.23.0 (Apr02, 2021)
FEATURES:
* Support Catalog Management
    - **Resources**
        - ibm_cm_offering_instance
        - ibm_cm_catalog
        - ibm_cm_offering
        - ibm_cm_version
    - **DataSources**
        - ibm_cm_catalog
        - ibm_cm_offering
	    - ibm_cm_version
		- ibm_cm_offering_instance
* Support IAM Account Management
    - **Resources**
        - ibm_iam_account_settings
    - **DataSources**
    	- ibm_iam_account_settings

* Support Enterprise Management
    - **Resources**
        - ibm_enterprise
        - ibm_enterprise_account_group
        - ibm_enterprise_account
    - **DataSources**
    	- ibm_enterprises
        - ibm_enterprise_account_groups
        - ibm_enterprise_accounts

BUGFIXES
* Fix the provision of classic Infrastructure VM to apply sshkeys, imageid and script ([#2448](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2448))
* Fix documentation updates ([#2443](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2443))
* Fix Dedicated host with status 'failed' throws error during destroy ([#2443](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2446))
* Fix while creating a DL Connect gateway do not wait for gateway to be provisioned for few providers ([#2458](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2458))

## 1.22.0 (Mar30, 2021)

FEATURES:

* Support VPC dedicated hosts
    - **Resources**
        - ibm_is_dedicated_host
		- ibm_is_dedicated_host_group
    - **DataSources**
        - ibm_is_dedicated_host
		- ibm_is_dedicated_hosts
		- ibm_is_dedicated_host_profile
		- ibm_is_dedicated_host_profiles
		- ibm_is_dedicated_host_group
		- ibm_is_dedicated_host_groups
* Support VPC reserved IP
     - **Resources**
        - ibm_is_subnet_reserved_ip
     - **DataSources**
        - ibm_is_subnet_reserved_ip
        - ibm_is_subnet_reserved_ips
* Support Push Notification chrome web
    - **Resources**
        - ibm_pn_application_chrome
    - **DataSources**
        - ibm_pn_application_chrome
        
* Support Key Management Alias and Rings
    - **Resources**
        - ibm_kms_key_alias
        - ibm_kms_key_rings
    - **DataSources**
        - ibm_kms_key_rings

* Support for reading secrets from IBM Cloud Secrets Manager
    - **DataSources**
        - ibm_secrets_manager_secrets
        - ibm_secrets_manager_secret

* Support Schematics
    - **Resources**
        - ibm_schematics_workspace
        - ibm_schematics_action
        - ibm_schematics_job
    - **DataSources**
        - ibm_schematics_action
        - ibm_schematics_job

* Support Observability
     - **Resources**
        - ibm_ob_logging
        - ibm_ob_monitoring

* Support for Satellite 
    - **Resources**
        - ibm_satellite_location
        - ibm_satellite_host
    - **DataSources**
        - ibm_satellite_location
        - ibm_satellite_attach_host_script

* Support for CIS Cache setting
    - **Datasource**
        - ibm_cis_cache_settings

PROVIDER

* Support `visibility` argument to control the visibility to IBM Cloud endpoint.

* `generation` argument is depreated. By default the provider targets to IBM Cloud VPC Infrastructure.


ENHANCEMENTS

* Support added DH group 19 and sha 512 for IKE and IPSec Policy([#2361](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2361))

* Support `delegate_vpc` action for VPC routing table ([#2355](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2355))

* Support tags for IBM Cloud VPC security group ([#2353](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2353))

* Support for renaming of default Network ACL, Security Group and Routing Table ([#2216](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2353))

* Support for allow control of Security Groups on Load Balancer ([#2324](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2324))

* Support recover the public key from an SSH key in IBM Cloud VPC ([#2388](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2388))

* Support ibm_is_vpc_routing_table_route to accept a VPN connection ID as next_hop ([#2270](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2270))

* Support to provision classic Infrastructure Virtual instance using quoteID ([#2433](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2433))

* Support `serve_stale_content` for CIS Cache settings ([#2219](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2219))

* Support filtering of subnets based on metro for IKS kubernetes Cluster ([#2403](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2403))

* Support `alias` argument to filter the keys in ibm_kms_key and ibm_kms_keys datasources and aliases attribute ([#2293](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2293))

* Support `key_ring_id` in ibm_kms_key resource and datasources ([#2378](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2378))

BUGFIXES

* Fix increase in panics while refreshing resources for cos_bucket ([#2373](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2373))

* Fix Cloud Function action runtimes version ([#2424](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2424))

* Fix COS buckets allow modifying key_protect after creation, but they should not ([#2310](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2310))


## 1.21.2 (Mar15, 2021)

PROVIDER:
* Updgrade Terraform SDK to v2

FEATURES:

* Support checksum argument for VPC Images ([#2227](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2161))

* Support iam_id argument cross Account iam_service_policy ([#2331](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2331))

* Support tags argument for VPC subnet ([#2321](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2332))

* Support tags argument for VPC network acl ([#2343](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2343))

* Add resource schema timeouts for classic infrastructure compute VM ([#2291](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2291))

* Support accept_proxy_protocol argument for vpc loadbalancer listener ([#2325](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2325))


BUGFIXES

* Fix logging not supported for VPC Network Loadbalancer ([#2332](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2332))

* Fix addons not being enabled post-cluster creation ([#2346](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2346))

* Fix ibm_iam_user_policy data source produces no results ([#2312](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2312))


## 1.21.1 (Mar03, 2021)

FEATURES:

* Support sort argument for IAM service policies and IAM user policies ([#2227](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2227))

* Support default_routing_table attribute for VPC resource and datasource ([#2286](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2286))

* Support logging argument for VPC load balancer ([#2228](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2228))

* Add transactionID for IAM authentication error messages ([#2304](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2304))

BUGFIXES

* Fix the provision of instance template with boot volume ([#2205](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2205))

* Fix ibm_resource_key is not tainted by change to role ([#2182](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2182))

* Fix ibm_cis, ibm_database, ibm_resource_instance not tainted by change to resource groupID ([#2297](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2297))

* Fix ibm_container_addons resource not detecting version change ([#2295](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2295))

* Fix error when trying to use data source ibm_container_cluster for existing lite IKS ([#2300](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2300))

## 1.21.0 (Feb12, 2021)

FEATURES:

* Support datasource for VPC volume profiles `ibm_is_volume_profile`, `ibm_is_volume_profiles`

* Support datasource to list power instance catalog images `ibm_pi_catalog_images`

ENHANCEMENTS

* Support `lunid` attribute for block classic block storage ([#1491]https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1491)

* Support `HTTP_COOKIE` session_affinity for lbass ([#2218](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2218)) 

* Support `auto_delete_volume` argument to delete data volumes of VPC instance [#646](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/646)) 

* Support `wait_till` argument for ibm_containar_cluster to control the behaviour of waiting for cluster ([#2232]https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2232)

* Enable retries on authnetication failures ([#2248]https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2248)


BUGFIXES

* Fix the nil pointer exception on lbs of vpc_cluster destroy ([#2226](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/22276)) 

* Fix the nil pointer exception on is_instance_group ([#2247](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2247)) 

* Fix the nil pointer on iam_service_api_key ([#2259](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2259)) 

* Fix the validation for LB listener policy rule ([#2257](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2257)) 

* Fix the patch_version update for kubernetes clusters ([#2217](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2217)


## 1.20.1 (Jan27, 2021)

BUGFIXES
Fix the regression issue provisioning of contianer clusters ([#2206](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2206)) 

## 1.20.0 (Jan25, 2021)

FEATURES:

* Support directlink provider gateway resource `ibm_dl_provider_gateway`

* Support directlink provider gateways, ports datasource `ibm_dl_provider_gateways`, `ibm_dl_provider_ports`

ENHANCEMENTS

* Support provision file storage size from 10TB to 13TB ([#2158](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2158))

* Support for pod-subnet and service-subnet for `ibm_container_cluster` resource ([#1196](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1196))

* Support the ability to retrieve the instances in a specific VPC ([#1961](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1961))

* Support for patch update for cluster worker nodes ([#1978](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1978)) 

* Support architecture attribute for VPC instance profiles ([#2002](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2002)) 

BUGFIXES

* Fix the nil pointer exception for cos bucket import scenario ([#2151](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2151))

* Fix Transit gateway Connection creation fails for cross account ([#2170](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2170))

* Fix Terraform crash when subnet not found ([#2058](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2058))

* Fix VPC LB creation with count greater than 1 ([#2168](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2168))

# 1.19.0 (Jan08, 2020)

FEATURES:
* Support Contianer Registry resource and datasource `ibm_cr_namespace`, `ibm_cr_namespaces` ([#2119](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2119))

* Support reset APIkey for cluster `ibm_container_api_key_reset` ([#2118](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2118))

* Support APIkey for serviceID `ibm_iam_service_api_key` ([#666](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/666))

ENHANCEMENTS:

* Move next_hop from optional to required for `ibm_is_vpc_routing_table_route` ([#2141](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2141))

* Support jp-osa endpoints for `ibm_cos_bucket` [#2149](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2149))

* Support crn in target attribute for `ibm_is_virtual_endpoint_gateway` ([#2147](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2147))

## 1.18.0 (Dec22, 2020)

FEATURES:

* Support CIS Certificate resources `ibm_cis_certificate_order`, `ibm_cis_certificate_upload`

* Support CIS Certificate datasources `ibm_cis_certificates`, `ibm_cis_custom_certificates`

* Support CIS DNS Records import and export `ibm_cis_dns_records_import`, `ibm_cis_dns_records`

* Support virtual private endpoint gateways `ibm_is_virtual_endpoint_gateway`, `ibm_is_virtual_endpoint_gateway_ip` resources and `ibm_is_virtual_endpoint_gateways`, `ibm_is_virtual_endpoint_gateway_ips`, `ibm_is_virtual_endpoint_gateway` datasources


ENHANCEMENTS:

* resource: Support `labels` argument and updates for kubernetes clusters ([#2109](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2109)) 

* resource: Support `namespace` argument and `persistence` and `status` attributes ([#2097](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2097)) 



BUGFIXES:

* Fix an ibm_resource_key that is removed outside of Terraform is not being recreated ([#2125](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2125))

* Fix cluster addon fails on apply after timeout ([#2129](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2129))


## 1.17.0 (Dec10, 2020)

FEATURES:

* Support CIS WAF resources `ibm_cis_range_app`, `ibm_cis_waf_package`, `ibm_cis_waf_group`, `ibm_cis_waf_rule`

* Support CIS WAF datasources `ibm_cis_range_apps`, `ibm_cis_waf_packages`, `ibm_cis_waf_groups`, `ibm_cis_waf_rules`

ENHANCEMENTS:

* resource: Support `force_delete` argument for ibm_cos_bucket ([#2017](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2017)) 

* resource: Move `bgp_base_cidr` as optiona argument ([#2087](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2087)) 

* resource: Support `expire_rule` argument for ibm_cos_bucket ([#1590](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1590))

* resource: Add validate function for resoure_instance_id argument to ibm_cos_bucket ([#2103](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2103))

* resource: Support Route and Profile based VPN gateways ([#2094](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2094))


BUGFIXES:

* Fix users not found when adding to access group ([#2034](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2034))

* Set the `instance_id` in ibm_kms_key ([#2106](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2106))

* Fix multiple cis_domain leads into inconsistency ([#2086](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2086))


## 1.16.1 (Dec02, 2020)

BUGFIXES:

* Fix issue when trying to delete a ibm_container_alb_cert ([#2067](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2067))


## 1.16.0 (Nov30, 2020)

FEATURES:

* Support VPC Routing Table `ibm_is_vpc_routing_table` and VPC Routing Table Route  `ibm_is_vpc_routing_table_route` resources ([#1395](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1395))

* Support `ibm_is_vpc_default_routing_table`, `ibm_is_vpc_routing_tables` and `ibm_is_vpc_routing_table_routes` datasources ([#1395](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1395))


ENHANCEMENTS:

* resource: Extend CIS firewall resource to support `access_rules` and `ua_rules` ([#2025](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2025)) 

* resource: Support anti-spoofing `allow_ip_spoofing` for ibm_is_instance ([#1396](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1396)) 

* resource: Support `routing_table` and `ip_version` agruments for ibm_is_subnet ([#1395](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1395))

* data: Support `policies` attribute for ibm_kms_keys and ibm_kms_key ([#1928](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1928))

* resource: Support `number_of_invited_users` and `invited_users` attribute for ibm_iam_user_invite ([#2053](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2053))

BUGFIXES:

* Fix the upgrade of kube_version for master and worker nodes of cluster ([#1952](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1952))

* Fix issue when trying to provision a new ibm_container_alb_cert ([#2067](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/2067))

## 1.15.0 (Nov24, 2020)

FEATURES:

* Support for subnet network interface attachment `ibm_is_subnet_network_acl_attachment` resource ([#1941]

* Support `ibm_cis_routing` resource ([#1991](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1991))

* Support `ibm_cis_cache_settings` resource ([#1995](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1995))

* Support `ibm_cis_global_load_balancers` datasource ([#1981](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1981))

* Supoort `ibm_cis_custom_page`resource and `ibm_cis_custom_pages` datasoruce ([#1997](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1997))

* Support `ibm_dns_glb`, `ibm_dns_glb_monitor`, `ibm_dns_glb_pool` resource for IBM Cloud PDNS Service ([#1887](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1887))

* Support `ibm_dns_glbs`, `ibm_dns_glb_monitors`, `ibm_dns_glb_pools` datasources for IBM Cloud PDNS Service ([#1887](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1887))

ENHANCEMENTS:

* resource: Support `public_ip` attribute in ibm_pi_network_port resource ([#1930](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1930)) 

* resource: Support encrypted images `encrypted_data_key` and `encryption_key` arguments in ibm_is_image resource ([#1938](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1938)) 

* resource: Support archive rule `archive_rule` for ibm_cos_bucket ([#1950](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1950)) 

* resource: Support Polcies for ibm_kms_key ([#1928](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1928))


* data: Support `list_bounded_services` argument for ibm_container_cluster ([#2051](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2051))

BUGFIXES:

* Fix provision of cloud funciton resources ([#837](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/837))

* Fix the destroy of ibm_pi_instance wait ([#2047](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/2047))

## 1.14.0 (Oct28, 2020)

FEATURES:

* Support for subnet network interface attachment `ibm_is_subnet_network_acl_attachment` resource ([#1941](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1941))

* Support for CIS tls settigns `ibm_cis_tls_settings` resource ([#1954](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1954))

* Support `ibm_cis_origin_pools` datasource ([#1959](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1959))

ENHANCEMENTS:

* resource: Support additional domain settings (max_upload, cipher, minify, security_header, mobile_redirect, challenge_ttl, dnssec, browser_check) for `ibm_cis_domain_settings` ([#1939](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1939))

* resource: Support `expiration_date` argument to ibm_kms_key resource ([#1967](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1967))

* resource: Support `wait_for_worker_update` argument to `ibm_container_cluster` and `ibm_container_voc_cluster` to control the upgrade of worker nodes ([#1969](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1969))

* data: Support `cert_file_path` attribute to `ibm_database` ([#1985](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1985))

BUGFIXES

* Remove forcenew for IS instance group ([#1951](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1951))

* Support resource instance's parameter to be an array type ([#1953](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1953))

* Fix the tags attachemnt for ibm_databse resource ([#1971](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1971))

* Fix changing allowed_ip to a list of IPs to nothing leads to an error when configuring a COS bucket ([#1661](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1661))

* Fix the update of VPC worker nodes kube verison ([#1952](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1952))


## 1.13.1 (Oct07, 2020)

ENHANCEMENTS:

* resource: Support endpoint_type argument and endpoint environmental variable for COS bucket ([#1945](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1945))

* resource: Support Direct Link Connect Type([#1927](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1927))

* doc: Update supported parameters for Event Streams([#1946](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1946))

BUGFIXES

* Fix the nil pointer exception for transist gateway delete ([#1943](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1943))

## 1.13.0 (Oct01, 2020)

FEATURES:

**VPC NLB Feature**: 
* Support for provisioning of NLB load balancers ([#1937](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1937))
    * data/ibm_is_lb_profiles
    * data/ibm_is_lbs

**CIS Edge Functions**: 
* Support for CIS Edge Functions ([#1873](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1873))
    * resource/ibm_cis_edge_functions_actions
    * resource/ibm_cis_edge_functions_trigger
    * data/ibm_cis_edge_functions_actions
    * data/ibm_cis_edge_functions_triggers


ENHANCEMENTS:

* datasource: Support `pools` attribute for is_lb datasource ([#1895](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1895))

* resource: Support of renew certificate in certificates manager ([#1909](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1909))

* resource: Support update of parameters for resource instance ([#1705](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1705))

* resource: Support for NLB load balancers in VPC ([#1937](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1937))

* resource: Migrate HPCS endpoints to cloud.ibm.com domain ([#1932](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1932))

* resource: Support retry of VPC instance to recover from a perpetual "starting" or "stopping" state by using "force_recovery_time" argument ([#1934](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1934))

* resource: Support customer health check request headers for Cloud Internet Services ([#1844](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1844))

* resource: Support ibm_iam_service_id (data / resouce) should return iam_id ([#1820](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1820))

* resource: Support ICD Service endpoint doesn't exist for region: "che01" ([#1894](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1894))

BUGFIXES

* Fix the instance template destroy error ([#1886](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1886))

* Fix the delete of ibm_cdn resource ([#1925](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1925))

* Fix the provision of free cluster ([#1901](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1901))

* Fix the ibm_container_addons not working on other resource_group !=Default ([#1920](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1920))

* Fix the crash of ibm_is_subnet datasource with empty identifier ([#1933](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1933))

* Fix Instance Group/AutoScale Max count should be 1000 not 100 ([#1889](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1889))

* Fix ibm_pi_instance not failing on ERROR state ([#1879](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1879))

## 1.12.0 (Sep14, 2020)

FEATURES:

**VPC Flow Logs**: 
* Support for IBM Cloud VPC Flow Logs ([#1356](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1356))
    * resource/ibm_is_flow_logs
    * data/ibm_is_flow_logs

**VPC Auto Scale**: 
* Support for IBM Cloud VPC Auto Scale ([#1357](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1357))
    * resource/ibm_is_instance_group
    * resource/ibm_is_instance_group_manager
    * resource/ibm_is_instance_group_manager_policy
    * resource/ibm_is_instance_template
    * data/ibm_is_instance_group
    * data/ibm_is_instance_group_manager
    * data/ibm_is_instance_group_managers
    * data/ibm_is_instance_group_manager_policies
    * data/ibm_is_instance_group_manager_policy
    * data/ibm_is_instance_templates
    * data/ibm_is_instance_profiles
    * data/ibm_is_instance_profile

**Power Instance**:
* Support for IBM Cloud Power Instance network port attachement and snapshot ([#1867](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1867))
    * resource/ibm_pi_snapshot
    * resource/ibm_pi_network_port_attach

**Cluster Addons**:
* Support for addons for container cluster ([#721](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/721))
    * resource/ibm_container_addons
    * data/ibm_container_addons

* data/ibm_is_lb: Support for ibm_is_lb [#1849](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1849))

* data/ibm_container_alb: Support for ibm_container_alb [#1850](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1849))

* data/ibm_container_alb_cert: Support for ibm_container_alb_cert [#1850](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1849))

* data/ibm_container_bind_service: Support for ibm_container_bind_service [#1850](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1849))


ENHANCEMENTS:

* resource: Support key protect configuraton for Container clusters ([#673] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/673))

* resource: Support delete of PVC Storage for Container clusters([#1847] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1847))

BUGFIXES

* Fix the diff on classsic VM instance ([#1828] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1828))

## 1.11.2 (Sep 01, 2020)

ENHANCEMENTS:
* resource: Support ProxyFromEnvironment for honouring for HPCS service.
* resource: Support auto scaling for IBM Cloud database service.


## 1.11.1 (Aug 26, 2020)

ENHANCEMENTS:

* resource: Assign IP address to a VSI on provisioning for VPC ([#1830](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1830))

* resource: Support kr-seo region for database ([#1831](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1831))

BUGFIXES

* Fix ibm_is_vpc datasource for Gen1 ([#1834](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1834))

* Fix provision of ibm_pi_instance ([#1833](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1833))

* Fix provision of ibm_pi_network_port ([#1823](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1823))

## 1.11.0 (Aug 24, 2020)

FEATURES:

* data/ibm_is_public_gateway: Support for ibm_is_public_gateway ([#1745](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1745))
* data/ibm_is_floating_ip: Support for ibm_is_floating_ip ([#1794](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1794))

ENHANCEMENTS:

* resource: Allow configuration of the key used to encrypt IBM Cloud Databases backups ([#1761] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1761))

* Support customer managed volume encryption for VPC Nextgen ([#1673] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1673))

* Support virtual cores capability for power instance instance ([#1798] (https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1798))

* Support for interconnecting two ibm cloud functions by target_url ([#1526] (https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1526))

*Support for cross account Transist Gateway ([#1021] (https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1021))

BUGFIXES

* ibm_tg_gateway delete is not complete when it has reported deleted ([#1783] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1783))

* Fix the IAM IP address restriction for invited user ([#1780] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1780))

* VPC instance datasource failing to fetch the correct instance ([#1801] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1801))

* Fix ibm_is_network_acl Validate name input field of ACL Rules ([#1262] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1262))

* Fix for datasource ibm_cis_domain only retrieve 20 domains ([#1804] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1804))

* Fix iam_access_group_policy with addition of account_management doesn't apply ([#1551] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1551))

* Fix error received - Current user does not have access to team directory ([#1536] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1536))

## 1.10.0 (Aug 06, 2020)

FEATURES:

**Transist Gateway**: 
* Support for Trasist Gateway Service ([#1021](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1021))
    * resource/ibm_tg_gateway
    * resource/ibm_tg_connection
    * data/ibm_tg_gateway
    * data/ibm_tg_gateways
    * data/ibm_tg_locations
    * data/ibm_tg_location

**DirectLink Gateway**: 
* Support for DirectLink Gateway Service ([#1349](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1349))
    * resource/ibm_dl_gateway
    * resource/ibm_dl_virtual_connection
    * data/ibm_dl_gateways
    * data/ibm_dl_offering_speeds
    * data/ibm_dl_port
    * data/ibm_dl_ports
    * data/ibm_dl_gateway
    * data/ibm_dl_locations
    * data/ibm_dl_routers

**CloudFunction**		
* Support for Cloud Function Namespace ([#682](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/682))
    * resource/ibm_function_namespace
    * data/ibm_function_namespace

**KMS (keyprotect/hpcs crypto)**
* Support for Key management (key protect/HPCS Crypto Service) ([#1353](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1353))
    * resource/ibm_kms_key
    * data/ibm_kms_key
    * data/ibm_kms_keys

**IAM User Setting**
* Support for IAM User Management settings ([#1780](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1780))
    * resource/ibm_iam_user_settings
    * data/ibm_iam_users
    * data/ibm_iam_user_profile

**Event Stream**
* Support for IBM Event Stream Topic ([#1781](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1781))
    * resource/ibm_event_streams_topic
    * data/ibm_event_streams_topic

* data/ibm_is_security_group: Support for ibm_is_security_group ([#1223](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1223))

* data/ibm_container_worker_pool: Support for ibm_container_worker_pool ([#1751](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1751))

* data/ibm_container_vpc_cluster_worker_pool: Support for ibm_container_vpc_cluster_worker_pool ([#1773](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1773))

* data/ibm_container_vpc_cluster_alb: Support for ibm_container_vpc_cluster_alb  ([#1775](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1775))

* data/ibm_certificate_manager_certificate: Support for ibm_certificate_manager_certificate  ([#1679](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1679))

ENHANCEMENTS:

* resource: Support configure geo routes in cis_glb ([#985] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/985))

* data: Retrieve icd disk encryption details ([#1742] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1742))

* resource: auto_renew_enabled support for ibm_certificate_manager_order ([#1657] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1657))

* resource: ibm_container_alb_cert destroy not synchronous ([#1712] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1712))


BUGFIXES

* Fix ibm_container_vpc_worker_pool resource forces a replace if resource_group_id not set ([#1748] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1748))

* Fix IBM Container VPC Cluster with no default worker pool crashes on destroy ([#1733] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1733))

## 1.9.0 (July 22, 2020)

FEATURES:

* data/ibm_is_instances: Support for ibm_is_instances ([#1454](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1454))
* data/ibm_is_instance: Support for ibm_is_instance ([#1454](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1454))

ENHANCEMENTS:

* resource: Support ibm_function_action, ibm_function_package, ibm_function_trigger, ibm_function_rule resources for IAM and CF based namespace ([#837](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/837))
**Note** - The provider level argument `function_namespace` is deprecated.The namespace is a required argument part of the function resources. The users need to update the templates to add the `namespace` argument to function resources.

* resource: Support update of adding additional zones to VPC cluster and worker pool resource ([#1546] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1546))

* Support update_all_workers flag to control the update of workers for VPC clusters ([#1681] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1681))

* resource: Support extension attribute for ibm_resource_instance ([#1686] (https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1686))

* resource: Support dashboard_url attributes for ibm_resource_instance ([#1682] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1682))

* resource: Support for update of key_protect_key parameter in ibm_database ([#1622] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1622))

* Support for resource synchronization of private dns permitted network ([#1674] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1674))

* data/ibm_resource_instance: Support guid attribute for ibm_resource_instance datasource ([#1724] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1724))

* resource: Support label argument for default worker pool ibm_container_cluster  ([#775] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/775))

BUGFIXES

* ibm_cis_domain_settings does not allow for Standard plans ([#1623] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1623))

* Fix the update of attachment of public gateway to VPC subnet ([#1626] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1626))

* Fix RHCOS via ibm_pi_instance timeout waiting for networ ([#1620] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1620))

* Fix Gateway enabled cluster recreated on every apply ([#1706] (https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1706))


## 1.8.1 (June 30, 2020)

ENHANCEMENTS:

* datasource: Support for aggregation of VPC images([#1580](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1580))

BUGFIXES

* resource: Fix the destroy of virtual instance with volumes
* resource: Fix the mutex of pdns resource records([#1601](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1601))

## 1.8.0 (June 23, 2020)

FEATURES:

* New Datasource: ([ibm_iam_roles](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1433))
* New Datasource: ([ibm_iam_role_actions](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1433))

ENAHANCEMENTS:

*resource: Support for provisioning ROKS on VPC Gen2 ([#1437](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1437))
*resource: Support for intergating firewall IP, activity tracker, metric monitoring to COS bucket ([#1487](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1487))
*data: Add new attribute `output_json` for ibm_scheamtics_output ([#1413](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1413))

BUGFIXES
*resource: Fix the list of network rules attached to network acl([#1547](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1547))

## 1.7.1 (June 11, 2020)

ENHANCEMENTS:

* resource/ibm_cis_domain_settings: Support additional domain settings ([#1475](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1475))
* resource/resource_ibm_certificate_manager_order: Added key_algorithm to order certificate in CMS ([#1512](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1512))
* data/ibm_is_vpc: Add zone name to data source vpc.subnets outputs ([#1450](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1450))

BUG FIXES:

* docs: Add documentation for is_instance and is_security_groups( [#1522](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1522))
* data/ibm_is_vpc: Regression on vpc_source_addresses support ([#1530](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1530))
* data/ibm_container_vpc_cluster: container_vpc_cluster fails for OCP VPC cluster with ALB error ([#1528](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1528))
* resource/ibm_iam_access_group_dynamic_rule : Fix dynamic Rule returning wrong resource ([#1535](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1535))
* resource/ibm_is_image: Fix the nil pointer exception for is image ([#1540](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1540))

## 1.7.0 (May 28, 2020)

ENHANCEMENTS:

* resource/ibm_cis_rate_limit: Support for CIS Rate Limiting ( [#1271](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1466))
* data/ibm_cis_rate_limit: Support for CIS Rate Limiting ( [#1271](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1466))

BUG FIXES:

* resource/ibm_is_security_group_rule: Gen1-Security Group Rule fix: allow 'Any' type for ICMP, TCP, UDP( [#1499](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1499))
* resource/ibm_dns_resource_record : Changes to lock resource record id and zone id ( [#1430](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1490))
* resource/ibm_is_vpc: Resource level Timeout updation and docs for vpc resources (is_vpc, is_vpc_route, is_vpn_gateway, is_vpn_gateway_connection )( [#1442](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1442))
* resource/ibm_is_vpn_gateway: Fix for deletion of VPN gateway( [#1495](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1495))
* resource/ibm_private_dns: Fix for provisioning of private dns resource records( [#1476](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1476 ))
* data/source_ibm_is_subnets: Fix for ibm_is_subnets output duplicates( [#1500](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1500 ))


## 1.6.0 (May 20, 2020)
FEATURES:

* New Resource: ([ibm_iam_custom_role](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1433))
* New Datasource: ([ibm_dns_zones](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/958))
* New Datasource: ([ibm_dns_permitted_networks](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/958))
* New Datasource: ([ibm_dns_resource_records](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/958))

ENAHANCEMENTS:

*resource: Adopt custom roles to IAM policies ([#1433](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1433))

## 1.5.3 (May 19, 2020)

ENAHANCEMENTS:

* resource :  Support for pi_pin_policy argument  in ibm_pi_instance ([#1469](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1469))
* resource :  Support for wait_till_albs argument  in ibm_container_workerpool_zone_attachment ([#1463](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1463))
* data : Support for state_store_json attribute in ibm_schematics_state ([#1411](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1411))

BUG FIXES:

* resource : Fix nil pointer if apikey not given for VPC ([#1427](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1427))
* data : CMS issuance_info update ([#1277](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1277))

## 1.5.2 (May 07, 2020)

ENAHANCEMENTS:

* resource : Support for entitlement argument for IKS Classic ROKS cluster (ibm_container_cluster) and worker pool(ibm_container_worker_pool)([#1350](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1350))

* resource : Support for source_resource_group_id and target_resource_group_id([#1364](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1364))

BUG FIXES:

* resource : Error deleting instance with data volume ([#1412](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1412))
* resource : Add force_new true for cidr argument of ibm_is_address_prefix ([#1416](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1416))
* resource : Fix import of ibm_container_cluster ([#1360](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1360))

## 1.5.1 (May 04, 2020)
BUG FIXES:

* resource : Fix VPC subnets created in incorrect resource group([#1398](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1398))

## 1.5.0 (April 29, 2020)
FEATURES:

* New Resource: ([ibm_is_lb_listener_policy_rule](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1147) )
* New Datasource: ([ibm_certificate_manager_certificates](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1277) )

ENAHANCEMENTS:

* resource : Support for auto-generate client_id and client_id for API gateway endpoint([#1390](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1390))
* resource: Support point_in_time_recovery_time and point_in_time_recovery_deployment_id arguments for ICD database([#1259](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1259))
* resource: Support for pending_reclamination for database and CIS instances ([#1242](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1242))


BUG FIXES:

* resource : Fix VPC Load Balancer resource ID is appended to Pool/Listener/Listener Policy ID ([#1359](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1359))
* resource : Fix domainID for CIS firewall resource ([#1201](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1201))
* resource : Fix the update of private dns resource record TTL ([#1331](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1331))
* resource : ibm_container_worker_pool_zone_attachment should wait for ALBs to finish in a new zone ([#1372](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1372))


## 1.4.0 (April 16, 2020)
  
NOTE :  For creating either vpc-classic (generation=1) or vpc-Gen2 (generation=1) IKS cluster, generation parameter needs to be set either in provider block or export via environment variable “IC_GENERATION”. By default the generation value is 2. 

FEATURES:

* New Resource: ([Terraform support for DNS service (beta service ) ibm_dns_zone, ibm_dns_permitted_network, ibm_dns_resource_record](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/958 ))
* New Resource: ([ibm_cis_firewall (lockdown)](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1201) )
* New Resource: ([ibm_lb_listener_policy](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1147) )

ENAHANCEMENTS:

* resource : Add support for resource group argument in ibm_is_network_acl ([#1265](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1265))
* resource : Support for IKS on Gen-2  (beta service) ([#1321](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1321))
* resource : Update functionality support for cis resources  ([#1180](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1180))
* resource : Add support for crn attribute for is_vpc   ([#1315](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1315) )
* data : Add support for crn attribute for is_vpc   ([#1317](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1317))

BUG FIXES:

* resource :  Fix the nil pointer exception for ibm_is_lb_listener resource ([#1289](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1289))


## 1.3.0 (April 02, 2020)

FEATURES:

* New Resource: ([ibm_iam_access_group_dynamic_rule](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/691))
* New Resource: ([ibm_api_gateway_endpoint](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1247))
* New Resource: ([ibm_api_gateway_endpoint_subscription](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1247))
* New DataSource: ([ibm_iam_access_group](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/953))
* New DataSource: ([ibm_api_gateway](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1247))

BUG FIXES:

* resource : Fix the destroy of cloudantnosqldb service([#1242](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1242)) 
* resource : Fix the ICD service endpoint for osl01([#1158](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1158)) 


## 1.2.6 (March 26, 2020)

ENHANCEMENTS:

* resource : Added support for cse_source_addresses  attribute for ibm_is_vpc  ([#1165](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1165)) 
* data : Added support for cse_source_addresses  attribute for ibm_is_vpc ([#1165](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1165)) 
* resource: Added support for new storage class smart for COS bucket  ([#1184](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1184))
* resource:  Allow deletion of non-existing resources like is_vpc, is_subnet, is_vpc_address_prefix and is_instance  ([#1229](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1229))
* resource:  Added support for force_delete argument for ibm_kp_key ([#1214](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1214))


## 1.2.5 (March 19, 2020)

ENHANCEMENTS:

* Provider : Adapt IAM access resources to v2 version ([#1183](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1183)) 
* resource: Added support for GUID attribute for ibm_cis and ibm_database ([#1169](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1169)) 
* data: Added support for GUID attribute for ibm_cis and ibm_database ([#1169](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1169))

BUG FIXES:

* resources : Updated the status string for `ibm_resource_instance, ibm_database and ibm_cis` to be inline with resource controller API changes ([#1190](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1190)) 
* resource/ibm_compute_bare_metal: Fix the order of provisioning of `bare metal` for processor capacity restriction type and SAP servers ([#1189](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1189)) 
* resource/ibm_resource_instance: Fix the order of provisioning of `block chain` platform service ([#1186](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1186)) 
* resource/ibm_container_cluster: Fix the force new for deprecated `billing` argument([#1187](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1187))  


## 1.2.4 (March 11, 2020)

ENHANCEMENTS:

* Provider: Added new parameter `zone` to support power virtual resources and data sources to work in multi-zone environment ([#1141](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1141))
* resource/ibm_pi_volume: Updated the list of volume types for power virtual volume ([#1149](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1149))
* resource/ibm_container_vpc_cluster : Added support for `ingress_hostname` and `ingress_secret` attributes ([#1167](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1167))
* data/ibm_container_vpc_cluster : Added support for `ingress_hostname` and `ingress_secret` attributes ([#1167](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1167))
* resource/ibm_is_floating_ip : Handle the case when floating IP is deleted manually ([#1160](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1160))

BUG FIXES:

* resources : Handle the case where the resource might be already deleted (manually) for ibm_iam_access_policies, ibm_iam_authorization_policies, ibm_iam_service_policies ([#1162](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1162))
* resource/ibm_is_inetwork_acl: Fix the order of creation of network acl ([#1123](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1123))
* resource/ibm_container_vpc_cluster: Added new attribute `wait_till` to control the cluster creation. Now user can control the cluster creation until master is ready / any one worker node is ready / ingress_hostname is  
  assigned.  ([#1143](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1143))
* resource/ibm_pi_instance: Fix the timeout configuration for create ([#1178](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1178))
* doc/ibm_cis_ip_addresses : Fix the description of data source ([#1178](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1178))

## 1.2.3 (March 03, 2020)

BUG FIXES:

* data/ibm_container_cluster_config : Fix the error to download the cluster config for VPC clusters ([#1150](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1150))

## 1.2.2 (February 26, 2020)

ENHANCEMENTS:

* resource/ibm_is_vpc: Improved error message for VPC creation ([#1106](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1106))
* resource/ibm_is_ssh_key: Improved error message for VPC SSH Key creation ([#1105](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1105))
* resource/ibm_container_cluster : Added gateway feature support for IKS clusters. This feature helps to create a cluster with a gateway worker pool of two gateway worker nodes that are connected to public and private VLANs to provide limited public access, and a compute worker pool of compute worker nodes that are connected to the private VLAN only. 
([#1125](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1125))
* data/ibm_conatiner_cluster_config : Extended the data source to provide additional attribute like admin_key, admin_certificate, ca_certificate, host and token. This attributes helps to connect to other providers like Kubernetes and Helm without loading cluster config file. ([#895](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/895))

BUG FIXES:

* doc/ibm_certificate_manager_order: Changed the type of rotate_key from string to bool ([#1110](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1110))
* resource/ibm_is_instance: Fix for updating security group for primary network interface for vpc instance. Now users can add or delete security groups([#1078](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1078))
* doc/ibm_resource_key : Provided an example in the docs as a workaround to create credentials using serviceID parameter ([#1121](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1121))
* resource/ibm_is_network_acl : Fix for crash during the update of rules. Fix for the order of rules creation. Now users can add or delete rules for network_acl ([#1117](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1117))
* resource/ibm_is_public_gateway : Added support for resource group and tags parameters ([#1102](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1102))
* resource/ibm_is_floating_ip : Added support for tags parameters ([#1131](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1131))
* resource/ibm_database : Parameters remote_leader_id, key_protect_instance and key_protect_key can’t be updated after creation. ([#1111](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1111))
* example/ibm-key-protect : Updated example to create an authorisation policy between COS and Key Protect instance([#1133](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1133))
* resource/ibm_resource_group: Removed suppression of error during deletion ([#1108](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1108))
* resource/ibm_iam_user_invite : Fix for inviting user from IBM Cloud lite account. ([#1114](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1114))


