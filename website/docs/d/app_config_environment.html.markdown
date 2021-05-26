---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration environment'
description: |-
  Get information about environment
---

# ibm_app_config_environment

Provides a read-only data source for `environment`. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_app_config_environment" "app_config_environment" {
	guid = "guid"
	expand = "expand"
	environment_id = "environment_id"
}
```

## Argument Reference

The following arguments are supported:

- `guid` - (Required, string) guid of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `environment_id` - (Required, string) Environment Id.
- `expand` - (optional, bool) If set to `true`, returns expanded view of the resource details.

## Attribute Reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the app-config-environment.
- `name` - Environment name.
- `description` - Environment description.
- `tags` - Tags associated with the environment.
- `color_code` - Color code to distinguish the environment. The Hex code for the color. For example `#FF0000` for `red`.
- `created_time` - Creation time of the environment.
- `updated_time` - Last modified time of the environment data.
- `href` - Environment URL.
