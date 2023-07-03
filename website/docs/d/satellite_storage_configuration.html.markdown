---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_storage_configuration"
description: |-
  Get information about an IBM Cloud Satellite Storage Configuration.
---

# ibm_satellite_storage_configuration
Retrieve information of an existing Satellite Storage Configuration. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax. For more information, about IBM Cloud Satellite Storage Configurations see [Satellite Storage](https://cloud.ibm.com/docs/satellite?topic=satellite-storage-template-ov&interface=ui).


## Example usage

```terraform
data "ibm_satellite_storage_configuration" "storage_configuration" {
  config_name = var.config_name
  location  = var.location
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `location` - (Required, String) The name of the location where the storage configuration is created.
- `config_name` - (Required, String) The name of the storage configuration.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `storage_template_name` - (String) The storage template name used by the storage configuration.
- `storage_template_version` - (String) The version of the storage template.
- `user_config_parameters` - (Map) The Storage Configuration parameters of a particular storage template in the form of a Map.
- `storage_class_parameters` - (List(Map)) A list of the different storage classes available to the storage configuration, each in the form of a map.
- `uuid` - (String) The Universally Unique IDentifier (UUID) of the Storage Configuration.
- `config_version` - (String) The current version of the storage configuration.
- `id` - (String) The ID of the storage configuration data source.