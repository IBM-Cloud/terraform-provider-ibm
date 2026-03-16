# Quick Start Guide - Testing IBM Cloud Logs Extensions APIs

## Prerequisites Checklist

- [ ] IBM Cloud Logs instance created
- [ ] IBM Cloud API key with appropriate permissions
- [ ] Terraform provider built locally (`make build`)
- [ ] Provider binary in `$GOPATH/bin` (usually `~/go/bin`)

## Setup (5 minutes)

### 1. Configure Local Provider

Create or update `~/.terraformrc`:

```hcl
provider_installation {
  dev_overrides {
    "IBM-Cloud/ibm" = "/Users/suruthiganesankalavathy/go/bin"
  }
  direct {}
}
```

### 2. Configure Variables

```bash
cd examples/ibm-logs
cp terraform.tfvars.template terraform.tfvars
```

Edit `terraform.tfvars`:
```hcl
ibmcloud_api_key = "YOUR_API_KEY"
region           = "us-south"  # or your region
logs_instance_id = "YOUR_INSTANCE_GUID"
```

## Quick Test (Automated)

Run all tests automatically:

```bash
./test-all-apis.sh
```

This will test all 6 APIs in sequence with prompts between each step.

## Manual Testing

### Test 1: List All Extensions

```bash
terraform init
terraform apply -target=data.ibm_logs_extensions.all_extensions -auto-approve
terraform output all_extensions
```

**API:** `GET /v1/extensions`

### Test 2: Get Extension Details

```bash
terraform apply -target=data.ibm_logs_extension.cloudant_extension -auto-approve
terraform output cloudant_extension_details
```

**API:** `GET /v1/extensions/IBMCloudant`

### Test 3: Create Deployment

```bash
terraform apply -auto-approve
terraform output deployment_id
```

**API:** `POST /v1/extensions/IBMCloudant/deployment`

### Test 4: Read Deployment

```bash
terraform output read_deployment_details
```

**API:** `GET /v1/extensions/IBMCloudant/deployment`

### Test 5: Update Deployment

Edit `main.tf` and add:
```hcl
resource "ibm_logs_extension_deployment" "cloudant_deployment" {
  # ... existing config ...
  applications = ["my-app"]
  subsystems   = ["my-subsystem"]
}
```

Then:
```bash
terraform apply -auto-approve
terraform output deployment_details
```

**API:** `PUT /v1/extensions/IBMCloudant/deployment`

### Test 6: Import Deployment

```bash
# Get the resource ID format: region/instance_id/extension_id
terraform import ibm_logs_extension_deployment.cloudant_deployment \
  us-south/YOUR_INSTANCE_GUID/IBMCloudant
```

### Test 7: Delete Deployment

```bash
terraform destroy -auto-approve
```

**API:** `DELETE /v1/extensions/IBMCloudant/deployment`

## Verify Each API Call

After each operation, check:

1. **Exit Code**: Should be 0 for success
2. **Output**: Should show expected data
3. **Terraform State**: `terraform show` to see current state

## Common Issues

### "Provider not found"
- Check `~/.terraformrc` path matches your `$GOPATH/bin`
- Run `which terraform-provider-ibm` to verify binary location
- Run `terraform init` again

### "Authentication failed"
- Verify API key in `terraform.tfvars`
- Check API key has Cloud Logs permissions
- Ensure region matches your instance location

### "Extension not found"
- Run `terraform output all_extensions` to see available extensions
- Check extension ID spelling (case-sensitive)

### "Validation error on update"
- This is fixed in the code - version and item_ids are always included
- If you still see this, rebuild the provider: `make build`

## Testing Different Extensions

To test with a different extension (e.g., IBMElasticsearch):

1. List available extensions:
   ```bash
   terraform output all_extensions
   ```

2. Update `main.tf`:
   ```hcl
   data "ibm_logs_extension" "my_extension" {
     logs_extension_id = "IBMElasticsearch"  # Change this
     # ... rest of config
   }
   
   resource "ibm_logs_extension_deployment" "my_deployment" {
     extension_id = "IBMElasticsearch"  # Change this
     # ... rest of config
   }
   ```

3. Run tests again

## API Coverage Summary

| API Operation | HTTP Method | Terraform Resource | Status |
|--------------|-------------|-------------------|--------|
| List Extensions | GET | `data.ibm_logs_extensions` | ✅ |
| Get Extension | GET | `data.ibm_logs_extension` | ✅ |
| Create Deployment | POST | `resource.ibm_logs_extension_deployment` | ✅ |
| Read Deployment | GET | `data.ibm_logs_extension_deployment` | ✅ |
| Update Deployment | PUT | `resource.ibm_logs_extension_deployment` | ✅ |
| Delete Deployment | DELETE | `resource.ibm_logs_extension_deployment` | ✅ |

## Clean Up

```bash
terraform destroy -auto-approve
rm -rf .terraform terraform.tfstate*
```

## Next Steps

After successful testing:
1. Document any issues found
2. Test with different extensions
3. Test edge cases (invalid IDs, missing permissions, etc.)
4. Verify import/export functionality
5. Check state file consistency

## Support

For detailed information, see [README.md](README.md)