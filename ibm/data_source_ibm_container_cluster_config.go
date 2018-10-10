package ibm

import (
	"fmt"
	"log"
	"path/filepath"

	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/helpers"
	"github.com/hashicorp/terraform/helper/schema"
	homedir "github.com/mitchellh/go-homedir"
)

func dataSourceIBMContainerClusterConfig() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMContainerClusterConfigRead,

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
			"resource_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the resource group.",
			},
			"cluster_name_id": {
				Description: "The name/id of the cluster",
				Type:        schema.TypeString,
				Required:    true,
			},
			"config_dir": {
				Description: "The directory where the cluster config to be downloaded. Default is home directory ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"download": {
				Description: "If set to false will not download the config, otherwise they are downloaded each time but onto the same path for a given cluster name/id",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"admin": {
				Description: "If set to true will download the config for admin",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"config_file_path": {
				Description: "The absolute path to the kubernetes config yml file ",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceIBMContainerClusterConfigRead(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	csAPI := csClient.Clusters()
	name := d.Get("cluster_name_id").(string)
	download := d.Get("download").(bool)
	admin := d.Get("admin").(bool)
	configDir := d.Get("config_dir").(string)

	if len(configDir) == 0 {
		configDir, err = homedir.Dir()
		if err != nil {
			return fmt.Errorf("Error fetching homedir: %s", err)
		}
	}
	var configPath string
	if !download {
		log.Println("Skipping download of the cluster config", "Going to check if it already exists")
		expectedDir := v1.ComputeClusterConfigDir(configDir, name, admin)
		configPath = filepath.Join(expectedDir, "config.yml")
		if !helpers.FileExists(configPath) {
			return fmt.Errorf(`Couldn't  find the cluster config at expected path %s. Please set "download" to true to download the new config`, configPath)
		}

	} else {
		targetEnv, err := getClusterTargetHeader(d, meta)
		if err != nil {
			return err
		}
		configPath, err = csAPI.GetClusterConfig(name, configDir, admin, targetEnv)
		if err != nil {
			return fmt.Errorf("Error downloading the cluster config [%s]: %s", name, err)
		}
	}

	d.SetId(name)
	d.Set("config_dir", configDir)
	d.Set("config_file_path", configPath)
	return nil
}
