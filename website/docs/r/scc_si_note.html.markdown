---
layout: "ibm"
page_title: "IBM : sa_note"
sidebar_current: "docs-ibm-resource-sa-note"
description: |-
  Manages IBM Cloud Security Advisor Findings Note.
---

# ibm\_sa\_note

This resource is used to order a IBM Cloud Security Advisor Findings Note.

## Example Usage

```hcl
resource "ibm_sa_note" "note" {
  provider_id                    = "test"
  short_description              = "short"
  long_description               = "long"
  kind                           = "FINDING"
  related_url {
    label   = "rel_label"
    url     = "rel_url"
  }
  expiration_time                = "2006-01-02 15:04:11"
  create_time                    = "2006-01-02 15:04:00"
  update_time                    = "2006-01-02 15:04:11"
  note_id                        = "id"
  shared                         = "true"
  reported_by {
    id    = "rep_id"
    title = "rep_title"
    url   = "rep_url"
  }
  finding {
    severity = "HIGH"
    next_steps {
      url   = "next_url"
      title = "next_title"
    }
  }
}
```

## Argument Reference

The following arguments are supported:

- `short_description` - A one sentence description of this Note.
- `long_description` - A detailed description of this Note.
- `kind` - Kind of this note. Possible values: [ FINDING, KPI, CARD, CARD_CONFIGURED, SECTION ]
- `related_url` - List of metadata for any related URL information
  - `label` - Label to describe usage of the URL
  - `url` - Specific URL to associate with the note
- `expiration_time` - Time of expiration for this note, null if note does not expire.
- `create_time` - Output only. The time this note was created.
- `update_time` - Output only. The time this note was last updated.
- `id` - The id of the note
- `reported_by` - The entity reporting a note
  - `id` - The id of this reporter
  - `title` - The title of this reporter
  - `url` - The url of this reporter
- `shared` - True if this Note can be shared by multiple accounts.
- `finding` - FindingType provides details about a finding note.
  - `severity` - Note provider-assigned severity/impact ranking. Possible values: [ LOW, MEDIUM, HIGH, CRITICAL ]
  - `next_steps` - List of common remediation steps for the finding of this type
    - `title` - Title of this next step
    - `url` - The URL associated to this next steps
- `kpi` - KpiType provides details about a KPI note.
  - `aggregation_type` - The aggregation type of the KPI values. Possible values: [ SUM ]
- `card` - Card provides details about a card kind of note
  - `section` - The section this card belongs to
  - `title` - The title of this card
  - `subtitle` - The subtitle of this card
  - `order` - The order of the card in which it will appear on SA dashboard in the mentioned section. Possible values: [ 1, 2, 3, 4, 5, 6 ]
  - `finding_note_names` - List of the finding note names associated to this card
  - `requires_configuration` - True if the card requires configuration
  - `badge_text` - The text associated to the card’s badge
  - `badge_image` - The base64 content of the image associated to the card’s badge
  - `elements` - List of the elements of this card
    - `kind` - Kind of element. Possible values: [ NUMERIC, BREAKDOWN, TIME_SERIES ]
    - `default_time_range` - The default time range of this card element. Possible values: [ 1d, 2d, 3d, 4d ]
- `section` - Section provides details about a section of note
  - `title` - The title of this section
  - `image` - The image of this section

## Attribute Reference

The following attributes are exported:

- `note` - Object of Security Advisor Findings Note.
  - `short_description` - A one sentence description of this Note.
  - `long_description` - A detailed description of this Note.
  - `kind` - Kind of this note. Possible values: [ FINDING, KPI, CARD, CARD_CONFIGURED, SECTION ]
  - `related_url` - List of metadata for any related URL information
    - `label` - Label to describe usage of the URL
    - `url` - Specific URL to associate with the note
  - `expiration_time` - Time of expiration for this note, null if note does not expire.
  - `create_time` - Output only. The time this note was created.
  - `update_time` - Output only. The time this note was last updated.
  - `id` - The id of the note
  - `reported_by` - The entity reporting a note
    - `id` - The id of this reporter
    - `title` - The title of this reporter
    - `url` - The url of this reporter
  - `shared` - True if this Note can be shared by multiple accounts.
  - `finding` - FindingType provides details about a finding note.
    - `severity` - Note provider-assigned severity/impact ranking. Possible values: [ LOW, MEDIUM, HIGH, CRITICAL ]
    - `next_steps` - List of common remediation steps for the finding of this type
      - `title` - Title of this next step
      - `url` - The URL associated to this next steps
  - `kpi` - KpiType provides details about a KPI note.
    - `aggregation_type` - The aggregation type of the KPI values. Possible values: [ SUM ]
  - `card` - Card provides details about a card kind of note
    - `section` - The section this card belongs to
    - `title` - The title of this card
    - `subtitle` - The subtitle of this card
    - `order` - The order of the card in which it will appear on SA dashboard in the mentioned section. Possible values: [ 1, 2, 3, 4, 5, 6 ]
    - `finding_note_names` - List of the finding note names associated to this card
    - `requires_configuration` - True if the card requires configuration
    - `badge_text` - The text associated to the card’s badge
    - `badge_image` - The base64 content of the image associated to the card’s badge
    - `elements` - List of the elements of this card
      - `kind` - Kind of element. Possible values: [ NUMERIC, BREAKDOWN, TIME_SERIES ]
      - `default_time_range` - The default time range of this card element. Possible values: [ 1d, 2d, 3d, 4d ]
  - `section` - Section provides details about a section of note
    - `title` - The title of this section
    - `image` - The image of this section
