# Gen2 Database Task Data Source Implementation

## Overview

This document describes the implementation of Gen2 support for the `ibm_database_task` data source in the IBM Cloud Terraform Provider. The implementation follows the backend pattern established in PR #6808 for database connections.

## Problem Statement

The existing `ibm_database_task` data source only supported Classic (Gen1) IBM Cloud Databases using the Cloud Databases V5 API. With the introduction of Gen2 databases, a new approach was needed because:

1. **Gen2 databases don't have traditional task APIs** - They use the Resource Controller API instead
2. **Different data models** - Gen2 uses `instance.LastOperation` while Classic uses dedicated task objects
3. **Need for backward compatibility** - Existing Classic database users should not be affected

## Solution Architecture

### Backend Pattern

Following PR #6808's connection implementation, we implemented a backend abstraction pattern:

```
┌─────────────────────────────────────────────────────────┐
│ data_source_ibm_database_task.go                        │
├─────────────────────────────────────────────────────────┤
│ • dataSourceIBMDatabaseTaskBackend (interface)          │
│ • pickDataSourceTaskBackend() → Gen2 detection          │
│ • dataSourceIBMDatabaseTaskRead() → delegates to backend│
│ • dataSourceIBMDatabaseTaskClassicBackend (Classic impl)│
└─────────────────────────────────────────────────────────┘
                         │
         ┌───────────────┴───────────────┐
         ▼                               ▼
┌──────────────────────┐      ┌──────────────────────────┐
│ Classic Backend      │      │ Gen2 Backend             │
├──────────────────────┤      ├──────────────────────────┤
│ Cloud Databases V5   │      │ Resource Controller V2   │
│ GetTask()            │      │ GetResourceInstance()    │
│ • task.Task.ID       │      │ • instance.LastOperation │
│ • task.Task.Status   │      │ • instance.State         │
│ • task.Task.Progress │      │ • Mapped to task fields  │
└──────────────────────┘      └──────────────────────────┘
```

## Implementation Details

### Files Modified/Created

1. **`data_source_ibm_database_task.go`**
   - Added backend interface and picker function
   - Refactored main read function to delegate to backends
   - Implemented Classic backend inline

2. **`data_source_ibm_database_task_gen2.go`** - NEW
   - Implements Gen2 backend using Resource Controller V2 API
   - Maps RC instance states to task statuses
   - Calculates progress from instance state

3. **`data_source_ibm_database_task_gen2_ac_test.go`** - NEW
   - Comprehensive acceptance tests for Gen2 functionality
   - Tests for error handling, status mapping, and progress calculation

### Key Components

#### Backend Interface
- Defines common `Read()` method for both Classic and Gen2 implementations
- Enables polymorphic behavior based on database generation

#### Backend Picker
- Automatically detects Gen2 vs Classic databases using `isGen2Plan()` helper
- Fetches instance details from Resource Controller API
- Returns appropriate backend implementation

#### Gen2 Backend Implementation
- Extracts task information from Resource Controller instance data
- Maps instance states to task statuses
- Calculates progress percentage from instance state

### Data Mapping

| Task Field | Classic Source | Gen2 Source |
|------------|----------------|-------------|
| `task_id` | `task.Task.ID` | `instance.ID` |
| `deployment_id` | `task.Task.DeploymentID` | `instance.ID` |
| `description` | `task.Task.Description` | `instance.LastOperation.Description` or Type |
| `status` | `task.Task.Status` | Mapped from `instance.State` |
| `progress_percent` | `task.Task.ProgressPercent` | Calculated from `instance.State` |
| `created_at` | `task.Task.CreatedAt` | `instance.UpdatedAt` or `CreatedAt` |

### State to Status Mapping

| RC Instance State | Task Status |
|-------------------|-------------|
| `active` | `completed` |
| `provisioning` | `running` |
| `in progress` | `running` |
| `failed` | `failed` |
| `inactive` | `queued` |
| `removed` | `completed` |

### Progress Calculation

| RC Instance State | Progress % |
|-------------------|------------|
| `active` | 100 |
| `provisioning` | 50 |
| `in progress` | 75 |
| `failed` / `removed` | 100 |
| `inactive` | 0 |

## Usage Examples

### Gen2 Database Task

