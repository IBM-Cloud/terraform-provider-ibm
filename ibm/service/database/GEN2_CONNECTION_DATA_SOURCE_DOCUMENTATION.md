# IBM Cloud Terraform Provider - Database Connection Gen2 Data Source Support

## Overview

This documentation covers the enhancement to the IBM Cloud Terraform Provider that adds support for **Generation 2 (Gen2) database instances** in the `ibm_database_connection` data source. This update enables users to seamlessly retrieve connection information for both Classic and Gen2 database deployments using the same data source with automatic detection and appropriate handling of generation-specific connection retrieval methods.

---

## Table of Contents

- [What's New](#whats-new)
- [Understanding Classic vs Gen2 Connection Retrieval](#understanding-classic-vs-gen2-connection-retrieval)
- [Architecture Overview](#architecture-overview)
- [Important Differences for Gen2](#important-differences-for-gen2)
- [Usage Examples](#usage-examples)
- [Migration Guide](#migration-guide)
- [Supported Attributes](#supported-attributes)
- [Testing](#testing)
- [Troubleshooting](#troubleshooting)
- [Best Practices](#best-practices)
- [Additional Resources](#additional-resources)

---

## What's New

### Gen2 Database Connection Support in Data Source

The `ibm_database_connection` data source now **automatically detects** whether a database instance is Classic or Gen2 and retrieves connection information using the appropriate method. This provides a unified interface for accessing database connection details regardless of their generation.

### Key Features

✅ **Automatic Detection**: The provider automatically identifies whether a database is Classic or Gen2 based on the plan type  
✅ **Unified Interface**: Use the same `ibm_database_connection` data source for both Classic and Gen2 instances  
✅ **Resource Key Integration**: Gen2 connections are retrieved through IBM Cloud resource keys  
✅ **Backward Compatible**: Existing Terraform configurations continue to work without modification  
✅ **Flexible Key Selection**: Automatically selects the appropriate resource key or falls back to the first available key  
✅ **Comprehensive Testing**: Full test coverage including acceptance tests and unit tests for transformation logic  

---

## Understanding Classic vs Gen2 Connection Retrieval

### Classic Databases (Standard Plans)

- **Connection Method**: Direct API access via Cloud Databases V5 API
- **Authentication**: Uses deployment ID, user type, user ID, and endpoint type
- **User Management**: Direct user creation and management through the database API
- **Connection Retrieval**: Fetches connection strings directly from the database deployment

**Example Classic Flow**:
```
User Request → Cloud Databases API → Get Connection for User → Return Connection Details
```

### Gen2 Databases (Standard-Gen2 Plans)

- **Connection Method**: Resource Controller API with resource keys
- **Authentication**: Uses resource keys as the credential mechanism
- **User Management**: Credentials are managed through IBM Cloud resource keys
- **Connection Retrieval**: Extracts connection information from resource key credentials

**Example Gen2 Flow**:
```
User Request → Resource Controller API → List Resource Keys → Extract Connection from Key Credentials
```

### Key Differences

| Aspect | Classic | Gen2 |
|--------|---------|------|
| **API Used** | Cloud Databases V5 | Resource Controller V2 |
| **Credential Source** | Database user | Resource key |
| **User ID Parameter** | Database username | Resource key name |
| **Connection Structure** | Direct from API | Nested in key credentials |
| **Key Requirement** | Not required | Required (must exist) |

---

## Architecture Overview

### Backend Selection Pattern

The data source uses a **strategy pattern** to handle different database generations:

```
User Query → Get Deployment ID → Fetch Instance → Check Plan Type → Select Backend
                                                                          ↓
                                                    ┌─────────────────────┴─────────────────┐
                                                    ↓                                       ↓
                                          Classic Backend                            Gen2 Backend
                                    (Cloud Databases API)                    (Resource Controller API)
                                          ↓                                              ↓
                                    Get Connection                              List Resource Keys
                                    for User/Endpoint                                    ↓
                                                                              Find Matching Key
                                                                                         ↓
                                                                              Extract Connection
                                                                              from Credentials
```

### Implementation Components

1. **`data_source_ibm_database_connection.go`**: Main data source with backend selection logic
2. **`data_source_ibm_database_connection_classic.go`**: Classic database backend implementation (existing)
3. **`data_source_ibm_database_connection_gen2.go`**: Gen2 database backend implementation (new)
4. **`data_source_ibm_database_connection_gen2_ac_test.go`**: Acceptance tests for Gen2 scenarios
5. **`data_source_ibm_database_connection_gen2_unit_test.go`**: Unit tests for transformation functions

### Gen2 Backend Implementation Details

The Gen2 backend (`data_source_ibm_database_connection_gen2.go`) implements:

- **Resource Key Discovery**: Lists all resource keys for the deployment
- **Key Selection Logic**: 
  - First attempts to find a key matching the `user_id` parameter
  - Falls back to the first available key if no match is found
  - Returns an error if no keys exist
- **Connection Transformation**: Converts Gen2 API structure to Terraform schema format
- **Multiple Connection Types**: Supports all database connection protocols (postgres, mongodb, redis, etc.)

---

## Important Differences for Gen2

### Resource Key Requirement

**Gen2 databases require resource keys to access connection information**. Unlike Classic databases where connection information is directly available through the database API, Gen2 stores connection details in resource key credentials.

#### Creating a Resource Key

Before querying connection information for a Gen2 database, you must create a resource key:

```hcl
resource "ibm_resource_key" "database_key" {
  name                 = "my-database-key"
  resource_instance_id = ibm_database.my_gen2_db.id
}
```

### User ID Parameter Behavior

The `user_id` parameter has different meanings for Classic and Gen2:

| Generation | user_id Meaning | Example |
|------------|-----------------|---------|
| **Classic** | Database username | `admin`, `user123` |
| **Gen2** | Resource key name | `my-database-key`, `app-credentials` |

### Fallback Behavior

If the specified `user_id` (resource key name) is not found for a Gen2 database, the data source will:

1. Log a debug message indicating the key was not found
2. Use the first available resource key
3. Update the `user_id` attribute in state to reflect the actual key used

This fallback behavior ensures flexibility while maintaining transparency about which credentials are being used.

### Connection Data Structure

Gen2 connection information is nested under a `connection` key in the resource key credentials:

```json
{
  "connection": {
    "postgres": { ... },
    "cli": { ... }
  }
}
```

The Gen2 backend automatically handles this nesting and extracts the appropriate connection types.

---

## Usage Examples

### Example 1: Basic Gen2 Connection Query

```hcl
# Create a Gen2 PostgreSQL database
resource "ibm_database" "postgres_gen2" {
  name              = "my-postgres-gen2"
  service           = "databases-for-postgresql"
  plan              = "standard-gen2"
  location          = "us-south"
  resource_group_id = data.ibm_resource_group.default.id

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

# Create a resource key for connection credentials
resource "ibm_resource_key" "postgres_key" {
  name                 = "postgres-credentials"
  resource_instance_id = ibm_database.postgres_gen2.id
}

# Query connection information
data "ibm_database_connection" "postgres_connection" {
  deployment_id = ibm_database.postgres_gen2.id
  user_type     = "database"
  user_id       = ibm_resource_key.postgres_key.name
  endpoint_type = "private"

  depends_on = [ibm_resource_key.postgres_key]
}

# Use the connection information
output "postgres_connection_string" {
  value     = data.ibm_database_connection.postgres_connection.postgres[0].composed[0]
  sensitive = true
}

output "postgres_host" {
  value = data.ibm_database_connection.postgres_connection.postgres[0].hosts[0].hostname
}

output "postgres_port" {
  value = data.ibm_database_connection.postgres_connection.postgres[0].hosts[0].port
}
```

### Example 2: Multiple Resource Keys with Explicit Selection

```hcl
# Create multiple resource keys for different purposes
resource "ibm_resource_key" "app_key" {
  name                 = "application-credentials"
  resource_instance_id = ibm_database.postgres_gen2.id
}

resource "ibm_resource_key" "admin_key" {
  name                 = "admin-credentials"
  resource_instance_id = ibm_database.postgres_gen2.id
}

# Query connection using specific key
data "ibm_database_connection" "app_connection" {
  deployment_id = ibm_database.postgres_gen2.id
  user_type     = "database"
  user_id       = ibm_resource_key.app_key.name
  endpoint_type = "private"

  depends_on = [
    ibm_resource_key.app_key,
    ibm_resource_key.admin_key
  ]
}
```

### Example 3: Gen2 MongoDB Connection

```hcl
resource "ibm_database" "mongodb_gen2" {
  name              = "my-mongodb-gen2"
  service           = "databases-for-mongodb"
  plan              = "standard-gen2"
  location          = "us-south"
  resource_group_id = data.ibm_resource_group.default.id

  group {
    group_id = "member"
    members {
      allocation_count = 3
    }
    host_flavor {
      id = "bx3d.4x20"
    }
    disk {
      allocation_mb = 20480
    }
  }
}

resource "ibm_resource_key" "mongodb_key" {
  name                 = "mongodb-credentials"
  resource_instance_id = ibm_database.mongodb_gen2.id
}

data "ibm_database_connection" "mongodb_connection" {
  deployment_id = ibm_database.mongodb_gen2.id
  user_type     = "database"
  user_id       = ibm_resource_key.mongodb_key.name
  endpoint_type = "private"

  depends_on = [ibm_resource_key.mongodb_key]
}

output "mongodb_connection_string" {
  value     = data.ibm_database_connection.mongodb_connection.mongodb[0].composed[0]
  sensitive = true
}

output "mongodb_replica_set" {
  value = lookup(
    data.ibm_database_connection.mongodb_connection.mongodb[0].query_options,
    "replicaSet",
    ""
  )
}
```

### Example 4: Fallback to First Available Key

```hcl
# This configuration will use the first available key
# if "nonexistent-key" is not found
data "ibm_database_connection" "fallback_connection" {
  deployment_id = ibm_database.postgres_gen2.id
  user_type     = "database"
  user_id       = "nonexistent-key"  # Will fallback to first available key
  endpoint_type = "private"

  depends_on = [ibm_resource_key.postgres_key]
}

# The actual key name used will be reflected in the state
output "actual_key_used" {
  value = data.ibm_database_connection.fallback_connection.user_id
}
```

### Example 5: Using CLI Connection Information

```hcl
data "ibm_database_connection" "postgres_connection" {
  deployment_id = ibm_database.postgres_gen2.id
  user_type     = "database"
  user_id       = ibm_resource_key.postgres_key.name
  endpoint_type = "private"

  depends_on = [ibm_resource_key.postgres_key]
}

# Extract CLI connection details
output "cli_command" {
  value     = data.ibm_database_connection.postgres_connection.cli[0].composed[0]
  sensitive = true
}

output "cli_binary" {
  value = data.ibm_database_connection.postgres_connection.cli[0].bin
}

output "cli_environment" {
  value     = data.ibm_database_connection.postgres_connection.cli[0].environment
  sensitive = true
}
```

### Example 6: Conditional Logic Based on Plan Type

```hcl
locals {
  is_gen2 = can(regex("gen2", ibm_database.my_database.plan))
}

# Create resource key only for Gen2
resource "ibm_resource_key" "db_key" {
  count                = local.is_gen2 ? 1 : 0
  name                 = "database-key"
  resource_instance_id = ibm_database.my_database.id
}

data "ibm_database_connection" "db_connection" {
  deployment_id = ibm_database.my_database.id
  user_type     = "database"
  user_id       = local.is_gen2 ? ibm_resource_key.db_key[0].name : "admin"
  endpoint_type = "private"

  depends_on = [ibm_resource_key.db_key]
}
```

### Example 7: Multiple Connection Types (Redis)

```hcl
resource "ibm_database" "redis_gen2" {
  name              = "my-redis-gen2"
  service           = "databases-for-redis"
  plan              = "standard-gen2"
  location          = "us-south"
  resource_group_id = data.ibm_resource_group.default.id

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

resource "ibm_resource_key" "redis_key" {
  name                 = "redis-credentials"
  resource_instance_id = ibm_database.redis_gen2.id
}

data "ibm_database_connection" "redis_connection" {
  deployment_id = ibm_database.redis_gen2.id
  user_type     = "database"
  user_id       = ibm_resource_key.redis_key.name
  endpoint_type = "private"

  depends_on = [ibm_resource_key.redis_key]
}

output "redis_connection_string" {
  value     = data.ibm_database_connection.redis_connection.rediss[0].composed[0]
  sensitive = true
}

output "redis_host" {
  value = data.ibm_database_connection.redis_connection.rediss[0].hosts[0].hostname
}
```

---

## Migration Guide

### Upgrading from Classic to Gen2

When migrating from a Classic database to Gen2, you need to update your Terraform configuration to include resource key creation.

#### Step 1: Update Database Plan

```hcl
# Before (Classic)
resource "ibm_database" "my_database" {
  name     = "my-database"
  service  = "databases-for-postgresql"
  plan     = "standard"  # Classic plan
  location = "us-south"
  # ... other configuration
}

# After (Gen2)
resource "ibm_database" "my_database" {
  name     = "my-database"
  service  = "databases-for-postgresql"
  plan     = "standard-gen2"  # Gen2 plan
  location = "us-south"
  # ... other configuration with host_flavor
  
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
```

#### Step 2: Add Resource Key Creation

```hcl
# New resource required for Gen2
resource "ibm_resource_key" "database_key" {
  name                 = "database-credentials"
  resource_instance_id = ibm_database.my_database.id
}
```

#### Step 3: Update Connection Data Source

```hcl
# Before (Classic)
data "ibm_database_connection" "connection" {
  deployment_id = ibm_database.my_database.id
  user_type     = "database"
  user_id       = "admin"  # Database username
  endpoint_type = "private"
}

# After (Gen2)
data "ibm_database_connection" "connection" {
  deployment_id = ibm_database.my_database.id
  user_type     = "database"
  user_id       = ibm_resource_key.database_key.name  # Resource key name
  endpoint_type = "private"
  
  depends_on = [ibm_resource_key.database_key]
}
```

#### Step 4: Update Dependent Resources

Ensure any resources or modules that depend on the connection information are updated to handle the new structure if needed. The connection attributes remain the same, so most downstream usage should work without changes.

### State Management During Migration

When migrating from Classic to Gen2:

1. **Plan Carefully**: Review the Terraform plan to understand what will be recreated
2. **Backup Data**: Ensure you have backups before destroying the Classic instance
3. **Update All References**: Update all places where the database connection is referenced
4. **Test Thoroughly**: Test the new Gen2 configuration in a non-production environment first

---

## Supported Attributes

### Input Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `deployment_id` | string | Yes | The CRN or ID of the database deployment |
| `user_type` | string | Yes | Type of user (typically "database") |
| `user_id` | string | Yes | For Gen2: Resource key name. For Classic: Database username |
| `endpoint_type` | string | Yes | Endpoint type: "public" or "private" |
| `certificate_root` | string | No | Optional certificate root path |

### Output Attributes - Connection Types

All connection type attributes are **computed** and available for both Classic and Gen2 (when applicable):

#### PostgreSQL Connection (`postgres`)

| Attribute | Type | Description |
|-----------|------|-------------|
| `type` | string | Connection type (e.g., "uri") |
| `composed` | list(string) | Complete connection strings |
| `scheme` | string | Connection scheme (e.g., "postgres") |
| `hosts` | list(object) | List of host objects with `hostname` and `port` |
| `path` | string | Database path |
| `database` | string | Database name |
| `authentication` | list(object) | Authentication details with `method`, `username`, `password` |
| `certificate` | list(object) | Certificate information with `name` and `certificate_base64` |
| `query_options` | map(string) | Query parameters as key-value pairs |

#### MongoDB Connection (`mongodb`)

Same structure as PostgreSQL, with MongoDB-specific values.

#### Redis Connection (`rediss`)

Same structure as PostgreSQL, with Redis-specific values.

#### MySQL Connection (`mysql`)

Same structure as PostgreSQL, with MySQL-specific values.

#### RabbitMQ Connections

- `amqps`: AMQP over SSL connection
- `mqtts`: MQTT over SSL connection
- `stomp_ssl`: STOMP over SSL connection
- `https`: HTTPS connection (for management API)

#### Elasticsearch Connections

- `https`: HTTPS connection
- `grpc`: gRPC connection

#### DataStax Connections

- `bi_connector`: BI Connector connection
- `analytics`: Analytics connection
- `ops_manager`: Ops Manager connection
- `emp`: EMP connection

#### CLI Connection (`cli`)

| Attribute | Type | Description |
|-----------|------|-------------|
| `type` | string | Connection type ("cli") |
| `composed` | list(string) | Complete CLI commands |
| `bin` | string | Binary/executable name |
| `arguments` | list(string) | Command-line arguments |
| `environment` | map(string) | Environment variables |
| `certificate` | list(object) | Certificate information |

---

## Testing

### Test Coverage

The Gen2 connection data source implementation includes comprehensive testing:

#### Acceptance Tests (`data_source_ibm_database_connection_gen2_ac_test.go`)

1. **TestAccIBMDatabaseConnectionGen2DataSourceRead**
   - Creates a Gen2 PostgreSQL database
   - Creates a resource key
   - Queries connection information
   - Validates all connection attributes are populated

2. **TestAccIBMDatabaseConnectionGen2DataSourceInvalidID**
   - Tests error handling for invalid deployment IDs
   - Validates appropriate error messages

3. **TestAccIBMDatabaseConnectionGen2DataSourceMissingResourceKey**
   - Creates a Gen2 database without resource keys
   - Validates error message when no keys exist

4. **TestAccIBMDatabaseConnectionGen2DataSourceFallsBackToFirstKey**
   - Creates multiple resource keys
   - Requests a non-existent key name
   - Validates fallback to first available key
   - Confirms `user_id` is updated to reflect actual key used

#### Unit Tests (`data_source_ibm_database_connection_gen2_unit_test.go`)

Comprehensive unit tests for transformation functions:

1. **Connection Transformation Tests**
   - PostgreSQL connection transformation
   - MongoDB connection transformation
   - Redis connection transformation
   - MySQL connection transformation
   - RabbitMQ connections (AMQPS, MQTTS, STOMP SSL)
   - Elasticsearch connections (HTTPS, gRPC)
   - DataStax connections (Analytics, BI Connector, Ops Manager, EMP)

2. **CLI Transformation Tests**
   - PostgreSQL CLI transformation
   - MongoDB CLI transformation
   - Redis CLI transformation

3. **Data Type Handling Tests**
   - Query options type conversion (boolean, int, float to string)
   - Port number type handling (float64, int, int64)
   - Empty input handling
   - Missing optional fields handling

### Running Tests

#### Prerequisites

```bash
# Set required environment variables
export IC_API_KEY="your-ibm-cloud-api-key"
export IBMCLOUD_TIMEOUT=900
```

#### Run All Gen2 Connection Tests

```bash
# Acceptance tests
cd ibm/service/database
go test -v -run TestAccIBMDatabaseConnectionGen2

# Unit tests
go test -v -run TestTransformGen2
```

#### Run Specific Test

```bash
# Run a specific acceptance test
go test -v -run TestAccIBMDatabaseConnectionGen2DataSourceRead

# Run a specific unit test
go test -v -run TestTransformGen2ConnectionToSchema_PostgreSQL
```

#### Run with Verbose Output

```bash
# Enable debug logging
export TF_LOG=DEBUG
go test -v -run TestAccIBMDatabaseConnectionGen2DataSourceRead
```

### Test Results Interpretation

- **PASS**: Test completed successfully
- **FAIL**: Test failed - check error messages for details
- **SKIP**: Test was skipped (usually due to missing prerequisites)

---

## Troubleshooting

### Common Issues and Solutions

#### Issue 1: "No resource keys found for Gen2 database"

**Symptom**:
```
Error: No resource keys found for Gen2 database (deployment_id: crn:...).
Gen2 databases require resource keys to access connection information.
Please create a resource key using ibm_resource_key resource first.
```

**Cause**: The Gen2 database doesn't have any resource keys created.

**Solution**:
```hcl
# Create a resource key before querying connection
resource "ibm_resource_key" "db_key" {
  name                 = "database-credentials"
  resource_instance_id = ibm_database.my_gen2_db.id
}

data "ibm_database_connection" "connection" {
  deployment_id = ibm_database.my_gen2_db.id
  user_type     = "database"
  user_id       = ibm_resource_key.db_key.name
  endpoint_type = "private"
  
  depends_on = [ibm_resource_key.db_key]
}
```

#### Issue 2: "Resource key credentials not available"

**Symptom**:
```
Error: Resource key credentials not available
Resource key 'my-key' does not contain credentials.
This may be due to insufficient permissions or the key being in an invalid state.
```

**Cause**: The resource key exists but doesn't have credentials populated.

**Solutions**:

1. **Wait for key to be ready**:
```hcl
resource "ibm_resource_key" "db_key" {
  name                 = "database-credentials"
  resource_instance_id = ibm_database.my_gen2_db.id
}

# Add a delay to ensure key is fully provisioned
resource "time_sleep" "wait_for_key" {
  depends_on      = [ibm_resource_key.db_key]
  create_duration = "30s"
}

data "ibm_database_connection" "connection" {
  deployment_id = ibm_database.my_gen2_db.id
  user_type     = "database"
  user_id       = ibm_resource_key.db_key.name
  endpoint_type = "private"
  
  depends_on = [time_sleep.wait_for_key]
}
```

2. **Check IAM permissions**: Ensure you have the necessary permissions to create and access resource keys.

#### Issue 3: Wrong Key Used (Fallback Behavior)

**Symptom**: The connection uses a different resource key than expected.

**Cause**: The specified `user_id` (key name) doesn't exist, triggering fallback to the first available key.

**Solution**:

1. **Verify key name**:
```bash
# List resource keys for the instance
ibmcloud resource service-keys --instance-name my-database-gen2
```

2. **Use exact key name**:
```hcl
data "ibm_database_connection" "connection" {
  deployment_id = ibm_database.my_gen2_db.id
  user_type     = "database"
  user_id       = "exact-key-name"  # Use exact name from list
  endpoint_type = "private"
}
```

3. **Check state for actual key used**:
```bash
terraform state show data.ibm_database_connection.connection | grep user_id
```

#### Issue 4: "Instance is not a Gen2 database"

**Symptom**:
```
Error: Instance is not a Gen2 database
Instance crn:... is not a Gen2 database
```

**Cause**: Attempting to use Gen2 backend with a Classic database, or vice versa.

**Solution**: The data source should automatically detect the correct backend. If you see this error, verify:

1. **Check the plan**:
```bash
ibmcloud resource service-instance my-database --output json | jq '.resource_plan_id'
```

2. **Ensure plan contains "gen2"** for Gen2 databases or doesn't contain it for Classic.

#### Issue 5: Connection Type Not Available

**Symptom**: Expected connection type (e.g., `postgres`, `mongodb`) is empty or not set.

**Cause**: The connection type might not be available in the resource key credentials, or there's a mismatch in connection type names.

**Solutions**:

1. **Check available connection types**:
```bash
# Get resource key details
ibmcloud resource service-key my-key --output json | jq '.credentials.connection | keys'
```

2. **Verify database service type**: Ensure you're querying the correct connection type for your database service.

3. **Enable debug logging**:
```bash
export TF_LOG=DEBUG
terraform plan
# Check logs for "Available connection types in credentials.connection:"
```

#### Issue 6: Certificate Issues

**Symptom**: SSL/TLS connection failures when using the connection string.

**Cause**: Certificate not properly extracted or configured.

**Solution**:

1. **Extract certificate from connection data**:
```hcl
# Save certificate to file
resource "local_file" "db_cert" {
  content  = base64decode(
    data.ibm_database_connection.connection.postgres[0].certificate[0].certificate_base64
  )
  filename = "${path.module}/ca-certificate.crt"
}
```

2. **Use certificate in connection**:
```bash
psql "$(terraform output -raw postgres_connection_string)" \
  --set=sslrootcert=ca-certificate.crt
```

### Debug Mode

Enable detailed logging to troubleshoot issues:

```bash
# Enable Terraform debug logging
export TF_LOG=DEBUG
export TF_LOG_PATH=./terraform-debug.log

# Run Terraform command
terraform plan

# Review logs
grep "Gen2 database connection" terraform-debug.log
grep "Available connection types" terraform-debug.log
grep "Found.*connection data" terraform-debug.log
```

**Key Debug Messages to Look For**:

- `"Gen2 database connection information retrieved from resource key: <key-name>"`
- `"Available connection types in credentials.connection:"`
- `"Found <type> connection data under '<key>' key"`
- `"Successfully set <type> connection data"`
- `"No <type> connection data found in credentials"`

---

## Best Practices

### 1. Always Create Resource Keys for Gen2

**Why**: Gen2 databases require resource keys to access connection information.

**Good Practice**:
```hcl
resource "ibm_database" "gen2_db" {
  name = "my-gen2-database"
  plan = "standard-gen2"
  # ... configuration
}

resource "ibm_resource_key" "db_key" {
  name                 = "app-credentials"
  resource_instance_id = ibm_database.gen2_db.id
}

data "ibm_database_connection" "connection" {
  deployment_id = ibm_database.gen2_db.id
  user_id       = ibm_resource_key.db_key.name
  # ... other parameters
  
  depends_on = [ibm_resource_key.db_key]
}
```

**Avoid**:
```hcl
# Don't query Gen2 connection without creating a key first
data "ibm_database_connection" "connection" {
  deployment_id = ibm_database.gen2_db.id
  user_id       = "some-user"  # This will fail for Gen2
  # ...
}
```

### 2. Use Descriptive Resource Key Names

**Why**: Makes it easier to identify and manage credentials, especially with multiple keys.

**Good Practice**:
```hcl
resource "ibm_resource_key" "app_key" {
  name                 = "${var.environment}-app-credentials"
  resource_instance_id = ibm_database.gen2_db.id
}

resource "ibm_resource_key" "admin_key" {
  name                 = "${var.environment}-admin-credentials"
  resource_instance_id = ibm_database.gen2_db.id
}
```

### 3. Leverage Depends_on for Proper Ordering

**Why**: Ensures resource keys are fully created before querying connection information.

**Good Practice**:
```hcl
data "ibm_database_connection" "connection" {
  deployment_id = ibm_database.gen2_db.id
  user_id       = ibm_resource_key.db_key.name
  endpoint_type = "private"
  
  depends_on = [ibm_resource_key.db_key]
}
```

### 4. Handle Sensitive Connection Information Properly

**Why**: Connection strings contain credentials that should not be exposed.

**Good Practice**:
```hcl
output "connection_string" {
  value     = data.ibm_database_connection.connection.postgres[0].composed[0]
  sensitive = true
}

output "connection_host" {
  value = data.ibm_database_connection.connection.postgres[0].hosts[0].hostname
  # Not sensitive - just the hostname
}
```

### 5. Use Separate Keys for Different Purposes

**Why**: Improves security through credential separation and easier rotation.

**Good Practice**:
```hcl
# Application credentials
resource "ibm_resource_key" "app_key" {
  name                 = "application-credentials"
  resource_instance_id = ibm_database.gen2_db.id
}

# Admin/maintenance credentials
resource "ibm_resource_key" "admin_key" {
  name                 = "admin-credentials"
  resource_instance_id = ibm_database.gen2_db.id
}

# Read-only credentials (if supported)
resource "ibm_resource_key" "readonly_key" {
  name                 = "readonly-credentials"
  resource_instance_id = ibm_database.gen2_db.id
}
```

### 6. Extract and Store Certificates Properly

**Why**: Many databases require SSL/TLS certificates for secure connections.

**Good Practice**:
```hcl
# Extract certificate
locals {
  db_certificate = try(
    base64decode(
      data.ibm_database_connection.connection.postgres[0].certificate[0].certificate_base64
    ),
    ""
  )
}

# Save to file if needed
resource "local_file" "db_cert" {
  count    = local.db_certificate != "" ? 1 : 0
  content  = local.db_certificate
  filename = "${path.module}/certificates/db-ca.crt"
}
```

### 7. Use Locals for Connection Details

**Why**: Makes connection information reusable and easier to reference.

**Good Practice**:
```hcl
locals {
  db_connection = data.ibm_database_connection.connection.postgres[0]
  
  db_host     = local.db_connection.hosts[0].hostname
  db_port     = local.db_connection.hosts[0].port
  db_name     = local.db_connection.database
  db_username = local.db_connection.authentication[0].username
  db_password = local.db_connection.authentication[0].password
}

# Use in other resources
resource "kubernetes_secret" "db_credentials" {
  metadata {
    name = "database-credentials"
  }
  
  data = {
    host     = local.db_host
    port     = local.db_port
    database = local.db_name
    username = local.db_username
    password = local.db_password
  }
}
```

### 8. Implement Proper Error Handling

**Why**: Gracefully handle cases where connection information might not be available.

**Good Practice**:
```hcl
locals {
  has_postgres_connection = length(data.ibm_database_connection.connection.postgres) > 0
  
  connection_string = local.has_postgres_connection ? (
    data.ibm_database_connection.connection.postgres[0].composed[0]
  ) : ""
}

output "connection_available" {
  value = local.has_postgres_connection
}
```

### 9. Document Generation-Specific Behavior

**Why**: Helps team members understand the differences between Classic and Gen2.

**Good Practice**:
```hcl
# Gen2 PostgreSQL Database
# Note: Gen2 databases require resource keys for connection information.
# The user_id parameter should reference the resource key name, not a database username.
resource "ibm_database" "postgres_gen2" {
  name = "my-postgres-gen2"
  plan = "standard-gen2"
  # ...
}

resource "ibm_resource_key" "postgres_key" {
  name                 = "postgres-credentials"
  resource_instance_id = ibm_database.postgres_gen2.id
}

data "ibm_database_connection" "postgres_connection" {
  deployment_id = ibm_database.postgres_gen2.id
  user_id       = ibm_resource_key.postgres_key.name  # Resource key name for Gen2
  # ...
}
```

### 10. Test Both Public and Private Endpoints

**Why**: Ensures your application can connect through the appropriate network path.

**Good Practice**:
```hcl
# Private endpoint connection (recommended for production)
data "ibm_database_connection" "private_connection" {
  deployment_id = ibm_database.gen2_db.id
  user_id       = ibm_resource_key.db_key.name
  endpoint_type = "private"
  
  depends_on = [ibm_resource_key.db_key]
}

# Public endpoint connection (for development/testing)
data "ibm_database_connection" "public_connection" {
  deployment_id = ibm_database.gen2_db.id
  user_id       = ibm_resource_key.db_key.name
  endpoint_type = "public"
  
  depends_on = [ibm_resource_key.db_key]
}
```

---

## Additional Resources

### Official Documentation

- [IBM Cloud Databases Documentation](https://cloud.ibm.com/docs/databases-for-postgresql)
- [IBM Cloud Terraform Provider Documentation](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs)
- [Resource Keys Documentation](https://cloud.ibm.com/docs/account?topic=account-service_credentials)

### Related Resources

- `ibm_database` - Create and manage database instances
- `ibm_resource_key` - Create and manage resource keys for service credentials
- `ibm_database` data source - Query database instance information

### Community and Support

- [IBM Cloud Terraform Provider GitHub](https://github.com/IBM-Cloud/terraform-provider-ibm)
- [IBM Cloud Support](https://cloud.ibm.com/unifiedsupport/supportcenter)
- [Terraform IBM Cloud Provider Issues](https://github.com/IBM-Cloud/terraform-provider-ibm/issues)

### Learning Resources

- [Getting Started with IBM Cloud Databases](https://cloud.ibm.com/docs/databases-for-postgresql?topic=databases-for-postgresql-getting-started)
- [Terraform Best Practices](https://www.terraform.io/docs/cloud/guides/recommended-practices/index.html)
- [IBM Cloud Database Security](https://cloud.ibm.com/docs/databases-for-postgresql?topic=databases-for-postgresql-security-compliance)

---

## Changelog

### Version 1.71.0+ (Current)

**Added**:
- Gen2 support for `ibm_database_connection` data source
- Automatic backend selection based on database plan type
- Resource key-based connection retrieval for Gen2 databases
- Fallback mechanism to first available key when specified key not found
- Comprehensive transformation functions for all connection types
- Support for nested connection structure in Gen2 credentials
- Flexible key matching for connection types (e.g., "mongodb" or "mongo")

**Changed**:
- `user_id` parameter now accepts resource key names for Gen2 databases
- Connection retrieval method differs between Classic and Gen2
- Gen2 backend uses Resource Controller API instead of Cloud Databases API

**Testing**:
- Added acceptance tests for Gen2 connection scenarios
- Added unit tests for transformation functions
- Added tests for error handling and fallback behavior

---

## FAQ

### Q: How do I know if my database supports Gen2 connections?

**A**: Check the plan name of your database instance. Gen2 plans include "gen2" in the name:
- Gen2 plans: `standard-gen2`, `enterprise-gen2`
- Classic plans: `standard`, `enterprise`, `platinum`

You can check your database plan using:
```bash
ibmcloud resource service-instance my-database --output json | jq '.resource_plan_id'
```

### Q: Can I use the same data source for both Classic and Gen2?

**A**: Yes! The `ibm_database_connection` data source automatically detects whether your database is Classic or Gen2 and uses the appropriate backend. You don't need to specify which type you're using.

### Q: What happens if I don't create a resource key for Gen2?

**A**: The data source will return an error indicating that no resource keys were found. Gen2 databases require resource keys to access connection information. You must create at least one resource key using the `ibm_resource_key` resource.

### Q: Can I have multiple resource keys for one Gen2 database?

**A**: Yes! You can create multiple resource keys for different purposes (e.g., application credentials, admin credentials, read-only access). Each key provides independent credentials.

### Q: What happens if I specify a non-existent resource key name?

**A**: The data source will fall back to using the first available resource key and update the `user_id` attribute in state to reflect the actual key used. A debug message will be logged indicating the fallback occurred.

### Q: How do I rotate credentials for Gen2 databases?

**A**: Create a new resource key, update your application to use the new credentials, then delete the old resource key:

```hcl
# Create new key
resource "ibm_resource_key" "new_key" {
  name                 = "new-credentials"
  resource_instance_id = ibm_database.gen2_db.id
}

# Update connection to use new key
data "ibm_database_connection" "connection" {
  deployment_id = ibm_database.gen2_db.id
  user_id       = ibm_resource_key.new_key.name
  # ...
}

# After migration, delete old key
# resource "ibm_resource_key" "old_key" { ... }  # Remove or comment out
```

### Q: Are all connection types available for all database services?

**A**: No. Available connection types depend on the database service:
- PostgreSQL: `postgres`, `cli`
- MongoDB: `mongodb` (or `mongo`), `cli`
- Redis: `rediss`, `cli`
- MySQL: `mysql`, `cli`
- RabbitMQ: `amqps`, `mqtts`, `stomp_ssl`, `https`, `cli`
- Elasticsearch: `https`, `grpc`, `cli`
- DataStax: `analytics`, `bi_connector`, `ops_manager`, `emp`, `cli`

### Q: How do I handle the certificate for SSL/TLS connections?

**A**: Extract the certificate from the connection data and decode it:

```hcl
locals {
  certificate_base64 = data.ibm_database_connection.connection.postgres[0].certificate[0].certificate_base64
  certificate_decoded = base64decode(local.certificate_base64)
}

resource "local_file" "ca_cert" {
  content  = local.certificate_decoded
  filename = "${path.module}/ca-certificate.crt"
}
```

### Q: Can I use Gen2 connection data source with Classic databases?

**A**: Yes! The data source automatically detects the database type and uses the appropriate backend. Your existing Classic database configurations will continue to work without modification.

### Q: What's the difference between user_id for Classic vs Gen2?

**A**: 
- **Classic**: `user_id` is the database username (e.g., "admin", "user123")
- **Gen2**: `user_id` is the resource key name (e.g., "app-credentials", "admin-key")

The data source handles this difference automatically based on the detected database type.

### Q: How do I debug connection issues?

**A**: Enable debug logging to see detailed information about the connection retrieval process:

```bash
export TF_LOG=DEBUG
terraform plan 2>&1 | grep -A 10 "Gen2 database connection"
```

Look for messages about:
- Available connection types
- Which resource key is being used
- Successfully set connection data
- Any errors or warnings

---

## Support and Contribution

### Getting Help

If you encounter issues or have questions:

1. **Check this documentation** for common issues and solutions
2. **Enable debug logging** to gather detailed information
3. **Search existing issues** on [GitHub](https://github.com/IBM-Cloud/terraform-provider-ibm/issues)
4. **Open a new issue** with:
   - Terraform version
   - Provider version
   - Database service and plan
   - Sanitized configuration (remove sensitive data)
   - Error messages and debug logs

### Contributing

Contributions are welcome! To contribute:

1. **Fork the repository**
2. **Create a feature branch**
3. **Make your changes** with tests
4. **Run tests** to ensure everything works
5. **Submit a pull request** with:
   - Clear description of changes
   - Test coverage for new functionality
   - Documentation updates

### Code Style

Follow the existing code style:
- Use meaningful variable names
- Add comments for complex logic
- Include error handling
- Write comprehensive tests
- Update documentation

### Testing Guidelines

When adding new features:

1. **Write unit tests** for transformation functions
2. **Write acceptance tests** for end-to-end scenarios
3. **Test error cases** and edge conditions
4. **Verify backward compatibility** with existing configurations

---

**Document Version**: 1.0  
**Last Updated**: 2026-05-28  
**Provider Version**: 1.71.0+  
**Related PR**: [#6808](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/6808)

---

*This documentation is maintained by the IBM Cloud Terraform Provider team. For questions or feedback, please open an issue on GitHub.*