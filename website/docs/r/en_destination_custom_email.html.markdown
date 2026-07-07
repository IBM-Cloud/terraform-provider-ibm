---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_destination_custom_email'
description: |-
  Manages Event Notification Custom Email destinations.
---

# ibm_en_destination_custom_email

Create, update, or delete a Custom Email destination by using IBM Cloud™ Event Notifications.

## Example usage

### Production Custom Email Destination

```terraform
resource "ibm_en_destination_custom_email" "custom_domain_en_destination" {
  instance_guid         = ibm_resource_instance.en_terraform_test_resource.guid
  name                  = "Custom Email EN Destination"
  type                  = "smtp_custom"
  collect_failed_events = true
  is_sandbox            = false
  description           = "Production custom email destination for event notification"
  config {
    params {
      domain = "mailx.com"
    }
  }
}
```

### Sandbox Custom Email Destination

```terraform
resource "ibm_en_destination_custom_email" "sandbox_custom_domain_en_destination" {
  instance_guid         = ibm_resource_instance.en_terraform_test_resource.guid
  name                  = "Sandbox Custom Email EN Destination"
  type                  = "smtp_custom"
  collect_failed_events = false
  is_sandbox            = true
  description           = "Sandbox custom email destination for testing"
  config {
    params {
      domain = "sandbox.example.com"
    }
  }
}
```

### Upgrading Sandbox to Production

```terraform
# Step 1: Create as sandbox
resource "ibm_en_destination_custom_email" "upgradeable_destination" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
  name          = "Upgradeable Custom Email Destination"
  type          = "smtp_custom"
  is_sandbox    = true
  description   = "Initially created as sandbox"
  config {
    params {
      domain = "upgrade.example.com"
    }
  }
}

# Step 2: Upgrade to production by changing is_sandbox to false
# After applying this change, the destination will be upgraded to production
# and DKIM/SPF records will be generated for domain verification
resource "ibm_en_destination_custom_email" "upgradeable_destination" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
  name          = "Upgradeable Custom Email Destination"
  type          = "smtp_custom"
  is_sandbox    = false  # Changed from true to false
  description   = "Upgraded to production"
  config {
    params {
      domain = "upgrade.example.com"
    }
  }
}
```

## Sandbox vs Production Destinations

### Sandbox Mode (`is_sandbox = true`)
- **Purpose**: Testing and development without domain verification
- **Type**: `smtp_custom_sandbox`
- **Domain Verification**: Not required
- **DKIM/SPF**: Not generated
- **Subscriptions**: Do not require `from_name` and `from_email` attributes
- **Use Case**: Quick testing, development, proof of concept
- **Limitations**: May have sending limits, not suitable for production use

### Production Mode (`is_sandbox = false`, default)
- **Purpose**: Production email sending with verified domain
- **Type**: `smtp_custom`
- **Domain Verification**: Required (SPF and DKIM)
- **DKIM/SPF**: Automatically generated upon creation
- **Subscriptions**: Require `from_name` and `from_email` attributes
- **Use Case**: Production workloads, verified email sending
- **Upgrade Path**: Can be upgraded from sandbox mode

### Important Notes

**Sandbox Destinations:**
- Domain can be updated/changed at any time
- No DNS verification required
- Ideal for testing before production deployment
- Can be upgraded to production (one-way operation)

**Production Destinations:**
- Domain is **immutable** after creation (cannot be changed)
- Requires DNS verification (SPF and DKIM records)
- Cannot be downgraded to sandbox mode
- Must complete domain verification before sending emails

**Upgrade Process:**
- Change `is_sandbox` from `true` to `false`
- Domain verification records (DKIM/SPF) will be automatically generated
- Complete DNS verification process (see below)
- Downgrade from production to sandbox is **not supported**

**NOTE:**
- To perform the verification for SPF and DKIM please follow the instructions here: https://cloud.ibm.com/docs/event-notifications?topic=event-notifications-en-destinations-custom-email
- `verification_type` is Custom Email Destination update parameter which can be used to verify the status of verification depending on the type of verification.

## Domain Configuration and Verification Process (Production Only)

This process is only required for production destinations (`is_sandbox = false`).

1. Select the configure overflow menu for the destination you want to verify.

2. Create Sender Policy Framework (SPF), which is used to authenticate the sender of an email. SPF specifies the mail servers that are allowed to send email for your domain.
    - Open your DNS hosting provider for the domain name configured
    - Create a new TXT record with your domain name registrar with the name and value provided in the configure screen for SPF

3. Create DomainKeys Identified Mail (DKIM), which allows an organization to take responsibility for transmitting a message by signing it. DKIM allows the receiver to check the email that claimed to have come from a specific domain, is authorized by the owner of that domain.
    - Open your DNS hosting provider for the domain name configured
    - Create a new TXT record with your domain name registrar with the name and value provided in the configure screen for DKIM

4. Save the TXT records

5. In the destination verify screen, click on Verify buttons for both SPF and DKIM.

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `name` - (Required, String) The Destintion name.

- `description` - (Optional, String) The Destination description.

- `type` - (Required, String) smtp_custom.

- `collect_failed_events` - (Boolean) Toggle switch to enable collect failed event in Cloud Object Storage bucket.

- `is_sandbox` - (Optional, Boolean) Toggle switch to enable sandbox mode. Default value is `false`.
  - When `true`: Creates a sandbox destination (`smtp_custom_sandbox` type) for testing without domain verification
  - When `false`: Creates a production destination (`smtp_custom` type) requiring domain verification
  - **Upgrade**: Can be changed from `true` to `false` to upgrade sandbox to production
  - **Downgrade**: Cannot be changed from `false` to `true` (production to sandbox downgrade is not supported)
  - **Domain Updates**:
    - Sandbox destinations: Domain can be updated at any time
    - Production destinations: Domain is immutable after creation

- `config` - (Optional, List) Payload describing a destination configuration.

  Nested scheme for **config**:

  - `params` - (Required, List)

  Nested scheme for **params**:

  - `domain` - (Required, String) The Custom Domain.
  - `spf` - (Optional, List) The SPF attributes.
		Nested schema for **spf**:
			* `txt_name` - (Optional, String) spf text name.
			  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.
			* `txt_value` - (Optional, String) spf text value.
			  * Constraints: The maximum length is `500` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.
			* `verification` - (Optional, String) spf verification.
			  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.
  - `dkim` - (Optional, List) The DKIM attributes.
		Nested schema for **dkim**:
			* `public_key` - (Optional, String) dkim public key.
			  * Constraints: The maximum length is `500` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.
			* `selector` - (Optional, String) dkim selector.
			  * Constraints: The maximum length is `500` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.
			* `verification` - (Optional, String) dkim verification.
			  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.      

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `custom_domain_en_destination`.
- `destination_id` - (String) The unique identifier of the created destination.
- `subscription_count` - (Integer) Number of subscriptions.
  - Constraints: The minimum value is `0`.
- `subscription_names` - (List) List of subscriptions.
- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_destination_custom_email` resource by using `id`.

The `id` property can be formed from `instance_guid`, and `destination_id` in the following format:

```
<instance_guid>/<destination_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.

- `destination_id`: A string. Unique identifier for Destination.

**Example**

```
$ terraform import ibm_en_destination_custom_email.custom_domain_email_en_destination <instance_guid>/<destination_id>
```
