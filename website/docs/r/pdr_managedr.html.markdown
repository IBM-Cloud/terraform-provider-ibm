---
layout: "ibm"
page_title: "IBM : ibm_pdr_managedr"
description: |-
  Manages pdr_managedr.
subcategory: "HA and DR Automation for IBM® Power® Virtual Server API reference"
---

# ibm_pdr_managedr

Creates DR Deployment by creating Orchestrator instance in the given PowerVS workspace and configuration. Orchestrator instance can be used to manage multiple virtual servers and ensure continuous availability. For more details, refer Deploying the Orchestrator -https://cloud.ibm.com/docs/dr-automation-powervs?topic=dr-automation-powervs-idep-the-orch

## Example Usage

```hcl
resource "ibm_pdr_managedr" "pdr_managedr_instance" {
  location_id                = "dal10"
  orchestrator_location_type = "off-premises"
  orchestrator_name          = "drautomationprimary"
  orchestrator_password      = "password"
  orchestrator_workspace_id  = "75cbf05b-78f6-406e-afe7-a904f646d798"

  machine_type                = "s922"
  orchestrator_ha              = false
  proxy_ip                    = "10.30.40.10:8888"
  ssh_key_name                = "vijaykey"
  standby_machine_type        = "s922"
  standby_orchestrator_name   = "drautomationstandby"
  standby_orchestrator_workspace_id = "71027b79-0e31-44f6-a499-63eca1a66feb"
  tenant_name                 = "xxx.ibm.com"
  tier                        = "tier1"
}
```
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
  ssh_key_name                        = "samplekey"
  action                              = "done"
  api_key                             = "apikey is required"
  standby_orchestrator_network_ids   =["0f635vae-xxxx-xxxx-xxxx-43f2e55127b9","0f6354ae-xxxx-xxxx-xxxx-43f2e551v7b0"]
  orchestrator_network_ids            = ["0f635vae-xxxx-xxxx-xxxx-43f2e55127b9","0f6354ae-xxxx-xxxx-xxxx-43f2e551v7b0"]
  standby_ssh_key_name                = "samplekey"
  managed_apikey              = "xxxxxx-xxxx-xxxx-xxxxx"

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

  #secure connection(CA certificates) --> Secrets manager
  primary_orch_ca_certs_secrets      ={
                        "ca_certificate_secret_id": "dxxxx-4b72-xxxx-2a22-xxxx",
                        "ca_secret_manager_guid": "xxxxx-9f66-xxxx-a750-xxxxxx"
                        }
  "standby_orch_ca_certs_secrets":{
                      "ca_certificate_secret_id": "xxxxxx-4b72-xxxx-2a22-xxxxxxx"
                      },
  #secure connection(CA certificates) --> Bash64 encoded file content
  primary_orch_ca_certs_server_cert   = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  primary_orch_ca_certs_server_key    = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  standby_orch_ca_certs_server_cert   = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  standby_orch_ca_certs_server_key    = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
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
  orchestrator_password             = "Everytxxxxxxxxxxxxxxx@1"
  machine_type                      = "s922"
  tier                              = "tier1"
  guid                              = "397dc20d-9f66-46dc-a750-d15392872023"
  secret_group                      = "12345-714f-86a6-6a50-2f128a4e7ac2"
  secret                            = "12345-997c-1d0d-5503-27ca856f2b5a"
  region_id                         = "us-south"
  action                            = "done"
  api_key                           = "apikey is required"
  standby_orchestrator_network_ids  =["0f635vae-xxxx-xxxx-xxxx-43f2e55127b9","0f6354ae-xxxx-xxxx-xxxx-43f2e551v7b0"]
  orchestrator_network_ids          = ["0f635vae-xxxx-xxxx-xxxx-43f2e55127b9","0f6354ae-xxxx-xxxx-xxxx-43f2e551v7b0"]
  standby_ssh_key_name              = "samplekey"
  managed_apikey                    = "xxxxxx-xxxx-xxxx-xxxxx"

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

  #secure connection(Optional, CA certificates) --> Secrets manager
  primary_orch_ca_certs_secrets      ={
                        "ca_certificate_secret_id": "dxxxx-4b72-xxxx-2a22-xxxx",
                        "ca_secret_manager_guid": "xxxxx-9f66-xxxx-a750-xxxxxx"
                        }
  "standby_orch_ca_certs_secrets":{
                      "ca_certificate_secret_id": "xxxxxx-4b72-xxxx-2a22-xxxxxxx"
                      },
  #secure connection(Optional, CA certificates) --> Bash64 encoded file content
  primary_orch_ca_certs_server_cert   = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  primary_orch_ca_certs_server_key    = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  standby_orch_ca_certs_server_cert   = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  standby_orch_ca_certs_server_key    = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
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
  ssh_key_name                = "samplekey"
  action                      = "done"
  api_key                     = "apikey is required"
  orchestrator_network_ids    = ["0f635vae-xxxx-xxxx-xxxx-43f2e55127b9","0f6354ae-xxxx-xxxx-xxxx-43f2e551v7b0"]
  managed_apikey              = "xxxxxx-xxxx-xxxx-xxxxx"

  # MFA (Multi-Factor Authentication)
  client_id                   = "123abcd-97d2-4b14-bf62-8eaecc67a122"
  client_secret               = "abcdefgT5rS8wK6qR9dD7vF1hU4sA3bE2jG0pL9oX7yC"
  tenant_name                 = "xxx.ibm.com"
  proxy_ip                    = "10.3.41.4:443"

  #secure connection(Optional, CA certificates) --> Secrets manager
  primary_orch_ca_certs_secrets      ={
                        "ca_certificate_secret_id": "dxxxx-4b72-xxxx-2a22-xxxx",
                        "ca_secret_manager_guid": "xxxxx-9f66-xxxx-a750-xxxxxx"
                        }
  "standby_orch_ca_certs_secrets":{
                      "ca_certificate_secret_id": "xxxxxx-4b72-xxxx-2a22-xxxxxxx"
                      },
  #secure connection(Optional, CA certificates) --> Bash64 encoded file content
  primary_orch_ca_certs_server_cert   = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  primary_orch_ca_certs_server_key    = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  standby_orch_ca_certs_server_cert   = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  standby_orch_ca_certs_server_key    = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
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
  orchestrator_network_ids    = ["0f635vae-xxxx-xxxx-xxxx-43f2e55127b9","0f6354ae-xxxx-xxxx-xxxx-43f2e551v7b0"]
  managed_apikey              = "xxxxxx-xxxx-xxxx-xxxxx"

  # MFA (Multi-Factor Authentication)
  client_id                   = "123abcd-97d2-4b14-bf62-8eaecc67a122"
  client_secret               = "abcdefgT5rS8wK6qR9dD7vF1hU4sA3bE2jG0pL9oX7yC"
  tenant_name                 = "xxx.ibm.com"
  proxy_ip                    = "10.3.41.4:443"

  #secure connection(Optional, CA certificates) --> Secrets manager
  primary_orch_ca_certs_secrets      ={
                        "ca_certificate_secret_id": "dxxxx-4b72-xxxx-2a22-xxxx",
                        "ca_secret_manager_guid": "xxxxx-9f66-xxxx-a750-xxxxxx"
                        }
  "standby_orch_ca_certs_secrets":{
                      "ca_certificate_secret_id": "xxxxxx-4b72-xxxx-2a22-xxxxxxx"
                      },
  #secure connection(Optional, CA certificates) --> Bash64 encoded file content
  primary_orch_ca_certs_server_cert   = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  primary_orch_ca_certs_server_key    = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  standby_orch_ca_certs_server_cert   = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  standby_orch_ca_certs_server_key    = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
