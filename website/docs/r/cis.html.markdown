---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM : Cloud Internet Services instance"
description: |-
  Manages IBM Cloud Internet Services instance.
---

# ibm_cis
Create, update, or delete an IBM Cloud Internet Services instance. The ibmcloud_api_key used by Terraform must have been granted sufficient rights to create IBM Cloud resources and have access to the resource group the CIS instance will be associated with. For more information, about CIS instance, see [getting started with CIS](https://cloud.ibm.com/docs/cis?topic=cis-getting-started).

If `resource_group_id` is not specified, the CIS instance is created in the default resource group. The API_KEY must have been assigned permissions for this group.

## Example usage

```terraform
data "ibm_resource_group" "group" {
  name = "test"
}

resource "ibm_cis" "cis_instance" {
  name              = "test"
  plan              = "standard"
  resource_group_id = data.ibm_resource_group.group.id
  tags              = ["tag1", "tag2"]
  location          = "global"

  //User can increase timeouts
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}
```

## Timeouts

`ibm_cis` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

**Create**: The creation of the IBM Cloud Internet Services instance is considered failed if no response is received for 10 minutes.
**Update**: The update of the IBM Cloud Internet Services instance is considered failed if no response is received for 10 minutes.
**Delete**: The deletion of the IBM Cloud Internet Services instance is considered failed if no response is received for 10 minutes.

## Argument reference
Review the argument references that you can specify for your resource.

- `location` - (Required, String) The target location where you want to create your instance.
- `name` - (Required, String) A descriptive name for your IBM Cloud Internet Services instance.
- `parameters` (Optional, Map) Arbitrary parameters to create instance. The value must be a JSON object.
- `plan` - (Required, String) The name of the plan for your instance. To retrieve this value, run `ibmcloud catalog service internet-svcs` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).
- `resource_group_id` - (Optional, Forces new resource, String) The ID of the resource group where you want to create the service. To retrieve this value, run `ibmcloud resource groups` or use the `ibm_resource_group` data source. If no value is specified, the `default` resource group is used.
- `tags` - (Optional, Array of strings) A list of tags that you want to associate with the instance.

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `guid` - (String) The unique identifier of the CIS instance.
- `id` - (String) The CRN of the CIS instance.
* `service` - (String) The service type of the instance.
* `status` - (String) The status of the CIS instance.

## Import

The `ibm_cis` resource can be imported by using the `id`. The ID is formed from the Cloud Resource Name (CRN)  from the **Overview** page of the Internet Services instance. These will be located under the **Domain** heading. 

* CRN is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

**Syntax**

```
$ terraform import ibm_cis.myorg <crn>

```

**Example**

```
$ terraform import ibm_cis.myorg crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
