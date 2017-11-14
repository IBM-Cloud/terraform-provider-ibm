package ibm

import (
	"fmt"

	"github.com/IBM-Bluemix/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Bluemix/bluemix-go/bmxerror"
	"github.com/IBM-Bluemix/bluemix-go/helpers"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIBMOrg() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMOrgCreate,
		Read:     resourceIBMOrgRead,
		Delete:   resourceIBMOrgDelete,
		Update:   resourceIBMOrgUpdate,
		Exists:   resourceIBMOrgExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Org name, for example myorg@domain",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_quota_definition_guid": {
				Description: "Org quota guid",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},

			"billing_managers": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The IBMID of the users who will have billing manager role in this org, ex - user@example.com",
			},
			"managers": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The IBMID of the users who will have manager role in this org, ex - user@example.com",
			},
			"auditors": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The IBMID of the users who will have auditor role in this org, ex - user@example.com",
			},
			"users": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The IBMID of the users who will have user role in this org, ex - user@example.com",
			},

			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}
func resourceIBMOrgCreate(d *schema.ResourceData, meta interface{}) error {
	cfAPI, err := meta.(ClientSession).MccpAPI()

	if err != nil {
		return err
	}
	orgAPI := cfAPI.Organizations()
	orgName := d.Get("name").(string)

	req := mccpv2.OrgCreateRequest{
		Name: orgName,
	}

	if orgQuotaDefinitionGUID, ok := d.GetOk("org_quota_definition_guid"); ok {

		req.OrgQuotaDefinitionGUID = orgQuotaDefinitionGUID.(string)
	}

	orgFields, err := orgAPI.Create(req)
	if err != nil {
		return fmt.Errorf("Error creating organisation: %s", err)
	}

	orgGUID := orgFields.Metadata.GUID
	d.SetId(orgGUID)

	if billingManagerSet := d.Get("billing_managers").(*schema.Set); len(billingManagerSet.List()) > 0 {
		billingManagers := expandStringList(billingManagerSet.List())
		for _, d := range billingManagers {
			_, err := orgAPI.AssociateBillingManager(orgGUID, d)
			if err != nil {
				return fmt.Errorf("Error associating billing manager %s with org %s : %s", d, orgGUID, err)
			}
		}
	}

	if auditorSet := d.Get("auditors").(*schema.Set); len(auditorSet.List()) > 0 {
		auditors := expandStringList(auditorSet.List())
		for _, d := range auditors {
			_, err := orgAPI.AssociateAuditor(orgGUID, d)
			if err != nil {
				return fmt.Errorf("Error associating auditor %s with org %s : %s", d, orgGUID, err)
			}
		}

	}
	if managerSet := d.Get("managers").(*schema.Set); len(managerSet.List()) > 0 {
		managers := expandStringList(managerSet.List())
		for _, d := range managers {
			_, err := orgAPI.AssociateManager(orgGUID, d)
			if err != nil {
				return fmt.Errorf("Error associating manager %s with org %s : %s", d, orgGUID, err)
			}
		}
	}

	if userSet := d.Get("users").(*schema.Set); len(userSet.List()) > 0 {
		users := expandStringList(userSet.List())
		for _, d := range users {
			_, err := orgAPI.AssociateUser(orgGUID, d)
			if err != nil {
				return fmt.Errorf("Error associating users %s with org %s : %s", d, orgGUID, err)
			}
		}
	}

	return resourceIBMOrgRead(d, meta)
}
func resourceIBMOrgRead(d *schema.ResourceData, meta interface{}) error {
	cfAPI, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return err
	}
	orgAPI := cfAPI.Organizations()
	id := d.Id()
	if err != nil {
		return err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	orgOwnerID := userDetails.userEmail
	orgFields, err := orgAPI.Get(id)
	if err != nil {
		return fmt.Errorf("Error retrieving organisation: %s", err)
	}

	d.Set("name", orgFields.Entity.Name)

	billingManager, err := orgAPI.ListBillingManager(id)
	if err != nil {
		return fmt.Errorf("Error retrieving billing manager in the org: %s", err)
	}

	managers, err := orgAPI.ListManager(id)
	if err != nil {
		return fmt.Errorf("Error retrieving managers in the org: %s", err)
	}

	auditors, err := orgAPI.ListAuditors(id)
	if err != nil {
		return fmt.Errorf("Error retrieving auditors in space: %s", err)
	}

	users, err := orgAPI.ListUsers(id)
	if err != nil {
		return fmt.Errorf("Error retrieving users in space: %s", err)
	}

	d.Set("auditors", flattenOrgRole(auditors))
	d.Set("managers", flattenOrgRoleManagerUsers(managers, orgOwnerID))
	d.Set("billing_managers", flattenOrgRole(billingManager))
	d.Set("users", flattenOrgRoleManagerUsers(users, orgOwnerID))

	if orgFields.Entity.OrgQuotaDefinitionGUID != "" {
		d.Set("org_quota_definition_guid", orgFields.Entity.OrgQuotaDefinitionGUID)
	}

	d.Set("status", orgFields.Entity.Status)
	d.Set("billing_enabled", orgFields.Entity.BillingEnabled)
	return nil
}

