package ibm

import (
	"fmt"
	"strings"
	"time"

	"github.com/IBM-Cloud/bluemix-go/models"

	"github.com/hashicorp/terraform/flatmap"

	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/controller"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/IBM-Cloud/bluemix-go/crn"
	"github.com/hashicorp/terraform/helper/schema"
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
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Name of the user role.Valid roles are Writer, Reader, Manager, Administrator, Operator, Viewer, Editor.",
				ValidateFunc: validateRole,
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
				Type:        schema.TypeMap,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Arbitrary parameters to pass. Must be a JSON object",
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
		},
	}
}

func resourceIBMResourceKeyCreate(d *schema.ResourceData, meta interface{}) error {
	rsContClient, err := meta.(ClientSession).ResourceControllerAPI()
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

	var keyParams map[string]interface{}

	if parameters, ok := d.GetOk("parameters"); ok {
		keyParams = parameters.(map[string]interface{})
	} else {
		keyParams = make(map[string]interface{})
	}

	resourceInstance, sourceCRN, err := getResourceInstanceAndCRN(d, meta)
	if err != nil {
		return fmt.Errorf("Error creating resource key: %s", err)
	}

	serviceID := resourceInstance.ServiceID

	rsCatClient, err := meta.(ClientSession).ResourceCatalogAPI()
	if err != nil {
		return fmt.Errorf("Error creating resource key: %s", err)
	}

	service, err := rsCatClient.ResourceCatalog().Get(serviceID, true)
	if err != nil {
		return fmt.Errorf("Error creating resource key: %s", err)
	}

	serviceRole, err := getRoleFromName(role, service.Name, meta)
	if err != nil {
		return fmt.Errorf("Error creating resource key: %s", err)
	}
	keyParams["role_crn"] = serviceRole.ID

	request := controller.CreateServiceKeyRequest{
		Name:       name,
		SourceCRN:  *sourceCRN,
		Parameters: keyParams,
	}

	resourceKey, err := rsContClient.ResourceServiceKey().CreateKey(request)
	if err != nil {
		return fmt.Errorf("Error creating resource key: %s", err)
	}

	d.SetId(resourceKey.ID)

	return resourceIBMResourceKeyRead(d, meta)
}

func resourceIBMResourceKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMResourceKeyRead(d *schema.ResourceData, meta interface{}) error {
	rsContClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return err
	}
	resourceKeyID := d.Id()

	resourceKey, err := rsContClient.ResourceServiceKey().GetKey(resourceKeyID)
	if err != nil {
		return fmt.Errorf("Error retrieving resource key: %s", err)
	}
	d.Set("credentials", flatmap.Flatten(resourceKey.Credentials))
	d.Set("name", resourceKey.Name)
	d.Set("status", resourceKey.State)

	if roleCrn, ok := resourceKey.Parameters["role_crn"].(string); ok {
		d.Set("role", roleCrn[strings.LastIndex(roleCrn, ":")+1:])
	} else if roleCrn, ok := resourceKey.Credentials["iam_role_crn"].(string); ok {
		d.Set("role", roleCrn[strings.LastIndex(roleCrn, ":")+1:])
	}

	d.Set("parameters", flatmap.Flatten(filterResourceKeyParameters(resourceKey.Parameters)))

	return nil
}

func resourceIBMResourceKeyDelete(d *schema.ResourceData, meta interface{}) error {
	rsContClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return err
	}

	resourceKeyID := d.Id()

	err = rsContClient.ResourceServiceKey().DeleteKey(resourceKeyID)
	if err != nil {
		return fmt.Errorf("Error deleting resource key: %s", err)
	}

	d.SetId("")

	return nil
}

func resourceIBMResourceKeyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	rsContClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return false, err
	}
	resourceKeyID := d.Id()

	resourceKey, err := rsContClient.ResourceServiceKey().GetKey(resourceKeyID)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}

	return resourceKey.ID == resourceKeyID, nil
}

func getResourceInstanceAndCRN(d *schema.ResourceData, meta interface{}) (*models.ServiceInstance, *crn.CRN, error) {
	rsContClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return nil, nil, err
	}
	if insID, ok := d.GetOk("resource_instance_id"); ok {
		instance, err := rsContClient.ResourceServiceInstance().GetInstance(insID.(string))
		if err != nil {
			return nil, nil, err
		}
		return &instance, &instance.Crn, nil

	}

	alias, err := rsContClient.ResourceServiceAlias().Alias(d.Get("resource_alias_id").(string))
	if err != nil {
		return nil, nil, err
	}
	instance, err := rsContClient.ResourceServiceInstance().GetInstance(alias.ServiceInstanceID)
	if err != nil {
		return nil, nil, err
	}
	return &instance, &instance.Crn, nil

}

func getRoleFromName(roleName, serviceName string, meta interface{}) (models.PolicyRole, error) {

	iamClient, err := meta.(ClientSession).IAMAPI()
	if err != nil {
		return models.PolicyRole{}, err
	}

	iamRepo := iamClient.ServiceRoles()

	var roles []models.PolicyRole

	if serviceName == "" {
		roles, err = iamRepo.ListSystemDefinedRoles()
	} else {
		roles, err = iamRepo.ListServiceRoles(serviceName)
	}
	if err != nil {
		return models.PolicyRole{}, err
	}
	role, err := findRoleByName(roles, roleName)
	if err != nil {
		return models.PolicyRole{}, err
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
