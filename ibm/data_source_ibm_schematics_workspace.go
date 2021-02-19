/*
* IBM Confidential
* Object Code Only Source Materials
* 5747-SM3
* (c) Copyright IBM Corp. 2017,2021
*
* The source code for this program is not published or otherwise divested
* of its trade secrets, irrespective of what has been deposited with the
* U.S. Copyright Office.
 */

package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSchematicsWorkspace() *schema.Resource {
	return &schema.Resource{
		Read: resourceIBMSchematicsWorkspaceRead,

		Schema: map[string]*schema.Schema{
			"workspace_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of workspace",
			},
			"template_id": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The id of templates",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of workspace",
			},
			"resource_group": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group of workspace",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of workspace",
			},
			"types": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"is_frozen": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_locked": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"location": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The location of workspace",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The description of workspace",
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "cloud resource name of the workspace",
			},
			"catalog_ref": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Catalog references",
			},
			ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this workspace",
			},
		},
	}

}

func resourceIBMSchematicsWorkspaceRead(d *schema.ResourceData, meta interface{}) error {
	scClient, err := meta.(ClientSession).SchematicsAPI()
	if err != nil {
		return err
	}

	wrkAPI := scClient.Workspaces()
	workspaceID := d.Get("workspace_id").(string)
	// templateID := d.Get("template_id").(string)
	WorkspaceInfo, err := wrkAPI.GetWorkspaceByID(workspaceID)
	if err != nil {
		return fmt.Errorf("Error retreiving Workspace details: %s", err)
	}

	templateID := []string{}
	for _, items := range WorkspaceInfo.TemplateData {
		templateID = append(templateID, items.TemplateID)
	}

	d.SetId(WorkspaceInfo.ID)
	d.Set("name", WorkspaceInfo.Name)
	d.Set("resource_group", WorkspaceInfo.ResourceGroup)
	d.Set("status", WorkspaceInfo.Status)
	d.Set("is_frozen", WorkspaceInfo.WorkspaceStatus.Frozen)
	d.Set("is_locked", WorkspaceInfo.WorkspaceStatus.Locked)
	d.Set("location", WorkspaceInfo.Location)
	d.Set("description", WorkspaceInfo.Description)
	d.Set("crn", WorkspaceInfo.CRN)
	d.Set("template_id", templateID)

	types := WorkspaceInfo.Type
	d.Set("types", types)

	//Update content catalog info of the workspace
	d.Set("catalog_ref", flattenCatalogRef(WorkspaceInfo.CatalogRef))

	tags := WorkspaceInfo.Tags
	d.Set("tags", tags)

	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/schematics")

	return nil

}
