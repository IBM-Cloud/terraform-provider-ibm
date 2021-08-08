---
layout: "ibm"
page_title: "IBM : cloudant_replication"
description: |-
  Manages cloudant_replication.
subcategory: "Cloudant"
---

# ibm\_cloudant_replication

Provides a resource for cloudant_replication. This allows cloudant_replication to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cloudant_replication" "cloudant_replication_doc" {
  cloudant_guid = var.cloudant_guid
  doc_id        = var.doc_id

  replication_document {
    id            = var.doc_id
    create_target = var.create_target
    continuous    = var.continuous
    cancel        = var.cancel

    source {
      auth {
        iam {
          api_key = var.source_api_key
        }
      }
      url = "https://${var.source_host}/${var.db_name}"
    }

    target {
      auth {
        iam {
          api_key = var.target_api_key
        }
      }
      url = "https://${var.target_host}/${var.db_name}"
    }
  }

}
```

## Argument Reference

The following arguments are supported:

* `cloudant_guid` - (Required, string) Path parameter to specify the cloudant instance GUID.
* `doc_id` - (Required, Forces new resource, string) Path parameter to specify the document ID.
  * Constraints: The value must match regular expression `/[^_].*/`
* `replication_document` - (Required, Forces new resource, List) HTTP request body for replication operations.
  * `attachments` - (Optional, map[string]interface{}) Schema for a map of attachment name to attachment metadata.
  * `conflicts` - (Optional, []interface{}) Schema for a list of document revision identifiers.
  * `deleted` - (Optional, bool) Deletion flag. Available if document was removed.
  * `deleted_conflicts` - (Optional, []interface{}) Schema for a list of document revision identifiers.
  * `id` - (Optional, string) Document ID.
  * `local_seq` - (Optional, string) Document's update sequence in current database. Available if requested with local_seq=true query parameter.
  * `rev` - (Optional, string) Schema for a document revision identifier.
    * Constraints: The value must match regular expression `/^[1-9]\\d*-[\\da-f]{32}$/`
  * `revisions` - (Optional, Revisions) Schema for list of revision information.
  * `revs_info` - (Optional, []interface{}) Schema for a list of objects with information about local revisions and their status.
  * `cancel` - (Optional, bool) Cancels the replication.
  * `checkpoint_interval` - (Optional, int) Defines replication checkpoint interval in milliseconds.
    * Constraints: The minimum value is `0`.
  * `connection_timeout` - (Optional, int) HTTP connection timeout per replication. Even for very fast/reliable networks it might need to be increased if a remote database is too busy.
    * Constraints: The minimum value is `0`.
  * `continuous` - (Optional, bool) Configure the replication to be continuous.
    * Constraints: The default value is `false`.
  * `create_target` - (Optional, bool) Creates the target database. Requires administrator privileges on target server.
    * Constraints: The default value is `false`.
  * `create_target_params` - (Optional, ReplicationCreateTargetParameters) Request parameters to use during target database creation.
  * `doc_ids` - (Optional, []interface{}) Schema for a list of document IDs.
  * `filter` - (Optional, string) The name of a filter function which is defined in a design document in the source database in {ddoc_id}/{filter} format. It determines which documents get replicated. Using the selector option provides performance benefits when compared with using the filter option. Use the selector option when possible.
    * Constraints: The value must match regular expression `/[^\/]+\/.+/`
  * `http_connections` - (Optional, int) Maximum number of HTTP connections per replication.
    * Constraints: The minimum value is `1`.
  * `query_params` - (Optional, map[string]interface{}) Schema for a map of string key value pairs, such as query parameters.
  * `retries_per_request` - (Optional, int) Number of times a replication request is retried. The requests are retried with a doubling exponential backoff starting at 0.25 seconds, with a cap at 5 minutes.
    * Constraints: The minimum value is `0`.
  * `selector` - (Optional, map[string]interface{}) JSON object describing criteria used to select documents. The selector specifies fields in the document, and provides an expression to evaluate with the field content or other data.The selector object must:  * Be structured as valid JSON.  * Contain a valid query expression.Using a selector is significantly more efficient than using a JavaScript filter function, and is the recommended option if filtering on document attributes only.Elementary selector syntax requires you to specify one or more fields, and the corresponding values required for those fields. You can create more complex selector expressions by combining operators.Operators are identified by the use of a dollar sign `$` prefix in the name field.There are two core types of operators in the selector syntax:* Combination operators: applied at the topmost level of selection. They are used to combine selectors. In addition to the common boolean operators (`$and`, `$or`, `$not`, `$nor`) there are three combination operators: `$all`, `$elemMatch`, and `$allMatch`. A combination operator takes a single argument. The argument is either another selector, or an array of selectors.* Condition operators: are specific to a field, and are used to evaluate the value stored in that field. For instance, the basic `$eq` operator matches when the specified field contains a value that is equal to the supplied argument.
  * `since_seq` - (Optional, string) Start the replication at a specific sequence value.
  * `socket_options` - (Optional, string) Replication socket options.
  * `source` - (Required, ReplicationDatabase) Schema for a replication source or target database.
  * `source_proxy` - (Optional, string) Address of a (http or socks5 protocol) proxy server through which replication with the source database should occur.
  * `target` - (Required, ReplicationDatabase) Schema for a replication source or target database.
  * `target_proxy` - (Optional, string) Address of a (http or socks5 protocol) proxy server through which replication with the target database should occur.
  * `use_checkpoints` - (Optional, bool) Specify if checkpoints should be saved during replication. Using checkpoints means a replication can be efficiently resumed.
    * Constraints: The default value is `true`.
  * `user_ctx` - (Optional, UserContext) Schema for the user context of a session.
  * `worker_batch_size` - (Optional, int) Controls how many documents are processed. After each batch a checkpoint is written so this controls how frequently checkpointing occurs.
    * Constraints: The minimum value is `1`.
  * `worker_processes` - (Optional, int) Controls how many separate processes will read from the changes manager and write to the target. A higher number can improve throughput.
    * Constraints: The minimum value is `1`.
* `if_match` - (Optional, string) Header parameter to specify the document revision. Alternative to rev query parameter.
* `batch` - (Optional, string) Query parameter to specify whether to store in batch mode. The server will respond with a HTTP 202 Accepted response code immediately.
  * Constraints: Allowable values are: ok
* `new_edits` - (Optional, bool) Query parameter to specify whether to prevent insertion of conflicting document revisions. If false, a well-formed _rev must be included in the document. False is used by the replicator to insert documents into the target database even if that leads to the creation of conflicts.
  * Constraints: The default value is `false`.
* `rev` - (Optional, Forces new resource, string) Query parameter to specify a document revision.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the cloudant_replication.
* `version` - Version of the cloudant_replication.

## Import

You can import the `cloudant_replication` resource by using `ID`.
The `ID` property can be formed from `cloudant_guid`, and `doc_id` in the following format:

```
<cloudant_guid>/<doc_id>
```
* `cloudant_guid`: A string. Path parameter to specify the cloudant instance GUID.
* `doc_id`: A string. Path parameter to specify the document ID.

```
$ terraform import ibm_cloudant_replication.cloudant_replication <cloudant_guid>/<doc_id>
```
