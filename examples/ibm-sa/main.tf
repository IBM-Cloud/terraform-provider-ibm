#################################

# Get note 
data "ibm_sa_note" "note" {
  provider_id = var.provider_id
  note_id = var.note_id
}
output "ibm_sa_note" {
  value = data.ibm_sa_note.note
}

#################################

# List notes
data "ibm_sa_notes" "notes" {
  provider_id = var.provider_id
}
output "ibm_sa_notes" {
  value = data.ibm_sa_notes.notes
}

#################################

# Get channel
data "ibm_sa_notification_channel" "channel" {
  channel_id = var.channel_id
}
output "ibm_sa_notification_channel" {
  value = data.ibm_sa_notification_channel.channel
}

#################################

# List channels

data "ibm_sa_notification_channels" "channels" {
}
output "ibm_sa_notification_channels" {
  value = data.ibm_sa_notification_channels.channels
}

#################################

# Create Channel

resource "ibm_sa_notification_channel" "create" {
  name              = "hello"
  type              = "Webhook"
  endpoint          = "http://cloud.ibm.com"
  enabled           = "true"
  description       = "channel"
  severity          = ["low", "medium"]
  alert_source {
    provider_name = "ALL"
    finding_types = ["ALL"]
  }
}
output "ibm_sa_notification_channel" {
  value = ibm_sa_notification_channel.create
}

#################################

# Create Note

resource "ibm_sa_note" "create" {
  provider_id                    = "test"
  short_description              = "short"
  long_description               = "long"
  kind                           = "FINDING"
  related_url {
    label   = "rel_label"
    url     = "rel_url"
  }
  expiration_time                = "2006-01-02 15:04:11"
  create_time                    = "2006-01-02 15:04:00"
  update_time                    = "2006-01-02 15:04:11"
  note_id                        = "id"
  shared                         = "true"
  reported_by {
    id    = "rep_id"
    title = "rep_title"
    url   = "rep_url"
  }
  finding {
    severity = "HIGH"
    next_steps {
      url   = "next_url"
      title = "next_title"
    }
  }
}
output "ibm_sa_note" {
  value = ibm_sa_note.create
}

#################################