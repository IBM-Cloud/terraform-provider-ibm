package ibm

import (
	"fmt"
	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"reflect"
)

func resourceIBMCISHealthCheck() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cis_id": {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
			},
			"path": {
				Type:        schema.TypeString,
				Description: "path",
				Optional:    true,
				Default:     "/",
			},
			"expected_body": {
				Type:        schema.TypeString,
				Description: "expected_body",
				Required:    true,
			},
			"expected_codes": {
				Type:        schema.TypeString,
				Description: "expected_codes",
				Required:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "description",
				Optional:    true,
			},
			"type": {
				Type:         schema.TypeString,
				Description:  "type",
				Optional:     true,
				Default:      "http",
				ValidateFunc: validateAllowedStringValue([]string{"http", "https"}),
			},
			"method": {
				Type:        schema.TypeString,
				Description: "method",
				Optional:    true,
				Default:     "GET",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Description: "timeout",
				Optional:    true,
				Default:     5,
				//ValidateFunc: validation.IntBetween(1, 10),
			},
			"retries": {
				Type:        schema.TypeInt,
				Description: "retries",
				Optional:    true,
				Default:     2,
				//ValidateFunc: validation.IntBetween(1, 5),
			},
			"interval": {
				Type:        schema.TypeInt,
				Description: "interval",
				Optional:    true,
				Default:     60,
			},
			"follow_redirects": {
				Type:        schema.TypeBool,
				Description: "follow_redirects",
				Optional:    true,
				Default:     true,
			},
			"allow_insecure": {
				Type:        schema.TypeBool,
				Description: "allow_insecure",
				Optional:    true,
				Default:     true,
			},
			"created_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// Supporting code omitted
			// "header": {
			//     Type:     schema.TypeSet,
			//     Optional: true,
			//     Elem: &schema.Resource{
			//         Schema: map[string]*schema.Schema{
			//             "header": {
			//                 Type:     schema.TypeString,
			//                 Required: true,
			//             },
			//             "values": {
			//                 Type:     schema.TypeSet,
			//                 Required: true,
			//                 Elem: &schema.Schema{
			//                     Type: schema.TypeString,
			//                 },
			//             },
			//         },
			//     },
			//     Set: HashByMapKey("header"),

		},

		Create: resourceCIShealthCheckCreate,
		Read:   resourceCIShealthCheckRead,
		Update: resourceCIShealthCheckUpdate,
		Delete: resourceCIShealthCheckDelete,
	}
}

func resourceCIShealthCheckCreate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	log.Printf("   client %v\n", cisClient)
	if err != nil {
		return err
	}

	cisId := d.Get("cis_id").(string)
	monitorPath := d.Get("path").(string)
	expCodes := d.Get("expected_codes").(string)
	expBody := d.Get("expected_body").(string)

	var monitors *[]v1.Monitor
	monitors, err = cisClient.Monitors().ListMonitors(cisId)
	if err != nil {
		log.Printf("ListMonitors Failed %s\n", err)
		return err
	}

	var monitorPaths []string
	monitorsObj := *monitors
	for _, monitor := range monitorsObj {
		monitorPaths = append(monitorPaths, monitor.Path)
	}

	log.Println(string("Existing monitor paths"))
	log.Println(monitorPaths)

	var monitor *v1.Monitor
	var monitorObj v1.Monitor

	index := indexOf(monitorPath, monitorPaths)
	if index == -1 {

		monitorNew := v1.MonitorBody{
			ExpCodes: expCodes,
			ExpBody:  expBody,
			Path:     monitorPath,
		}

		if monType, ok := d.GetOk("type"); ok {
			monitorNew.MonType = monType.(string)
		}
		if method, ok := d.GetOk("method"); ok {
			monitorNew.Method = method.(string)
		}
		if timeout, ok := d.GetOk("timeout"); ok {
			monitorNew.Timeout = timeout.(int)
		}
		if retries, ok := d.GetOk("retries"); ok {
			monitorNew.Retries = retries.(int)
		}
		if interval, ok := d.GetOk("interval"); ok {
			monitorNew.Interval = interval.(int)
		}
		if follow_redirects, ok := d.GetOk("follow_redirects"); ok {
			monitorNew.FollowRedirects = follow_redirects.(bool)
		}
		if allow_insecure, ok := d.GetOk("allow_insecure"); ok {
			monitorNew.AllowInsecure = allow_insecure.(bool)
		}

		monitor, err = cisClient.Monitors().CreateMonitor(cisId, monitorNew)
		if err != nil {
			log.Printf("CreateMonitors Failed %s\n", err)
			return err
		}
		monitorObj = *monitor
	} else {
		return fmt.Errorf("Resource with name %s already exists", monitorPath)
	}

	d.SetId(monitorObj.Id)
	d.Set("path", monitorObj.Path)

	return resourceCIShealthCheckRead(d, meta)
}

func resourceCIShealthCheckRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}

	var monitorId string

	monitorId = d.Id()
	cisId := d.Get("cis_id").(string)

	log.Printf("resourceCIShealthCheckRead - Getting Monitor %v\n", monitorId)
	var monitor *v1.Monitor

	monitor, err = cisClient.Monitors().GetMonitor(cisId, monitorId)
	if err != nil {
		log.Printf("resourceCIhealthCheckRead - ListMonitors Failed %s\n", err)
		return err
	} else {
		log.Printf("resourceCIShealthCheckRead - Retrieved Monitor %v\n", monitor)

		monitorObj := *monitor
		d.Set("path", monitorObj.Path)
		d.Set("expected_body", monitorObj.ExpBody)
		d.Set("expected_codes", monitorObj.ExpCodes)
		d.Set("type", monitorObj.MonType)
		d.Set("method", monitorObj.Method)
		d.Set("timeout", monitorObj.Timeout)
		d.Set("retries", monitorObj.Retries)
		d.Set("interval", monitorObj.Interval)
		d.Set("follow_redirects", monitorObj.FollowRedirects)
		d.Set("allow_insecure", monitorObj.AllowInsecure)
		// }
	}
	return nil
}

func resourceCIShealthCheckUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceCIShealthCheckRead(d, meta)
}

func resourceCIShealthCheckDelete(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	monitorId := d.Id()
	cisId := d.Get("cis_id").(string)
	var monitor *v1.Monitor
	emptyMonitor := new(v1.Monitor)

	log.Println("Getting Monitor to delete")
	monitor, err = cisClient.Monitors().GetMonitor(cisId, monitorId)
	if err != nil {
		log.Printf("GetMonitor Failed %s\n", err)
		return err
	}

	monitorObj := *monitor
	if !reflect.DeepEqual(emptyMonitor, monitorObj) {
		log.Println("Deleting Monitor")
		err = cisClient.Monitors().DeleteMonitor(cisId, monitorId)
		if err != nil {
			log.Printf("DeleteMonitor Failed %s\n", err)
			return err
		}
	}

	d.SetId("")
	return nil
}
