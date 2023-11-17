# satellite-ibm

Use this terraform automation to set up IBM Cloud satellite location for Virtual Server Instances of IBM VPC Infrastructure.

This example uses below modules to set up the satellite location with IBM environment.

1. [satellite-location](modules/location) This module `creates satellite location` for the specified zone|location|region and `generates script` named addhost.sh in the working directory by performing attach host.The generated script is used by `ibm_is_instance` as `user_data` and runs the script. At this stage all the VMs that has run addhost.sh will be attached to the satellite location and will be in unassigned state.
2. [satellite-host](modules/host) This module assigns 3 host to setup the location control plane.
3. [satellite-cluster](modules/cluster) This module will create ROKS satellite cluster and worker pool, attach zone to worker pool.
4. [satellite-link](modules/link) This module will create satellite link.
5. [satellite-route](modules/route) This module will create openshift route.
6. [satellite-endpoint](modules/endpoint) This module will create satellite endpoint.
7. [satellite-dns](modules/dns) This module will register public IPs to control plane & open-shit cluster subdomain DNS records.
8. [satellite-storage-configuration](modules/configuration) This module will create and manage storage configurations in your satellite location.
9. [satellite-storage-assignment](modules/assignment) This module will assign your storage configurations to clusters or cluster groups.
 
## Usage

```
terraform init
```
```
terraform plan
```
```
terraform apply
```
```
terraform destroy
```
## Example Usage
``` hcl

module "satellite-location" {
  source            = "./modules/location"

  is_location_exist = var.is_location_exist
  location          = var.location
  managed_from      = var.managed_from
  location_zones    = var.location_zones
  location_bucket   = var.location_bucket
  ibm_region        = var.ibm_region
  resource_group    = var.resource_group
  host_labels       = var.host_labels
  tags              = var.tags
}

module "satellite-host" {
  source            = "./modules/host"

  host_count        = var.host_count
  location          = module.satellite-location.location_id
  host_vms          = ibm_is_instance.satellite_instance[*].name
  location_zones    = var.location_zones
  host_labels       = var.host_labels
  host_provider     = "ibm"
}

module "satellite-cluster" {
  source               = "./modules/cluster"

  cluster              = var.cluster
  location             = module.satellite-host.location
  kube_version         = var.kube_version
  default_wp_labels    = var.default_wp_labels
  zones                = var.cluster_zones
  resource_group       = var.resource_group
  worker_pool_name     = var.worker_pool_name
  worker_count         = var.worker_count
  workerpool_labels    = var.workerpool_labels
  cluster_tags         = var.cluster_tags
  host_labels          = var.host_labels
  zone_name            = var.zone_name
}

module "satellite-dns" {
  source = "./modules/dns"

  location          = var.location
  cluster           = var.cluster_server_url
  control_plane_ips = var.control_plane_ips
  cluster_ips       = var.cluster_ips
}

module "satellite-link" {
  source = "./modules/link"

  location       = var.location
  crn            = var.location_crn
}

module "satellite-route" {
  source = "./modules/route"

  ibmcloud_api_key   = var.ibmcloud_api_key
  cluster_master_url = var.cluster_master_url
  route_name         = var.route_name
}

module "satellite-endpoint" {
  source = "./modules/endpoint"

  location           = var.location
  connection_type    = var.connection_type
  display_name       = var.display_name
  server_host        = var.server_host
  server_port        = var.server_port
  sni                = var.sni
  client_protocol    = var.client_protocol
  client_mutual_auth = var.client_mutual_auth
  server_protocol    = var.server_protocol
  server_mutual_auth = var.server_mutual_auth
  reject_unauth      = var.reject_unauth
  timeout            = var.timeout
  created_by         = var.created_by
  client_certificate = var.client_certificate
} 

module "satellite-storage-configuration" {
  source = "./modules/configuration"

  location = var.location
  config_name = var.config_name
  storage_template_name = var.storage_template_name
  storage_template_version = var.storage_template_version
  user_config_parameters = var.user_config_parameters
  user_secret_parameters = var.user_secret_parameters
  storage_class_parameters = var.storage_class_parameters
}

module "satellite-storage-assignment"{
  source = "./modules/assignment"

  assignment_name = var.assignment_name
  cluster = var.cluster
  config = var.config
  controller = var.controller
}
```

## Note

* `satellite-location` module creates new location or use existing location ID to process.
   If user pass the location name which is already exist, `satellite-location` module will error out and exit the module.
   In such cases user has to set `is_location_exist` value to true. so that module will use existing location for processing.
