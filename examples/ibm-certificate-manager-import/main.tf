resource "ibm_resource_instance" "cm" {
  name     = "testname"
  location = "us-south"
  service  = "cloudcerts"
  plan     = "free"
}

// template file to generate ssl certificate and key and import generated certificate...
resource "ibm_certificate_manager_import" "cert" {
  certificate_manager_instance_id = "${ibm_resource_instance.cm.id}"
  name                            = "test"

  data = {
    content = "${file(var.certfile_path)}"
  }
}

// template file to generate ssl certificate and key and import generated certificate...
#null resource for generating ssl key and certificate...
resource "null_resource" "import" {
  provisioner "local-exec" {
    command = <<EOT
openssl req -x509  \
          -newkey rsa:1024 \
          -keyout "${var.key}" \
          -out "${var.cert}" \
          -days 1 -nodes \
          -subj "/C=us/ST="${var.region}"/L=Dal-10/O=IBM/OU=CloudCerts/CN="${var.host}"" 
      EOT
  }
}

#datasource to read local certificate file...
data "local_file" "cert" {
  filename   = "${path.module}/${var.cert}"
  depends_on = ["null_resource.import"]
}

#datasource to read local priv_key file...
data "local_file" "key" {
  filename   = "${path.module}/${var.key}"
  depends_on = ["null_resource.import"]
}

resource "ibm_certificate_manager_import" "cert" {
  certificate_manager_instance_id = "${ibm_resource_instance.cm.id}"
  name                            = "test"

  data = {
    content  = "${data.local_file.cert.content}"
    priv_key = "${data.local_file.key.content}"
  }
}
