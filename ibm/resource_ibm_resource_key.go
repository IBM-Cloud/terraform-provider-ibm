// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv2"
	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM-Cloud/bluemix-go/utils"
)

func resourceIBMResourceKey() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMResourceKeyCreate,
		Read:     resourceIBMResourceKeyRead,
		Update:   resourceIBMResourceKeyUpdate,
		Delete:   resourceIBMResourceKeyDelete,
		Exists:   resourceIBMResourceKeyExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the resource key",
			},

			"role": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the user role.Valid roles are Writer, Reader, Manager, Administrator, Operator, Viewer, Editor and Custom Roles.",
				// ValidateFunc: validateRole,
			},

			"resource_instance_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				Description:   "The id of the resource instance for which to create resource key",
				ConflictsWith: []string{"resource_alias_id"},
			},

			"resource_alias_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				Description:   "The id of the resource alias for which to create resource key",
				ConflictsWith: []string{"resource_instance_id"},
			},

			"parameters": {
				Type:             schema.TypeMap,
				Optional:         true,
				DiffSuppressFunc: applyOnce,
				Description:      "Arbitrary parameters to pass. Must be a JSON object",
			},

			"credentials": {
				Description: "Credentials asociated with the key",
				Type:        schema.TypeMap,
				Sensitive:   true,
				Computed:    true,
			},

			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of resource key",
			},

			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "crn of resource key",
			},

			"guid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "When you create a new key, a globally unique identifier (GUID) is assigned.",
			},

			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "When you created a new key, a relative URL path is created identifying the location of the key.",
			},

			"account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An alpha-numeric value identifying the account ID.",
			},

			"resource_group_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The short ID of the resource group.",
			},

			"source_crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN of resource instance or alias associated to the key.",
			},

			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of the key.",
			},

			"iam_compatible": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Specifies whether the key’s credentials support IAM.",
			},

			"resource_instance_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The relative path to the resource.",
			},

			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date when the key was created.",
			},

			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date when the key was last updated.",
			},

			"deleted_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date when the key was deleted.",
			},

			"created_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The subject who created the key.",
			},

			"updated_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The subject who updated the key.",
			},

			"deleted_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The subject who deleted the key.",
			},
		},
	}
}

func resourceIBMResourceKeyCreate(d *schema.ResourceData, meta interface{}) error {
	rsContClient, err := meta.(ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}
	name := d.Get("name").(string)
	role := d.Get("role").(string)

	var instanceID, aliasID string
	if insID, ok := d.GetOk("resource_instance_id"); ok {
		instanceID = insID.(string)
	}

	if aliID, ok := d.GetOk("resource_alias_id"); ok {
		aliasID = aliID.(string)
	}

	if instanceID == "" && aliasID == "" {
		return fmt.Errorf("Provide either `resource_instance_id` or `resource_alias_id`")
	}

	keyParameters := rc.ResourceKeyPostParameters{}

	if parameters, ok := d.GetOk("parameters"); ok {
		temp := parameters.(map[string]interface{})
		for k, v := range temp {
			if v == "true" || v == "false" {
				b, _ := strconv.ParseBool(v.(string))
				keyParameters.SetProperty(k, b)
			} else {
				keyParameters.SetProperty(k, v)
			}
		}
	}

	resourceInstance, sourceCRN, err := getResourceInstanceAndCRN(d, meta)
	if err != nil {
		return fmt.Errorf("Error creating resource key when get instance and CRN: %s", err)
	}

	serviceID := resourceInstance.ResourceID

	rsCatClient, err := meta.(ClientSession).ResourceCatalogAPI()
	if err != nil {
		return fmt.Errorf("Error creating resource key when get ResourceCatalogAPI: %s", err)
	}

	service, err := rsCatClient.ResourceCatalog().Get(*serviceID, true)
	if err != nil {
		return fmt.Errorf("Error creating resource key when get service: %s", err)
	}
	serviceRole, err := getRoleFromName(role, service.Name, meta)
	if err != nil {
		return fmt.Errorf("Error creating resource key when get role: %s", err)
	}

	keyParameters.SetProperty("role_crn", serviceRole.Crn)

	resourceKeyCreate := rc.CreateResourceKeyOptions{
		Name:       &name,
		Source:     sourceCRN,
		Role:       &serviceRole.Crn,
		Parameters: &keyParameters,
	}
	resourceKey, resp, err := rsContClient.CreateResourceKey(&resourceKeyCreate)
	if err != nil {
		return fmt.Errorf("Error creating resource key: %s with resp code: %s", err, resp)
	}

	d.SetId(*resourceKey.ID)

	return resourceIBMResourceKeyRead(d, meta)
}

func resourceIBMResourceKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMResourceKeyRead(d *schema.ResourceData, meta interface{}) error {
	rsContClient, err := meta.(ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}
	resourceKeyID := d.Id()
	resourceKeyGet := rc.GetResourceKeyOptions{
		ID: &resourceKeyID,
	}

	resourceKey, resp, err := rsContClient.GetResourceKey(&resourceKeyGet)
	if err != nil || resourceKey == nil {
		return fmt.Errorf("Error retrieving resource key: %s with resp : %s", err, resp)
	}
	var credInterface map[string]interface{}
	cred, _ := json.Marshal(resourceKey.Credentials)
	json.Unmarshal(cred, &credInterface)
	d.Set("credentials", Flatten(credInterface))
	d.Set("name", *resourceKey.Name)
	d.Set("status", *resourceKey.State)
	if resourceKey.Credentials != nil && resourceKey.Credentials.IamRoleCRN != nil {
		roleCrn := *resourceKey.Credentials.IamRoleCRN
		roleName := roleCrn[strings.LastIndex(roleCrn, ":")+1:]

		if strings.Contains(roleCrn, ":customRole:") {
			iampapv2Client, err := meta.(ClientSession).IAMPAPAPIV2()
			if err == nil {
				var resourceCRN string
				if resourceKey.CRN != nil {
					serviceName := strings.Split(*resourceKey.CRN, ":")
					if len(serviceName) > 4 {
						resourceCRN = serviceName[4]
					}
				}
				roles, err := iampapv2Client.IAMRoles().ListCustomRoles(*resourceKey.AccountID, resourceCRN)
				if err == nil && len(roles) > 0 {
					for _, role := range roles {
						if role.Name == roleName {
							customRoleName := role.DisplayName
							d.Set("role", customRoleName)
						}
					}
				}
			}
		} else {
			d.Set("role", roleName)
		}
	}

	sCrn := *resourceKey.SourceCRN
	if sCrn != "" {
		d.Set("resource_instance_id", sCrn)
	}

	d.Set("crn", *resourceKey.CRN)

	d.Set("guid", *resourceKey.GUID)
	d.Set("url", *resourceKey.URL)
	d.Set("account_id", *resourceKey.AccountID)
	d.Set("resource_group_id", *resourceKey.ResourceGroupID)
	d.Set("source_crn", *resourceKey.SourceCRN)
	d.Set("state", *resourceKey.State)
	d.Set("iam_compatible", *resourceKey.IamCompatible)
	d.Set("resource_instance_url", *resourceKey.ResourceInstanceURL)
	if resourceKey.CreatedAt != nil {
		d.Set("created_at", resourceKey.CreatedAt.String())
	} else {
		d.Set("created_at", "")
	}
	if resourceKey.UpdatedAt != nil {
		d.Set("updated_at", resourceKey.UpdatedAt.String())
	} else {
		d.Set("updated_at", "")
	}
	if resourceKey.DeletedAt != nil {
		d.Set("deleted_at", resourceKey.DeletedAt.String())
	} else {
		d.Set("deleted_at", "")
	}
	d.Set("created_by", *resourceKey.CreatedBy)
	d.Set("updated_by", *resourceKey.UpdatedBy)
	d.Set("deleted_by", *resourceKey.DeletedBy)

	return nil
}

func resourceIBMResourceKeyDelete(d *schema.ResourceData, meta interface{}) error {
	rsContClient, err := meta.(ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}

	resourceKeyID := d.Id()
	resourceKeyDelete := rc.DeleteResourceKeyOptions{
		ID: &resourceKeyID,
	}

	resp, err := rsContClient.DeleteResourceKey(&resourceKeyDelete)
	if err != nil {
		return fmt.Errorf("Error deleting resource key: %s with resp code: %s", err, resp)
	}

	d.SetId("")

	return nil
}

func resourceIBMResourceKeyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	rsContClient, err := meta.(ClientSession).ResourceControllerV2API()
	if err != nil {
		return false, err
	}
	resourceKeyID := d.Id()
	resourceKeyGet := rc.GetResourceKeyOptions{
		ID: &resourceKeyID,
	}

	resourceKey, resp, err := rsContClient.GetResourceKey(&resourceKeyGet)
	if err != nil {
		if resp != nil && (resp.StatusCode == 404 || resp.StatusCode == 410) {
			return false, nil
		}
		return false, fmt.Errorf("Error communicating with the API: %s with resp code: %s", err, resp)
	}
	if err == nil && *resourceKey.State == "removed" {
		return false, nil
	}

	return *resourceKey.ID == resourceKeyID, nil
}

func getResourceInstanceAndCRN(d *schema.ResourceData, meta interface{}) (*rc.ResourceInstance, *string, error) {
	rsContClient, err := meta.(ClientSession).ResourceControllerV2API()
	if err != nil {
		return nil, nil, err
	}
	if insID, ok := d.GetOk("resource_instance_id"); ok {
		insIdString := insID.(string)
		resourceInstanceGet := rc.GetResourceInstanceOptions{
			ID: &insIdString,
		}
		instance, resp, err := rsContClient.GetResourceInstance(&resourceInstanceGet)
		if err != nil {
			log.Printf("Error when get resource instance in getResourceInstanceAndCRN: %s with resp code: %s", err, resp)
			return nil, nil, err
		}
		return instance, instance.CRN, nil
	}

	aliasID := d.Get("resource_alias_id").(string)
	resourceAliasGet := rc.GetResourceAliasOptions{
		ID: &aliasID,
	}
	alias, resp, err := rsContClient.GetResourceAlias(&resourceAliasGet)
	if err != nil {
		log.Printf("Error when get resource alias in getResourceInstanceAndCRN: %s with resp code: %s", err, resp)
		return nil, nil, err
	}
	resourceInstanceGet := rc.GetResourceInstanceOptions{
		ID: alias.ResourceInstanceID,
	}
	instance, resp, err := rsContClient.GetResourceInstance(&resourceInstanceGet)
	if err != nil {
		log.Printf("Error when get resource instance in getResourceInstanceAndCRN: %s with resp code: %s", err, resp)
		return nil, nil, err
	}
	return instance, instance.CRN, nil

}

func getRoleFromName(roleName, serviceName string, meta interface{}) (iampapv2.Role, error) {

	iamClient, err := meta.(ClientSession).IAMPAPAPIV2()
	if err != nil {
		return iampapv2.Role{}, err
	}

	iamRepo := iamClient.IAMRoles()

	var roles []iampapv2.Role

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return iampapv2.Role{}, err
	}
	query := iampapv2.RoleQuery{
		AccountID:   userDetails.userAccount,
		ServiceName: serviceName,
	}
	if serviceName == "" {
		roles, err = iamRepo.ListSystemDefinedRoles()
	} else {
		roles, err = iamRepo.ListAll(query)
	}
	if err != nil {
		return iampapv2.Role{}, err
	}

	role, err := utils.FindRoleByNameV2(roles, roleName)
	if err != nil {
		return iampapv2.Role{}, err
	}
	return role, nil

}

func findRoleByName(supported []models.PolicyRole, name string) (models.PolicyRole, error) {
	for _, role := range supported {
		if role.DisplayName == name {
			return role, nil
		}
	}

	return models.PolicyRole{}, fmt.Errorf("Role [%s] was not found. Valid roles are %s", name, getSupportedRolesString(supported))
}

func getSupportedRolesString(supported []models.PolicyRole) string {
	rolesStr := ""
	for index, role := range supported {
		if index != 0 {
			rolesStr += ", "
		}
		rolesStr += role.DisplayName
	}
	return rolesStr
}
