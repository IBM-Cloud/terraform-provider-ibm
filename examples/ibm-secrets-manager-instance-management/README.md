# Examples for IBM Cloud Secrets Manager Instance Management API

These examples illustrate how to use the resources and data sources associated with IBM Cloud Secrets Manager Instance Management API.

The following resources are supported:
* ibm_sm_destination
* ibm_sm_hsm_destination
* ibm_sm_hsm_ha_group
* ibm_sm_hsm_path
* ibm_sm_hsm_path_build
* ibm_sm_hsm_server_ca
* ibm_sm_hsm_client_certificate

The following data sources are supported:
* ibm_sm_instance
* ibm_sm_destinations
* ibm_sm_destination
* ibm_sm_hsm_destination
* ibm_sm_hsm_destinations
* ibm_sm_hsm_destination_health
* ibm_sm_hsm_ha_group
* ibm_sm_hsm_ha_groups
* ibm_sm_hsm_path
* ibm_sm_hsm_paths
* ibm_sm_hsm_server_ca
* ibm_sm_hsm_server_cas
* ibm_sm_hsm_client_certificate
* ibm_sm_hsm_client_certificates

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## IBM Cloud Secrets Manager Instance Management API resources

### Resource: ibm_sm_destination

```hcl
resource "ibm_sm_destination" "sm_destination_instance" {
  instance_id = var.sm_destination_instance_id
  name = var.sm_destination_name
  type = var.sm_destination_type
  description = var.sm_destination_description
  crn = var.sm_destination_crn
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| instance_id | Secrets Manager instance ID. | `string` | true |
| name | Destination name. | `string` | true |
| type | Destination type. | `string` | true |
| description | Optional description. | `string` | false |
| crn | IBM Cloud Database service instance CRN. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| href | The URL of the destination resource. |
| state | Destination state:- `not_started`: Initial state before the first provisioning attempt begins- `provisioning`: Provisioning in progress — poll until `succeeded` or `failed`- `succeeded`: Destination ready and usable- `failed`: Terminal state reached when provisioning or deletion fails. A `failed` destination still counts against the per-instance quota until deleted.- `deleting`: Deletion in progress. |
| created_at | Timestamp when the destination was created. |
| updated_at | Timestamp when the destination was last updated. |
| created_by | IAM identity that created the destination. |
| message | Human-readable message providing additional context about the current state. Present only when non-empty — set when `state` is `failed`, describing why provisioning or deletion failed. |
| destination_id | Destination ID. |

### Resource: ibm_sm_hsm_destination

```hcl
resource "ibm_sm_hsm_destination" "sm_hsm_destination_instance" {
  instance_id = var.sm_hsm_destination_instance_id
  name = var.sm_hsm_destination_name
  description = var.sm_hsm_destination_description
  vendor = var.sm_hsm_destination_vendor
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| instance_id | Secrets Manager instance ID. | `string` | true |
| name | Human-readable display name. | `string` | true |
| description | Optional description. Omitted if not set. | `string` | false |
| vendor | HSM vendor. Immutable after creation. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| href | URL of the destination resource. |
| account_id | IBM Cloud account ID. System-managed. Required for account policy setup. |
| provisioning_state | Lifecycle state of the destination:- `completed`: Destination created and available for configuration. Initial state on POST.- `deleting`: Hard delete initiated; asynchronous cleanup in progress.- `deletion_failed`: Asynchronous delete failed; the destination still exists. Retry DELETE to re-initiate. |
| health_state | AS OF (point-in-time) health state. Written exclusively by POST …/health and POST …/apply:- `unknown`: Health check not yet performed, or destination is deleting/deletion_failed.- `evaluating`: A health check (POST …/health) is in progress. Set on acceptance; replaced when the check completes.- `in_progress`: An apply operation (POST …/apply) is running. Set on acceptance; replaced by the health check result once apply completes (or set to `unknown` on failure).- `not_configured`: Prerequisites not met (needs ≥1 HA Group, ≥1 path with active non-expired Server CA, and ≥1 active non-expired client certificate).- `healthy`: All non-`not_configured` HSM HA Groups are `healthy` or `degraded`.- `degraded`: Some non-`not_configured` HSM HA Groups are `unhealthy`/`unknown` AND at least one is `healthy`.- `unhealthy`: All non-`not_configured` HSM HA Groups are `unhealthy` or `unknown`. |
| health_as_of | Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health. |
| created_at | Computed field. When the resource was created. |
| updated_at | Computed field. Omitted until first update. |
| created_by | IAM identity of creator. Computed field. |
| updated_by | IAM identity of last updater. Computed field. Omitted until first update. |
| destination_id | Destination ID. Immutable. System-managed. |

### Resource: ibm_sm_hsm_ha_group

```hcl
resource "ibm_sm_hsm_ha_group" "sm_hsm_ha_group_instance" {
  instance_id = var.sm_hsm_ha_group_instance_id
  destination_id = var.sm_hsm_ha_group_destination_id
  name = var.sm_hsm_ha_group_name
  description = var.sm_hsm_ha_group_description
  group_label = var.sm_hsm_ha_group_group_label
  group_serial = var.sm_hsm_ha_group_group_serial
  member_serials = var.sm_hsm_ha_group_member_serials
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| instance_id | Secrets Manager instance ID. | `string` | true |
| destination_id | Destination ID. | `string` | true |
| name | Human-readable display name. | `string` | true |
| description | Optional description. Omitted if not set. | `string` | false |
| group_label | HA Group virtual token label. Immutable after creation. | `string` | true |
| group_serial | HA Group serial number. Immutable after creation. | `string` | true |
| member_serials | Ordered list of HSM node member serial numbers for this HA Group. Max 100 unique members. | `list(string)` | true |

#### Outputs

| Name | Description |
|------|-------------|
| href | URL of the HA Group resource. |
| provisioning_state | Lifecycle state of the HA Group:- `completed`: HA Group has been created and its config has been staged in the HSM config file. Initial and only state.**Note:** New values may be added to this enumeration in future API versions. Clients MUST handle unknown values gracefully and MUST NOT fail on unrecognised states. |
| health_state | AS OF health state. Written exclusively by POST …/health and POST …/apply:- `unknown`: Initial state on creation. No health check has been performed yet.- `evaluating`: A health check is in progress for this resource.- `in_progress`: An apply operation is running across the destination.- `healthy`: Configured AND all member HSM nodes are healthy.- `degraded`: Configured AND any one member is unhealthy (but not all).- `unhealthy`: Configured AND all members are unhealthy or unknown.- `not_configured`: Prerequisites not met (destination needs ≥1 path with active non-expired Server CA AND ≥1 active non-expired client certificate). |
| health_as_of | Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health. |
| created_at | Computed field. |
| updated_at | Computed field. Omitted until first update. |
| created_by | Computed field. |
| updated_by | Computed field. Omitted until first update. |
| ha_group_id | HA Group ID. Immutable. System-managed. |

### Resource: ibm_sm_hsm_path

```hcl
resource "ibm_sm_hsm_path" "sm_hsm_path_instance" {
  instance_id = var.sm_hsm_path_instance_id
  destination_id = var.sm_hsm_path_destination_id
  name = var.sm_hsm_path_name
  description = var.sm_hsm_path_description
  crn = var.sm_hsm_path_crn
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| instance_id | Secrets Manager instance ID. | `string` | true |
| destination_id | Destination ID. | `string` | true |
| name | Human-readable display name. | `string` | true |
| description | Optional description. Omitted if not set. | `string` | false |
| crn | CRN of the gateway through which the HSM is reached. Immutable. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| href | URL of the path resource. |
| provisioning_state | Lifecycle state of the path:- `pending_server_ca`: Path created. Waiting for a Server CA to be added and activated.- `pending_account_policy`: Server CA active. Customer must review and confirm the Secrets Manager account policy for the gateway before proceeding.- `provisioning`: Connectivity setup in progress (async). Triggered by POST /build.- `pending_approval`: Connectivity created. Waiting for customer approval.- `completed`: Connectivity established. Path fully operational.- `vpe_provisioning_failed`: Connectivity setup failed. Call POST /reset then retry POST /build.- `crn_already_in_use`: Connectivity setup failed — the gateway CRN is already in use by another Secrets Manager instance in the same region. Obtain a new CRN and update via PATCH, then call POST /reset.- `approval_declined`: Customer declined the pending approval. Call POST /reset.- `deleting`: Delete initiated; async cleanup in progress.- `deletion_failed`: Async deletion failed. Retry DELETE to re-initiate. |
| health_state | AS OF health state. Written exclusively by POST …/health and POST …/apply:- `unknown`: Initial state on creation. Health check not yet performed, or path is deleting/deletion_failed.- `evaluating`: A health check is in progress for this path.- `in_progress`: An apply operation is running across the destination.- `no_valid_server_ca`: Pre-check failed — no Server CA is active, or the active Server CA has expired.- `not_configured`: Path is not yet provisioned (pending_account_policy, provisioning, pending_approval, approval_declined, crn_already_in_use, vpe_provisioning_failed) — AND active Server CA is valid. If active CA is missing or expired, no_valid_server_ca takes priority.- `healthy`: Active Server CA valid AND connectivity check to the HSM succeeded.- `unhealthy`: Active Server CA valid AND connectivity check to the HSM failed. |
| health_as_of | Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health. |
| created_at | Computed field. |
| updated_at | Computed field. Omitted until first update. |
| created_by | Computed field. |
| updated_by | Computed field. Omitted until first update. |
| path_id | Path ID. Immutable. System-managed. |

### Resource: ibm_sm_hsm_path_build

```hcl
resource "ibm_sm_hsm_path_build" "sm_hsm_path_build_instance" {
  instance_id = var.sm_hsm_path_build_instance_id
  destination_id = var.sm_hsm_path_build_destination_id
  path_id = var.sm_hsm_path_build_path_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| instance_id | Secrets Manager instance ID. | `string` | true |
| destination_id | Destination ID. | `string` | true |
| path_id | Path ID. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| provisioning_state | Always `provisioning` on a 202 response. |
| path_id | ID of the path being provisioned. |

### Resource: ibm_sm_hsm_server_ca

```hcl
resource "ibm_sm_hsm_server_ca" "sm_hsm_server_ca_instance" {
  instance_id = var.sm_hsm_server_ca_instance_id
  destination_id = var.sm_hsm_server_ca_destination_id
  path_id = var.sm_hsm_server_ca_path_id
  name = var.sm_hsm_server_ca_name
  description = var.sm_hsm_server_ca_description
  certificate = var.sm_hsm_server_ca_certificate
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| instance_id | Secrets Manager instance ID. | `string` | true |
| destination_id | Destination ID. | `string` | true |
| path_id | Path ID. | `string` | true |
| name | Human-readable display name. | `string` | true |
| description | Optional description. Omitted if not set. | `string` | false |
| certificate | PEM-encoded Server CA certificate. Immutable. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| expires_at | Expiry parsed from certificate NotAfter field. System-managed. Populated on creation. |
| status | Only one Server CA can be active per path. Starts inactive on creation. |
| health_state | AS OF health state of the Server CA certificate. Written exclusively by POST …/health:- `unknown`: Initial state on creation. Health check not yet performed.- `evaluating`: A health check is in progress for this Server CA.- `valid`: Certificate is within its validity period (expires_at is in the future).- `expired`: Certificate has passed its expires_at date. |
| health_as_of | Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health. |
| created_at | Computed field. |
| updated_at | Computed field. Omitted until first update. |
| created_by | Computed field. |
| updated_by | Computed field. Omitted until first update. |
| server_ca_id | Server CA ID. Immutable. |

### Resource: ibm_sm_hsm_client_certificate

```hcl
resource "ibm_sm_hsm_client_certificate" "sm_hsm_client_certificate_instance" {
  instance_id = var.sm_hsm_client_certificate_instance_id
  destination_id = var.sm_hsm_client_certificate_destination_id
  name = var.sm_hsm_client_certificate_name
  description = var.sm_hsm_client_certificate_description
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| instance_id | Secrets Manager instance ID. | `string` | true |
| destination_id | Destination ID. | `string` | true |
| name | Human-readable display name. | `string` | true |
| description | Optional description. Omitted if not set. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| href | URL of the client certificate resource. |
| status | Request-only. Set to `active` in a PATCH to activate this certificate. Never returned; read the current state from `provisioning_state`. |
| provisioning_state | Lifecycle state of the client certificate:- `csr_generated`: CSR and private key generated successfully. `csr` and `csr_cn` are populated. Initial state on POST. Waiting for signed_cert to be uploaded via PATCH.- `inactive`: signed_cert has been uploaded but the certificate has not been activated yet.- `active`: This certificate is the active client certificate for the destination. Only one can be active at a time. |
| health_state | AS OF health state of the client certificate. Written exclusively by POST …/health:- `unknown`: Initial state on creation, and for csr_generated certs with no signed_cert.- `evaluating`: A health check is in progress for this client certificate.- `valid`: Certificate is within its validity period (expires_at is in the future).- `expired`: Certificate has passed its expires_at date. |
| health_as_of | Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health. |
| csr | PEM-encoded CSR. Platform-generated. Immutable once set. Populated on creation (always present once the resource exists). |
| csr_cn | Common Name embedded in the CSR. Set to the SM instance UUID at generation time. Immutable once set. Populated on creation (always present once the resource exists). |
| signed_cert | CA-signed certificate (PEM). Omitted until uploaded via PATCH. |
| expires_at | Parsed from signed_cert NotAfter field. Omitted until signed_cert is uploaded. |
| created_at | Computed field. |
| updated_at | Computed field. Omitted until first update. |
| created_by | Computed field. |
| updated_by | Computed field. Omitted until first update. |
| cert_id | Certificate ID. Immutable. |

## IBM Cloud Secrets Manager Instance Management API data sources

### Data source: ibm_sm_instance

```hcl
data "ibm_sm_instance" "sm_instance_instance" {
  instance_id = var.sm_instance_instance_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | Secrets Manager instance ID. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| name | The instance name. |
| instance_crn | The instance CRN identifier. |
| plan | Instance plan name. |
| vault_cluster | Vault cluster information for Vault Dedicated instances. |
| endpoints | Instance endpoints for Vault Dedicated instances. |
| encryption | Vault encryption configuration for Vault Dedicated instances. |
| href | The URL of the instance resource. |

### Data source: ibm_sm_destinations

```hcl
data "ibm_sm_destinations" "sm_destinations_instance" {
  instance_id = var.sm_destinations_instance_id
  state = var.sm_destinations_state
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | Secrets Manager instance ID. | `string` | true |
| state | Filter by destination state. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| destinations | List of destinations. |
| total | Total number of destinations. Maximum 20 per instance. |

### Data source: ibm_sm_destination

```hcl
data "ibm_sm_destination" "sm_destination_instance" {
  instance_id = var.data_sm_destination_instance_id
  destination_id = var.data_sm_destination_destination_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | Secrets Manager instance ID. | `string` | true |
| destination_id | Destination ID. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| href | The URL of the destination resource. |
| name | Destination name. |
| type | Destination type. |
| description | Optional description. |
| state | Destination state:- `not_started`: Initial state before the first provisioning attempt begins- `provisioning`: Provisioning in progress — poll until `succeeded` or `failed`- `succeeded`: Destination ready and usable- `failed`: Terminal state reached when provisioning or deletion fails. A `failed` destination still counts against the per-instance quota until deleted.- `deleting`: Deletion in progress. |
| created_at | Timestamp when the destination was created. |
| updated_at | Timestamp when the destination was last updated. |
| created_by | IAM identity that created the destination. |
| message | Human-readable message providing additional context about the current state. Present only when non-empty — set when `state` is `failed`, describing why provisioning or deletion failed. |
| crn | IBM Cloud Database service instance CRN. |

### Data source: ibm_sm_hsm_destination

```hcl
data "ibm_sm_hsm_destination" "sm_hsm_destination_instance" {
  instance_id = var.data_sm_hsm_destination_instance_id
  destination_id = var.data_sm_hsm_destination_destination_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | Secrets Manager instance ID. | `string` | true |
| destination_id | Destination ID. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| href | URL of the destination resource. |
| account_id | IBM Cloud account ID. System-managed. Required for account policy setup. |
| name | Human-readable display name. |
| description | Optional description. Omitted if not set. |
| vendor | HSM vendor. Immutable after creation. |
| provisioning_state | Lifecycle state of the destination:- `completed`: Destination created and available for configuration. Initial state on POST.- `deleting`: Hard delete initiated; asynchronous cleanup in progress.- `deletion_failed`: Asynchronous delete failed; the destination still exists. Retry DELETE to re-initiate. |
| health_state | AS OF (point-in-time) health state. Written exclusively by POST …/health and POST …/apply:- `unknown`: Health check not yet performed, or destination is deleting/deletion_failed.- `evaluating`: A health check (POST …/health) is in progress. Set on acceptance; replaced when the check completes.- `in_progress`: An apply operation (POST …/apply) is running. Set on acceptance; replaced by the health check result once apply completes (or set to `unknown` on failure).- `not_configured`: Prerequisites not met (needs ≥1 HA Group, ≥1 path with active non-expired Server CA, and ≥1 active non-expired client certificate).- `healthy`: All non-`not_configured` HSM HA Groups are `healthy` or `degraded`.- `degraded`: Some non-`not_configured` HSM HA Groups are `unhealthy`/`unknown` AND at least one is `healthy`.- `unhealthy`: All non-`not_configured` HSM HA Groups are `unhealthy` or `unknown`. |
| health_as_of | Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health. |
| created_at | Computed field. When the resource was created. |
| updated_at | Computed field. Omitted until first update. |
| created_by | IAM identity of creator. Computed field. |
| updated_by | IAM identity of last updater. Computed field. Omitted until first update. |

### Data source: ibm_sm_hsm_destinations

```hcl
data "ibm_sm_hsm_destinations" "sm_hsm_destinations_instance" {
  instance_id = var.sm_hsm_destinations_instance_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | Secrets Manager instance ID. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| total_count | Total number of HSM Destinations in the collection. |
| hsm_destinations | The list of HSM Destinations. At most 5 destinations may exist per SM instance. |

### Data source: ibm_sm_hsm_destination_health

```hcl
data "ibm_sm_hsm_destination_health" "sm_hsm_destination_health_instance" {
  instance_id = var.sm_hsm_destination_health_instance_id
  destination_id = var.sm_hsm_destination_health_destination_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | Secrets Manager instance ID. | `string` | true |
| destination_id | Destination ID. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| href | URL of the health snapshot endpoint for this destination. |
| health_state | AS OF (point-in-time) health state. Written exclusively by POST …/health and POST …/apply:- `unknown`: Health check not yet performed, or destination is deleting/deletion_failed.- `evaluating`: A health check (POST …/health) is in progress. Set on acceptance; replaced when the check completes.- `in_progress`: An apply operation (POST …/apply) is running. Set on acceptance; replaced by the health check result once apply completes (or set to `unknown` on failure).- `not_configured`: Prerequisites not met (needs ≥1 HA Group, ≥1 path with active non-expired Server CA, and ≥1 active non-expired client certificate).- `healthy`: All non-`not_configured` HSM HA Groups are `healthy` or `degraded`.- `degraded`: Some non-`not_configured` HSM HA Groups are `unhealthy`/`unknown` AND at least one is `healthy`.- `unhealthy`: All non-`not_configured` HSM HA Groups are `unhealthy` or `unknown`. |
| health_as_of | Timestamp of this health check — always the current time of the live probe. |
| client_certificates | All client certificates for this destination (active, inactive, csr_generated). health_state freshly evaluated for all certs with a signed_cert. |
| ha_groups | All HSM HA Groups for this destination. Empty array if none exist. |
| paths | All paths for this destination. Empty array if none exist. |

### Data source: ibm_sm_hsm_ha_group

```hcl
data "ibm_sm_hsm_ha_group" "sm_hsm_ha_group_instance" {
  instance_id = var.data_sm_hsm_ha_group_instance_id
  destination_id = var.data_sm_hsm_ha_group_destination_id
  ha_group_id = var.data_sm_hsm_ha_group_ha_group_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | Secrets Manager instance ID. | `string` | true |
| destination_id | Destination ID. | `string` | true |
| ha_group_id | The identifier of the HA Group to operate on. Matches the `id` field in the `HSMHAGroup` response schema. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| href | URL of the HA Group resource. |
| name | Human-readable display name. |
| description | Optional description. Omitted if not set. |
| provisioning_state | Lifecycle state of the HA Group:- `completed`: HA Group has been created and its config has been staged in the HSM config file. Initial and only state.**Note:** New values may be added to this enumeration in future API versions. Clients MUST handle unknown values gracefully and MUST NOT fail on unrecognised states. |
| health_state | AS OF health state. Written exclusively by POST …/health and POST …/apply:- `unknown`: Initial state on creation. No health check has been performed yet.- `evaluating`: A health check is in progress for this resource.- `in_progress`: An apply operation is running across the destination.- `healthy`: Configured AND all member HSM nodes are healthy.- `degraded`: Configured AND any one member is unhealthy (but not all).- `unhealthy`: Configured AND all members are unhealthy or unknown.- `not_configured`: Prerequisites not met (destination needs ≥1 path with active non-expired Server CA AND ≥1 active non-expired client certificate). |
| health_as_of | Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health. |
| group_label | HA Group virtual token label. Immutable after creation. |
| group_serial | HA Group serial number. Immutable after creation. |
| member_serials | Ordered list of HSM node member serial numbers for this HA Group. Max 100 unique members. |
| created_at | Computed field. |
| updated_at | Computed field. Omitted until first update. |
| created_by | Computed field. |
| updated_by | Computed field. Omitted until first update. |

### Data source: ibm_sm_hsm_ha_groups

```hcl
data "ibm_sm_hsm_ha_groups" "sm_hsm_ha_groups_instance" {
  instance_id = var.sm_hsm_ha_groups_instance_id
  destination_id = var.sm_hsm_ha_groups_destination_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | Secrets Manager instance ID. | `string` | true |
| destination_id | Destination ID. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| ha_groups | HSM HA Groups for this destination. At most 20 HSM HA Groups may exist per destination. |

### Data source: ibm_sm_hsm_path

```hcl
data "ibm_sm_hsm_path" "sm_hsm_path_instance" {
  instance_id = var.data_sm_hsm_path_instance_id
  destination_id = var.data_sm_hsm_path_destination_id
  path_id = var.data_sm_hsm_path_path_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | Secrets Manager instance ID. | `string` | true |
| destination_id | Destination ID. | `string` | true |
| path_id | Path ID. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| href | URL of the path resource. |
| name | Human-readable display name. |
| description | Optional description. Omitted if not set. |
| crn | CRN of the gateway through which the HSM is reached. Immutable. |
| provisioning_state | Lifecycle state of the path:- `pending_server_ca`: Path created. Waiting for a Server CA to be added and activated.- `pending_account_policy`: Server CA active. Customer must review and confirm the Secrets Manager account policy for the gateway before proceeding.- `provisioning`: Connectivity setup in progress (async). Triggered by POST /build.- `pending_approval`: Connectivity created. Waiting for customer approval.- `completed`: Connectivity established. Path fully operational.- `vpe_provisioning_failed`: Connectivity setup failed. Call POST /reset then retry POST /build.- `crn_already_in_use`: Connectivity setup failed — the gateway CRN is already in use by another Secrets Manager instance in the same region. Obtain a new CRN and update via PATCH, then call POST /reset.- `approval_declined`: Customer declined the pending approval. Call POST /reset.- `deleting`: Delete initiated; async cleanup in progress.- `deletion_failed`: Async deletion failed. Retry DELETE to re-initiate. |
| health_state | AS OF health state. Written exclusively by POST …/health and POST …/apply:- `unknown`: Initial state on creation. Health check not yet performed, or path is deleting/deletion_failed.- `evaluating`: A health check is in progress for this path.- `in_progress`: An apply operation is running across the destination.- `no_valid_server_ca`: Pre-check failed — no Server CA is active, or the active Server CA has expired.- `not_configured`: Path is not yet provisioned (pending_account_policy, provisioning, pending_approval, approval_declined, crn_already_in_use, vpe_provisioning_failed) — AND active Server CA is valid. If active CA is missing or expired, no_valid_server_ca takes priority.- `healthy`: Active Server CA valid AND connectivity check to the HSM succeeded.- `unhealthy`: Active Server CA valid AND connectivity check to the HSM failed. |
| health_as_of | Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health. |
| created_at | Computed field. |
| updated_at | Computed field. Omitted until first update. |
| created_by | Computed field. |
| updated_by | Computed field. Omitted until first update. |

### Data source: ibm_sm_hsm_paths

```hcl
data "ibm_sm_hsm_paths" "sm_hsm_paths_instance" {
  instance_id = var.sm_hsm_paths_instance_id
  destination_id = var.sm_hsm_paths_destination_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | Secrets Manager instance ID. | `string` | true |
| destination_id | Destination ID. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| paths | Paths for this destination. At most 20 paths may exist per destination. |

### Data source: ibm_sm_hsm_server_ca

```hcl
data "ibm_sm_hsm_server_ca" "sm_hsm_server_ca_instance" {
  instance_id = var.data_sm_hsm_server_ca_instance_id
  destination_id = var.data_sm_hsm_server_ca_destination_id
  path_id = var.data_sm_hsm_server_ca_path_id
  server_ca_id = var.data_sm_hsm_server_ca_server_ca_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | Secrets Manager instance ID. | `string` | true |
| destination_id | Destination ID. | `string` | true |
| path_id | Path ID. | `string` | true |
| server_ca_id | The identifier of the Server CA to operate on. Matches the `id` field in the `ServerCa` response schema. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| name | Human-readable display name. |
| description | Optional description. Omitted if not set. |
| certificate | PEM-encoded Server CA certificate. Immutable. |
| expires_at | Expiry parsed from certificate NotAfter field. System-managed. Populated on creation. |
| status | Only one Server CA can be active per path. Starts inactive on creation. |
| health_state | AS OF health state of the Server CA certificate. Written exclusively by POST …/health:- `unknown`: Initial state on creation. Health check not yet performed.- `evaluating`: A health check is in progress for this Server CA.- `valid`: Certificate is within its validity period (expires_at is in the future).- `expired`: Certificate has passed its expires_at date. |
| health_as_of | Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health. |
| created_at | Computed field. |
| updated_at | Computed field. Omitted until first update. |
| created_by | Computed field. |
| updated_by | Computed field. Omitted until first update. |

### Data source: ibm_sm_hsm_server_cas

```hcl
data "ibm_sm_hsm_server_cas" "sm_hsm_server_cas_instance" {
  instance_id = var.sm_hsm_server_cas_instance_id
  destination_id = var.sm_hsm_server_cas_destination_id
  path_id = var.sm_hsm_server_cas_path_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | Secrets Manager instance ID. | `string` | true |
| destination_id | Destination ID. | `string` | true |
| path_id | Path ID. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| server_cas | The list of Server CA certificates for this path. |

### Data source: ibm_sm_hsm_client_certificate

```hcl
data "ibm_sm_hsm_client_certificate" "sm_hsm_client_certificate_instance" {
  instance_id = var.data_sm_hsm_client_certificate_instance_id
  destination_id = var.data_sm_hsm_client_certificate_destination_id
  cert_id = var.data_sm_hsm_client_certificate_cert_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | Secrets Manager instance ID. | `string` | true |
| destination_id | Destination ID. | `string` | true |
| cert_id | The identifier of the Client Certificate to operate on. Matches the `id` field in the `ClientCertificate` response schema. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| href | URL of the client certificate resource. |
| name | Human-readable display name. |
| description | Optional description. Omitted if not set. |
| status | Request-only. Set to `active` in a PATCH to activate this certificate. Never returned; read the current state from `provisioning_state`. |
| provisioning_state | Lifecycle state of the client certificate:- `csr_generated`: CSR and private key generated successfully. `csr` and `csr_cn` are populated. Initial state on POST. Waiting for signed_cert to be uploaded via PATCH.- `inactive`: signed_cert has been uploaded but the certificate has not been activated yet.- `active`: This certificate is the active client certificate for the destination. Only one can be active at a time. |
| health_state | AS OF health state of the client certificate. Written exclusively by POST …/health:- `unknown`: Initial state on creation, and for csr_generated certs with no signed_cert.- `evaluating`: A health check is in progress for this client certificate.- `valid`: Certificate is within its validity period (expires_at is in the future).- `expired`: Certificate has passed its expires_at date. |
| health_as_of | Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health. |
| csr | PEM-encoded CSR. Platform-generated. Immutable once set. Populated on creation (always present once the resource exists). |
| csr_cn | Common Name embedded in the CSR. Set to the SM instance UUID at generation time. Immutable once set. Populated on creation (always present once the resource exists). |
| signed_cert | CA-signed certificate (PEM). Omitted until uploaded via PATCH. |
| expires_at | Parsed from signed_cert NotAfter field. Omitted until signed_cert is uploaded. |
| created_at | Computed field. |
| updated_at | Computed field. Omitted until first update. |
| created_by | Computed field. |
| updated_by | Computed field. Omitted until first update. |

### Data source: ibm_sm_hsm_client_certificates

```hcl
data "ibm_sm_hsm_client_certificates" "sm_hsm_client_certificates_instance" {
  instance_id = var.sm_hsm_client_certificates_instance_id
  destination_id = var.sm_hsm_client_certificates_destination_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | Secrets Manager instance ID. | `string` | true |
| destination_id | Destination ID. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| client_certificates | The list of client certificates for this destination. |

## Assumptions

1. TODO

## Notes

1. TODO

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |
