package ibm

import (
	"log"

	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	cisGLBHealthCheckID              = "monitor_id"
	cisGLBHealthCheckPath            = "path"
	cisGLBHealthCheckPort            = "port"
	cisGLBHealthCheckExpectedBody    = "expected_body"
	cisGLBHealthCheckExpectedCodes   = "expected_codes"
	cisGLBHealthCheckDesc            = "description"
	cisGLBHealthCheckType            = "type"
	cisGLBHealthCheckMethod          = "method"
	cisGLBHealthCheckTimeout         = "timeout"
	cisGLBHealthCheckRetries         = "retries"
	cisGLBHealthCheckInterval        = "interval"
	cisGLBHealthCheckFollowRedirects = "follow_redirects"
	cisGLBHealthCheckAllowInsecure   = "allow_insecure"
	cisGLBHealthCheckCreatedOn       = "create_on"
	cisGLBHealthCheckModifiedOn      = "modified_on"
)

func resourceIBMCISHealthCheck() *schema.Resource {
	return &schema.Resource{

		Create:   resourceCISHealthCheckCreate,
		Read:     resourceCISHealthCheckRead,
		Update:   resourceCISHealthCheckUpdate,
		Delete:   resourceCISHealthCheckDelete,
		Exists:   resourceCISHealthCheckExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
			},
			cisGLBHealthCheckID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GLB Monitor/Health check id",
			},
			cisGLBHealthCheckPath: {
				Type:         schema.TypeString,
				Description:  "path",
				Optional:     true,
				ValidateFunc: validateURLPath,
			},
			cisGLBHealthCheckExpectedBody: {
				Type:        schema.TypeString,
				Description: "expected_body",
				Optional:    true,
			},
			cisGLBHealthCheckExpectedCodes: {
				Type:        schema.TypeString,
				Description: "expected_codes",
				Optional:    true,
			},
			cisGLBHealthCheckDesc: {
				Type:        schema.TypeString,
				Description: "description",
				Optional:    true,
			},
			cisGLBHealthCheckType: {
				Type:        schema.TypeString,
				Description: "type",
				Optional:    true,
				Default:     "http",
				ValidateFunc: validateAllowedStringValue(
					[]string{
						"http",
						"https",
						"tcp",
					},
				),
			},
			cisGLBHealthCheckMethod: {
				Type:        schema.TypeString,
				Description: "method",
				Optional:    true,
				ValidateFunc: validateAllowedStringValue(
					[]string{
						"GET",
						"HEAD",
					},
				),
			},
			cisGLBHealthCheckTimeout: {
				Type:         schema.TypeInt,
				Description:  "timeout",
				Optional:     true,
				Default:      2,
				ValidateFunc: validateTimeout,
			},
			cisGLBHealthCheckRetries: {
				Type:        schema.TypeInt,
				Description: "retries",
				Optional:    true,
				Default:     1,
			},
			cisGLBHealthCheckInterval: {
				Type:        schema.TypeInt,
				Description: "interval",
				Optional:    true,
				Default:     validateInterval,
			},
			cisGLBHealthCheckFollowRedirects: {
				Type:        schema.TypeBool,
				Description: "follow_redirects",
				Optional:    true,
			},
			cisGLBHealthCheckAllowInsecure: {
				Type:        schema.TypeBool,
				Description: "allow_insecure",
				Optional:    true,
				Default:     false,
			},
			cisGLBHealthCheckCreatedOn: {
				Type:     schema.TypeString,
				Computed: true,
			},
			cisGLBHealthCheckModifiedOn: {
				Type:     schema.TypeString,
				Computed: true,
			},
			cisGLBHealthCheckPort: {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceCISHealthCheckCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).CisGLBHealthCheckClientSession()
	if err != nil {
		return err
	}

	crn := d.Get(cisID).(string)
	log.Printf("\n\n crn : %s \n\n", crn)
	sess.Crn = core.StringPtr(crn)

	opt := sess.NewCreateLoadBalancerMonitorOptions()

	if expCodes, ok := d.GetOk(cisGLBHealthCheckExpectedCodes); ok {
		opt.SetExpectedCodes(expCodes.(string))
	}
	if expBody, ok := d.GetOk(cisGLBHealthCheckExpectedBody); ok {
		opt.SetExpectedBody(expBody.(string))
	}
	if monPath, ok := d.GetOk(cisGLBHealthCheckPath); ok {
		opt.SetPath(monPath.(string))
	}
	if description, ok := d.GetOk(cisGLBHealthCheckDesc); ok {
		opt.SetDescription(description.(string))
	}
	if method, ok := d.GetOk(cisGLBHealthCheckMethod); ok {
		opt.SetMethod(method.(string))
	}
	if timeout, ok := d.GetOk(cisGLBHealthCheckTimeout); ok {
		opt.SetTimeout(int64(timeout.(int)))
	}
	if retries, ok := d.GetOk(cisGLBHealthCheckRetries); ok {
		opt.SetRetries(int64(retries.(int)))
	}
	if interval, ok := d.GetOk(cisGLBHealthCheckInterval); ok {
		opt.SetInterval(int64(interval.(int)))
	}
	if followRedirects, ok := d.GetOk(cisGLBHealthCheckFollowRedirects); ok {
		opt.SetFollowRedirects(followRedirects.(bool))
	}
	if allowInsecure, ok := d.GetOk(cisGLBHealthCheckAllowInsecure); ok {
		opt.SetAllowInsecure(allowInsecure.(bool))
	}
	if port, ok := d.GetOk(cisGLBHealthCheckPort); ok {
		opt.SetPort(int64(port.(int)))
	}

	result, resp, err := sess.CreateLoadBalancerMonitor(opt)
	if err != nil {
		log.Printf("create global load balancer health check failed %s", resp)
		return err
	}
	d.SetId(convertCisToTfTwoVar(*result.Result.ID, crn))
	d.Set(cisGLBHealthCheckID, result.Result.ID)
	d.Set(cisID, crn)
	d.Set(cisGLBHealthCheckDesc, result.Result.Description)
	d.Set(cisGLBHealthCheckPath, result.Result.Path)
	d.Set(cisGLBHealthCheckExpectedBody, result.Result.ExpectedBody)
	d.Set(cisGLBHealthCheckExpectedCodes, result.Result.ExpectedCodes)
	d.Set(cisGLBHealthCheckType, result.Result.Type)
	d.Set(cisGLBHealthCheckMethod, result.Result.Method)
	d.Set(cisGLBHealthCheckTimeout, result.Result.Timeout)
	d.Set(cisGLBHealthCheckRetries, result.Result.Retries)
	d.Set(cisGLBHealthCheckInterval, result.Result.Interval)
	d.Set(cisGLBHealthCheckFollowRedirects, result.Result.FollowRedirects)
	d.Set(cisGLBHealthCheckAllowInsecure, result.Result.AllowInsecure)
	d.Set(cisGLBHealthCheckCreatedOn, result.Result.CreatedOn)
	d.Set(cisGLBHealthCheckModifiedOn, result.Result.ModifiedOn)
	d.Set(cisGLBHealthCheckPort, result.Result.Port)

	return nil
}

func resourceCISHealthCheckRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).CisGLBHealthCheckClientSession()
	if err != nil {
		return err
	}

	monitorID, crn, err := convertTftoCisTwoVar(d.Id())
	if err != nil {
		return err
	}
	sess.Crn = core.StringPtr(crn)

	opt := sess.NewGetLoadBalancerMonitorOptions(monitorID)

	result, resp, err := sess.GetLoadBalancerMonitor(opt)
	if err != nil {
		log.Printf("Error reading global load balancer health check detail: %s", resp)
		return err
	}
	d.Set(cisGLBHealthCheckID, result.Result.ID)
	d.Set(cisID, crn)
	d.Set(cisGLBHealthCheckDesc, result.Result.Description)
	d.Set(cisGLBHealthCheckPath, result.Result.Path)
	d.Set(cisGLBHealthCheckExpectedBody, result.Result.ExpectedBody)
	d.Set(cisGLBHealthCheckExpectedCodes, result.Result.ExpectedCodes)
	d.Set(cisGLBHealthCheckType, result.Result.Type)
	d.Set(cisGLBHealthCheckMethod, result.Result.Method)
	d.Set(cisGLBHealthCheckTimeout, result.Result.Timeout)
	d.Set(cisGLBHealthCheckRetries, result.Result.Retries)
	d.Set(cisGLBHealthCheckInterval, result.Result.Interval)
	d.Set(cisGLBHealthCheckFollowRedirects, result.Result.FollowRedirects)
	d.Set(cisGLBHealthCheckAllowInsecure, result.Result.AllowInsecure)
	d.Set(cisGLBHealthCheckPort, result.Result.Port)
	d.Set(cisGLBHealthCheckCreatedOn, result.Result.CreatedOn)
	d.Set(cisGLBHealthCheckModifiedOn, result.Result.ModifiedOn)

	return nil
}

func resourceCISHealthCheckUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).CisGLBHealthCheckClientSession()
	if err != nil {
		return err
	}

	monitorID, crn, err := convertTftoCisTwoVar(d.Id())
	if err != nil {
		return err
	}
	sess.Crn = core.StringPtr(crn)

	opt := sess.NewEditLoadBalancerMonitorOptions(monitorID)
	if d.HasChange(cisGLBHealthCheckType) ||
		d.HasChange(cisGLBHealthCheckDesc) ||
		d.HasChange(cisGLBHealthCheckPort) ||
		d.HasChange(cisGLBHealthCheckExpectedCodes) ||
		d.HasChange(cisGLBHealthCheckExpectedCodes) ||
		d.HasChange(cisGLBHealthCheckMethod) ||
		d.HasChange(cisGLBHealthCheckTimeout) ||
		d.HasChange(cisGLBHealthCheckRetries) ||
		d.HasChange(cisGLBHealthCheckInterval) ||
		d.HasChange(cisGLBHealthCheckFollowRedirects) ||
		d.HasChange(cisGLBHealthCheckAllowInsecure) ||
		d.HasChange(cisGLBHealthCheckPort) {
		if monType, ok := d.GetOk(cisGLBHealthCheckType); ok {
			opt.SetType(monType.(string))
		}
		if expCodes, ok := d.GetOk(cisGLBHealthCheckExpectedCodes); ok {
			opt.SetExpectedCodes(expCodes.(string))
		}
		if expBody, ok := d.GetOk(cisGLBHealthCheckExpectedBody); ok {
			opt.SetExpectedBody(expBody.(string))
		}
		if monPath, ok := d.GetOk(cisGLBHealthCheckPath); ok {
			opt.SetPath(monPath.(string))
		}
		if description, ok := d.GetOk(cisGLBHealthCheckDesc); ok {
			opt.SetDescription(description.(string))
		}
		if method, ok := d.GetOk(cisGLBHealthCheckMethod); ok {
			opt.SetMethod(method.(string))
		}
		if timeout, ok := d.GetOk(cisGLBHealthCheckTimeout); ok {
			opt.SetTimeout(int64(timeout.(int)))
		}
		if retries, ok := d.GetOk(cisGLBHealthCheckRetries); ok {
			opt.SetRetries(int64(retries.(int)))
		}
		if interval, ok := d.GetOk(cisGLBHealthCheckInterval); ok {
			opt.SetInterval(int64(interval.(int)))
		}
		if followRedirects, ok := d.GetOk(cisGLBHealthCheckFollowRedirects); ok {
			opt.SetFollowRedirects(followRedirects.(bool))
		}
		if allowInsecure, ok := d.GetOk(cisGLBHealthCheckAllowInsecure); ok {
			opt.SetAllowInsecure(allowInsecure.(bool))
		}
		if port, ok := d.GetOk(cisGLBHealthCheckPort); ok {
			opt.SetPort(int64(port.(int)))
		}
		result, resp, err := sess.EditLoadBalancerMonitor(opt)
		if err != nil {
			log.Printf("Error updating global load balancer health check detail: %s", resp)
			return err
		}
		d.Set(cisGLBHealthCheckID, result.Result.ID)
		d.Set(cisID, crn)
		d.Set(cisGLBHealthCheckDesc, result.Result.Description)
		d.Set(cisGLBHealthCheckPath, result.Result.Path)
		d.Set(cisGLBHealthCheckExpectedBody, result.Result.ExpectedBody)
		d.Set(cisGLBHealthCheckExpectedCodes, result.Result.ExpectedCodes)
		d.Set(cisGLBHealthCheckType, result.Result.Type)
		d.Set(cisGLBHealthCheckMethod, result.Result.Method)
		d.Set(cisGLBHealthCheckTimeout, result.Result.Timeout)
		d.Set(cisGLBHealthCheckRetries, result.Result.Retries)
		d.Set(cisGLBHealthCheckInterval, result.Result.Interval)
		d.Set(cisGLBHealthCheckFollowRedirects, result.Result.FollowRedirects)
		d.Set(cisGLBHealthCheckAllowInsecure, result.Result.AllowInsecure)
		d.Set(cisGLBHealthCheckPort, result.Result.Port)
		d.Set(cisGLBHealthCheckCreatedOn, result.Result.CreatedOn)
		d.Set(cisGLBHealthCheckModifiedOn, result.Result.ModifiedOn)
	}

	return nil
}

func resourceCISHealthCheckDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).CisGLBHealthCheckClientSession()
	if err != nil {
		return err
	}

	monitorID, crn, err := convertTftoCisTwoVar(d.Id())
	if err != nil {
		return err
	}
	sess.Crn = core.StringPtr(crn)

	opt := sess.NewDeleteLoadBalancerMonitorOptions(monitorID)

	result, resp, err := sess.DeleteLoadBalancerMonitor(opt)
	if err != nil {
		log.Printf("Error deleting global load balancer health check detail: %s", resp)
		return err
	}
	log.Printf("Monitor ID: %s", *result.Result.ID)
	return nil
}

func resourceCISHealthCheckExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).CisGLBHealthCheckClientSession()
	if err != nil {
		return false, err
	}

	monitorID, crn, err := convertTftoCisTwoVar(d.Id())
	if err != nil {
		return false, err
	}
	sess.Crn = core.StringPtr(crn)

	opt := sess.NewGetLoadBalancerMonitorOptions(monitorID)

	result, resp, err := sess.GetLoadBalancerMonitor(opt)
	if err != nil {
		log.Printf("global load balancer health check does not exist: %s", resp)
		return false, err
	}
	log.Printf("global load balancer health check exists: %s", *result.Result.ID)
	return true, nil
}
