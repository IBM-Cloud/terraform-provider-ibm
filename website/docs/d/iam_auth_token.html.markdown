---
layout: "ibm"
page_title: "IBM: ibm_iam_auth_token"
sidebar_current: "docs-ibm-datasource-iam-auth-token"
description: |-
  Get information about an IBM Cloud IAM and UAA tokens.
---

# ibm\_iam_auth_token

Import the details of an existing IBM Cloud authentication tokens as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_auth_token" "tokendata" {}
```

## Attribute Reference

The following attributes are exported:

* `iam_access_token` - IAM access Token. 
* `iam_refresh_token` - IAM Refresh Token.
* `uaa_access_token`- UAA access Token. 
* `uaa_refresh_token` -  UAA Refresh Token.
  