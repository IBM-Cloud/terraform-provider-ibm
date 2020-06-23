package ibm

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/networking-go-sdk/transitgatewayapisv1"
)

const (
	tgGateways      = "transit_gateways"
	tgResourceGroup = "resource_group"
	tgID            = "id"
	tgCrn           = "crn"
	tgName          = "name"
	tgLocation      = "location"
	tgCreatedAt     = "created_at"
	tgGlobal        = "global"
	tgStatus        = "status"
	tgUpdatedAt     = "updated_at"

	isTransitGatewayProvisioning     = "provisioning"
	isTransitGatewayProvisioningDone = "done"
	isTransitGatewayDeleting         = "deleting"
	isTransitGatewayDeleted          = "done"
)

func resourceIBMTransitGateway() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMTransitGatewayCreate,
		Read:     resourceIBMTransitGatewayRead,
		Delete:   resourceIBMTransitGatewayDelete,
		Exists:   resourceIBMTransitGatewayExists,
		Update:   resourceIBMTransitGatewayUpdate,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

		Schema: map[string]*schema.Schema{
			tgLocation: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			tgName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			tgGlobal: {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
			},

			tgResourceGroup: {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},

			ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the resource",
			},

			ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func transitgatewayClient(meta interface{}) (*transitgatewayapisv1.TransitGatewayApIsV1, error) {
	sess, err := meta.(ClientSession).TransitGatewayV1API()
	return sess, err
}

func resourceIBMTransitGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	client, err := transitgatewayClient(meta)
	if err != nil {
		return err
	}

	location := d.Get(tgLocation).(string)
	name := d.Get(tgName).(string)
	global := d.Get(tgGlobal).(bool)

	createTransitGatewayOptions := &transitgatewayapisv1.CreateTransitGatewayOptions{}

	createTransitGatewayOptions.Name = &name
	createTransitGatewayOptions.Location = &location
	createTransitGatewayOptions.Global = &global

	if _, ok := d.GetOk(tgResourceGroup); ok {
		resourceGroup := d.Get(tgResourceGroup).(string)
		createTransitGatewayOptions.ResourceGroup = &transitgatewayapisv1.ResourceGroupIdentity{ID: &resourceGroup}
	}

	//log.Println("going to create tgw now with options", *createTransitGatewayOptions.ResourceGroup)
	tgw, response, err := client.CreateTransitGateway(createTransitGatewayOptions)
	log.Println("Creation Success")

	if err != nil {
		log.Printf("[DEBUG] Create Transit Gateway err %s\n%s", err, response)
		return err
	}
	d.SetId(*tgw.ID)

	_, err = isWaitForTransitGatewayAvailable(client, d.Id(), d.Timeout(schema.TimeoutCreate))

	return resourceIBMTransitGatewayUpdate(d, meta)
}

func isWaitForTransitGatewayAvailable(client *transitgatewayapisv1.TransitGatewayApIsV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for transit gateway (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isTransitGatewayProvisioning},
		Target:     []string{isTransitGatewayProvisioningDone, ""},
		Refresh:    isTransitGatewayRefreshFunc(client, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isTransitGatewayRefreshFunc(client *transitgatewayapisv1.TransitGatewayApIsV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		gettgwoptions := &transitgatewayapisv1.DetailTransitGatewayOptions{
			ID: &id,
		}
		transitGateway, response, err := client.DetailTransitGateway(gettgwoptions)
		if err != nil {
			return nil, "", fmt.Errorf("Error Getting Transit Gateway: %s\n%s", err, response)
		}

		if *transitGateway.Status == "available" || *transitGateway.Status == "failed" {
			return transitGateway, isTransitGatewayProvisioningDone, nil
		}

		return transitGateway, isTransitGatewayProvisioning, nil
	}
}

func resourceIBMTransitGatewayRead(d *schema.ResourceData, meta interface{}) error {
	ID := d.Id()
	return transitGatewayGet(d, meta, ID)
}

func transitGatewayGet(d *schema.ResourceData, meta interface{}, id string) error {
	client, err := transitgatewayClient(meta)
	if err != nil {
		return err
	}
	tgOptions := &transitgatewayapisv1.DetailTransitGatewayOptions{}
	if id != "" {
		tgOptions.ID = &id
	}

	log.Println("Inside get detail tgw now")
	tgw, _, err := client.DetailTransitGateway(tgOptions)
	if err != nil {
		return err
	}

	log.Println("Got detail tgw now")
	d.SetId(*tgw.ID)
	d.Set("crn", tgw.Crn)
	d.Set("name", tgw.Name)
	d.Set("location", tgw.Location)
	d.Set("created_at", tgw.CreatedAt.String())

	log.Println("Updated Field is:", tgw.UpdatedAt)
	if tgw.UpdatedAt != nil {
		d.Set("updated_at", tgw.UpdatedAt.String())
	}
	d.Set("global", tgw.Global)
	d.Set("status", tgw.Status)

	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}

	d.Set(ResourceControllerURL, controller+"/interconnectivity/transit")
	d.Set(ResourceName, *tgw.Name)
	d.Set(ResourceCRN, *tgw.Crn)
	d.Set(ResourceStatus, *tgw.Status)
	if tgw.ResourceGroup != nil {
		rg := tgw.ResourceGroup
		d.Set("resource_group", *rg.ID)
		d.Set(ResourceGroupName, *rg.ID)
	}
	return nil
}

