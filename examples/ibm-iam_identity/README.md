# Example for IBM IAM Account Settings

This example illustrates how to use the IAM Account Settings to configure settings on their account with regards to MFA, session lifetimes, access control for creating new identities, and enforcing IP restrictions on token creation.

These types of resources are supported:

* iam_account_settings

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## IAM Account Setting Resource

Account Setting Resource:

```hcl
resource "iam_account_settings" "iam_account_settings_instance" {
  include_history = var.iam_account_settings_include_history
}
```

##  IAM Account Setting Data Source
Lists all Policies of a particular IBM ID user.

```hcl
data "ibm_iam_account_settings" "iam_account_settings_source" {
}
```

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | n/a |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| include_history | Defines if the entity history is included in the response. | `bool` | false |
| if_match | Version of the account settings to be updated, if no value is supplied then the default value `*` will be used to indicate to update any version available. This might result in stale updates. | `string` | false |
| restrict_create_service_id | Defines whether or not creating a Service Id is access controlled. | `string` | false |
| restrict_create_platform_apikey | Defines whether or not creating platform API keys is access controlled. | `string` | false |
| allowed_ip_addresses | Defines the IP addresses and subnets from which IAM tokens can be created for the account. | `string` | false |
| mfa | Defines the MFA trait for the account. | `string` | false |
| session_expiration_in_seconds | Defines the session expiration in seconds for the account. | `string` | false |
| session_invalidation_in_seconds | Defines the period of time in seconds in which a session will be invalidated due  to inactivity. | `string` | false |
| max_sessions_per_identity | Defines the max allowed sessions per identity required by the account. | `string` | false |




## Outputs

| Name | Description |
|------|-------------|
| iam_account_settings | iam_account_settings object |
