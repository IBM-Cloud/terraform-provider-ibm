package ibm

import (
	"fmt"
	"log"

	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/managementv2"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIBMResourceGroup() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMResourceGroupCreate,
		Read:     resourceIBMResourceGroupRead,
		Update:   resourceIBMResourceGroupUpdate,
		Delete:   resourceIBMResourceGroupDelete,
		Exists:   resourceIBMResourceGroupExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the resource group",
			},

			"quota_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the quota",
				Removed:     "This field is removed",
			},

			"default": {
				Description: "Specifies whether its default resource group or not",
				Type:        schema.TypeBool,
				Computed:    true,
			},

			"state": {
				Type:        schema.TypeString,
				Description: "State of the resource group",
				Computed:    true,
			},

			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func resourceIBMResourceGroupCreate(d *schema.ResourceData, meta interface{}) error {
	rMgtClient, err := meta.(ClientSession).ResourceManagementAPIv2()
	if err != nil {
		return err
	}
	name := d.Get("name").(string)

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	accountID := userDetails.userAccount

	resourceGroupCreate := models.ResourceGroupv2{
		ResourceGroup: models.ResourceGroup{
			Name:      name,
			AccountID: accountID,
		},
	}

	resourceGroup, err := rMgtClient.ResourceGroup().Create(resourceGroupCreate)
	if err != nil {
		return fmt.Errorf("Error creating resource group: %s", err)
	}

	d.SetId(resourceGroup.ID)

	return resourceIBMResourceGroupRead(d, meta)
}

func resourceIBMResourceGroupRead(d *schema.ResourceData, meta interface{}) error {
	rMgtClient, err := meta.(ClientSession).ResourceManagementAPIv2()
	if err != nil {
		return err
	}
	resourceGroupID := d.Id()

	resourceGroup, err := rMgtClient.ResourceGroup().Get(resourceGroupID)
	if err != nil {
		return fmt.Errorf("Error retrieving resource group: %s", err)
	}

	d.Set("name", resourceGroup.Name)
	d.Set("state", resourceGroup.State)
	d.Set("default", resourceGroup.Default)

	return nil
}

func resourceIBMResourceGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	rMgtClient, err := meta.(ClientSession).ResourceManagementAPIv2()
	if err != nil {
		return err
	}

	resourceGroupID := d.Id()

	updateReq := managementv2.ResourceGroupUpdateRequest{}
	hasChange := false
	if d.HasChange("name") {
		updateReq.Name = d.Get("name").(string)
		hasChange = true
	}

	if hasChange {
		_, err := rMgtClient.ResourceGroup().Update(resourceGroupID, &updateReq)
		if err != nil {
			return fmt.Errorf("Error updating resource group: %s", err)
		}

	}
	return resourceIBMResourceGroupRead(d, meta)
}

func resourceIBMResourceGroupDelete(d *schema.ResourceData, meta interface{}) error {

	rMgtClient, err := meta.(ClientSession).ResourceManagementAPIv2()
	if err != nil {
		return err
	}

	resourceGroupID := d.Id()

	err = rMgtClient.ResourceGroup().Delete(resourceGroupID)
	if err != nil {
		log.Fatal(err)
	}
	d.SetId("")

	return nil
}

func resourceIBMResourceGroupExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	rMgtClient, err := meta.(ClientSession).ResourceManagementAPIv2()
	if err != nil {
		return false, err
	}
	resourceGroupID := d.Id()

	resourceGroup, err := rMgtClient.ResourceGroup().Get(resourceGroupID)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}

	return resourceGroup.ID == resourceGroupID, nil
}