```

```hcl
ServiceInstanceManageDr Non-HA without Secrets and SSH key
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
  region_id                   = "us-south"
  action                      = "done"
  api_key                     = "apikey is required"
  orchestrator_network_ids    = ["0f635vae-xxxx-xxxx-xxxx-43f2e55127b9","0f6354ae-xxxx-xxxx-xxxx-43f2e551v7b0"]
  managed_apikey              = "xxxxxx-xxxx-xxxx-xxxxx"

  # MFA (Multi-Factor Authentication)
  client_id                   = "123abcd-97d2-4b14-bf62-8eaecc67a122"
  client_secret               = "abcdefgT5rS8wK6qR9dD7vF1hU4sA3bE2jG0pL9oX7yC"
  tenant_name                 = "xxx.ibm.com"
  proxy_ip                    = "10.3.41.4:443"

  #secure connection(Optional, CA certificates) --> Secrets manager
  primary_orch_ca_certs_secrets      ={
                        "ca_certificate_secret_id": "dxxxx-4b72-xxxx-2a22-xxxx",
                        "ca_secret_manager_guid": "xxxxx-9f66-xxxx-a750-xxxxxx"
                        }
  "standby_orch_ca_certs_secrets":{
                      "ca_certificate_secret_id": "xxxxxx-4b72-xxxx-2a22-xxxxxxx"
                      },
  #secure connection(Optional, CA certificates) --> Bash64 encoded file content
  primary_orch_ca_certs_server_cert   = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  primary_orch_ca_certs_server_key    = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  standby_orch_ca_certs_server_cert   = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  standby_orch_ca_certs_server_key    = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
