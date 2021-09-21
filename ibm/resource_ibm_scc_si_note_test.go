// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/scc-go-sdk/findingsv1"
)

func TestAccIBMSccSiNoteCardNumeric(t *testing.T) {
	var conf findingsv1.APINote
	providerID := fmt.Sprintf("tf_provider_id_%d", acctest.RandIntRange(10, 100))
	shortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	longDescription := fmt.Sprintf("tf_long_description_%d", acctest.RandIntRange(10, 100))
	kind := "CARD"
	noteID := fmt.Sprintf("tf_note_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSccSiNoteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccSiNoteCardNumericBasic(scc_si_account, providerID, shortDescription, longDescription, kind, noteID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccSiNoteExists("ibm_scc_si_note.scc_si_note", conf),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "provider_id", providerID),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "short_description", shortDescription),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "long_description", longDescription),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "kind", kind),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "note_id", noteID),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "card.0.section", "section"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "card.0.title", "title"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "card.0.subtitle", "subtitle"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "card.0.elements.0.kind", "NUMERIC"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "card.0.elements.0.value_type.0.kind", "FINDING_COUNT"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "card.0.elements.0.value_type.0.finding_note_names.0", fmt.Sprintf("providers/%s/notes/note123", providerID)),
				),
			},
		},
	})
}

func TestAccIBMSccSiNoteCardBreakdown(t *testing.T) {
	var conf findingsv1.APINote
	providerID := fmt.Sprintf("tf_provider_id_%d", acctest.RandIntRange(10, 100))
	shortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	longDescription := fmt.Sprintf("tf_long_description_%d", acctest.RandIntRange(10, 100))
	kind := "CARD"
	noteID := fmt.Sprintf("tf_note_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSccSiNoteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccSiNoteCardBreakdownBasic(scc_si_account, providerID, shortDescription, longDescription, kind, noteID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccSiNoteExists("ibm_scc_si_note.scc_si_note", conf),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "provider_id", providerID),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "short_description", shortDescription),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "long_description", longDescription),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "kind", kind),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "note_id", noteID),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "card.0.section", "section"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "card.0.title", "title"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "card.0.subtitle", "subtitle"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "card.0.elements.0.kind", "BREAKDOWN"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "card.0.elements.0.value_types.0.kind", "FINDING_COUNT"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "card.0.elements.0.value_types.0.finding_note_names.0", fmt.Sprintf("providers/%s/notes/note123", providerID)),
				),
			},
		},
	})
}

func TestAccIBMSccSiNoteCardTimeSeries(t *testing.T) {
	var conf findingsv1.APINote
	providerID := fmt.Sprintf("tf_provider_id_%d", acctest.RandIntRange(10, 100))
	shortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	longDescription := fmt.Sprintf("tf_long_description_%d", acctest.RandIntRange(10, 100))
	kind := "CARD"
	noteID := fmt.Sprintf("tf_note_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSccSiNoteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccSiNoteCardTimeSeriesBasic(scc_si_account, providerID, shortDescription, longDescription, kind, noteID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccSiNoteExists("ibm_scc_si_note.scc_si_note", conf),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "provider_id", providerID),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "short_description", shortDescription),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "long_description", longDescription),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "kind", kind),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "note_id", noteID),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "card.0.section", "section"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "card.0.title", "title"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "card.0.subtitle", "subtitle"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "card.0.elements.0.kind", "TIME_SERIES"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "card.0.elements.0.value_types.0.kind", "FINDING_COUNT"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "card.0.elements.0.value_types.0.finding_note_names.0", fmt.Sprintf("providers/%s/notes/note123", providerID)),
				),
			},
		},
	})
}

func TestAccIBMSccSiNoteBasic(t *testing.T) {
	var conf findingsv1.APINote
	providerID := fmt.Sprintf("tf_provider_id_%d", acctest.RandIntRange(10, 100))
	shortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	longDescription := fmt.Sprintf("tf_long_description_%d", acctest.RandIntRange(10, 100))
	kind := "FINDING"
	noteID := fmt.Sprintf("tf_note_id_%d", acctest.RandIntRange(10, 100))
	shortDescriptionUpdate := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	longDescriptionUpdate := fmt.Sprintf("tf_long_description_%d", acctest.RandIntRange(10, 100))
	kindUpdate := "FINDING"
	noteIDUpdate := noteID

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSccSiNoteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccSiNoteConfigBasic(scc_si_account, providerID, shortDescription, longDescription, kind, noteID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccSiNoteExists("ibm_scc_si_note.scc_si_note", conf),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "provider_id", providerID),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "short_description", shortDescription),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "long_description", longDescription),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "kind", kind),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "note_id", noteID),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSccSiNoteConfigBasic(scc_si_account, providerID, shortDescriptionUpdate, longDescriptionUpdate, kindUpdate, noteIDUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "provider_id", providerID),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "short_description", shortDescriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "long_description", longDescriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "kind", kindUpdate),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "note_id", noteIDUpdate),
				),
			},
		},
	})
}

