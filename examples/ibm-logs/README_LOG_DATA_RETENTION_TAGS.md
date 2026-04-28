# IBM Cloud Logs Log Data Retention Tags - Terraform Example

This guide explains how to test the new feature Log Data retention tags in the IBM Cloud Logs Terraform provider.

## Overview of Changes

1. **New Resource**: `ibm_logs_log_data_retention_tags` - Manages log data retention tags
2. **Updated Field**: `archive_retention_tag` in `ibm_logs_policy` - Replaces the old `archive_retention` object

## Prerequisites

Before testing these changes, ensure you have:

1. **IBM Cloud Logs Instance**: A provisioned IBM Cloud Logs instance
2. **Archive Bucket**: **CRITICAL** - Archive bucket must be configured and attached to your Logs instance BEFORE attempting to configure retention tags. Without an archive bucket, Terraform will fail with a 412 Precondition Failed error.
3. **IBM Cloud API Key**: Valid API key with permissions to manage Logs resources
4. **Terraform**: Version 1.0 or later

## Setup

### 1. Configure Variables

Create or update `terraform.tfvars` with your instance details:

```hcl
ibmcloud_api_key = "your-api-key-here"
logs_instance_id = "your-logs-instance-guid"
region           = "us-south"  # or your instance region
```

### 2. Initialize Terraform

```bash
cd examples/ibm-logs
terraform init
```

## Testing Scenarios

### Scenario 1: Test New Retention Tags Resource

This tests the new `ibm_logs_log_data_retention_tags` resource.

**File**: `pr188-retention-tags-example.tf`

**Steps**:

1. Review the retention tags configuration:
   ```bash
   terraform plan -target=ibm_logs_log_data_retention_tags.example
   ```

2. Apply the retention tags:
   ```bash
   terraform apply -target=ibm_logs_log_data_retention_tags.example
   ```

3. Verify the tags were created:
   ```bash
   terraform show
   ```

**Expected Result**: 
- 3 retention tags created with IDs 1, 2, 3
- Tag names: "short-term", "medium-term", "long-term"

### Scenario 2: Test Policy with Archive Retention Tag

This tests the updated `archive_retention_tag` field in policies.

**File**: `pr188-retention-tags-example.tf`

**Steps**:

1. Plan the policy creation:
   ```bash
   terraform plan -target=ibm_logs_policy.example_with_retention_tag
   ```

2. Apply the policy:
   ```bash
   terraform apply -target=ibm_logs_policy.example_with_retention_tag
   ```

3. Verify the policy uses the retention tag:
   ```bash
   terraform output policy_archive_retention_tag
   ```

**Expected Result**:
- Policy created successfully
- `archive_retention_tag` field set to "medium-term"
- Policy references the retention tag by name (not ID)

### Scenario 3: Test Data Sources

This tests reading retention tags and policies via data sources.

**Steps**:

1. Apply all resources:
   ```bash
   terraform apply
   ```

2. View the outputs:
   ```bash
   terraform output retention_tags_configuration
   terraform output policy_details
   ```

**Expected Result**:
- Data sources successfully read the resources
- Outputs display retention tags and policy details
- `archive_retention_tag` field visible in policy data source

### Scenario 4: Test Update Operations

This tests updating retention tag names and policy configurations.

**Steps**:

1. Modify retention tag names in `pr188-retention-tags-example.tf`:
   ```hcl
   retention_tags {
     tag_id   = 1
     tag_name = "weekly"  # changed from "short-term"
   }
   ```

2. Apply the changes:
   ```bash
   terraform apply
   ```

3. Verify the policy still references the correct tag:
   ```bash
   terraform output policy_archive_retention_tag
   ```

**Expected Result**:
- Retention tag name updated successfully
- Policy automatically uses the new tag name

### Scenario 5: Test Validation

This tests the validation rules for retention tags and policies.

**Test Cases**:

1. **Invalid Tag ID** (should fail):
   ```hcl
   retention_tags {
     tag_id   = 0  # Reserved, should fail
     tag_name = "test"
   }
   ```

2. **Invalid Tag Name Pattern** (should fail):
   ```hcl
   retention_tags {
     tag_id   = 1
     tag_name = "invalid tag!"  # Contains invalid characters
   }
   ```

3. **Non-existent Retention Tag** (should fail):
   ```hcl
   archive_retention_tag = "non-existent-tag"
   ```

**Expected Result**:
- Terraform validation errors for invalid configurations
- Clear error messages explaining the validation rules

## Cleanup

To remove all test resources:

```bash
terraform destroy
```

## Troubleshooting

### Error: Archive bucket not configured

**Symptom**: Error when creating retention tags
```
Error: Archive bucket must be configured before using retention tags
```

**Solution**: Configure an archive bucket in your IBM Cloud Logs instance before using retention tags.

### Error: Invalid retention tag name

**Symptom**: Validation error for tag name
```
Error: archive_retention_tag must match pattern: ^[a-zA-Z0-9_-]+$
```

**Solution**: Use only alphanumeric characters, hyphens, and underscores in tag names.

### Error: Policy references non-existent tag

**Symptom**: Error when creating policy
```
Error: The specified retention tag does not exist
```

**Solution**: Ensure the retention tag is created before referencing it in a policy. Use `depends_on` if needed.

## API Endpoints Tested

These examples test the following API endpoints from PR #188:

1. **GET /v1/log-data-retention-tags** - Read retention tags
2. **PUT /v1/log-data-retention-tags** - Update retention tags
3. **GET /v1/policies** - Read policies (with new `archive_retention_tag` field)
4. **POST /v1/policies** - Create policy (with new `archive_retention_tag` field)
5. **PUT /v1/policies/{id}** - Update policy (with new `archive_retention_tag` field)

## Additional Resources

- [OpenAPI PR #188](https://github.com/observability-c/dragonlog-openapi/pull/188)
- [Terraform Provider Documentation](../../website/docs/r/logs_log_data_retention_tags.html.markdown)
- [Manual Changes Runbook](https://github.com/observability-c/dragonlog-runbooks/blob/main/runbooks/terraform/ManualChangesOnProvider.md)

## Reporting Issues

If you encounter any issues while testing:

1. Check the Terraform logs: `TF_LOG=DEBUG terraform apply`
2. Verify your IBM Cloud Logs instance configuration
3. Ensure archive bucket is properly configured
4. Report issues with full error messages and steps to reproduce