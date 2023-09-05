---
layout: "ibm"
page_title: "IBM : ibm_cd_tekton_pipeline_trigger_property"
description: |-
  Manages cd_tekton_pipeline_trigger_property.
subcategory: "Continuous Delivery"
---

# ibm_cd_tekton_pipeline_trigger_property

Create, update, and delete cd_tekton_pipeline_trigger_propertys with this resource.

## Example Usage

```hcl
resource "ibm_cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property_instance" {
  name = "prop1"
  pipeline_id = "94619026-912b-4d92-8f51-6c74f0692d90"
  trigger_id = "1bb892a1-2e04-4768-a369-b1159eace147"
  type = "text"
  value = "https://github.com/open-toolchain/hello-tekton.git"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `enum` - (Optional, List) Options for `single_select` property type. Only needed for `single_select` property type.
  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`. The maximum length is `256` items. The minimum length is `0` items.
* `name` - (Required, Forces new resource, String) Property name.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
* `path` - (Optional, String) A dot notation path for `integration` type properties only, that selects a value from the tool integration. If left blank the full tool integration data will be used.
  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/^[-0-9a-zA-Z_.]*$/`.
* `pipeline_id` - (Required, Forces new resource, String) The Tekton pipeline ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `trigger_id` - (Required, Forces new resource, String) The trigger ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `type` - (Required, Forces new resource, String) Property type.
  * Constraints: Allowable values are: `secure`, `text`, `integration`, `single_select`, `appconfig`.
* `value` - (Optional, String) Property value. Any string value is valid.
  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the cd_tekton_pipeline_trigger_property.
* `href` - (String) API URL for interacting with the trigger property.
  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

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
* `pipeline_id`: A string in the format `94619026-912b-4d92-8f51-6c74f0692d90`. The Tekton pipeline ID.
* `trigger_id`: A string in the format `1bb892a1-2e04-4768-a369-b1159eace147`. The trigger ID.
* `property_name`: A string in the format `debug-pipeline`. The property name.

# Syntax
```
$ terraform import ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property <pipeline_id>/<trigger_id>/<property_name>
```
