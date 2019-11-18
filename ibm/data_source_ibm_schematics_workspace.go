package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
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
	d.Set("template_id", templateID)

	types := WorkspaceInfo.Type
	d.Set("types", types)

	tags := WorkspaceInfo.Tags
	d.Set("tags", tags)

	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/schematics")

	return nil

}
