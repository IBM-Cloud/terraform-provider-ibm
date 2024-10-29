---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration Snapshots'
description: |-
  Manages snapshots.
---

# ibm_app_config_snapshots

Provides a resource for `snapshot`. This allows snapshot to be created, updated and deleted. For more information, about App Configuration snapshots, see [segments](https://cloud.ibm.com/docs/app-configuration?topic=app-configuration-ac-snapshots).

## Example usage

```terraform
resource "ibm_app_config_snapshot" "app_config_snapshot" {
  guid = "guid"
  region = "region"
  collection_id = "collection_id"
  environment_id = "environment_id"
  git_config_id = "git_config_id"
  git_config_name = "git_config_name"
  git_url = "git_url"
  git_branch = "git_branch"
  git_file_path = "git_file_path"
  git_token = "git_token"
}
```

## Argument reference

Review the argument reference that you can specify for your resource. 

- `guid` - (Required, String) The GUID of the App Configuration service. Fetch GUID from the service instance credentials section of the dashboard.
- `collection_id`  - (Required, String) Collection ID
- `environment_id` - (Required, String) Environment Id
- `git_config_name` - (Required, String) Git config name. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
- `git_config_id` - (Required, String) Git config id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only
- `git_url`  - (Required, String) Git url which will be used to connect to the github account. The url must be formed in this format, https://api.github.com/repos/{owner}/{repo_name} for the personal git account.
- `git_branch`  - (Required, String) Branch name to which you need to write or update the configuration.
- `git_file_path`  - (Required, String) Git file path, this is a path where your configuration file will be written. The path must contain the file name with `json` extension.
- `git_token`  - (Required, String) Git token, this needs to be provided with enough permission to write and update the file.


## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `created_time` - (Timestamp) Creation time of the segment.
- `updated_time` - (Timestamp) Last modified time of the segment data.
- `href` - (String) Git config URL.


## Import

The `ibm_app_config_snapshot` resource can be imported by using `guid` of the App Configuration instance and `git_config_id`. Get the `guid` from the service instance credentials section of the dashboard.

**Syntax**

```
terraform import ibm_app_config_snapshot.sample  <guid/git_config_id>

```

**Example**

```
terraform import ibm_app_config_snapshot.app_config_snapshot 272111153-c118-4116-8116-b811fbc31132/sample
```