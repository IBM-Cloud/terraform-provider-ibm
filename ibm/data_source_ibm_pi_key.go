package ibm

import (
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.ibm.com/Bluemix/power-go-client/clients/instance"
	"github.ibm.com/Bluemix/power-go-client/helpers"
	"log"
)

func dataSourceIBMPIKey() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMPIKeysRead,
		Schema: map[string]*schema.Schema{

			helpers.PIKeyName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "SSHKey Name to be used for pvminstances",
				ValidateFunc: validation.NoZeroValues,
			},
			helpers.PICloudInstanceId: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			//Computed Attributes
			"creationdate": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sshkey": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMPIKeysRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).IBMPISession()

	if err != nil {
		return err
	}

	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	log.Printf("Calling the ibm-pi-key datasource with the %s instanceid ", powerinstanceid)
	sshkeyC := instance.NewIBMPIKeyClient(sess, powerinstanceid)
	sshkeydata, err := sshkeyC.Get(d.Get(helpers.PIKeyName).(string), powerinstanceid)

	if err != nil {
		return err
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)
	d.Set("name", sshkeydata.Name)
	d.Set("sshkey", sshkeydata.SSHKey)
	d.Set("creationdate", sshkeydata.CreationDate)

	return nil
	//return fmt.Errorf("No Image found with name %s", imagedata.)

}
