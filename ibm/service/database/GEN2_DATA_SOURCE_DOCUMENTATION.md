# IBM Cloud Terraform Provider - Database Gen2 Data Source Support

## Overview

This documentation covers the enhancement to the IBM Cloud Terraform Provider that adds support for **Generation 2 (Gen2) database instances** in the `ibm_database` data source. This update enables users to seamlessly query both Classic and Gen2 database deployments using the same data source with automatic detection and appropriate handling of generation-specific features.

---

## Table of Contents

- [What's New](#whats-new)
- [Understanding Classic vs Gen2 Databases](#understanding-classic-vs-gen2-databases)
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

### Gen2 Database Support in Data Source

The `ibm_database` data source now **automatically detects** whether a database instance is Classic or Gen2 and retrieves the appropriate information. This provides a unified interface for querying database instances regardless of their generation.

### Key Features

✅ **Automatic Detection**: The provider automatically identifies whether a database is Classic or Gen2 based on the plan type  
✅ **Unified Interface**: Use the same `ibm_database` data source for both Classic and Gen2 instances  
✅ **Backward Compatible**: Existing Terraform configurations continue to work without modification  
✅ **Proper State Management**: Gen2-specific behavior ensures clean state management when switching between instance types  
✅ **Comprehensive Testing**: Full test coverage for Gen2 scenarios including edge cases  

---

## Understanding Classic vs Gen2 Databases

### Classic Databases (Standard Plans)

- **Deployment Model**: Traditional IBM Cloud Databases architecture
- **Plan Names**: `standard`, `enterprise`, `platinum`
- **Features**: Full feature set including:
  - Default admin user management
  - IP allowlists for network access control
  - Auto-scaling configuration
  - Direct user management through API
  - Configuration schema access

### Gen2 Databases (Standard-Gen2 Plans)

- **Deployment Model**: Next-generation architecture with enhanced capabilities
- **Plan Names**: `standard-gen2`, `enterprise-gen2`
- **Benefits**:
  - Enhanced performance and scalability
  - Improved resource isolation
  - Modern infrastructure
  - Better integration with IBM Cloud platform services
- **Management Model**: Different approach to credentials and access control

---

## Architecture Overview

### Backend Selection Pattern

The data source uses a **strategy pattern** to handle different database generations:

```
User Query → findInstance() → Detect Plan Type → Select Backend
                                                      ↓
                                    ┌─────────────────┴─────────────────┐
                                    ↓                                   ↓
                          Classic Backend                        Gen2 Backend
                          (Full API Access)                (Metadata-based Access)
```

### Implementation Components

1. **`data_source_ibm_database.go`**: Main data source with backend selection logic
2. **`data_source_ibm_database_classic.go`**: Classic database backend implementation
3. **`data_source_ibm_database_gen2.go`**: Gen2 database backend implementation
4. **`helpers.go`**: Shared utility functions for both backends
5. **Test Files**: Comprehensive test coverage for both generations

---

## Important Differences for Gen2

### Attributes Not Available in Gen2

The following attributes are **not supported** for Gen2 database instances and will return empty/zero values:

| Attribute | Classic Behavior | Gen2 Behavior | Reason |
|-----------|------------------|---------------|--------|
| `adminuser` | Returns default admin username | Returns empty string | No default admin user in Gen2 |
| `adminpassword` | Returns admin password | Returns empty string | Credentials managed via service keys |
| `users` | Returns list of database users | Returns empty list `[]` | User management through service keys |
| `allowlist` | Returns IP allowlist entries | Returns empty list `[]` | Network access via VPC/CBR |
| `auto_scaling` | Returns auto-scaling config | Returns empty list `[]` | Not currently supported |
| `configuration_schema` | Returns JSON schema | Returns empty string | Not available in Gen2 |
| `backup_encryption_key_crn` | Returns backup encryption key | Not set | Different encryption model |

### Attributes Available in Gen2

These attributes work identically in both Classic and Gen2:

✅ `name` - Database instance name  
✅ `service` - Database service type (e.g., `databases-for-postgresql`)  
✅ `plan` - Service plan (e.g., `standard-gen2`)  
✅ `status` - Instance status  
✅ `version` - Database version  
✅ `location` - Region/location  
✅ `resource_group_id` - Resource group identifier  
✅ `guid` - Unique instance identifier  
✅ `groups` - Scaling group configurations (memory, disk, CPU)  
✅ `platform_options.disk_encryption_key_crn` - Disk encryption key  
✅ `tags` - Instance tags  

---

## Usage Examples

### Example 1: Basic Data Source Query

Works for both Classic and Gen2 databases:

```hcl
# Query any database instance
data "ibm_database" "database" {
  name = "my-database-instance"
}

# Access common attributes
output "database_info" {
  value = {
    version  = data.ibm_database.database.version
    location = data.ibm_database.database.location
    plan     = data.ibm_database.database.plan
    status   = data.ibm_database.database.status
  }
}
```

### Example 2: Gen2 Database with Credential Management

For Gen2 databases, use `ibm_resource_key` for credentials:

```hcl
# Query Gen2 database instance
data "ibm_database" "my_gen2_db" {
  name = "my-postgres-gen2"
}

# Create service credentials for Gen2 database
resource "ibm_resource_key" "db_credentials" {
  name                 = "my-db-credentials"
  resource_instance_id = data.ibm_database.my_gen2_db.id
  role                 = "Administrator"
}

# Access connection details
output "database_connection" {
  value = {
    host     = ibm_resource_key.db_credentials.credentials["connection.postgres.hosts.0.hostname"]
    port     = ibm_resource_key.db_credentials.credentials["connection.postgres.hosts.0.port"]
    database = ibm_resource_key.db_credentials.credentials["connection.postgres.database"]
    username = ibm_resource_key.db_credentials.credentials["connection.postgres.authentication.username"]
    password = ibm_resource_key.db_credentials.credentials["connection.postgres.authentication.password"]
  }
  sensitive = true
}
```

### Example 3: Querying with Resource Group Filter

```hcl
data "ibm_resource_group" "group" {
  name = "production"
}

data "ibm_database" "database" {
  name              = "my-database"
  resource_group_id = data.ibm_resource_group.group.id
}
```

### Example 4: Querying with Location Filter

```hcl
data "ibm_database" "database" {
  name     = "my-database"
  location = "us-south"
}

# Useful when multiple databases share the same name in different regions
```

### Example 5: Accessing Scaling Groups Information

```hcl
data "ibm_database" "database" {
  name = "my-database"
}

# Memory allocation
output "memory_allocation_mb" {
  value = data.ibm_database.database.groups[0].memory[0].allocation_mb
}

# Disk allocation
output "disk_allocation_mb" {
  value = data.ibm_database.database.groups[0].disk[0].allocation_mb
}

# CPU allocation (if available)
output "cpu_allocation_count" {
  value = data.ibm_database.database.groups[0].cpu[0].allocation_count
}

# Check if resources are adjustable
output "can_scale_memory" {
  value = data.ibm_database.database.groups[0].memory[0].is_adjustable
}
```

### Example 6: Conditional Logic Based on Plan Type

```hcl
data "ibm_database" "database" {
  name = "my-database"
}

locals {
  is_gen2 = can(regex("gen2", data.ibm_database.database.plan))
}

# Use different credential management based on generation
resource "ibm_resource_key" "credentials" {
  count = local.is_gen2 ? 1 : 0
  
  name                 = "gen2-credentials"
  resource_instance_id = data.ibm_database.database.id
  role                 = "Administrator"
}

output "credential_source" {
  value = local.is_gen2 ? "Using ibm_resource_key for Gen2" : "Using adminuser for Classic"
}
```

### Example 7: Platform Options and Encryption

```hcl
data "ibm_database" "database" {
  name = "my-encrypted-database"
}

# Access disk encryption key
output "disk_encryption_key" {
  value = try(
    data.ibm_database.database.platform_options[0].disk_encryption_key_crn,
    "No encryption key configured"
  )
}
```

---

## Migration Guide

### Upgrading from Classic to Gen2

If you're migrating from a Classic database to Gen2, follow these steps:

#### Step 1: Update Plan Reference

```hcl
# Before (Classic)
data "ibm_database" "db" {
  name = "my-database"
}
# Plan will be "standard"

# After (Gen2)
data "ibm_database" "db" {
  name = "my-database-gen2"
}
# Plan will be "standard-gen2"
```

#### Step 2: Remove Classic-Only Attribute References

```hcl
# Before (Classic) - REMOVE THESE
output "admin_user" {
  value = data.ibm_database.db.adminuser
}

output "admin_password" {
  value     = data.ibm_database.db.adminpassword
  sensitive = true
}

# After (Gen2) - Use resource keys instead
resource "ibm_resource_key" "credentials" {
  name                 = "db-credentials"
  resource_instance_id = data.ibm_database.db.id
  role                 = "Administrator"
}

output "credentials" {
  value     = ibm_resource_key.credentials.credentials
  sensitive = true
}
```

#### Step 3: Update Network Access Control

```hcl
# Before (Classic) - Allowlist
# This will return empty for Gen2
output "allowlist" {
  value = data.ibm_database.db.allowlist
}

# After (Gen2) - Use VPC Security Groups or Context-Based Restrictions
# Network access is managed at the VPC/network level
```

#### Step 4: Handle Auto-Scaling Differences

```hcl
# Before (Classic)
output "autoscaling_config" {
  value = data.ibm_database.db.auto_scaling
}

# After (Gen2)
# Auto-scaling not currently available in Gen2
# Monitor and manually adjust resources as needed
```

### State Management During Migration

The data source automatically handles state transitions:

1. **Automatic Cleanup**: When switching from Classic to Gen2 instance queries, unsupported attributes are automatically cleared
2. **No Manual Intervention**: No manual state manipulation required
3. **Plan Visibility**: Terraform will show the attribute changes in the plan output

```bash
# After switching to Gen2, run:
terraform plan

# You'll see changes like:
# ~ adminuser = "admin" -> ""
# ~ users     = [...] -> []
# ~ allowlist = [...] -> []
```

---

## Supported Attributes

### Input Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `name` | String | Yes | Database instance name |
| `resource_group_id` | String | No | Filter by resource group ID |
| `location` | String | No | Filter by region/location |
| `service` | String | No | Filter by service type |

### Output Attributes - Common (Both Generations)

| Attribute | Type | Description |
|-----------|------|-------------|
| `id` | String | Resource instance ID |
| `name` | String | Database instance name |
| `service` | String | Database service type |
| `plan` | String | Service plan (includes `-gen2` suffix for Gen2) |
| `status` | String | Instance status (active, provisioning, etc.) |
| `version` | String | Database version |
| `location` | String | Region/location |
| `resource_group_id` | String | Resource group ID |
| `guid` | String | Unique instance identifier |
| `resource_name` | String | Resource name |
| `resource_crn` | String | Cloud Resource Name |
| `resource_status` | String | Resource status |
| `resource_group_name` | String | Resource group name |
| `resource_controller_url` | String | Dashboard URL |
| `tags` | Set(String) | Instance tags |

### Output Attributes - Groups (Both Generations)

| Attribute | Type | Description |
|-----------|------|-------------|
| `groups` | List | Scaling group configurations |
| `groups[].group_id` | String | Group identifier (e.g., "member") |
| `groups[].count` | Integer | Number of members in group |
| `groups[].memory` | List | Memory configuration |
| `groups[].memory[].allocation_mb` | Integer | Current memory allocation in MB |
| `groups[].memory[].minimum_mb` | Integer | Minimum allowed memory |
| `groups[].memory[].step_size_mb` | Integer | Memory scaling increment |
| `groups[].memory[].is_adjustable` | Boolean | Whether memory can be adjusted |
| `groups[].memory[].can_scale_down` | Boolean | Whether memory can scale down |
| `groups[].disk` | List | Disk configuration |
| `groups[].disk[].allocation_mb` | Integer | Current disk allocation in MB |
| `groups[].disk[].minimum_mb` | Integer | Minimum allowed disk |
| `groups[].disk[].step_size_mb` | Integer | Disk scaling increment |
| `groups[].disk[].is_adjustable` | Boolean | Whether disk can be adjusted |
| `groups[].disk[].can_scale_down` | Boolean | Whether disk can scale down |
| `groups[].cpu` | List | CPU configuration |
| `groups[].cpu[].allocation_count` | Integer | Current CPU allocation |
| `groups[].cpu[].minimum_count` | Integer | Minimum allowed CPUs |
| `groups[].cpu[].step_size_count` | Integer | CPU scaling increment |
| `groups[].cpu[].is_adjustable` | Boolean | Whether CPU can be adjusted |
| `groups[].cpu[].can_scale_down` | Boolean | Whether CPU can scale down |

### Output Attributes - Platform Options (Both Generations)

| Attribute | Type | Gen2 Support | Description |
|-----------|------|--------------|-------------|
| `platform_options` | Set | Yes | Platform-specific options |
| `platform_options[].disk_encryption_key_crn` | String | Yes | Disk encryption key CRN |
| `platform_options[].backup_encryption_key_crn` | String | No | Backup encryption key (Classic only) |

### Output Attributes - Classic Only

| Attribute | Type | Gen2 Behavior | Description |
|-----------|------|---------------|-------------|
| `adminuser` | String | Empty string | Default admin username (Classic only) |
| `adminpassword` | String | Empty string | Default admin password (Classic only) |
| `users` | Set | Empty list | Database users (Classic only) |
| `allowlist` | Set | Empty list | IP allowlist (Classic only) |
| `auto_scaling` | List | Empty list | Auto-scaling configuration (Classic only) |
| `configuration_schema` | String | Empty string | Configuration schema JSON (Classic only) |

---

## Testing

### Test Coverage

The implementation includes comprehensive test coverage for Gen2 scenarios:

#### Test Cases

1. **TestAccIBMDatabaseDataSourceGen2Basic**
   - Validates successful data retrieval for Gen2 instances
   - Verifies common attributes (name, service, plan, location)
   - Confirms Gen2-specific behavior (empty adminuser, users, allowlist)
   - Checks groups information (memory, disk allocations)

2. **TestAccIBMDatabaseDataSourceGen2WithResourceGroupID**
   - Tests filtering by resource group
   - Validates resource group ID is properly set
   - Ensures correct instance is retrieved when multiple exist

3. **TestAccIBMDatabaseDataSourceGen2InvalidInput**
   - Tests error handling for non-existent databases
   - Validates appropriate error messages
   - Ensures graceful failure

4. **TestAccIBMDatabaseDataSourceGen2InvalidID**
   - Tests error handling for malformed identifiers
   - Validates input validation
   - Ensures proper error reporting

### Running Tests

#### Prerequisites

```bash
# Set required environment variables
export IC_API_KEY="your-ibm-cloud-api-key"
export IAAS_CLASSIC_API_KEY="your-classic-infrastructure-key"
export IAAS_CLASSIC_USERNAME="your-classic-username"

# Optional: Set specific test database
export ICD_DB_GEN2_DEPLOYMENT_ID="your-gen2-database-name"
```

#### Run All Gen2 Data Source Tests

```bash
make testacc TEST=./ibm/service/database TESTARGS='-run=TestAccIBMDatabaseDataSourceGen2'
```

#### Run Specific Test

```bash
make testacc TEST=./ibm/service/database TESTARGS='-run=TestAccIBMDatabaseDataSourceGen2Basic'
```

#### Run with Verbose Output

```bash
TF_LOG=DEBUG make testacc TEST=./ibm/service/database TESTARGS='-run=TestAccIBMDatabaseDataSourceGen2'
```

### Test Results Interpretation

✅ **PASS**: All attributes correctly populated, Gen2-specific attributes empty  
❌ **FAIL**: Check error messages for specific issues  
⚠️ **SKIP**: Test skipped (usually due to missing prerequisites)  

---

## Troubleshooting

### Common Issues and Solutions

#### Issue 1: "No resource instance found"

**Symptoms:**
```
Error: No resource instance found with name [my-database]
```

**Possible Causes:**
- Database name doesn't exist
- Database is in a different resource group
- Database is in a different region
- Typo in database name

**Solutions:**
```hcl
# Solution 1: Verify database exists
# Check IBM Cloud console or CLI:
ibmcloud resource service-instances --service-name databases-for-postgresql

# Solution 2: Specify resource group explicitly
data "ibm_database" "db" {
  name              = "my-database"
  resource_group_id = "your-resource-group-id"
}

# Solution 3: Specify location
data "ibm_database" "db" {
  name     = "my-database"
  location = "us-south"
}

# Solution 4: Add service filter
data "ibm_database" "db" {
  name    = "my-database"
  service = "databases-for-postgresql"
}
```

#### Issue 2: "The database instance was not found in the region"

**Symptoms:**
```
Error: The database instance was not found in the region set for the Provider
```

**Cause:** Database exists in a different region than configured in provider

**Solution:**
```hcl
# Option 1: Update provider region
provider "ibm" {
  region = "eu-gb"  # Match your database region
}

# Option 2: Use provider alias for multi-region
provider "ibm" {
  alias  = "us_south"
  region = "us-south"
}

provider "ibm" {
  alias  = "eu_gb"
  region = "eu-gb"
}

data "ibm_database" "us_db" {
  provider = ibm.us_south
  name     = "my-us-database"
}

data "ibm_database" "eu_db" {
  provider = ibm.eu_gb
  name     = "my-eu-database"
}
```

#### Issue 3: Stale Classic Attributes in State

**Symptoms:**
- Gen2 database showing non-empty `adminuser` or `allowlist`
- State shows Classic attributes for Gen2 instance

**Cause:** Data source previously resolved to Classic, now resolves to Gen2

**Solution:**
```bash
# Refresh state to clear stale attributes
terraform refresh

# Or force re-read
terraform apply -refresh-only
```

#### Issue 4: Multiple Instances Found

**Symptoms:**
```
Error: More than one resource instance found with name matching [my-database]
```

**Cause:** Multiple databases with same name in different resource groups or regions

**Solution:**
```hcl
# Be more specific with filters
data "ibm_database" "db" {
  name              = "my-database"
  resource_group_id = data.ibm_resource_group.prod.id
  location          = "us-south"
  service           = "databases-for-postgresql"
}
```

#### Issue 5: Empty Groups Information

**Symptoms:**
- `groups` attribute is empty or missing data

**Cause:** Database is still provisioning or metadata not yet available

**Solution:**
```hcl
# Add depends_on to ensure database is ready
resource "ibm_database" "db" {
  # ... configuration
}

data "ibm_database" "db_info" {
  name = ibm_database.db.name
  
  depends_on = [ibm_database.db]
}

# Or add explicit wait
resource "time_sleep" "wait_for_db" {
  depends_on = [ibm_database.db]
  create_duration = "5m"
}

data "ibm_database" "db_info" {
  name = ibm_database.db.name
  
  depends_on = [time_sleep.wait_for_db]
}
```

### Debug Mode

Enable debug logging for detailed troubleshooting:

```bash
# Set Terraform log level
export TF_LOG=DEBUG
export TF_LOG_PATH=./terraform-debug.log

# Run Terraform command
terraform plan

# Review log file
cat terraform-debug.log | grep -i database
```

---

## Best Practices

### 1. Always Specify Resource Group

**Why:** Avoids ambiguity when multiple databases share the same name

```hcl
# ✅ Good
data "ibm_database" "db" {
  name              = "my-database"
  resource_group_id = data.ibm_resource_group.prod.id
}

# ❌ Avoid
data "ibm_database" "db" {
  name = "my-database"
  # May match wrong instance if multiple exist
}
```

### 2. Use Location Filter for Multi-Region Deployments

**Why:** Ensures you're querying the correct regional instance

```hcl
# ✅ Good
data "ibm_database" "db" {
  name     = "my-database"
  location = "us-south"
}
```

### 3. Check Plan Type to Determine Generation

**Why:** Enables conditional logic based on database generation

```hcl
# ✅ Good
data "ibm_database" "db" {
  name = var.database_name
}

locals {
  is_gen2 = can(regex("-gen2$", data.ibm_database.db.plan))
}

# Use different credential strategies
resource "ibm_resource_key" "gen2_creds" {
  count = local.is_gen2 ? 1 : 0
  # ... Gen2 credential configuration
}
```

### 4. Manage Gen2 Credentials Separately

**Why:** Gen2 doesn't provide default admin credentials

```hcl
# ✅ Good - Gen2 approach
data "ibm_database" "gen2_db" {
  name = "my-gen2-database"
}

resource "ibm_resource_key" "credentials" {
  name                 = "app-credentials"
  resource_instance_id = data.ibm_database.gen2_db.id
  role                 = "Administrator"
  
  parameters = {
    role_crn = var.service_id_crn  # Optional: bind to service ID
  }
}

# ❌ Avoid - Won't work for Gen2
output "admin_user" {
  value = data.ibm_database.gen2_db.adminuser  # Will be empty!
}
```

### 5. Use Outputs for Reusable Information

**Why:** Makes database information available to other modules

```hcl
# ✅ Good
data "ibm_database" "db" {
  name = var.database_name
}

output "database_info" {
  value = {
    id       = data.ibm_database.db.id
    version  = data.ibm_database.db.version
    location = data.ibm_database.db.location
    plan     = data.ibm_database.db.plan
    is_gen2  = can(regex("-gen2$", data.ibm_database.db.plan))
  }
}
```

### 6. Handle Optional Attributes Safely

**Why:** Prevents errors when attributes may not be set

```hcl
# ✅ Good - Use try() for optional attributes
output "disk_encryption" {
  value = try(
    data.ibm_database.db.platform_options[0].disk_encryption_key_crn,
    "No encryption key configured"
  )
}

# ✅ Good - Check before accessing
output "memory_allocation" {
  value = length(data.ibm_database.db.groups) > 0 ? (
    length(data.ibm_database.db.groups[0].memory) > 0 ? 
      data.ibm_database.db.groups[0].memory[0].allocation_mb : 
      null
  ) : null
}
```

### 7. Document Generation-Specific Behavior

**Why:** Helps team members understand limitations

```hcl
# ✅ Good - Add comments
data "ibm_database" "db" {
  name = var.database_name
}

# Note: For Gen2 databases, use ibm_resource_key for credentials
# adminuser and adminpassword will be empty for Gen2 instances
resource "ibm_resource_key" "credentials" {
  count = can(regex("-gen2$", data.ibm_database.db.plan)) ? 1 : 0
  
  name                 = "app-credentials"
  resource_instance_id = data.ibm_database.db.id
  role                 = "Administrator"
}
```

### 8. Monitor and Alert on Status

**Why:** Ensures database health and availability

```hcl
# ✅ Good - Check status
data "ibm_database" "db" {
  name = var.database_name
}

resource "null_resource" "status_check" {
  triggers = {
    status = data.ibm_database.db.status
  }
  
  provisioner "local-exec" {
    command = <<-EOT
      if [ "${data.ibm_database.db.status}" != "active" ]; then
        echo "Warning: Database status is ${data.ibm_database.db.status}"
      fi
    EOT
  }
}
```

### 9. Use Data Source for Read-Only Operations

**Why:** Data sources are for querying existing resources, not managing them

```hcl
# ✅ Good - Query existing database
data "ibm_database" "existing_db" {
  name = "production-database"
}

# Use the data for configuration
resource "ibm_resource_key" "app_key" {
  resource_instance_id = data.ibm_database.existing_db.id
  # ...
}

# ❌ Avoid - Don't try to modify via data source
# Data sources are read-only
```

### 10. Version Control Your Configurations

**Why:** Track changes and enable rollback

```hcl
# ✅ Good - Include version constraints
terraform {
  required_version = ">= 1.0"
  
  required_providers {
    ibm = {
      source  = "IBM-Cloud/ibm"
      version = "~> 1.60"  # Use version with Gen2 support
    }
  }
}
```

---

## Additional Resources

### Official Documentation

- **IBM Cloud Databases**: [https://cloud.ibm.com/docs/databases-for-postgresql](https://cloud.ibm.com/docs/databases-for-postgresql)
- **Terraform IBM Provider**: [https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs)
- **IBM Cloud Databases API**: [https://cloud.ibm.com/apidocs/cloud-databases-api](https://cloud.ibm.com/apidocs/cloud-databases-api)

### Related Resources

- **ibm_database Resource**: [https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/database](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/database)
- **ibm_resource_key Resource**: [https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/resource_key](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/resource_key)
- **ibm_resource_instance Data Source**: [https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/resource_instance](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/resource_instance)

### Community and Support

- **GitHub Repository**: [https://github.com/IBM-Cloud/terraform-provider-ibm](https://github.com/IBM-Cloud/terraform-provider-ibm)
- **GitHub Issues**: [https://github.com/IBM-Cloud/terraform-provider-ibm/issues](https://github.com/IBM-Cloud/terraform-provider-ibm/issues)
- **IBM Cloud Support**: [https://cloud.ibm.com/unifiedsupport/supportcenter](https://cloud.ibm.com/unifiedsupport/supportcenter)
- **Terraform Community**: [https://discuss.hashicorp.com/c/terraform-providers](https://discuss.hashicorp.com/c/terraform-providers)

### Learning Resources

- **Terraform Best Practices**: [https://www.terraform.io/docs/cloud/guides/recommended-practices/index.html](https://www.terraform.io/docs/cloud/guides/recommended-practices/index.html)
- **IBM Cloud Architecture Center**: [https://www.ibm.com/cloud/architecture](https://www.ibm.com/cloud/architecture)
- **Database Security Best Practices**: [https://cloud.ibm.com/docs/databases-for-postgresql?topic=databases-for-postgresql-security-compliance](https://cloud.ibm.com/docs/databases-for-postgresql?topic=databases-for-postgresql-security-compliance)

---

## Changelog

### Version 1.60.0+ (Current)

**Added:**
- ✅ Gen2 database support in `ibm_database` data source
- ✅ Automatic backend selection based on plan type
- ✅ Comprehensive test coverage for Gen2 scenarios
- ✅ Proper state management for generation transitions

**Changed:**
- 📝 Updated documentation to reflect Gen2 differences
- 📝 Added migration guide for Classic to Gen2 transition

**Deprecated:**
- ⚠️ Classic-only attributes return empty values for Gen2 instances

---

## FAQ

### Q: How do I know if my database is Gen2?

**A:** Check the `plan` attribute. Gen2 plans have a `-gen2` suffix (e.g., `standard-gen2`, `enterprise-gen2`).

```hcl
data "ibm_database" "db" {
  name = "my-database"
}

output "is_gen2" {
  value = can(regex("-gen2$", data.ibm_database.db.plan))
}
```

### Q: Can I use the same data source for both Classic and Gen2?

**A:** Yes! The data source automatically detects the generation and handles it appropriately.

### Q: What happens to Classic-only attributes when querying Gen2?

**A:** They return empty values (empty strings, empty lists) to maintain schema compatibility.

### Q: How do I get credentials for a Gen2 database?

**A:** Use the `ibm_resource_key` resource to create service credentials:

```hcl
resource "ibm_resource_key" "credentials" {
  name                 = "my-credentials"
  resource_instance_id = data.ibm_database.db.id
  role                 = "Administrator"
}
```

### Q: Will my existing Terraform configurations break?

**A:** No. The enhancement is fully backward compatible. Existing configurations continue to work without modification.

### Q: Can I migrate from Classic to Gen2 without downtime?

**A:** Database migration requires creating a new Gen2 instance and migrating data. Consult IBM Cloud documentation for migration strategies.

### Q: Are all database services available in Gen2?

**A:** Gen2 availability varies by service. Check IBM Cloud documentation for current Gen2 service availability.

### Q: How do I handle network access control for Gen2?

**A:** Gen2 uses VPC security groups and context-based restrictions instead of IP allowlists. Configure these at the VPC/network level.

---

## Support and Contribution

### Getting Help

If you encounter issues or have questions:

1. **Check Documentation**: Review this guide and official documentation
2. **Search Issues**: Look for similar issues on GitHub
3. **Create Issue**: Open a new issue with detailed information
4. **Contact Support**: Reach out to IBM Cloud Support for account-specific issues

### Contributing

Contributions are welcome! To contribute:

1. Fork the repository
2. Create a feature branch
3. Make your changes with tests
4. Submit a pull request

See [CONTRIBUTING.md](../../../CONTRIBUTING.md) for detailed guidelines.

---

**Document Version**: 1.0  
**Last Updated**: 2026-05-28  
**Terraform Provider Version**: 1.60.0+  

---

**Note**: This enhancement maintains full backward compatibility with existing Terraform configurations while providing seamless support for next-generation database deployments. Gen2 databases represent the future of IBM Cloud Databases with improved performance, scalability, and integration with IBM Cloud platform services.