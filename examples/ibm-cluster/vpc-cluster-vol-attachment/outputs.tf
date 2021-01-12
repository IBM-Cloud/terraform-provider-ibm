output "volume_attachment_id" {
  value = ibm_container_vpc_worker_volume.volume_attach.id
}

output "volume_attachment_status" {
    value = ibm_container_vpc_worker_volume.volume_attach.status
}