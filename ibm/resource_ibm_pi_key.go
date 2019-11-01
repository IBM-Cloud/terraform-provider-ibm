package ibm

import (
	"fmt"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"time"
)

func resourceIBMPIKey() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMPIKeyCreate,
		Read:     resourceIBMPIKeyRead,
		Update:   resourceIBMPIKeyUpdate,
		Delete:   resourceIBMPIKeyDelete,
		Exists:   resourceIBMPIKeyExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			helpers.PIKeyId: {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
				Optional: true,
			},

			helpers.PIKeyName: {
				Type:     schema.TypeString,
				Required: true,
			},

			helpers.PIKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			helpers.PIKeyDate: {
				Type:     schema.TypeString,
				Optional: true,
			},

			helpers.PICloudInstanceId: {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIBMPIKeyCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}

	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	name := d.Get(helpers.PIKeyName).(string)
	sshkey := d.Get(helpers.PIKey).(string)
	client := st.NewIBMPIKeyClient(sess, powerinstanceid)

	sshResponse, _, err := client.Create(name, sshkey, powerinstanceid)

	if err != nil {
		return err
	}

	log.Printf("Printing the sshkey %+v", &sshResponse)

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)

	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}

	return resourceIBMPIKeyRead(d, meta)
}

func resourceIBMPIKeyRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	sshkeyC := st.NewIBMPIKeyClient(sess, powerinstanceid)
	sshkeydata, err := sshkeyC.Get(d.Get("name").(string), powerinstanceid)

	if err != nil {
		return err
	}

	d.Set("name", sshkeydata.Name)
	d.Set("sshkey", sshkeydata.SSHKey)
	d.Set("creationdate", sshkeydata.CreationDate)

	return nil

}

func resourceIBMPIKeyUpdate(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMPIKeyDelete(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMPIKeyExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return false, err
	}
	//id := d.Id()
	name := d.Get("name")
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	client := st.NewIBMPIKeyClient(sess, powerinstanceid)

	key, err := client.Get(d.Get("name").(string), powerinstanceid)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}
	return key.Name == name, nil
}
