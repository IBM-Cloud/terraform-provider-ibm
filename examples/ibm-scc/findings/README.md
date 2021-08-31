# Example for FindingsV1

This example illustrates how to use the FindingsV1

These types of resources are supported:

* scc_si_note

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## FindingsV1 resources

scc_si_note resource:

```hcl
resource "ibm_scc_si_note" "ts-card-finding" {
  provider_id       = "scc"
  short_description = "Security Threat"
  long_description  = "Security Threat found in your account"
  kind              = "CARD"
  note_id           = "ts-card-finding"
  reported_by {
    id    = "scc-si-terraform"
    title = "SCC SI Terraform"
    url   = "https://cloud.ibm.com"
  }
  card {
    section            = "Security"
    title              = "Threats"
    subtitle           = "Summary of Security Threats"
    finding_note_names = ["providers/scc/notes/finding"]
    elements {
      kind               = "TIME_SERIES"
      text               = "count"
      default_time_range = "3d"
      value_types {
        text               = "count"
        finding_note_names = ["providers/scc/notes/finding"]
        kind               = "FINDING_COUNT"
      }
    }
  }
}
```

## FindingsV1 Data sources

scc_si_providers data source:

```hcl
data "ibm_scc_si_providers" "providers" {
  limit = 4
}
```
scc_si_notes data source:

```hcl
data "ibm_scc_si_notes" "notes" {
  page_size = 3
}
```
scc_si_note data source:

```hcl
data "ibm_scc_si_note" "scc_si_note" {
	note_id = "note_id"
	provider_id = "provider_id"
}
```

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| provider_id | Part of the parent. This field contains the provider ID. For example: providers/{provider_id}. | `string` | true |
| short_description | A one sentence description of your note. | `string` | true |
| long_description | A more detailed description of your note. | `string` | true |
| kind | The type of note. Use this field to filter notes and occurences by kind. - FINDING&#58; The note and occurrence represent a finding. - KPI&#58; The note and occurrence represent a KPI value. - CARD&#58; The note represents a card showing findings and related metric values. - CARD_CONFIGURED&#58; The note represents a card configured for a user account. - SECTION&#58; The note represents a section in a dashboard. | `string` | true |
| note_id | The ID of the note. | `string` | true |
| reported_by | The entity reporting a note. | `` | true |
| related_url |  | `list()` | false |
| shared | True if this note can be shared by multiple accounts. | `bool` | false |
| finding | FindingType provides details about a finding note. | `` | false |
| kpi | KpiType provides details about a KPI note. | `` | false |
| card | Card provides details about a card kind of note. | `` | false |
| section | Card provides details about a card kind of note. | `` | false |
| id | The ID of the provider. | `string` | false |
| provider_id | Part of the parent. This field contains the provider ID. For example: providers/{provider_id}. | `string` | true |
| note_id | Second part of note `name`: providers/{provider_id}/notes/{note_id}. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| scc_si_note | scc_si_note object |
