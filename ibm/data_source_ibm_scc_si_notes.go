package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIBMNotes() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMNotesRead,
		Schema: map[string]*schema.Schema{
			"provider_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"notes": {
				Type:        schema.TypeList,
				Description: "Collection of SA Notes",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"short_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"long_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"related_url": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"label": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"url": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"expiration_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"create_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"update_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shared": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"reported_by": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"url": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"finding": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"next_steps": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"url": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"title": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"kpi": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aggregation_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"card": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"section": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"subtitle": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"order": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"finding_note_names": {
										Type:     schema.TypeList,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"requires_configuration": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"badge_text": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"badge_image": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"elements": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"kind": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"text": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"value_type": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"kind": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"finding_note_names": {
																Type:     schema.TypeList,
																Computed: true,
																Elem:     &schema.Schema{Type: schema.TypeString},
															},
															"kpi_note_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"text": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"value_types": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"kind": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"finding_note_names": {
																Type:     schema.TypeList,
																Computed: true,
																Elem:     &schema.Schema{Type: schema.TypeString},
															},
															"kpi_note_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"text": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"default_time_range": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"default_interval": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"section": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"image": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMNotesRead(d *schema.ResourceData, meta interface{}) error {
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

	listNotesOptions := sess.NewListNotesOptions(accountID, providerID)
	notes, _, err := sess.ListNotes(listNotesOptions)

	if err != nil {
		return fmt.Errorf("error occurred while listing notes: %v", err)
	}

	notesList := make([]map[string]interface{}, 0)
	for _, note := range notes.Notes {
		noteObject := map[string]interface{}{}
		noteObject["id"] = note.ID
		noteObject["kind"] = note.Kind
		noteObject["short_description"] = note.ShortDescription
		noteObject["long_description"] = note.LongDescription
		if note.ExpirationTime != nil {
			noteObject["expiration_time"] = note.ExpirationTime.String()
		}
		noteObject["create_time"] = note.CreateTime.String()
		noteObject["update_time"] = note.UpdateTime.String()
		noteObject["shared"] = "false"
		if note.Shared != nil {
			noteObject["shared"] = "true"
		}
		noteRelatedURLList := make([]map[string]interface{}, 0)
		for _, noteRelatedURL := range note.RelatedURL {
			noteRelatedURLObject := map[string]interface{}{}
			noteRelatedURLObject["label"] = noteRelatedURL.Label
			noteRelatedURLObject["url"] = noteRelatedURL.URL
			noteRelatedURLList = append(noteRelatedURLList, noteRelatedURLObject)
		}
		noteObject["related_url"] = noteRelatedURLList
		noteReportedByObjectList := make([]map[string]interface{}, 1)
		noteReportedByObject := map[string]interface{}{}
		noteReportedByObject["id"] = note.ReportedBy.ID
		noteReportedByObject["title"] = note.ReportedBy.Title
		noteReportedByObject["url"] = note.ReportedBy.URL
		noteReportedByObjectList[0] = noteReportedByObject
		noteObject["reported_by"] = noteReportedByObjectList

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
			noteFindingTypeObjectList[0] = noteFindingTypeObject
			noteObject["finding"] = noteFindingTypeObjectList
		}

		if note.Kpi != nil {
			noteKpiTypeObjectList := make([]map[string]interface{}, 1)
			noteKpiTypeObject := map[string]interface{}{}
			noteKpiTypeObject["aggregation_type"] = note.Kpi.AggregationType
			noteKpiTypeObjectList[0] = noteKpiTypeObject
			noteObject["kpi"] = noteKpiTypeObjectList
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
			noteObject["card"] = noteCardObjectList
		}

		if note.Section != nil {
			noteSectionObjectList := make([]map[string]interface{}, 1)
			noteSectionObject := map[string]interface{}{}
			noteSectionObject["image"] = note.Section.Image
			noteSectionObject["title"] = note.Section.Title
			noteSectionObjectList[0] = noteSectionObject
			noteObject["section"] = noteSectionObjectList
		}
		notesList = append(notesList, noteObject)
	}

	d.Set("notes", notesList)
	d.SetId(accountID)

	return nil
}
