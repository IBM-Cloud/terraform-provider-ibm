---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_custom_list"
description: |-
  Provides an IBM CIS custom list resource.
---

# ibm_cis_custom_list

Provides an IBM Cloud Internet Services custom list resource to create, update, and delete the custom list of an instance. For more information about the IBM Cloud Internet Services custom list, see [custom list](https://cloud.ibm.com/docs/cis).

## Example usage

```terraform

  resource ibm_cis_custom_list custom_list {
    cis_id      = ibm_cis.instance.id
    kind        = var.list.kind
    name        = var.list.name
    description = var.list.description
  }

```

## Argument reference

Review the argument references that you can specify for your resource.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `description` - (Optional, string) Description of the custom list.
- `kind` - (Required, string) The kind of the custom list. Allowed values are `ip`, `asn`, `hostname`.
- `name` - (Required, string) Name of the custom list. Use this name in filter and rule expressions.

## Attribute reference

In addition to the argument reference list, you can access the following attribute reference after your resource is created.

- `list_id` - (String) ID of the custom list.
- `num_items` - (int) The number of items in the custom list.
- `num_referencing_filters` - (int) The number of filters referencing the custom list.

## Import

The `ibm_cis_custom_list` resource is imported by using the ID. The ID is formed from the list ID and the Cloud Resource Name (CRN) concatenated using a `:` character.

The CRN is located on the **Overview** page of the Internet Services instance of the domain heading of the console, or by using the `ibm cis` CLI commands.

- **List Id** is a 32-digit character string of the form: `77bc00aa67184d0b8b7233b131c432cf`.

- **CRN** is a 120-digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`.

### Syntax

``` terraform
terraform import ibm_cis_custom_list.custom_list <list-id>:<crn>
```

### Example

``` terraform
terraform import ibm_cis_custom_list.custom_list 77bc00aa67184d0b8b7233b131c432cf:crn:v1:staging:public:internet-svcs-ci:global:a/01652b251c3ae2787110a995d8db0135:1a9174b6-0106-417a-844b-c8eb43a72f63::
```
