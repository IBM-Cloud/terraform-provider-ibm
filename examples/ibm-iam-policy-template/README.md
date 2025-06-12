# Example for IamPolicyManagementV1

This example illustrates how to use the IamPolicyManagementV1

The following types of resources are supported:

* policy_template

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## IamPolicyManagementV1 resources

policy_template resource:

```hcl
resource "ibm_iam_policy_template" "template" {
	name = "TestTemplates"
	account_id = "xxxxxx"
	policy {
		type = "access"
		description = "description"
		resource {
			attributes {
				key = "serviceName"
				operator = "stringEquals"
				value = "is"
			}
			tags {
				key = "terraform"
				operator = "stringEquals"
				value = "terraform_test"
					
			}
		}
		roles = ["Operator", "Bare Metal Console Admin" , "Viewer"]
	}
	description = "Base Template"
	committed = "true"
}

resource "ibm_iam_policy_template_version" "template_version" {
	template_id = ibm_iam_policy_template.template.template_id 
	policy {
		type = "access"
		description = "description"
		resource {
			attributes {
				key = "service_group_id"
				operator = "stringEquals"
				value = "IAM"
			}
		}
		roles = ["Service ID creator", "Operator", "Viewer" ]
		rule_conditions {
			key = "{{environment.attributes.day_of_week}}"
			operator = "dayOfWeekAnyOf"
			value = ["1+00:00","2+00:00","3+00:00","4+00:00", "5+00:00"]
		}
	  pattern = "time-based-conditions:weekly:all-day"
	}
	description = "Template version2"
	committed = "true"
}
```

## IamPolicyManagementV1 data sources

ibm_iam_policy_template data source:

```hcl
data "ibm_iam_policy_template" "template" {
	account_id = "xxxxx"
}
data "ibm_iam_policy_template_version" "template" {
	policy_template_id = "policyTemplate-xxx-xx-xxxx-xxx-xxxxx"
	version = "2"
}
```

## Assumptions

1. TODO

## Notes

1. TODO

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
| policy_template_id | The policy template ID. | `string` | true |
| description | Description of the policy template. This is shown to users in the enterprise account. Use this to describe the purpose or context of the policy for enterprise users managing IAM templates. | `string` | false |
| committed | Committed status of the template version. | `bool` | false |
| policy | The core set of properties associated with the template's policy objet. | `` | true |
| template_id | The policy template ID. | `string` | true |
| version | The policy template version. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| policy_template | policy_template object |
| policy_template | policy_template object |
