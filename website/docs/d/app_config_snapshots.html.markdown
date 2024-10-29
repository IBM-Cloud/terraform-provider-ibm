---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration Snapshots'
description: |-
  Get information about Snapshots
---

# ibm_app_config_segment
Retrieve information about an existing IBM Cloud App Configuration snapshots. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration snapshot, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_snapshots" "app_config_snapshots" {
  guid = "guid"
  region = "region"
  sort = "sort"
  collection_id = "collection_id"
  environment_id = "environment_id"
  limit = "limit"
  offset = "offset"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) The GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `region` - (Required, String)The region of the App Configuration Instance
- `collection_id` - (Optional, String) Filters the response based on the specified collection_id.
- `environment_id` - (Optional, String) Filters the response based on the specified environment_id.
- `limit` - (Optional, String) The number of records to retrieve. By default, the list operation return the first 10 records.
- `offset` - (Optional, String) The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value.


## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `limit`  - (Integer) Number of records returned.
- `offset` - (Integer) Skipped number of records.
- `total_count` - (Integer) Total number of records

- `first` - (List) The URL to navigate to the first page of records.

  Nested scheme for `first`:
    - `href` - (String) The first `href` URL.

- `last` - (List) The URL to navigate to the last page of records.

  Nested scheme for `last`:
    - `href` - (String) The last `href` URL.

- `previous` - (List) The URL to navigate to the previous list of records.

  Nested scheme for `previous`:
    - `href` - (String) The previous `href` URL.

- `next` - (List) The URL to navigate to the next list of records

  Nested scheme for `next`:
    - `href` - (String) The next `href` URL.


- `git_config` - Array of git_configs.

    Nested scheme for `git_config`:
  - `git_config_name` - (String) Git config name. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
  - `git_config_id` - (String) Git config id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only
  - `git_url`  - (String) Git url which will be used to connect to the github account. The url must be formed in this format, https://api.github.com/repos/{owner}/{repo_name} for the personal git account.
  - `git_branch`  - (String) Branch name to which you need to write or update the configuration.
  - `git_file_path`  - (String) Git file path, this is a path where your configuration file will be written. The path must contain the file name with `json` extension.
  - `created_time` - (Timestamp) Creation time of the git config.
  - `updated_time` - (Timestamp) Last modified time of the git config data.
  - `last_sync_time` - (Timestamp) Latest time when the snapshot was synced to git.
  - `href` - (String) Git config URL.

  - `collection` - (Object) Collection object will be returned containing attributes collection_id, collection_name.

    Nested scheme for `collection`:
    - `collection_id`  - (String) Collection id.
    - `name`  - (String) Name of the collection.

  - `environment`  - (Object) Environment object will be returned containing attributes environment_id, environment_name, color_code.

    Nested scheme for `environment` :
    - `environment_id`  - (String)  Environment Id.
    - `environment_name` - (String) Environment name. 
    - `color_code` - (String) Color code to distinguish the environment.
  


    