---
layout: "ibm"
page_title: "IBM : ibm_restapi_data"
description: |-
  Reads data from an arbitrary IBM Cloud REST API.
subcategory: "REST API"
---

# ibm_restapi_data

Read data from an arbitrary IBM Cloud REST API. Use this data source for IBM Cloud services that do not have a dedicated Terraform data source yet, instead of adding a third party REST provider to your configuration.

The request is authenticated with the credentials that are already configured on the `ibm` provider, so you do not need to manage an API key or a bearer token separately.

## Example usage

```hcl
data "ibm_restapi_data" "resource_groups" {
  url = "https://resource-controller.cloud.ibm.com/v2/resource_groups"

  query_params = {
    account_id = var.account_id
  }
}

locals {
  resource_group_names = [
    for group in jsondecode(data.ibm_restapi_data.resource_groups.response_body).resources : group.name
  ]
}
```

Call a search style API that takes a query body.

```hcl
data "ibm_restapi_data" "search" {
  url    = "https://api.global-search-tagging.cloud.ibm.com/v3/resources/search"
  method = "POST"

  request_body = jsonencode({
    query = "type:cf-application"
  })
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `accept_status_codes` - (Optional, List of Integers) Status codes outside the 2xx range that are treated as success instead of an error, for example `[404]`.
- `headers` - (Optional, Map of Strings) Additional HTTP headers that are sent with the request. `Accept` defaults to `application/json`.
- `method` - (Optional, String) The HTTP method used to read the data. One of `GET`, `POST`, `HEAD`. The default value is `GET`.
- `query_params` - (Optional, Map of Strings) Query parameters that are appended to the request URL. Parameters that are already present on `url` are kept, and entries here win when the same key is given twice.
- `request_body` - (Optional, String) The JSON body that is sent with the request. Only useful when `method` is `POST`.
- `sensitive_headers` - (Optional, Map of Strings) Additional HTTP headers that are sent with the request and are kept out of the Terraform output. Values here override values with the same key in `headers`.
- `timeout_seconds` - (Optional, Integer) How long the HTTP request may take before it is cancelled. Between `1` and `3600`. The default value is `60`.
- `url` - (Required, String) The URL that is read, for example `https://resource-controller.cloud.ibm.com/v2/resource_instances`.
- `use_iam_auth` - (Optional, Boolean) Whether the IBM Cloud credentials that are configured on the provider are used to authenticate the request. Set it to `false` when the target API is not IAM protected, or when you supply an `Authorization` header through `sensitive_headers`. The default value is `true`.

## Attribute reference

In addition to all argument references, you can access the following attribute references after your data source is read.

- `id` - (String) The unique identifier of the read, derived from the request URL, method and body.
- `is_json` - (Boolean) Whether the response body parsed as JSON. Check it before you call `jsondecode` on `response_body`.
- `response_body` - (String) The raw response body. Use the `jsondecode` function to work with it.
- `response_headers` - (Map of Strings) The headers of the response.
- `response_status_code` - (Integer) The status code of the response.
