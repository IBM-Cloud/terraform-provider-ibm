// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const rhaiiProjectDataSourceName = "ibm_rhaii_project"

// DataSourceIBMRhaiiProject reads an existing Red Hat AI Inference (RHAII) project
// by name or by ID.
func DataSourceIBMRhaiiProject() *schema.Resource {
	s := map[string]*schema.Schema{
		"name": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ExactlyOneOf: []string{"name", "identifier"},
			Description:  "The name of the project.",
		},
		"identifier": {
			Type:          schema.TypeString,
			Optional:      true,
			ExactlyOneOf:  []string{"name", "identifier"},
			ConflictsWith: []string{"resource_group_id", "location"},
			Description:   "The ID (GUID) or CRN of the project.",
		},
		"resource_group_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The ID of the resource group to search in.",
		},
		"location": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The region of the project.",
		},
		"plan": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The pricing plan of the project.",
		},
		"tags": {
			Type:        schema.TypeSet,
			Computed:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Set:         flex.ResourceIBMVPCHash,
			Description: "The user tags of the project.",
		},
		"access_tags": {
			Type:        schema.TypeSet,
			Computed:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Set:         flex.ResourceIBMVPCHash,
			Description: "The access management tags of the project.",
		},
	}
	for k, v := range rhaiiProjectComputedSchema() {
		s[k] = v
	}

	return &schema.Resource{
		ReadContext: dataSourceIBMRhaiiProjectRead,
		Schema:      s,
	}
}

func dataSourceIBMRhaiiProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rcClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return rhaiiDiagFor(err, rhaiiProjectDataSourceName, "read", "initialize-client")
	}

	var instance *rc.ResourceInstance
	if id, ok := d.GetOk("identifier"); ok {
		instance, _, err = rcClient.GetResourceInstanceWithContext(ctx, &rc.GetResourceInstanceOptions{ID: ptr(id.(string))})
		if err != nil {
			return rhaiiDiagFor(fmt.Errorf("GetResourceInstanceWithContext failed: %s", err), rhaiiProjectDataSourceName, "read", "get-instance")
		}
		if !isRhaiiInstance(instance) {
			return rhaiiDiagFor(fmt.Errorf("resource instance %s is not a Red Hat AI Inference project (service is %q, expected %q)", id, strDeref(instance.ResourceID), rhaiiServiceName), rhaiiProjectDataSourceName, "read", "check-service")
		}
	} else {
		instance, err = rhaiiFindProjectByName(ctx, rcClient, d.Get("name").(string), d.Get("resource_group_id").(string), d.Get("location").(string))
		if err != nil {
			return rhaiiDiagFor(err, rhaiiProjectDataSourceName, "read", "find-instance")
		}
	}

	d.SetId(*instance.ID)
	if err := setRhaiiProjectAttributes(ctx, d, meta, instance); err != nil {
		return rhaiiDiagFor(err, rhaiiProjectDataSourceName, "read", "set-attributes")
	}

	tags, err := flex.GetTagsUsingCRN(meta, *instance.CRN)
	if err != nil {
		log.Printf("[ERROR] Error on get of %s (%s) tags: %s", rhaiiProjectDataSourceName, d.Id(), err)
	}
	d.Set("tags", tags)

	accessTags, err := flex.GetGlobalTagsUsingCRN(meta, *instance.CRN, "", rhaiiAccessTagType)
	if err != nil {
		log.Printf("[ERROR] Error on get of %s (%s) access tags: %s", rhaiiProjectDataSourceName, d.Id(), err)
	}
	d.Set("access_tags", accessTags)

	return nil
}

// rhaiiFindProjectByName lists the RHAII instances with a name and returns the
// only match. Removed and pending reclamation instances are ignored.
func rhaiiFindProjectByName(ctx context.Context, rcClient *rc.ResourceControllerV2, name, resourceGroupID, location string) (*rc.ResourceInstance, error) {
	listOptions := &rc.ListResourceInstancesOptions{
		Name:       &name,
		ResourceID: ptr(rhaiiServiceName),
	}
	if resourceGroupID != "" {
		listOptions.ResourceGroupID = &resourceGroupID
	}

	var matches []rc.ResourceInstance
	for {
		list, _, err := rcClient.ListResourceInstancesWithContext(ctx, listOptions)
		if err != nil {
			return nil, fmt.Errorf("ListResourceInstancesWithContext failed: %s", err)
		}
		for i := range list.Resources {
			instance := list.Resources[i]
			if isRhaiiInstanceGone(&instance) {
				continue
			}
			if location != "" && strDeref(instance.RegionID) != location {
				continue
			}
			matches = append(matches, instance)
		}
		start, err := rhaiiNextStart(list.NextURL)
		if err != nil {
			return nil, fmt.Errorf("error parsing next_url of the resource instance list: %s", err)
		}
		if start == "" {
			break
		}
		listOptions.Start = &start
	}

	switch len(matches) {
	case 0:
		return nil, fmt.Errorf("no Red Hat AI Inference project found with name %q. Check the name, resource_group_id and location", name)
	case 1:
		return &matches[0], nil
	default:
		return nil, fmt.Errorf("found %d Red Hat AI Inference projects with name %q. Set resource_group_id or location to narrow the search, or use identifier", len(matches), name)
	}
}
