---
layout: "ibm"
page_title: "IBM : ibm_pdr_managedr"
description: |-
  Manages pdr_managedr.
subcategory: "DrAutomation Service"
---

# ibm_pdr_managedr

Creates DR Deployment by creating Orchestrator instance in the given PowerVS workspace and configuration. Orchestrator instance can be used to manage multiple virtual servers and ensure continuous availability. For more details, refer Deploying the Orchestrator -https://cloud.ibm.com/docs/dr-automation-powervs?topic=dr-automation-powervs-idep-the-orch

## Example Usage

```hcl
ServiceInstanceManageDr HA with sshkey
resource "ibm_pdr_managedr" "pdr_managedr_instance" {
  instance_id                         = "050ebe3b-13f4-4db8-8ece-501a3c13be80mh1"
  orchestrator_ha                     = true
  orchestrator_location_type          = "off-premises"
  location_id                         = "dal10"
  orchestrator_workspace_id           = "75cbf05b-78f6-406e-afe7-a904f646d798"
  orchestrator_name                   = "drautomationprimarymh1"
  orchestrator_password               = "EverytimeNewPassword@1"
  machine_type                        = "s922"
  tier                                = "tier1"
  ssh_key_name                        = "vijaykey"
  action                              = "done"
  api_key                             = "apikey is required"

  # Standby configuration (applicable only for HA setup)
  standby_orchestrator_name           = "drautomationstandbymh1"
  standby_orchestrator_workspace_id   = "71027b79-0e31-44f6-a499-63eca1a66feb"
  standby_machine_type                = "s922"
  standby_tier                        = "tier1"
  standby_redeploy                    = false

  # MFA (multi-factor authentication) details
  client_id                           = "123abcd-97d2-4b14-bf62-8eaecc67a122"
  client_secret                       = "abcdefgT5rS8wK6qR9dD7vF1hU4sA3bE2jG0pL9oX7yC"
  tenant_name                         = "xxx.ibm.com"
  proxy_ip                            = "10.3.41.4:443"
}

```
```hcl
ServiceInstanceManageDr HA with secrets
resource "ibm_pdr_managedr" "pdr_managedr_instance" {
  instance_id                       = "050ebe3b-13f4-4db8-8ece-501a3c13be80mh3"
  orchestrator_ha                   = true
  orchestrator_location_type        = "off-premises"
  location_id                       = "dal10"
  orchestrator_workspace_id         = "75cbf05b-78f6-406e-afe7-a904f646d798"
  orchestrator_name                 = "drautomationprimarymh3"
  orchestrator_password             = "EverytimeNewPassword@1"
  machine_type                      = "s922"
  tier                              = "tier1"
  guid                              = "397dc20d-9f66-46dc-a750-d15392872023"
  secret_group                      = "12345-714f-86a6-6a50-2f128a4e7ac2"
  secret                            = "12345-997c-1d0d-5503-27ca856f2b5a"
  region_id                         = "us-south"
  action                            = "done"
  api_key                           = "apikey is required"

  # Standby configuration (for HA setup)
  standby_orchestrator_name         = "drautomationstandbymh3"
  standby_orchestrator_workspace_id = "71027b79-0e31-44f6-a499-63eca1a66feb"
  standby_machine_type              = "s922"
  standby_tier                      = "tier1"
  standby_redeploy                  = false

  # MFA (Multi-Factor Authentication)
  client_id                         = "123abcd-97d2-4b14-bf62-8eaecc67a122"
  client_secret                     = "abcdefgT5rS8wK6qR9dD7vF1hU4sA3bE2jG0pL9oX7yC"
  tenant_name                       = "xxx.ibm.com"
  proxy_ip                          = "10.3.41.4:443"
}
```
```hcl
ServiceInstanceManageDr Non-HA with sshkey
resource "ibm_pdr_managedr" "pdr_managedr_instance" {
  instance_id                 = "050ebe3b-13f4-4db8-8ece-501a3c13be80mnh5"
  orchestrator_ha             = false
  orchestrator_location_type  = "off-premises"
  location_id                 = "dal10"
  orchestrator_workspace_id   = "75cbf05b-78f6-406e-afe7-a904f646d798"
  orchestrator_name           = "drautomationprimarymnh5"
  orchestrator_password       = "EverytimeNewPassword@1"
  machine_type                = "s922"
  tier                        = "tier1"
  ssh_key_name                = "vijaykey"
  action                      = "done"
  api_key                     = "apikey is required"

  # MFA (Multi-Factor Authentication)
  client_id                   = "123abcd-97d2-4b14-bf62-8eaecc67a122"
  client_secret               = "abcdefgT5rS8wK6qR9dD7vF1hU4sA3bE2jG0pL9oX7yC"
  tenant_name                 = "xxx.ibm.com"
  proxy_ip                    = "10.3.41.4:443"
}
```
```hcl
ServiceInstanceManageDr Non-HA with secrets
resource "ibm_pdr_managedr" "pdr_managedr_instance" {
  instance_id                 = "050ebe3b-13f4-4db8-8ece-501a3c13be80mnh7"
  orchestrator_ha             = false
  orchestrator_location_type  = "off-premises"
  location_id                 = "dal10"
  orchestrator_workspace_id   = "75cbf05b-78f6-406e-afe7-a904f646d798"
  orchestrator_name           = "drautomationprimarymnh7"
  orchestrator_password       = "EverytimeNewPassword@1"
  machine_type                = "s922"
  tier                        = "tier1"
  guid                        = "397dc20d-9f66-46dc-a750-d15392872023"
  secret_group                = "12345-714f-86a6-6a50-2f128a4e7ac2"
  secret                      = "12345-997c-1d0d-5503-27ca856f2b5a"
  region_id                   = "us-south"
  action                      = "done"
  api_key                     = "apikey is required"

  # MFA (Multi-Factor Authentication)
  client_id                   = "123abcd-97d2-4b14-bf62-8eaecc67a122"
  client_secret               = "abcdefgT5rS8wK6qR9dD7vF1hU4sA3bE2jG0pL9oX7yC"
  tenant_name                 = "xxx.ibm.com"
  proxy_ip                    = "10.3.41.4:443"
}
```

