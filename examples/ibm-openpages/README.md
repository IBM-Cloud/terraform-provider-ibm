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
      "initialContentType": "_no_samples",
      "selectedSolutions": ["ORM"]
    }
  EOF

  timeouts {
    create = "200m"
  }
}

To provison your instance with the correct configuration, update the `parameters_json` field with the appropriate values.
```

## Dependencies

- The owner of the `ibmcloud_api_key` has permission to create an OpenPages instance under the specified resource group.

## Configuration

- `ibmcloud_api_key` - An API key for IBM Cloud services. If you don't have one already, go to https://cloud.ibm.com/iam/#/apikeys and create a new key.

## Running the configuration

For the planning phase

```bash
terraform init
terraform plan
```

For the apply phase

```bash
terraform apply
```

For the destroy

```bash
terraform destroy
```