provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

# Provision scc_si_providers data source instance

data "ibm_scc_si_providers" "providers" {
  limit = 4
}

# Provision scc_si_notes data source instance

data "ibm_scc_si_notes" "notes" {
  provider_id = var.provider_id
  page_size   = 3
}

# Provision scc_si_note data source instance

data "ibm_scc_si_note" "note" {
  provider_id = var.provider_id
  note_id = var.note_id
}

# Provision scc_si_occurrences data source instance
data "ibm_scc_si_occurrences" "occurrences" {
  provider_id = var.provider_id
  page_size   = 4
}

# Provision scc_si_occurrence data source instance

data "ibm_scc_si_occurrence" "occurrence" {
  provider_id = var.provider_id
  occurrence_id = var.occurrence_id
}

# Provision scc_si_note resource instance - Kind FINDING
resource "ibm_scc_si_note" "finding" {
  provider_id       = var.provider_id
  short_description = "Security Threat"
  long_description  = "Security Threat found in your account"
  kind              = "FINDING"
  note_id           = "finding"
  reported_by {
    id    = "scc-si-terraform"
    title = "SCC SI Terraform"
    url   = "https://cloud.ibm.com"
  }
  finding {
    severity = "LOW"
    next_steps {
      title = "Security Threat"
      url   = "https://cloud.ibm.com/security-compliance/findings"
    }
  }
}

# Provision scc_si_note resource instance - Kind KPI
resource "ibm_scc_si_note" "kpi" {
  provider_id       = var.provider_id
  short_description = "Security Threat"
  long_description  = "Security Threat found in your account"
  kind              = "KPI"
  note_id           = "kpi"
  reported_by {
    id    = "scc-si-terraform"
    title = "SCC SI Terraform"
    url   = "https://cloud.ibm.com"
  }
  kpi {
    aggregation_type = "SUM"
  }
}

# Provision scc_si_note resource instance - Kind Card (NUMERIC - FINDING_COUNT)
resource "ibm_scc_si_note" "num-card-finding" {
  provider_id       = var.provider_id
  short_description = "Security Threat"
  long_description  = "Security Threat found in your account"
  kind              = "CARD"
  note_id           = "num-card-finding"
  reported_by {
    id    = "scc-si-terraform"
    title = "SCC SI Terraform"
    url   = "https://cloud.ibm.com"
  }
  card {
    section            = "Terraform Insights"
    title              = "NUMERIC Finding Card"
    subtitle           = "Summary of Findings"
    finding_note_names = ["${var.account_id}/providers/scc/notes/finding"]
    elements {
      kind = "NUMERIC"
      text = "Issue Count"
      value_type {
        finding_note_names = ["${var.account_id}/providers/scc/notes/finding"]
        kind               = "FINDING_COUNT"
      }
    }
  }
}

# Provision scc_si_note resource instance - Kind Card (NUMERIC - KPI)
resource "ibm_scc_si_note" "num-card-kpi" {
  provider_id       = var.provider_id
  short_description = "Security Threat"
  long_description  = "Security Threat found in your account"
  kind              = "CARD"
  note_id           = "num-card-kpi"
  reported_by {
    id    = "scc-si-terraform"
    title = "SCC SI Terraform"
    url   = "https://cloud.ibm.com"
  }
  card {
    section            = "Terraform Insights"
    title              = "NUMERIC KPI Card"
    subtitle           = "Summary of KPIs"
    finding_note_names = ["${var.account_id}/providers/scc/notes/finding"]
    elements {
      kind = "NUMERIC"
      text = "Issue Count"
      value_type {
        kpi_note_name = "${var.account_id}/providers/scc/notes/kpi"
        kind          = "KPI"
      }
    }
  }
}

