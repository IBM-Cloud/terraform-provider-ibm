package ibm

import (
	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv2"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceIBMIAMRole() *schema.Resource {
	return &schema.Resource{
		Read: datasourceIBMIAMRoleRead,

		Schema: map[string]*schema.Schema{
			"service": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Service Name",
				ForceNew:    true,
			},
			"roles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}

}

func datasourceIBMIAMRoleRead(d *schema.ResourceData, meta interface{}) error {
	iampapv2Client, err := meta.(ClientSession).IAMPAPAPIV2()
	if err != nil {
		return err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	var serviceName string
	var customRoles, serviceRoles, systemRoles []iampapv2.Role
	if service, ok := d.GetOk("service"); ok {
		serviceName = service.(string)

		customRoles, err = iampapv2Client.IAMRoles().ListCustomRoles(userDetails.userAccount, serviceName)
		if err != nil {
			return err
		}

		serviceRoles, err = iampapv2Client.IAMRoles().ListServiceRoles(serviceName)
		if err != nil {
			return err
		}
	}

	d.SetId(userDetails.userAccount)

	systemRoles, err = iampapv2Client.IAMRoles().ListSystemDefinedRoles()
	if err != nil {
		return err
	}

	var roles []map[string]string

	roles = append(flattenRoleData(systemRoles, "platform"), append(flattenRoleData(serviceRoles, "service"), flattenRoleData(customRoles, "custom")...)...)

	d.Set("roles", roles)

	return nil
}
