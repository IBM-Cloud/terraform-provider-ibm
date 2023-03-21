---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : security_group"
description: |-
  Manages IBM Cloud Security Group.
---

# ibm_is_vpc_default_security_group

Provides a resource to manage a default security group of a VPC. Manage the default security group created alongwith the vpc creation with this resource. Provides a networking security group resource that controls access to the public and private interfaces of a virtual server instance. To create rules for the security group, use the `rules` parameter. 

~> **NOTE:** This is an advanced resource with special caveats. Please read this document in its entirety before using this resource. The `ibm_is_vpc_default_security_group` resource behaves differently from normal resources. Terraform does not _create_ this resource but instead attempts to "adopt" it into management.

Every VPC has a default security group that can be managed but not destroyed. When Terraform first adopts a default security group, it **immediately removes all defined rules**. It then proceeds to create any rules specified in the configuration. This step is required so that only the rules specified in the configuration exist in the default security group.

For more information, about VPC, see [getting started with Virtual Private Cloud](https://cloud.ibm.com/docs/vpc?topic=vpc-getting-started). For more information, about updating default security group, see [updating a VPC's default security group rules](https://cloud.ibm.com/docs/vpc?topic=vpc-updating-the-default-security-group&interface=ui).

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

resource "ibm_is_vpc_default_security_group" "example" {
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
- `rules` - (Optional, List of Objects) A nested block describes the rules of this security group. Nested `rules` blocks have the following structure.

  Nested scheme for `rules`:
  - `code` - (Optional, String) The `ICMP` traffic code to allow.
  - `direction`-  (Optional, String) The direction of the traffic either `inbound` or `outbound`.
  - `ip_version` - (Optional, String) IP version: `ipv4`
  - `protocol` - (Optional, String) The type of the protocol `all`, `icmp`, `tcp`, `udp`.
  - `port_max`- (Optional, Integer) The `TCP/UDP` port range that includes the maximum bound.
  - `port_min`- (Optional, Integer) The `TCP/UDP` port range that includes the minimum bound.
  - `remote` - (Optional, String) Security group id, an IP address, a `CIDR` block, or a single security group identifier.
  - `type` - (Optional, String) The `ICMP` traffic type to allow.
- `tags`- (Optional, List of Strings) The tags associated with an instance.
- `vpc` - (Required, Forces new resource, String) The VPC ID.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN of the security group.
- `id` - (String) The ID of the security group.
- `resource_group` - (String) The resource group ID where the security group resides.

## Import
The `ibm_is_vpc_default_security_group` resource can be imported by using security group ID. 

**Example**

```
$ terraform import ibm_is_vpc_default_security_group.example a1aaa111-1111-111a-1a11-a11a1a11a11a
```
