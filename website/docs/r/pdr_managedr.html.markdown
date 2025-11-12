---
layout: "ibm"
page_title: "IBM : ibm_pdr_managedr"
description: |-
  Manages pdr_managedr.
subcategory: "DrAutomation Service"
---

# ibm_pdr_managedr

 Creates Orchestrator VM in the given workspace and configuration. Orchestrator VM can be used to manage multiple virtual servers and help ensure continuous availability. For more details, refer Deploying the Orchestrator -
 https://test.cloud.ibm.com/docs/dr-automation-powervs?topic=dr-automation-powervs-idep-the-orch

## Example Usage

```hcl
ServiceInstanceManageDr HA with sshkey
resource "ibm_pdr_managedr" "pdr_managedr_instance" {
  instance_id                         = "crn:v1:staging:public:power-dr-automation:global:a/a09202c1bfb04ceebfb4a9fd38c87721:050ebe3b-13f4-4db8-8ece-501a3c13be80mh1::"
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
}

```
```hcl
ServiceInstanceManageDr HA with secrets
resource "ibm_pdr_managedr" "pdr_managedr_instance" {
  instance_id                       = "crn:v1:staging:public:power-dr-automation:global:a/a09202c1bfb04ceebfb4a9fd38c87721:050ebe3b-13f4-4db8-8ece-501a3c13be80mh3::"
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
}
```
```hcl
ServiceInstanceManageDr Non-HA with sshkey
resource "ibm_pdr_managedr" "pdr_managedr_instance" {
  instance_id                 = "crn:v1:staging:public:power-dr-automation:global:a/a09202c1bfb04ceebfb4a9fd38c87721:050ebe3b-13f4-4db8-8ece-501a3c13be80mnh5::"
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
}
```
```hcl
ServiceInstanceManageDr Non-HA with secrets
resource "ibm_pdr_managedr" "pdr_managedr_instance" {
  instance_id                 = "crn:v1:staging:public:power-dr-automation:global:a/a09202c1bfb04ceebfb4a9fd38c87721:050ebe3b-13f4-4db8-8ece-501a3c13be80mnh7::"
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
}
```

## Argument Reference

You can specify the following arguments for this resource:

* `instance_id` - (Required, Forces new resource, String) The CRN (Cloud Resource Name) of the Power DR Automation service instance.
* `action` - (Optional, String) Indicates whether to proceed with asynchronous operation after all configuration details are updated in the database.
* `api_key` - (Required, String, Sensitive) The API key associated with the IBM Cloud service instance.
* `client_id` - (Optional, String) Client ID for MFA (Multi-Factor Authentication).
* `client_secret` - (Optional, String, Sensitive) Client secret for MFA (Multi-Factor Authentication).
* `tenant_name` - (Optional, String) Tenant name for MFA authentication.
* `guid` - (Optional, String) The globally unique identifier of the service instance.
* `region_id` - (Optional, String) Cloud region where the service instance is deployed.
* `location_id` - (Optional, String) Location or data center identifier where the service instance is deployed.
* `machine_type` - (Optional, String) Machine type or flavor used for virtual machines in the service instance.
* `tier` - (Optional, String) Tier of the service instance.
* `ssh_key_name` - (Optional, String) Name of the SSH key stored in the cloud provider.
* `ssh_public_key` - (Optional, String) SSH public key for accessing virtual machines in the service instance.
* `orchestrator_ha` - (Optional, Boolean) Flag to enable or disable High Availability (HA) for the service instance.
* `orchestrator_location_type` - (Optional, String) Type of orchestrator cluster used in the service instance.
* `orchestrator_name` - (Optional, String) Username for the orchestrator management interface.
* `orchestrator_password` - (Optional, String, Sensitive) Password for the orchestrator management interface.
* `orchestrator_workspace_id` - (Optional, String) ID of the orchestrator workspace.
* `orchestrator_workspace_location` - (Optional, String) Location of the orchestrator workspace.
* `resource_instance` - (Optional, String) ID of the associated IBM Cloud resource instance.
* `secondary_workspace_id` - (Optional, String) ID of the secondary workspace used for redundancy or disaster recovery.
* `secret_group` - (Optional, String) Secret group name in IBM Cloud Secrets Manager containing sensitive data for * the service instance.
* `secret` - (Optional, String) Secret name or identifier used for retrieving credentials from Secrets Manager.
* `standby_orchestrator_name` - (Optional, String) Username for the standby orchestrator management interface.
* `standby_orchestrator_workspace_id` - (Optional, String) ID of the standby orchestrator workspace.
* `standby_orchestrator_workspace_location` - (Optional, String) Location of the standby orchestrator workspace.
* `standby_machine_type` - (Optional, String) Machine type or flavor used for standby virtual machines.
* `standby_tier` - (Optional, String) Tier of the standby service instance.
* `stand_by_redeploy` - (Optional, String) Flag to indicate if the standby environment should be redeployed. Must be "true" or "false".

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the pdr_managedr.
* `dashboard_url` - (String) URL to the dashboard for managing the DR service instance in IBM Cloud.
* `instance_id` - (String) The CRN (Cloud Resource Name) of the DR service instance.

* `etag` - ETag identifier for pdr_managedr.

## Import

You can import the `ibm_pdr_managedr` resource by using `id`.
The `id` property can be formed from `instance_id`, and `instance_id` in the following format:

<pre>
&lt;instance_id&gt;/&lt;instance_id&gt;
</pre>
* `instance_id`: A string in the format `crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::`. instance id of instance to provision.
* `instance_id`: A string in the format `crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::`. The CRN (Cloud Resource Name) of the DR service instance.

# Syntax
<pre>
$ terraform import ibm_pdr_managedr.pdr_managedr &lt;instance_id&gt;/&lt;instance_id&gt;
</pre>


