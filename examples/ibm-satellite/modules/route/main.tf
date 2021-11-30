resource "null_resource" "get_oc_token" {
  count = var.is_endpoint_provision ? 1 : 0

  provisioner "local-exec" {
    when        = create
    interpreter = ["/bin/bash", "-c"]
    command     = <<-EOT
      sleep 200
      curl -u "apikey:${var.ibmcloud_api_key}" -H "X-CSRF-Token: a" "$(curl ${var.cluster_master_url}/.well-known/oauth-authorization-server | jq -r .issuer)/oauth/authorize?client_id=openshift-challenging-client&response_type=token" -vvv &> /dev/stdout | tee -a resp.log
      token=$(awk -v FS="(#access_token=|&expires_in)" '{print $2}' resp.log)
      echo $token > token.log
      rm -f resp.log
    EOT
  }

  provisioner "local-exec" {
    when    = destroy
    command = "rm -f token.log"
  }
}

data "local_file" "token_file" {
  count    = var.is_endpoint_provision ? 1 : 0
  filename = "token.log"

  depends_on = [null_resource.get_oc_token]
}

// Provision route
resource "restapi_object" "create_route" {
  count = var.is_endpoint_provision ? 1 : 0

  object_id = var.route_name
  path      = "/apis/route.openshift.io/v1/namespaces/default/routes"
  data      = "{ \"kind\":\"Route\",\"apiVersion\": \"route.openshift.io/v1\", \"metadata\": { \"name\": \"${var.route_name}\", \"creationTimestamp\":null }, \"spec\": { \"to\": { \"kind\": \"\", \"name\": \"nginx-service\", \"weight\": null}, \"port\": { \"targetPort\": \"https\"}, \"tls\": { \"termination\": \"passthrough\"}}, \"status\":{} }"

  depends_on = [null_resource.get_oc_token]
}

output "token_data" {
  value = chomp(format("Bearer %v", data.local_file.token_file.*.content))
}