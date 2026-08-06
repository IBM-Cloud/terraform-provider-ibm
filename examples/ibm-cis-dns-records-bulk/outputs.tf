output "result_posts" {
  description = "DNS records created by the batch post operation"
  value       = ibm_cis_dns_records_batch.posts.result_posts
}

output "result_puts" {
  description = "DNS records replaced by the batch put operation"
  value       = ibm_cis_dns_records_batch.updates.result_puts
}

output "result_patches" {
  description = "DNS records updated by the batch patch operation"
  value       = ibm_cis_dns_records_batch.updates.result_patches
}

output "result_deletes" {
  description = "DNS records removed by the batch delete operation"
  value       = ibm_cis_dns_records_batch.updates.result_deletes
}
