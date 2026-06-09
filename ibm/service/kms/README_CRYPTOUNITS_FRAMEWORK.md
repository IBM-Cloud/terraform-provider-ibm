# IBM KMS Crypto Units - Terraform Plugin Framework Implementation

## Overview

This is the Terraform Plugin Framework implementation of the `ibm_kms_cryptounits` resource, designed to support modern Terraform features including true ephemeral attributes.

## File

`resource_ibm_kms_cryptounits_framework.go`

## Key Features

### 1. Type-Safe Resource Model

The Framework version uses strongly-typed Go structs for the resource model:

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

### 2. Nested Attribute Models

Nested structures are defined with their own models:

```go
type signatureKeyModel struct {
    Filepath   types.String `tfsdk:"filepath"`
    Passphrase types.String `tfsdk:"passphrase"`
    Owner      types.String `tfsdk:"owner"`
}

type masterBackupKeyModel struct {
    KeyShareFiles types.Set    `tfsdk:"keysharefile"`
    KeyName       types.String `tfsdk:"keyname"`
}

type keyShareFileModel struct {
    Filepath types.String `tfsdk:"filepath"`
    Token    types.String `tfsdk:"token"`
}
```

### 3. Modern Schema Definition

Uses Framework's declarative schema with better semantics:

```go
"signature_key": schema.SingleNestedAttribute{
    Required:    true,
    Description: "Credentials for the user to create sessions with the crypto units.",
    Attributes: map[string]schema.Attribute{
        "filepath": schema.StringAttribute{
            Required:    true,
            Description: "The filepath to store the signature key.",
        },
        "passphrase": schema.StringAttribute{
            Required:    true,
            Sensitive:   true,
            Description: "The passphrase of the signature key.",
        },
        "owner": schema.StringAttribute{
            Required:    true,
            Description: "The owner of the signature key.",
        },
    },
},
```

### 4. Plan Modifiers

Framework supports plan modifiers for better resource lifecycle management:

```go
"instance_id": schema.StringAttribute{
    Required:    true,
    Description: "Key protect or HPCS instance GUID or CRN.",
    PlanModifiers: []planmodifier.String{
        stringplanmodifier.RequiresReplace(),
    },
},
```

### 5. Structured Error Handling

Better error reporting with context:

```go
resp.Diagnostics.AddError(
    "Failed to Initialize KMS Crypto Unit Client",
    fmt.Sprintf("Unable to create KMS crypto unit client: %s", err.Error()),
)
```

## Resource Interface Implementation

The resource implements these Framework interfaces:

- `resource.Resource` - Basic resource operations
- `resource.ResourceWithConfigure` - Provider configuration
- `resource.ResourceWithImportState` - State import support

## CRUD Operations

### Create

1. Extracts configuration from plan using type-safe model
2. Parses and validates signature key configuration
3. Parses and validates master backup key configuration
4. Initializes crypto units via API
5. Reads current state and updates model
6. Saves state

### Read

1. Retrieves current state
2. Lists crypto units from API
3. Updates crypto units map in state
4. Saves updated state

### Update

1. Reads plan configuration
2. Refreshes crypto units state
3. Updates state (resource is mostly immutable)

### Delete

1. Lists all crypto units
2. Zeroizes each crypto unit
3. Removes resource from state

## Configuration Parsing

### Master Backup Key

The `parseMasterBackupKey` method:
- Extracts key share files from nested set
- Validates filepath and token for each entry
- Resolves relative paths to absolute paths
- Checks for duplicate filepaths
- Validates key name (max 8 characters)
- Returns `MBKKeySpec` for API call

### Signature Key

The `parseSignatureKey` method:
- Extracts filepath, passphrase, and owner
- Resolves relative path to absolute path
- Validates required fields
- Returns `GenerateSignatureKeyRequest` for API call

## Path Resolution

The `resolveRelativePathFramework` function:
- Converts relative paths to absolute based on current working directory
- Returns absolute paths unchanged
- Ensures consistent path handling across platforms

## Client Integration

Uses the existing `conns.ClientSession` for API client initialization:

```go
func (r *kmsCryptoUnitsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
    client, ok := req.ProviderData.(conns.ClientSession)
    if !ok {
        resp.Diagnostics.AddError(
            "Unexpected Resource Configure Type",
            fmt.Sprintf("Expected conns.ClientSession, got: %T", req.ProviderData),
        )
        return
    }
    r.client = client
}
```

## State Management

### Crypto Units Map

The resource maintains a map of crypto unit IDs to their states:

```go
cryptoUnitsMap := make(map[string]attr.Value)
if cryptoUnitsResponse != nil && cryptoUnitsResponse.CryptoUnits != nil {
    for _, cu := range cryptoUnitsResponse.CryptoUnits {
        if cu.ID != "" && cu.State != "" {
            cryptoUnitsMap[cu.ID] = types.StringValue(string(cu.State))
        }
    }
}
```

## Import Support

The resource supports state import using the instance ID:

```go
func (r *kmsCryptoUnitsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
```

## Usage Example

```hcl
resource "ibm_kms_cryptounits" "example" {
  instance_id = "12345678-1234-1234-1234-123456789012"
  region      = "us-south"

  signature_key {
    filepath   = "./signature_key.p12"
    passphrase = "my-secure-passphrase"
    owner      = "admin"
  }

  master_backup_key {
    keyname = "MBK01"
    
    keysharefile {
      filepath = "./keyshare1.p12"
      token    = "token1"
    }
    
    keysharefile {
      filepath = "./keyshare2.p12"
      token    = "token2"
    }
  }
}
```

## Advantages Over SDK v2 Version

1. **Type Safety**: Compile-time type checking prevents runtime errors
2. **Better Validation**: Framework provides built-in validation
3. **Cleaner Code**: Less boilerplate, more declarative
4. **Future-Ready**: Prepared for ephemeral attributes
5. **Better Testing**: Framework testing utilities
6. **Improved Diagnostics**: Structured error reporting
7. **Plan Modifiers**: Better lifecycle management

## Future Enhancements

### Ephemeral Attributes

Once Terraform fully supports ephemeral attributes, sensitive fields can be marked as ephemeral:

```go
"passphrase": schema.StringAttribute{
    Required:  true,
    Sensitive: true,
    Ephemeral: true,  // Not stored in state
},
```

This will ensure:
- Sensitive data never written to state file
- Data only available during resource operations
- Automatic cleanup after use
- Enhanced security posture

## Testing Considerations

When testing this resource:

1. **Unit Tests**: Test parsing functions independently
2. **Integration Tests**: Test full CRUD lifecycle
3. **State Tests**: Verify state management
4. **Import Tests**: Test state import functionality
5. **Error Tests**: Verify error handling paths
6. **Validation Tests**: Test input validation

## Compatibility

- Fully compatible with existing SDK v2 version
- Same resource name and attributes
- Existing state files work without migration
- Can coexist with SDK v2 version during transition

## Registration

The resource is registered in the Framework provider:

```go
// In ibm/provider_framework/provider.go
func (p *frameworkProvider) Resources(ctx context.Context) []func() resource.Resource {
    return []func() resource.Resource{
        kms.NewKmsCryptoUnitsResource,
    }
}
```

## References

- Original SDK v2 implementation: `resource_ibm_kms_cryptounits.go`
- Migration guide: `FRAMEWORK_MIGRATION.md`
- [Terraform Plugin Framework](https://developer.hashicorp.com/terraform/plugin/framework)