## Argument Reference

You can specify the following arguments for this resource:

* `instance_id` - (Required, Forces new resource, String) The ID of the Power DR Automation service instance.
* `action` - (Optional, String) Indicates whether to proceed with asynchronous operation after all configuration details are updated in the database.
* `api_key` - (Required, String, Sensitive) The api Key of the service instance for deploying the disaster recovery service.
* `client_id` - (Optional, String) The Client Id created for MFA authentication API.
* `client_secret` - (Optional, String, Sensitive) The client secret created for MFA authentication API.
* `tenant_name` - (Optional, String) The tenant name for MFA authentication API.
* `proxy_ip` - (Optional, String) The Proxy IP for the Communication between Orchestrator and Service.
* `guid` - (Optional, String) The global unique identifier of the service instance.
* `region_id` - (Optional, String) The power virtual server region where the service instance is deployed.
* `location_id` - (Required, String) The Location or data center identifier where the service instance is deployed. you can fetch locations using data_source "ibm_pdr_get_dr_locations". 
* `machine_type` - (Required, String) The machine type used for deploying orchestrator. you can fetch machine types use  data_source "ibm_pdr_get_machine_types".
* `tier` - (Required, String) The storage tier used for deploying primary orchestrator (e.g., tier1, tier3, etc).
* `ssh_key_name` - (Optional, String) The name of the SSH key used for deploying the orchestator.
* `orchestrator_ha` - (Required, Boolean) Indicates whether the orchestrator High Availability (HA) is enabled for the service instance.
* `orchestrator_location_type` - (Required, String) The cloud location where your orchestator need to be created.(eg., "off-premises", "on-premises")
* `orchestrator_name` - (Required, String) Username for the orchestrator management interface.
* `orchestrator_password` - (Required, String, Sensitive) The password that you can use to access your orchestrator.
* `orchestrator_workspace_id` - (Required, String) The unique identifier orchestrator workspace.
* `secret_group` - (Optional, String) The secret group name in IBM Cloud Secrets Manager containing sensitive data for the service instance.
* `secret` - (Optional, String) Secret name or identifier used for retrieving credentials from Secrets Manager.
* `standby_orchestrator_name` - (Optional, String) Username for the standby orchestrator management interface.
* `standby_orchestrator_workspace_id` - (Optional, String) The unique identifier of the standby orchestrator workspace.
* `standby_orchestrator_workspace_location` - (Optional, String) Location of the standby orchestrator workspace.
* `standby_machine_type` - (Optional, String) The machine type used for deploying standby virtual machines.
* `standby_tier` - (Optional, String) The storage tier used for deploying standby orchestrator.
* `stand_by_redeploy` - (Optional, String)  Flag to indicate if standby should be redeployed only for HA case (must be "true" or "false").

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the pdr_managedr.
* `dashboard_url` - (String) URL to the dashboard for managing the DR service instance in IBM Cloud.
* `instance_id` - (String) The CRN (Cloud Resource Name) of the DR service instance.


