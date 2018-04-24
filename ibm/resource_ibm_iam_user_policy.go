package ibm

import (
	"fmt"
	"reflect"
	"strings"

	v1 "github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform/helper/schema"
)

const allIAMEnabledServices = "All Identity and Access enabled services"

var roleNameToID = make(map[string]string)
var roleIDToName = make(map[string]string)

func init() {

	roleNameToID = map[string]string{
		"viewer":        "crn:v1:bluemix:public:iam::::role:Viewer",
		"editor":        "crn:v1:bluemix:public:iam::::role:Editor",
		"operator":      "crn:v1:bluemix:public:iam::::role:Operator",
		"administrator": "crn:v1:bluemix:public:iam::::role:Administrator",
	}
	for k, v := range roleNameToID {
		roleIDToName[v] = k
	}
}

func resourceIBMIAMUserPolicy() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMIAMUserPolicyCreate,
		Read:     resourceIBMIAMUserPolicyRead,
		Update:   resourceIBMIAMUserPolicyUpdate,
		Delete:   resourceIBMIAMUserPolicyDelete,
		Exists:   resourceIBMIAMUserPolicyExists,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			"account_guid": {
				Description: "The bluemix account guid",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"ibm_id": {
				Description: "The ibm id or email of user",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"resources": {
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"service_instance": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"region": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"resource": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"space_guid": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"organization_guid": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"roles": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				MaxItems: 4,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
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

func resourceIBMIAMUserPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	iamClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}
	accountGUID := d.Get("account_guid").(string)
	userEmail := d.Get("ibm_id").(string)
	userID, err := getIBMID(accountGUID, userEmail, meta)
	if err != nil {
		return err
	}
	roleIDSet := d.Get("roles").(*schema.Set)
	roles, err := getRoles(roleIDSet)
	if err != nil {
		return err
	}

	policyServices := d.Get("resources").(*schema.Set)
	resources, err := expandResources(policyServices, iamClient, accountGUID)
	if err != nil {
		return err
	}

	params := v1.AccessPolicyRequest{
		Roles:     roles,
		Resources: resources,
	}

	accessPolicyResponse, etag, err := iamClient.IAMPolicy().Create(accountGUID, userID, params)
	if err != nil {
		return err
	}
	d.SetId(accessPolicyResponse.ID)
	d.Set("account_guid", accountGUID)
	d.Set("ibm_id", userEmail)
	d.Set("etag", etag)

	return resourceIBMIAMUserPolicyRead(d, meta)
}

func resourceIBMIAMUserPolicyRead(d *schema.ResourceData, meta interface{}) error {
	iamClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}
	accountGUID := d.Get("account_guid").(string)
	userEmail := d.Get("ibm_id").(string)
	userID, err := getIBMID(accountGUID, userEmail, meta)
	if err != nil {
		return err
	}
	policyID := d.Id()
	iamPolicy, err := iamClient.IAMPolicy().Get(accountGUID, userID, policyID)
	if err != nil {
		return fmt.Errorf("Unable to read policy:%s", err)
	}
	resources, err := flattenIAMPolicyResource(iamPolicy.Resources, iamClient)
	if err != nil {
		return err
	}
	roles := iamPolicy.Roles
	d.Set("roles", flattenIAMPolicyRoles(roles))
	d.Set("resources", resources)
	return nil
}

func resourceIBMIAMUserPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	iamClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}
	policyID := d.Id()
	accountGUID := d.Get("account_guid").(string)
	etag := d.Get("etag").(string)
	userEmail := d.Get("ibm_id").(string)
	userID, err := getIBMID(accountGUID, userEmail, meta)
	if err != nil {
		return err
	}
	if d.HasChange("roles") || d.HasChange("resources") {
		roles, err := getRoles(d.Get("roles").(*schema.Set))
		if err != nil {
			return err
		}
		policyServices := d.Get("resources").(*schema.Set)
		resources, err := expandResources(policyServices, iamClient, accountGUID)
		if err != nil {
			return err
		}
		accessPolicy := v1.AccessPolicyRequest{
			Roles:     roles,
			Resources: resources,
		}
		_, etag, err = iamClient.IAMPolicy().Update(accountGUID, userID, policyID, etag, accessPolicy)
		if err != nil {
			return fmt.Errorf("Unable to update policy:%s", err)
		}
		d.Set("account_guid", accountGUID)
		d.Set("etag", etag)
	}
	return resourceIBMIAMUserPolicyRead(d, meta)
}

func resourceIBMIAMUserPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	iamClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}
	accountGUID := d.Get("account_guid").(string)
	userEmail := d.Get("ibm_id").(string)
	userID, err := getIBMID(accountGUID, userEmail, meta)
	if err != nil {
		return err
	}
	policyID := d.Id()
	err = iamClient.IAMPolicy().Delete(accountGUID, userID, policyID)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceIBMIAMUserPolicyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	iamClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return false, err
	}
	accClient, err := meta.(ClientSession).BluemixAcccountAPI()
	if err != nil {
		return false, err
	}

	accountGUID := d.Get("account_guid").(string)
	userEmail := d.Get("ibm_id").(string)
	account, err := accClient.Accounts().Get(accountGUID)
	if err != nil {
		return false, fmt.Errorf("Error retrieving account: %s", err)
	}
	userID, err := getIBMID(accountGUID, userEmail, meta)
	if userID == "" || err != nil {
		return false, fmt.Errorf("User %q does not exist in the account:%q", userEmail, account.Name)
	}
	policyID := d.Id()

	accessPolicyResponse, err := iamClient.IAMPolicy().Get(accountGUID, userID, policyID)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}

	return policyID == accessPolicyResponse.ID, nil
}

func expandResources(policyServices *schema.Set, iamClient v1.IAMPAPAPI, accountGUID string) ([]v1.Resources, error) {
	var resources []v1.Resources
	for _, policyService := range policyServices.List() {
		rpm, _ := policyService.(map[string]interface{})
		serviceInstancesList := expandStringList(rpm["service_instance"].([]interface{}))
		sName := strings.TrimSpace(rpm["service_name"].(string))
		if sName == "" {
			return resources, fmt.Errorf("Error service_name cannot be empty")
		}
		serviceName, err := iamClient.IAMService().GetServiceName(sName)
		if err != nil {
			return resources, fmt.Errorf("Error retrieving service %s: %s", sName, err)
		}
		serviceInstance := ""
		if len(serviceInstancesList) > 0 {
			serviceInstance = serviceInstancesList[0]
		}
		if serviceName == allIAMEnabledServices {
			if serviceInstance != "" {
				return nil, fmt.Errorf("For the service %s you must not specify any service_instance. Found following service_instance %s", allIAMEnabledServices, serviceInstancesList)
			}
		}
		resources = append(resources, generateResource(rpm, serviceName, serviceInstance, accountGUID))
	}
	return resources, nil
}

func generateResource(rpm map[string]interface{}, serviceName, serviceInstance, accountGUID string) v1.Resources {
	resourceParam := v1.Resources{
		AccountId:       accountGUID,
		ServiceInstance: serviceInstance,
		Region:          rpm["region"].(string),
		ResourceType:    rpm["resource_type"].(string),
		Resource:        rpm["resource"].(string),
		SpaceId:         rpm["space_guid"].(string),
		OrganizationId:  rpm["organization_guid"].(string),
	}
	if serviceName != allIAMEnabledServices {
		resourceParam.ServiceName = serviceName
	}
	return resourceParam
}

func getIBMID(accountGUID, userEmail string, meta interface{}) (string, error) {
	accClient, err := meta.(ClientSession).BluemixAcccountAPI()
	if err != nil {
		return "", err
	}
	account, err := accClient.Accounts().Get(accountGUID)
	if err != nil {
		return "", fmt.Errorf("Error retrieving account: %s", err)
	}

	accountv1Client, err := meta.(ClientSession).BluemixAcccountv1API()
	if err != nil {
		return "", err
	}
	accUsers, err := accountv1Client.Accounts().GetAccountUsers(accountGUID)
	if err != nil {
		return "", err
	}
	for _, accUser := range accUsers {
		if accUser.Email == userEmail {
			return accUser.IbmUniqueId, nil
		}
	}
	return "", fmt.Errorf("User %q does not exist in the account:%q", userEmail, account.Name)
}

func getRoles(roleIDSet *schema.Set) ([]v1.Roles, error) {
	roleIDS := make([]v1.Roles, 0, roleIDSet.Len())
	for _, elem := range roleIDSet.List() {
		roleID := elem.(string)
		id, ok := roleNameToID[roleID]
		if !ok {
			return roleIDS, fmt.Errorf("The given role %q is not valid. Valid roles are %q", roleID, reflect.ValueOf(roleNameToID).MapKeys())
		}
		role := v1.Roles{
			ID: id,
		}
		roleIDS = append(roleIDS, role)
	}
	return roleIDS, nil
}
