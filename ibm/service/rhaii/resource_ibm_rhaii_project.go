// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii

import (
	"context"
	"maps"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/resourcecontroller"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceIBMRhaiiProject manages a Red Hat AI Inference (RHAII) project.
// A project is a resource controller instance of the "instructlab" service,
// so create, update and delete reuse the resource controller implementation.
func ResourceIBMRhaiiProject() *schema.Resource {
	// Clone to avoid mutating the shared resource-controller schema map.
	rhaiiSchema := maps.Clone(resourcecontroller.ResourceIBMResourceInstance().Schema)

	// The service is always "instructlab", so it is not user configurable.
	rhaiiSchema["service"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
		Description: "The service type of the instance. Always `instructlab`.",
	}

	rhaiiSchema["plan"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Default:     rhaiiDefaultPlan,
		Description: "The pricing plan of the project.",
	}

	location := rhaiiSchema["location"]
	rhaiiSchema["location"] = &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		ForceNew:     true,
		Default:      rhaiiDefaultLocation,
		ValidateFunc: location.ValidateFunc,
		Description:  "The region where the project is created.",
	}

	maps.Copy(rhaiiSchema, rhaiiProjectSchema())

	return &schema.Resource{
		Create: resourceIBMRhaiiProjectCreate,
		Read:   resourceIBMRhaiiProjectRead,
		Update: resourceIBMRhaiiProjectUpdate,
		Delete: resourcecontroller.ResourceIBMResourceInstanceDelete,
		Exists: resourcecontroller.ResourceIBMResourceInstanceExists,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, _ interface{}) error {
				return flex.ResourceTagsCustomizeDiff(diff)
			},
		),

		Schema: rhaiiSchema,
	}
}

func resourceIBMRhaiiProjectCreate(d *schema.ResourceData, meta interface{}) error {
	if err := d.Set("service", rhaiiServiceName); err != nil {
		return flex.FmtErrorf("[ERROR] Error setting service: %s", err)
	}
	if err := resourcecontroller.ResourceIBMResourceInstanceCreate(d, meta); err != nil {
		return err
	}
	return setRhaiiProjectAttributes(d)
}

func resourceIBMRhaiiProjectRead(d *schema.ResourceData, meta interface{}) error {
	if err := resourcecontroller.ResourceIBMResourceInstanceRead(d, meta); err != nil {
		return err
	}
	return setRhaiiProjectAttributes(d)
}

func resourceIBMRhaiiProjectUpdate(d *schema.ResourceData, meta interface{}) error {
	if err := d.Set("service", rhaiiServiceName); err != nil {
		return flex.FmtErrorf("[ERROR] Error setting service: %s", err)
	}
	if err := resourcecontroller.ResourceIBMResourceInstanceUpdate(d, meta); err != nil {
		return err
	}
	return setRhaiiProjectAttributes(d)
}
