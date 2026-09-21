data "ibm_resource_group" "group" {
  name = var.resource_group
}

data "ibm_cis" "cis" {
  name              = var.cis_instance_name
  resource_group_id = data.ibm_resource_group.group.id
}

data "ibm_cis_domain" "cis_domain" {
  cis_id = data.ibm_cis.cis.id
  domain = var.domain
}

# Batch-create DNS records in a single API call
resource "ibm_cis_dns_records_batch" "posts" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.id

  posts {
    name    = "batch-a"
    type    = "A"
    content = "1.2.3.4"
    ttl     = 120
    proxied = false
  }

  posts {
    name    = "batch-txt"
    type    = "TXT"
    content = "v=spf1 include:example.com ~all"
    ttl     = 300
  }

  posts {
    name    = "batch-mx"
    type    = "MX"
    content = "mail.example.com"
    ttl     = 300
    priority = 10
  }
}

# Batch-update DNS records: full replace, partial update, and delete
resource "ibm_cis_dns_records_batch" "updates" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.id

  # Full replace of the A record
  puts {
    id      = ibm_cis_dns_records_batch.posts.result_posts[0].id
    name    = "batch-a"
    type    = "A"
    content = "5.6.7.8"
    ttl     = 240
    proxied = false
  }

  # Partial update of the TXT record content only
  patches {
    id      = ibm_cis_dns_records_batch.posts.result_posts[1].id
    content = "v=spf1 include:updated.com ~all"
  }

  # Delete the MX record
  deletes {
    id = ibm_cis_dns_records_batch.posts.result_posts[2].id
  }
}
