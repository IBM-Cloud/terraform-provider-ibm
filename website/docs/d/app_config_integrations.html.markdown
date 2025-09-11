---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration Integrations'
description: |-
  Get information about Integrations
---

# ibm_app_config_integrations

Retrieve information about an existing IBM Cloud App Configuration Integrations. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_app_config_integrations" "app_config_integrations" {
  guid = "guid"
  limit = 1
  offset = 1
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) The GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `limit` - (optional, Integer) The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different set of records, use `limit` with `offset` to page through the available records.
- `offset` - (optional, Integer) The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. Use `offset` with `limit` to page through the available records.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `total_count` - (String) Number of records returned in the current response.\

- `integrations` - (List) The list of integrations.
  - `integration_type` - (String) The type of integration.
  - `integration_id` - (String) The id of integration.

- `first` - (List) The URL to navigate to the first page of records.
  - `href` - (String) The first `href` URL.

- `previous` - (List) The URL to navigate to the previous list of records.
  - `href` - (String) The previous `href` URL.

- `last` - (List) The URL to navigate to the last list of records.
  - `href` - (String) The last `href` URL.

- `next` - (List) The URL to navigate to the next list of records.
  - `href` - (String) The next `href` URL.
