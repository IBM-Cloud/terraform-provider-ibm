---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_attach_host_script"
description: |-
  Generate host script to attach host to Satellite location.
---

# ibm_satellite_attach_host_script
Retrieve information of an existing IBM Satellite location registration script as a data source. Creates a script to run on a Red Hat Enterprise Linux 7 or AWS EC2 host in your on-premises infrastructure. The script attaches the host to your IBM Cloud Satellite location. The host must have access to the public network in order for the script to complete. For more information, about setting up Satellite hosts, see [Satellite hosts](https://cloud.ibm.com/docs/satellite?topic=satellite-hosts).

## Example usage

###  Sample to create satellite host script to attach IBM host to Satellite control plane

```terraform
data "ibm_satellite_attach_host_script" "script" {
  location          = var.location
  labels            = var.labels
  host_provider     = "ibm"
}
```

###  Sample to create satellite host script to attach AWS EC2 host to Satellite control plane

```terraform
data "ibm_satellite_attach_host_script" "script" {
  location          = var.location
  labels            = var.labels
  script_dir        = "/tmp"
  host_provider     = "aws"
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `location` - (Required, String) The name or ID of the Satellite location.

## Attributes reference
In addition to the argument reference list, you can access the following attribute reference after your resource is created.

- `id` - The unique identifier of the location.
- `labels` - (Strings) The key-value pairs to label the host, such as `cpu=4` to describe the host capabilities.
- `script_dir` - (String) The directory path to store the generated script.
- `host_provider` - (String) The name of host provider, such as `ibm`, `aws` or `azure`.
- `script_path` -  (String) Directory path to store the generated script.
- `host_script` -  (String) The raw content of the script file that was read.

