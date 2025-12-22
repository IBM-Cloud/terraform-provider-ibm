---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : security_group"
description: |-
  Manages IBM Cloud Security Group.
---

# ibm_is_security_group
Create, delete, and update a security group. Provides a networking security group resource that controls access to the public and private interfaces of a virtual server instance. To create rules for the security group, use the `is_security_group_rule` resource. For more information, about security group, see API Docs(https://cloud.ibm.com/docs/vpc?topic=vpc-using-security-groups).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_security_group" "example" {
  name = "example-security-group"
  vpc  = ibm_is_vpc.example.id
}
```


## Argument reference
Review the argument references that you can specify for your resource. 

- `access_tags`  - (Optional, List of Strings) A list of access management tags to attach to the security group.

  ~> **Note:** 
  **&#x2022;** You can attach only those access tags that already exists.</br>
  **&#x2022;** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#create-access-console).</br>
  **&#x2022;** You must have the access listed in the [Granting users access to tag resources](https://cloud.ibm.com/docs/account?topic=account-access) for `access_tags`</br>
  **&#x2022;** `access_tags` must be in the format `key:value`.
- `name` - (Optional, String) The security group name.
- `resource_group` - (Optional, String) The resource group ID where the security group to be created.
- `tags`- (Optional, List of Strings) The tags associated with an instance.
- `vpc` - (Required, Forces new resource, String) The VPC ID.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN of the security group.
- `id` - (String) The ID of the security group.
- `rules` - (List of Objects) A nested block describes the rules of this security group. Nested `rules` blocks have the following structure.

  Nested scheme for `rules`:
  - `code` - (String) The `ICMP` traffic code to allow.
  - `direction`-  (String) The direction of the traffic either `inbound` or `outbound`.
  - `ip_version` - (String) IP version: `ipv4`
  - `local` - (String) 	The local IP address or range of local IP addresses to which this rule will allow inbound traffic (or from which, for outbound traffic). A CIDR block of 0.0.0.0/0 allows traffic to all local IP addresses (or from all local IP addresses, for outbound rules). an IP address, a `CIDR` block.
  - `protocol` - (String) The type of the protocol `all`, `icmp`, `tcp`, `udp`.
  - `port_max`- (Integer) The `TCP/UDP` port range that includes the maximum bound.
  - `port_min`- (Integer) The `TCP/UDP` port range that includes the minimum bound.
  - `remote` - (String) Security group id, an IP address, a `CIDR` block, or a single security group identifier.
  - `type` - (String) The `ICMP` traffic type to allow.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_security_group` resource by using `id`.
The `id` property can be formed from `security group ID`. For example:

```terraform
import {
  to = ibm_is_security_group.example
  id = "<security_group_id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_security_group.example <security_group_id>
```