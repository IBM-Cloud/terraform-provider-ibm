---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID SAML IDP"
description: |-
  Provides AppID SAML IDP resource.
---

# ibm_appid_idp_saml

Update or reset an IBM Cloud AppID Management Services SAML IDP configuration. For more information, see [SAML](https://cloud.ibm.com/docs/appid?topic=appid-enterprise)

## Example usage

```terraform
resource "ibm_appid_idp_saml" "saml" {
  tenant_id = var.tenant_id
  is_active = true
  config {
    entity_id = "https://test-saml-idp"
    sign_in_url = "https://test-saml-idp/login"
    display_name = "%s"
    encrypt_response = true
    sign_request = false
    certificates = [
      <<EOF
MIIDNTCCAh2gAwIBAgIRAPbl3OBL5oXq47d98l2s/3IwDQYJKoZIhvcNAQELBQAw
KDEQMA4GA1UEChMHRXhhbXBsZTEUMBIGA1UEAxMLZXhhbXBsZS5jb20wHhcNMjAw
OTIxMTEyMjI1WhcNMzAwOTI5MTEyMjI1WjAoMRAwDgYDVQQKEwdFeGFtcGxlMRQw
EgYDVQQDEwtleGFtcGxlLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC
ggEBAJqMcqnms1XpCuKz+CIVrqppMog9aerAQEV5wY6XuvakZ/w89zrA7YX3vwgi
+0ZO9ldDBh5Wvl8Li8vDFALJc42MxxyENk4qB6zee1O+zYu1Bwynkp7nIxqyKKRd
+0tvc+WHUbPFHvXc94rajT/csHOvBRiLmABMBx/IqF1nEAG/+KAEh7+KZYbvQ6wk
OoiPZlW+B0HR/DL/uO/v1Q7eq2Z8pAVTGikHefckvolkOqiCIRZx8HDe8DxTojEm
ygiR1aeT29XV8frI3Y2C8e7vgDpuZ8nV+0JUzqi5tAfl8bUfuq/W0eng6BYk2hBD
uuS66fHb1hnW96WIaExlK6T096sCAwEAAaNaMFgwDgYDVR0PAQH/BAQDAgKkMA8G
A1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFP3hKtk3MVXSF1H79ukO7oBVwwAkMBYG
A1UdEQQPMA2CC2V4YW1wbGUuY29tMA0GCSqGSIb3DQEBCwUAA4IBAQA9TbumFQHA
SHS6DBzzJz8GeX451AelW8UtpIuc5mRDFvTFEuNn/wMikxi+m8SQgkcuO5wfQi+0
FzLQO8DYH6fnAo1BYqooT1bt4lXflt74FnyUYbZ75yUhsddYF00FYOX6eOxrAU/U
qaPXw2N/e6S859hsUMq79/g3ES9sdNedtiwgiQv7roh4WNSvgTLh+sD32Ehl+x/I
eE80MljFLf5bfu2bQqV7C17lszGxTQWI2Xj56gLr2jcITjltcHCuBwnRDyXJkNhq
/2KRyIGAaRkkCOJAJxiz82wxkuQ8aL4sD3dctfGNu2Qe1JXHB65M1P2m0j/IcrLT
iUCoFQ0xO5VC
EOF
    ]
  }
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `is_active` (Required, Boolean) SAML IDP activation
- `config` (List of Object, Max: 1) SAML IDP configuration

    Nested scheme for `config`:
    - `entity_id` - (Required, String) Unique name for an Identity Provider
    - `sign_in_url` - (Required, String) SAML SSO url
    - `certificates` - (Required, List of String) List of certificates, primary and optional secondary
    - `display_name` - (Optional, String) Optional provider name
    - `encrypt_response` - (Optional, Bool) `true` if SAML responses should be encrypted
    - `sign_request` - (Optional, Bool) `true` if SAML requests should be signed
    - `include_scoping` - (Optional, Bool) `true` if scopes are included
    - `authn_context` - (Optional, List of Object, Max: 1) SAML authNContext configuration

        Nested scheme for `authn_context`:
        `class` - (List of String) List of `authnContext` classes 
        `comparison` - (String) Allowed values: `exact`, `maximum`, `minimum`, `better`


## Import

The `ibm_appid_idp_saml` resource can be imported by using the AppID tenant ID.

**Syntax**

```bash
$ terraform import ibm_appid_idp_saml.saml <tenant_id>
```
**Example**

```bash
$ terraform import ibm_appid_idp_saml.saml 5fa344a8-d361-4bc2-9051-58ca253f4b2b
```
