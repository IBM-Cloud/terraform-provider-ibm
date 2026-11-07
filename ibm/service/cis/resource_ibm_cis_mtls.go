// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"context"
	"crypto/md5"
	"encoding/pem"
	"fmt"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/mtlsv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	cisMtlsID            = "mtls_id"
	cisMtlsCert          = "certificate"
	cisMtlsHostNames     = "associated_hostnames"
	cisMtlsCertExpireOn  = "expires_on"
	cisMtlsZone          = "zone"
	cisMtlsZoneHostNames = "hostnames"
)

func ResourceIBMCISMtls() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMCISMtlsCreate,
		ReadContext:   resourceIBMCISMtlsRead,
		UpdateContext: resourceIBMCISMtlsUpdate,
		DeleteContext: resourceIBMCISMtlsDelete,
		Importer:      &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
				ValidateFunc: validate.InvokeValidator("ibm_cis_mtls",
					"cis_id"),
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Description:      "Associated CIS domain",
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{cisDomainID, cisMtlsZone},
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			cisMtlsID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Mtls transaction ID",
			},
			cisMtlsCert: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Certificate contents",
				Sensitive:   true,
				// The API never returns the PEM, so after import the configured value is matched against the fingerprint.
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					fingerprint := d.Get(cisMtlsCertFingerprint).(string)
					return old == "" && d.Id() != "" && fingerprint != "" && cisMtlsCertFingerprintOf(new) == fingerprint
				},
			},
			cisMtlsCertName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Certificate name",
			},
			cisMtlsHostNames: {
				Type:          schema.TypeList,
				Elem:          &schema.Schema{Type: schema.TypeString},
				Optional:      true,
				ConflictsWith: []string{cisMtlsZone},
				Description:   "Host name list to be associated",
			},
			cisMtlsZone: {
				Type:          schema.TypeList,
				Optional:      true,
				ConflictsWith: []string{cisMtlsHostNames},
				Description:   "Zones of the CIS instance and the host names to associate in each",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						cisDomainID: {
							Type:             schema.TypeString,
							Required:         true,
							Description:      "CIS domain ID",
							DiffSuppressFunc: suppressDomainIDDiff,
						},
						cisMtlsZoneHostNames: {
							Type:        schema.TypeSet,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Required:    true,
							MinItems:    1,
							Description: "Host names of this zone to be associated",
						},
					},
				},
			},
			cisMtlsCertFingerprint: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate fingerprint",
			},
			cisMtlsCertCreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate Created At",
			},
			cisMtlsCertUpdatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate Updated At",
			},
			cisMtlsCertExpireOn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate Expires on",
			},
			cisMtlsCertID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate ID",
			},
		},
	}
}
func ResourceIBMCISMtlsValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cis_id",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			CloudDataType:              "resource_instance",
			CloudDataRange:             []string{"service:internet-svcs"},
			Required:                   true})
	ibmCISMtlsValidator := validate.ResourceValidator{
		ResourceName: "ibm_cis_mtls",
		Schema:       validateSchema}
	return &ibmCISMtlsValidator
}

// cisMtlsCertFingerprintOf renders the first certificate of a PEM bundle the way the API reports its fingerprint.
func cisMtlsCertFingerprintOf(certPEM string) string {
	block, _ := pem.Decode([]byte(certPEM))
	if block == nil {
		return ""
	}
	sum := md5.Sum(block.Bytes)
	hexBytes := make([]string, len(sum))
	for i, b := range sum {
		hexBytes[i] = fmt.Sprintf("%02X", b)
	}
	return "MD5 Fingerprint=" + strings.Join(hexBytes, ":")
}

type cisMtlsZoneHosts struct {
	zoneID    string
	hostnames []string
}

func resourceIBMCISMtlsCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).CisMtlsSession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err,
			fmt.Sprintf("resourceIBMCISMtlsCreate CisMtlsSession initialization failed: %s", err.Error()),
			"ibm_cis_mtls", "create")
		return tfErr.GetDiag()
	}
	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)

	zones := expandCisMtlsZones(d.Get(cisMtlsZone).([]interface{}))
	zoneID := d.Get(cisDomainID).(string)
	hostnames := flex.ExpandStringList(d.Get(cisMtlsHostNames).([]interface{}))
	if len(zones) > 0 {
		// A hostname can only be added through its own zone: create through the first zone, then add the others.
		zoneID = zones[0].zoneID
		hostnames = zones[0].hostnames
	}

	options := sess.NewCreateAccessCertificateOptions(zoneID)
	options.SetName(d.Get(cisMtlsCertName).(string))
	options.SetCertificate(d.Get(cisMtlsCert).(string))
	options.SetAssociatedHostnames(hostnames)

	result, resp, err := sess.CreateAccessCertificate(options)
	if err != nil || result == nil {
		tfErr := flex.TerraformErrorf(err,
			fmt.Sprintf("resourceIBMCISMtlsCreate CreateAccessCertificate failed: %v \nResponse: %v", err, resp),
			"ibm_cis_mtls", "create")
		return tfErr.GetDiag()
	}

	d.SetId(flex.ConvertCisToTfThreeVar(*result.Result.ID, zoneID, crn))

	if len(zones) > 1 {
		owners, err := cisMtlsHostOwners(meta, crn, zones)
		if err != nil {
			tfErr := flex.TerraformErrorf(err,
				fmt.Sprintf("resourceIBMCISMtlsCreate GetZone failed: %s", err.Error()),
				"ibm_cis_mtls", "create")
			return tfErr.GetDiag()
		}
		calls := planCisMtlsZoneUpdates(result.Result.AssociatedHostnames, zones, owners)
		if err := applyCisMtlsZoneUpdates(sess, *result.Result.ID, d.Get(cisMtlsCertName).(string), calls); err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_cis_mtls", "create")
			return tfErr.GetDiag()
		}
	}
	return resourceIBMCISMtlsRead(context, d, meta)

}

func resourceIBMCISMtlsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).CisMtlsSession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err,
			fmt.Sprintf("resourceIBMCISMtlsRead CisMtlsSession initialization failed: %s", err.Error()),
			"ibm_cis_mtls", "read")
		return tfErr.GetDiag()
	}

	certID, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(d.Id())
	sess.Crn = core.StringPtr(crn)
	getOptions := sess.NewGetAccessCertificateOptions(zoneID, certID)
	result, response, err := sess.GetAccessCertificate(getOptions)

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err,
			fmt.Sprintf("resourceIBMCISMtlsRead GetAccessCertificate failed: %s \nResponse: %v \nResult: %v", err.Error(), response, result),
			"ibm_cis_mtls", "read")
		return tfErr.GetDiag()
	}

	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)
	d.Set(cisMtlsID, *result.Result.ID)
	d.Set(cisMtlsCertName, result.Result.Name)
	d.Set(cisMtlsCertFingerprint, result.Result.Fingerprint)
	d.Set(cisMtlsCertCreatedAt, *result.Result.CreatedAt)
	d.Set(cisMtlsCertUpdatedAt, *result.Result.UpdatedAt)
	d.Set(cisMtlsCertExpireOn, *result.Result.ExpiresOn)
	d.Set(cisMtlsCertID, *result.Result.ID)

	zones := expandCisMtlsZones(d.Get(cisMtlsZone).([]interface{}))
	if len(zones) == 0 {
		prior := flex.ExpandStringList(d.Get(cisMtlsHostNames).([]interface{}))
		d.Set(cisMtlsHostNames, orderCisMtlsHostnames(result.Result.AssociatedHostnames, prior))
		return nil
	}

	// Hosts outside the configured zones are left out of state and are never changed by this resource.
	owners, err := cisMtlsHostOwners(meta, crn, zones)
	if err != nil {
		tfErr := flex.TerraformErrorf(err,
			fmt.Sprintf("resourceIBMCISMtlsRead GetZone failed: %s", err.Error()),
			"ibm_cis_mtls", "read")
		return tfErr.GetDiag()
	}
	zoneList := make([]interface{}, 0, len(zones))
	for _, z := range zones {
		hosts := []string{}
		for _, h := range result.Result.AssociatedHostnames {
			if owners(h) == z.zoneID {
				hosts = append(hosts, h)
			}
		}
		zoneList = append(zoneList, map[string]interface{}{
			cisDomainID:          z.zoneID,
			cisMtlsZoneHostNames: flex.NewStringSet(schema.HashString, hosts),
		})
	}
	d.Set(cisMtlsZone, zoneList)
	d.Set(cisMtlsHostNames, nil)

	return nil
}

func resourceIBMCISMtlsUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).CisMtlsSession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err,
			fmt.Sprintf("resourceIBMCISMtlsUpdate CisMtlsSession initialization failed: %s", err.Error()),
			"ibm_cis_mtls", "update")
		return tfErr.GetDiag()
	}

	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)

	certID, zoneID, _, _ := flex.ConvertTfToCisThreeVar(d.Id())
	name := d.Get(cisMtlsCertName).(string)

	zones := expandCisMtlsZones(d.Get(cisMtlsZone).([]interface{}))
	if len(zones) > 0 {
		if !d.HasChange(cisMtlsCertName) && !d.HasChange(cisMtlsZone) {
			return resourceIBMCISMtlsRead(context, d, meta)
		}
		oldZones, _ := d.GetChange(cisMtlsZone)
		removed := []cisMtlsZoneHosts{}
		for _, oz := range expandCisMtlsZones(oldZones.([]interface{})) {
			if !cisMtlsZonesContain(zones, oz.zoneID) {
				removed = append(removed, cisMtlsZoneHosts{zoneID: oz.zoneID})
			}
		}
		owners, err := cisMtlsHostOwners(meta, crn, append(zones, removed...))
		if err != nil {
			tfErr := flex.TerraformErrorf(err,
				fmt.Sprintf("resourceIBMCISMtlsUpdate GetZone failed: %s", err.Error()),
				"ibm_cis_mtls", "update")
			return tfErr.GetDiag()
		}
		current, resp, err := sess.GetAccessCertificate(sess.NewGetAccessCertificateOptions(zoneID, certID))
		if err != nil {
			tfErr := flex.TerraformErrorf(err,
				fmt.Sprintf("resourceIBMCISMtlsUpdate GetAccessCertificate failed: %s \nResponse: %v", err.Error(), resp),
				"ibm_cis_mtls", "update")
			return tfErr.GetDiag()
		}
		calls := planCisMtlsZoneUpdates(current.Result.AssociatedHostnames, append(zones, removed...), owners)
		if len(calls) == 0 && d.HasChange(cisMtlsCertName) {
			calls = []cisMtlsZoneHosts{{zoneID: zoneID, hostnames: current.Result.AssociatedHostnames}}
		}
		if err := applyCisMtlsZoneUpdates(sess, certID, name, calls); err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_cis_mtls", "update")
			return tfErr.GetDiag()
		}
		if !cisMtlsZonesContain(zones, zoneID) {
			d.SetId(flex.ConvertCisToTfThreeVar(certID, zones[0].zoneID, crn))
		}
		return resourceIBMCISMtlsRead(context, d, meta)
	}

	if d.HasChange(cisMtlsCertName) ||
		d.HasChange(cisMtlsHostNames) {

		updateOption := sess.NewUpdateAccessCertificateOptions(zoneID, certID)
		// Allowing hostnames to be set empty. Resource cannot be deleted until hostname is attached.
		updateOption.SetAssociatedHostnames(flex.ExpandStringList(d.Get(cisMtlsHostNames).([]interface{})))
		updateOption.SetName(name)

		_, updateResp, updateErr := sess.UpdateAccessCertificate(updateOption)
		if updateErr != nil {
			tfErr := flex.TerraformErrorf(updateErr,
				fmt.Sprintf("resourceIBMCISMtlsUpdate UpdateAccessCertificate failed: %s \nResponse: %v", updateErr.Error(), updateResp),
				"ibm_cis_mtls", "update")
			return tfErr.GetDiag()
		}
	}

	return resourceIBMCISMtlsRead(context, d, meta)
}

