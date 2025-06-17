---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : ibm_dl_gateway_macsec_cak"
description: |-
  Manages IBM Cloud Infrastructure Direct Link Gateway Macsec CAK.
---

# ibm_dl_gateway_macsec_cak

Create/Update/Delete a CAK associated with the MACsec configuration of a direct link.
A connectivity association key (CAK) used in the MACsec Key Agreement (MKA) protocol. MACsec CAKs consist of both a name and key. The CAK's name must be a hexadecimal string of even lengths between 2 to 64 inclusive. The CAK's key must be a [Hyper Protect Crypto Service Standard Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) type=standard with key material a hexadecimal string exactly 64 characters in length.

For more information, about IBM Cloud Direct Link, see [getting started with IBM Cloud Direct Link](https://cloud.ibm.com/docs/dl?topic=dl-get-started-with-ibm-cloud-dl).


## Example usage

---
```terraform
resource "ibm_dl_gateway_macsec_cak" "test" {
    gateway = "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
    version = "2019-12-13"
    key {
        crn = "crn:v1:bluemix:public:hs-crypto:us-south:a/4111d05f36894e3cb9b46a43556d9000:abc111b8-37aa-4034-9def-f2607c87aaaa:key:bbb222bc-430a-4de9-9aad-84e5bb022222"
    }
    name = "AA01"
    session = "fallback"
}
```
---
## Argument reference
Review the argument reference that you can specify for your resource. 

- `gateway` - (Required, String) Direct Link gateway identifier.
- `version` - (Required, String) Requests the version of the API as a date in the format `YYYY-MM-DD`. Any date from 2019-12-13 up to the current date may be provided. Specify the current date to request the latest version.
- `key` - (Required, List) A reference to a Hyper Protect Crypto Service Standard Key.
    Nested schema for `key`:
    - `crn` - (String) The CRN of the referenced key.
- `name` - (Required, String) The name identifies the connectivity association key (CAK) within the MACsec key chain. The CAK's `name` must be a hexadecimal string of even lengths between 2 to 64 inclusive. This value, along with the material of the `key`, must match on the MACsec peers.
- `session` - (Required, String) The intended session the key will be used to secure. If the `primary` MACsec session fails due to a key/key name mismatch on the peers, the `fallback` session can take over. There must be a `primary` session CAK. A `fallback` CAK is optional


## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created.


- `active_delta` - (List) This field will be present when the status of the MACsec CAK is rotating or inactive. It may be present when the CAK status is failed.
    Nested schema for `active_delta`:
    - `key` - (List) A reference to a Hyper Protect Crypto Service Standard Key.
        Nested schema for `key`:
        - `crn` - (String) The CRN of the referenced key.
    - `name` - (String) The name identifies the connectivity association key (CAK) within the MACsec key chain. The CAK's name must be a hexadecimal string of even lengths between 2 to 64 inclusive.This value, along with the material of the key, must match on the MACsec peers.
    - `status` - (String) Current status of the CAK.
- `created_at` - (String) The date and time the resource was created.
- `key` - (List) A reference to a Hyper Protect Crypto Service Standard Key.
    Nested schema for `key`:
    - `crn` - (String) The CRN of the referenced key.
- `name` - (String) The name identifies the connectivity association key (CAK) within the MACsec key chain. The CAK's `name` must be a hexadecimal string of even lengths between 2 to 64 inclusive. This value, along with the material of the `key`, must match on the MACsec peers.
- `session` - (String) The intended session the key will be used to secure. If the `primary` MACsec session fails due to a key/key name mismatch on the peers, the `fallback` session can take over. There must be a `primary` session CAK. A `fallback` CAK is optional
- `status` - (String) Current status of the CAK.
    - Status `operational` is returned when the CAK is configured - successfully.
    - Status `rotating` is returned during a key rotation. The CAK defined by `active_delta` is still configured on the device and could be securing the MACsec session. In the case of a primary CAK, the status will be `rotating` for a period of time while waiting for the MACsec session to be secured with the new CAK. After that time, the CAK will either enter `active` or `inactive` status.
    - Status `active` is returned when the CAK is configured successfully and is currently used to secure the MACsec session.
    - Status `inactive` is returned when the CAK is configured successfully, but is not currently used to secure the MACsec session. The CAK may enter `rotating` status, and ultimately the `active` status, if it is found to be used to secure the MACsec session. The CAK may never leave this status on its own (e.g. if there is a key/key name mismatch). You are allowed to patch the CAK in this state to start the rotation procedure again.
    - Status `failed` is returned when the CAK cannot be configured. To recover, first resolve any issues with your HPCS key, then patch this CAK with the same or new key. Alternatively, you can delete this CAK if used for the `fallback` session.
- `updated_at` - (String) The date and time the resource was last updated.
