package ibm

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	pdnsVpcCRN                     = "vpc_crn"
	pdnsNetworkType                = "type"
	pdnsPermittedNetworkID         = "permitted_network_id"
	pdnsPermittedNetworkCreatedOn  = "created_on"
	pdnsPermittedNetworkModifiedOn = "modified_on"
	pdnsPermittedNetworkState      = "state"
)

var allowedNetworkTypes = []string{
	"vpc",
}

func resourceIBMPrivateDNSPermittedNetwork() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMPrivateDnsPermittedNetworkCreate,
		Read:     resourceIBMPrivateDnsPermittedNetworkRead,
		Delete:   resourceIBMPrivateDnsPermittedNetworkDelete,
		Exists:   resourceIBMPrivateDnsPermittedNetworkExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			pdnsPermittedNetworkID: {
				Type:     schema.TypeString,
				Computed: true,
			},

			pdnsInstanceID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			pdnsZoneID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			pdnsNetworkType: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: func(val interface{}, field string) (warnings []string, errors []error) {
					value := val.(string)
					for _, rtype := range allowedNetworkTypes {
						if value == rtype {
							return
						}
					}

					errors = append(
						errors,
						fmt.Errorf("%s is not one of the valid domain record types: %s",
							value, strings.Join(allowedNetworkTypes, ", "),
						),
					)
					return
				},
			},

			pdnsVpcCRN: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			pdnsPermittedNetworkCreatedOn: {
				Type:     schema.TypeString,
				Computed: true,
			},

			pdnsPermittedNetworkModifiedOn: {
				Type:     schema.TypeString,
				Computed: true,
			},

			pdnsPermittedNetworkState: {
				Type:     schema.TypeString,
				Computed: true,
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
	nwType := d.Get(pdnsNetworkType).(string)

	createPermittedNetworkOptions := sess.NewCreatePermittedNetworkOptions(instanceID, zoneID)
	permittedNetworkCrn, err := sess.NewPermittedNetworkVpc(vpcCRN)
	if err != nil {
		return err
	}

	createPermittedNetworkOptions.SetPermittedNetwork(permittedNetworkCrn)
	createPermittedNetworkOptions.SetType(nwType)
	response, _, err := sess.CreatePermittedNetwork(createPermittedNetworkOptions)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", instanceID, zoneID, *response.ID))
	d.Set(pdnsPermittedNetworkID, *response.ID)

	return resourceIBMPrivateDnsPermittedNetworkRead(d, meta)
}

func resourceIBMPrivateDnsPermittedNetworkRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")
	getPermittedNetworkOptions := sess.NewGetPermittedNetworkOptions(id_set[0], id_set[1], id_set[2])
	response, _, err := sess.GetPermittedNetwork(getPermittedNetworkOptions)

	if err != nil {
		return err
	}

	d.Set("id", response.ID)
	d.Set(pdnsInstanceID, id_set[0])
	d.Set(pdnsZoneID, id_set[1])
	d.Set(pdnsPermittedNetworkID, response.ID)
	d.Set(pdnsPermittedNetworkCreatedOn, response.CreatedOn)
	d.Set(pdnsPermittedNetworkModifiedOn, response.ModifiedOn)
	d.Set(pdnsVpcCRN, response.PermittedNetwork)
	d.Set(pdnsNetworkType, response.Type)
	d.Set(pdnsPermittedNetworkState, response.State)

	return nil
}

func resourceIBMPrivateDnsPermittedNetworkDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")
	deletePermittedNetworkOptions := sess.NewDeletePermittedNetworkOptions(id_set[0], id_set[1], id_set[2])
	_, _, err = sess.DeletePermittedNetwork(deletePermittedNetworkOptions)

	if err != nil {
		return err
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
	_, response, err := sess.GetPermittedNetwork(getPermittedNetworkOptions)
	if err != nil && response.StatusCode != 404 {
		if response.StatusCode == 404 {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
