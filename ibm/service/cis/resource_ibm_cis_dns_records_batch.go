// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/dnsrecordsv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	cisDNSRecordsBatchDeletes = "deletes"
	cisDNSRecordsBatchPatches = "patches"
	cisDNSRecordsBatchPosts   = "posts"
	cisDNSRecordsBatchPuts    = "puts"
)

func ResourceIBMCISDNSRecordsBatch() *schema.Resource {
	recordSchema := map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		cisDNSRecordName: {
			Type:     schema.TypeString,
			Optional: true,
		},
		cisDNSRecordType: {
			Type:     schema.TypeString,
			Optional: true,
		},
		cisDNSRecordTTL: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		cisDNSRecordContent: {
			Type:     schema.TypeString,
			Optional: true,
		},
		cisDNSRecordPriority: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		cisDNSRecordProxied: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		cisDNSRecordData: {
			Type:     schema.TypeMap,
			Optional: true,
		},
	}

	deleteSchema := map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Required: true,
		},
	}

	resultRecordSchema := map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		cisDNSRecordName: {
			Type:     schema.TypeString,
			Computed: true,
		},
		cisDNSRecordType: {
			Type:     schema.TypeString,
			Computed: true,
		},
		cisDNSRecordTTL: {
			Type:     schema.TypeInt,
			Computed: true,
		},
		cisDNSRecordContent: {
			Type:     schema.TypeString,
			Computed: true,
		},
		cisDNSRecordProxied: {
			Type:     schema.TypeBool,
			Computed: true,
		},
		cisDNSRecordProxiable: {
			Type:     schema.TypeBool,
			Computed: true,
		},
		cisDNSRecordCreatedOn: {
			Type:     schema.TypeString,
			Computed: true,
		},
		cisDNSRecordModifiedOn: {
			Type:     schema.TypeString,
			Computed: true,
		},
	}

	return &schema.Resource{
		Create:   resourceIBMCISDNSRecordsBatchCreate,
		Read:     resourceIBMCISDNSRecordsBatchRead,
		Update:   resourceIBMCISDNSRecordsBatchCreate,
		Delete:   resourceIBMCISDNSRecordsBatchDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance CRN",
				Required:    true,
				ValidateFunc: validate.InvokeValidator("ibm_cis_dns_records_batch",
					"cis_id"),
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Description:      "Associated CIS domain",
				Required:         true,
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			cisDNSRecordsBatchDeletes: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: deleteSchema,
				},
			},
			cisDNSRecordsBatchPatches: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: recordSchema,
				},
			},
			cisDNSRecordsBatchPosts: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: recordSchema,
				},
			},
			cisDNSRecordsBatchPuts: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: recordSchema,
				},
			},
			"result_deletes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: resultRecordSchema},
			},
			"result_patches": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: resultRecordSchema},
			},
			"result_posts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: resultRecordSchema},
			},
			"result_puts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: resultRecordSchema},
			},
		},
	}
}

func ResourceIBMCISDNSRecordsBatchValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cis_id",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			CloudDataType:              "resource_instance",
			CloudDataRange:             []string{"service:internet-svcs"},
			Required:                   true})
	return &validate.ResourceValidator{
		ResourceName: "ibm_cis_dns_records_batch",
		Schema:       validateSchema,
	}
}

func resourceIBMCISDNSRecordsBatchCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisDNSRecordClientSession()
	if err != nil {
		return err
	}

	crn := d.Get(cisID).(string)
	zoneID, _, _ := flex.ConvertTftoCisTwoVar(d.Get(cisDomainID).(string))
	sess.Crn = core.StringPtr(crn)
	sess.ZoneIdentifier = core.StringPtr(zoneID)

	opt := sess.NewBatchDnsRecordsOptions()

	if v, ok := d.GetOk(cisDNSRecordsBatchDeletes); ok {
		items := v.([]interface{})
		deletes := make([]dnsrecordsv1.BatchDnsRecordsRequestDeletesItem, 0, len(items))
		for _, item := range items {
			m := item.(map[string]interface{})
			del, err := sess.NewBatchDnsRecordsRequestDeletesItem(m["id"].(string))
			if err != nil {
				return err
			}
			deletes = append(deletes, *del)
		}
		opt.SetDeletes(deletes)
	}

	if v, ok := d.GetOk(cisDNSRecordsBatchPatches); ok {
		items := v.([]interface{})
		patches := make([]dnsrecordsv1.BatchDnsRecordsRequestPatchesItem, 0, len(items))
		for _, item := range items {
			m := item.(map[string]interface{})
			patch, err := sess.NewBatchDnsRecordsRequestPatchesItem(m["id"].(string))
			if err != nil {
				return err
			}
			if v, ok := m[cisDNSRecordName].(string); ok && v != "" {
				patch.Name = core.StringPtr(v)
			}
			if v, ok := m[cisDNSRecordType].(string); ok && v != "" {
				patch.Type = core.StringPtr(v)
			}
			if v, ok := m[cisDNSRecordContent].(string); ok && v != "" {
				patch.Content = core.StringPtr(v)
			}
			if v, ok := m[cisDNSRecordTTL].(int); ok && v != 0 {
				patch.TTL = core.Int64Ptr(int64(v))
			}
			if v, ok := m[cisDNSRecordPriority].(int); ok && v != 0 {
				patch.Priority = core.Int64Ptr(int64(v))
			}
			if v, ok := m[cisDNSRecordProxied].(bool); ok {
				patch.Proxied = core.BoolPtr(v)
			}
			if dataRaw, ok := m[cisDNSRecordData]; ok {
				if dataMap, ok := dataRaw.(map[string]interface{}); ok && len(dataMap) > 0 {
					patch.Data = dataMap
				}
			}
			patches = append(patches, *patch)
		}
		opt.SetPatches(patches)
	}

	if v, ok := d.GetOk(cisDNSRecordsBatchPosts); ok {
		items := v.([]interface{})
		posts := make([]dnsrecordsv1.DnsrecordInput, 0, len(items))
		for _, item := range items {
			m := item.(map[string]interface{})
			post := dnsrecordsv1.DnsrecordInput{}
			if v, ok := m[cisDNSRecordName].(string); ok && v != "" {
				post.Name = core.StringPtr(v)
			}
			if v, ok := m[cisDNSRecordType].(string); ok && v != "" {
				post.Type = core.StringPtr(v)
			}
			if v, ok := m[cisDNSRecordContent].(string); ok && v != "" {
				post.Content = core.StringPtr(v)
			}
			if v, ok := m[cisDNSRecordTTL].(int); ok && v != 0 {
				post.TTL = core.Int64Ptr(int64(v))
			}
			if v, ok := m[cisDNSRecordPriority].(int); ok && v != 0 {
				post.Priority = core.Int64Ptr(int64(v))
			}
			if v, ok := m[cisDNSRecordProxied].(bool); ok {
				post.Proxied = core.BoolPtr(v)
			}
			if dataRaw, ok := m[cisDNSRecordData]; ok {
				if dataMap, ok := dataRaw.(map[string]interface{}); ok && len(dataMap) > 0 {
					post.Data = dataMap
				}
			}
			posts = append(posts, post)
		}
		opt.SetPosts(posts)
	}

	if v, ok := d.GetOk(cisDNSRecordsBatchPuts); ok {
		items := v.([]interface{})
		puts := make([]dnsrecordsv1.BatchDnsRecordsRequestPutsItem, 0, len(items))
		for _, item := range items {
			m := item.(map[string]interface{})
			put, err := sess.NewBatchDnsRecordsRequestPutsItem(
				m["id"].(string),
				m[cisDNSRecordName].(string),
				m[cisDNSRecordType].(string),
				int64(m[cisDNSRecordTTL].(int)),
				m[cisDNSRecordContent].(string),
			)
			if err != nil {
				return err
			}
			if v, ok := m[cisDNSRecordPriority].(int); ok && v != 0 {
				put.Priority = core.Int64Ptr(int64(v))
			}
			if v, ok := m[cisDNSRecordProxied].(bool); ok {
				put.Proxied = core.BoolPtr(v)
			}
			if dataRaw, ok := m[cisDNSRecordData]; ok {
				if dataMap, ok := dataRaw.(map[string]interface{}); ok && len(dataMap) > 0 {
					put.Data = dataMap
				}
			}
			puts = append(puts, *put)
		}
		opt.SetPuts(puts)
	}

	result, response, err := sess.BatchDnsRecords(opt)
	if err != nil {
		log.Printf("Error executing batch DNS records: %s, error %s", response, err)
		return err
	}

	d.SetId(flex.ConvertCisToTfTwoVar(zoneID, crn))

	if result.Result != nil {
		d.Set("result_deletes", flattenBatchDNSRecordDetails(result.Result.Deletes))
		d.Set("result_patches", flattenBatchDNSRecordDetails(result.Result.Patches))
		d.Set("result_posts", flattenBatchDNSRecordDetails(result.Result.Posts))
		d.Set("result_puts", flattenBatchDNSRecordDetails(result.Result.Puts))
	}

	return nil
}

func resourceIBMCISDNSRecordsBatchRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMCISDNSRecordsBatchDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}

func flattenBatchDNSRecordDetails(records []dnsrecordsv1.BatchDnsRecordDetails) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(records))
	for _, r := range records {
		m := map[string]interface{}{}
		if r.ID != nil {
			m["id"] = *r.ID
		}
		if r.Name != nil {
			m[cisDNSRecordName] = *r.Name
		}
		if r.Type != nil {
			m[cisDNSRecordType] = *r.Type
		}
		if r.TTL != nil {
			m[cisDNSRecordTTL] = *r.TTL
		}
		if r.Content != nil {
			m[cisDNSRecordContent] = *r.Content
		}
		if r.Proxied != nil {
			m[cisDNSRecordProxied] = *r.Proxied
		}
		if r.Proxiable != nil {
			m[cisDNSRecordProxiable] = *r.Proxiable
		}
		if r.CreatedOn != nil {
			m[cisDNSRecordCreatedOn] = *r.CreatedOn
		}
		if r.ModifiedOn != nil {
			m[cisDNSRecordModifiedOn] = *r.ModifiedOn
		}
		result = append(result, m)
	}
	return result
}
