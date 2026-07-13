---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_vpc_default_network_acl"
description: |-
  Manages IBM VPC default network ACL.
---

# ibm_is_vpc_default_network_acl

Provides a resource to manage the default network access control list (ACL) of a VPC. This resource allows you to manage the default network ACL that is automatically created when a VPC is created. Unlike custom network ACLs, the default network ACL cannot be deleted but can be configured to control network traffic behavior.

~> **NOTE:** This is an advanced resource with special caveats. Please read this document in its entirety before using this resource. The `ibm_is_vpc_default_network_acl` resource behaves differently from normal resources. Terraform does not _create_ this resource but instead attempts to "adopt" it into management.

Every VPC has a default network ACL that can be managed but not destroyed. When Terraform first adopts a default network ACL, it manages the configuration of the existing default network ACL. This resource allows you to configure the name and tags for the default network ACL.

For more information, about network ACL, see [setting up network ACLs](https://cloud.ibm.com/docs/vpc?topic=vpc-using-acls).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

### Basic usage
```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_vpc_default_network_acl" "example" {
  vpc = ibm_is_vpc.example.id
}
```

### Configure with custom name
```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_vpc_default_network_acl" "example" {
  vpc  = ibm_is_vpc.example.id
  name = "my-custom-default-acl"
}
```

### Example usage with tags
```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_vpc_default_network_acl" "example" {
  vpc         = ibm_is_vpc.example.id
  name        = "my-custom-default-acl"
  tags        = ["env:production", "team:networking"]
  access_tags = ["project:web-app"]
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `access_tags` - (Optional, List of Strings) A list of access management tags to attach to the default network ACL.

  ~> **Note:** 
  **&#x2022;** You can attach only those access tags that already exists.</br>
  **&#x2022;** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#create-access-console).</br>
  **&#x2022;** You must have the access listed in the [Granting users access to tag resources](https://cloud.ibm.com/docs/account?topic=account-access) for `access_tags`</br>
  **&#x2022;** `access_tags` must be in the format `key:value`.
- `name` - (Optional, String) The name of the default network ACL. If unspecified, the name will be automatically assigned by IBM Cloud.
- `tags` - (Optional, List of Strings) Tags associated with the default network ACL.
- `vpc` - (Required, Forces new resource, String) The VPC ID.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN of the default network ACL.
- `default_network_acl` - (String) The ID of the default network ACL.
- `id` - (String) The unique identifier of the default network ACL. The ID is composed of `<vpc_id>/<network_acl_id>`.
- `resource_group` - (List) The resource group for this default network ACL.

  Nested scheme for `resource_group`:
  - `href` - (String) The URL for this resource group.
  - `id` - (String) The unique identifier for this resource group.
  - `name` - (String) The user-defined name for this resource group.
- `subnets` - (List) The subnets to which this default network ACL is attached.

  Nested scheme for `subnets`:
  - `id` - (String) The unique identifier for this subnet.
  - `name` - (String) The user-defined name for this subnet.

## Import
The `ibm_is_vpc_default_network_acl` resource can be imported by using VPC ID and the default network ACL ID.

**Syntax**

```
$ terraform import ibm_is_vpc_default_network_acl.example <vpc_id>/<network_acl_id>
```

**Example**

```
$ terraform import ibm_is_vpc_default_network_acl.example 56738c92-4631-4eb5-8938-8af9211a6ea4/d7bec597-4726-451f-8a63-e0ba6a5a11ba
```

## Differences from custom network ACLs

The default network ACL resource differs from custom network ACLs (`ibm_is_network_acl`) in several key ways:

- **Cannot be deleted**: The default network ACL is permanent and cannot be destroyed
- **Always exists**: Created automatically when a VPC is created
- **Only one per VPC**: Each VPC has exactly one default network ACL
- **No rule management**: Individual ACL rules are managed separately using `ibm_is_network_acl_rule` resources

## Managing ACL rules

To manage individual rules in the default network ACL, use the `ibm_is_network_acl_rule` resource with the default network ACL ID:

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_vpc_default_network_acl" "example" {
  vpc  = ibm_is_vpc.example.id
  name = "my-default-acl"
}

# Manage individual rules separately
resource "ibm_is_network_acl_rule" "allow_ssh" {
  network_acl = ibm_is_vpc_default_network_acl.example.default_network_acl
  name        = "allow-ssh"
  action      = "allow"
  source      = "0.0.0.0/0"
  destination = "0.0.0.0/0"
  direction   = "inbound"
  tcp {
    port_min = 22
    port_max = 22
  }
}

resource "ibm_is_network_acl_rule" "allow_outbound" {
  network_acl = ibm_is_vpc_default_network_acl.example.default_network_acl
  name        = "allow-outbound"
  action      = "allow"
  source      = "0.0.0.0/0"
  destination = "0.0.0.0/0"
  direction   = "outbound"
}
```

## Notes

- The default network ACL cannot be deleted. When you run `terraform destroy`, the resource will be removed from Terraform state, but the actual network ACL will remain in IBM Cloud.
- The name of the default network ACL can be customized, unlike some other default resources.
- To manage individual ACL rules in the default network ACL, use the `ibm_is_network_acl_rule` resource with the default network ACL ID.
- Changes to network ACL rules may affect network traffic flow to your VPC. Plan changes carefully.
- The default network ACL is automatically attached to any subnet that doesn't specify a custom network ACL.