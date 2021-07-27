data "ibm_resource_group" "rg" {
	name = var.resource_group
}

resource "ibm_satellite_cluster" "create_cluster" {
	name                   = var.cluster
	location               = var.location
	resource_group_id      = data.ibm_resource_group.rg.id
	enable_config_admin    = true
	kube_version           = var.kube_version
	wait_for_worker_update = true
	worker_count		   = var.worker_count 
	host_labels        	   = var.host_labels
	
	dynamic "zones" {
		for_each = var.zones
		content {
			id	= zones.value
		}
	}

	default_worker_pool_labels = var.default_wp_labels
	tags = var.cluster_tags
}

data "ibm_satellite_cluster" "read_cluster" {
	name	= ibm_satellite_cluster.create_cluster.id
}

resource "ibm_satellite_cluster_worker_pool" "create_cluster_wp" {
	name               = var.worker_pool_name
	cluster	           = data.ibm_satellite_cluster.read_cluster.id
	resource_group_id  = data.ibm_resource_group.rg.id
	worker_count       = var.worker_count 
	host_labels        = var.host_labels

	dynamic "zones" {
	for_each = var.zones
		content {
			id	= zones.value
		}
	}

	worker_pool_labels = var.workerpool_labels
}	

data "ibm_satellite_cluster_worker_pool" "read_cluster_wp" {
	name	= ibm_satellite_cluster_worker_pool.create_cluster_wp.name
	cluster	= data.ibm_satellite_cluster.read_cluster.id

	depends_on = [ibm_satellite_cluster_worker_pool.create_cluster_wp]
}

output "cluster_id" {
  value  = data.ibm_satellite_cluster.read_cluster.id
}

output "worker_pool_id" {
  value  = data.ibm_satellite_cluster_worker_pool.read_cluster_wp.id
}