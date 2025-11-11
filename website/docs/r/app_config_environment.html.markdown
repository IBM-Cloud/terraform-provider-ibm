---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration environments'
description: |-
  Manages environments.
---

# ibm_app_config_environment

Create, update, or delete an environment by using IBM Cloud™ App Configuration. For more information, about App Configuration, see [Getting started with App Configuration](https://cloud.ibm.com/docs/app-configuration?topic=app-configuration-getting-started).

## Example usage

```terraform
resource "ibm_app_config_environment" "app_config_environment" {
  guid = "guid"
  environment_id = "environment_id"
  name = "name"
  description = "description"
  tags = "tag1,tag2"
  color_code = "color_code"
}
```

## Argument reference

Review the argument reference that you can specify for your resource. 

- `guid` - (Required, String) The GUID of the App Configuration service. Fetch GUID from the service instance credentials section of the dashboard.
- `name` - (Required, String) The environment name.
- `environment_id` - (Required, String) The environment ID.
- `description` - (Optional, String) The environment description.
- `tags` - (Optional, String) The tags associated with an environment.
- `color_code` - (Optional, String) The color code to distinguish an environment in the Hexademical code format. For example, `#FF0000` for `red`.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of an environment resource.
- `created_time` - (Timestamp) the creation time of an environment.
- `updated_time` - (Timestamp) the last modified time of an environment data.
- `href` - (String) the environment URL.

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
