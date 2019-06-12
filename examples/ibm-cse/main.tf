# Need to run export IC_API_KEY="IBM Cloud API Key" firt
privider "ibm" {}

# Create a CSE instance with the required parameters
resource "ibm_cse_instance" "instance1" {
  service = "terraform-1"
  customer = "customer1"
  service_addresses = ["10.102.33.131", "10.102.33.133"]
  region = "us-south"
  data_centers = ["dal10", "dal13"]
  tcp_ports = [8080, 80]
}

# Create a CSE instance with the all parameters
# For update, only below fields are supported:
#   service_addresses,
#   data_centers,
#   tcp_ports,
#   udp_ports,
#   tcp_range,
#   udp_range,
#   estado_proto,
#   estado_port,
#   estado_path,
#   acl
resource "ibm_cse_instance" "instance2" {
  service = "terraform-2"
  customer = "customer-2"
  service_addresses = ["10.102.33.131"]
  region = "us-south"
  data_centers = ["dal10"]
  tcp_ports = [8080,80]
  udp_ports = [12345]
  tcp_range = "8000-9000"
  udp_range = "50000-51000"
  max_speed = "1g"
  estado_proto = "http"
  estado_port = 80
  estado_path = "/healthcheck"
  dedicated = 0
  multi_tenant = 1
  acl = ["10.10.10.1/24"]
}
