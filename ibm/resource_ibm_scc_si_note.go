package ibm

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/ibm-cloud-security/security-advisor-sdk-go/findingsapiv1"
)

func resourceIBMNote() *schema.Resource {
	return &schema.Resource{
		Create: resourceIBMNoteCreate,
		Read:   resourceIBMNoteRead,
		Update: resourceIBMNoteUpdate,
		Delete: resourceIBMNoteDelete,
		Exists: resourceIBMNoteExists,

		Schema: map[string]*schema.Schema{
			"provider_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"note_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"short_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"long_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"kind": {
				Type:     schema.TypeString,
				Required: true,
			},
			"related_url": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"label": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"url": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"expiration_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"update_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"shared": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reported_by": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"title": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"url": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"finding": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"severity": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"next_steps": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"title": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"kpi": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aggregation_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"card": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"section": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"title": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"subtitle": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"order": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"finding_note_names": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"requires_configuration": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"badge_text": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"badge_image": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"elements": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"kind": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"text": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"value_type": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"kind": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"finding_note_names": {
													Type:     schema.TypeList,
													Optional: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
												},
												"kpi_note_name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"text": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"value_types": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"kind": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"finding_note_names": {
													Type:     schema.TypeList,
													Optional: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
												},
												"kpi_note_name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"text": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"default_time_range": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"default_interval": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"section": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"image": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"title": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceIBMNoteCreate(d *schema.ResourceData, meta interface{}) error {
	shortDescription := d.Get("short_description").(string)
	longDescription := d.Get("long_description").(string)
	kind := d.Get("kind").(string)
	relatedURLInterface := d.Get("related_url").([]interface{})
	expirationTime := d.Get("expiration_time").(string)
	createTime := d.Get("create_time").(string)
	updateTime := d.Get("update_time").(string)
	noteID := d.Get("note_id").(string)
	shared, err := strconv.ParseBool(d.Get("shared").(string))
	if err != nil {
		return fmt.Errorf("error occurred while creating note: %v", err)
	}
	reportedByInterface := d.Get("reported_by").([]interface{})
	findingInterface := d.Get("finding").([]interface{})
	kpiInterface := d.Get("kpi").([]interface{})
	cardInterface := d.Get("card").([]interface{})
	sectionInterface := d.Get("section").([]interface{})

	createNoteOptions := findingsapiv1.CreateNoteOptions{}

	if len(relatedURLInterface) != 0 {
		relatedURLObjectList := make([]findingsapiv1.ApiNoteRelatedURL, 0)
		for _, relatedURL := range relatedURLInterface {
			relatedURLObject := findingsapiv1.ApiNoteRelatedURL{}
			label := relatedURL.(map[string]interface{})["label"].(string)
			url := relatedURL.(map[string]interface{})["url"].(string)
			relatedURLObject.Label = &label
			relatedURLObject.URL = &url
			relatedURLObjectList = append(relatedURLObjectList, relatedURLObject)
		}
		createNoteOptions.RelatedURL = relatedURLObjectList
	}

	if len(reportedByInterface) != 0 {
		reportedByObject := findingsapiv1.Reporter{}
		reporterID := reportedByInterface[0].(map[string]interface{})["id"].(string)
		reportedTitle := reportedByInterface[0].(map[string]interface{})["title"].(string)
		reporterURL := reportedByInterface[0].(map[string]interface{})["url"].(string)
		reportedByObject.ID = &reporterID
		reportedByObject.Title = &reportedTitle
		reportedByObject.URL = &reporterURL
		createNoteOptions.ReportedBy = &reportedByObject
	}

	if len(findingInterface) != 0 {
		findingObject := findingsapiv1.FindingType{}
		findingSeverity := findingInterface[0].(map[string]interface{})["severity"].(string)
		findingObject.Severity = &findingSeverity
		findingNextSteps := make([]findingsapiv1.RemediationStep, 0)
		for _, nextStep := range findingInterface[0].(map[string]interface{})["next_steps"].([]interface{}) {
			nextStepObject := findingsapiv1.RemediationStep{}
			remStepURL := nextStep.(map[string]interface{})["url"].(string)
			remStepTitle := nextStep.(map[string]interface{})["title"].(string)
			nextStepObject.URL = &remStepURL
			nextStepObject.Title = &remStepTitle
			findingNextSteps = append(findingNextSteps, nextStepObject)
		}
		findingObject.NextSteps = findingNextSteps
		createNoteOptions.Finding = &findingObject
	}

	if len(kpiInterface) != 0 {
		kpiObject := findingsapiv1.KpiType{}
		kpiAggregationType := kpiInterface[0].(map[string]interface{})["aggregation_type"].(string)
		kpiObject.AggregationType = &kpiAggregationType
		createNoteOptions.Kpi = &kpiObject
	}

	if len(cardInterface) != 0 {
		cardObject := findingsapiv1.Card{}
		cardSection := cardInterface[0].(map[string]interface{})["section"].(string)
		cardTitle := cardInterface[0].(map[string]interface{})["title"].(string)
		cardSubtitle := cardInterface[0].(map[string]interface{})["subtitle"].(string)
		cardOrder := int64(cardInterface[0].(map[string]interface{})["order"].(int))
		cardRequiresConf, _ := strconv.ParseBool(cardInterface[0].(map[string]interface{})["requires_configuration"].(string))
		cardBadgeText := cardInterface[0].(map[string]interface{})["badge_text"].(string)
		cardBadgeImage := cardInterface[0].(map[string]interface{})["badge_image"].(string)
		cardObject.Section = &cardSection
		cardObject.Title = &cardTitle
		cardObject.Subtitle = &cardSubtitle
		cardObject.Order = &cardOrder
		cardObject.RequiresConfiguration = &cardRequiresConf
		cardObject.BadgeText = &cardBadgeText
		cardObject.BadgeImage = &cardBadgeImage
		cardObjectFindingNoteNamesList := make([]string, 0)
		for _, findingNoteName := range cardInterface[0].(map[string]interface{})["finding_note_names"].([]interface{}) {
			cardObjectFindingNoteNamesList = append(cardObjectFindingNoteNamesList, findingNoteName.(string))
		}
		cardObject.FindingNoteNames = cardObjectFindingNoteNamesList
		cardElementsList := make([]findingsapiv1.CardElement, 0)
		for _, cardElement := range cardInterface[0].(map[string]interface{})["elements"].([]interface{}) {
			cardElementObject := findingsapiv1.CardElement{}
			cardElementKind := cardElement.(map[string]interface{})["kind"].(string)
			cardElementText := cardElement.(map[string]interface{})["text"].(string)
			cardElementDefTime := cardElement.(map[string]interface{})["default_time_range"].(string)
			cardElementDefInterval := cardElement.(map[string]interface{})["default_interval"].(string)
			cardElementObject.Kind = &cardElementKind
			cardElementObject.Text = &cardElementText
			cardElementObject.DefaultTimeRange = &cardElementDefTime
			cardElementObject.DefaultInterval = &cardElementDefInterval

			cardValueTypeObject := findingsapiv1.CardValueType{}
			cardValueTypeInterface := cardElement.(map[string]interface{})["value_type"].([]interface{})
			cardValueTypeKind := cardValueTypeInterface[0].(map[string]interface{})["kind"].(string)
			cardValueTypeKpiNoteName := cardValueTypeInterface[0].(map[string]interface{})["kpi_note_name"].(string)
			cardValueTypeText := cardValueTypeInterface[0].(map[string]interface{})["text"].(string)
			cardValueTypeObject.Kind = &cardValueTypeKind
			cardValueTypeObject.KpiNoteName = &cardValueTypeKpiNoteName
			cardValueTypeObject.Text = &cardValueTypeText
			cardValueTypeFindingNoteNames := make([]string, 0)
			for _, cardValueTypeFindingNoteName := range cardValueTypeInterface[0].(map[string]interface{})["finding_note_names"].([]interface{}) {
				cardValueTypeFindingNoteNames = append(cardValueTypeFindingNoteNames, cardValueTypeFindingNoteName.(string))
			}
			cardValueTypeObject.FindingNoteNames = cardValueTypeFindingNoteNames
			cardElementObject.ValueType = &cardValueTypeObject

			cardValueTypesObjectList := make([]findingsapiv1.CardValueType, 0)
			for _, cardValueType := range cardElement.(map[string]interface{})["value_types"].([]interface{}) {
				cardValueTypeObject := findingsapiv1.CardValueType{}
				cardValueTypeKind := cardValueType.(map[string]interface{})["kind"].(string)
				cardValueTypeKpiNoteName := cardValueType.(map[string]interface{})["kpi_note_name"].(string)
				cardValueTypeText := cardValueType.(map[string]interface{})["text"].(string)
				cardValueTypeObject.Kind = &cardValueTypeKind
				cardValueTypeObject.KpiNoteName = &cardValueTypeKpiNoteName
				cardValueTypeObject.Text = &cardValueTypeText
				cardValueTypeFindingNoteNames := make([]string, 0)
				for _, cardValueTypeFindingNoteName := range cardValueType.(map[string]interface{})["finding_note_names"].([]interface{}) {
					cardValueTypeFindingNoteNames = append(cardValueTypeFindingNoteNames, cardValueTypeFindingNoteName.(string))
				}
				cardValueTypeObject.FindingNoteNames = cardValueTypeFindingNoteNames
				cardValueTypesObjectList = append(cardValueTypesObjectList, cardValueTypeObject)
			}
			cardElementsList = append(cardElementsList, cardElementObject)
			cardElementObject.ValueTypes = cardValueTypesObjectList
		}
		cardObject.Elements = cardElementsList
		createNoteOptions.Card = &cardObject
	}

	if len(sectionInterface) != 0 {
		sectionObject := findingsapiv1.Section{}
		sectionTitle := sectionInterface[0].(map[string]interface{})["title"].(string)
		sectionImage := sectionInterface[0].(map[string]interface{})["image"].(string)
		sectionObject.Title = &sectionTitle
		sectionObject.Image = &sectionImage
		createNoteOptions.Section = &sectionObject
	}

	if shortDescription != "" {
		createNoteOptions.ShortDescription = &shortDescription
	}

	if longDescription != "" {
		createNoteOptions.LongDescription = &longDescription
	}

	createNoteOptions.Kind = &kind
	currentTime := strfmt.DateTime(time.Now())
	createNoteOptions.ExpirationTime = &currentTime
	if expirationTime != "" {
		noteExpirationTime, err := strfmt.ParseDateTime(expirationTime)
		if err != nil {
			return fmt.Errorf("error occurred while creating note: %v", err)
		}
		createNoteOptions.ExpirationTime = &noteExpirationTime
	}
	if createTime != "" {
		noteCreateTime, err := strfmt.ParseDateTime(createTime)
		if err != nil {
			return fmt.Errorf("error occurred while creating note: %v", err)
		}
		createNoteOptions.CreateTime = &noteCreateTime
	}
	if updateTime != "" {
		noteUpdateTime, err := strfmt.ParseDateTime(updateTime)
		if err != nil {
			return fmt.Errorf("error occurred while creating note: %v", err)
		}
		createNoteOptions.UpdateTime = &noteUpdateTime
	}
	createNoteOptions.ID = &noteID

	if &shared != nil {
		createNoteOptions.Shared = &shared
	}

	sess, err := meta.(ClientSession).FindingsV1API()
	if err != nil {
		return err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	accountID := userDetails.userAccount
	providerID := d.Get("provider_id").(string)

	createNoteOptions.AccountID = &accountID
	createNoteOptions.ProviderID = &providerID

	note, _, err := sess.CreateNote(&createNoteOptions)

	if err != nil {
		return fmt.Errorf("error occurred while creating note: %v", err)
	}
	if note.ID != nil {
		d.SetId(*note.ID)
		return resourceIBMNoteRead(d, meta)
	}

	return nil
}

func resourceIBMNoteRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).FindingsV1API()
	if err != nil {
		return err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	accountID := userDetails.userAccount
	providerID := d.Get("provider_id").(string)
	noteID := d.Id()

	getNoteOptions := sess.NewGetNoteOptions(accountID, providerID, noteID)
	note, _, err := sess.GetNote(getNoteOptions)

	if err != nil {
		return fmt.Errorf("error occurred while reading note: %v", err)
	}

	if note.ID == nil {
		return fmt.Errorf("no such note found: %v", noteID)
	}

	noteExpirationTime := ""
	noteShared := "false"
	noteRelatedURL := make([]map[string]interface{}, 0)
	noteReportedBy := make([]map[string]interface{}, 0)
	noteFinding := make([]map[string]interface{}, 0)
	noteKpi := make([]map[string]interface{}, 0)
	noteCard := make([]map[string]interface{}, 0)
	noteSection := make([]map[string]interface{}, 0)

	if note.ExpirationTime != nil {
		noteExpirationTime = note.ExpirationTime.String()
	}

	if note.Shared != nil {
		noteShared = "true"
	}

	noteRelatedURLList := make([]map[string]interface{}, 0)
	for _, noteRelatedURL := range note.RelatedURL {
		noteRelatedURLObject := map[string]interface{}{}
		noteRelatedURLObject["label"] = noteRelatedURL.Label
		noteRelatedURLObject["url"] = noteRelatedURL.URL
		noteRelatedURLList = append(noteRelatedURLList, noteRelatedURLObject)
	}
	noteRelatedURL = noteRelatedURLList

	noteReportedByObjectList := make([]map[string]interface{}, 1)
	noteReportedByObject := map[string]interface{}{}
	noteReportedByObject["id"] = note.ReportedBy.ID
	noteReportedByObject["title"] = note.ReportedBy.Title
	noteReportedByObject["url"] = note.ReportedBy.URL
	noteReportedByObjectList[0] = noteReportedByObject
	noteReportedBy = noteReportedByObjectList

	if note.Finding != nil {
		noteFindingTypeObjectList := make([]map[string]interface{}, 1)
		noteFindingTypeObject := map[string]interface{}{}
		noteFindingTypeObject["severity"] = note.Finding.Severity
		noteFindingTypeNextStepsList := make([]map[string]interface{}, 0)
		for _, noteFindingNextStep := range note.Finding.NextSteps {
			noteFindingTypeNextStepObject := map[string]interface{}{}
			noteFindingTypeNextStepObject["title"] = noteFindingNextStep.Title
			noteFindingTypeNextStepObject["url"] = noteFindingNextStep.URL
			noteFindingTypeNextStepsList = append(noteFindingTypeNextStepsList, noteFindingTypeNextStepObject)
		}
		noteFindingTypeObject["next_steps"] = noteFindingTypeNextStepsList
		noteFindingTypeObjectList = append(noteFindingTypeObjectList, noteFindingTypeObject)
		noteFinding = noteFindingTypeObjectList
	}

	if note.Kpi != nil {
		noteKpiTypeObjectList := make([]map[string]interface{}, 1)
		noteKpiTypeObject := map[string]interface{}{}
		noteKpiTypeObject["aggregation_type"] = note.Kpi.AggregationType
		noteKpiTypeObjectList[0] = noteKpiTypeObject
		noteKpi = noteKpiTypeObjectList
	}

	if note.Card != nil {
		noteCardObjectList := make([]map[string]interface{}, 1)
		noteCardObject := map[string]interface{}{}
		noteCardObject["section"] = note.Card.Section
		noteCardObject["title"] = note.Card.Title
		noteCardObject["subtitle"] = note.Card.Subtitle
		noteCardObject["order"] = note.Card.Order
		noteCardObject["finding_note_names"] = note.Card.FindingNoteNames
		noteCardObject["requires_configuration"] = "false"
		if note.Card.RequiresConfiguration != nil && *note.Card.RequiresConfiguration {
			noteCardObject["requires_configuration"] = "true"
		}
		noteCardObject["badge_text"] = note.Card.BadgeText
		noteCardObject["badge_image"] = note.Card.BadgeImage
		noteCardElementsList := make([]map[string]interface{}, 0)
		for _, noteCardElement := range note.Card.Elements {
			noteCardElementObject := map[string]interface{}{}
			noteCardElementObject["kind"] = noteCardElement.Kind
			noteCardElementObject["text"] = noteCardElement.Text
			noteCardElementObject["default_time_range"] = noteCardElement.DefaultTimeRange
			noteCardElementObject["default_interval"] = noteCardElement.DefaultInterval
			if noteCardElement.ValueType != nil {
				noteCardElementValueTypeObjectList := make([]map[string]interface{}, 1)
				noteCardElementValueTypeObject := map[string]interface{}{}
				noteCardElementValueTypeObject["kind"] = noteCardElement.ValueType.Kind
				noteCardElementValueTypeObject["finding_note_names"] = noteCardElement.ValueType.FindingNoteNames
				noteCardElementValueTypeObject["kpi_note_name"] = noteCardElement.ValueType.KpiNoteName
				noteCardElementValueTypeObject["text"] = noteCardElement.ValueType.Text
				noteCardElementValueTypeObjectList[0] = noteCardElementValueTypeObject
				noteCardElementObject["value_type"] = noteCardElementValueTypeObjectList
			}
			if noteCardElement.ValueTypes != nil {
				noteCardElementValueTypesList := make([]map[string]interface{}, 0)
				for _, noteCardElementValueTypes := range noteCardElement.ValueTypes {
					noteCardElementValueTypesObject := map[string]interface{}{}
					noteCardElementValueTypesObject["kind"] = noteCardElementValueTypes.Kind
					noteCardElementValueTypesObject["finding_note_names"] = noteCardElementValueTypes.FindingNoteNames
					noteCardElementValueTypesObject["kpi_note_name"] = noteCardElementValueTypes.KpiNoteName
					noteCardElementValueTypesObject["text"] = noteCardElementValueTypes.Text
				}
				noteCardElementObject["value_types"] = noteCardElementValueTypesList
				noteCardElementsList = append(noteCardElementsList, noteCardElementObject)
			}
		}
		noteCardObject["elements"] = noteCardElementsList
		noteCardObjectList[0] = noteCardObject
		noteCard = noteCardObjectList
	}

	if note.Section != nil {
		noteSectionObjectList := make([]map[string]interface{}, 1)
		noteSectionObject := map[string]interface{}{}
		noteSectionObject["image"] = note.Section.Image
		noteSectionObject["title"] = note.Section.Title
		noteSectionObjectList[0] = noteSectionObject
		noteSection = noteSectionObjectList
	}

	d.Set("short_description", note.ShortDescription)
	d.Set("long_description", note.LongDescription)
	d.Set("kind", note.Kind)
	d.Set("expiration_time", noteExpirationTime)
	d.Set("create_time", note.CreateTime)
	d.Set("update_time", note.UpdateTime)
	d.Set("shared", noteShared)
	d.Set("related_url", noteRelatedURL)
	d.Set("reported_by", noteReportedBy)
	d.Set("finding", noteFinding)
	d.Set("kpi", noteKpi)
	d.Set("card", noteCard)
	d.Set("section", noteSection)
	d.SetId(*note.ID)

	return nil
}

func resourceIBMNoteExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).FindingsV1API()
	if err != nil {
		return false, err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return false, err
	}
	accountID := userDetails.userAccount
	providerID := d.Get("provider_id").(string)
	noteID := d.Id()

	getNoteOptions := sess.NewGetNoteOptions(accountID, providerID, noteID)
	_, resp, err := sess.GetNote(getNoteOptions)

	if err != nil {
		return false, fmt.Errorf("error occurred while reading note: %v", err)
	}

	if resp.StatusCode == 404 {
		return false, nil
	}

	return true, nil
}

func resourceIBMNoteUpdate(d *schema.ResourceData, meta interface{}) error {
	shortDescription := d.Get("short_description").(string)
	longDescription := d.Get("long_description").(string)
	kind := d.Get("kind").(string)
	relatedURLInterface := d.Get("related_url").([]interface{})
	expirationTime := d.Get("expiration_time").(string)
	createTime := d.Get("create_time").(string)
	updateTime := d.Get("update_time").(string)
	shared, err := strconv.ParseBool(d.Get("shared").(string))
	if err != nil {
		return fmt.Errorf("error occurred while creating note: %v", err)
	}
	reportedByInterface := d.Get("reported_by").([]interface{})
	findingInterface := d.Get("finding").([]interface{})
	kpiInterface := d.Get("kpi").([]interface{})
	cardInterface := d.Get("card").([]interface{})
	sectionInterface := d.Get("section").([]interface{})

	updateNoteOptions := findingsapiv1.UpdateNoteOptions{}

	if len(relatedURLInterface) != 0 {
		relatedURLObjectList := make([]findingsapiv1.ApiNoteRelatedURL, 0)
		for _, relatedURL := range relatedURLInterface {
			relatedURLObject := findingsapiv1.ApiNoteRelatedURL{}
			label := relatedURL.(map[string]interface{})["label"].(string)
			url := relatedURL.(map[string]interface{})["url"].(string)
			relatedURLObject.Label = &label
			relatedURLObject.URL = &url
			relatedURLObjectList = append(relatedURLObjectList, relatedURLObject)
		}
		updateNoteOptions.RelatedURL = relatedURLObjectList
	}

	if len(reportedByInterface) != 0 {
		reportedByObject := findingsapiv1.Reporter{}
		reporterID := reportedByInterface[0].(map[string]interface{})["id"].(string)
		reportedTitle := reportedByInterface[0].(map[string]interface{})["title"].(string)
		reporterURL := reportedByInterface[0].(map[string]interface{})["url"].(string)
		reportedByObject.ID = &reporterID
		reportedByObject.Title = &reportedTitle
		reportedByObject.URL = &reporterURL
		updateNoteOptions.ReportedBy = &reportedByObject
	}

	if len(findingInterface) != 0 {
		findingObject := findingsapiv1.FindingType{}
		findingSeverity := findingInterface[0].(map[string]interface{})["severity"].(string)
		findingObject.Severity = &findingSeverity
		findingNextSteps := make([]findingsapiv1.RemediationStep, 0)
		for _, nextStep := range findingInterface[0].(map[string]interface{})["next_steps"].([]interface{}) {
			nextStepObject := findingsapiv1.RemediationStep{}
			remStepURL := nextStep.(map[string]interface{})["url"].(string)
			remStepTitle := nextStep.(map[string]interface{})["title"].(string)
			nextStepObject.URL = &remStepURL
			nextStepObject.Title = &remStepTitle
			findingNextSteps = append(findingNextSteps, nextStepObject)
		}
		findingObject.NextSteps = findingNextSteps
		updateNoteOptions.Finding = &findingObject
	}

	if len(kpiInterface) != 0 {
		kpiObject := findingsapiv1.KpiType{}
		kpiAggregationType := kpiInterface[0].(map[string]interface{})["aggregation_type"].(string)
		kpiObject.AggregationType = &kpiAggregationType
		updateNoteOptions.Kpi = &kpiObject
	}

	if len(cardInterface) != 0 {
		cardObject := findingsapiv1.Card{}
		cardSection := cardInterface[0].(map[string]interface{})["section"].(string)
		cardTitle := cardInterface[0].(map[string]interface{})["title"].(string)
		cardSubtitle := cardInterface[0].(map[string]interface{})["subtitle"].(string)
		cardOrder := int64(cardInterface[0].(map[string]interface{})["order"].(int))
		cardRequiresConf, _ := strconv.ParseBool(cardInterface[0].(map[string]interface{})["requires_configuration"].(string))
		cardBadgeText := cardInterface[0].(map[string]interface{})["badge_text"].(string)
		cardBadgeImage := cardInterface[0].(map[string]interface{})["badge_image"].(string)
		cardObject.Section = &cardSection
		cardObject.Title = &cardTitle
		cardObject.Subtitle = &cardSubtitle
		cardObject.Order = &cardOrder
		cardObject.RequiresConfiguration = &cardRequiresConf
		cardObject.BadgeText = &cardBadgeText
		cardObject.BadgeImage = &cardBadgeImage
		cardObjectFindingNoteNamesList := make([]string, 0)
		for _, findingNoteName := range cardInterface[0].(map[string]interface{})["finding_note_names"].([]interface{}) {
			cardObjectFindingNoteNamesList = append(cardObjectFindingNoteNamesList, findingNoteName.(string))
		}
		cardObject.FindingNoteNames = cardObjectFindingNoteNamesList
		cardElementsList := make([]findingsapiv1.CardElement, 0)
		for _, cardElement := range cardInterface[0].(map[string]interface{})["elements"].([]interface{}) {
			cardElementObject := findingsapiv1.CardElement{}
			cardElementKind := cardElement.(map[string]interface{})["kind"].(string)
			cardElementText := cardElement.(map[string]interface{})["text"].(string)
			cardElementDefTime := cardElement.(map[string]interface{})["default_time_range"].(string)
			cardElementDefInterval := cardElement.(map[string]interface{})["default_interval"].(string)
			cardElementObject.Kind = &cardElementKind
			cardElementObject.Text = &cardElementText
			cardElementObject.DefaultTimeRange = &cardElementDefTime
			cardElementObject.DefaultInterval = &cardElementDefInterval

			cardValueTypeObject := findingsapiv1.CardValueType{}
			cardValueTypeInterface := cardElement.(map[string]interface{})["value_type"].([]interface{})
			cardValueTypeKind := cardValueTypeInterface[0].(map[string]interface{})["kind"].(string)
			cardValueTypeKpiNoteName := cardValueTypeInterface[0].(map[string]interface{})["kpi_note_name"].(string)
			cardValueTypeText := cardValueTypeInterface[0].(map[string]interface{})["text"].(string)
			cardValueTypeObject.Kind = &cardValueTypeKind
			cardValueTypeObject.KpiNoteName = &cardValueTypeKpiNoteName
			cardValueTypeObject.Text = &cardValueTypeText
			cardValueTypeFindingNoteNames := make([]string, 0)
			for _, cardValueTypeFindingNoteName := range cardValueTypeInterface[0].(map[string]interface{})["finding_note_names"].([]interface{}) {
				cardValueTypeFindingNoteNames = append(cardValueTypeFindingNoteNames, cardValueTypeFindingNoteName.(string))
			}
			cardValueTypeObject.FindingNoteNames = cardValueTypeFindingNoteNames
			cardElementObject.ValueType = &cardValueTypeObject

			cardValueTypesObjectList := make([]findingsapiv1.CardValueType, 0)
			for _, cardValueType := range cardElement.(map[string]interface{})["value_types"].([]interface{}) {
				cardValueTypeObject := findingsapiv1.CardValueType{}
				cardValueTypeKind := cardValueType.(map[string]interface{})["kind"].(string)
				cardValueTypeKpiNoteName := cardValueType.(map[string]interface{})["kpi_note_name"].(string)
				cardValueTypeText := cardValueType.(map[string]interface{})["text"].(string)
				cardValueTypeObject.Kind = &cardValueTypeKind
				cardValueTypeObject.KpiNoteName = &cardValueTypeKpiNoteName
				cardValueTypeObject.Text = &cardValueTypeText
				cardValueTypeFindingNoteNames := make([]string, 0)
				for _, cardValueTypeFindingNoteName := range cardValueType.(map[string]interface{})["finding_note_names"].([]interface{}) {
					cardValueTypeFindingNoteNames = append(cardValueTypeFindingNoteNames, cardValueTypeFindingNoteName.(string))
				}
				cardValueTypeObject.FindingNoteNames = cardValueTypeFindingNoteNames
				cardValueTypesObjectList = append(cardValueTypesObjectList, cardValueTypeObject)
			}
			cardElementsList = append(cardElementsList, cardElementObject)
			cardElementObject.ValueTypes = cardValueTypesObjectList
		}
		cardObject.Elements = cardElementsList
		updateNoteOptions.Card = &cardObject
	}

	if len(sectionInterface) != 0 {
		sectionObject := findingsapiv1.Section{}
		sectionTitle := sectionInterface[0].(map[string]interface{})["title"].(string)
		sectionImage := sectionInterface[0].(map[string]interface{})["image"].(string)
		sectionObject.Title = &sectionTitle
		sectionObject.Image = &sectionImage
		updateNoteOptions.Section = &sectionObject
	}

	if shortDescription != "" {
		updateNoteOptions.ShortDescription = &shortDescription
	}

	if longDescription != "" {
		updateNoteOptions.LongDescription = &longDescription
	}

	updateNoteOptions.Kind = &kind
	currentTime := strfmt.DateTime(time.Now())
	updateNoteOptions.ExpirationTime = &currentTime
	if expirationTime != "" {
		noteExpirationTime, err := strfmt.ParseDateTime(expirationTime)
		if err != nil {
			return fmt.Errorf("error occurred while creating note: %v", err)
		}
		updateNoteOptions.ExpirationTime = &noteExpirationTime
	}
	if createTime != "" {
		noteCreateTime, err := strfmt.ParseDateTime(createTime)
		if err != nil {
			return fmt.Errorf("error occurred while creating note: %v", err)
		}
		updateNoteOptions.CreateTime = &noteCreateTime
	}
	if updateTime != "" {
		noteUpdateTime, err := strfmt.ParseDateTime(updateTime)
		if err != nil {
			return fmt.Errorf("error occurred while creating note: %v", err)
		}
		updateNoteOptions.UpdateTime = &noteUpdateTime
	}

	if &shared != nil {
		updateNoteOptions.Shared = &shared
	}

	sess, err := meta.(ClientSession).FindingsV1API()
	if err != nil {
		return err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	accountID := userDetails.userAccount
	providerID := d.Get("provider_id").(string)
	noteID := d.Get("note_id").(string)

	updateNoteOptions.AccountID = &accountID
	updateNoteOptions.ProviderID = &providerID
	updateNoteOptions.ID = &noteID
	updateNoteOptions.NoteID = &noteID

	note, _, err := sess.UpdateNote(&updateNoteOptions)

	if err != nil {
		return fmt.Errorf("error occurred while creating note: %v", err)
	}
	if note.ID != nil {
		d.SetId(*note.ID)
		return resourceIBMNoteRead(d, meta)
	}

	return nil
}

func resourceIBMNoteDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).FindingsV1API()
	if err != nil {
		return err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	accountID := userDetails.userAccount
	providerID := d.Get("provider_id").(string)
	noteID := d.Id()

	deleteNoteOptions := sess.NewDeleteNoteOptions(accountID, providerID, noteID)
	_, err = sess.DeleteNote(deleteNoteOptions)

	if err != nil {
		return fmt.Errorf("error occurred while deleting note: %v", err)
	}

	d.SetId("")

	return nil
}
