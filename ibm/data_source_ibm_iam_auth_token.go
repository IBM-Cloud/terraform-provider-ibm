/*
* IBM Confidential
* Object Code Only Source Materials
* 5747-SM3
* (c) Copyright IBM Corp. 2017,2021
*
* The source code for this program is not published or otherwise divested
* of its trade secrets, irrespective of what has been deposited with the
* U.S. Copyright Office.
 */

package ibm

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIBMIAMAuthToken() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMIAMAuthTokenRead,

		Schema: map[string]*schema.Schema{

			"iam_access_token": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"iam_refresh_token": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uaa_access_token": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"uaa_refresh_token": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMIAMAuthTokenRead(d *schema.ResourceData, meta interface{}) error {
	bmxSess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return err
	}
	d.SetId(dataSourceIBMIAMAuthTokenID(d))

	d.Set("iam_access_token", bmxSess.Config.IAMAccessToken)
	d.Set("iam_refresh_token", bmxSess.Config.IAMRefreshToken)
	d.Set("uaa_access_token", bmxSess.Config.UAAAccessToken)
	d.Set("uaa_refresh_token", bmxSess.Config.UAARefreshToken)

	return nil
}

func dataSourceIBMIAMAuthTokenID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
