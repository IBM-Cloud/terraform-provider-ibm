---
layout: "ibm"
page_title: "IBM : ibm_cm_validation"
description: |-
  Manages ibm_cm_validation.
subcategory: "Catalog Management"
---

# ibm_cm_validation

Provides a resource for ibm_cm_validation. This allows ibm_cm_validation to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cm_validation" "cm_version_validation" {
  version_locator = ibm_cm_version.my_cm_version_tf.version_locator
  revalidate_if_validated = false
  override_values = {
    <example_override_key1> = <example_override_value1>
    <example_override_key2> = <example_override_value2>
  }
  mark_version_consumable = true
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `version_locator` - (Required, Forces new resource, String) Version locator - the version that will be validated.
* `region` - (Optional, Forces new resource, String) Validation region.
* `override_values` - (Optional, Forces new resource, Map) Map of override values to be used in validation.
* `environment_variables` - (List) List of environment variables to pass to Schematics.
Nested scheme for **environment_variables**:
	* `name` - (Optional, String) Name of the environment variable.
	* `value` - (Optional, String) Value of the environment variable.
	* `secure` - (Optional, Bool) If the environment variablel should be secure.
* `schematics` - (List) Other values to pass to Schematics.
Nested scheme for **schematics**:
	* `name` - (Optional, String) Name for the schematics workspace.
	* `description` - (Optional, String) Description for the schematics workspace.
	* `resource_group_id` - (Optional, String) The resource group ID.
	* `terraform_version` - (Optional, String) Version of terraform to use in schematics.
	* `region` - (Optional, String) Region to use for the schematics installation.
	* `tags` - (Optional, List) List of tags for the schematics workspace.
* `revalidate_if_validated` - (Optional, Forces new resource, Bool) If the version should be revalidated if it is already validated.
* `mark_version_consumable` - (Optional, Bool) If the version should be marked as consumable after validation, aka \"ready to share\".

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `version_locator` - (String) Version locator - the version that will be validated.
* `region` - (String) Validation region.
* `override_values` - (Map) Map of override values to be used in validation.
* `environment_variables` - (List) List of environment variables to pass to Schematics.
Nested scheme for **environment_variables**:
	* `name` - (String) Name of the environment variable.
	* `value` - (String) Value of the environment variable.
	* `secure` - (Bool) If the environment variablel should be secure.
* `schematics` - (List) Other values to pass to Schematics.
Nested scheme for **schematics**:
	* `name` - (String) Name for the schematics workspace.
	* `description` - (String) Description for the schematics workspace.
	* `resource_group_id` - (String) The resource group ID.
	* `terraform_version` - (String) Version of terraform to use in schematics.
	* `region` - (String) Region to use for the schematics installation.
	* `tags` - (List) List of tags for the schematics workspace.
* `validated` - (String) Data and time of last successful validation.
* `requested` - (String) Data and time of last validation request.
* `state` - (String) Current validation state - <empty>, in_progress, valid, invalid, expired.
* `last_operation` - (String) Last operation (e.g. submit_deployment, generate_installer, install_offering.
* `target` - (Map) Validation target information (e.g. cluster_id, region, namespace, etc).  Values will vary by Content type.
* `message` - (String) Any message needing to be conveyed as part of the validation job.
* `revalidate_if_validated` - (Bool) If the version should be revalidated if it is already validated.
* `mark_version_consumable` - (Bool) If the version should be marked as consumable after validation, aka \"ready to share\".

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
export IC_API_KEY="api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="api_key"
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

For more informaton, see [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#authentication).
