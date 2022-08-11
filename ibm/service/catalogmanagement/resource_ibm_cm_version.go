// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement

import (
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
)

func ResourceIBMCmVersion() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMCmVersionCreate,
		Read:     resourceIBMCmVersionRead,
		Delete:   resourceIBMCmVersionDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"catalog_identifier": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Catalog identifier.",
			},
			"offering_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Offering identification.",
			},
			"tags": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Tags array.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"target_kinds": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Target kinds.  Current valid values are 'iks', 'roks', 'vcenter', and 'terraform'.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"content": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "byte array representing the content to be imported.  Only supported for OVA images at this time.",
			},
			"zipurl": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "URL path to zip location.  If not specified, must provide content in the body of this call.",
			},
			"target_version": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The semver value for this new version, if not found in the zip url package content.",
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Version's CRN.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Version of content type.",
			},
			"sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "hash of the content.",
			},
			"created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time this version was created.",
			},
			"updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time this version was last updated.",
			},
			"catalog_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Catalog ID.",
			},
			"kind_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Kind ID.",
			},
			"repo_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Content's repo URL.",
			},
			"source_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Content's source URL (e.g git repo).",
			},
			"tgz_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "File used to on-board this version.",
			},
		},
	}
}

func resourceIBMCmVersionCreate(d *schema.ResourceData, meta interface{}) error {
	catalogManagementClient, err := meta.(conns.ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}

	importOfferingVersionOptions := catalogManagementClient.NewImportOfferingVersionOptions(d.Get("catalog_identifier").(string), d.Get("offering_id").(string))

	if _, ok := d.GetOk("tags"); ok {
		importOfferingVersionOptions.SetTags(d.Get("tags").([]string))
	}
	if _, ok := d.GetOk("target_kinds"); ok {
		list := flex.ExpandStringList(d.Get("target_kinds").([]interface{}))
		importOfferingVersionOptions.SetTargetKinds(list)

	}
	if _, ok := d.GetOk("content"); ok {
		importOfferingVersionOptions.SetContent([]byte(d.Get("content").(string)))
	}
	if _, ok := d.GetOk("zipurl"); ok {
		importOfferingVersionOptions.SetZipurl(d.Get("zipurl").(string))
	}
	if _, ok := d.GetOk("target_version"); ok {
		importOfferingVersionOptions.SetTargetVersion(d.Get("target_version").(string))
	}

	offering, response, err := catalogManagementClient.ImportOfferingVersion(importOfferingVersionOptions)

	if err != nil {
		log.Printf("[DEBUG] ImportOfferingVersion failed %s\n%s", err, response)
		return err
	}

	versionLocator := *offering.Kinds[0].Versions[0].VersionLocator

	d.SetId(versionLocator)

	return resourceIBMCmVersionRead(d, meta)
}

func resourceIBMCmVersionRead(d *schema.ResourceData, meta interface{}) error {
	catalogManagementClient, err := meta.(conns.ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}

	getVersionOptions := &catalogmanagementv1.GetVersionOptions{}

	getVersionOptions.SetVersionLocID(d.Id())

	offering, response, err := catalogManagementClient.GetVersion(getVersionOptions)
	version := offering.Kinds[0].Versions[0]

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetVersion failed %s\n%s", err, response)
		return err
	}

	if err = d.Set("crn", version.CRN); err != nil {
		return fmt.Errorf("[ERROR] Error setting crn: %s", err)
	}
	if err = d.Set("version", version.Version); err != nil {
		return fmt.Errorf("[ERROR] Error setting version: %s", err)
	}
	if err = d.Set("sha", version.Sha); err != nil {
		return fmt.Errorf("[ERROR] Error setting sha: %s", err)
	}
	if err = d.Set("created", version.Created.String()); err != nil {
		return fmt.Errorf("[ERROR] Error setting created: %s", err)
	}
	if err = d.Set("updated", version.Updated.String()); err != nil {
		return fmt.Errorf("[ERROR] Error setting updated: %s", err)
	}
	if err = d.Set("catalog_id", version.CatalogID); err != nil {
		return fmt.Errorf("[ERROR] Error setting catalog_id: %s", err)
	}
	if err = d.Set("kind_id", version.KindID); err != nil {
		return fmt.Errorf("[ERROR] Error setting kind_id: %s", err)
	}
	if err = d.Set("repo_url", version.RepoURL); err != nil {
		return fmt.Errorf("[ERROR] Error setting repo_url: %s", err)
	}
	if err = d.Set("source_url", version.SourceURL); err != nil {
		return fmt.Errorf("[ERROR] Error setting source_url: %s", err)
	}
	if err = d.Set("tgz_url", version.TgzURL); err != nil {
		return fmt.Errorf("[ERROR] Error setting tgz_url: %s", err)
	}

	return nil
}

func resourceIBMCmVersionDelete(d *schema.ResourceData, meta interface{}) error {
	catalogManagementClient, err := meta.(conns.ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}

	deleteVersionOptions := &catalogmanagementv1.DeleteVersionOptions{}
	deleteVersionOptions.SetVersionLocID(d.Id())

	response, err := catalogManagementClient.DeleteVersion(deleteVersionOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteVersion failed %s\n%s", err, response)
		return err
	}

	d.SetId("")

	return nil
}