func expandCisMtlsZones(in []interface{}) []cisMtlsZoneHosts {
	zones := make([]cisMtlsZoneHosts, 0, len(in))
	for _, raw := range in {
		m, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}
		zones = append(zones, cisMtlsZoneHosts{
			zoneID:    strings.Split(m[cisDomainID].(string), ":")[0],
			hostnames: flex.ExpandStringList(m[cisMtlsZoneHostNames].(*schema.Set).List()),
		})
	}
	return zones
}

func cisMtlsZonesContain(zones []cisMtlsZoneHosts, zoneID string) bool {
	for _, z := range zones {
		if z.zoneID == zoneID {
			return true
		}
	}
	return false
}

// cisMtlsHostOwners resolves the zone names and returns a lookup of the zone a hostname belongs to, or "" if none.
func cisMtlsHostOwners(meta interface{}, crn string, zones []cisMtlsZoneHosts) (func(string) string, error) {
	client, err := meta.(conns.ClientSession).CisZonesV1ClientSession()
	if err != nil {
		return nil, err
	}
	client.Crn = core.StringPtr(crn)
	names := map[string]string{}
	for _, z := range zones {
		if _, ok := names[z.zoneID]; ok {
			continue
		}
		result, resp, err := client.GetZone(client.NewGetZoneOptions(z.zoneID))
		if err != nil || result == nil || result.Result == nil || result.Result.Name == nil {
			return nil, fmt.Errorf("zone %s: %v \nResponse: %v", z.zoneID, err, resp)
		}
		names[z.zoneID] = *result.Result.Name
	}
	return cisMtlsHostOwnerLookup(names), nil
}

func cisMtlsHostOwnerLookup(zoneNames map[string]string) func(string) string {
	return func(host string) string {
		host = strings.ToLower(strings.TrimSuffix(host, "."))
		owner, best := "", 0
		for id, name := range zoneNames {
			name = strings.ToLower(strings.TrimSuffix(name, "."))
			if (host == name || strings.HasSuffix(host, "."+name)) && len(name) > best {
				owner, best = id, len(name)
			}
		}
		return owner
	}
}

// planCisMtlsZoneUpdates returns the update calls, one per zone in order, that turn current into the
// desired hostnames: each zone's own hostnames are added and removed through that zone, and hosts
// outside the given zones are carried along unchanged.
func planCisMtlsZoneUpdates(current []string, zones []cisMtlsZoneHosts, owner func(string) string) []cisMtlsZoneHosts {
	desired := map[string]bool{}
	for _, z := range zones {
		for _, h := range z.hostnames {
			desired[h] = true
		}
	}
	cumulative := append([]string{}, current...)
	calls := []cisMtlsZoneHosts{}
	for _, z := range zones {
		present := map[string]bool{}
		target := []string{}
		changed := false
		for _, h := range cumulative {
			if !desired[h] && owner(h) == z.zoneID {
				changed = true
				continue
			}
			present[h] = true
			target = append(target, h)
		}
		for _, h := range z.hostnames {
			if !present[h] {
				present[h] = true
				target = append(target, h)
				changed = true
			}
		}
		if changed {
			calls = append(calls, cisMtlsZoneHosts{zoneID: z.zoneID, hostnames: target})
			cumulative = target
		}
	}
	return calls
}

