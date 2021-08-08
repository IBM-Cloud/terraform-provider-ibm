// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

func dataSourceIbmCloudantReplication() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmCloudantReplicationRead,

		Schema: map[string]*schema.Schema{
			"doc_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: InvokeValidator("ibm_cloudant_replication", "doc_id"),
				Description:  "Path parameter to specify the document ID.",
			},
			"cloudant_guid": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cloudant Instance GUID.",
			},
			"rev": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Query parameter to specify a document revision.",
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"replication_document": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "HTTP request body for replication operations.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Document ID.",
						},
						"local_seq": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Document's update sequence in current database. Available if requested with local_seq=true query parameter.",
						},
						"rev": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Schema for a document revision identifier.",
						},
						"cancel": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Cancels the replication.",
						},
						"checkpoint_interval": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Defines replication checkpoint interval in milliseconds.",
						},
						"connection_timeout": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "HTTP connection timeout per replication. Even for very fast/reliable networks it might need to be increased if a remote database is too busy.",
						},
						"continuous": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Configure the replication to be continuous.",
						},
						"create_target": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Creates the target database. Requires administrator privileges on target server.",
						},
						"create_target_params": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Request parameters to use during target database creation.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"n": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Schema for the number of replicas of a database in a cluster.",
									},
									"partitioned": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Parameter to specify whether to enable database partitions when creating the target database.",
									},
									"q": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Schema for the number of shards in a database. Each shard is a partition of the hash value range.",
									},
								},
							},
						},
						"doc_ids": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Schema for a list of document IDs.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"filter": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of a filter function which is defined in a design document in the source database in {ddoc_id}/{filter} format. It determines which documents get replicated. Using the selector option provides performance benefits when compared with using the filter option. Use the selector option when possible.",
						},
						"http_connections": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum number of HTTP connections per replication.",
						},
						"query_params": &schema.Schema{
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Schema for a map of string key value pairs, such as query parameters.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"retries_per_request": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of times a replication request is retried. The requests are retried with a doubling exponential backoff starting at 0.25 seconds, with a cap at 5 minutes.",
						},
						"selector": &schema.Schema{
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "JSON object describing criteria used to select documents. The selector specifies fields in the document, and provides an expression to evaluate with the field content or other data.The selector object must:  * Be structured as valid JSON.  * Contain a valid query expression.Using a selector is significantly more efficient than using a JavaScript filter function, and is the recommended option if filtering on document attributes only.Elementary selector syntax requires you to specify one or more fields, and the corresponding values required for those fields. You can create more complex selector expressions by combining operators.Operators are identified by the use of a dollar sign `$` prefix in the name field.There are two core types of operators in the selector syntax:* Combination operators: applied at the topmost level of selection. They are used to combine selectors. In addition to the common boolean operators (`$and`, `$or`, `$not`, `$nor`) there are three combination operators: `$all`, `$elemMatch`, and `$allMatch`. A combination operator takes a single argument. The argument is either another selector, or an array of selectors.* Condition operators: are specific to a field, and are used to evaluate the value stored in that field. For instance, the basic `$eq` operator matches when the specified field contains a value that is equal to the supplied argument.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"since_seq": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Start the replication at a specific sequence value.",
						},
						"socket_options": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Replication socket options.",
						},
						"source": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Schema for a replication source or target database.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Schema for replication source or target database authentication.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"iam": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Schema for an IAM API key for replication database authentication.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"api_key": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "IAM API key.",
															},
														},
													},
												},
											},
										},
									},
									"url": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Replication database URL.",
									},
								},
							},
						},
						"source_proxy": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Address of a (http or socks5 protocol) proxy server through which replication with the source database should occur.",
						},
						"target": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Schema for a replication source or target database.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Schema for replication source or target database authentication.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"iam": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Schema for an IAM API key for replication database authentication.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"api_key": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "IAM API key.",
															},
														},
													},
												},
											},
										},
									},
									"url": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Replication database URL.",
									},
								},
							},
						},
						"target_proxy": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Address of a (http or socks5 protocol) proxy server through which replication with the target database should occur.",
						},
						"use_checkpoints": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specify if checkpoints should be saved during replication. Using checkpoints means a replication can be efficiently resumed.",
						},
						"user_ctx": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Schema for the user context of a session.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"db": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Database name in the context of the provided operation.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "User name.",
									},
									"roles": &schema.Schema{
										Type:        schema.TypeList,
										Required:    true,
										Description: "List of user roles.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
						"worker_batch_size": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Controls how many documents are processed. After each batch a checkpoint is written so this controls how frequently checkpointing occurs.",
						},
						"worker_processes": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Controls how many separate processes will read from the changes manager and write to the target. A higher number can improve throughput.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmCloudantReplicationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cloudantClient, err := meta.(ClientSession).CloudantV1()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudantInstId := d.Get("cloudant_guid").(string)
	docID := d.Get("doc_id").(string)
	cUrl, err := getCloudantInstanceUrl(cloudantInstId, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	cloudantClient.Service.Options.URL = cUrl

	getReplicationDocumentOptions := &cloudantv1.GetReplicationDocumentOptions{}
	getReplicationDocumentOptions.SetDocID(docID)

	replicationDocument, response, err := cloudantClient.GetReplicationDocumentWithContext(context, getReplicationDocumentOptions)
	if err != nil {
		log.Printf("[DEBUG] GetReplicationDocumentWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetReplicationDocumentWithContext failed %s\n%s", err, response))
	}

	replicationDocumentMap := resourceIbmCloudantReplicationReplicationDocumentToMap(*replicationDocument)
	d.SetId(docID)
	d.Set("replication_document", []map[string]interface{}{replicationDocumentMap})

	return nil
}
