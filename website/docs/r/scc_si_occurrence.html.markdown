---
layout: "ibm"
subcategory: "Security and Compliance Center"
page_title: "IBM : ibm_scc_si_occurrence"
description: |-
  Manages Security and Compliance Center occurrence.
---

# DEPRECATED
Security and Compliance Center - Security Insights has now deprecated, backend services are no longer available. The docs will be removed in next release.

# ibm_scc_si_occurrence

Create, update, or delete for a Security and Compliance Center occurrence. For more information, about Security and Compliance Center, see [getting started with Security and Compliance Center](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-getting-started).

## Example usage

#### FINDING

```terraform
resource "ibm_scc_si_occurrence" "finding-occurrence" {
  provider_id   = var.provider_id
  note_name     = var.note_name
  kind          = "FINDING"
  occurrence_id = "finding-occ"
  resource_url  = "https://cloud.ibm.com"
  remediation   = "Limit the cluster access"
  finding {
    severity  = "LOW"
    certainty = "LOW"
    next_steps {
      title = "Security Threat"
      url   = "https://cloud.ibm.com/security-compliance/findings"
    }
  }
}
```

#### KPI

```terraform
resource "ibm_scc_si_occurrence" "kpi-occurrence" {
  provider_id   = var.provider_id
  note_name     = var.note_name
  kind          = "KPI"
  occurrence_id = "kpi-occ"
  resource_url  = "https://cloud.ibm.com"
  remediation   = "Limit the cluster access"
  kpi {
    value = 40
    total = 100
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `context` - (Optional, List) 

  Nested scheme for **context**:
  - `component_name` - (Optional, String) The name of the component the occurrence applies to.
  - `environment_name` - (Optional, String) The name of the environment the occurrence applies to.
  - `region` - (Optional, String) The IBM Cloud region.
  - `resource_crn` - (Optional, String) The resource CRN (e.g. certificate CRN, image CRN).
  - `resource_id` - (Optional, String) The resource ID, in case the CRN is not available.
  - `resource_name` - (Optional, String) The user-friendly resource name.
  - `resource_type` - (Optional, String) The resource type name (e.g. Pod, Cluster, Certificate, Image).
  - `service_crn` - (Optional, String) The service CRN (e.g. CertMgr Instance CRN).
  - `service_name` - (Optional, String) The service name (e.g. CertMgr).
  - `toolchain_id` - (Optional, String) The id of the toolchain the occurrence applies to.
- `finding` - (Optional, List) Finding provides details about a finding occurrence.

   Nested scheme for **finding**:
	- `certainty` - (Optional, String) Note provider-assigned confidence on the validity of an occurrence- LOW&#58; Low Certainty- MEDIUM&#58; Medium Certainty- HIGH&#58; High Certainty.
	  - Constraints: Allowable values are: `LOW`, `MEDIUM`, `HIGH`.
	- `data_transferred` - (Optional, List) It provides details about data transferred between clients and servers.
	  
	  Nested scheme for **data_transferred**:
	  - `client_bytes` - (Optional, Integer) The number of client bytes transferred.
	  - `client_packets` - (Optional, Integer) The number of client packets transferred.
	  - `server_bytes` - (Optional, Integer) The number of server bytes transferred.
	  - `server_packets` - (Optional, Integer) The number of server packets transferred.
	- `network_connection` - (Optional, List) It provides details about a network connection.
	
	Nested scheme for **network_connection**:
		* `client` - (Optional, List) It provides details about a socket address.
		Nested scheme for **client**:
			* `address` - (Required, String) The IP address of this socket address.
			* `port` - (Optional, Integer) The port number of this socket address.
		* `direction` - (Optional, String) The direction of this network connection.
		* `protocol` - (Optional, String) The protocol of this network connection.
		* `server` - (Optional, List) It provides details about a socket address.
		Nested scheme for **server**:
			* `address` - (Required, String) The IP address of this socket address.
			* `port` - (Optional, Integer) The port number of this socket address.
	* `next_steps` - (Optional, List) Remediation steps for the issues reported in this finding. They override the note's next steps.
	Nested scheme for **next_steps**:
		* `title` - (Optional, String) Title of this next step.
		* `url` - (Optional, String) The URL associated to this next steps.
	* `severity` - (Optional, String) Note provider-assigned severity/impact ranking- LOW&#58; Low Impact- MEDIUM&#58; Medium Impact- HIGH&#58; High Impact- CRITICAL&#58; Critical Impact.
	  * Constraints: Allowable values are: `LOW`, `MEDIUM`, `HIGH`, `CRITICAL`.
* `kind` - (Required, String) The type of note. Use this field to filter notes and occurrences by kind. - FINDING&#58; The note and occurrence represent a finding. - KPI&#58; The note and occurrence represent a KPI value. - CARD&#58; The note represents a card showing findings and related metric values. - CARD_CONFIGURED&#58; The note represents a card configured for a user account. - SECTION&#58; The note represents a section in a dashboard.
  * Constraints: Allowable values are: `FINDING`, `KPI`, `CARD`, `CARD_CONFIGURED`, `SECTION`.
* `kpi` - (Optional, List) Kpi provides details about a KPI occurrence.
Nested scheme for **kpi**:
	* `total` - (Optional, Float) The total value of this KPI.
	* `value` - (Required, Float) The value of this KPI.
* `note_name` - (Required, String) An analysis note associated with this image, in the form "{account_id}/providers/{provider_id}/notes/{note_id}" This field can be used as a filter in list requests.
* `account_id` - (Optional, Forces new resource, String) Account ID is optional, if not provided value will be inferred from the token retrieved from the IBM Cloud API key.
* `occurrence_id` - (Required, Forces new resource, String) The ID of the occurrence.
* `provider_id` - (Required, Forces new resource, String) Part of the parent. This field contains the provider ID. For example: providers/{provider_id}.
* `remediation` - (Optional, String) A description of actions that can be taken to remedy the `Note`.
* `replace_if_exists` - (Optional, Boolean) When set to true, an existing occurrence is replaced rather than duplicated.
* `resource_url` - (Optional, String) The unique URL of the resource, image or the container, for which the `Occurrence` applies. For example, https://gcr.io/provider/image@sha256:foo. This field can be used as a filter in list requests.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the scc_si_occurrence.
- `create_time` - (Optional, String) Output only. The time this `Occurrence` was created.
- `update_time` - (Optional, String) Output only. The time this `Occurrence` was last updated.

## Import

You can import the `ibm_scc_si_occurrence` resource by using `id`.
The `id` property can be formed from `provider_id`, and `occurrence_id` in the following format:

```sh
<account_id>/<provider_id>/<occurrence_id>
```
- `account_id` - A string. AccountID from the resource has to be imported.
- `provider_id`: A string. Part of the parent. This field contains the provider ID. For example: **providers/{provider_id}**.
- `occurrence_id`: A string. Second part of occurrence `name`: **providers/{provider_id}/occurrences/{occurrence_id}**.

# Syntax

```sh
$ terraform import ibm_scc_si_occurrence.scc_si_occurrence <account_id>/<provider_id>/<occurrence_id>
```
