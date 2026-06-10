# IBM Cloud Logs Extensions - Terraform Example

This example demonstrates how to use the IBM Cloud Logs Extensions APIs with Terraform.

## Important Note

The main `main.tf` file in this directory has syntax errors in the existing examples. To test the extensions functionality, use only the `extensions-example.tf` file.

## Quick Start

### 1. Rename main.tf temporarily
```bash
cd examples/ibm-logs
mv main.tf main.tf.backup
```

### 2. Configure credentials
```bash
cp terraform.tfvars.template terraform.tfvars
# Edit terraform.tfvars with your values:
# - ibmcloud_api_key
# - logs_instance_id  
# - region
```

### 3. Test extensions
```bash
terraform init
terraform plan
terraform apply
```

### 4. View outputs
```bash
terraform output all_extensions
terraform output deployment_details
```

### 5. Clean up
```bash
terraform destroy
mv main.tf.backup main.tf
```

## What Gets Tested

The `extensions-example.tf` file tests all 5 extension-related APIs:

| API | Method | Resource/Data Source |
|-----|--------|---------------------|
| List Extensions | GET | `data.ibm_logs_extensions` |
| Get Extension | GET | `data.ibm_logs_extension` |
| Deploy Extension | PUT | `resource.ibm_logs_extension_deployment` |
| Read Deployment | GET | `data.ibm_logs_extension_deployment` |
| Delete Deployment | DELETE | `resource.ibm_logs_extension_deployment` (destroy) |

## Files

- `extensions-example.tf` - Complete working example for extensions
- `variables.tf` - Variable definitions (already exists)
- `versions.tf` - Terraform version requirements (already exists)
- `terraform.tfvars.template` - Template for your credentials
- `QUICK_START.md` - Quick reference guide
- `test-all-apis.sh` - Automated test script

## Extension Used

The example uses the **IBMCloudant** extension which is available in all Cloud Logs instances.

## Troubleshooting

### "Duplicate provider configuration"
- Make sure you renamed `main.tf` to `main.tf.backup`
- Only `extensions-example.tf` should be active

### "Extension not found"
- Run `terraform output all_extensions` to see available extensions
- Verify the extension ID spelling (case-sensitive)

### "Authentication failed"
- Verify your API key in `terraform.tfvars`
- Ensure the API key has Cloud Logs permissions

## For More Details

See `QUICK_START.md` for step-by-step testing instructions.