func resourceIBMOrgUpdate(d *schema.ResourceData, meta interface{}) error {
	cfAPI, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return err
	}
	orgAPI := cfAPI.Organizations()
	id := d.Id()

	d.Partial(true)

	req := mccpv2.OrgUpdateRequest{}

	if d.HasChange("name") {
		req.Name = helpers.String(d.Get("name").(string))
	}

	_, err = orgAPI.Update(id, req)

	if err != nil {
		return fmt.Errorf("Error updating organisation: %s", err)
	}

	err = updateOrgBillingManagers(orgAPI, id, d)
	if err != nil {
		return err
	}
	err = updateOrgManagers(orgAPI, id, d)
	if err != nil {
		return err
	}
	err = updateOrgAuditors(orgAPI, id, d)
	if err != nil {
		return err
	}
	err = updateOrgUsers(orgAPI, id, d)
	if err != nil {
		return err
	}

	d.Partial(false)
	return resourceIBMOrgRead(d, meta)
}

func resourceIBMOrgDelete(d *schema.ResourceData, meta interface{}) error {
	cfAPI, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return err
	}
	orgAPI := cfAPI.Organizations()
	id := d.Id()
	err = orgAPI.Delete(id)
	if err != nil {
		return fmt.Errorf("Error deleting organisation: %s", err)
	}

	d.SetId("")

	return nil
}

func resourceIBMOrgExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	cfClient, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return false, err
	}
	id := d.Id()

	org, err := cfClient.Organizations().Get(id)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}

	return org.Metadata.GUID == id, nil
}

func updateOrgBillingManagers(api mccpv2.Organizations, orgGUID string, d *schema.ResourceData) error {
	if !d.HasChange("billing_managers") {
		return nil
	}
	var remove, add []string
	o, n := d.GetChange("billing_managers")
	os := o.(*schema.Set)
	ns := n.(*schema.Set)

	remove = expandStringList(os.Difference(ns).List())
	add = expandStringList(ns.Difference(os).List())

	if len(add) > 0 {
		for _, d := range add {
			_, err := api.AssociateBillingManager(orgGUID, d)
			if err != nil {
				return fmt.Errorf("Error associating billing manager %s with org %s : %s", d, orgGUID, err)
			}
		}
	}
	if len(remove) > 0 {
		for _, d := range remove {
			err := api.DisassociateBillingManager(orgGUID, d)
			if err != nil {
				return fmt.Errorf("Error dis-associating billing manager %s with org %s : %s", d, orgGUID, err)
			}
		}
	}
	d.SetPartial("billing_managers")
	return nil
}

func updateOrgManagers(api mccpv2.Organizations, orgGUID string, d *schema.ResourceData) error {
	if !d.HasChange("managers") {
		return nil
	}
	var remove, add []string
	o, n := d.GetChange("managers")
	os := o.(*schema.Set)
	ns := n.(*schema.Set)

	remove = expandStringList(os.Difference(ns).List())
	add = expandStringList(ns.Difference(os).List())

	if len(add) > 0 {
		for _, d := range add {
			_, err := api.AssociateManager(orgGUID, d)
			if err != nil {
				return fmt.Errorf("Error associating manager %s with org %s : %s", d, orgGUID, err)
			}
		}
	}
	if len(remove) > 0 {
		for _, d := range remove {
			err := api.DisassociateManager(orgGUID, d)
			if err != nil {
				return fmt.Errorf("Error dis-associating manager %s with org %s : %s", d, orgGUID, err)
			}
		}
	}
	d.SetPartial("managers")
	return nil
}
func updateOrgAuditors(api mccpv2.Organizations, orgGUID string, d *schema.ResourceData) error {
	if !d.HasChange("auditors") {
		return nil
	}
	var remove, add []string
	o, n := d.GetChange("auditors")
	os := o.(*schema.Set)
	ns := n.(*schema.Set)

	remove = expandStringList(os.Difference(ns).List())
	add = expandStringList(ns.Difference(os).List())

	if len(add) > 0 {
		for _, d := range add {
			_, err := api.AssociateAuditor(orgGUID, d)
			if err != nil {
				return fmt.Errorf("Error associating auditor %s with org %s : %s", d, orgGUID, err)
			}
		}
	}
	if len(remove) > 0 {
		for _, d := range remove {
			err := api.DisassociateAuditor(orgGUID, d)
			if err != nil {
				return fmt.Errorf("Error dis-associating auditor %s with org %s : %s", d, orgGUID, err)
			}
		}
	}
	d.SetPartial("auditors")
	return nil
}

func updateOrgUsers(api mccpv2.Organizations, orgGUID string, d *schema.ResourceData) error {
	if !d.HasChange("users") {
		return nil
	}
	var remove, add []string
	o, n := d.GetChange("users")
	os := o.(*schema.Set)
	ns := n.(*schema.Set)

	remove = expandStringList(os.Difference(ns).List())
	add = expandStringList(ns.Difference(os).List())

	if len(add) > 0 {
		for _, d := range add {
			_, err := api.AssociateUser(orgGUID, d)
			if err != nil {
				return fmt.Errorf("Error associating user %s with org %s : %s", d, orgGUID, err)
			}
		}
	}
	if len(remove) > 0 {
		for _, d := range remove {
			err := api.DisassociateUser(orgGUID, d)
			if err != nil {
				return fmt.Errorf("Error dis-associating user %s with org %s : %s", d, orgGUID, err)
			}
		}
	}
	d.SetPartial("users")
	return nil
}