* `satellite-route` and `satellite-endpoint` module creates OC route and satellite endpoint.
   If user provisioning route & endpoint make sure ROKS cluster is accessible.


<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name                           | Description                                                       | Type     | Default | Required |
|--------------------------------|-------------------------------------------------------------------|----------|---------|----------|
| ibmcloud_api_key               | IBM Cloud API Key.                                                | string   | n/a     | yes      |
| resource_group                 | Resource Group Name that has to be targeted.                      | string   | n/a     | yes      |
| ibm_region                     | The location or the region in which VM instance exists.           | string   | us-east | no       |
| location                       | Name of the Location that has to be created                       | string   | satellite-ibm      | no  |
| managed_from                   | The IBM Cloud region to manage your Satellite location from.      | string   | n/a     | yes      |
| location_zones                 | Allocate your hosts across these three zones                      | list     | ["us-east-1", "us-east-2", "us-east-3"]     | no       |
| location_bucket                | COS bucket name                                                   | string   | n/a     | no       |
| is_location_exist              | Determines if the location has to be created or not               | bool     | false   | yes      |
| is_prefix                      | Prefix to the Names of all VSI Resources                          | string   | n/a     | yes      |
| public_key                     | Public SSH key used to provision Host/VSI                         | string   | n/a     | no       |
| cluster                        | The name for the new IBM Cloud Satellite cluster                  | string   | satellite-ibm-cluster  | no |
| cluster_zones                  | cluster zones                                                     | list     | ["us-east-1", "us-east-2", "us-east-3"]     | no       |
| kube_version                   | The OpenShift Container Platform version                          | string   | 4.6_openshift     | no       |
| default_wp_labels              | Labels on the default worker pool                                 | map      | n/a     | no       |
| worker_pool_name               | Public SSH key used to provision Host/VSI                         | string   | satellite-ibm-cluster-wp     | no       |
| workerpool_labels              | Labels on the worker pool                                         | map      | n/a     | no       |
| zone_name                      | creates new zone on workerpool                                    | string   | n/a     | no       |
| cluster_tags                   | List of tags for the cluster resource                             | list     | [ "env:cluster" ]     | no       |
| connection_type                | The type of the endpoint.                                         | string   | n/a     | no |
| is_endpoint_provision          | Determines if the route and endpoint has to be created or not.    | bool   | false |   no      |
| display_name                   | The display name of the endpoint. Endpoint names must start with a letter and end with an alphanumeric character, can contain letters, numbers, and hyphen (-), and must be 63 characters or fewer. | string | n/a | yes |
| server_host                    | The host name or IP address of the server endpoint. For 'http-tunnel' protocol, server_host can start with '*.' , which means a wildcard to it's sub domains. Such as '*.example.com' can accept request to 'api.example.com' and 'www.example.com'. | string | n/a | no |
| server_port                    | The port number of the server endpoint. For 'http-tunnel' protocol, server_port can be 0, which means any port. Such as 0 is good for 80 (http) and 443 (https). | number | 443 | no |
| sni                            | The server name indicator (SNI) which used to connect to the server endpoint. Only useful if server side requires SNI. | string | n/a | no |
| client_protocol                | The protocol in the client application side. | string | n/a | no |
| client_mutual_auth             | Whether enable mutual auth in the client application side, when client_protocol is 'tls' or 'https', this field is required. | bool | n/a | no |
| server_protocol                | The protocol in the server application side. This parameter will change to default value if it is omitted even when using PATCH API. If client_protocol is 'udp', server_protocol must be 'udp'. If client_protocol is 'tcp'/'http', server_protocol could be 'tcp'/'tls' and default to 'tcp'. If client_protocol is 'tls'/'https', server_protocol could be 'tcp'/'tls' and default to 'tls'. If client_protocol is 'http-tunnel', server_protocol must be 'tcp'. | string | n/a | no |
| server_mutual_auth             | Whether enable mutual auth in the server application side, when client_protocol is 'tls', this field is required. | bool | n/a | no |
| reject_unauth                  | Whether reject any connection to the server application which is not authorized with the list of supplied CAs in the fields certs.server_cert. | bool | n/a | no |
| timeout                        | The inactivity timeout in the Endpoint side.                                      | number | n/a | no  |
| created_by                     | The service or person who created the endpoint. Must be 1000 characters or fewer. | string | n/a | no |
| client_certificate             | The certs.                                                                        | string | n/a | no |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->