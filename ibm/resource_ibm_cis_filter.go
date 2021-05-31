// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/networking-go-sdk/filtersv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ibmCISFilters        = "ibm_cis_filter"
	cisFilterExpression  = "expression"
	cisFilterPaused      = "paused"
	cisFilterDescription = "description"
	cisFilterID          = "filterid"
)

type FilterCreateUpdate struct {
	Success bool `json:"success"`
	Errors  []struct {
	} `json:"errors"`
	Messages []struct {
	} `json:"messages"`
	Result []struct {
		ID          string `json:"id"`
		Paused      bool   `json:"paused"`
		Description string `json:"description"`
		Expression  string `json:"expression"`
		CreatedOn   struct {
		} `json:"created_on"`
		ModifiedOn struct {
		} `json:"modified_on"`
	} `json:"result"`
}
type FilterRead struct {
	Success bool `json:"success"`
	Errors  []struct {
	} `json:"errors"`
	Messages []struct {
	} `json:"messages"`
	Result struct {
		ID          string `json:"id"`
		Paused      bool   `json:"paused"`
		Description string `json:"description"`
		Expression  string `json:"expression"`
		CreatedOn   struct {
		} `json:"created_on"`
		ModifiedOn struct {
		} `json:"modified_on"`
	} `json:"result"`
}