```

### Path Parameters

* `instance_id` - (Required, Forces new resource, String) The unique identifier of the DR service instance.
  * Constraints: The maximum length is `1048` characters. The minimum length is `36` characters. The value must match regular expression `/^.*$/`.

### Query Parameters

* `accept_language` - (Optional, Forces new resource, String) The language in which the response should be returned.
* `accepts_incomplete` - (Optional, Forces new resource, Boolean) Indicates whether the request can be accepted before the operation is complete.
* `stand_by_redeploy` - (Optional, Forces new resource, Boolean) Indicates whether the standby orchestrator should be redeployed.

## Argument Reference

You can specify the following arguments for this resource.

* `location_id` - (Required, Forces new resource, String) Location or data center identifier for the DR deployment.
 * Constraints: The maximum length is `32` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
* `orchestrator_name` - (Required, Forces new resource, String) Name of the primary orchestrator.
 * Constraints: The maximum length is `1048` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
* `orchestrator_password` - (Required, Forces new resource, String) Password for the primary orchestrator management interface.
 * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[\x20-\x7E]*$/`.
* `orchestrator_workspace_id` - (Required, Forces new resource, String) PowerVS workspace ID for the primary orchestrator.
 * Constraints: The maximum length is `1048` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9\-_]*$/`.
* `orchestrator_location_type` - (Required, Forces new resource, String) Type of location where the orchestrator is deployed, such as off-premises or VPC.
 * Constraints: The maximum length is `32` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9\-_]*$/`.
* `primary_orch_ca_certs_server_cert` - (Optional, Forces new resource, String) PEM-encoded CA certificate for the primary orchestrator server certificate.
 * Constraints: The maximum length is `65536` characters. The minimum length is `0` characters. The value must match regular expression `/^[\x20-\x7E]*$/`.
* `primary_orch_ca_certs_server_key` - (Optional, Forces new resource, String) PEM-encoded private key for the primary orchestrator CA certificate.
 * Constraints: The maximum length is `65536` characters. The minimum length is `0` characters. The value must match regular expression `/^[\x20-\x7E]*$/`.
* `standby_orch_ca_certs_server_cert` - (Optional, Forces new resource, String) PEM-encoded CA certificate for the standby orchestrator server certificate.
 * Constraints: The maximum length is `65536` characters. The minimum length is `0` characters. The value must match regular expression `/^[\x20-\x7E]*$/`.
* `standby_orch_ca_certs_server_key` - (Optional, Forces new resource, String) PEM-encoded private key for the standby orchestrator CA certificate.
 * Constraints: The maximum length is `65536` characters. The minimum length is `0` characters. The value must match regular expression `/^[\x20-\x7E]*$/`.
* `api_key` - (Optional, Forces new resource, String) IBM Cloud API key used for DR automation operations.
 * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9\-_]*$/`.
* `ssh_key_name` - (Optional, Forces new resource, String) SSH key name used to access the primary orchestrator VM.
 * Constraints: The maximum length is `1048` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9\-_]*$/`.
* `standby_ssh_key_name` - (Optional, Forces new resource, String) SSH key name used to access the standby orchestrator VM.
 * Constraints: The maximum length is `1048` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
