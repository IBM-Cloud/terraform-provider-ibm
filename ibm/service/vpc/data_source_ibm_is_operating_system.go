// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isOperatingSystemName                   = "name"
	isOperatingSystemArchitecture           = "architecture"
	isOperatingSystemDHOnly                 = "dedicated_host_only"
	isOperatingSystemDisplayName            = "display_name"
	isOperatingSystemFamily                 = "family"
	isOperatingSystemHref                   = "href"
	isOperatingSystemVendor                 = "vendor"
	isOperatingSystemVersion                = "version"
	isOperatingSystemAllowUserImageCreation = "allow_user_image_creation"
	isOperatingSystemUserDataFormat         = "user_data_format"
)

func DataSourceIBMISOperatingSystem() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISOperatingSystemRead,

		Schema: map[string]*schema.Schema{
			isOperatingSystemAllowUserImageCreation: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Users may create new images with this operating system",
			},
			isOperatingSystemName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The globally unique name for this operating system",
			},

			isOperatingSystemArchitecture: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The operating system architecture",
			},

			isOperatingSystemVersion: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The major release version of this operating system",
			},
			isOperatingSystemDHOnly: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Flag which shows images with this operating system can only be used on dedicated hosts or dedicated host groups",
			},
			isOperatingSystemDisplayName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A unique, display-friendly name for the operating system",
			},
			isOperatingSystemFamily: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the software family this operating system belongs to",
			},
			isOperatingSystemHref: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this operating system",
			},

			isOperatingSystemVendor: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The vendor of the operating system",
			},
			isOperatingSystemUserDataFormat: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user data format for this operating system",
			},
		},
	}
}

func dataSourceIBMISOperatingSystemRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get(isOperatingSystemName).(string)
	err := osGet(context, d, meta, name)
	if err != nil {
		return err
	}
	return nil
}

func osGet(context context.Context, d *schema.ResourceData, meta interface{}, name string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_operating_system", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	getOperatingSystemOptions := &vpcv1.GetOperatingSystemOptions{
		Name: &name,
	}
	operatingSystem, _, err := sess.GetOperatingSystemWithContext(context, getOperatingSystemOptions)
	if err != nil || operatingSystem == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetOperatingSystemWithContext failed: %s", err.Error()), "(Data) ibm_is_operating_system", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set("name", operatingSystem.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_operating_system", "read", "set-name").GetDiag()
	}
	d.SetId(*operatingSystem.Name)
	if err = d.Set("dedicated_host_only", operatingSystem.DedicatedHostOnly); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dedicated_host_only: %s", err), "(Data) ibm_is_operating_system", "read", "set-dedicated_host_only").GetDiag()
	}
	if err = d.Set("architecture", operatingSystem.Architecture); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting architecture: %s", err), "(Data) ibm_is_operating_system", "read", "set-architecture").GetDiag()
	}
	if err = d.Set("display_name", operatingSystem.DisplayName); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting display_name: %s", err), "(Data) ibm_is_operating_system", "read", "set-display_name").GetDiag()
	}
	if err = d.Set("family", operatingSystem.Family); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting family: %s", err), "(Data) ibm_is_operating_system", "read", "set-family").GetDiag()
	}
	if err = d.Set("href", operatingSystem.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_operating_system", "read", "set-href").GetDiag()
	}
	if err = d.Set("vendor", operatingSystem.Vendor); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vendor: %s", err), "(Data) ibm_is_operating_system", "read", "set-vendor").GetDiag()
	}
	if err = d.Set("version", operatingSystem.Version); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting version: %s", err), "(Data) ibm_is_operating_system", "read", "set-version").GetDiag()
	}
	if err = d.Set("allow_user_image_creation", operatingSystem.AllowUserImageCreation); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting allow_user_image_creation: %s", err), "(Data) ibm_is_operating_system", "read", "set-allow_user_image_creation").GetDiag()
	}
	if err = d.Set("user_data_format", operatingSystem.UserDataFormat); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting user_data_format: %s", err), "(Data) ibm_is_operating_system", "read", "set-user_data_format").GetDiag()
	}
	return nil
}