# Provision scc_si_note resource instance - Kind Card (BREAKDOWN - FINDING_COUNT)
resource "ibm_scc_si_note" "bkd-card-finding" {
  provider_id       = var.provider_id
  short_description = "Security Threat Breakdown Card"
  long_description  = "Security Threat found in your account"
  kind              = "CARD"
  note_id           = "bkd-card-finding"
  reported_by {
    id    = "scc-si-terraform"
    title = "SCC SI Terraform"
    url   = "https://cloud.ibm.com"
  }
  card {
    section            = "Terraform Insights"
    title              = "BREAKDOWN Finding Card"
    subtitle           = "Summary of Findings"
    finding_note_names = ["${var.account_id}/providers/scc/notes/finding"]
    elements {
      kind = "BREAKDOWN"
      text = "Issue Count"
      value_types {
        text               = "Issue Count"
        finding_note_names = ["${var.account_id}/providers/scc/notes/finding"]
        kind               = "FINDING_COUNT"
      }
    }
  }
}

# Provision scc_si_note resource instance - Kind Card (BREAKDOWN - KPI)
resource "ibm_scc_si_note" "bkd-card-kpi" {
  provider_id       = var.provider_id
  short_description = "Security Threat"
  long_description  = "Security Threat found in your account"
  kind              = "CARD"
  note_id           = "bkd-card-kpi"
  reported_by {
    id    = "scc-si-terraform"
    title = "SCC SI Terraform"
    url   = "https://cloud.ibm.com"
  }
  card {
    section            = "Terraform Insights"
    title              = "BREAKDOWN KPI Card"
    subtitle           = "Summary of KPIs"
    finding_note_names = ["${var.account_id}/providers/scc/notes/finding"]
    elements {
      kind = "BREAKDOWN"
      text = "Issue Count"
      value_types {
        text          = "Issue Count"
        kpi_note_name = "${var.account_id}/providers/scc/notes/kpi"
        kind          = "KPI"
      }
    }
  }
}

# Provision scc_si_note resource instance - Kind Card (TIME_SERIES - FINDING_COUNT)
resource "ibm_scc_si_note" "ts-card-finding" {
  provider_id       = var.provider_id
  short_description = "Security Threat"
  long_description  = "Security Threat found in your account"
  kind              = "CARD"
  note_id           = "ts-card-finding"
  reported_by {
    id    = "scc-si-terraform"
    title = "SCC SI Terraform"
    url   = "https://cloud.ibm.com"
  }
  card {
    section            = "Terraform Insights"
    title              = "TIME_SERIES Finding Card"
    subtitle           = "Summary of Findings"
    finding_note_names = ["${var.account_id}/providers/scc/notes/finding"]
    elements {
      kind               = "TIME_SERIES"
      text               = "Issue Count"
      default_time_range = "3d"
      value_types {
        text               = "Issue Count"
        finding_note_names = ["${var.account_id}/providers/scc/notes/finding"]
        kind               = "FINDING_COUNT"
      }
    }
  }
}


// Provision scc_si_occurrence resource instance - Kind FINDING

resource "ibm_scc_si_occurrence" "finding-occurrence" {
  provider_id   = var.provider_id
  note_name     = "${var.account_id}/providers/${var.provider_id}/notes/${var.note_id}"
  kind          = "FINDING"
  occurrence_id = "finding-occ"
  resource_url  = "https://cloud.ibm.com"
  remediation   = "Limit the cluster access"
  finding {
    severity  = "HIGH"
    certainty = "LOW"
    next_steps {
      title = "Security Threat"
      url   = "https://cloud.ibm.com/security-compliance/findings"
    }
  }
}

// Provision scc_si_occurrence resource instance - Kind KPI

resource "ibm_scc_si_occurrence" "kpi-occurrence" {
  provider_id   = var.provider_id
  note_name     = "${var.account_id}/providers/${var.provider_id}/notes/${var.note_id}"
  kind          = "KPI"
  occurrence_id = "kpi-occ"
  resource_url  = "https://cloud.ibm.com"
  remediation   = "Limit the cluster access"
  kpi {
    value = 40
    total = 100
  }
}