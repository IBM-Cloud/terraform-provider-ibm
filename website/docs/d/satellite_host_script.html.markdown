---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_attach_host_script"
description: |-
  Generate host script to attach host to satellite location.
---

# ibm\_satellite_host

Import the details of an existing IBM satellite location registration script as a data source. Creates a script to run on a Red Hat Enterprise Linux 7 or AWS EC2 host in your on-premises infrastructure. The script attaches the host to your IBM Cloud Satellite location. The host must have access to the public network in order for the script to complete.

## Example Usage

###  Create satellite host script to attach IBM host to satellite control plane

```hcl
data "ibm_satellite_attach_host_script" "script" {
  location          = var.location
  labels            = var.labels
  host_provider     = "ibm"
}
```

###  Create satellite host script to attach AWS EC2 host to satellite control plane

```hcl
data "ibm_satellite_attach_host_script" "script" {
  location          = var.location
  labels            = var.labels
  script_dir        = "/tmp"
  host_provider     = "aws"
}
```

## Argument Reference

The following arguments are supported:

* `location` - (Required, string) The name or ID of the Satellite location.

## Attributes Reference

The following attributes are exported:

* `id` - The unique identifier of the location.
* `labels` - (Optional, array of strings) Key-value pairs to label the host, such as cpu=4 to describe the host capabilities.
* `script_dir` - (Optional, string) Directory path to store the generated script.
* `host_provider` - (Required, string) The name of host provider, such as ibm, aws or azure.
* `script_path` -  Directory path to store the generated script.
* `host_script` -  The raw content of the script file that was read.

