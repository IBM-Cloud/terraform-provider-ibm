---
layout: "ibm"
page_title: "IBM : ibm_hpcs_keystore"
description: |-
  Get information about keystore
subcategory: "Hyper Protect Crypto Services"
---

# ibm_hpcs_keystore

Provides a read-only data source for keystore. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_hpcs_keystore" "keystore" {
  instance_id = "76195d24-8a31-4c6d-9050-c35f09375cfb"
  region = "us-east"
  keystore_id = "d8cc1ef7-d13b-4731-95be-1f7c98c9f524"
  uko_vault = ibm_hpcs_vault.vault.vault_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `instance_id` - (Required, String) ID of UKO Instance
  * Constraints: Must match the ID of the UKO instance you are trying to work with.
* `region` - (Required, String) Region of the UKO Instance
  * Constraints: Allowable values are: `au-syd`, `in-che`, `jp-osa`, `jp-tok`, `kr-seo`, `eu-de`, `eu-gb`, `ca-tor`, `us-south`, `us-south-test`, `us-east`, `br-sao`.
* `keystore_id` - (Required, Forces new resource, String) UUID of the keystore.
* `uko_vault` - (Required, String) The UUID of the Vault in which the update is to take place.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `keystore_id` - The unique identifier of the keystore.
* `aws_access_key_id` - (String) The access key id used for connecting to this instance of AWS KMS.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-_]*$/`.

* `aws_region` - (String) AWS Region.
  * Constraints: Allowable values are: `af_south_1`, `ap_east_1`, `ap_northeast_1`, `ap_northeast_2`, `ap_south_1`, `ap_southeast_1`, `ap_southeast_2`, `aws_cn_global`, `aws_global`, `aws_iso_global`, `aws_iso_b_global`, `aws_us_gov_global`, `ca_central_1`, `cn_north_1`, `cn_northwest_1`, `eu_central_1`, `eu_west_1`, `eu_west_2`, `eu_west_3`, `me_south_1`, `sa_east_1`, `us_east_1`, `us_east_2`, `us_gov_east_1`, `us_gov_west_1`, `us_iso_east_1`, `us_isob_east_1`, `us_west_1`, `us_west_2`.

* `aws_secret_access_key` - (String) The secret access key used for connecting to this instance of AWS KMS.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.

* `azure_environment` - (String) Azure environment, usually 'Azure'.
  * Constraints: Allowable values are: `azure`, `azure_china`, `azure_germany`, `azure_us_government`.

* `azure_location` - (String) Location of the Azure Key Vault.
  * Constraints: Allowable values are: `asia_east`, `asia_southeast`, `australia_central`, `australia_central_2`, `australia_east`, `australia_southeast`, `brazil_south`, `canada_central`, `canada_east`, `china_east`, `china_east_2`, `china_north`, `china_north_2`, `europe_north`, `europe_west`, `france_central`, `france_south`, `germany_central`, `germany_northeast`, `india_central`, `india_south`, `india_west`, `japan_east`, `japan_west`, `korea_central`, `korea_south`, `south_africa_north`, `south_africa_west`, `uk_south`, `uk_west`, `us_central`, `us_dod_central`, `us_dod_east`, `us_east`, `us_east_2`, `us_gov_arizona`, `us_gov_iowa`, `us_gov_texas`, `us_gov_virginia`, `us_north_central`, `us_south_central`, `us_west`, `us_west_2`, `us_west_central`.

* `azure_resource_group` - (String) Resource group in Azure.
  * Constraints: The maximum length is `90` characters. The minimum length is `1` character. The value must match regular expression `/^[-\\w\\._\\(\\)]*[^\\.]$/`.

* `azure_service_name` - (String) Service name of the key vault instance from the Azure portal.
  * Constraints: The maximum length is `24` characters. The minimum length is `3` characters. The value must match regular expression `/^[A-Za-z0-9-]+$/`.

* `azure_service_principal_client_id` - (String) Azure service principal client ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z]+$/`.

* `azure_service_principal_password` - (String) Azure service principal password.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.

* `azure_subscription_id` - (String) Subscription ID in Azure.
  * Constraints: The maximum length is `36` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z]+$/`.

* `azure_tenant` - (String) Azure tenant that the Key Vault is associated with,.
  * Constraints: The maximum length is `36` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z]+$/`.

* `created_at` - (String) Date and time when the target keystore was created.

* `created_by` - (String) ID of the user that created the key.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-]+$/`.

* `description` - (String) Description of the keystore.
  * Constraints: The maximum length is `200` characters. The minimum length is `0` characters. The value must match regular expression `/(.|\\n)*/`.

* `google_credentials` - (String) The value of the JSON key represented in the Base64 format.
  * Constraints: The maximum length is `524288` characters. The minimum length is `1` character. The value must match regular expression `/^(?:[A-Za-z0-9+\/]{4})*(?:[A-Za-z0-9+\/]{2}==|[A-Za-z0-9+\/]{3}=)?$/`.

* `google_key_ring` - (String) A key ring organizes keys in a specific Google Cloud location and allows you to manage access control on groups of keys.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-]*$/`.

* `google_location` - (String) Location represents the geographical region where a Cloud KMS resource is stored and can be accessed. A key's location impacts the performance of applications using the key.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-]*$/`.

* `google_private_key_id` - (String) The private key id associated with this keystore.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-_]*$/`.

* `google_project_id` - (String) The project id associated with this keystore.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-_]*$/`.

* `groups` - (List) List of groups that this keystore belongs to.
  * Constraints: The list items must match regular expression `/^[A-Za-z0-9][A-Za-z0-9-_ ]+$/`. The maximum length is `128` items. The minimum length is `1` item.

* `href` - (String) A URL that uniquely identifies your cloud resource.
  * Constraints: The maximum length is `200` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._~:\/?&=-]+$/`.

* `ibm_api_endpoint` - (String) API endpoint of the IBM Cloud keystore.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._~:\/-]+$/`.

* `ibm_api_key` - (String) The IBM Cloud API key to be used for connecting to this IBM Cloud keystore.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-_&.]*$/`.

* `ibm_iam_endpoint` - (String) Endpoint of the IAM service for this IBM Cloud keystore.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._~:\/-]+$/`.

* `ibm_instance_id` - (String) The instance ID of the IBM Cloud keystore.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-]*$/`.

* `ibm_key_ring` - (String) The key ring of an IBM Cloud KMS Keystore.
  * Constraints: The default value is `Default`. The maximum length is `100` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.

* `ibm_variant` - (String) Possible IBM Cloud KMS variants.
  * Constraints: Allowable values are: `hpcs`, `internal`, `key_protect`.

* `location` - (String) Geographic location of the keystore, if available.
  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9][A-Za-z0-9 ._-]*$/`.

* `name` - (String) Name of the target keystore. It can be changed in the future.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9][A-Za-z0-9 ._-]*$/`.

* `type` - (String) Type of keystore.
  * Constraints: Allowable values are: `aws_kms`, `azure_key_vault`, `ibm_cloud_kms`, `google_kms`.

* `updated_at` - (String) Date and time when the target keystore was last updated.

* `updated_by` - (String) ID of the user that last updated the key.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-]+$/`.

* `vault` - (List) Reference to a vault.
Nested scheme for **vault**:
	* `href` - (String) A URL that uniquely identifies your cloud resource.
	  * Constraints: The maximum length is `200` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._~:\/?&=-]+$/`.
	* `id` - (String) The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
	* `name` - (String) Name of the referenced vault.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9#@!$%'_-][A-Za-z0-9#@!$% '_-]*$/`.

