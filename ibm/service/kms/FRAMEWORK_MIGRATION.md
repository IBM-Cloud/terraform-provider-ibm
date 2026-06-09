# KMS Crypto Units Framework Migration

## Overview

This document describes the migration of `ibm_kms_cryptounits` resource from Terraform Plugin SDK v2 to Terraform Plugin Framework.

## Files

- **Original (SDK v2)**: `resource_ibm_kms_cryptounits.go`
- **Framework Version**: `resource_ibm_kms_cryptounits_framework.go`

## Key Differences

### 1. Framework Support

**SDK v2 Version:**
- Uses `github.com/hashicorp/terraform-plugin-sdk/v2`
- Does NOT support true ephemeral attributes
- Sensitive data (tokens, passphrases) stored in state file

**Framework Version:**
- Uses `github.com/hashicorp/terraform-plugin-framework`
- Supports true ephemeral attributes (future enhancement)
- Better type safety with structured models
- More explicit schema definitions

### 2. Resource Registration

**SDK v2:**
```go
// In ibm/provider/provider.go
"ibm_kms_cryptounits": kms.ResourceIBMKmsCryptoUnits(),
```

**Framework:**
```go
// In ibm/provider_framework/provider.go
func (p *frameworkProvider) Resources(ctx context.Context) []func() resource.Resource {
    return []func() resource.Resource{
        kms.NewKmsCryptoUnitsResource,
    }
}
```

### 3. Schema Definition

**SDK v2:**
```go
Schema: map[string]*schema.Schema{
    "signature_key": {
        Type:     schema.TypeSet,
        Required: true,
        MaxItems: 1,
        Elem: &schema.Resource{
            Schema: map[string]*schema.Schema{
                "passphrase": {
                    Type:      schema.TypeString,
                    Required:  true,
                    Sensitive: true,
                },
            },
        },
    },
}
```

**Framework:**
```go
"signature_key": schema.SingleNestedAttribute{
    Required: true,
    Attributes: map[string]schema.Attribute{
        "passphrase": schema.StringAttribute{
            Required:  true,
            Sensitive: true,
        },
    },
},
```

### 4. Data Model

**SDK v2:**
- Uses `*schema.ResourceData` directly
- Type assertions required for nested structures
- Manual parsing with `GetOk()` and type casting

**Framework:**
- Strongly-typed Go structs with `tfsdk` tags
- Automatic marshaling/unmarshaling
- Type-safe access to attributes

```go
type kmsCryptoUnitsResourceModel struct {
    ID               types.String `tfsdk:"id"`
    InstanceID       types.String `tfsdk:"instance_id"`
    Region           types.String `tfsdk:"region"`
    CryptoUnits      types.Map    `tfsdk:"cryptounits"`
    SignatureKey     types.Object `tfsdk:"signature_key"`
    MasterBackupKey  types.Object `tfsdk:"master_backup_key"`
}
```

### 5. CRUD Operations

**SDK v2:**
```go
func resourceIBMKmsCryptoUnitsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    region, ok := d.GetOk("region")
    // Manual type assertions and error handling
}
```

**Framework:**
```go
func (r *kmsCryptoUnitsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    var plan kmsCryptoUnitsResourceModel
    resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
    // Type-safe access to plan data
    region := plan.Region.ValueString()
}
```

### 6. Error Handling

**SDK v2:**
```go
return diag.Errorf("failed to initialize: %v", err)
```

**Framework:**
```go
resp.Diagnostics.AddError(
    "Failed to Initialize KMS Crypto Unit Client",
    fmt.Sprintf("Unable to create KMS crypto unit client: %s", err.Error()),
)
```

## Future Enhancements

### Ephemeral Attributes

The Framework version is designed to support ephemeral attributes in the future:

```go
// Future enhancement - mark sensitive fields as ephemeral
"signature_key": schema.SingleNestedAttribute{
    Required: true,
    Attributes: map[string]schema.Attribute{
        "passphrase": schema.StringAttribute{
            Required:  true,
            Sensitive: true,
            Ephemeral: true,  // Future: Not stored in state
        },
    },
},
```

This would ensure that sensitive data like passphrases and tokens are:
- Never written to the state file
- Only available during resource creation
- Automatically cleared after use

## Migration Path

### For Users

The Framework version maintains full compatibility with the SDK v2 version:
- Same resource name: `ibm_kms_cryptounits`
- Same attributes and behavior
- Existing state files remain compatible

### For Developers

To use the Framework version:

1. The resource is automatically registered in the Framework provider
2. Both versions can coexist during transition
3. The Framework version will be used for new deployments
4. Existing deployments continue using SDK v2 version

## Testing

Both versions should be tested to ensure:
- Identical behavior for all operations (Create, Read, Update, Delete)
- Proper handling of sensitive data
- Correct state management
- Error handling consistency

## Benefits of Framework Version

1. **Better Type Safety**: Compile-time type checking for resource models
2. **Cleaner Code**: Less boilerplate, more declarative
3. **Future-Proof**: Ready for ephemeral attributes and other Framework features
4. **Better Diagnostics**: More structured error reporting
5. **Improved Testing**: Framework provides better testing utilities

## Implementation Notes

### Nested Attributes

The Framework version uses `SingleNestedAttribute` instead of `TypeSet` with `MaxItems: 1`:
- More semantically correct
- Better type safety
- Clearer intent in schema

### Path Resolution

Both versions use the same `resolveRelativePath` logic, but the Framework version has it renamed to `resolveRelativePathFramework` to avoid conflicts.

### Client Configuration

Both versions use the same `conns.ClientSession` for API client initialization, ensuring consistent behavior.

## References

- [Terraform Plugin Framework Documentation](https://developer.hashicorp.com/terraform/plugin/framework)
- [Migration Guide from SDK to Framework](https://developer.hashicorp.com/terraform/plugin/framework/migrating)
- [Ephemeral Resources RFC](https://github.com/hashicorp/terraform/issues/31796)