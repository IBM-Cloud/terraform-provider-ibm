---
layout: "ibm"
page_title: "IBM : ibm_metrics_router_settings"
description: |-
  Manages metrics_router_settings
subcategory: "IBM Cloud Metrics Routing"
---

# ibm_metrics_router_settings

Provides a resource for metrics_router_settings. This allows metrics_router_settings to be updated.

## Example Usage

### Example with no default_targets

```hcl
resource "ibm_metrics_router_settings" "metrics_router_settings_instance" {
  permitted_target_regions = ["us-south"]
  primary_metadata_region = "us-south"
  backup_metadata_region = "us-east"
  private_api_endpoint_only = false
}
```

### Example with single default_targets

```hcl
resource "ibm_metrics_router_settings" "metrics_router_settings_instance" {
  default_targets {
		id = "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"
  }
  permitted_target_regions = ["us-south"]
  primary_metadata_region = "us-south"
  backup_metadata_region = "us-east"
  private_api_endpoint_only = false
}
```

### Example with multiple default_targets

```hcl
resource "ibm_metrics_router_settings" "metrics_router_settings_instance" {
  default_targets {
		id = "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"
  }
  default_targets {
		id = "c3af557f-fb0e-2222-85c3-0889e7fe7bc4"
  }
  permitted_target_regions = ["us-south", "us-east"]
  primary_metadata_region = "us-south"
  backup_metadata_region = "us-east"
  private_api_endpoint_only = false
}
```

### Example without permitted_target_regions

```hcl
resource "ibm_metrics_router_settings" "metrics_router_settings_instance" {
  permitted_target_regions = []
  primary_metadata_region = "us-south"
  backup_metadata_region = "us-east"
  private_api_endpoint_only = false
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `backup_metadata_region` - (Optional, String) To backup all your meta data in a different region.
  * Constraints: The maximum length is `256` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 \\-_]+$/`.
* `default_targets` - (Optional, List) A list of default target references.
  * Constraints: The maximum length is `2` items. The minimum length is `0` items.
Nested scheme for **default_targets**:
	* `id` - (Required, String) The target uuid for a pre-defined metrics router target.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 \\-._:]+$/`.
* `permitted_target_regions` - (Optional, List) If present then only these regions may be used to define a target.
  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9 \\-_]+$/`. The maximum length is `16` items. The minimum length is `0` items.
* `primary_metadata_region` - (Optional, String) To store all your meta data in a single region. For new accounts, all target / route creation will fail until primary_metadata_region is set.
  * Constraints: The maximum length is `256` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 \\-_]+$/`.
  * Note: In case of _ibm_metrics_router_settings_, we don't use a traditional ID to track it since its a resource without one. However, terraform expects there will be one. To work around this, we set the primary_metadata_region to be used as the ID.
* `private_api_endpoint_only` - (Optional, Boolean) If you set this true then you cannot access api through public network.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the metrics_router_settings.

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

You can import the `ibm_metrics_router_settings` resource by using `primary_metadata_region`. To store all your meta data in a single region.

# Syntax
```
$ terraform import ibm_metrics_router_settings.metrics_router_settings <primary_metadata_region>
```

# Example
```
$ terraform import ibm_metrics_router_settings.metrics_router_settings us-south
```