* `orchestrator_ha` - (Optional, Forces new resource, Boolean) Whether high availability is enabled for the orchestrator.
* `resource_instance` - (Optional, Forces new resource, String) CRN or identifier of the associated resource instance.
 * Constraints: The maximum length is `1048` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
* `secret_group` - (Optional, Forces new resource, String) Secrets Manager secret group containing deployment secrets.
 * Constraints: The maximum length is `1048` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
* `secret` - (Optional, Forces new resource, String) Secrets Manager secret name or identifier.
 * Constraints: The maximum length is `1048` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
* `region_id` - (Optional, Forces new resource, String) PowerVS region ID for the deployment.
 * Constraints: The maximum length is `1048` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
* `guid` - (Optional, Forces new resource, String) Globally unique identifier for the service instance.
 * Constraints: The maximum length is `36` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
* `machine_type` - (Optional, Forces new resource, String) Machine type for the primary orchestrator.
 * Constraints: The maximum length is `32` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9\-_]*$/`.
* `tier` - (Optional, Forces new resource, String) Storage tier for the primary orchestrator.
 * Constraints: The maximum length is `32` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9\-_]*$/`.
* `standby_tier` - (Optional, Forces new resource, String) Storage tier for the standby orchestrator.
 * Constraints: The maximum length is `32` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9\-_]*$/`.
* `standby_machine_type` - (Optional, Forces new resource, String) Machine type for the standby orchestrator.
 * Constraints: The maximum length is `32` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9\-_]*$/`.
* `client_id` - (Optional, Forces new resource, String) OAuth client ID used for authentication.
 * Constraints: The maximum length is `1048` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
* `client_secret` - (Optional, Forces new resource, String) OAuth client secret used for authentication.
 * Constraints: The maximum length is `1048` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
* `tenant_name` - (Optional, Forces new resource, String) Tenant name used for MFA authentication.
 * Constraints: The maximum length is `1048` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
* `proxy_ip` - (Optional, Forces new resource, String) Proxy IP address and port for orchestrator-to-service communication.
 * Constraints: The maximum length is `1048` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
* `dedicated_host_enabled` - (Optional, Forces new resource, Boolean) Whether dedicated host deployment is enabled.

* `primary_deployment_target` - (Optional, Forces new resource, Object) Deployment target for the primary virtual server.
* Nested schema for **primary_deployment_target**:
    * `id` - (String) ID of the dedicated host where the virtual server will be deployed.
     * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
    * `type` - (String) Type of the deployment target.
     * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.

* `standby_deployment_target` - (Optional, Forces new resource, Object) Deployment target for the standby virtual server.
* Nested schema for **standby_deployment_target**:
    * `id` - (String) ID of the dedicated host where the virtual server will be deployed.
     * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
    * `type` - (String) Type of the deployment target.
     * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `orchestrator_network_ids` - (Optional, Forces new resource, List of String) Network IDs to attach to the primary orchestrator.
  * Constraints: The maximum number of items is `10`. The minimum number of items is `0`.
  * Nested item constraints: The maximum length of each item is `36` characters. The minimum length of each item is `0` characters. Each item must match regular expression `/^[a-zA-Z0-9\-_]*$/`.
* `standby_orchestrator_network_ids` - (Optional, Forces new resource, List of String) Network IDs to attach to the standby orchestrator.
  * Constraints: The maximum number of items is `10`. The minimum number of items is `0`.
  * Nested item constraints: The maximum length of each item is `36` characters. The minimum length of each item is `0` characters. Each item must match regular expression `/^[a-zA-Z0-9\-_]*$/`.
* `managed_apikey` - (Optional, Forces new resource, String) IBM Cloud API key managed by the service for workload access.
  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9\-_]*$/`.

* `primary_orch_ca_certs_secrets` - (Optional, Forces new resource, Object) Secrets manager details for primary orchestrator CA certificates.
* Nested schema for **primary_orch_ca_certs_secrets**:
    * `ca_certificate_secret_id` - (String) ID of the secret.
     * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9\-]+$/`.
    * `ca_secret_manager_guid` - (String) ID of the Primary CA certificate secret.
     * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9\-]+$/`.

* `standby_orch_ca_certs_secrets` - (Optional, Forces new resource, Object) Secrets manager details for standby orchestrator CA certificates.
* Nested schema for **standby_orch_ca_certs_secrets**:
    * `ca_certificate_secret_id` - (String) ID of the Standby CA certificate secret.
     * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9\-]+$/`.

