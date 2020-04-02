package ibm

import (
	"fmt"
	"strings"
	"time"

	//"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
)

var allowedPrivateDomainRecordTypes = []string{
	"A", "AAAA", "CNAME", "MX", "PTR", "SRV", "TXT",
}

const (
	pdnsRecordType  = "type"
	pdnsRecordTTL   = "ttl"
	pdnsIPv4Address = "ipv4_address"
	pdnsRecordName  = "name"
)

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

			pdnsRecordTTL: {
				Type:     schema.TypeInt,
				Required: true,
				DefaultFunc: func() (interface{}, error) {
					return 900, nil
				},
			},

			pdnsRecordName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			pdnsIPv4Address: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			pdnsRecordType: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: func(val interface{}, field string) (warnings []string, errors []error) {
					value := val.(string)
					for _, rtype := range allowedPrivateDomainRecordTypes {
						if value == rtype {
							return
						}
					}

					errors = append(
						errors,
						fmt.Errorf("%s is not one of the valid domain record types: %s",
							value, strings.Join(allowedPrivateDomainRecordTypes, ", "),
						),
					)
					return
				},
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
	recordType := d.Get(pdnsRecordType).(string)
	ttl := d.Get(pdnsRecordTTL).(int)
	ipv4 := d.Get(pdnsIPv4Address).(string)
	name := d.Get(pdnsRecordName).(string)

	createResourceRecordOptions := sess.NewCreateResourceRecordOptions(instanceID, zoneID)
	createResourceRecordOptions.SetName(name)
	createResourceRecordOptions.SetType(recordType)
	createResourceRecordOptions.SetTTL(int64(ttl))
	resourceRecordAData, _ := sess.NewResourceRecordInputRdataRdataARecord(ipv4)
	createResourceRecordOptions.SetRdata(resourceRecordAData)
	response, _, err := sess.CreateResourceRecord(createResourceRecordOptions)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", instanceID, zoneID, *response.ID))

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
