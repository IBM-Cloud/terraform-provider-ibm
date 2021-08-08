// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

func resourceIbmCloudantReplication() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmCloudantReplicationCreate,
		ReadContext:   resourceIbmCloudantReplicationRead,
		DeleteContext: resourceIbmCloudantReplicationDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"doc_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: InvokeValidator("ibm_cloudant_replication", "doc_id"),
				Description:  "Path parameter to specify the document ID.",
			},
			"cloudant_guid": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Cloudant Instance GUID.",
			},
			"replication_document": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "HTTP request body for replication operations.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Document ID.",
						},
						"local_seq": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Document's update sequence in current database. Available if requested with local_seq=true query parameter.",
						},
						"rev": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Schema for a document revision identifier.",
						},
						"cancel": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Cancels the replication.",
						},
						"checkpoint_interval": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Defines replication checkpoint interval in milliseconds.",
						},
						"connection_timeout": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "HTTP connection timeout per replication. Even for very fast/reliable networks it might need to be increased if a remote database is too busy.",
						},
						"continuous": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Configure the replication to be continuous.",
						},
						"create_target": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Creates the target database. Requires administrator privileges on target server.",
						},
						"create_target_params": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Request parameters to use during target database creation.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"n": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Schema for the number of replicas of a database in a cluster.",
									},
									"partitioned": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "Parameter to specify whether to enable database partitions when creating the target database.",
									},
									"q": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Schema for the number of shards in a database. Each shard is a partition of the hash value range.",
									},
								},
							},
						},
						"doc_ids": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "Schema for a list of document IDs.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"filter": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of a filter function which is defined in a design document in the source database in {ddoc_id}/{filter} format. It determines which documents get replicated. Using the selector option provides performance benefits when compared with using the filter option. Use the selector option when possible.",
						},
						"http_connections": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Maximum number of HTTP connections per replication.",
						},
						"query_params": &schema.Schema{
							Type:        schema.TypeMap,
							Optional:    true,
							Computed:    true,
							Description: "Schema for a map of string key value pairs, such as query parameters.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"retries_per_request": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Number of times a replication request is retried. The requests are retried with a doubling exponential backoff starting at 0.25 seconds, with a cap at 5 minutes.",
						},
						"selector": &schema.Schema{
							Type:        schema.TypeMap,
							Optional:    true,
							Computed:    true,
							Description: "JSON object describing criteria used to select documents. The selector specifies fields in the document, and provides an expression to evaluate with the field content or other data.The selector object must:  * Be structured as valid JSON.  * Contain a valid query expression.Using a selector is significantly more efficient than using a JavaScript filter function, and is the recommended option if filtering on document attributes only.Elementary selector syntax requires you to specify one or more fields, and the corresponding values required for those fields. You can create more complex selector expressions by combining operators.Operators are identified by the use of a dollar sign `$` prefix in the name field.There are two core types of operators in the selector syntax:* Combination operators: applied at the topmost level of selection. They are used to combine selectors. In addition to the common boolean operators (`$and`, `$or`, `$not`, `$nor`) there are three combination operators: `$all`, `$elemMatch`, and `$allMatch`. A combination operator takes a single argument. The argument is either another selector, or an array of selectors.* Condition operators: are specific to a field, and are used to evaluate the value stored in that field. For instance, the basic `$eq` operator matches when the specified field contains a value that is equal to the supplied argument.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"since_seq": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Start the replication at a specific sequence value.",
						},
						"socket_options": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Replication socket options.",
						},
						"source": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "Schema for a replication source or target database.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Schema for replication source or target database authentication.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"iam": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
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
										Required:    true,
										Description: "Replication database URL.",
									},
								},
							},
						},
						"source_proxy": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Address of a (http or socks5 protocol) proxy server through which replication with the source database should occur.",
						},
						"target": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "Schema for a replication source or target database.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Schema for replication source or target database authentication.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"iam": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
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
										Required:    true,
										Description: "Replication database URL.",
									},
								},
							},
						},
						"target_proxy": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Address of a (http or socks5 protocol) proxy server through which replication with the target database should occur.",
						},
						"use_checkpoints": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "Specify if checkpoints should be saved during replication. Using checkpoints means a replication can be efficiently resumed.",
						},
						"user_ctx": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Schema for the user context of a session.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"db": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
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
							Optional:    true,
							Description: "Controls how many documents are processed. After each batch a checkpoint is written so this controls how frequently checkpointing occurs.",
						},
						"worker_processes": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Controls how many separate processes will read from the changes manager and write to the target. A higher number can improve throughput.",
						},
					},
				},
			},
			"if_match": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Header parameter to specify the document revision. Alternative to rev query parameter.",
			},
			"batch": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: InvokeValidator("ibm_cloudant_replication", "batch"),
				Description:  "Query parameter to specify whether to store in batch mode. The server will respond with a HTTP 202 Accepted response code immediately.",
			},
			"new_edits": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Query parameter to specify whether to prevent insertion of conflicting document revisions. If false, a well-formed _rev must be included in the document. False is used by the replicator to insert documents into the target database even if that leads to the creation of conflicts.",
			},
			"rev": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Query parameter to specify a document revision.",
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIbmCloudantReplicationValidator() *ResourceValidator {
	validateSchema := make([]ValidateSchema, 1)
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "doc_id",
			ValidateFunctionIdentifier: ValidateRegexp,
			Type:                       TypeString,
			Required:                   true,
			Regexp:                     `[^_].*`,
		},
		ValidateSchema{
			Identifier:                 "batch",
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Optional:                   true,
			AllowedValues:              "ok",
		},
	)

	resourceValidator := ResourceValidator{ResourceName: "ibm_cloudant_replication", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmCloudantReplicationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cloudantClient, err := meta.(ClientSession).CloudantV1()
	if err != nil {
		return diag.FromErr(err)
	}

	instanceId := d.Get("cloudant_guid").(string)
	cUrl, err := getCloudantInstanceUrl(instanceId, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	cloudantClient.Service.Options.URL = cUrl

	putReplicationDocumentOptions := &cloudantv1.PutReplicationDocumentOptions{}
	putReplicationDocumentOptions.SetDocID(d.Get("doc_id").(string))

	if replicationDoc, ok := d.GetOk("replication_document"); ok {
		repDocList := replicationDoc.([]interface{})
		for _, l := range repDocList {
			replicationMap, _ := l.(map[string]interface{})
			replicationDocument := resourceIbmCloudantReplicationMapToReplicationDocument(replicationMap)
			putReplicationDocumentOptions.SetReplicationDocument(&replicationDocument)
		}

	}

	if _, ok := d.GetOk("batch"); ok {
		putReplicationDocumentOptions.SetBatch(d.Get("batch").(string))
	}
	if _, ok := d.GetOk("new_edits"); ok {
		putReplicationDocumentOptions.SetNewEdits(d.Get("new_edits").(bool))
	}
	if _, ok := d.GetOk("rev"); ok {
		putReplicationDocumentOptions.SetRev(d.Get("rev").(string))
	}

	documentResult, response, err := cloudantClient.PutReplicationDocumentWithContext(context, putReplicationDocumentOptions)
	if err != nil {
		log.Printf("[DEBUG] PutReplicationDocumentWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("PutReplicationDocumentWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", instanceId, *documentResult.ID))

	return resourceIbmCloudantReplicationRead(context, d, meta)
}

func resourceIbmCloudantReplicationMapToReplicationDocument(replicationDocumentMap map[string]interface{}) cloudantv1.ReplicationDocument {

	replicationDocument := cloudantv1.ReplicationDocument{}

	if v, found := replicationDocumentMap["id"]; found {
		replicationDocument.ID = core.StringPtr(v.(string))
	}
	if v, found := replicationDocumentMap["rev"]; found {
		if len(v.(string)) > 0 {
			replicationDocument.Rev = core.StringPtr(v.(string))
		}
	}
	if replicationDocumentMap["source"] != nil {
		replicationDocument.Source = resourceIbmCloudantReplicationMapToReplicationDatabase(replicationDocumentMap["source"])
	}
	if replicationDocumentMap["target"] != nil {
		replicationDocument.Target = resourceIbmCloudantReplicationMapToReplicationDatabase(replicationDocumentMap["target"])
	}
	if v, found := replicationDocumentMap["cancel"]; found {
		replicationDocument.Cancel = core.BoolPtr(v.(bool))
	}
	if v, found := replicationDocumentMap["checkpoint_interval"]; found {
		replicationDocument.CheckpointInterval = core.Int64Ptr(int64(v.(int)))
	}
	if v, found := replicationDocumentMap["connection_timeout"]; found {
		replicationDocument.ConnectionTimeout = core.Int64Ptr(int64(v.(int)))
	}
	if v, found := replicationDocumentMap["continuous"]; found {
		replicationDocument.Continuous = core.BoolPtr(v.(bool))
	}
	if v, found := replicationDocumentMap["create_target"]; found {
		replicationDocument.CreateTarget = core.BoolPtr(v.(bool))
	}
	if v, found := replicationDocumentMap["create_target_params"]; found {
		if reflect.ValueOf(v).IsNil() {
			replicationDocument.CreateTargetParams = resourceIbmCloudantReplicationMapToReplicationCreateTargetParameters(v.(map[string]interface{}))
		}
	}
	if v, found := replicationDocumentMap["doc_ids"]; found {
		if reflect.ValueOf(v).IsNil() {
			docIds := []string{}
			for _, docIdsItem := range v.([]interface{}) {
				docIds = append(docIds, docIdsItem.(string))
			}
			replicationDocument.DocIds = docIds
		}
	}
	if v, found := replicationDocumentMap["filter"]; found {
		if len(v.(string)) > 0 {
			replicationDocument.Filter = core.StringPtr(v.(string))
		}
	}
	if v, found := replicationDocumentMap["http_connections"]; found {
		replicationDocument.HTTPConnections = core.Int64Ptr(int64(v.(int)))
	}
	if v, found := replicationDocumentMap["query_params"]; found {
		if reflect.ValueOf(v).IsNil() {
			replicationDocument.QueryParams = flattenMapInterfaceVal(v.(map[string]interface{}))
		}
	} else {
		replicationDocument.QueryParams = make(map[string]string)
	}
	if v, found := replicationDocumentMap["retries_per_request"]; found {
		replicationDocument.RetriesPerRequest = core.Int64Ptr(int64(v.(int)))
	}
	if v, found := replicationDocumentMap["since_seq"]; found {
		if len(v.(string)) > 0 {
			replicationDocument.SinceSeq = core.StringPtr(v.(string))
		}
	}
	if v, found := replicationDocumentMap["socket_options"]; found {
		if len(v.(string)) > 0 {
			replicationDocument.SocketOptions = core.StringPtr(v.(string))
		}
	}
	if v, found := replicationDocumentMap["socket_options"]; found {
		if len(v.(string)) > 0 {
			replicationDocument.SourceProxy = core.StringPtr(v.(string))
		}
	}
	if v, found := replicationDocumentMap["target_proxy"]; found {
		if len(v.(string)) > 0 {
			replicationDocument.TargetProxy = core.StringPtr(v.(string))
		}
	}
	if v, found := replicationDocumentMap["use_checkpoints"]; found {
		replicationDocument.UseCheckpoints = core.BoolPtr(v.(bool))
	}
	if v, found := replicationDocumentMap["user_ctx"]; found {
		if reflect.ValueOf(v).IsNil() {
			replicationDocument.UserCtx = resourceIbmCloudantReplicationMapToUserContext(v.(map[string]interface{}))
		}
	}
	if v, found := replicationDocumentMap["worker_batch_size"]; found {
		replicationDocument.WorkerBatchSize = core.Int64Ptr(int64(v.(int)))
	}
	if v, found := replicationDocumentMap["worker_processes"]; found {
		replicationDocument.WorkerProcesses = core.Int64Ptr(int64(v.(int)))
	}

	return replicationDocument
}

func resourceIbmCloudantReplicationMapToDocumentRevisionStatus(documentRevisionStatusMap map[string]interface{}) cloudantv1.DocumentRevisionStatus {
	documentRevisionStatus := cloudantv1.DocumentRevisionStatus{}

	documentRevisionStatus.Rev = core.StringPtr(documentRevisionStatusMap["rev"].(string))
	documentRevisionStatus.Status = core.StringPtr(documentRevisionStatusMap["status"].(string))

	return documentRevisionStatus
}

func resourceIbmCloudantReplicationMapToReplicationCreateTargetParameters(replicationCreateTargetParametersMap map[string]interface{}) *cloudantv1.ReplicationCreateTargetParameters {
	replicationCreateTargetParameters := *&cloudantv1.ReplicationCreateTargetParameters{}

	if replicationCreateTargetParametersMap != nil {
		if replicationCreateTargetParametersMap["n"] != nil {
			replicationCreateTargetParameters.N = core.Int64Ptr(int64(replicationCreateTargetParametersMap["n"].(int)))
		}
		if replicationCreateTargetParametersMap["partitioned"] != nil {
			replicationCreateTargetParameters.Partitioned = core.BoolPtr(replicationCreateTargetParametersMap["partitioned"].(bool))
		}
		if replicationCreateTargetParametersMap["q"] != nil {
			replicationCreateTargetParameters.Q = core.Int64Ptr(int64(replicationCreateTargetParametersMap["q"].(int)))
		}
	}

	return &replicationCreateTargetParameters
}

func resourceIbmCloudantReplicationMapToReplicationDatabase(replicationDatabaseMap interface{}) *cloudantv1.ReplicationDatabase {
	replicationDatabase := cloudantv1.ReplicationDatabase{}

	for _, item := range replicationDatabaseMap.([]interface{}) {
		if item.(map[string]interface{})["auth"] != nil {
			replicationDatabase.Auth = resourceIbmCloudantReplicationMapToReplicationDatabaseAuth(item.(map[string]interface{})["auth"])
		}
		replicationDatabase.URL = core.StringPtr(item.(map[string]interface{})["url"].(string))
	}

	return &replicationDatabase
}

func resourceIbmCloudantReplicationMapToReplicationDatabaseAuth(replicationDatabaseAuthMap interface{}) *cloudantv1.ReplicationDatabaseAuth {
	replicationDatabaseAuth := *&cloudantv1.ReplicationDatabaseAuth{}

	for _, item := range replicationDatabaseAuthMap.([]interface{}) {
		if item.(map[string]interface{})["iam"] != nil {
			replicationDatabaseAuth.Iam = resourceIbmCloudantReplicationMapToReplicationDatabaseAuthIam(item.(map[string]interface{})["iam"])
		}
	}

	return &replicationDatabaseAuth
}

func resourceIbmCloudantReplicationMapToReplicationDatabaseAuthIam(replicationDatabaseAuthIamMap interface{}) *cloudantv1.ReplicationDatabaseAuthIam {
	replicationDatabaseAuthIam := *&cloudantv1.ReplicationDatabaseAuthIam{}

	for _, item := range replicationDatabaseAuthIamMap.([]interface{}) {
		if item.(map[string]interface{})["api_key"] != nil {
			apiKey := item.(map[string]interface{})["api_key"].(string)
			replicationDatabaseAuthIam.ApiKey = &apiKey
		}
	}

	return &replicationDatabaseAuthIam
}

func resourceIbmCloudantReplicationMapToUserContext(userContextMap map[string]interface{}) *cloudantv1.UserContext {
	userContext := *&cloudantv1.UserContext{}

	if userContextMap["db"] != nil {
		userContext.Db = core.StringPtr(userContextMap["db"].(string))
	}
	userContext.Name = core.StringPtr(userContextMap["name"].(string))
	roles := []string{}
	for _, rolesItem := range userContextMap["roles"].([]interface{}) {
		roles = append(roles, rolesItem.(string))
	}
	userContext.Roles = roles

	return &userContext
}

func resourceIbmCloudantReplicationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cloudantClient, err := meta.(ClientSession).CloudantV1()
	if err != nil {
		return diag.FromErr(err)
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	instanceId := parts[0]
	cUrl, err := getCloudantInstanceUrl(instanceId, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	cloudantClient.Service.Options.URL = cUrl

	getReplicationDocumentOptions := &cloudantv1.GetReplicationDocumentOptions{}
	getReplicationDocumentOptions.SetDocID(parts[1])

	replicationDocument, response, err := cloudantClient.GetReplicationDocumentWithContext(context, getReplicationDocumentOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetReplicationDocumentWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetReplicationDocumentWithContext failed %s\n%s", err, response))
	}

	d.Set("cloudant_guid", instanceId)
	if err = d.Set("doc_id", *replicationDocument.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting doc_id: %s", err))
	}
	replicationDocumentMap := resourceIbmCloudantReplicationReplicationDocumentToMap(*replicationDocument)
	if err = d.Set("replication_document", []map[string]interface{}{replicationDocumentMap}); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting replication_document: %s", err))
	}

	return nil
}

func resourceIbmCloudantReplicationReplicationDocumentToMap(replicationDocument cloudantv1.ReplicationDocument) map[string]interface{} {
	replicationDocumentMap := map[string]interface{}{}
	replicationDocumentMap["id"] = *replicationDocument.ID

	if replicationDocument.LocalSeq != nil {
		replicationDocumentMap["_local_seq"] = *replicationDocument.LocalSeq
	}
	if replicationDocument.Rev != nil {
		replicationDocumentMap["rev"] = *replicationDocument.Rev
	}
	if replicationDocument.RevsInfo != nil {
		revsInfo := []map[string]interface{}{}
		for _, revsInfoItem := range replicationDocument.RevsInfo {
			revsInfoItemMap := resourceIbmCloudantReplicationDocumentRevisionStatusToMap(revsInfoItem)
			revsInfo = append(revsInfo, revsInfoItemMap)
		}
		replicationDocumentMap["_revs_info"] = revsInfo
	}
	if replicationDocument.Cancel != nil {
		replicationDocumentMap["cancel"] = *replicationDocument.Cancel
	}
	replicationDocumentMap["checkpoint_interval"] = intValue(replicationDocument.CheckpointInterval)
	replicationDocumentMap["connection_timeout"] = intValue(replicationDocument.ConnectionTimeout)
	replicationDocumentMap["continuous"] = *replicationDocument.Continuous
	if replicationDocument.CreateTarget != nil {
		replicationDocumentMap["create_target"] = replicationDocument.CreateTarget
	}
	if replicationDocument.CreateTargetParams != nil {
		CreateTargetParamsMap := resourceIbmCloudantReplicationReplicationCreateTargetParametersToMap(*replicationDocument.CreateTargetParams)
		replicationDocumentMap["create_target_params"] = []map[string]interface{}{CreateTargetParamsMap}
	}
	if replicationDocument.DocIds != nil {
		replicationDocumentMap["doc_ids"] = replicationDocument.DocIds
	}
	if replicationDocument.Filter != nil {
		replicationDocumentMap["filter"] = *replicationDocument.Filter
	}
	if replicationDocument.HTTPConnections != nil {
		replicationDocumentMap["http_connections"] = intValue(replicationDocument.HTTPConnections)
	}
	if replicationDocument.QueryParams != nil {
		replicationDocumentMap["query_params"] = replicationDocument.QueryParams
	}
	if replicationDocument.RetriesPerRequest != nil {
		replicationDocumentMap["retries_per_request"] = intValue(replicationDocument.RetriesPerRequest)
	}
	if replicationDocument.Selector != nil {
		replicationDocumentMap["selector"] = replicationDocument.Selector
	}
	if replicationDocument.SinceSeq != nil {
		replicationDocumentMap["since_seq"] = *replicationDocument.SinceSeq
	}
	if replicationDocument.SocketOptions != nil {
		replicationDocumentMap["socket_options"] = *replicationDocument.SocketOptions
	}
	if replicationDocument.Source != nil {
		SourceMap := resourceIbmCloudantReplicationReplicationDatabaseToMap(*replicationDocument.Source)
		replicationDocumentMap["source"] = []map[string]interface{}{SourceMap}
	}
	if replicationDocument.SourceProxy != nil {
		replicationDocumentMap["source_proxy"] = *replicationDocument.SourceProxy
	}
	if replicationDocument.Target != nil {
		TargetMap := resourceIbmCloudantReplicationReplicationDatabaseToMap(*replicationDocument.Target)
		replicationDocumentMap["target"] = []map[string]interface{}{TargetMap}
	}
	if replicationDocument.TargetProxy != nil {
		replicationDocumentMap["target_proxy"] = *replicationDocument.TargetProxy
	}
	if replicationDocument.UseCheckpoints != nil {
		replicationDocumentMap["use_checkpoints"] = *replicationDocument.UseCheckpoints
	}
	if replicationDocument.UserCtx != nil {
		UserCtxMap := resourceIbmCloudantReplicationUserContextToMap(*replicationDocument.UserCtx)
		replicationDocumentMap["user_ctx"] = []map[string]interface{}{UserCtxMap}
	}
	if replicationDocument.WorkerBatchSize != nil {
		replicationDocumentMap["worker_batch_size"] = intValue(replicationDocument.WorkerBatchSize)
	}
	if replicationDocument.WorkerProcesses != nil {
		replicationDocumentMap["worker_processes"] = intValue(replicationDocument.WorkerProcesses)
	}

	return replicationDocumentMap
}

func resourceIbmCloudantReplicationRevisionsToMap(revisions cloudantv1.Revisions) map[string]interface{} {
	revisionsMap := map[string]interface{}{}

	revisionsMap["ids"] = revisions.Ids
	revisionsMap["start"] = intValue(revisions.Start)

	return revisionsMap
}

func resourceIbmCloudantReplicationDocumentRevisionStatusToMap(documentRevisionStatus cloudantv1.DocumentRevisionStatus) map[string]interface{} {
	documentRevisionStatusMap := map[string]interface{}{}

	documentRevisionStatusMap["rev"] = documentRevisionStatus.Rev
	documentRevisionStatusMap["status"] = documentRevisionStatus.Status

	return documentRevisionStatusMap
}

func resourceIbmCloudantReplicationReplicationCreateTargetParametersToMap(replicationCreateTargetParameters cloudantv1.ReplicationCreateTargetParameters) map[string]interface{} {
	replicationCreateTargetParametersMap := map[string]interface{}{}

	replicationCreateTargetParametersMap["n"] = intValue(replicationCreateTargetParameters.N)
	replicationCreateTargetParametersMap["partitioned"] = replicationCreateTargetParameters.Partitioned
	replicationCreateTargetParametersMap["q"] = intValue(replicationCreateTargetParameters.Q)

	return replicationCreateTargetParametersMap
}

func resourceIbmCloudantReplicationReplicationDatabaseToMap(replicationDatabase cloudantv1.ReplicationDatabase) map[string]interface{} {
	replicationDatabaseMap := map[string]interface{}{}

	if replicationDatabase.Auth != nil {
		AuthMap := resourceIbmCloudantReplicationReplicationDatabaseAuthToMap(*replicationDatabase.Auth)
		replicationDatabaseMap["auth"] = []map[string]interface{}{AuthMap}
	}
	replicationDatabaseMap["url"] = replicationDatabase.URL

	return replicationDatabaseMap
}

func resourceIbmCloudantReplicationReplicationDatabaseAuthToMap(replicationDatabaseAuth cloudantv1.ReplicationDatabaseAuth) map[string]interface{} {
	replicationDatabaseAuthMap := map[string]interface{}{}

	if replicationDatabaseAuth.Iam != nil {
		IamMap := resourceIbmCloudantReplicationReplicationDatabaseAuthIamToMap(*replicationDatabaseAuth.Iam)
		replicationDatabaseAuthMap["iam"] = []map[string]interface{}{IamMap}
	}

	return replicationDatabaseAuthMap
}

func resourceIbmCloudantReplicationReplicationDatabaseAuthIamToMap(replicationDatabaseAuthIam cloudantv1.ReplicationDatabaseAuthIam) map[string]interface{} {
	replicationDatabaseAuthIamMap := map[string]interface{}{}

	replicationDatabaseAuthIamMap["api_key"] = replicationDatabaseAuthIam.ApiKey

	return replicationDatabaseAuthIamMap
}

func resourceIbmCloudantReplicationUserContextToMap(userContext cloudantv1.UserContext) map[string]interface{} {
	userContextMap := map[string]interface{}{}

	userContextMap["db"] = userContext.Db
	userContextMap["name"] = userContext.Name
	userContextMap["roles"] = userContext.Roles

	return userContextMap
}

func resourceIbmCloudantReplicationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cloudantClient, err := meta.(ClientSession).CloudantV1()
	if err != nil {
		return diag.FromErr(err)
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	instanceId := parts[0]
	cUrl, err := getCloudantInstanceUrl(instanceId, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	cloudantClient.Service.Options.URL = cUrl

	getReplicationDocumentOptions := &cloudantv1.GetReplicationDocumentOptions{}
	getReplicationDocumentOptions.SetDocID(parts[1])

	replicationDocument, response, err := cloudantClient.GetReplicationDocumentWithContext(context, getReplicationDocumentOptions)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceIbmCloudantReplicationDelete failed to get Document %s\n%s", err, response))
	}

	deleteReplicationDocumentOptions := &cloudantv1.DeleteReplicationDocumentOptions{}
	deleteReplicationDocumentOptions.SetDocID(parts[1])
	deleteReplicationDocumentOptions.SetRev(*replicationDocument.Rev)

	_, response, err = cloudantClient.DeleteReplicationDocumentWithContext(context, deleteReplicationDocumentOptions)
	if err != nil {
		if response.StatusCode == 404 || response.StatusCode == 409 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] DeleteReplicationDocumentWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteReplicationDocumentWithContext failed %s\n%s", err, response))
	}

	return nil
}
