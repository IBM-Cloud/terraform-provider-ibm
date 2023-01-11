
data "ibm_resource_group" "group" {
  name = "Default"
}

#### Scenario 1: Create Planning Analytics service instance
resource "ibm_resource_instance" "pa_instance" {
  name              = "terraform-automation"
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
