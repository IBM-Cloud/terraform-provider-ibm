---
layout: "ibm"
page_title: "IBM : ibm_sm_arbitrary_secret"
description: |-
  Manages ArbitrarySecret.
subcategory: "Secrets Manager"
---

# ibm_sm_arbitrary_secret

Provides a resource for ArbitrarySecret. This allows ArbitrarySecret to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_sm_arbitrary_secret" "sm_arbitrary_secret" {
  instance_id   = "6ebc4224-e983-496a-8a54-f40a0bfa9175"
  region        = "us-south"
  custom_metadata = {"key":"value"}
  description = "Extended description for this secret."
  expiration_date = "2022-04-12T23:20:50.520Z"
  labels = ["my-label"]
  payload = "secret-credentials"
  secret_group_id = "default"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `custom_metadata` - (Optional, Map) The secret metadata that a user can customize.
* `description` - (Optional, String) An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/(.*?)/`.
* `expiration_date` - (Optional, Forces new resource, String) The date a secret is expired. The date format follows RFC 3339.
* `labels` - (Optional, List) Labels that you can use to search for secrets in your instance.Up to 30 labels can be created.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `30` items. The minimum length is `0` items.
* `name` - (Required, String) The human-readable name of your secret.
  * Constraints: The maximum length is `256` characters. The minimum length is `2` characters. The value must match regular expression `/^\\w(([\\w-.]+)?\\w)?$/`.
* `payload` - (Required, Forces new resource, String) The arbitrary secret's data payload.
  * Constraints: The maximum length is `100000` characters. The minimum length is `0` characters. The value must match regular expression `/(.*?)/`.
* `secret_group_id` - (Optional, Forces new resource, String) A v4 UUID identifier, or `default` secret group.
  * Constraints: The maximum length is `36` characters. The minimum length is `7` characters. The value must match regular expression `/^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|default)$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `secret_id` - The unique identifier of the ArbitrarySecret.
* `created_at` - (String) The date when a resource was created. The date format follows RFC 3339.
* `created_by` - (String) The unique identifier that is associated with the entity that created the secret.
  * Constraints: The maximum length is `128` characters. The minimum length is `4` characters.
* `crn` - (String) A CRN that uniquely identifies an IBM Cloud resource.
  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
* `downloaded` - (Boolean) Indicates whether the secret data that is associated with a secret version was retrieved in a call to the service API.
* `locks_total` - (Integer) The number of locks of the secret.
  * Constraints: The maximum value is `1000`. The minimum value is `0`.
* `secret_type` - (String) The secret type. Supported types are arbitrary, certificates (imported, public, and private), IAM credentials, key-value, and user credentials.
  * Constraints: Allowable values are: `arbitrary`, `imported_cert`, `public_cert`, `iam_credentials`, `kv`, `username_password`, `private_cert`.
* `state` - (Integer) The secret state that is based on NIST SP 800-57. States are integers and correspond to the `Pre-activation = 0`, `Active = 1`,  `Suspended = 2`, `Deactivated = 3`, and `Destroyed = 5` values.
  * Constraints: Allowable values are: `0`, `1`, `2`, `3`, `5`.
* `state_description` - (String) A text representation of the secret state.
  * Constraints: Allowable values are: `pre_activation`, `active`, `suspended`, `deactivated`, `destroyed`.
* `updated_at` - (String) The date when a resource was recently modified. The date format follows RFC 3339.
* `versions_total` - (Integer) The number of versions of the secret.
  * Constraints: The maximum value is `50`. The minimum value is `0`.

## Provider Configuration

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

To find which credentials are required for this resource, see the service table [here](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-provider-reference#required-parameters).

### Static credentials

You can provide your static credentials by adding the `ibmcloud_api_key`, `iaas_classic_username`, and `iaas_classic_api_key` arguments in the IBM Cloud provider block.

Usage:
```
provider "ibm" {
    ibmcloud_api_key = ""
    iaas_classic_username = ""
    iaas_classic_api_key = ""
}
```

### Environment variables

You can provide your credentials by exporting the `IC_API_KEY`, `IAAS_CLASSIC_USERNAME`, and `IAAS_CLASSIC_API_KEY` environment variables, representing your IBM Cloud platform API key, IBM Cloud Classic Infrastructure (SoftLayer) user name, and IBM Cloud infrastructure API key, respectively.

```
provider "ibm" {}
```

Usage:
```
export IC_API_KEY="ibmcloud_api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="iaas_classic_api_key"
terraform plan
```

Note:

1. Create or find your `ibmcloud_api_key` and `iaas_classic_api_key` [here](https://cloud.ibm.com/iam/apikeys).
  - Select `My IBM Cloud API Keys` option from view dropdown for `ibmcloud_api_key`
  - Select `Classic Infrastructure API Keys` option from view dropdown for `iaas_classic_api_key`
2. For iaas_classic_username
  - Go to [Users](https://cloud.ibm.com/iam/users)
  - Click on user.
  - Find user name in the `VPN password` section under `User Details` tab

For more information, see [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#authentication).

## Import

You can import the `ibm_sm_arbitrary_secret` resource by using `region`, `instance_id`, and `secret_id`.
For more information, see [the documentation](https://cloud.ibm.com/docs/secrets-manager)

# Syntax
```
$ terraform import ibm_sm_arbitrary_secret.sm_arbitrary_secret <region>/<instance_id>/<secret_id>
```

# Example
```
$ terraform import ibm_sm_arbitrary_secret.sm_arbitrary_secret us-east/6ebc4224-e983-496a-8a54-f40a0bfa9175/b49ad24d-81d4-5ebc-b9b9-b0937d1c84d5
```
