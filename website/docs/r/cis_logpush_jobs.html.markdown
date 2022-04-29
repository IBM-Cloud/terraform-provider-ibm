---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_logpush_jobs"
description: |-
  Get information on an IBM Cloud Internet Services logpush jobs.
---

# ibm_cis_logpush_jobs

Retrieve information about an IBM Cloud Internet Services logpush jobs data sources. For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis?topic=cis-about-ibm-cloud-internet-services-cis).

## Example usage

```terraform
resource "ibm_cis_logpush_job" "test" {
		cis_id          = data.ibm_cis.cis.id
		domain_id       = data.ibm_cis_domain.cis_domain.domain_id
		name            = "MylogpushJob"
		enabled         = false
		logpull_options = "timestamps=rfc3339&timestamps=rfc3339"
		dataset         = "http_requests"
		frequency       = "low"
		logdna =<<LOG
			{
				"hostname": "examplse.cistest-load.com",
				"ingress_key": "e2f72cxxxxxxxxxxxxa0b87859e",
				"region": "in-che"
		}
		LOG
	}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Required, String) The Domain ID of the CIS service instance.
- `name` - (Required, String) Logpush Job Name.
- `enabled` - (Required, Boolean) Whether the logpush job enabled or not.
- `logpull_options` - (Required, String) Configuration string.
- `dataset` - (Optional, String) Dataset to be pulled,Option for dataset`http_requests`,`range_events`,`firewall_events`
- `frequency` - (Optional, String) The frequency at which CIS sends batches of logs to your destination.`high`, `low`
- `logdna` - (Required, String)Information to identify the LogDNA instance the data will be pushed. Must provided in JSON format. Which need hostename,ingress_key and region (https://cloud.ibm.com/docs/cis?topic=cis-logpush&interface=api)


## Attributes Reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of logpush job resource. It is a combination of <`job-id`>:<`crn`> attributes concatenated with ":".
- `job_id` - (String) Unique identifier for the each LogpushJob.


## Import

The `ibm_cis_logpus_jobs` resource can be imported using the `id`. The ID is formed from the `Job ID`and the `CRN` (Cloud Resource Name) concatentated usinga `:` character.

The CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` CLI commands.

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **Job ID** is a 32 digit character string of the form: `52bfa670237f49ecb68473033c569649`. 

**Syntax**

```
$ terraform import ibm_cis_logpus_jobs.myorg <job_id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_logpus_jobs.myorg
crn:v1:bluemix:public:internet-svcs-ci:global:a/01652b251c3ae2787110a995d8db0135:9054ad06-3485-421a-9300-fe3fb4b79e1d::
```
