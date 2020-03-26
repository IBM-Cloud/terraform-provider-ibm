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

const ()

func resourceIBMPrivateDNSRecords() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMPrivateDnsRecordCreate,
		Read:     resourceIBMPrivateDnsRecordRead,
		Update:   resourceIBMPrivateDnsRecordUpdate,
		Delete:   resourceIBMPrivateDnsRecordDelete,
		Exists:   resourceIBMPrivateDnsRecordExists,
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
func resourceIBMPrivateDnsRecordCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return err
	}

	instanceID := d.Get(pdnsInstanceID).(string)
	zoneID := d.Get(pdnsZoneID).(string)

	createResourceRecordOptions := sess.NewCreateResourceRecordOptions(instanceID, zoneID)
	createResourceRecordOptions.SetName("exmaple")
	createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_A)
	resourceRecordAData, _ := sess.NewResourceRecordInputRdataRdataARecord("1.1.1.1")
	createResourceRecordOptions.SetRdata(resourceRecordAData)
	response, _, err := sess.CreateResourceRecord(createResourceRecordOptions)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", instanceID, zoneID, *response.ID))

	log.Printf("[DEBUG] TEST5")

	return resourceIBMPrivateDnsRecordRead(d, meta)
}

func resourceIBMPrivateDnsRecordRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")
	getResourceRecordOptions := sess.NewGetResourceRecordOptions(id_set[0], id_set[1], id_set[2])
	_, _, err = sess.GetResourceRecord(getResourceRecordOptions)
	if err != nil {
		return err
	}

	return nil
}

func resourceIBMPrivateDnsRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceIBMPrivateDnsRecordRead(d, meta)
}

func resourceIBMPrivateDnsRecordDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")
	deleteResourceRecordOptions := sess.NewDeleteResourceRecordOptions(id_set[0], id_set[1], id_set[2])
	_, err = sess.DeleteResourceRecord(deleteResourceRecordOptions)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourceIBMPrivateDnsRecordExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return false, err
	}

	id_set := strings.Split(d.Id(), "/")
	getResourceRecordOptions := sess.NewGetResourceRecordOptions(id_set[0], id_set[1], id_set[2])
	_, _, err = sess.GetResourceRecord(getResourceRecordOptions)

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
