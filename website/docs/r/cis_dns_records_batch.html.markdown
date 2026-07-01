---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_dns_records_batch"
description: |-
  Provides a IBM CIS DNS records batch resource.
---

# ibm_cis_dns_records_batch

Provides an IBM Cloud Internet Services DNS records batch resource. This resource allows you to create, update, and delete multiple DNS records in a single API call, improving efficiency when managing large numbers of DNS records. This resource is associated with an IBM Cloud Internet Services instance and a CIS domain resource. For more information, about CIS DNS records, see [managing DNS records](https://cloud.ibm.com/docs/cis?topic=cis-set-up-your-dns-for-cis).

## Example usage

### Create multiple DNS records

```terraform
data "ibm_resource_group" "group" {
  name = "Default"
}

data "ibm_cis" "cis" {
  name              = "my-cis-instance"
  resource_group_id = data.ibm_resource_group.group.id
}

data "ibm_cis_domain" "cis_domain" {
  cis_id = data.ibm_cis.cis.id
  domain = "example.com"
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
```

### Update and delete DNS records

```terraform
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
```

## Argument reference

Review the argument references that you can specify for your resource.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Required, String) The ID of the domain to add the DNS records.
- `posts` - (Optional, List) A list of DNS records to create. Each record supports the following attributes:
  - `name` - (Optional, String) The name of the DNS record.
  - `type` - (Optional, String) The type of the DNS record (e.g., A, AAAA, CNAME, TXT, MX, etc.).
  - `content` - (Optional, String) The value of the DNS record.
  - `ttl` - (Optional, Integer) The TTL (Time to Live) of the DNS record in seconds.
  - `priority` - (Optional, Integer) The priority of the record (for MX and SRV records).
  - `proxied` - (Optional, Boolean) Whether the record is proxied through CIS.
  - `data` - (Optional, Map) Additional data for the DNS record (for SRV, LOC, and CAA records).
- `puts` - (Optional, List) A list of DNS records to fully replace. Each record requires all fields and supports:
  - `id` - (Required, String) The ID of the DNS record to replace.
  - `name` - (Required, String) The name of the DNS record.
  - `type` - (Required, String) The type of the DNS record.
  - `content` - (Required, String) The value of the DNS record.
  - `ttl` - (Required, Integer) The TTL of the DNS record in seconds.
  - `priority` - (Optional, Integer) The priority of the record.
  - `proxied` - (Optional, Boolean) Whether the record is proxied through CIS.
  - `data` - (Optional, Map) Additional data for the DNS record.
- `patches` - (Optional, List) A list of DNS records to partially update. Only specified fields will be updated:
  - `id` - (Required, String) The ID of the DNS record to update.
  - `name` - (Optional, String) The name of the DNS record.
  - `type` - (Optional, String) The type of the DNS record.
  - `content` - (Optional, String) The value of the DNS record.
  - `ttl` - (Optional, Integer) The TTL of the DNS record in seconds.
  - `priority` - (Optional, Integer) The priority of the record.
  - `proxied` - (Optional, Boolean) Whether the record is proxied through CIS.
  - `data` - (Optional, Map) Additional data for the DNS record.
- `deletes` - (Optional, List) A list of DNS record IDs to delete:
  - `id` - (Required, String) The ID of the DNS record to delete.

## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your resource is created.

- `id` - (String) The resource ID, a combination of `<domain_id>:<cis_id>`.
- `result_posts` - (List) The DNS records created by the batch post operation. Each record contains:
  - `id` - (String) The unique identifier of the DNS record.
  - `name` - (String) The name of the DNS record.
  - `type` - (String) The type of the DNS record.
  - `content` - (String) The value of the DNS record.
  - `ttl` - (Integer) The TTL of the DNS record.
  - `proxied` - (Boolean) Whether the record is proxied through CIS.
  - `proxiable` - (Boolean) Whether the record can be proxied.
  - `created_on` - (String) The timestamp when the record was created.
  - `modified_on` - (String) The timestamp when the record was last modified.
- `result_puts` - (List) The DNS records replaced by the batch put operation. Contains the same attributes as `result_posts`.
- `result_patches` - (List) The DNS records updated by the batch patch operation. Contains the same attributes as `result_posts`.
- `result_deletes` - (List) The DNS records removed by the batch delete operation. Contains the same attributes as `result_posts`.

## Notes

- The batch operations (`posts`, `puts`, `patches`, `deletes`) are executed in a single API call, making it more efficient than creating individual DNS records.
- Use `posts` to create new DNS records.
- Use `puts` to completely replace existing DNS records (all fields must be provided).
- Use `patches` to partially update existing DNS records (only specified fields are updated).
- Use `deletes` to remove DNS records.
- Multiple operation types can be combined in a single resource.
- The resource does not perform a traditional "read" operation; it only executes the batch operations and stores the results.

## Import

The `ibm_cis_dns_records_batch` resource does not support import as it is designed for batch operations rather than managing individual record state.