func applyCisMtlsZoneUpdates(sess *mtlsv1.MtlsV1, certID, name string, calls []cisMtlsZoneHosts) error {
	for _, c := range calls {
		opt := sess.NewUpdateAccessCertificateOptions(c.zoneID, certID)
		opt.SetName(name)
		// Sent even when empty, so the last hostname can be removed.
		opt.AssociatedHostnames = c.hostnames
		if opt.AssociatedHostnames == nil {
			opt.AssociatedHostnames = []string{}
		}
		_, resp, err := sess.UpdateAccessCertificate(opt)
		if err != nil {
			return fmt.Errorf("UpdateAccessCertificate through zone %s failed: %s \nResponse: %v", c.zoneID, err.Error(), resp)
		}
	}
	return nil
}

// orderCisMtlsHostnames keeps the prior order of hostnames that are still associated, then appends new ones.
func orderCisMtlsHostnames(actual, prior []string) []string {
	remaining := map[string]bool{}
	for _, h := range actual {
		remaining[h] = true
	}
	ordered := make([]string, 0, len(actual))
	for _, h := range prior {
		if remaining[h] {
			ordered = append(ordered, h)
			delete(remaining, h)
		}
	}
	for _, h := range actual {
		if remaining[h] {
			ordered = append(ordered, h)
			delete(remaining, h)
		}
	}
	return ordered
}

func resourceIBMCISMtlsDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).CisMtlsSession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err,
			fmt.Sprintf("resourceIBMCISMtlsDelete CisMtlsSession initialization failed: %s", err.Error()),
			"ibm_cis_mtls", "delete")
		return tfErr.GetDiag()
	}

	certID, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(d.Id())
	//crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)
	// zoneID := d.Get(cisDomainID).(string)

	// The API refuses to delete a certificate that still has associated hostnames.
	current, getResp, err := sess.GetAccessCertificate(sess.NewGetAccessCertificateOptions(zoneID, certID))
	if err != nil {
		if getResp != nil && getResp.StatusCode == 404 {
			return nil
		}
		tfErr := flex.TerraformErrorf(err,
			fmt.Sprintf("resourceIBMCISMtlsDelete GetAccessCertificate failed: %s \nResponse: %v", err.Error(), getResp),
			"ibm_cis_mtls", "delete")
		return tfErr.GetDiag()
	}
	if len(current.Result.AssociatedHostnames) > 0 {
		calls := []cisMtlsZoneHosts{{zoneID: zoneID}}
		if zones := expandCisMtlsZones(d.Get(cisMtlsZone).([]interface{})); len(zones) > 0 {
			for i := range zones {
				zones[i].hostnames = nil
			}
			owners, err := cisMtlsHostOwners(meta, crn, zones)
			if err != nil {
				tfErr := flex.TerraformErrorf(err,
					fmt.Sprintf("resourceIBMCISMtlsDelete GetZone failed: %s", err.Error()),
					"ibm_cis_mtls", "delete")
				return tfErr.GetDiag()
			}
			calls = planCisMtlsZoneUpdates(current.Result.AssociatedHostnames, zones, owners)
		}
		if err := applyCisMtlsZoneUpdates(sess, certID, *current.Result.Name, calls); err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_cis_mtls", "delete")
			return tfErr.GetDiag()
		}
	}

	delOpt := sess.NewDeleteAccessCertificateOptions(zoneID, certID)
	var delResp *core.DetailedResponse
	// Detached hostnames take a few seconds to clear; until then the delete returns 409.
	delErr := retry.RetryContext(context, 2*time.Minute, func() *retry.RetryError {
		var err error
		_, delResp, err = sess.DeleteAccessCertificate(delOpt)
		if err != nil && delResp != nil && delResp.StatusCode == 409 {
			return retry.RetryableError(err)
		}
		if err != nil {
			return retry.NonRetryableError(err)
		}
		return nil
	})
	if delErr != nil {
		tfErr := flex.TerraformErrorf(delErr,
			fmt.Sprintf("resourceIBMCISMtlsDelete DeleteAccessCertificate failed: %s \nResponse: %v", delErr.Error(), delResp),
			"ibm_cis_mtls", "delete")
		return tfErr.GetDiag()
	}

	return nil

}
