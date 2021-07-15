---

copyright:
  years: 2021
lastupdated: "2021"

keywords: terraform

subcollection: terraform

---

# IBM Cloud Shell resources
{: #ibm-cloud-shell-resources}

Update IBM Cloud Shell resources.
You can reference the output parameters for each resource in other resources or data sources by using Terraform interpolation syntax.

Before you start working with your resource, make sure to review the [required parameters](/docs/terraform?topic=terraform-provider-reference#required-parameters) 
that you need to specify in the `provider` block of your Terraform configuration file.
{: important}

## `ibm_cloud_shell_account_settings`
{: #cloud_shell_account_settings}

Update cloud_shell_account_settings.
{: shortdesc}

### Sample Terraform code
{: #cloud_shell_account_settings-sample}

```
resource "ibm_cloud_shell_account_settings" "cloud_shell_account_settings" {
  account_id = "12345678-abcd-1a2b-a1b2-1234567890ab"
  rev = "130-1bc9ec83d7b9b049890c6d4b74dddb2a"
  default_enable_new_features = true
  default_enable_new_regions = true
  enabled = true
  features {
  	enabled = true
  	key = "server.file_manager"
  }
  features {
  	enabled = true
  	key = "server.web_preview"
  }
  regions {
  	enabled = true
  	key = "eu-de"
  }
  regions {
  	enabled = true
  	key = "jp-tok"
  }
  regions {
  	enabled = true
  	key = "us-south"
  }
}
```

### Input parameters
{: #cloud_shell_account_settings-input}

Review the input parameters that you can specify for your resource. {: shortdesc}

|Name|Data type|Required/optional|Description|Forces new resource|
|----|-----------|-------|----------|--------------------|
|`account_id`|String|Required|The account ID in which the account settings belong to.|Yes|
|`rev`|String|Optional|Unique revision number for the settings object.|No|
|`default_enable_new_features`|Boolean|Optional|You can choose which Cloud Shell features are available in the account and whether any new features are enabled as they become available. The feature settings apply only to the enabled Cloud Shell locations.|No|
|`default_enable_new_regions`|Boolean|Optional|Set whether Cloud Shell is enabled in a specific location for the account. The location determines where user and session data are stored. By default, users are routed to the nearest available location.|No|
|`enabled`|Boolean|Optional|When enabled, Cloud Shell is available to all users in the account.|No|
|`features`|List|Optional|List of Cloud Shell features.|No|
|`regions`|List|Optional|List of Cloud Shell region settings.|No|

### Output parameters
{: #cloud_shell_account_settings-output}

Review the output parameters that you can access after your resource is updated. {: shortdesc}

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

### Import
{: #cloud_shell_account_settings-import}

`ibm_cloud_shell_account_settings` can be imported by ID

```
$ terraform import ibm_cloud_shell_account_settings.example sample-id
```

