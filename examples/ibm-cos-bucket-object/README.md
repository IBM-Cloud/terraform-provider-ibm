# Example for IBM Cloud Object Storage - Objects

This example illustrates how to use IBM Cloud Object Storage to create objects in a bucket.

These types of resources are supported:

* ibm_cos_bucket_object

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
$ terraform destroy
```

## COS object resource

ibm_cos_bucket_object resource:

```hcl
resource "ibm_cos_bucket_object" "object" {
  bucket_crn      = var.cos_bucket_crn
  bucket_location = var.cos_bucket_location
  content         = "Hello World"
  key             = "file.txt"
}
```

## COS object data source

ibm_cos_bucket_object data source:

```hcl
data "ibm_cos_bucket_object" "object" {
  bucket_crn      = var.cos_bucket_crn
  bucket_location = var.cos_bucket_location
  key             = "file.txt"
}
```

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | TBD |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| bucket\_crn | The CRN of the COS bucket. | `string` | true |
| bucket\_location | The location of the COS bucket. | `string` | true |
| content | Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text. Conflicts with `content_base64` and `content_file`. | `string` | false |
| content\_base64 | Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content. Conflicts with `content` and `content_file`. | `string` | false |
| content\_file | The path to a file that will be read and uploaded as raw bytes for the object content. Conflicts with `content` and `content_base64`. | `string` | false |
| endpoint\_type | The type of endpoint used to access COS. Valid options are "public", "private", and "direct". Defaults to "public". | `string` | false |
| etag | MD5 hexdigest used to trigger updates. The only meaningful value is `filemd5("path/to/file")`. | `string` | false |
| key | The name of the object in the COS bucket. | `string` | true |

## Outputs

| Name | Description | Type |
|------|-------------|------|
| body | Literal string value of the object content. Only supported for `text/*` and `application/json` content types. | `string` |
| content\_length | Size of the object body, in bytes. | `integer` |
| content\_type | A standard MIME type describing the format of the object data. | `string` |
| etag | Computed MD5 hexdigest of the object content. | `string` |
| last\_modified | Last modified date of the object. A GMT formatted date. | `string` |