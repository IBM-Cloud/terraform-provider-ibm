package ibm

import (
	"fmt"
	"log"
	"strings"
	"time"

	//"fmt"

	"github.com/IBM/dns-svcs-go-sdk/dnssvcsv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
)

const (
	pdnsZoneID = "zone_id"
	pdnsVpcCRN = "vpc_crn"
)

func resourceIBMPrivateDNSPermittedNetwork() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMPrivateDnsPermittedNetworkCreate,
		Read:     resourceIBMPrivateDnsPermittedNetworkRead,
		Update:   resourceIBMPrivateDnsPermittedNetworkUpdate,
		Delete:   resourceIBMPrivateDnsPermittedNetworkDelete,
		Exists:   resourceIBMPrivateDnsPermittedNetworkExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

		Schema: map[string]*schema.Schema{
			pdnsInstanceID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			pdnsZoneID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			pdnsVpcCRN: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
		},
	}
}
func resourceIBMPrivateDnsPermittedNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return err
	}

	instanceID := d.Get(pdnsInstanceID).(string)
	zoneID := d.Get(pdnsZoneID).(string)
	vpcCRN := d.Get(pdnsVpcCRN).(string)

	createPermittedNetworkOptions := sess.NewCreatePermittedNetworkOptions(instanceID, zoneID)
	permittedNetworkCrn, err := sess.NewPermittedNetworkVpc(vpcCRN)

	createPermittedNetworkOptions.SetPermittedNetwork(permittedNetworkCrn)
	createPermittedNetworkOptions.SetType(dnssvcsv1.CreatePermittedNetworkOptions_Type_Vpc)
	response, _, err := sess.CreatePermittedNetwork(createPermittedNetworkOptions)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", instanceID, zoneID, *response.ID))
	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}

	log.Printf("[DEBUG] TEST5")

	return resourceIBMPrivateDnsPermittedNetworkRead(d, meta)
}

func resourceIBMPrivateDnsPermittedNetworkRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")
	getPermittedNetworkOptions := sess.NewGetPermittedNetworkOptions(id_set[0], id_set[1], id_set[2])
	_, _, err = sess.GetPermittedNetwork(getPermittedNetworkOptions)

	if err != nil {
		return err
	}

	return nil
}

func resourceIBMPrivateDnsPermittedNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceIBMPrivateDnsPermittedNetworkRead(d, meta)
}

func resourceIBMPrivateDnsPermittedNetworkDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")
	deletePermittedNetworkOptions := sess.NewDeletePermittedNetworkOptions(id_set[0], id_set[1], id_set[2])
	_, _, reqErr := sess.DeletePermittedNetwork(deletePermittedNetworkOptions)

	if reqErr != nil {
		return reqErr
	}

	d.SetId("")
	return nil
}

func resourceIBMPrivateDnsPermittedNetworkExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return false, err
	}

	id_set := strings.Split(d.Id(), "/")
	getPermittedNetworkOptions := sess.NewGetPermittedNetworkOptions(id_set[0], id_set[1], id_set[2])
	_, _, err = sess.GetPermittedNetwork(getPermittedNetworkOptions)
	if err != nil {
		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}