func resourceIBMCISFilter() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMCISFilterCreate,
		Read:     resourceIBMCISFilterRead,
		Update:   resourceIBMCISFilterUpdate,
		Delete:   resourceIBMCISFilterDelete,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Description:      "Associated CIS domain",
				Required:         true,
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			cisFilterPaused: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Filter Paused",
			},
			cisFilterID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Filter ID",
			},
			cisFilterExpression: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Filter Expression",
			},
			cisFilterDescription: {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Filter Description",
				ValidateFunc: InvokeValidator(ibmCISFilters, cisFilterDescription),
			},
		},
	}
}
func resourceIBMCISFilterCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return fmt.Errorf("Error while Getting IAM Access Token using BluemixSession %s", err)
	}
	xAuthtoken := sess.Config.IAMAccessToken

	cisClient, err := meta.(ClientSession).CisFiltersSession()
	if err != nil {
		return fmt.Errorf("Error while getting the CisFiltersSession %s", err)
	}

	crn := d.Get(cisID).(string)
	zoneID, _, err := convertTftoCisTwoVar(d.Get(cisDomainID).(string))

	var newfilter filtersv1.FilterInput

	if _, ok := d.GetOk(cisFilterPaused); ok {
		paused := d.Get(cisFilterPaused).(bool)
		newfilter.Paused = &paused
	}
	if _, ok := d.GetOk(cisFilterDescription); ok {
		description := d.Get(cisFilterDescription).(string)
		newfilter.Description = &description
	}
	if _, ok := d.GetOk(cisFilterExpression); ok {
		expression := d.Get(cisFilterExpression).(string)
		newfilter.Expression = &expression
	}

	opt := cisClient.NewCreateFilterOptions(xAuthtoken, crn, zoneID)

	filetrInput := &filtersv1.FilterInput{
		Expression:  newfilter.Expression,
		Paused:      newfilter.Paused,
		Description: newfilter.Description,
	}

	opt.SetFilterInput([]filtersv1.FilterInput{*filetrInput})

	result, response, err := cisClient.CreateFilter(opt)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		return fmt.Errorf("Error creating Filter for zone %q: %s", zoneID, err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(result)
	newStr := buf.String()

	filter_res := FilterCreateUpdate{}
	json_err := json.Unmarshal([]byte(newStr), &filter_res)

	if json_err != nil {
		return fmt.Errorf("Error unmarshal the jason string %s", json_err)
	}

	d.SetId(convertCisToTfThreeVar(filter_res.Result[0].ID, zoneID, crn))
	return resourceIBMCISFilterRead(d, meta)

}
func resourceIBMCISFilterRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return fmt.Errorf("Error while Getting IAM Access Token using BluemixSession %s", err)
	}
	xAuthtoken := sess.Config.IAMAccessToken

	cisClient, err := meta.(ClientSession).CisFiltersSession()
	if err != nil {
		return fmt.Errorf("Error while getting the CisFiltersSession %s", err)
	}
	filterid, zoneID, crn, err := convertTfToCisThreeVar(d.Id())
	if err != nil {
		return err
	}

	opt := cisClient.NewGetFilterOptions(xAuthtoken, crn, zoneID, filterid)

	filters, response, err := cisClient.GetFilter(opt)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			log.Printf("Error GetFilter not found ")
			return nil
		}
		return fmt.Errorf("Error finding GetFilter %q: %s", d.Id(), err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(filters)
	newStr := buf.String()

	filter_res := FilterRead{}
	json_err := json.Unmarshal([]byte(newStr), &filter_res)
	if json_err != nil {
		return fmt.Errorf("Error unmarshal the json string resourceIBMCISFilterRead %s", json_err)
	}
	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)
	d.Set(cisFilterID, filter_res.Result.ID)
	d.Set(cisFilterPaused, filter_res.Result.Paused)
	d.Set(cisFilterDescription, filter_res.Result.Description)
	d.Set(cisFilterExpression, filter_res.Result.Expression)

	return nil
}
func resourceIBMCISFilterUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return fmt.Errorf("Error while Getting IAM Access Token using BluemixSession %s", err)
	}
	xAuthtoken := sess.Config.IAMAccessToken

	cisClient, err := meta.(ClientSession).CisFiltersSession()
	if err != nil {
		return fmt.Errorf("Error while getting the CisFiltersSession %s", err)
	}

	filterid, zoneID, crn, err := convertTfToCisThreeVar(d.Id())
	if err != nil {
		return err
	}
	var updatefilter filtersv1.FilterUpdateInput
	updatefilter.ID = &filterid

	if _, ok := d.GetOk(cisFilterPaused); ok {
		paused := d.Get(cisFilterPaused).(bool)
		updatefilter.Paused = &paused
	}
	if _, ok := d.GetOk(cisFilterDescription); ok {
		description := d.Get(cisFilterDescription).(string)
		updatefilter.Description = &description
	}
	if _, ok := d.GetOk(cisFilterExpression); ok {
		expression := d.Get(cisFilterExpression).(string)
		updatefilter.Expression = &expression
	}

	opt := cisClient.NewUpdateFiltersOptions(xAuthtoken, crn, zoneID)

	filterUpdateInput := &filtersv1.FilterUpdateInput{
		ID:          updatefilter.ID,
		Expression:  updatefilter.Expression,
		Paused:      updatefilter.Paused,
		Description: updatefilter.Description,
	}
	opt.SetFilterUpdateInput([]filtersv1.FilterUpdateInput{*filterUpdateInput})

	result, _, err := cisClient.UpdateFilters(opt)
	if err != nil {
		return fmt.Errorf("Error updating Filter for zone %q: %s", zoneID, err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(result)
	newStr := buf.String()

	filter_res := FilterCreateUpdate{}
	json_err := json.Unmarshal([]byte(newStr), &filter_res)
	if json_err != nil {
		return fmt.Errorf("Error unmarshal the json string resourceIBMCISFilterUpdate %s", json_err)
	}
	if filter_res.Result[0].ID == "" {
		return fmt.Errorf("Error failed to find id in Update response; resource was empty")
	}
	return resourceIBMCISFilterRead(d, meta)
}
func resourceIBMCISFilterDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return err
	}
	xAuthtoken := sess.Config.IAMAccessToken
	cisClient, err := meta.(ClientSession).CisFiltersSession()
	if err != nil {
		return err
	}
	filterid, zoneID, crn, err := convertTfToCisThreeVar(d.Id())
	if err != nil {
		return err
	}
	opt := cisClient.NewDeleteFiltersOptions(xAuthtoken, crn, zoneID, filterid)
	_, _, err = cisClient.DeleteFilters(opt)
	if err != nil {
		return fmt.Errorf("Error deleting Filter: %s", err)
	}

	return nil
}

func resourceIBMCISFilterValidator() *ResourceValidator {
	validateSchema := make([]ValidateSchema, 1)
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 cisFilterDescription,
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Required:                   true,
			AllowedValues:              "Filter-creation"})

	ibmCISFiltersResourceValidator := ResourceValidator{ResourceName: ibmCISFilters, Schema: validateSchema}
	return &ibmCISFiltersResourceValidator
}
