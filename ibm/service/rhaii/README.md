# Terraform IBM Provider Red Hat AI Inference
<!-- markdownlint-disable MD026 -->
This area is primarily for IBM provider contributors and maintainers. For information on _using_ Terraform and the IBM provider, see the links below.

## What is in this package

| Name | Type | API used |
|------|------|----------|
| `ibm_rhaii_project` | Resource and data source | Global Catalog, Resource Controller and Global Tagging |
| `ibm_rhaii_inference_models` | Data source | Red Hat AI Inference API (`rhaiiv1`) |
| `ibm_rhaii_inference_model` | Data source | Red Hat AI Inference API (`rhaiiv1`) |

All four are implemented with the Terraform Plugin Framework and registered in
`ibm/provider_framework/provider.go`. The framework provider is muxed with the
SDKv2 provider in `main.go`, so users see one `ibm` provider. A type name must be
registered in only one of the two providers; `TestRhaiiMuxRegistration` checks it.

Mapping of the SDKv2 patterns used by the other resources:

| SDKv2 | Here |
|-------|------|
| `Default` + `ForceNew` | `stringdefault.StaticString` + `stringplanmodifier.RequiresReplace()` |
| Optional + Computed tags (unset keeps server tags, `[]` detaches) | Optional + Computed + `setplanmodifier.UseStateForUnknown()` |
| `flex.ResourceTagsCustomizeDiff` (`IC_ENV_TAGS`) | `ModifyPlan` adds the env tags to the planned tags |
| `flex.ResourceValidateAccessTags` | `ModifyPlan` checks the access tags exist (`rhaiiMissingAccessTags`) |
| validator dictionary (`validate.InvokeValidator`) | `terraform-plugin-framework-validators` |
| `timeouts` block | `terraform-plugin-framework-timeouts`, as a block, so the HCL is the same |

Two SDKv2 packages are still imported on purpose: `helper/retry` for the wait
loops, and `helper/schema` only to pass sets to the shared `flex` tag helpers.

A project is a resource controller instance of the `instructlab` service. The
project resource has its own create, read, update and delete logic and does not
call into the `resourcecontroller` package, so changes to `ibm_resource_instance`
do not affect it.

`rhaiiv1` is a small embedded client for the Red Hat AI Inference API. The public
spec marks the inference operations with `x-sdk-exclude`, so no published Go SDK
covers them. The client follows the layout of the generated IBM Cloud Go SDKs and
is created in `ibm/conns/config.go` with the shared provider authenticator,
retries and headers. Replace it with the official SDK if one is published.

The API endpoint follows the provider `visibility` setting: public
(`https://<region>.rhai.ibm.com/v1`), private (`https://private.<region>.rhai.ibm.com/v1`)
or public-and-private (private when available). The rules are in
`rhaiiv1.GetServiceURLForVisibility`, which both `ibm/conns/config.go` and the
data sources use. The endpoint can be overridden with `IBMCLOUD_RHAII_API_ENDPOINT`
(environment variable or endpoints file).

## Tests

* Unit tests (no cloud access): `go test ./ibm/service/rhaii/... -run 'TestRhaii|TestFlatten|TestList|TestGet|TestOptions|TestService|TestClone'`
  * `TestRhaiiProjectPlan*` and `TestRhaiiProjectValidateConfig` run plan and validation through the real mux server, without credentials
* Acceptance tests use `acc.TestAccProtoV6ProviderFactories()`: `TF_ACC=1 IC_API_KEY=<key> go test ./ibm/service/rhaii/ -run TestAcc -timeout 60m`
  * `TestAccIBMRhaiiProject_accessTags` runs only when `IBM_RHAII_ACCESS_TAG` is set to an existing access management tag.

## Handy Links
* [Find out about contributing](../../../CONTRIBUTING.md) to the IBM provider!
* IBM Provider Docs: [Home](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs)
* IBM API Docs: [IBM API Docs for Red Hat AI Inference](https://cloud.ibm.com/docs/apis/inference)
* IBM Docs: [Red Hat AI Inference on IBM Cloud](https://cloud.ibm.com/docs/inference)
