Create a policy for the load balancer listener

The following arguments are supported:

action - (Required, Forces new resource, string) The policy action. Allowable values: [forward,redirect,reject]
priority - (Required, Forces new resource, integer). Priority of the policy. Lower value indicates higher priority. 1 ≤ value ≤ 10
name - (Optional, string) The user-defined name for this policy. Names must be unique within the load balancer listener the policy resides in.
rules - (Optinal, list) The list of rules ( condition, type, value and field ) of this policy
		condition : Allowable values: [contains,equals,matches_regex]
		type : Allowable values: [header,hostname,path]
		value : Constraints: 1 ≤ length ≤ 128
		filed: Constraints: 1 ≤ length ≤ 128
target-id - (Optional, integer) The unique identifier for this load balancer pool, specified with 'forward' action
target-http-status-code -(Optional, integer) The http status code in the redirect response, one of [301, 302, 303, 307, 308], specified with 'redirect' action
target-url - (Optional, integer) The redirect target URL, specified with 'redirect' action 

Note : 

When action is forward, target-id should specify which pool the load balancer forwards the traffic to. When action is redirect,target-url should specify the url and target-http-status-code specify the code used in the redirect response.

Attribute Reference

The following attributes are exported:
* id - The unique identifier of the load balancer listener policy.
* status - The status of load balancer listener policy.

To run, configure your IBM Cloud provider

Running the example

For planning phase

terraform plan
For apply phase

terraform apply
For destroy

terraform destroy