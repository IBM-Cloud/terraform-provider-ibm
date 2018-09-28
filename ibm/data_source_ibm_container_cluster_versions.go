package ibm

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceIBMContainerClusterVersions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMContainerClusterVersionsRead,

		Schema: map[string]*schema.Schema{
			"org_guid": {
				Description: "The bluemix organization guid this cluster belongs to",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"space_guid": {
				Description: "The bluemix space guid this cluster belongs to",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"account_guid": {
				Description: "The bluemix account guid this cluster belongs to",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The cluster region",
			},
			"valid_kube_versions": {
				Description: "List supported kube-versions",
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceIBMContainerClusterVersionsRead(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	verAPI := csClient.KubeVersions()
	targetEnv, err := getClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}

	availableVersions, _ := verAPI.List(targetEnv)
	versions := make([]string, len(availableVersions))
	for i, version := range availableVersions {
		versions[i] = fmt.Sprintf("%d%s%d%s%d", version.Major, ".", version.Minor, ".", version.Patch)
	}
	d.SetId(time.Now().UTC().String())
	d.Set("valid_kube_versions", versions)
	return nil
}
