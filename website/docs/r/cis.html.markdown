---
layout: "ibm"
page_title: "IBM : Cloud Internet Services instance"
sidebar_current: "docs-ibm-resource-cis"
description: |-
  Manages IBM Cloud Internet Services Instance.
---

# ibm\_cis

Provides a Cloud Internet Services instance resource. This allows CIS instances to be created, updated, and deleted. The Bluemix_API_KEY used by Terraform must have been granted sufficient rights to create IBM Cloud Resources and have access to the Resource Group the CIS instance will be associated with. 

If no resource_group_id is specified, the CIS instance is created under the default resource group. The API_KEY must have been assigned permissions for this group.  

## Example Usage

```hcl
data "ibm_resource_group" "group" {
  name = "test"
}

resource "ibm_cis" "cis_instance" {
  name              = "test"
  plan              = "standard"
  resource_group_id = "${data.ibm_resource_group.group.id}"
  tags              = ["tag1", "tag2"]

  //User can increase timeouts 
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}
```

## Timeouts

ibm_cis provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for Creating Instance.
* `update` - (Default 10 minutes) Used for Updating Instance.
* `delete` - (Default 10 minutes) Used for Deleting Instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required, string) A descriptive name used to identify the CIS instance.
* `plan` - (Required, string) The name of the plan type for Cloud Internet Services. You can retrieve the value by running the `ibmcloud catalog service internet-svcs` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `resource_group_id` - (Optional, string) The ID of the resource group where you want to create the service. You can retrieve the value from data source `ibm_resource_group`. If not provided creates the service in default resource group.
* `tags` - (Optional, array of strings) Tags associated with the instance.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the new CIS instance.
* `status` - Status of resource instance.


## Import

The `cis` resource can be imported using the `id`. The ID is formed from the `CRN` (Cloud Resource Name) from the **Overview** page of the Internet Services instance. These will be located under the **Domain** heading. 
* CRN is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

```
$ terraform import cis_domain.myorg <crn>

$ terraform import cis_domain.myorg crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
