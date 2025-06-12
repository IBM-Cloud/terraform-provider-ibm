---
layout: "ibm"
page_title: "IBM : ibm_metrics_router_route"
description: |-
  Manages metrics_router_route
subcategory: "IBM Cloud Metrics Routing"
---

# ibm_metrics_router_route

Provides a resource for metrics_router_route. This allows metrics_router_route to be created, updated and deleted.

## Example Usage

### Example for a single rule and single inclusion_filter

```hcl
resource "ibm_metrics_router_route" "metrics_router_route_instance" {
	name = "my-route"
	rules {
		action = "send"
		targets {
			id = "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"
		}
		inclusion_filters {
			operand = "location"
			operator = "is"
			values = [ "us-south" ]
		}
  	}
}
```

### Example for multiple rules and multiple inclusion_filters

```hcl
resource "ibm_metrics_router_route" "metrics_router_route_instance" {
	name = "my-route"
	rules {
		action = "send"
		targets {
			id = "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"
		}
		inclusion_filters {
			operand = "location"
			operator = "is"
			values = [ "us-south" ]
		}
		inclusion_filters {
			operand = "service_name"
			operator = "in"
			values = ["metrics-router"]
		}
  	}
	rules {
		action = "send"
		targets {
			id = "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"
		}
		inclusion_filters {
			operand = "resource_type"
			operator = "is"
			values = ["worker"]
		}
  	}
}
```

### Example of route with drop action

```hcl
resource "ibm_metrics_router_route" "metrics_router_route_instance" {
	name = "my-route"
	rules {
		action = "drop"
		inclusion_filters {
			operand = "location"
			operator = "is"
			values = [ "us-south" ]
		}
	}
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `name` - (Required, String) The name of the route. The name must be 1000 characters or less and cannot include any special characters other than `(space) - . _ :`. Do not include any personal identifying information (PII) in any resource names.
  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 \\-._:]+$/`.
* `rules` - (Required, List) Routing rules that will be evaluated in their order of the array.
  * Constraints: The maximum length is `4` items. The minimum length is `0` items.
Nested scheme for **rules**:
	* `action` - (Optional, String) The action if the inclusion_filters matches, default is `send` action.
	  * Constraints: Allowable values are: `send`, `drop`.
	* `inclusion_filters` - (Required, List) A list of conditions to be satisfied for routing metrics to pre-defined target.
	  * Constraints: The maximum length is `7` items. The minimum length is `0` items.
	Nested scheme for **inclusion_filters**:
		* `operand` - (Required, String) Part of CRN that can be compared with values.
		  * Constraints: Allowable values are: `location`, `service_name`, `service_instance`, `resource_type`, `resource`.
		* `operator` - (Required, String) The operation to be performed between operand and the provided values. 'is' to be used with one value and 'in' can support upto 20 values in the array.
		  * Constraints: Allowable values are: `is`, `in`.
		* `values` - (Required, List) The provided string values of the operand to be compared with.
		  * Constraints: The maximum length is `20` items. The minimum length is `1` item.
	* `targets` - (Required, List) A collection of targets with ID in the request.
	  * Constraints: The maximum length is `3` items. The minimum length is `0` items.
	Nested scheme for **targets**:
		* `id` - (Required, String) The target uuid for a pre-defined metrics router target.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 \\-._:]+$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the metrics_router_route.
* `created_at` - (String) The timestamp of the route creation time.
* `crn` - (String) The crn of the route resource.
* `updated_at` - (String) The timestamp of the route last updated time.

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

You can import the `ibm_metrics_router_route` resource by using `id`. The UUID of the route resource.

# Syntax
```
$ terraform import ibm_metrics_router_route.metrics_router_route <id>
```

# Example
```
$ terraform import ibm_metrics_router_route.metrics_router_route c3af557f-fb0e-4476-85c3-0889e7fe7bc4
```
