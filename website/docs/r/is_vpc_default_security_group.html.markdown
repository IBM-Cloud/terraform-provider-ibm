---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_vpc_default_security_group"
description: |-
  Manages IBM VPC default security group.
---

# ibm_is_vpc_default_security_group

Provides a resource to manage the default security group of a VPC. This resource allows you to manage the default security group that is automatically created when a VPC is created. Unlike custom security groups, the default security group cannot be deleted but can be configured to control network access behavior.

~> **NOTE:** This is an advanced resource with special caveats. Please read this document in its entirety before using this resource. The `ibm_is_vpc_default_security_group` resource behaves differently from normal resources. Terraform does not _create_ this resource but instead attempts to "adopt" it into management.

Every VPC has a default security group that can be managed but not destroyed. When Terraform first adopts a default security group, it manages the configuration of the existing default security group. This resource allows you to configure the name and tags for the default security group.

For more information, about VPC security groups, see [using security groups](https://cloud.ibm.com/docs/vpc?topic=vpc-using-security-groups). For more information, about updating default security group, see [updating a VPC's default security group rules](https://cloud.ibm.com/docs/vpc?topic=vpc-updating-the-default-security-group&interface=ui).

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

resource "ibm_is_vpc_default_security_group" "example" {
  vpc = ibm_is_vpc.example.id
}
```

### Configure with custom name
```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_vpc_default_security_group" "example" {
  vpc  = ibm_is_vpc.example.id
  name = "my-custom-default-sg"
}
```

### Example usage with tags
```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_vpc_default_security_group" "example" {
  vpc         = ibm_is_vpc.example.id
  name        = "my-custom-default-sg"
  tags        = ["env:production", "team:security"]
  access_tags = ["project:web-app"]
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `access_tags` - (Optional, List of Strings) A list of access management tags to attach to the default security group.

  ~> **Note:** 
  **&#x2022;** You can attach only those access tags that already exists.</br>
  **&#x2022;** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#create-access-console).</br>
  **&#x2022;** You must have the access listed in the [Granting users access to tag resources](https://cloud.ibm.com/docs/account?topic=account-access) for `access_tags`</br>
  **&#x2022;** `access_tags` must be in the format `key:value`.
- `name` - (Optional, String) The name of the default security group. If unspecified, the name will be automatically assigned by IBM Cloud.
- `tags` - (Optional, List of Strings) Tags associated with the default security group.
- `vpc` - (Required, Forces new resource, String) The VPC ID.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN of the default security group.
- `default_security_group` - (String) The ID of the default security group.
- `id` - (String) The unique identifier of the default security group. The ID is composed of `<vpc_id>/<security_group_id>`.
- `resource_group` - (List) The resource group for this default security group.

  Nested scheme for `resource_group`:
  - `href` - (String) The URL for this resource group.
  - `id` - (String) The unique identifier for this resource group.
  - `name` - (String) The user-defined name for this resource group.
- `targets` - (List) The targets attached to this default security group.

  Nested scheme for `targets`:
  - `id` - (String) The unique identifier for this target.
  - `name` - (String) The user-defined name for this target.
  - `resource_type` - (String) The resource type of the target.

## Import
The `ibm_is_vpc_default_security_group` resource can be imported by using VPC ID and the default security group ID.

**Syntax**

```
$ terraform import ibm_is_vpc_default_security_group.example <vpc_id>/<security_group_id>
```

**Example**

```
$ terraform import ibm_is_vpc_default_security_group.example 56738c92-4631-4eb5-8938-8af9211a6ea4/a1aaa111-1111-111a-1a11-a11a1a11a11a
```

## Differences from custom security groups

The default security group resource differs from custom security groups (`ibm_is_security_group`) in several key ways:

- **Cannot be deleted**: The default security group is permanent and cannot be destroyed
- **Always exists**: Created automatically when a VPC is created
- **Only one per VPC**: Each VPC has exactly one default security group
- **No rule management**: Individual security group rules are managed separately using `ibm_is_security_group_rule` resources

## Managing security group rules

To manage individual rules in the default security group, use the `ibm_is_security_group_rule` resource with the default security group ID:

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_vpc_default_security_group" "example" {
  vpc  = ibm_is_vpc.example.id
  name = "my-default-sg"
}

# Manage individual rules separately
resource "ibm_is_security_group_rule" "allow_ssh" {
  group     = ibm_is_vpc_default_security_group.example.default_security_group
  direction = "inbound"
  remote    = "0.0.0.0/0"
  tcp {
    port_min = 22
    port_max = 22
  }
}

resource "ibm_is_security_group_rule" "allow_outbound" {
  group     = ibm_is_vpc_default_security_group.example.default_security_group
  direction = "outbound"
  remote    = "0.0.0.0/0"
}
```

## Assigning resources to the default security group

When creating resources like instances or load balancers, you can assign them to the default security group using the `default_security_group` attribute:

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_vpc_default_security_group" "example" {
  vpc  = ibm_is_vpc.example.id
  name = "my-default-sg"
}

resource "ibm_is_subnet" "example" {
  name                     = "example-subnet"
  vpc                      = ibm_is_vpc.example.id
  zone                     = "us-south-1"
  total_ipv4_address_count = 256
}

resource "ibm_is_instance" "example" {
  name    = "example-instance"
  image   = "r006-14140f94-fcc4-11e9-96e7-a72723715315"
  profile = "bx2-2x8"

  primary_network_interface {
    subnet          = ibm_is_subnet.example.id
    security_groups = [ibm_is_vpc_default_security_group.example.default_security_group]
  }

  vpc  = ibm_is_vpc.example.id
  zone = "us-south-1"
  keys = [ibm_is_ssh_key.example.id]
}
```

## Notes

- The default security group cannot be deleted. When you run `terraform destroy`, the resource will be removed from Terraform state, but the actual security group will remain in IBM Cloud.
- The name of the default security group can be customized, unlike some other default resources.
- To manage individual security group rules in the default security group, use the `ibm_is_security_group_rule` resource with the default security group ID.
- Changes to security group rules may affect network traffic flow to your VPC. Plan changes carefully.
- The default security group is automatically assigned to any network interface that doesn't specify custom security groups.
- You can view which resources are currently attached to the default security group through the `targets` attribute.