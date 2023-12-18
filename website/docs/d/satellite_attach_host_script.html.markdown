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

###  Sample to read satellite host script to attach IBM host to Satellite control plane

```terraform
data "ibm_satellite_attach_host_script" "script" {
  location          = var.location
  labels            = ["cpu:4"]
  host_provider     = "ibm"
}
```

###  Sample to read satellite host script to attach AWS EC2 host to Satellite control plane and uses reduced firewall requirements.

```terraform
data "ibm_satellite_attach_host_script" "script" {
  location                 = var.location
  labels                   = var.labels
  script_dir               = "/tmp"
  host_provider            = "aws"
  host_link_agent_endpoint = "c-01-ws.us-south.link.satellite.cloud.ibm.com"
}
```
###  Sample to read satellite host script to attach IBM host to Satellite control plane

```terraform
data "ibm_satellite_attach_host_script" "script" {
  location      = var.location
  custom_script = <<EOF
subscription-manager refresh
subscription-manager release --set=8
subscription-manager repos --enable rhel-8-for-x86_64-baseos-rpms 
subscription-manager repos --enable rhel-8-for-x86_64-appstream-rpms
subscription-manager repos --disable='*eus*'
yum install container-selinux -y
EOF
}

```

## Argument reference
Review the argument references that you can specify for your data source.

- `coreos_host`	   = (Optional, Bool) True if attaching a CoreOS host to a CoreOS-enabled location. Host attach script will be in ignition file format. If attaching a RHEL host to a location, then the value is false.
- `custom_script` - (Optional, String) RHEL hosts only. The custom script that has to be appended to generated host script file. Either `custom_script` or `host_provider` is required. This `custom_script` will be appended to the downloaded host attach script. Find custom scripts for respective cloud providers [aws](https://cloud.ibm.com/docs/satellite?topic=satellite-aws#aws-host-attach), [google](https://cloud.ibm.com/docs/satellite?topic=satellite-gcp#gcp-host-attach), [azure](https://cloud.ibm.com/docs/satellite?topic=satellite-azure#azure-host-attach), [ibm](https://cloud.ibm.com/docs/satellite?topic=satellite-ibm#ibm-host-attach).
- `location` - (Required, String) The name or ID of the Satellite location.
- `host_provider` - (Optional, String) The name of host provider, such as `ibm`, `aws` or `azure`.
- `labels` - (Optional, Set(Strings)) The set of key-value pairs to label the host, such as `["cpu:4"]` to describe the host capabilities.
- `script_dir` - (Optional, String) The directory path to store the generated script.
- `host_link_agent_endpoint` - (Optional, String) The endpoint that the link agent uses to connect to the link tunnel server. Required for reduced firewall support.

## Attributes reference
In addition to the argument reference list, you can access the following attribute reference after your resource is created.

- `id` - The unique identifier of the location.
- `script_path` -  (String) Directory path to store the generated script.
- `host_script` -  (String) The raw content of the script file that was read.