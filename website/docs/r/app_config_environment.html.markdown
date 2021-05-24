---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration environments'
description: |-
  Manages environments.
---

# ibm_app_config_environment

Provides a resource for `environment`. This allows environment to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_app_config_environment" "app_config_environment" {
  guid = "guid"
  environment_id = "environment_id"
  name = "name"
  description = "description"
  tags = "tags"
  color_code = "color_code"
}
```

## Argument Reference

The following arguments are supported:

- `guid` - (Required, string) guid of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `name` - (Required, string) Environment name.
- `environment_id` - (Required, string) Environment id.
- `description` - (Optional, string) Environment description.
- `tags` - (Optional, string) Tags associated with the environment.
- `color_code` - (Optional, string) Color code to distinguish the environment. The Hex code for the color. For example `#FF0000` for `red`.

## Attribute Reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the environment resource.
- `created_time` - Creation time of the environment.
- `updated_time` - Last modified time of the environment data.
- `href` - Environment URL.

## Import

The `ibm_app_config_environment` resource can be imported by using `guid` of the App Configuration instance and `environmentId`. Get `guid` from the service instance credentials section of the dashboard.

**Syntax**

```
terraform import ibm_app_config_environment.sample  <guid/environmentId>

```

**Example**

```
terraform import ibm_app_config_environment.sample 272111153-c118-4116-8116-b811fbc31132/dev
```
