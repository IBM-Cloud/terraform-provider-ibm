---
layout: "ibm"
subcategory: "Security and Compliance Center"
page_title: "IBM : ibm_scc_si_note"
description: |-
  Get information about scc_si_note
---

# DEPRECATED
Security and Compliance Center - Security Insights has now deprecated, backend services are no longer available. The docs will be removed in next release.

# ibm_scc_si_note

Provides a read-only data source for scc_si_note. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_scc_si_note" "scc_si_note" {
	note_id = "note_id"
	provider_id = "provider_id"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

* `account_id` - (Optional, String) Account ID is optional, if not provided value will be inferred from the token retrieved from the IBM Cloud API key.
* `note_id` - (Required, Forces new resource, String) Second part of note `name`: providers/{provider_id}/notes/{note_id}.
* `provider_id` - (Required, Forces new resource, String) Part of the parent. This field contains the provider ID. For example: providers/{provider_id}.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the scc_si_note.
* `card` - (Optional, List) Card provides details about a card kind of note.
Nested scheme for **card**:
	* `section` - (Required, String) The section this card belongs to.
	  * Constraints: The maximum length is `30` characters.
	* `title` - (Required, String) The title of this card.
	  * Constraints: The maximum length is `28` characters.
	* `subtitle` - (Required, String) The subtitle of this card.
	  * Constraints: The maximum length is `30` characters.
	* `order` - (Optional, Integer) The order of the card in which it will appear on SA dashboard in the mentioned section.
	  * Constraints: Allowable values are: 1, 2, 3, 4, 5, 6
	* `finding_note_names` - (Required, List) The finding note names associated to this card.
	* `requires_configuration` - (Optional, Boolean)
	  * Constraints: The default value is `false`.
	* `badge_text` - (Optional, String) The text associated to the card's badge.
	* `badge_image` - (Optional, String) The base64 content of the image associated to the card's badge.
	* `elements` - (Required, List) The elements of this card.
	Nested scheme for **elements**:
		* `text` - (Optional, String) The text of this card element.
		  * Constraints: The maximum length is `60` characters.
		* `kind` - (Optional, String) Kind of element- NUMERIC&#58; Single numeric value- BREAKDOWN&#58; Breakdown of numeric values- TIME_SERIES&#58; Time-series of numeric values.
		  * Constraints: The default value is `NUMERIC`. Allowable values are: NUMERIC, BREAKDOWN, TIME_SERIES
		* `default_time_range` - (Optional, String) The default time range of this card element.
		  * Constraints: The default value is `4d`. Allowable values are: 1d, 2d, 3d, 4d
		* `value_type` - (Optional, List)
		Nested scheme for **value_type**:
			* `kind` - (Optional, String) Kind of element- KPI&#58; Kind of value derived from a KPI occurrence.
			  * Constraints: Allowable values are: KPI
			* `kpi_note_name` - (Optional, String) The name of the kpi note associated to the occurrence with the value for this card element value type.
			* `text` - (Optional, String) The text of this element type.
			  * Constraints: The default value is `label`. The maximum length is `22` characters.
			* `finding_note_names` - (Optional, List) the names of the finding note associated that act as filters for counting the occurrences.
		* `value_types` - (Optional, List) the value types associated to this card element.
		Nested scheme for **value_types**:
			* `kind` - (Optional, String) Kind of element- KPI&#58; Kind of value derived from a KPI occurrence.
			  * Constraints: Allowable values are: KPI
			* `kpi_note_name` - (Optional, String) The name of the kpi note associated to the occurrence with the value for this card element value type.
			* `text` - (Optional, String) The text of this element type.
			  * Constraints: The default value is `label`. The maximum length is `22` characters.
			* `finding_note_names` - (Optional, List) the names of the finding note associated that act as filters for counting the occurrences.
		* `default_interval` - (Optional, String) The default interval of the time series.
		  * Constraints: The default value is `d`.

* `create_time` - (Optional, String) Output only. The time this note was created. This field can be used as a filter in list requests.

* `finding` - (Optional, List) FindingType provides details about a finding note.
Nested scheme for **finding**:
	* `severity` - (Required, String) Note provider-assigned severity/impact ranking- LOW&#58; Low Impact- MEDIUM&#58; Medium Impact- HIGH&#58; High Impact- CRITICAL&#58; Critical Impact.
	  * Constraints: Allowable values are: LOW, MEDIUM, HIGH, CRITICAL
	* `next_steps` - (Optional, List) Common remediation steps for the finding of this type.
	Nested scheme for **next_steps**:
		* `title` - (Optional, String) Title of this next step.
		* `url` - (Optional, String) The URL associated to this next steps.

* `kind` - (Required, String) The type of note. Use this field to filter notes and occurences by kind. - FINDING&#58; The note and occurrence represent a finding. - KPI&#58; The note and occurrence represent a KPI value. - CARD&#58; The note represents a card showing findings and related metric values. - CARD_CONFIGURED&#58; The note represents a card configured for a user account. - SECTION&#58; The note represents a section in a dashboard.
  * Constraints: Allowable values are: FINDING, KPI, CARD, CARD_CONFIGURED, SECTION

* `kpi` - (Optional, List) KpiType provides details about a KPI note.
Nested scheme for **kpi**:
	* `aggregation_type` - (Required, String) The aggregation type of the KPI values. - SUM&#58; A single-value metrics aggregation type that sums up numeric values  that are extracted from KPI occurrences.
	  * Constraints: The default value is `SUM`. Allowable values are: SUM

* `long_description` - (Required, String) A more detailed description of your note.

* `related_url` - (Optional, List) 
Nested scheme for **related_url**:
	* `label` - (Required, String) Label to describe usage of the URL.
	* `url` - (Required, String) The URL that you want to associate with the note.

* `reported_by` - (Required, List) The entity reporting a note.
Nested scheme for **reported_by**:
	* `id` - (Required, String) The id of this reporter.
	* `title` - (Required, String) The title of this reporter.
	* `url` - (Optional, String) The url of this reporter.

* `section` - (Optional, List) Card provides details about a card kind of note.
Nested scheme for **section**:
	* `title` - (Required, String) The title of this section.
	* `image` - (Required, String) The image of this section.

* `shared` - (Optional, Boolean) True if this note can be shared by multiple accounts.
  * Constraints: The default value is `true`.

* `short_description` - (Required, String) A one sentence description of your note.

* `update_time` - (Optional, String) Output only. The time this note was last updated. This field can be used as a filter in list requests.