func TestAccIBMSccSiNoteAllArgs(t *testing.T) {
	var conf findingsv1.APINote
	providerID := fmt.Sprintf("tf_provider_id_%d", acctest.RandIntRange(10, 100))
	shortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	longDescription := fmt.Sprintf("tf_long_description_%d", acctest.RandIntRange(10, 100))
	kind := "FINDING"
	noteID := fmt.Sprintf("tf_note_id_%d", acctest.RandIntRange(10, 100))
	shared := "true"
	shortDescriptionUpdate := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	longDescriptionUpdate := fmt.Sprintf("tf_long_description_%d", acctest.RandIntRange(10, 100))
	kindUpdate := "FINDING"
	noteIDUpdate := noteID
	sharedUpdate := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSccSiNoteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccSiNoteConfig(scc_si_account, providerID, shortDescription, longDescription, kind, noteID, shared),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccSiNoteExists("ibm_scc_si_note.scc_si_note", conf),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "provider_id", providerID),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "short_description", shortDescription),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "long_description", longDescription),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "kind", kind),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "note_id", noteID),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "shared", shared),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSccSiNoteConfig(scc_si_account, providerID, shortDescriptionUpdate, longDescriptionUpdate, kindUpdate, noteIDUpdate, sharedUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "provider_id", providerID),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "short_description", shortDescriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "long_description", longDescriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "kind", kindUpdate),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "note_id", noteIDUpdate),
					resource.TestCheckResourceAttr("ibm_scc_si_note.scc_si_note", "shared", sharedUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_scc_si_note.scc_si_note",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMSccSiNoteInvalid(t *testing.T) {
	providerID := fmt.Sprintf("tf_provider_id_%d", acctest.RandIntRange(10, 100))
	shortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	longDescription := fmt.Sprintf("tf_long_description_%d", acctest.RandIntRange(10, 100))
	kind := "FINDING"
	noteID := fmt.Sprintf("tf_note_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSccSiNoteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMSccSiNoteConfigInvalid(scc_si_account, providerID, shortDescription, longDescription, kind, noteID),
				ExpectError: regexp.MustCompile("Invalid combination of arguments"),
			},
		},
	})
}

func TestAccIBMSccSiNoteEmpty(t *testing.T) {
	providerID := fmt.Sprintf("tf_provider_id_%d", acctest.RandIntRange(10, 100))
	shortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	longDescription := fmt.Sprintf("tf_long_description_%d", acctest.RandIntRange(10, 100))
	kind := "FINDING"
	noteID := fmt.Sprintf("tf_note_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSccSiNoteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMSccSiNoteConfigEmpty(scc_si_account, providerID, shortDescription, longDescription, kind, noteID),
				ExpectError: regexp.MustCompile("Invalid combination of arguments"),
			},
		},
	})
}

func testAccCheckIBMSccSiNoteCardNumericBasic(accountID string, providerID string, shortDescription string, longDescription string, kind string, noteID string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_si_note" "scc_si_note" {
			account_id = "%s"
			provider_id = "%s"
			short_description = "%s"
			long_description = "%s"
			kind = "%s"
			note_id = "%s"
			reported_by {
				id = "id"
				title = "title"
				url = "url"
			}
			card {
				section            = "section"
				title              = "title"
				subtitle           = "subtitle"
				finding_note_names = ["providers/%s/notes/note123"]
				elements {
				  kind = "NUMERIC"
				  text = "text"
				  value_type {
					finding_note_names = ["providers/%s/notes/note123"]
					kind               = "FINDING_COUNT"
				  }
				}
			}
		}
	`, accountID, providerID, shortDescription, longDescription, kind, noteID, providerID, providerID)
}

func testAccCheckIBMSccSiNoteCardBreakdownBasic(accountID string, providerID string, shortDescription string, longDescription string, kind string, noteID string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_si_note" "scc_si_note" {
			account_id = "%s"
			provider_id = "%s"
			short_description = "%s"
			long_description = "%s"
			kind = "%s"
			note_id = "%s"
			reported_by {
				id = "id"
				title = "title"
				url = "url"
			}
			card {
				section            = "section"
				title              = "title"
				subtitle           = "subtitle"
				finding_note_names = ["providers/%s/notes/note123"]
				elements {
				  kind = "BREAKDOWN"
				  text = "text"
				  default_time_range = "3d"
				  value_types {
					finding_note_names = ["providers/%s/notes/note123"]
					kind               = "FINDING_COUNT"
					text = "text"
				  }
				}
			}
		}
	`, accountID, providerID, shortDescription, longDescription, kind, noteID, providerID, providerID)
}

