---
subcategory: "Direct Link Gateway Macsec Config"
layout: "ibm"
page_title: "IBM : ibm_dl_gateway_macsec_config"
description: |-
  Manages IBM Cloud Infrastructure Direct Link Gateway Macsec Config.
---

# ibm_dl_gateway_macsec_config

Set/Unset/Update the MACsec configuration of a IBM Cloud Infrastructure Direct Link Gateway. For more information, about IBM Cloud Direct Link, see [getting started with IBM Cloud Direct Link](https://cloud.ibm.com/docs/dl?topic=dl-get-started-with-ibm-cloud-dl).


## Example usage

---
```terraform
resource "ibm_dl_gateway_macsec_config" "test" {
    gateway = "0a06fb9b-820f-4c44-8a31-77f1f0806d28"
    version = "2019-12-13"
}
```
---
## Argument reference
Review the argument reference that you can specify for your resource. 

- `gateway` - (Required, String) Direct Link gateway identifier.
- `version` - (Required, String) Requests the version of the API as a date in the format `YYYY-MM-DD`. Any date from 2019-12-13 up to the current date may be provided. Specify the current date to request the latest version.
- `active` - (Optional, Bool) Indicates if the MACsec feature is currently active (true) or inactive (false) for a gateway.
- `sak_rekey` - (Optional, List) Determines how SAK rekeying occurs. It is either timer based or based on the amount of used packet numbers.
    Nested scheme for `sak_rekey`:
    - `interval` - (Integer) The time, in seconds, to force a Secure Association Key (SAK) rekey.
    - `mode` - (String) Determines that the SAK rekey occurs based on a timer.
 - `security_policy` - (Optional, String) Determines how packets without MACsec headers are handled. `must_secure` - Packets without MACsec headers are dropped. This policy should be used to prefer security over network availability. `should_secure` - Packets without MACsec headers are allowed. This policy should be used to prefer network availability over security.
 - `window_size` - (Optional, Integer) The window size determines the number of frames in a window for replay protection. Replay protection is used to counter replay attacks. Frames within a window size can be out of order and are not replay protected.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `active` - (Bool) Indicates if the MACsec feature is currently active (true) or inactive (false) for a gateway.
- `cipher_suite` - (String) The cipher suite used in generating the security association key (SAK).
- `confidentiality_offset` - (Integer) The confidentiality offset determines the number of octets in an Ethernet frame that are not encrypted.
- `created_at` - (String) The date and time the resource was created.
- `key_server_priority` - (Integer) Used in the MACsec Key Agreement (MKA) protocol to determine which peer acts as the key server. Lower values indicate a higher preference to be the key server. The MACsec configuration on the direct link will always set this value to 255.
- `sak_rekey` - (List) Determines how SAK rekeying occurs. It is either timer based or based on the amount of used packet numbers.
    Nested scheme for `sak_rekey`:
    - `interval` - (Integer) The time, in seconds, to force a Secure Association Key (SAK) rekey.
    - `mode` - (String) Determines that the SAK rekey occurs based on a timer.
- `security_policy` - (String) Determines how packets without MACsec headers are handled. `must_secure` - Packets without MACsec headers are dropped. This policy should be used to prefer security over network availability. `should_secure` - Packets without MACsec headers are allowed. This policy should be used to prefer network availability over security.
- `status` - (String) Current status of MACsec on this direct link. Status `offline` is returned when MACsec is inactive and during direct link creation. Status `deleting` is returned when MACsec during removal of MACsec from the direct link and during direct link deletion. See `status_reasons[]` for possible remediation of the `failed` status.
- `status_reasons` - (List) Context for certain values of status.
    Nested Schema for `status_reasons`:
    - `code` - (String) A reason code for the status: `macsec_cak_failed` -  At least one of the connectivity association keys (CAKs) associated with the MACsec configuration was unable to be configured on the direct link gateway. Refer to the status of the CAKs associated with the MACsec configuration to find the the source of this reason.
    - `message` - (String) An explanation of the status reason.
    - `more_info` - (String) Link to documentation about this status reason.
- `updated_at` - (String) The date and time the resource was last updated.
- `window_size` - (Integer) The window size determines the number of frames in a window for replay protection. Replay protection is used to counter replay attacks. Frames within a window size can be out of order and are not replay protected.