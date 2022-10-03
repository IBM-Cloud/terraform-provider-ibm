---
layout: "ibm"
page_title: "IBM : ibm_cm_offering_instance"
description: |-
  Manages cm_offering_instance.
subcategory: "Catalog Management API"
---

# ibm_cm_offering_instance

Provides a resource for cm_offering_instance. This allows cm_offering_instance to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cm_offering_instance" "cm_offering_instance" {
  last_operation {
		operation = "operation"
		state = "state"
		message = "message"
		transaction_id = "transaction_id"
		updated = "2021-01-31T09:44:12Z"
		code = "code"
  }
  x_auth_refresh_token = "x_auth_refresh_token"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `account` - (Optional, String) The account this instance is owned by.
* `catalog_id` - (Optional, String) Catalog ID this instance was created from.
* `channel` - (Optional, String) Channel to pin the operator subscription to.
* `cluster_all_namespaces` - (Optional, Boolean) designate to install into all namespaces.
* `cluster_id` - (Optional, String) Cluster ID.
* `cluster_namespaces` - (Optional, List) List of target namespaces to install into.
* `cluster_region` - (Optional, String) Cluster region (e.g., us-south).
* `created` - (Optional, String) date and time create.
* `crn` - (Optional, String) platform CRN for this instance.
* `disabled` - (Optional, Boolean) Indicates if Resource Controller has disabled this instance.
* `install_plan` - (Optional, String) Type of install plan (also known as approval strategy) for operator subscriptions. Can be either automatic, which automatically upgrades operators to the latest in a channel, or manual, which requires approval on the cluster.
* `kind_format` - (Optional, String) the format this instance has (helm, operator, ova...).
* `kind_target` - (Optional, String) The target kind for the installed software version.
* `label` - (Optional, String) the label for this instance.
* `last_operation` - (Optional, List) the last operation performed and status.
Nested scheme for **last_operation**:
	* `code` - (Optional, String) Error code from the last operation, if applicable.
	* `message` - (Optional, String) additional information about the last operation.
	* `operation` - (Optional, String) last operation performed.
	* `state` - (Optional, String) state after the last operation performed.
	* `transaction_id` - (Optional, String) transaction id from the last operation.
	* `updated` - (Optional, String) Date and time last updated.
* `location` - (Optional, String) String location of OfferingInstance deployment.
* `metadata` - (Optional, Map) Map of metadata values for this offering instance.
* `offering_id` - (Optional, String) Offering ID this instance was created from.
* `resource_group_id` - (Optional, String) Id of the resource group to provision the offering instance into.
* `rev` - (Optional, String) Cloudant revision.
* `schematics_workspace_id` - (Optional, String) Id of the schematics workspace, for offering instances provisioned through schematics.
* `sha` - (Optional, String) The digest value of the installed software version.
* `updated` - (Optional, String) date and time updated.
* `url` - (Optional, String) url reference to this object.
* `version` - (Optional, String) The version this instance was installed from (semver - not version id).
* `version_id` - (Optional, String) The version id this instance was installed from (version id - not semver).
* `x_auth_refresh_token` - (Required, String) IAM Refresh token.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cm_offering_instance.

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

You can import the `ibm_cm_offering_instance` resource by using `id`. provisioned instance ID (part of the CRN).

# Syntax
```
$ terraform import ibm_cm_offering_instance.cm_offering_instance <id>
```