func resourceIBMTransitGatewayUpdate(d *schema.ResourceData, meta interface{}) error {
	client, err := transitgatewayClient(meta)

	if err != nil {
		return err
	}

	ID := d.Id()
	tgOptions := &transitgatewayapisv1.DetailTransitGatewayOptions{
		ID: &ID,
	}
	_, tgw, err := client.DetailTransitGateway(tgOptions)

	if err != nil {
		log.Printf("Error fetching Tranisit  Gateway: %s", tgw)
		return err
	}

	updateTransitGatewayOptions := &transitgatewayapisv1.UpdateTransitGatewayOptions{}
	updateTransitGatewayOptions.ID = &ID
	if d.HasChange(tgName) {
		if _, ok := d.GetOk(tgName); ok {
			name := d.Get(tgName).(string)
			updateTransitGatewayOptions.Name = &name
		}
	}
	if d.HasChange(tgGlobal) {
		if _, ok := d.GetOk(tgGlobal); ok {
			global := d.Get(tgGlobal).(bool)
			updateTransitGatewayOptions.Global = &global
		}
	}
	_, response, err := client.UpdateTransitGateway(updateTransitGatewayOptions)
	if err != nil {
		log.Printf("[DEBUG] Update Transit Gateway err %s\n%s", err, response)
		return err
	}

	return resourceIBMTransitGatewayRead(d, meta)
}

func resourceIBMTransitGatewayDelete(d *schema.ResourceData, meta interface{}) error {

	client, err := transitgatewayClient(meta)
	if err != nil {
		return err
	}

	ID := d.Id()
	delOptions := &transitgatewayapisv1.DeleteTransitGatewayOptions{
		ID: &ID,
	}
	response, err := client.DeleteTransitGateway(delOptions)

	if err != nil && response.StatusCode != 404 {
		log.Printf("Error deleting Transit Gateway: %s", response)
		return err
	}
	_, err = isWaitForTransitGatewayDeleted(client, ID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func isWaitForTransitGatewayDeleted(client *transitgatewayapisv1.TransitGatewayApIsV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for transit gateway (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isTransitGatewayDeleting},
		Target:     []string{"", isTransitGatewayDeleted},
		Refresh:    isTransitGatewayDeleteRefreshFunc(client, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isTransitGatewayDeleteRefreshFunc(client *transitgatewayapisv1.TransitGatewayApIsV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] delete function here")
		gettgwoptions := &transitgatewayapisv1.DetailTransitGatewayOptions{
			ID: &id,
		}
		transitGateway, response, err := client.DetailTransitGateway(gettgwoptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return client, isTransitGatewayDeleted, nil
			}
			return transitGateway, "", fmt.Errorf("Error Getting Transit Gateway: %s\n%s", err, response)
		}
		return transitGateway, isTransitGatewayDeleting, err
	}
}

func resourceIBMTransitGatewayExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	client, err := transitgatewayClient(meta)
	if err != nil {
		return false, err
	}

	ID := d.Id()

	tgOptions := &transitgatewayapisv1.DetailTransitGatewayOptions{}
	if ID != "" {
		tgOptions.ID = &ID
	}
	_, response, err := client.DetailTransitGateway(tgOptions)
	if err != nil {
		return false, fmt.Errorf("Error Getting Transit Gateway: %s\n%s", err, response)
	}

	if response.StatusCode == 404 {
		d.SetId("")
		return false, nil
	}
	return true, nil
}
