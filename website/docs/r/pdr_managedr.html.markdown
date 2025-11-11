---
layout: "ibm"
page_title: "IBM : ibm_pdr_managedr"
description: |-
  Manages pdr_managedr.
subcategory: "DrAutomation Service"
---

# ibm_pdr_managedr

Create, update, and delete pdr_managedrs with this resource.

## Example Usage

```hcl
resource "ibm_pdr_managedr" "pdr_managedr_instance" {
  instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `accept_language` - (Optional, Forces new resource, String) The language requested for the return document.
* `accepts_incomplete` - (Optional, Forces new resource, Boolean) A value of true indicates that both the IBM Cloud platform and the requesting client support asynchronous deprovisioning.
  * Constraints: The default value is `true`.
* `instance_id` - (Required, Forces new resource, String) instance id of instance to provision.
* `stand_by_redeploy` - (Optional, Forces new resource, String) Flag to indicate if standby should be redeployed (must be "true" or "false").

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



# ==================================================================================================
                                    # HA Cases
# ==================================================================================================
# Case 1: ManageDR with HA + schematic id + sshkey
# provider "ibm" {
#   region = "us-south"
# }

# resource "ibm_power_vs_ssh_key" "vijaykey" {
#   name       = "vijaykey"
#   public_key = file("~/.ssh/id_rsa.pub")
# }

# resource "ibm_drautomation_service_instance" "ha_dr_instance" {
    # name        = "service1234"
    # crn         = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf7::"
    # enable_flag = false  
    # dr_location_id                    = "dal10"
    # dr_orchestrator_name             = "drautomationprimary7a"
    # dr_orchestrator_password         = "Password1234567"
    # dr_orchestrator_workspace_id     = "75cbf05b-78f6-406e-afe7-a904f646d798"
    # machine_type                     = "s922"
    # orchestrator_cluster_type        = "off-premises"
    # schematic_workspace_id           = "us-south.workspace.projects-service.3ae96a02"
    # ssh_key_name                     = ibm_power_vs_ssh_key.vijaykey.name
    # standby_machine_type             = "s922"
    # standby_orchestrator_name        = "drautomationstandby7a"
    # standby_orchestrator_workspace_id = "71027b79-0e31-44f6-a499-63eca1a66feb"
    # tier                             = "tier1"
# }

# ================================================================================================

# Case 2: ManageDR with HA + custom VPC + sshkey

# resource "ibm_drautomation_service_instance" "ha_dr_instance" {
    # name        = "service1234"
    # crn         = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf7::"
    # enable_flag = false    
    # dr_location_id                     = "dal10"
    # dr_orchestrator_name              = "drautomationprimary7a"
    # dr_orchestrator_password          = "Password1234567"
    # dr_orchestrator_workspace_id      = "75cbf05904f646d798"
    # machine_type                      = "s922"
    # orchestrator_cluster_type         = "off-premises"
    # ssh_key_name                      = ibm_power_vs_ssh_key.vijaykey.name
    # standby_machine_type              = "s922"
    # standby_orchestrator_name         = "drautomationstandby7a"
    # standby_orchestrator_workspace_id = "71027b79-0e31-44f6-a499-63eca1a66feb"
    # tier                              = "tier1"
# }

# ==================================================================================================

# Case 3: ManageDR with HA + schematic id + secrets

# resource "ibm_drautomation_service_instance" "ha_dr_instance" {
#   name        = "service1234"
#   crn         = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf7::"
#   enable_flag = false   
#     dr_location_id                     = "dal10"
#     dr_orchestrator_name              = "drautomationprimary7a"
#     dr_orchestrator_password          = "Password1234567"  # Consider using a variable or secret reference
#     dr_orchestrator_workspace_id      = "75cbf05b-78f6-406e-afe7-a904f646d798"
#     machine_type                      = "s922"
#     orchestrator_cluster_type         = "off-premises"
#     schematic_workspace_id            = "us-south.workspace.projects-service.3ae96a02"
#     secret_group                      = "secret_group id"
#     secret                            = "secret id"
#     region_id                         = "us-south"
#     guid                              = "gu id"
#     standby_machine_type              = "s922"
#     standby_orchestrator_name         = "drautomationstandby7a"
#     standby_orchestrator_workspace_id = "71027b79-0e31-44f6-a499-63eca1a66feb"
#     tier                              = "tier1"
# }

# ==================================================================================================

# Case 4: ManageDR with HA + custom VPC + secrets

# resource "ibm_drautomation_service_instance" "ha_dr_instance" {
#   name        = "service1234"
#   crn         = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf7::"
#   enable_flag = false   
#     dr_location_id                     = "dal10"
#     dr_orchestrator_name              = "drautomationprimary7a"
#     dr_orchestrator_password          = "Password1234567"  # Use a variable or secret reference in production
#     dr_orchestrator_workspace_id      = "75cbf05b-78f6-406e-afe7-a904f646d798"
#     machine_type                      = "s922"
#     orchestrator_cluster_type         = "off-premises"
#     secret_group                      = "secret_group id"
#     secret                            = "secret id"
#     region_id                         = "us-south"
#     guid                              = "gu id"
#     standby_machine_type              = "s922"
#     standby_orchestrator_name         = "drautomationstandby7a"
#     standby_orchestrator_workspace_id = "71027b79-0e31-44f6-a499-63eca1a66feb"
#     tier                              = "tier1"
# }

# ==================================================================================================
                                    # Non-HA Cases
# ==================================================================================================

# Case 1: ManageDR without HA + schematic id + sshkey
# resource "ibm_drautomation_service_instance" "non_ha_dr_instance" {
#   name        = "service1234"
#   crn         = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf7::"
#   enable_flag = false   
#     dr_location_id                = "dal10"
#     dr_orchestrator_name         = "drautomationprimary7a"
#     dr_orchestrator_password     = "Password1234567"  # Use a variable or secret reference in production
#     dr_orchestrator_workspace_id = "75cbf05b-78f6-406e-afe7-a904f646d798"
#     machine_type                 = "s922"
#     orchestrator_cluster_type    = "off-premises"
#     schematic_workspace_id       = "us-south.workspace.projects-service.3ae96a02"
#     ssh_key_name                 = ibm_power_vs_ssh_key.vijaykey.name
#     tier                         = "tier1"
# }

# ==================================================================================================

# Case 2: ManageDR without HA + custom VPC + sshkey
# resource "ibm_drautomation_service_instance" "non_ha_dr_instance" {
#   name        = "service1234"
#   crn         = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf7::"
#   enable_flag = false   
#     dr_location_id                = "dal10"
#     dr_orchestrator_name         = "drautomationprimary7a"
#     dr_orchestrator_password     = "Password1234567"  # Use a variable or secret reference in production
#     dr_orchestrator_workspace_id = "75cbf05b-78f6-406e-afe7-a904f646d798"
#     machine_type                 = "s922"
#     orchestrator_cluster_type    = "off-premises"
#     ssh_key_name                 = ibm_power_vs_ssh_key.vijaykey.name
#     tier                         = "tier1"
# }

# ==================================================================================================

# Case  3: ManageDR without HA + schematic id + secrets
# resource "ibm_drautomation_service_instance" "non_ha_dr_instance" {
#   name        = "service1234"
#   crn         = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf7::"
#   enable_flag = false#   
#     dr_location_id                = "dal10"
#     dr_orchestrator_name         = "drautomationprimary7a"
#     dr_orchestrator_password     = "Password1234567"  # Use a variable or secret reference in production
#     dr_orchestrator_workspace_id = "75cbf05b-78f6-406e-afe7-a904f646d798"
#     machine_type                 = "s922"
#     orchestrator_cluster_type    = "off-premises"
#     schematic_workspace_id       = "us-south.workspace.projects-service.3ae96a02"
#     secret_group                 = "secret_group id"
#     secret                       = "secret id"
#     region_id                    = "us-south"
#     guid                         = "gu id"
#     tier                         = "tier1"
# }

# ==================================================================================================

# Case  4: ManageDR without HA + custom VPC + secrets
# resource "ibm_drautomation_service_instance" "non_ha_dr_instance" {
#   name        = "service1234"
#   crn         = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf7::"
#   enable_flag = false#   
#     dr_location_id                = "dal10"
#     dr_orchestrator_name         = "drautomationprimary7a"
#     dr_orchestrator_password     = "Password1234567"  # Use a variable or secret reference in production
#     dr_orchestrator_workspace_id = "75cbf05b-78f6-406e-afe7-a904f646d798"
#     machine_type                 = "s922"
#     orchestrator_cluster_type    = "off-premises"
#     secret_group                 = "secret_group id"
#     secret                       = "secret id"
#     region_id                    = "us-south"
#     guid                         = "gu id"
#     tier                         = "tier1"
# }


# ================================================================================

