# Testing Log Data Retention Tags Example

This guide explains how to run and test the `log-data-retention-tags-example.tf` file.

## Prerequisites

1. **IBM Cloud Account** with access to Cloud Logs service
2. **Cloud Logs Instance** already created with an archive bucket configured
3. **IBM Cloud API Key** with appropriate permissions
4. **Terraform** installed (v1.0 or later)
5. **Local Provider Build** (for testing unreleased features)

## Setup Steps

### 1. Build the Provider Locally

```bash
# From the terraform-provider-ibm root directory
make build

# This installs the provider to ~/go/bin/terraform-provider-ibm
```

### 2. Configure Local Provider

Copy the built provider to the Terraform plugins directory:

```bash
# Create the directory structure
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/ibm-cloud/ibm/2.0.0/darwin_arm64

# Copy the provider binary
cp ~/go/bin/terraform-provider-ibm ~/.terraform.d/plugins/registry.terraform.io/ibm-cloud/ibm/2.0.0/darwin_arm64/terraform-provider-ibm_v2.0.0

# Make it executable
chmod +x ~/.terraform.d/plugins/registry.terraform.io/ibm-cloud/ibm/2.0.0/darwin_arm64/terraform-provider-ibm_v2.0.0
```

**Note:** Adjust the path for your OS:
- macOS ARM: `darwin_arm64`
- macOS Intel: `darwin_amd64`
- Linux: `linux_amd64`
- Windows: `windows_amd64`

### 3. Configure Terraform to Use Local Provider

Create or update `~/.terraformrc`:

```hcl
provider_installation {
  filesystem_mirror {
    path    = "/Users/YOUR_USERNAME/.terraform.d/plugins"
    include = ["registry.terraform.io/ibm-cloud/ibm"]
  }
  direct {
    exclude = ["registry.terraform.io/ibm-cloud/ibm"]
  }
}
```

### 4. Set Up Variables

Copy the template and fill in your values:

```bash
cd examples/ibm-logs
cp terraform.tfvars.template terraform.tfvars
```

Edit `terraform.tfvars` with your values:

```hcl
ibmcloud_api_key    = "your-api-key-here"
region              = "us-south"  # or your preferred region
logs_instance_id    = "your-instance-id"
logs_instance_region = "us-south"
```

**Finding Your Instance ID:**
```bash
# List your Cloud Logs instances
ibmcloud resource service-instances --service-name logs

# Get instance details
ibmcloud resource service-instance "your-instance-name" --output json | jq -r '.guid'
```

## Running the Example

### Option 1: Test Only Retention Tags (Recommended for Quick Testing)

Create a minimal test file:

```bash
cat > test-retention-tags.tf << 'EOF'
terraform {
  required_providers {
    ibm = {
      source  = "ibm-cloud/ibm"
      version = "~> 2.0.0"
    }
  }
}

provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = var.region
}

# Configure retention tags
resource "ibm_logs_log_data_retention_tags" "example" {
  instance_id = var.logs_instance_id
  region      = var.logs_instance_region
  tags        = ["short-term", "medium-term", "long-term"]
}

# Read back the tags
data "ibm_logs_log_data_retention_tags" "example" {
  instance_id = var.logs_instance_id
  region      = var.logs_instance_region
  depends_on  = [ibm_logs_log_data_retention_tags.example]
}

output "retention_tags" {
  value = ibm_logs_log_data_retention_tags.example
}

output "retention_tags_data" {
  value = data.ibm_logs_log_data_retention_tags.example
}
EOF
```

### Option 2: Run the Full Example

The full example includes policies and other resources:

```bash
# Initialize Terraform
terraform init

# Review the plan
terraform plan

# Apply the configuration
terraform apply
```

### Option 3: Use the Test Script

Run the automated test script:

```bash
chmod +x test-pr188-retention-tags.sh
./test-pr188-retention-tags.sh
```

This script runs comprehensive tests including:
- Creating retention tags
- Updating tags
- Reading tags via data source
- Testing backward compatibility
- Verifying tag values

## Verification Steps

### 1. Check Terraform State

```bash
# View the created resource
terraform show

# Check specific resource
terraform state show ibm_logs_log_data_retention_tags.example
```

