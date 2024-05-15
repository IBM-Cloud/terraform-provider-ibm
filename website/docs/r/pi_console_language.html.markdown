---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_console_language"
description: |-
  Manages Instance Console Languages in the Power Virtual Server cloud.
---

# ibm_pi_console_language

Update the Console Language of a Server for your Power Systems Virtual Server instance. For more information, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage

The following example enables you to update the Console Languages of a Server:

```terraform
resource "ibm_pi_console_language" "example" {
  pi_cloud_instance_id  = "<value of the cloud_instance_id>"
  pi_instance_name      = "<instance name or id>"
  pi_language_code      = "<language code>"
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

## Timeouts

The `ibm_pi_console_language` provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 5 minutes) Used for setting a console language.
- **update** - (Default 5 minutes) Used for updating a console language.

## Argument reference

Review the argument references that you can specify for your resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_instance_name` - (Required, String) The unique identifier or name of the instance.
- `pi_language_code` - (Required, String) Language code.

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the instance console language. The ID is composed of `<pi_cloud_instance_id>/<pi_instance_name>`.
