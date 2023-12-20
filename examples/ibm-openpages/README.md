# OpenPages example

This example shows 1 usage scenario.

#### Scenario 1: Create an OpenPages service instance.

```terraform
resource "ibm_resource_instance" "openpages_instance" {
  name              = "terraform-automation"
  service           = "openpages"
  plan              = "essentials"
  location          = "global"
  resource_group_id = data.ibm_resource_group.default_group.id
  parameters_json   = <<EOF
    {
      "aws_region": "us-east-1",
      "baseCurrency": "USD",
      "selectedSolutions": ["ORM"]
    }
  EOF

  timeouts {
    create = "200m"
  }
}

To provison your instance with the correct configuraiton, update the `parameters_json` field with appropriate values.
```

## Dependencies

- The owner of the `ibmcloud_api_key` has permission to create an OpenPages instance under specified resource group.

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