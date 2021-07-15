---

copyright:
  years: 2021
lastupdated: "2021"

keywords: terraform

subcollection: terraform

---

# IBM Cloud Shell data sources
{: #ibm-cloud-shell-data-sources}

Review the data sources that you can use to retrieve information about your IBM Cloud Shell resources.
All data sources are imported as read-only information. You can reference the output parameters for each data source by using Terraform interpolation syntax.

Before you start working with your data source, make sure to review the [required parameters](/docs/terraform?topic=terraform-provider-reference#required-parameters)
that you need to specify in the `provider` block of your Terraform configuration file.
{: important}

## `ibm_cloud_shell_account_settings`
{: #cloud_shell_account_settings}

Retrieve information about cloud_shell_account_settings.
{: shortdesc}

### Sample Terraform code
{: #cloud_shell_account_settings-sample}

```
data "ibm_cloud_shell_account_settings" "cloud_shell_account_settings" {
  account_id = "account_id"
}
```

### Input parameters
{: #cloud_shell_account_settings-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`account_id`|String|Required|The account ID in which the account settings belong to.|

### Output parameters
{: #cloud_shell_account_settings-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|Unique id of the settings object.|
|`rev`|String|Unique revision number for the settings object.|
|`created_at`|Integer|Creation timestamp in Unix epoch time.|
|`created_by`|String|IAM ID of creator.|
|`default_enable_new_features`|Boolean|You can choose which Cloud Shell features are available in the account and whether any new features are enabled as they become available. The feature settings apply only to the enabled Cloud Shell locations.|
|`default_enable_new_regions`|Boolean|Set whether Cloud Shell is enabled in a specific location for the account. The location determines where user and session data are stored. By default, users are routed to the nearest available location.|
|`enabled`|Boolean|When enabled, Cloud Shell is available to all users in the account.|
|`features`|List|List of Cloud Shell features.|
|`features.enabled`|Boolean|State of the feature.|
|`features.key`|String|Name of the feature.|
|`regions`|List|List of Cloud Shell region settings.|
|`regions.enabled`|Boolean|State of the region.|
|`regions.key`|String|Name of the region.|
|`type`|String|Type of api response object.|
|`updated_at`|Integer|Timestamp of last update in Unix epoch time.|
|`updated_by`|String|IAM ID of last updater.|

