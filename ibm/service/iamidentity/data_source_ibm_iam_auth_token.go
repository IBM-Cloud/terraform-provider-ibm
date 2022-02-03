// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity

import (
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMIAMAuthToken() *schema.Resource {
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
	bmxSess, err := meta.(conns.ClientSession).BluemixSession()
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
