// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
)

func resourceIbmAppConfigCollection() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIbmAppConfigCollectionCreate,
		Read:     resourceIbmAppConfigCollectionRead,
		Update:   resourceIbmAppConfigCollectionUpdate,
		Delete:   resourceIbmAppConfigCollectionDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"guid": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Collection name.",
			},
			"collection_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Collection Id.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Collection description.",
			},
			"tags": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Tags associated with the collection.",
			},
			"created_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of the collection.",
			},
			"updated_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last updated time of the collection data.",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Collection URL.",
			},
		},
	}
}

func resourceIbmAppConfigCollectionCreate(d *schema.ResourceData, meta interface{}) error {
	guid := d.Get("guid").(string)
	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return err
	}

	options := &appconfigurationv1.CreateCollectionOptions{}

	options.SetName(d.Get("name").(string))
	options.SetCollectionID(d.Get("collection_id").(string))
	if _, ok := d.GetOk("description"); ok {
		options.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("tags"); ok {
		options.SetTags(d.Get("tags").(string))
	}

	result, response, err := appconfigClient.CreateCollection(options)
	if err != nil {
		log.Printf("[DEBUG] CreateCollection failed %s\n%s", err, response)
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", guid, *result.CollectionID))

	return resourceIbmAppConfigCollectionRead(d, meta)
}

func resourceIbmAppConfigCollectionRead(d *schema.ResourceData, meta interface{}) error {
	parts, err := idParts(d.Id())
	if err != nil {
		return nil
	}
	appconfigClient, err := getAppConfigClient(meta, parts[0])
	if err != nil {
		return err
	}

	options := &appconfigurationv1.GetCollectionOptions{}

	options.SetExpand(true)
	options.SetCollectionID(parts[1])

	result, response, err := appconfigClient.GetCollection(options)
	if err != nil {
		log.Printf("[DEBUG] GetCollection failed %s\n%s", err, response)
		return err
	}
	d.Set("guid", parts[0])
	if result.Name != nil {
		if err = d.Set("name", result.Name); err != nil {
			return fmt.Errorf("error setting name: %s", err)
		}
	}
	if result.CollectionID != nil {
		if err = d.Set("collection_id", result.CollectionID); err != nil {
			return fmt.Errorf("error setting collection_id: %s", err)
		}
	}
	if result.Description != nil {
		if err = d.Set("description", result.Description); err != nil {
			return fmt.Errorf("error setting description: %s", err)
		}
	}
	if result.Tags != nil {
		if err = d.Set("tags", result.Tags); err != nil {
			return fmt.Errorf("error setting tags: %s", err)
		}
	}
	if result.CreatedTime != nil {
		if err = d.Set("created_time", result.CreatedTime.String()); err != nil {
			return fmt.Errorf("error setting created_time: %s", err)
		}
	}
	if result.UpdatedTime != nil {
		if err = d.Set("updated_time", result.UpdatedTime.String()); err != nil {
			return fmt.Errorf("error setting updated_time: %s", err)
		}
	}
	if result.Href != nil {
		if err = d.Set("href", result.Href); err != nil {
			return fmt.Errorf("error setting href: %s", err)
		}
	}

	return nil
}

func resourceIbmAppConfigCollectionUpdate(d *schema.ResourceData, meta interface{}) error {
	if ok := d.HasChanges("name", "tags", "description"); ok {
		parts, err := idParts(d.Id())
		if err != nil {
			return nil
		}
		appconfigClient, err := getAppConfigClient(meta, parts[0])
		if err != nil {
			return err
		}
		options := &appconfigurationv1.UpdateCollectionOptions{}

		options.SetCollectionID(parts[1])
		options.SetName(d.Get("name").(string))
		if _, ok := d.GetOk("description"); ok {
			options.SetDescription(d.Get("description").(string))
		}
		if _, ok := d.GetOk("tags"); ok {
			options.SetTags(d.Get("tags").(string))
		}

		_, response, err := appconfigClient.UpdateCollection(options)
		if err != nil {
			log.Printf("[DEBUG] UpdateCollection failed %s\n%s", err, response)
			return err
		}

		return resourceIbmAppConfigCollectionRead(d, meta)
	}
	return nil
}

func resourceIbmAppConfigCollectionDelete(d *schema.ResourceData, meta interface{}) error {
	parts, err := idParts(d.Id())
	if err != nil {
		return nil
	}

	appconfigClient, err := getAppConfigClient(meta, parts[0])
	if err != nil {
		return err
	}
	options := &appconfigurationv1.DeleteCollectionOptions{}

	options.SetCollectionID(parts[1])

	response, err := appconfigClient.DeleteCollection(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] DeleteCollection failed %s\n%s", err, response)
		return err
	}

	d.SetId("")

	return nil
}