Expected output:
```hcl
resource "ibm_logs_log_data_retention_tags" "example" {
    id          = "us-south/your-instance-id"
    instance_id = "your-instance-id"
    region      = "us-south"
    tags        = [
        "short-term",
        "medium-term",
        "long-term",
    ]
}
```

### 2. Verify via IBM Cloud UI

1. Go to [IBM Cloud Console](https://cloud.ibm.com)
2. Navigate to **Observability** → **Logging** → **Cloud Logs**
3. Select your instance
4. Go to **Data Pipeline** → **Data Usage**
5. Check the **Retention Tags** section

### 3. Verify via API

```bash
# Get IAM token
TOKEN=$(ibmcloud iam oauth-tokens --output json | jq -r '.iam_token')

# Get retention tags
curl -X GET \
  "https://api.${REGION}.logs.cloud.ibm.com/v1/log_data_retention_tags" \
  -H "Authorization: ${TOKEN}" \
  -H "Content-Type: application/json"
```

### 4. Verify via Data Source

```bash
# Check data source output
terraform output retention_tags_data
```

## Testing Updates

### Update Tag Names

Edit your terraform file:

```hcl
resource "ibm_logs_log_data_retention_tags" "example" {
  instance_id = var.logs_instance_id
  region      = var.logs_instance_region
  tags        = ["tier-1", "tier-2", "tier-3"]  # Changed names
}
```

Apply the changes:

```bash
terraform apply
```

### Test with Policy

Create a policy using one of the retention tags:

```hcl
resource "ibm_logs_policy" "example" {
  instance_id = var.logs_instance_id
  region      = var.logs_instance_region
  name        = "test-policy"
  description = "Policy using retention tag"
  priority    = "type_high"
  
  application_rule {
    name         = "test-app"
    rule_type_id = "is"
  }
  
  log_rules {
    severities = ["info", "warning", "error"]
  }
  
  archive_retention_tag = "short-term"  # Use one of your tags
  
  depends_on = [ibm_logs_log_data_retention_tags.example]
}
```

## Troubleshooting

### Issue: Provider Not Found

**Error:** `provider registry.terraform.io/ibm-cloud/ibm was not found`

**Solution:**
1. Verify the provider binary exists in the correct location
2. Check `~/.terraformrc` configuration
3. Run `terraform init` again

### Issue: Resource Not Supported

**Error:** `The provider does not support resource type "ibm_logs_log_data_retention_tags"`

**Solution:**
1. Ensure you built the provider from the correct branch
2. Verify the provider version: `terraform version`
3. Check that the local provider is being used (not downloaded from registry)

### Issue: 412 Precondition Failed

**Error:** `UpdateLogDataRetentionTagsWithContext failed: 412 Precondition Failed`

**Solution:**
This means the archive bucket is not configured. You must:
1. Configure an archive bucket for your Cloud Logs instance first
2. Wait a few minutes for the configuration to propagate
3. Then run terraform apply again

### Issue: Tags Not Showing in UI

**Possible Causes:**
1. **UI Cache:** Hard refresh the browser (Cmd+Shift+R or Ctrl+Shift+R)
2. **Propagation Delay:** Wait 1-2 minutes and refresh
3. **Wrong Instance:** Verify you're viewing the correct Cloud Logs instance

**Verification:**
```bash
# Verify tags via API
curl -X GET \
  "https://api.${REGION}.logs.cloud.ibm.com/v1/log_data_retention_tags" \
  -H "Authorization: ${TOKEN}"
```

If the API returns the tags correctly, it's a UI display issue.

## Cleanup

To remove the resources:

```bash
# Destroy all resources
terraform destroy

# Or remove specific resource
terraform destroy -target=ibm_logs_log_data_retention_tags.example
```

**Note:** The retention tags resource cannot be truly deleted - it only removes it from Terraform state. The tags remain configured in the Cloud Logs instance until you remove the archive bucket.

## Debug Mode

For detailed logging:

```bash
# Enable debug logging
export TF_LOG=DEBUG
export TF_LOG_PATH=./terraform-debug.log

# Run terraform
terraform apply

# View logs
cat terraform-debug.log | grep -i "retention"
```

## Additional Resources

- [Cloud Logs Documentation](https://cloud.ibm.com/docs/cloud-logs)
- [Terraform IBM Provider Documentation](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs)
- [Example Files](./log-data-retention-tags-example.tf)
