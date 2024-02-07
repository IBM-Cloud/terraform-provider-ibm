# IBM Planning Analytics examples

This example shows 1 usage scenario.

#### Scenario 1: Create a Planning Analytics service instance.

```terraform
resource "ibm_resource_instance" "pa_automation_instance" {
  name              = "planning-analytics-instance-1"
  service           = "planning-analytics"
  plan              = "enterprise"
  location          = "global" 
  resource_group_id = data.ibm_resource_group.group.id
  parameters_json = <<PARAMETERS_JSON
    {
      "sublocation": "us-east-2",
      "planning-analytics": {
        "quota": {
          "memory": 16,
          "storage": 200,
          "users" : 5
        }
      }
    }
  PARAMETERS_JSON
}

To scale the resources assigned to your instance, update the `parameters_json` field with appropriate values.

```

## Dependencies

- The owner of the `ibmcloud_api_key` has permission to create Planning Analytics instance under specified resource group.

## Configuration

- `ibmcloud_api_key` - An API key for IBM Cloud services. If you don't have one already, go to https://cloud.ibm.com/iam/#/apikeys and create a new key.

## Running the configuration

For planning phase

```bash
terraform init
terraform plan
```

For apply phase

```bash
terraform apply
```

For destroy

```bash
terraform destroy
```
