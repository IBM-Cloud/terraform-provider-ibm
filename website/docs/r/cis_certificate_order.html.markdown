---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_certificate_order"
description: |-
  Provides a IBM CIS certificate order resource.
---

# ibm_cis_certificate_order

 Provides an IBM Cloud Internet Services certificate order resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS domain resource. It allows to order and delete dedicated certificates of a domain of a CIS instance. For more information about CIS certificate order, see [managing origin certificates](https://cloud.ibm.com/docs/cis?topic=cis-cis-origin-certificates).

## Example usage

```terraform
resource "ibm_cis_certificate_order" "test" {
	cis_id    = data.ibm_cis.cis.id
	domain_id = data.ibm_cis_domain.cis_domain.domain_id
	hosts     = ["example.com"]
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain.
- `hosts` - (Required, String) The hosts for the certificates to be ordered.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `certificate_id`- (String) The certificate ID.
- `id` - (String) The record ID. It is a combination of `<certificate_id>,<domain_id>,<cis_id>` attributes concatenated with `:`.
- `status`- (String) The certificate status.

## Import
The `ibm_cis_certificate_order` resource can be imported using the ID. The ID is formed from the certificate ID, the domain ID of the domain and the CRN  Concatenated  by using a `:` character.

The domain ID and CRN is located on the **Overview** page of the IBM Cloud Internet Services instance of the console domain heading, or by using the `ibmcloud cis` command line commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **Certificate ID** is a 32 digit character string of the form: `489d96f0da6ed76251b475971b097205c`.


**Syntax**

```
$ terraform import ibm_cis_certificate_order.myorg <certificate_id>:<domain-id>:<crn>
```


**Example**

```
$ terraform import ibm_cis_certificate_order.myorg certificate_order 48996f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
