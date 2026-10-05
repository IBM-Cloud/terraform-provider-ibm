// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// rhaiiServiceName is the global catalog name of the Red Hat AI Inference service.
	rhaiiServiceName = "instructlab"
	// rhaiiDefaultPlan is the only plan that the service offers today.
	rhaiiDefaultPlan = "instructlab-pricing-plan"
	// rhaiiDefaultLocation is the only region that the service is deployed to today.
	rhaiiDefaultLocation = "us-east"
)

// rhaiiProjectSchema holds the attributes that are specific to an RHAII project
// and that are added on top of the resource controller instance schema.
func rhaiiProjectSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"project_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The ID of the Red Hat AI Inference project. This is the value to use as `project_id` in the Red Hat AI Inference API.",
		},
		"endpoint": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The base URL of the Red Hat AI Inference API for this project.",
		},
	}
}

// rhaiiProjectEndpoint builds the Red Hat AI Inference API base URL for a project.
func rhaiiProjectEndpoint(location, projectID string) string {
	return fmt.Sprintf("https://%s.rhai.ibm.com/v1/projects/%s", location, projectID)
}

// setRhaiiProjectAttributes sets the RHAII specific attributes after the
// resource controller attributes were read. It fails when the instance is not
// an RHAII project, for example when a wrong instance is imported.
func setRhaiiProjectAttributes(d *schema.ResourceData) error {
	if service := d.Get("service").(string); service != rhaiiServiceName {
		return fmt.Errorf("[ERROR] Resource instance %s is not a Red Hat AI Inference project (service is %q, expected %q)", d.Id(), service, rhaiiServiceName)
	}

	projectID := d.Get("guid").(string)
	location := d.Get("location").(string)

	if err := d.Set("project_id", projectID); err != nil {
		return fmt.Errorf("[ERROR] Error setting project_id: %s", err)
	}
	if projectID != "" && location != "" {
		if err := d.Set("endpoint", rhaiiProjectEndpoint(location, projectID)); err != nil {
			return fmt.Errorf("[ERROR] Error setting endpoint: %s", err)
		}
	}
	return nil
}
