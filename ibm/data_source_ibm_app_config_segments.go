// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"
	"net/url"
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
)

func dataSourceIbmAppConfigSegments() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIbmAppConfigSegmentsRead,

		Schema: map[string]*schema.Schema{
			"guid": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.",
			},
			"sort": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Sort the segment details based on the specified attribute.",
			},
			"tags": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "filter the resources to be returned based on the associated tags. Returns resources associated with any of the specified tags.",
			},
			"expand": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If set to `true`, returns expanded view of the resource details.",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different set of records, use `limit` with `offset` to page through the available records.",
			},
			"offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. Use `offset` with `limit` to page through the available records.",
			},
			"includes": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Segment details to include the associated rules in the response.",
			},
			"segments": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Array of Segments.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Segment name.",
						},
						"segment_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Segment id.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Segment description.",
						},
						"tags": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Tags associated with the segments.",
						},
						"rules": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of rules that determine if the entity is part of the segment.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"attribute_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Attribute name.",
									},
									"operator": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Operator to be used for the evaluation if the entity is part of the segment.",
									},
									"values": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "List of values. Entities matching any of the given values will be considered to be part of the segment.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"created_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time of the segment.",
						},
						"updated_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Last modified time of the segment data.",
						},
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Segment URL.",
						},
					},
				},
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total number of records.",
			},
			"next": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the next list of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the response.",
						},
					},
				},
			},
			"first": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the first page of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the response.",
						},
					},
				},
			},
			"previous": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the previous list of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the response.",
						},
					},
				},
			},
			"last": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the last page of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the response.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmAppConfigSegmentsRead(d *schema.ResourceData, meta interface{}) error {
	guid := d.Get("guid").(string)

	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return err
	}

	options := &appconfigurationv1.ListSegmentsOptions{}
	if _, ok := d.GetOk("expand"); ok {
		options.SetExpand(d.Get("expand").(bool))
	}
	if _, ok := d.GetOk("tags"); ok {
		options.SetTags(d.Get("tags").(string))
	}
	if _, ok := d.GetOk("includes"); ok {
		options.SetTags(d.Get("includes").(string))
	}
	if _, ok := d.GetOk("sort"); ok {
		options.SetTags(d.Get("sort").(string))
	}
	var segmentsList *appconfigurationv1.SegmentsList
	var offset int64 = 0
	var limit int64 = 10
	finalList := []appconfigurationv1.Segment{}

	var isLimit bool
	if _, ok := d.GetOk("limit"); ok {
		isLimit = true
		limit = int64(d.Get("limit").(int))
	}
	options.SetLimit(limit)

	if _, ok := d.GetOk("offset"); ok {
		offset = int64(d.Get("offset").(int))
	}
	for {
		options.SetOffset(offset)
		result, response, err := appconfigClient.ListSegments(options)
		segmentsList = result
		if err != nil {
			log.Printf("[DEBUG] ListSegments failed %s\n%s", err, response)
			return err
		}
		if isLimit {
			offset = 0
		} else {
			offset = dataSourceSegmentsListGetNext(result.Next)
		}
		finalList = append(finalList, result.Segments...)
		if offset == 0 {
			break
		}
	}

	segmentsList.Segments = finalList

	d.SetId(guid)

	if segmentsList.Segments != nil {
		err = d.Set("segments", dataSourceSegmentsListFlattenSegments(segmentsList.Segments))
		if err != nil {
			return fmt.Errorf("error setting segments %s", err)
		}
	}
	if segmentsList.TotalCount != nil {
		if err = d.Set("total_count", segmentsList.TotalCount); err != nil {
			return fmt.Errorf("error setting total_count: %s", err)
		}
	}
	if segmentsList.Limit != nil {
		if err = d.Set("limit", segmentsList.Limit); err != nil {
			return fmt.Errorf("error setting limit: %s", err)
		}
	}
	if segmentsList.Offset != nil {
		if err = d.Set("offset", segmentsList.Offset); err != nil {
			return fmt.Errorf("error setting offset: %s", err)
		}
	}
	if segmentsList.First != nil {
		err = d.Set("first", dataSourceSegmentListFlattenPagination(*segmentsList.First))
		if err != nil {
			return fmt.Errorf("error setting first %s", err)
		}
	}

	if segmentsList.Previous != nil {
		err = d.Set("previous", dataSourceSegmentListFlattenPagination(*segmentsList.Previous))
		if err != nil {
			return fmt.Errorf("error setting previous %s", err)
		}
	}

	if segmentsList.Last != nil {
		err = d.Set("last", dataSourceSegmentListFlattenPagination(*segmentsList.Last))
		if err != nil {
			return fmt.Errorf("error setting last %s", err)
		}
	}
	if segmentsList.Next != nil {
		err = d.Set("next", dataSourceSegmentListFlattenPagination(*segmentsList.Next))
		if err != nil {
			return fmt.Errorf("error setting next %s", err)
		}
	}
	return nil
}

