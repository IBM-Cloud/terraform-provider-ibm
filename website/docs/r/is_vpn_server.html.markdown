---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : VPN-Server"
description: |-
  Manages IBM VPN Server.
---

# ibm_is_vpn_server

Provides a resource for VPNServer. This allows VPNServer to be created, updated and deleted. For more information, about VPN Server, see [Creating a VPN server](https://cloud.ibm.com/docs/vpc?topic=vpc-vpn-create-server&interface=ui).

## Example Usage
The following example creates a VPN Server:

```terraform
resource "ibm_is_vpn_server" "example" {
  certificate_crn = "crn:v1:bluemix:public:cloudcerts:us-south:a/efe5afc483594adaa8325e2b4d1290df:86f62739-f3a8-42ac-abea-f23255965983:certificate:00406b5615f95dba9bf7c2ab52bb3083"
  client_authentication {
    method    = "certificate"
    client_ca_crn = "crn:v1:bluemix:public:cloudcerts:us-south:a/efe5afc483594adaa8325e2b4d1290df:86f62739-f3a8-42ac-abea-f23255965983:certificate:00406b5615f95dba9bf7c2ab52bb3083"
  }
  client_ip_pool         = "10.5.0.0/21"
  client_dns_server_ips  = ["192.168.3.4"]
  client_idle_timeout    = 2800
  enable_split_tunneling = false
  name                   = "example-vpn-server"
  port                   = 443
  protocol               = "udp"
  subnets                = [ibm_is_subnet.subnet1.id]
}
```

## Timeouts
The `ibm_is_vpn_server` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create**: The creation of the VPN server is considered `failed` when no response is received for 10 minutes. 
- **update**: The update of the VPN server is considered `failed` when no response is received for 10 minutes. 
- **delete**: The deletion of the VPN server is considered `failed` when no response is received for 10 minutes. 

## Argument Reference
Review the argument references that you can specify for your resource. 

- `access_tags`  - (Optional, List of Strings) A list of access management tags to attach to the vpn server.

  ~> **Note:** 
  **&#x2022;** You can attach only those access tags that already exists.</br>
  **&#x2022;** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#create-access-console).</br>
  **&#x2022;** You must have the access listed in the [Granting users access to tag resources](https://cloud.ibm.com/docs/account?topic=account-access) for `access_tags`</br>
  **&#x2022;** `access_tags` must be in the format `key:value`.
- `certificate_crn` - (Required, String) The certificate CRN of secret from Secrets Manager for this VPN server. 

  !> **Removal Notification** Certificate Manager support is removed, please use Secrets Manager.

- `client_authentication` - (Required, List) The methods used to authenticate VPN clients to this VPN server.
  
  Nested scheme for **client_authentication**:
	- `method` - (Required, String) The type of authentication.
	  - Constraints: Allowable values are: certificate, username
    
   -> **NOTE:** 
      `identity_provider` and `client_ca_crn` are mutually exclusive, which means either one must be provided. When `method` has `certificate` as value `client_ca_crn` must be provided and when `method` has `username` as value `identity_provider` must be provided.

	- `identity_provider` - (Required, String) The type of identity provider to be used by VPN client.The type of identity provider to be used by the VPN client.- `iam`: IBM identity and access management The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the route on which the unexpected property value was encountered.
		  - Constraints: Allowable values are: iam
	- `client_ca_crn` - (Required, String)  The CRN of the certificate instance or CRN of the secret from secrets manager to use for the VPN client certificate authority (CA). As the usage of certificate CRN from Certificate Manager is getting deprecated, It is recommended to use Secret manger for same.
- `client_dns_server_ips` - (Optional, List) The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered, the DNS server addresses that will be provided to VPN clients connected to this VPN server.
- `client_idle_timeout` - (Optional, Integer) The seconds a VPN client can be idle before this VPN server will disconnect it.   Specify `0` to prevent the server from disconnecting idle clients.
  - Constraints: The maximum value is `28800`. The minimum value is `0`, default is `600`.
- `client_ip_pool` - (Required, String) The VPN client IPv4 address pool, expressed in CIDR format. The request must not overlap with any existing address prefixes in the VPC or any of the following reserved address ranges:  - `127.0.0.0/8` (IPv4 loopback addresses)  - `161.26.0.0/16` (IBM services)  - `166.8.0.0/14` (Cloud Service Endpoints)  - `169.254.0.0/16` (IPv4 link-local addresses)  - `224.0.0.0/4` (IPv4 multicast addresses)The prefix length of the client IP address pool's CIDR must be between`/9` (8,388,608 addresses) and `/22` (1024 addresses). A CIDR block that contains twice the number of IP addresses that are required to enable the maximum number of concurrent connections is recommended.
- `enable_split_tunneling` - (Optional, Boolean) Indicates whether the split tunneling is enabled on this VPN server.
  - Constraints: The default value is `false`.
- `name` - (Optional, String) The user-defined name for this VPN server. If unspecified, the name will be a hyphenated list of randomly-selected words. Names must be unique within the VPC this VPN server is serving.
- `port` - (Optional, Integer) The port number to use for this VPN server.
  - Constraints: The maximum value is `65535`. The minimum value is `1`.
- `protocol` - (Optional, String) The transport protocol to use for this VPN server.
  - Constraints: The default value is `udp`. Allowable values are: udp, tcp
- `resource_group` - (Optional, Forces new resource, String) The resource group (id), where the VPN gateway to be created.
- `security_groups` - (Optional, List) The security groups `ID` to use for this VPN server. If unspecified, the VPC's default security group is used.
- `subnets` - (Required, List) Comma-separated IDs of the subnets to provision this VPN server in.  Use subnets in different zones for high availability. User can also upgrade or downgrade the VPN server to high availability or standalone by adding/remove the subnets.
- `tags`- (Optional, Array of Strings) A list of user tags that you want to add to your VPN server. (https://cloud.ibm.com/apidocs/tagging#types-of-tags)

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the VPNServer.
- `vpn_server` - The unique identifier of the VPNServer.
- `client_auto_delete` - (Boolean) If set to `true`, disconnected VPN clients will be automatically deleted after the `client_auto_delete_timeout` time has passed.
- `client_auto_delete_timeout` - (Integer) Hours after which disconnected VPN clients will be automatically deleted. If `0`, disconnected VPN clients will be deleted immediately.
- `created_at` - (String) The date and time that the VPN server was created.
- `crn` - (String) The CRN for this VPN server.
- `health_reasons` - (List) The reasons for the current health_state (if any).

  Nested scheme for `health_reasons`:
  - `code` - (String) A snake case string succinctly identifying the reason for this health state.
  - `message` - (String) An explanation of the reason for this health state.
  - `more_info` - (String) Link to documentation about the reason for this health state.
- `health_state` - (String) The health of this resource.

  -> **Supported health_state values:** 
    </br>&#x2022; `ok`: Healthy
    </br>&#x2022; `degraded`: Suffering from compromised performance, capacity, or connectivity
    </br>&#x2022; `faulted`: Completely unreachable, inoperative, or otherwise entirely incapacitated
    </br>&#x2022; `inapplicable`: The health state does not apply because of the current lifecycle state. 
      **Note:** A resource with a lifecycle state of `failed` or `deleting` will have a health state of `inapplicable`. A `pending` resource may also have this state.
- `hostname` - (String) Fully qualified domain name assigned to this VPN server.
- `href` - (String) The URL for this VPN server.
- `lifecycle_reasons` - (List) The reasons for the current lifecycle_reasons (if any).

  Nested scheme for `lifecycle_reasons`:
  - `code` - (String) A snake case string succinctly identifying the reason for this lifecycle reason.
  - `message` - (String) An explanation of the reason for this lifecycle reason.
  - `more_info` - (String) Link to documentation about the reason for this lifecycle reason.
- `lifecycle_state` - (String) The lifecycle state of the VPN server.
- `private_ips` - (List) The reserved IPs bound to this VPN server.

  Nested scheme for `private_ips`:
    - `address` - (String) The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
    - `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
    Nested scheme for `deleted`:
      - `more_info` - (String) Link to documentation about deleted resources.
    - `href` - (String) The URL for this reserved IP.
    - `id` - (String) The unique identifier for this reserved IP.
    - `name` - (String) The user-defined or system-provided name for this reserved IP.
    - `resource_type` - (String) The resource type.

- `resource_type` - (String) The type of resource referenced.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_vpn_server` resource by using `id`.
The `id` property can be formed using the appropriate identifier(s). For example:

```terraform
import {
  to = ibm_is_vpn_server.is_vpn_server
  id = "<id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_vpn_server.is_vpn_server <id>
```