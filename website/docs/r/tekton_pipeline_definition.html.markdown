---
layout: "ibm"
page_title: "IBM : ibm_cd_tekton_pipeline_definition"
description: |-
  Manages tekton_pipeline_definition.
subcategory: "CD Tekton Pipeline"
---

# ibm_cd_tekton_pipeline_definition

Provides a resource for tekton_pipeline_definition. This allows tekton_pipeline_definition to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cd_tekton_pipeline_definition" "tekton_pipeline_definition" {
  pipeline_id = "94619026-912b-4d92-8f51-6c74f0692d90"
  scm_source {
		url = "url"
		branch = "branch"
		tag = "tag"
		path = "path"
  }
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `pipeline_id` - (Required, Forces new resource, String) The tekton pipeline ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `scm_source` - (Optional, List) Scm source for tekton pipeline defintion.
Nested scheme for **scm_source**:
	* `branch` - (Optional, String) A branch of the repo, branch field doesn't coexist with tag field.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
	* `path` - (Required, String) The path to the definitions yaml files.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
	* `tag` - (Optional, String) A tag of the repo.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_]{1,235}$/`.
	* `url` - (Required, String) General href URL.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the tekton_pipeline_definition.
* `definition_id` - (Required, String) UUID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `service_instance_id` - (Required, String) UUID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.

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

For more informaton, see [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#authentication).

## Import

You can import the `ibm_cd_tekton_pipeline_definition` resource by using `id`.
The `id` property can be formed from `pipeline_id`, and `definition_id` in the following format:

```
<pipeline_id>/<definition_id>
```
* `pipeline_id`: A string in the format `94619026-912b-4d92-8f51-6c74f0692d90`. The tekton pipeline ID.
* `definition_id`: A string in the format `94299034-d45f-4e9a-8ed5-6bd5c7bb7ada`. The definition ID.

# Syntax
```
$ terraform import ibm_cd_tekton_pipeline_definition.tekton_pipeline_definition <pipeline_id>/<definition_id>
```
