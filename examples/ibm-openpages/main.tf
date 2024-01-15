data "ibm_resource_group" "default_group" {
  is_default = true
}

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