func testAccCheckIBMSccSiNoteCardTimeSeriesBasic(accountID string, providerID string, shortDescription string, longDescription string, kind string, noteID string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_si_note" "scc_si_note" {
			account_id = "%s"
			provider_id = "%s"
			short_description = "%s"
			long_description = "%s"
			kind = "%s"
			note_id = "%s"
			reported_by {
				id = "id"
				title = "title"
				url = "url"
			}
			card {
				section            = "section"
				title              = "title"
				subtitle           = "subtitle"
				finding_note_names = ["providers/%s/notes/note123"]
				elements {
				  kind = "TIME_SERIES"
				  text = "text"
				  default_interval = "3d"
				  default_time_range = "3d"
				  value_types {
					finding_note_names = ["providers/%s/notes/note123"]
					kind               = "FINDING_COUNT"
					text = "text"
				  }
				}
			}
		}
	`, accountID, providerID, shortDescription, longDescription, kind, noteID, providerID, providerID)
}

func testAccCheckIBMSccSiNoteConfigBasic(accountID string, providerID string, shortDescription string, longDescription string, kind string, noteID string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_si_note" "scc_si_note" {
			account_id = "%s"
			provider_id = "%s"
			short_description = "%s"
			long_description = "%s"
			kind = "%s"
			note_id = "%s"
			reported_by {
				id = "id"
				title = "title"
				url = "url"
			}
			finding {
				severity = "LOW"
				next_steps {
					title = "title"
					url = "url"
				}
			}
		}
	`, accountID, providerID, shortDescription, longDescription, kind, noteID)
}

func testAccCheckIBMSccSiNoteConfig(accountID string, providerID string, shortDescription string, longDescription string, kind string, noteID string, shared string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_si_note" "scc_si_note" {
			account_id = "%s"
			provider_id = "%s"
			short_description = "%s"
			long_description = "%s"
			kind = "%s"
			note_id = "%s"
			reported_by {
				id = "id"
				title = "title"
				url = "url"
			}
			related_url {
				label = "label"
				url = "url"
			}
			shared = %s
			finding {
				severity = "LOW"
				next_steps {
					title = "title"
					url = "url"
				}
			}
		}
	`, accountID, providerID, shortDescription, longDescription, kind, noteID, shared)
}

func testAccCheckIBMSccSiNoteConfigInvalid(accountID string, providerID string, shortDescription string, longDescription string, kind string, noteID string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_si_note" "scc_si_note" {
			account_id = "%s"
			provider_id = "%s"
			short_description = "%s"
			long_description = "%s"
			kind = "%s"
			note_id = "%s"
			reported_by {
				id = "id"
				title = "title"
				url = "url"
			}
			related_url {
				label = "label"
				url = "url"
			}
			finding {
				severity = "LOW"
				next_steps {
					title = "title"
					url = "url"
				}
			}
			kpi {
				aggregation_type = "SUM"
			}
			card {
				section = "section"
				title = "title"
				subtitle = "subtitle"
				order = 1
				finding_note_names = [ "finding_note_names" ]
				badge_text = "badge_text"
				badge_image = "badge_image"
				elements {
					text = "text"
					kind = "NUMERIC"
					default_time_range = "1d"
					value_type {
						kind = "KPI"
						kpi_note_name = "kpi_note_name"
						text = "text"
						finding_note_names = [ "finding_note_names" ]
					}
					value_types {
						kind = "KPI"
						kpi_note_name = "kpi_note_name"
						text = "text"
						finding_note_names = [ "finding_note_names" ]
					}
					default_interval = "default_interval"
				}
			}
			section {
				title = "title"
				image = "image"
			}
		}
	`, accountID, providerID, shortDescription, longDescription, kind, noteID)
}

func testAccCheckIBMSccSiNoteConfigEmpty(accountID string, providerID string, shortDescription string, longDescription string, kind string, noteID string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_si_note" "scc_si_note" {
			account_id = "%s"
			provider_id = "%s"
			short_description = "%s"
			long_description = "%s"
			kind = "%s"
			note_id = "%s"
			reported_by {
				id = "id"
				title = "title"
				url = "url"
			}
			related_url {
				label = "label"
				url = "url"
			}
		}
	`, accountID, providerID, shortDescription, longDescription, kind, noteID)
}

func testAccCheckIBMSccSiNoteExists(n string, obj findingsv1.APINote) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		findingsClient, err := testAccProvider.Meta().(ClientSession).FindingsV1()
		if err != nil {
			return err
		}

		getNoteOptions := &findingsv1.GetNoteOptions{}

		parts, err := sepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		findingsClient.AccountID = &parts[0]

		getNoteOptions.SetProviderID(parts[1])
		getNoteOptions.SetNoteID(parts[2])

		apiNote, _, err := findingsClient.GetNote(getNoteOptions)
		if err != nil {
			return err
		}

		obj = *apiNote
		return nil
	}
}

func testAccCheckIBMSccSiNoteDestroy(s *terraform.State) error {
	findingsClient, err := testAccProvider.Meta().(ClientSession).FindingsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_si_note" {
			continue
		}

		getNoteOptions := &findingsv1.GetNoteOptions{}

		parts, err := sepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		findingsClient.AccountID = &parts[0]

		getNoteOptions.SetProviderID(parts[1])
		getNoteOptions.SetNoteID(parts[2])

		// Try to find the key
		_, response, err := findingsClient.GetNote(getNoteOptions)

		if err == nil {
			return fmt.Errorf("scc_si_note still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for scc_si_note (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