func dataSourceSegmentsListFlattenSegments(result []appconfigurationv1.Segment) (segments []map[string]interface{}) {
	for _, segmentsItem := range result {
		segments = append(segments, dataSourceSegmentsListSegmentsToMap(segmentsItem))
	}

	return segments
}

func dataSourceSegmentsListSegmentsToMap(segmentsItem appconfigurationv1.Segment) (segmentsMap map[string]interface{}) {
	segmentsMap = map[string]interface{}{}

	if segmentsItem.Name != nil {
		segmentsMap["name"] = segmentsItem.Name
	}
	if segmentsItem.SegmentID != nil {
		segmentsMap["segment_id"] = segmentsItem.SegmentID
	}
	if segmentsItem.Description != nil {
		segmentsMap["description"] = segmentsItem.Description
	}
	if segmentsItem.Tags != nil {
		segmentsMap["tags"] = segmentsItem.Tags
	}
	if segmentsItem.Rules != nil {
		rulesList := []map[string]interface{}{}
		for _, rulesItem := range segmentsItem.Rules {
			rulesList = append(rulesList, dataSourceSegmentsListSegmentsRulesToMap(rulesItem))
		}
		segmentsMap["rules"] = rulesList
	}
	if segmentsItem.CreatedTime != nil {
		segmentsMap["created_time"] = segmentsItem.CreatedTime.String()
	}
	if segmentsItem.UpdatedTime != nil {
		segmentsMap["updated_time"] = segmentsItem.UpdatedTime.String()
	}
	if segmentsItem.Href != nil {
		segmentsMap["href"] = segmentsItem.Href
	}

	return segmentsMap
}

func dataSourceSegmentsListSegmentsRulesToMap(rulesItem appconfigurationv1.Rule) (rulesMap map[string]interface{}) {
	rulesMap = map[string]interface{}{}

	if rulesItem.AttributeName != nil {
		rulesMap["attribute_name"] = rulesItem.AttributeName
	}
	if rulesItem.Operator != nil {
		rulesMap["operator"] = rulesItem.Operator
	}
	if rulesItem.Values != nil {
		rulesMap["values"] = rulesItem.Values
	}

	return rulesMap
}

func dataSourceSegmentsListGetNext(next interface{}) int64 {
	if reflect.ValueOf(next).IsNil() {
		return 0
	}

	u, err := url.Parse(reflect.ValueOf(next).Elem().FieldByName("Href").Elem().String())
	if err != nil {
		return 0
	}

	q := u.Query()
	var page string

	if q.Get("start") != "" {
		page = q.Get("start")
	} else if q.Get("offset") != "" {
		page = q.Get("offset")
	}

	convertedVal, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return 0
	}
	return convertedVal
}

func dataSourceSegmentListFlattenPagination(result appconfigurationv1.PageHrefResponse) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceSegmentListURLToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceSegmentListURLToMap(urlItem appconfigurationv1.PageHrefResponse) (urlMap map[string]interface{}) {
	urlMap = map[string]interface{}{}

	if urlItem.Href != nil {
		urlMap["href"] = urlItem.Href
	}

	return urlMap
}