## Attribute Reference

After your resource is created, you can read values from the following attributes.

* `id` - (String) The CRN (Cloud Resource Name) of the DR service instance.
 * Constraints: The maximum length is `1048` characters. The minimum length is `20` characters. The value must match regular expression `/^.*$/`.
* `href` - (String) Resource reference.
 * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `dashboard_url` - (String) URL to the dashboard for managing the DR service instance in IBM Cloud.
 * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^https?:\/\/[a-zA-Z0-9\-._~:\/?#[\]@!$&'()*+,;=]+$/`.
* `orchestrator_location_type` - (String) The cloud location where your orchestator need to be created.
 * Constraints: The maximum length is `512` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9\-_]+$/`.
* `location_id` - (String) The location or data center identifier where the service instance is deployed.
  * Constraints: The maximum length is `512` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9\-_]+$/`.
* `ssh_key_name` - (String) The name of the SSH key used for deploying the orchestator.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-_]+$/`.
* `standby_ssh_key_name` - (String) The name of the SSH key used for deploying the standby orchestator.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-_]+$/`.
* `orchestrator_name` - (String) The username used for the orchestrator.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-_]+$/`.
* `orchestrator_workspace_id` - (String) The unique identifier orchestrator workspace.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-_]+$/`.
* `standby_orchestrator_name` - (String) The username for the standby orchestrator management interface.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-_]+$/`.
* `standby_orchestrator_workspace_id` - (String) The unique identifier of the standby orchestrator workspace.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-_]+$/`.
* `orchestrator_ha` - (Boolean) Indicates whether the orchestrator High Availability (HA) is enabled for the service instance.
* `resource_instance` - (String) The uniquie identifier of the associated IBM Cloud resource instance.
 * Constraints: The maximum length is `512` characters. The minimum length is `20` characters. The value must match regular expression `/^crn:v1:[a-zA-Z0-9\-_]+:public:resource-controller:[a-zA-Z0-9\-_]+:[a-zA-Z0-9\-_\/]+:[a-zA-Z0-9\-_]+::$/`.
* `secret_group` - (String) The secret group name in IBM Cloud Secrets Manager containing sensitive data for the service instance.
 * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-_]+$/`.
* `secret` - (String) The secret name or identifier used for retrieving credentials from secrets manager.
 * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-_]+$/`.
* `region_id` - (String) The power virtual server region where the service instance is deployed.
 * Constraints: The maximum length is `512` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9\-_]+$/`.
* `guid` - (String) The global unique identifier of the service instance.
 * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-zA-Z0-9\-_]+$/`.
* `machine_type` - (String) The machine type used for deploying orchestrator.
 * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-_]+$/`.
* `tier` - (String) The storage tier used for deploying primary orchestrator.
 * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-_]+$/`.
* `standby_tier` - (String) The storage tier used for deploying standby orchestrator.
 * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-_]+$/`.
* `standby_machine_type` - (String) The machine type used for deploying standby virtual machines.
 * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-_]+$/`.
* `tenant_name` - (String) The tenant name for MFA authentication API.
 * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9.\-_:]+$/`.
* `proxy_ip` - (String) Proxy IP for the Communication between Orchestrator and Service broker.
 * Constraints: The maximum length is `512` characters. The minimum length is `7` characters. The value must match regular expression `/^[a-zA-Z0-9.\-_:]+$/`.


## Import

You can import the `ibm_pdr_managedr` resource by using `id`.

The `id` property can be formed from `instance_id` and the service instance CRN in the following format:

<pre>
&lt;instance_id&gt;/&lt;instance_id&gt;
</pre>

* `instance_id`: A string in the format `123456d3-1122-3344-b67d-4389b44b7bf9`. Service Instance ID.
* `instance_id`: A string in the format `crn:v1:staging:public:power-dr-automation:global:a/a123456fb04cefb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::`. The CRN (Cloud Resource Name) of the DR service instance.

# Syntax

<pre>
$ terraform import ibm_pdr_managedr.pdr_managedr &lt;instance_id&gt;/&lt;instance_id&gt;
</pre>

---
