# Examples for IBM Cloud Backup and Recovery Migration API

These examples illustrate how to use the resources and data sources associated with IBM Cloud Backup and Recovery Migration API.

The following resources are supported:
* ibm_brs_migration
* ibm_brs_migration_host
* ibm_brs_migration_volume
* ibm_brs_migration_workload
* ibm_brs_migration_workload_run
* ibm_brs_migration_discover

The following data sources are supported:
* ibm_brs_migration
* ibm_brs_migrations
* ibm_brs_migration_host
* ibm_brs_migration_hosts
* ibm_brs_migration_volume
* ibm_brs_migration_volumes
* ibm_brs_migration_workload
* ibm_brs_migration_workloads
* ibm_brs_migration_workload_run
* ibm_brs_migration_workload_runs
* ibm_brs_migration_workload_history
* ibm_brs_migration_discover
* ibm_brs_migration_discovers

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## IBM Cloud Backup and Recovery Migration API resources

### Resource: ibm_brs_migration

```hcl
resource "ibm_brs_migration" "brs_migration_instance" {
  name = var.brs_migration_name
  brs_crn = var.brs_migration_brs_crn
  description = var.brs_migration_description
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | Human-readable name for this migration project. | `string` | true |
| brs_crn | CRN of the IBM Cloud Backup and Recovery instance backing this migration. | `string` | true |
| description | Optional human-readable description. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| crn | Server-assigned CRN for this migration resource. |
| state | Current lifecycle state of the migration project. |
| created_at | Timestamp when this migration was created. |
| updated_at | Timestamp of the last update to this migration. |

### Resource: ibm_brs_migration_host

```hcl
resource "ibm_brs_migration_host" "brs_migration_host_instance" {
  migration_id = var.brs_migration_host_migration_id
  type = var.brs_migration_host_type
  env = var.brs_migration_host_env
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| migration_id | The migration project ID (mgr-{uuid4} format). | `string` | true |
| type | Whether the host is a Virtual Server Instance or bare metal server. | `string` | true |
| env | Infrastructure environment this host belongs to. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| compute | Enriched compute details. Schema variant matches the sibling `env` field: `classic` → `ClassicComputeDetails`, `vpc` → `VPCComputeDetails`. |
| volume_attachments | Per-volume attachment records for this host. |
| migrated | Set to true when `POST /migrations/{migration_id}/workloads/{workload_id}/complete` is called for a workload that includes this host. |
| workload_id | ID of the workload this host is associated with. Null when the host has not been added to any workload yet. |
| registered_at | Timestamp when this host was registered in the Migration API. |
| host_id | Migration service host ID (host-{uuid4} format). |

### Resource: ibm_brs_migration_volume

```hcl
resource "ibm_brs_migration_volume" "brs_migration_volume_instance" {
  migration_id = var.brs_migration_volume_migration_id
  env = var.brs_migration_volume_env
  storage_type = var.brs_migration_volume_storage_type
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| migration_id | The migration project ID (mgr-{uuid4} format). | `string` | true |
| env | Infrastructure environment this volume belongs to. | `string` | true |
| storage_type | Storage type of the volume. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| attachment_state | Migration-service-computed attachment status. Set to `attached` when `host_attachments` is non-empty, `unattached` when empty. |
| storage | Enriched storage details. |
| host_attachments | Per-host attachment records for this volume. |
| migrated | Set to true when the workload covering this volume is completed. |
| workload_id | ID of the workload this volume is associated with. |
| registered_at | Timestamp when this volume was registered in the Migration API. |
| volume_id | Migration service volume ID (vol-{uuid4} format). |

### Resource: ibm_brs_migration_workload

```hcl
resource "ibm_brs_migration_workload" "brs_migration_workload_instance" {
  migration_id = var.brs_migration_workload_migration_id
  name = var.brs_migration_workload_name
  payloads = var.brs_migration_workload_payloads
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| migration_id | The migration project ID (mgr-{uuid4} format). | `string` | true |
| name | Human-readable name for this workload. | `string` | false |
| payloads | List of source-to-destination payload mappings. | `list()` | true |

#### Outputs

| Name | Description |
|------|-------------|
| volume_ownership_map | Server-computed map of `vol-*` volume_id → owning payload_id. Only populated when the same `source.volume_id` appears in two or more payloads. Null in all other cases. |
| state | Current lifecycle state of the workload. |
| schedule | Populated once setup completes. Null while settingUp or in created/failed state. |
| scheduling_error | Non-empty when `state` is `failed` due to async setup failure. |
| created_at | Timestamp when this workload was created. |
| updated_at | Timestamp of the last update to this workload. |
| completed_at | Timestamp when `POST /migrations/{migration_id}/workloads/{workload_id}/complete` finished. |
| workload_id | Migration service workload ID (wl-{uuid4} format). |

### Resource: ibm_brs_migration_workload_run

```hcl
resource "ibm_brs_migration_workload_run" "brs_migration_workload_run_instance" {
  migration_id = var.brs_migration_workload_run_migration_id
  workload_id = var.brs_migration_workload_run_workload_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| migration_id | The migration project ID (mgr-{uuid4} format). | `string` | true |
| workload_id | The migration service workload ID (wl-{uuid4} format). | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| operation_type | Whether this run is a backup or a restore operation. |
| run_type | Whether this run was triggered on-demand or by the schedule. |
| status | Current execution status of the run. |
| started_at | Time the run started (ISO 8601 UTC). |
| completed_at | Time the run completed. Null if still in progress. |
| duration_seconds | Wall-clock duration of the run in seconds. |
| message | Human-readable status message or error detail. |
| stats | Data-transfer statistics for a workload run or payload. |
| payload_results | Per-payload breakdown of the run. |
| run_id | Unique run ID (run-{uuid4} format). |

### Resource: ibm_brs_migration_discover

```hcl
resource "ibm_brs_migration_discover" "brs_migration_discover_instance" {
  migration_id = var.brs_migration_discover_migration_id
  env = var.brs_migration_discover_env
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| migration_id | The migration project ID (mgr-{uuid4} format). | `string` | true |
| env | Infrastructure environment being discovered. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| state | Current lifecycle state of the discovery job. |
| start_time | Start of the time window used for this discovery run. |
| end_time | End of the time window used for this discovery run. |
| message | Human-readable status or error message. |
| summary | Counts of discovered resources by compute and storage type. |
| job_id | Unique discovery job ID (job-{uuid4} format). |

## IBM Cloud Backup and Recovery Migration API data sources

### Data source: ibm_brs_migration

```hcl
data "ibm_brs_migration" "brs_migration_instance" {
  migration_id = var.data_brs_migration_migration_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| migration_id | The migration project ID (mgr-{uuid4} format). | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| name | Human-readable name for this migration project. |
| brs_crn | CRN of the IBM Cloud Backup and Recovery instance backing this migration. |
| crn | Server-assigned CRN for this migration resource. |
| state | Current lifecycle state of the migration project. |
| description | Optional human-readable description. |
| created_at | Timestamp when this migration was created. |
| updated_at | Timestamp of the last update to this migration. |

### Data source: ibm_brs_migrations

```hcl
data "ibm_brs_migrations" "brs_migrations_instance" {
  state = var.brs_migrations_state
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| state | Filter by migration state. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| migrations | The list of migration projects on this page. |

### Data source: ibm_brs_migration_host

```hcl
data "ibm_brs_migration_host" "brs_migration_host_instance" {
  migration_id = var.data_brs_migration_host_migration_id
  host_id = var.data_brs_migration_host_host_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| migration_id | The migration project ID (mgr-{uuid4} format). | `string` | true |
| host_id | The migration service host ID (host-{uuid4} format). | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| type | Whether the host is a Virtual Server Instance or bare metal server. |
| env | Infrastructure environment this host belongs to. |
| compute | Enriched compute details. Schema variant matches the sibling `env` field: `classic` → `ClassicComputeDetails`, `vpc` → `VPCComputeDetails`. |
| volume_attachments | Per-volume attachment records for this host. |
| migrated | Set to true when `POST /migrations/{migration_id}/workloads/{workload_id}/complete` is called for a workload that includes this host. |
| workload_id | ID of the workload this host is associated with. Null when the host has not been added to any workload yet. |
| registered_at | Timestamp when this host was registered in the Migration API. |

### Data source: ibm_brs_migration_hosts

```hcl
data "ibm_brs_migration_hosts" "brs_migration_hosts_instance" {
  migration_id = var.brs_migration_hosts_migration_id
  env = var.brs_migration_hosts_env
  type = var.brs_migration_hosts_type
  migrated = var.brs_migration_hosts_migrated
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| migration_id | The migration project ID (mgr-{uuid4} format). | `string` | true |
| env | Filter by infrastructure environment. | `string` | false |
| type | Filter by compute type. | `string` | false |
| migrated | Filter by migration status. | `bool` | false |

#### Outputs

| Name | Description |
|------|-------------|
| hosts | The list of registered hosts on this page. |

### Data source: ibm_brs_migration_volume

```hcl
data "ibm_brs_migration_volume" "brs_migration_volume_instance" {
  migration_id = var.data_brs_migration_volume_migration_id
  volume_id = var.data_brs_migration_volume_volume_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| migration_id | The migration project ID (mgr-{uuid4} format). | `string` | true |
| volume_id | The migration service volume ID (vol-{uuid4} format). | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| env | Infrastructure environment this volume belongs to. |
| storage_type | Storage type of the volume. |
| attachment_state | Migration-service-computed attachment status. Set to `attached` when `host_attachments` is non-empty, `unattached` when empty. |
| storage | Enriched storage details. |
| host_attachments | Per-host attachment records for this volume. |
| migrated | Set to true when the workload covering this volume is completed. |
| workload_id | ID of the workload this volume is associated with. |
| registered_at | Timestamp when this volume was registered in the Migration API. |

### Data source: ibm_brs_migration_volumes

```hcl
data "ibm_brs_migration_volumes" "brs_migration_volumes_instance" {
  migration_id = var.brs_migration_volumes_migration_id
  env = var.brs_migration_volumes_env
  storage_type = var.brs_migration_volumes_storage_type
  migrated = var.brs_migration_volumes_migrated
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| migration_id | The migration project ID (mgr-{uuid4} format). | `string` | true |
| env | Filter by infrastructure environment. | `string` | false |
| storage_type | Filter by storage type. | `string` | false |
| migrated | Filter by migration status. | `bool` | false |

#### Outputs

| Name | Description |
|------|-------------|
| volumes | The list of registered volumes on this page. |

### Data source: ibm_brs_migration_workload

```hcl
data "ibm_brs_migration_workload" "brs_migration_workload_instance" {
  migration_id = var.data_brs_migration_workload_migration_id
  workload_id = var.data_brs_migration_workload_workload_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| migration_id | The migration project ID (mgr-{uuid4} format). | `string` | true |
| workload_id | The migration service workload ID (wl-{uuid4} format). | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| name | Human-readable name for this workload. |
| volume_ownership_map | Server-computed map of `vol-*` volume_id → owning payload_id. Only populated when the same `source.volume_id` appears in two or more payloads. Null in all other cases. |
| state | Current lifecycle state of the workload. |
| payloads | List of source-to-destination payload mappings. |
| schedule | Populated once setup completes. Null while settingUp or in created/failed state. |
| scheduling_error | Non-empty when `state` is `failed` due to async setup failure. |
| created_at | Timestamp when this workload was created. |
| updated_at | Timestamp of the last update to this workload. |
| completed_at | Timestamp when `POST /migrations/{migration_id}/workloads/{workload_id}/complete` finished. |

### Data source: ibm_brs_migration_workloads

```hcl
data "ibm_brs_migration_workloads" "brs_migration_workloads_instance" {
  migration_id = var.brs_migration_workloads_migration_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| migration_id | The migration project ID (mgr-{uuid4} format). | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| workloads | The list of workloads on this page. |

### Data source: ibm_brs_migration_workload_run

```hcl
data "ibm_brs_migration_workload_run" "brs_migration_workload_run_instance" {
  migration_id = var.data_brs_migration_workload_run_migration_id
  workload_id = var.data_brs_migration_workload_run_workload_id
  run_id = var.data_brs_migration_workload_run_run_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| migration_id | The migration project ID (mgr-{uuid4} format). | `string` | true |
| workload_id | The migration service workload ID (wl-{uuid4} format). | `string` | true |
| run_id | The migration service run ID (run-{uuid4} format). | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| operation_type | Whether this run is a backup or a restore operation. |
| run_type | Whether this run was triggered on-demand or by the schedule. |
| status | Current execution status of the run. |
| started_at | Time the run started (ISO 8601 UTC). |
| completed_at | Time the run completed. Null if still in progress. |
| duration_seconds | Wall-clock duration of the run in seconds. |
| message | Human-readable status message or error detail. |
| stats | Data-transfer statistics for a workload run or payload. |
| payload_results | Per-payload breakdown of the run. |

### Data source: ibm_brs_migration_workload_runs

```hcl
data "ibm_brs_migration_workload_runs" "brs_migration_workload_runs_instance" {
  migration_id = var.brs_migration_workload_runs_migration_id
  workload_id = var.brs_migration_workload_runs_workload_id
  status = var.brs_migration_workload_runs_status
  run_type = var.brs_migration_workload_runs_run_type
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| migration_id | The migration project ID (mgr-{uuid4} format). | `string` | true |
| workload_id | The migration service workload ID (wl-{uuid4} format). | `string` | true |
| status | Filter by run status. | `list(string)` | false |
| run_type | Filter by run type (scheduled or on-demand). | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| runs | List of runs ordered by `startedAt` descending. |

### Data source: ibm_brs_migration_workload_history

```hcl
data "ibm_brs_migration_workload_history" "brs_migration_workload_history_instance" {
  migration_id = var.brs_migration_workload_history_migration_id
  workload_id = var.brs_migration_workload_history_workload_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| migration_id | The migration project ID (mgr-{uuid4} format). | `string` | true |
| workload_id | The migration service workload ID (wl-{uuid4} format). | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| history | Workload execution history entries on this page. |

### Data source: ibm_brs_migration_discover

```hcl
data "ibm_brs_migration_discover" "brs_migration_discover_instance" {
  migration_id = var.data_brs_migration_discover_migration_id
  job_id = var.data_brs_migration_discover_job_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| migration_id | The migration project ID (mgr-{uuid4} format). | `string` | true |
| job_id | The unique ID of the discovery job (job-{uuid4} format). | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| env | Infrastructure environment being discovered. |
| state | Current lifecycle state of the discovery job. |
| start_time | Start of the time window used for this discovery run. |
| end_time | End of the time window used for this discovery run. |
| message | Human-readable status or error message. |
| summary | Counts of discovered resources by compute and storage type. |

### Data source: ibm_brs_migration_discovers

```hcl
data "ibm_brs_migration_discovers" "brs_migration_discovers_instance" {
  migration_id = var.brs_migration_discovers_migration_id
  env = var.brs_migration_discovers_env
  state = var.brs_migration_discovers_state
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| migration_id | The migration project ID (mgr-{uuid4} format). | `string` | true |
| env | Filter discovery jobs by infrastructure environment. | `string` | false |
| state | Filter discovery jobs by current state. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| discover | Discovery jobs on this page. |

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