```hcl
# Create a Gen2 PostgreSQL database
resource "ibm_database" "postgresql_gen2" {
  name              = "my-postgres-gen2"
  service           = "databases-for-postgresql"
  plan              = "standard-gen2"  # Gen2 plan
  location          = "us-south"
  
  group {
    group_id = "member"
    members {
      allocation_count = 2
    }
    host_flavor {
      id = "bx3d.4x20"
    }
    disk {
      allocation_mb = 10240
    }
  }
}

# Get task details for the Gen2 database
data "ibm_database_task" "task" {
  task_id = ibm_database.postgresql_gen2.id
}

output "task_status" {
  value = data.ibm_database_task.task.status
}

output "task_progress" {
  value = data.ibm_database_task.task.progress_percent
}
```

### Classic Database Task (unchanged)

```hcl
# Create a Classic PostgreSQL database
resource "ibm_database" "postgresql_classic" {
  name     = "my-postgres-classic"
  service  = "databases-for-postgresql"
  plan     = "standard"  # Classic plan
  location = "us-south"
}

# Get task details - automatically uses Classic backend
data "ibm_database_task" "task" {
  task_id = "task-id-from-operation"
}
```

## Acceptance Test Coverage

### Test Suite

| Test | Purpose | Duration |
|------|---------|----------|
| TestAccIBMDatabaseTaskGen2DataSourceRead | Validates all task fields populated correctly | ~45-60 min |
| TestAccIBMDatabaseTaskGen2DataSourceInvalidID | Tests error handling for invalid IDs | ~2-5 min |
| TestAccIBMDatabaseTaskGen2DataSourceStatusMapping | Verifies state-to-status mapping | ~45-60 min |
| TestAccIBMDatabaseTaskGen2DataSourceProgressPercent | Validates progress calculation | ~45-60 min |

### Coverage Summary

**Functional Coverage:**
- ✅ Gen2 database detection via `isGen2Plan()`
- ✅ Resource Controller API integration
- ✅ All data source fields (task_id, deployment_id, description, status, progress_percent, created_at)
- ✅ Error handling for invalid/non-existent resources
- ✅ State mapping (active, provisioning, failed, etc.)
- ✅ Progress calculation (0-100%)
- ✅ LastOperation field parsing
- ✅ Timestamp handling

**API Coverage:**
- ✅ ResourceControllerV2.GetResourceInstance
- ✅ Instance state retrieval
- ✅ LastOperation field access
- ✅ Plan ID validation

**Error Scenarios:**
- ✅ Invalid deployment ID
- ✅ Non-existent resource
- ✅ API failures
- ✅ Nil pointer handling

### Running Tests

```bash
# Set credentials
export IC_API_KEY="your-api-key"

# Run all Gen2 tests
TF_ACC=1 go test -v ./ibm/service/database -run TestAccIBMDatabaseTaskGen2 -timeout 120m

# Run specific test
TF_ACC=1 go test -v ./ibm/service/database -run TestAccIBMDatabaseTaskGen2DataSourceRead -timeout 120m
```

### Test Configuration

Tests use Gen2 PostgreSQL with:
- Plan: `standard-gen2`
- Location: `ca-mon`
- Host flavor: `bx3d.4x20`
- Disk: 10GB
- Members: 2

## Backward Compatibility

✅ **Fully backward compatible** - Classic database users experience no changes:
- Classic databases automatically use the Classic backend
- Same API calls and data structure
- No configuration changes required

## API Usage

### Classic Backend
- **API**: Cloud Databases V5
- **Endpoint**: `GET /deployments/{id}/tasks/{task_id}`
- **Authentication**: IAM token
- **Response**: Task object with all fields

### Gen2 Backend
- **API**: Resource Controller V2
- **Endpoint**: `GET /v2/resource_instances/{id}`
- **Authentication**: IAM token
- **Response**: Instance object with `LastOperation` field

## Error Handling

Both backends use consistent error handling with `flex.TerraformErrorf()` for uniform error messages and diagnostics.

## Debugging

Gen2 backend includes detailed debug logging for instance ID, state, and LastOperation details. Enable with: `export TF_LOG=DEBUG`

## Future Enhancements

Potential improvements for future iterations:

1. **Enhanced Progress Tracking**: More granular progress calculation based on operation type
2. **Operation History**: Support for retrieving multiple operations/tasks
3. **Real-time Updates**: Polling mechanism for long-running operations
4. **Custom Timeouts**: Per-operation timeout configuration

## References

- PR #6808: Gen2 Connection Data Source (pattern reference)
- PR #6802: Related Gen2 implementation
- IBM Cloud Databases API: https://cloud.ibm.com/apidocs/cloud-databases-api
- Resource Controller API: https://cloud.ibm.com/apidocs/resource-controller

## Contributors

Implementation follows the established patterns in the IBM Cloud Terraform Provider and maintains consistency with existing Gen2 implementations.