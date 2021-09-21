// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccSiNotesDataSourceBasic(t *testing.T) {
	apiNoteProviderID := fmt.Sprintf("tf_provider_id_%d", acctest.RandIntRange(10, 100))
	apiNoteShortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	apiNoteLongDescription := fmt.Sprintf("tf_long_description_%d", acctest.RandIntRange(10, 100))
	apiNoteKind := "FINDING"
	apiNoteID := fmt.Sprintf("tf_note_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccSiNotesDataSourceConfigBasic(scc_si_account, apiNoteProviderID, apiNoteShortDescription, apiNoteLongDescription, apiNoteKind, apiNoteID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "provider_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "notes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "notes.0.short_description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "notes.0.long_description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "notes.0.kind"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "notes.0.related_url.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "notes.0.create_time"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "notes.0.update_time"),
				),
			},
		},
	})
}

func TestAccIBMSccSiNotesDataSourceAllArgs(t *testing.T) {
	apiNoteProviderID := fmt.Sprintf("tf_provider_id_%d", acctest.RandIntRange(10, 100))
	apiNoteShortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	apiNoteLongDescription := fmt.Sprintf("tf_long_description_%d", acctest.RandIntRange(10, 100))
	apiNoteKind := "FINDING"
	apiNoteID := fmt.Sprintf("tf_note_id_%d", acctest.RandIntRange(10, 100))
	apiNoteShared := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccSiNotesDataSourceConfig(scc_si_account, apiNoteProviderID, apiNoteShortDescription, apiNoteLongDescription, apiNoteKind, apiNoteID, apiNoteShared),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "provider_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "notes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "notes.0.short_description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "notes.0.long_description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "notes.0.kind"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "notes.0.related_url.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "notes.0.related_url.0.label"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "notes.0.related_url.0.url"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "notes.0.create_time"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "notes.0.update_time"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "notes.0.shared"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "notes.0.reported_by.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_notes.scc_si_notes", "notes.0.finding.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSccSiNotesDataSourceConfigBasic(accountID string, apiNoteProviderID string, apiNoteShortDescription string, apiNoteLongDescription string, apiNoteKind string, apiNoteID string) string {
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

		data "ibm_scc_si_notes" "scc_si_notes" {
			account_id = ibm_scc_si_note.scc_si_note.account_id
			provider_id = ibm_scc_si_note.scc_si_note.provider_id
		} 
	`, accountID, apiNoteProviderID, apiNoteShortDescription, apiNoteLongDescription, apiNoteKind, apiNoteID)
}

func testAccCheckIBMSccSiNotesDataSourceConfig(accountID string, apiNoteProviderID string, apiNoteShortDescription string, apiNoteLongDescription string, apiNoteKind string, apiNoteID string, apiNoteShared string) string {
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

		data "ibm_scc_si_notes" "scc_si_notes" {
			account_id = ibm_scc_si_note.scc_si_note.account_id
			provider_id = ibm_scc_si_note.scc_si_note.provider_id
		}
	`, accountID, apiNoteProviderID, apiNoteShortDescription, apiNoteLongDescription, apiNoteKind, apiNoteID, apiNoteShared)
}
