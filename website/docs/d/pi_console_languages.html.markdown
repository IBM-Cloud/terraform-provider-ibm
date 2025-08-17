---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_console_languages"
description: |-
  Manages Instance Console Languages in the Power Virtual Server cloud.
---

# ibm_pi_console_languages

Retrieve information about all the available Console Languages for an Instance. For more information, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage

```terraform
data "ibm_pi_console_languages" "example" {
  pi_cloud_instance_id  = "<value of the cloud_instance_id>"
  pi_instance_name      = "<instance name or id>"
}
```

### Notes

- Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
- If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  - `region` - `lon`
  - `zone` - `lon04`

Example usage:

  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```
  
## Argument Reference

Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_instance_name` - (Required, String) The unique identifier or name of the instance.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `console_languages` - (List) List of all the Console Languages.

  Nested scheme for `console_languages`:
  - `code` - (String) Language code.
  - `language` - (String) Language description.
