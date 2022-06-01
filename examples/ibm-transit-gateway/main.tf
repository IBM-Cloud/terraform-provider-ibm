# Create a transit gateway
resource "ibm_tg_gateway" "new_tg_gw"{
        name=var.name
        location=var.location
        global=true
//      resource_group = data.ibm_resource_group.rg.id
}
resource "ibm_is_vpc" "test_tg_vpc" {
  name = var.vpc_name
}
# Add connection to a Transit Gateway
resource "ibm_tg_connection" "test_ibm_tg_connection"{
                gateway = "${ibm_tg_gateway.new_tg_gw.id}"
                network_type = var.network_type
                name = var.vc_name
                network_id = ibm_is_vpc.test_tg_vpc.resource_crn
}
# Add a prefix filter to a Transit Gateway connection
resource "ibm_tg_connection_prefix_filter" "test_tg_prefix_filter" {
                gateway = ibm_tg_gateway.new_tg_gw.id
                connection_id = ibm_tg_connection.test_ibm_tg_connection.connection_id
                action = "permit"
                prefix = "192.168.100.0/24"
                le = "0"
                ge = "32"
}
/*
# Create a transit gateway cross account connection
resource "ibm_tg_connection" "test_tg_cross_connection"{
                gateway = "${ibm_tg_gateway.new_tg_gw.id}"
                network_type = var.network_type
                name= var.vc_name
                # vpc crn from other account
                network_id = var.network_id
                network_account_id = var.network_account_id
}

# Create a transit gateway gre_tunnel connection
resource "ibm_tg_connection" "test_tg_gre_connection"{
                gateway = "${ibm_tg_gateway.new_tg_gw.id}"
                network_type = var.network_type
                name= var.vc_name
                # ID of the classic connection 
                base_connection_id = "ibm_tg_connection.classic_connection.id"
                remote_bgp_asn = "65010"
                local_gateway_ip = "192.168.100.1"
                local_tunnel_ip = "192.168.101.1"
                remote_gateway_ip = "10.242.63.12"
                remote_tunnel_ip = "192.168.101.2"
                zone = "us-south"
}

# Create a transit gateway directlink connection
resource "ibm_tg_connection" "test_tg_dl_connection"{
                gateway = "${ibm_tg_gateway.new_tg_gw.id}"
                network_type = var.network_type
                name= var.vc_name
                # directlink gateway crn
                network_id = var.network_id
}

# Create a transit gateway route report
resource ibm_tg_route_report" "test_tg_route_report" {
	gateway = ibm_tg_gateway.new_tg_gw.id
}

# Retrieves specified Transit Gateway
data "ibm_tg_gateway" "tg_gateway" {
        name= ibm_tg_gateway.new_tg_gw.name
}
# List all the Transit Gateways in the account.
data "ibm_tg_gateways" "all_tg_gws"{
}
# List all locations that support Transit Gateways
data "ibm_tg_locations" "tg_locations" {
}
# Get the details of a Transit Gateway Location.
data "ibm_tg_location" "tg_location" {
        name = "us-south"
}

# List all prefix filters for a Transit Gateway Connection
data "ibm_tg_connection_prefix_filters" "tg_prefix_filters" {
    gateway = ibm_tg_gateway.new_tg_gw.id
    connection_id = ibm_tg_connection.test_ibm_tg_connection.connection_id
}

# Retrieve specified Transit Gateway Connection Prefix Filter
data "ibm_tg_connection_prefix_filter" "tg_prefix_filter" {
    gateway = ibm_tg_gateway.new_tg_gw.id
    connection_id = ibm_tg_connection.test_ibm_tg_connection.connection_id
	filter_id = ibm_tg_connection_prefix_filter.test_tg_prefix_filter.filter_id
}

# List all route reports for a Transit Gateway
data "ibm_tg_route_reports" "tg_route_reports" {
	gateway = ibm_tg_gateway.new_tg_gw.id
}

# Retrieve specified Transit Gateway Route report
data "ibm_tg_route_report" "tg_route_report" {
	gateway = ibm_tg_gateway.new_tg_gw.
	route_report = ibm_tg_route_report_test_tg_route_report.route_report_id
}

*/