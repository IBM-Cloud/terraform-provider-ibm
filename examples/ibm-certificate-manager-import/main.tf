resource "ibm_resource_instance" "cm" {
  name     = "testname"
  location = "us-south"
  service  = "cloudcerts"
  plan     = "free"
}

resource "ibm_certificate_manager_import" "cert" {
  certificate_manager_instance_id = "${ibm_resource_instance.cm.id}"
  name                            = "test"

  data = {
    content = "${file(var.certfile_path)}"
  }
}

