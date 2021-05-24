---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration environments'
description: |-
  List all the environments.
---

# ibm_app_config_environments

Provides a read-only data source for `environments`. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_app_config_environments" "app_config_environments" {
  guid = "guid"
  tags = "tags"
  expand = "expand"
  limit = "limit"
  offset = "limit"
}
```

## Argument Reference

The following arguments are supported:

- `guid` - (Required, string) guid of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `tags` - (optional, string) Flter the resources to be returned based on the associated tags. Returns resources associated with any of the specified tags.
- `expand` - (optional, bool) If set to `true`, returns expanded view of the resource details.
- `limit` - (optional, int) The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different set of records, use `limit` with `offset` to page through the available records.
- `offset` - (optional, int) The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. Use `offset` with `limit` to page through the available records.

## Attribute Reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the EnvironmentList.
- `environments` - Array of environments. Nested `environments` blocks have the following structure:

  - `name` - Environment name.

  - `environment_id` - Environment id.

  - `description` - Environment description.

  - `tags` - Tags associated with the environment.

  - `color_code` - Color code to distinguish the environment. The Hex code for the color. For example `#FF2200` for `red`.

  - `created_time` - Creation time of the environment.

  - `updated_time` - Last modified time of the environment data.

  - `href` - Environment URL.

- `total_count` - Number of records returned in the current response.

- `first` - URL to navigate to the first page of records. Nested `first` blocks have the following structure:

  - `href` - URL.

- `previous` - URL to navigate to the previous list of records. Nested `previous` blocks have the following structure:

  - `href` - URL.

- `last` - URL to navigate to the last list of records. Nested `last` blocks have the following structure:

  - `href` - URL.

- `next` - URL to navigate to the next list of records.. Nested `next` blocks have the following structure:
  - `href` - URL.
