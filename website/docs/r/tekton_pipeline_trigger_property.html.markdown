---
layout: "ibm"
page_title: "IBM : ibm_cd_tekton_pipeline_trigger_property"
description: |-
  Manages tekton_pipeline_trigger_property.
subcategory: "CD Tekton Pipeline"
---

# ibm_cd_tekton_pipeline_trigger_property

Provides a resource for tekton_pipeline_trigger_property. This allows tekton_pipeline_trigger_property to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cd_tekton_pipeline_trigger_property" "tekton_pipeline_trigger_property" {
  name = "key1"
  pipeline_id = "94619026-912b-4d92-8f51-6c74f0692d90"
  trigger_id = "1bb892a1-2e04-4768-a369-b1159eace147"
  type = "TEXT"
  value = "https://github.com/IBM/tekton-tutorial.git"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `default` - (Optional, String) Default option for SINGLE_SELECT property type.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
* `enum` - (Optional, List) Options for SINGLE_SELECT property type.
  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
* `name` - (Optional, String) Property name.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,234}$/`.
* `path` - (Optional, String) property path for INTEGRATION type properties.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.
* `pipeline_id` - (Required, Forces new resource, String) The tekton pipeline ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `trigger_id` - (Required, Forces new resource, String) The trigger ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `type` - (Optional, String) Property type.
  * Constraints: Allowable values are: `SECURE`, `TEXT`, `INTEGRATION`, `SINGLE_SELECT`, `APPCONFIG`.
* `value` - (Optional, String) String format property value.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the tekton_pipeline_trigger_property.

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

You can import the `ibm_cd_tekton_pipeline_trigger_property` resource by using `name`.
The `name` property can be formed from `pipeline_id`, `trigger_id`, and `property_name` in the following format:

```
<pipeline_id>/<trigger_id>/<property_name>
```
* `pipeline_id`: A string in the format `94619026-912b-4d92-8f51-6c74f0692d90`. The tekton pipeline ID.
* `trigger_id`: A string in the format `1bb892a1-2e04-4768-a369-b1159eace147`. The trigger ID.
* `property_name`: A string in the format `debug-pipeline`. The property's name.

# Syntax
```
$ terraform import ibm_cd_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property <pipeline_id>/<trigger_id>/<property_name>
```
