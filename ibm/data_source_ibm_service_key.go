package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceIBMServiceKey() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMServiceKeyRead,

		Schema: map[string]*schema.Schema{
			"credentials": {
				Description: "Credentials asociated with the key",
				Type:        schema.TypeMap,
				Computed:    true,
			},

			"name": {
				Description: "The name of the service key",
				Type:        schema.TypeString,
				Required:    true,
			},
			"service_instance_name": {
				Description: "Service instance name for example, cleardbinstance",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceIBMServiceKeyRead(d *schema.ResourceData, meta interface{}) error {
	cfClient, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return err
	}
	siAPI := cfClient.ServiceInstances()
	skAPI := cfClient.ServiceKeys()
	serviceInstanceName := d.Get("service_instance_name").(string)
	name := d.Get("name").(string)
	inst, err := siAPI.FindByName(serviceInstanceName)
	if err != nil {
		return err
	}
	serviceInstance, err := siAPI.Get(inst.GUID)
	if err != nil {
		return fmt.Errorf("Error retrieving service: %s", err)
	}
	serviceKey, err := skAPI.FindByName(serviceInstance.Metadata.GUID, name)
	if err != nil {
		return fmt.Errorf("Error retrieving service key: %s", err)
	}
	d.SetId(serviceKey.GUID)
	d.Set("credentials", serviceKey.Credentials)

	return nil
}
