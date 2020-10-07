## 1.13.1 (Oct07, 2020)

ENHANCEMENTS:

* resource: Support endpoint_type argument and eu-fr2 region for COS bucket ([#1945](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1945))

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

* resource: Support eu-fr2 region for ICD [#1870](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1870))

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


