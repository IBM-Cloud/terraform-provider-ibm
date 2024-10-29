---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration environment'
description: |-
  Get information about environment
---

# ibm_app_config_environment

Retrieve information about an existing IBM Cloud App Configuration environment. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration environment, see [Managing access levels for App Configuration environments](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-service-access-level-management).

## Example usage

```terraform
data "ibm_app_config_environment" "app_config_environment" {
	guid = "guid"
    region = "region"
	expand = "expand"
	environment_id = "environment_id"
}
```

## Argument reference

The following arguments are supported:

- `guid` - (Required, String) The GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.

- `environment_id` - (Required, String) Environment ID.
- `expand` - (optional, Bool) If set to `true`, returns expanded view of the resource details.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the app-config-environment.
- `name` - (String) Environment name.
- `description` - (String) Environment description.
- `tags` - (String) Tags associated with the environment.
- `color_code` - (String) The color code to distinguish the environment. The Hexadecimal code for the color. For example `#FF0000` for `red`.
- `created_time` - (Timestamp) Creation time of the environment.
- `updated_time` - (Timestamp) Last modified time of the environment data.
- `href` - (String) Environment URL.
