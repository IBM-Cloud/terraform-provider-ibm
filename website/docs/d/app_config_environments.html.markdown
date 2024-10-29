---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration environments'
description: |-
  List all the environments.
---

# ibm_app_config_environments

Retrieve information about an existing IBM Cloud App Configuration environments. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration environments, see [Managing access levels for App Configuration environments](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-service-access-level-management).

## Example usage

```terraform
data "ibm_app_config_environments" "app_config_environments" {
  guid = "guid"
  tags = "tags"
  expand = "expand"
  limit = "limit"
  offset = "limit"
}
```

## Argument reference

The following arguments are supported:

- `guid` - (Required, String) The GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.

- `tags` - (Optional, String) Filter the resources to be returned based on the associated tags. Returns resources associated with any of the specified tags.
- `expand` - (Optional, Bool) If set to `true`, returns expanded view of the resource details.
- `limit` - (Optional, Integer) The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different set of records, use `limit` with `offset` to page through the available records.
- `offset` - (Optional, Integer) The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. Use `offset` with `limit` to page through the available records.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the Environment List.
- `environments` - (List) Array of environments. Nested `environments` blocks have the following structure:

  Nested scheme for `environments`:
  - `name` - (String) Environment name.
  - `environment_id` - (String) Environment ID.
  - `description` - (String) The environment description.
  - `tags` - (String) The tags associated with the environment.
  - `color_code` - (String) The color code to distinguish the environment. The Hexadecimal code for the color. For example `#FF2200` for `red`.
  - `created_time` - (Timestamp) The creation time of the environment.
  - `updated_time` - (Timestamp) The last modified time of the environment data.
  - `href` - (String) Environment URL.

- `total_count` - Number of records returned in the current response.
- `first` - (List) URL to navigate to the first page of records.

  Nested scheme for `first`:
  - `href` - (String) The first `href` URL.
- `previous` - (List) URL to navigate to the previous list of records.

  Nested scheme for `previous`:
  - `href` - (String) The previous `href` URL.
- `last` - (List) URL to navigate to the last list of records.

  Nested scheme for `last`:
  - `href` - (String) The last `href` URL.
- `next` - (List) URL to navigate to the next list of records.

  Nested scheme for `first`:
  - `href` - (String) The next `href` URL.
