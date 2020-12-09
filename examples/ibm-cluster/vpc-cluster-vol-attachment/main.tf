resource "ibm_container_vpc_worker_volume" "volume_attach"{
    volume = var.volume_id
    cluster = var.cluster_id
    worker =var.worker_id
}