---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : reservation"
description: |-
  Manages IBM reservation.
---

# ibm_is_reservation
Create, update, or delete a reservation.

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "us-south"
}
```


## Example usage

```terraform
resource "ibm_is_reservation" "example" {
  capacity {
    total = 5
  }
  committed_use {
    term = "one_year"
  }
  profile {
    name          = "ba2-2x8"
    resource_type = "instance_profile"
  }
  zone = "us-east-3"
  name = "reservation-terraform-1"
}
```


## Argument reference
Review the argument references that you can specify for your resource. 

- `affinity_policy`  - (Optional, String)  affinity policy to use for this reservation. [ **restricted** ]
- `capacity` - (List) The capacity reservation configuration to use.

  Nested scheme for `capacity`:
    - `total` - (Integer) The total amount to use for this capacity reservation.
- `committed_use` - (List) The committed use configuration to use for this reservation.

  Nested scheme for `committed_use`:
    - `expiration_policy` - (Optional, String) The policy to apply when the committed use term expires. [**release**, **renew**]
    ~> **Note:**
    **&#x2022;** `release` Release any available capacity and let the reservation expire.</br>
    **&#x2022;** `renew` Renew for another term, provided the term remains listed in the reservation_terms for the profile. Otherwise, let the reservation expire.</br>
    - `term` - (String) The term for this committed use reservation. [**one_year**, **three_year**]
- `name` - (Optional, String) The name for this reservation. The name must not be used by another reservation in the region. If unspecified, the name will be a hyphenated list of randomly-selected words.
- `profile` - (List) The virtual server instance profile to use for this reservation.

  Nested scheme for `profile`:
    - `name` - (String) The globally unique name of the profile. 
    - `resource_type` - (String) The resource type of the profile. [ **instance_profile** ]
- `resource_group` - (Optional, List) The resource group to use. If unspecified, the account's default resource group will be used.

  Nested scheme for `resource_group`: 
  - `id` - (String) The unique identifier for this resource group.
- `zone` - (String) The zone to use for this reservation.


## Attribute reference
You can access the following attribute references after your data source is created. 
- `affinity_policy`  - (String) The affinity policy to use for this reservation.
- `capacity` - (List) The capacity configuration for this reservation. If absent, this reservation has no assigned capacity.

  Nested scheme for `capacity`:
  - `allocated` - (Integer) The amount allocated to this capacity reservation.
  - `available` - (Integer) The amount of this capacity reservation available for new attachments.
  - `status` - (String) The status of the capacity reservation.

    ->**status** 
      </br>&#x2022; allocating: The capacity reservation is being allocated for use
      </br>&#x2022; allocated: The total capacity of the reservation has been allocated for use
      </br>&#x2022; degraded: The capacity reservation has been allocated for use, but some of the capacity is not available
      </br>&#x2022; unallocated: The capacity reservation is not allocated for use
  - `total` - (Integer) The total amount of this capacity reservation.
  - `used` - (Integer) The amount of this capacity reservation used by existing attachments.
- `committed_use` - (List) The committed use configuration for this reservation. If absent, this reservation has no commitment for use.

  Nested scheme for `committed_use`:
  - `expiration_at` - (Timestamp) The expiration date and time for this committed use reservation.
  - `expiration_policy` - (String) The policy to apply when the committed use term expires.

    ->**expiration_policy** 
      </br>&#x2022; release: Release any available capacity and let the reservation expire
      </br>&#x2022; renew: Renew for another term, provided the term remains listed in the reservation_terms for the profile. Otherwise, let the reservation expire
  - `term` - (String) The term for this committed use reservation.

    ->**term** 
      </br>&#x2022; one_year: 1 year
      </br>&#x2022; three_year: 3 years
- `created_at` - (Timestamp) The date and time that the reservation was created.
- `crn` - (String) The CRN for this reservation.
- `href` - (String) The URL for this reservation.
- `id` - (String) The unique identifier for this reservation.
- `lifecycle_state` - (String) The lifecycle state of this reservation.

   ->**lifecycle_state** 
      </br>&#x2022; deleting
      </br>&#x2022; failed
      </br>&#x2022; pending
      </br>&#x2022; stable
      </br>&#x2022; suspended
      </br>&#x2022; updating
      </br>&#x2022; waiting
- `profile` - (List) The virtual server instance profile this reservation. 

  Nested scheme for `profile`:
  - `href` - (String) The URL for this virtual server instance profile.
  - `name` - (String) The globally unique name for this virtual server instance profile.
  - `resource_type` - (string) The resource type

     ->**resource_type** 
      </br>&#x2022; instance_profile
- `resource_group` - (List) The resource group for this reservation. 

  Nested scheme for `resource_group`:
  - `href` - (String) The URL for this resource group.
  - `id` - (String) The unique identifier for this resource group.
  - `name` - (String) The name for this resource group.
- `resource_type` - (String) The resource type.

  ->**resource_type** 
    </br>&#x2022; reservation
- `status` - (String) The status of the reservation.

  ->**status** 
    </br>&#x2022; activating
    </br>&#x2022; active
    </br>&#x2022; deactivating
    </br>&#x2022; expired
    </br>&#x2022; failed
    </br>&#x2022; inactive
- `status_reasons` - (List) The reasons for the current status (if any).

  Nested scheme for `status_reasons`:
  - `code` - (String) A snake case string succinctly identifying the status reason.
  
    ->**code** 
      </br>&#x2022; cannot_activate_no_capacity_available
      </br>&#x2022; cannot_renew_unsupported_profile_term
  - `message` - (String) An explanation of the status reason.
  - `more_info` - (string) Link to documentation about this status reason
- `zone` - (String) The globally unique name for this zone.

## Import
The `ibm_is_reservation` resource can be imported by using the ID. 

**Syntax**

```
$ terraform import ibm_is_reservation.example <reservation_ID>
```

**Example**

```
$ terraform import ibm_is_reservation.example d7bec597-4726-451f-8a63-e62e6f12122c
```

## References 

* [IBM Cloud Terraform Docs](https://cloud.ibm.com/docs/vpc?topic=vpc-provisioning-reserved-capacity-vpc&interface=ui
https://cloud.ibm.com/docs/vpc?topic=vpc-about-reserved-virtual-servers-vpc)