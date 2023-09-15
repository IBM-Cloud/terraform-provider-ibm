---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_storage_configuration"
description: |-
  Manages IBM Cloud Satellite Storage Configuration.
---

# ibm_satellite_storage_configuration

Create, update, or delete [IBM Cloud Storage Configuration](https://cloud.ibm.com/docs/satellite?topic=satellite-storage-template-ov&interface=ui). By using storage templates, you can create storage configurations that can be consistently assigned, updated, and managed across the clusters, service clusters, and cluster groups in your location.

## Example usage

###  Sample to create a storage configuration by using the odf-remote storage template

```terraform
resource "ibm_satellite_storage_configuration" "storage_configuration" {
    location = "location-name"
    config_name = "config-name"
    storage_template_name = "odf-remote"
    storage_template_version = "4.12"
    user_config_parameters = {
        osd-size = "100Gi"
        osd-storage-class = "ibmc-vpc-block-metro-5iops-tier"
        billing-type = "advanced"
        cluster-encryption = "false"
        ibm-cos-endpoint = ""
        ibm-cos-location = ""
        ignore-noobaa = "false"
        kms-base-url = ""
        kms-encryption = "false"
        kms-instance-id = ""
        kms-instance-name = ""
        kms-token-url = ""
        num-of-osd = "1"
        odf-upgrade = "true"
        perform-cleanup = "false"
        worker-nodes = ""	
    }
    user_secret_parameters = {
        iam-api-key = "apikey"
    }
}
```
## Timeouts
The `ibm_satellite_storage_configuration` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 20 minutes) Used for creating Instance.
- **delete** - (Default 20 minutes) Used for deleting Instance.

## Argument reference
Review the argument references that you can specify for your resource. 

- `location` - (Required, String) The name of the location where the storage configuration is created.
- `config_name` - (Required, String) The name of the storage configuration to be created.
- `storage_template_name` - (Required, String) The storage template name to create the storage configuration such as `odf-remote`.
- `storage_template_version` - (Required, String) The version of the storage template you'd like to use.
- `user_config_parameters` - (Required, Map) The Storage Configuration parameters of a particular storage template to be passed to be passed as a Map of string key-value. Both the key-value i.e the configuration parameter name and the respective value must be entered based on the chosen storage template.
- `user_secret_parameters` - (Required, Map) The Storage Configuration secret parameters of a particular storage template to be passed as a Map of string key-value. Both the key-value i.e the secret parameter name and the respective value must be entered based on the chosen storage template.
- `storage_class_parameters` - (Optional, List(Map)) A list of Maps, users can enter custom storage classes if the storage template supports it. Each Map will require a key-value of type string.
- `update_assignments` - (Optional, String) If set to `true` it will auto-update all the configuration's assignments with the latest revision after a configuration update.
- `delete_assignments` - (Optional, String) If set to `true` it will auto-delete all the configuration's assignments before the configuration's deletion.

* To find out more about the different parameters and storage classes check the [available storage templates](https://cloud.ibm.com/docs/satellite?topic=satellite-storage-template-ov&interface=ui#storage-template-ov-providers) to know more.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `uuid` - (String) The Universally Unique IDentifier (UUID) of the Storage Configuration.
- `config_version` - (String) The current version of the storage configuration.
- `id` - (String) The ID of the storage configuration resource.