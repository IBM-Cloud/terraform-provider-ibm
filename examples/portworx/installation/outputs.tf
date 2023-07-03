output "release_pod_status" {
  value = data.helm_release_status.test.pod
}
output "release_service_status" {
  value = data.helm_release_status.test.service
}
