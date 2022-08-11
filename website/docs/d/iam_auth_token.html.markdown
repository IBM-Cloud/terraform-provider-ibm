---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM: ibm_iam_auth_token"
description: |-
  Get information about an IBM Cloud IAM and UAA tokens.
---

# ibm_iam_auth_token

Retrieve information about your IAM access token. You can use this token to authenticate with the IBM Cloud platform. For more information, about IAM and UAA token, see [access tokens](https://cloud.ibm.com/docs/appid?topic=appid-tokens).

## Example usage

```terraform
data "ibm_iam_auth_token" "tokendata" {}
```

## Attribute reference

You can access the following attribute references after your data source is created.

- `iam_access_token`  - (String) The IAM access token.
- `iam_refresh_token` - (String) The IAM refresh token.
- `uaa_access_token` - (String) The UAA access token.
- `uaa_refresh_token` - (String) The UAA refresh token.
  