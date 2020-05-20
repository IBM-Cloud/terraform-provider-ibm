---
layout: "ibm"
page_title: "IBM: ibm_cis_rate_limit"
sidebar_current: "docs-ibm-cis-rate-limit"
description: |-
  Provides a IBM CIS Rate Limit resource.
---

# ibm_cis_rate_limit

Provides a IBM CIS Ratelimiting resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS Domain resource. It allows to create, update, delete custom ratelimits of a domain of a CIS instance

## Example Usage

```hcl
# Add a rate limit to the domain

resource "ibm_cis_rate_limit" "ratelimit" {
    cis_id = "${data.ibm_cis.cis.id}"
    domain_id = "${data.ibm_cis_domain.cis_domain.id}"
    threshold = 20
    period = 900
    match =[{
        request =[{
            url = "*.example.org/path*"
            schemes = ["HTTP", "HTTPS"]
            methods = ["GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"]
        }]
        response=[{
            status = [200, 201, 202, 301, 429]
            origin_traffic = false
        }]
    }]
    action =[{
        mode = "ban"
        timeout = 43200
        response =[{
            content_type = "text/plain"
            body = "custom response body"
        }]
    }]
    correlate =[{
        by = "nat"
    }]
    disabled = false
    description = "example rate limit for a zone"
}
```

## Argument Reference

The following arguments are supported:

* `cis_id` - (Required,string) The ID of the CIS service instance
* `domain_id` - (Required,string) The ID of the domain to add the Rate Limit rule.
* `threshold` - (Required,int).  The threshold that triggers the rate limit mitigations, combined with period. For example, threshold per period. Min value: 2, max value: 1000000.
* `period` - (Required,int).  The time, in seconds, to count matching traffic. If the count exceeds threshold within this period the action is performed. Min value:1, max value: 3600.
* `match` - (Optional,list).  Determines which traffic the rate limiting rule counts towards the threshold.
    * `request` - (Optional,list).  Matches HTTP requests.If not provided API ll default request to * , [_ALL_], `_ALL_` respectively.
        * `url` - (Optional,string).   The URL pattern to match comprised of the host and path, for instance, example.org/path. Wildcards are expanded to match applicable traffic, query strings are not matched. Use * for all traffic to your zone. Max length is 1024.
        * `schemes` - (Optional,set(string)).  HTTP Schemes, can be one [HTTPS], both [HTTP,HTTPS] or all [_ALL_]. This field is not required.
        * `methods` - (Optional,set(string)). HTTP Methods, can be a subset [POST,PUT] or all [_ALL_]. This field is not required to create a rate limit rule. Valid values are `GET`, `POST`, `PUT`, `DELETE`, `PATCH`, `HEAD`, `_ALL_`.
    * `response` - (Optional,list).  Matches HTTP responses before they are returned to the client . If this is defined, then the entire counting of traffic occurs at this stage. 
        * `status` - (Optional,set(int)).  HTTP Status codes, can be one [403], many [401,403] or indicate all by not providing this value. This field is not required. Min value: 100, max value: 999.
        * `origin_traffic` - (Optional,bool).  Orrigin traffic.
        * `header` - (Optional,list).  Array of response headers to match. If a response does not meet the header criteria then the request is not counted towards the rate limiting rule. The header matching criteria includes following properties.
            * `name` - (Optional,string).  The name of the response header to match.
            * `op` - (Optional,string).  The operator when matching, eq means equals, ne means not equals. Valid values are [`eq`] and [`ne`].
            * `value` - (Optional,string).  WThe value of the header, which is exactly matched.
* `action` - (Required,list).  The action performed when the threshold of matched traffic within the period defined is exceeded.
    * `mode` - (Required,string).  The type of action performed. Valid values are: [`simulate`], [`ban`], [`challenge`], [`js_challenge`].
    * `timeout` - (Optional,int).  The time, in seconds, as an integer to perform the mitigation action. Timeout be the same or greater than the period. This field is valid only when mode is [`simulate`] or [`ban`]. Min value: 10, max value: 86400.
    * `response` - (Optional,list).  Custom content-type and body to return. This overrides the custom error for the zone. Omission results in the default HTML error page. This field is valid only when mode is [`simulate`] or [`ban`].
        * `content_type` - (Optional,string).  The content-type of the body, which must be one of the following: [`text/plain`], [`text/xml`], [`application/json`].
        * `body` - (Optional,string).  The body to return. The content here must conform to the `content_type`. Max length is 10240.
* `disabled` - (Optional,bool).  Whether this rate limiting rule is currently disabled.
* `description` - (Optional,string).  A note that you can use to describe the reason for a rate limiting rule.
* `correlate` - (Optional,list).  Whether to enable NAT based rate limiting.
    * `by` - (Optional,string).  Valid values: [`nat`].
* `bypass` - (Optional,list).  Criteria that allows the rate limit to be bypassed. For example, to express that you shouldn’t apply a rate limit to a given set of URLs.
    * `name` - (Optional,string).  Valid values is [`url`].
    * `value` - (Optional,string).  The url to bypass.

**NOTE:**  To create a custom rate limit rule the CIS instance should be a `enterprise` plan

## Attributes Reference

The following attributes are exported:

* `id` - The record ID. It is a combination of <`rule_id`>,<`domain_id`>,<`cis_id`> attributes concatenated with ":".
* `rule_id` - The Rate Limit Rule ID.


## Import

The `ibm_cis_rate_limit` resource can be imported using the `id`. The ID is formed from the  `Rate Limit rule ID`, the `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated using a `:` character.  

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `bx cis` CLI commands.

* **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

* **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

* **Rate Limit rule ID** is a 32 digit character string of the form: `489d96f0da6ed76251b475971b097205c`.


```
$ terraform import ibm_cis_rate_limit.ratelimit <rule_id>:<domain-id>:<crn>

$ terraform import ibm_cis_rate_limit.ratelimit 48996f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::