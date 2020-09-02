provider "tls" {}

resource "tls_private_key" "ca" {
  algorithm = "RSA"
}

resource "tls_self_signed_cert" "ca" {
  key_algorithm         = "RSA"
  private_key_pem       = tls_private_key.ca.private_key_pem
  validity_period_hours = var.ca_cert_validity_period_days * 24
  early_renewal_hours   = var.ca_cert_early_renewal_days * 24
  is_ca_certificate     = true

  allowed_uses = ["digital_signature", "cert_signing", "key_encipherment"]

  dns_names = ["example.com"]

  subject {
    common_name  = "example.com"
    organization = "Example"
  }
}
provider "ibm" {
}
resource "ibm_resource_instance" "cm" {
  name     = var.cms_name
  location = var.region
  service  = "cloudcerts"
  plan     = "free"
}
resource "ibm_certificate_manager_import" "cert" {
  certificate_manager_instance_id = ibm_resource_instance.cm.id
  name                            = var.import_name
  data = {
    content      = tls_self_signed_cert.ca.cert_pem
    priv_key     = tls_private_key.ca.private_key_pem
    intermediate = ""
  }
}

# // template to import existing certificate...
# resource "ibm_certificate_manager_import" "cert" {
#   certificate_manager_instance_id = ibm_resource_instance.cm.id
#   name                            = var.import_name
#   data = {
#     content = file(var.cert_file_path)
#   }
# }

# // template file to generate ssl certificate and key and import generated certificate...
# #null resource for generating ssl key and certificate...
# resource "null_resource" "import" {
#   provisioner "local-exec" {
#     command = <<EOT
# openssl req -x509  \
#           -newkey rsa:1024 \
#           -keyout "${var.ssl_key}" \
#           -out "${var.ssl_cert}" \
#           -days 1 -nodes \
#           -subj "/C=us/ST="${var.ssl_region}"/L=Dal-10/O=IBM/OU=CloudCerts/CN="${var.host}"" 
#       EOT
#   }
# }
# #datasource to read local certificate file...
# data "local_file" "cert" {
#   filename   = "${path.module}/${var.ssl_cert}"
#   depends_on = [null_resource.import]
# }
# #datasource to read local priv_key file...
# data "local_file" "key" {
#   filename   = "${path.module}/${var.ssl_key}"
#   depends_on = [null_resource.import]
# }
# resource "ibm_certificate_manager_import" "cert" {

#   certificate_manager_instance_id = ibm_resource_instance.cm.id
#   name                            = var.import_name
#   data = {
#     content  = data.local_file.cert.content
#     priv_key = data.local_file.key.content
#   }
# }
