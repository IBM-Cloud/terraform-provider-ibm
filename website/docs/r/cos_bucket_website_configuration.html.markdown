---

subcategory: "Object Storage"
layout: "ibm"
page_title: "IBM : Cloud Object Storage Static Web Hosting"
description: 
  "Manages IBM Cloud Object Storage Static Web Hosting"
---

# ibm_cos_bucket_website_configuration
Provides an  Static web hosting configuration resource. This resource is used to  configure the website to use your documents as an index for the site and to potentially display errors.It can also be used to configure more advanced options including routing rules and request redirect for your domain.For more information about static web hosting please refer [Static web hsoting with COS](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-static-website-options) To configure a website_redirect for an object please refer [ibm_cos_bucket_object](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/cm_object).



**Note:**
Enabling public access to the COS bucket required to access the hosted website.Granting public access to this bucket will allow anyone to access the bucket.

---

## Example usage
The following example demonstrates creating a bucket and adding the website configuration to host a static website.

```terraform
data "ibm_resource_group" "cos_group" {
  name = "cos-resource-group"
}

resource "ibm_resource_instance" "cos_instance" {
  name              = "cos-instance"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "cloud-object-storage"
  plan              = "standard"
  location          = "global"
}

resource "ibm_cos_bucket" "cos_bucket_website_configuration" {
  bucket_name           = var.bucket_name
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = var.regional_loc
  storage_class         = var.standard_storage_class

}

data "ibm_iam_access_group" "public_access_group" {
  access_group_name = "Public Access"
}

# Give public access to above mentioned bucket
 
resource "ibm_iam_access_group_policy" "policy" { 
  depends_on = [ibm_cos_bucket.cos_bucket_website_configuration] 
  access_group_id = data.ibm_iam_access_group.public_access_group.groups[0].id 
  roles = ["Object Reader"] 

  resources { 
    service = "cloud-object-storage" 
    resource_type = "bucket" 
    resource_instance_id = "COS instance guid"  # eg : 94xxxxxx-3xxx-4xxx-8xxx-7xxxxxxxxx7
    resource = ibm_cos_bucket.cos_bucket_website_configuration.bucket_name
  } 
} 

resource ibm_cos_bucket_website_configuration "website" {
  bucket_crn = "bucket_crn"
  bucket_location = data.ibm_cos_bucket.cos_bucket_website_configuration.regional_location
  website_configuration {
    error_document{
      key = "error.html"
    }
    index_document{
      suffix = "index.html"
    }
  }
}

```
# Adding website configuration on a bucket with redirect requests.
All the requests can be redirected to a specific host.To redirect the request , first create a COS bucket , grant public access to the bucket and  use the `redirect_all_requests_to` argument as shown in the example:

## Example usage

```terraform
// To redirect all the incoming redirect requests

resource ibm_cos_bucket_website_configuration "website" {
  bucket_crn = "bucket_crn"
  bucket_location = data.ibm_cos_bucket.cos_bucket_website_configuration.regional_location
  website_configuration {
      redirect_all_requests_to{
			host_name = "exampleBucketName" or "www.domain.com"
			protocol = "https"
		}

  }
}
```

# Adding website configuration with routing_rule configured.
Routing rules define the individual rules that process incoming requests for specific pages. To configure routinge rule, first create a COS bucket , grant public access to the bucket and  use the `routing_rule` argument as shown in the example: 

## Example usage

```terraform
// To redirect all the incoming redirect requests

resource ibm_cos_bucket_website_configuration "website" {
  bucket_crn = "bucket_crn"
  bucket_location = data.ibm_cos_bucket.cos_bucket_website_configuration.regional_location
 website_configuration {
    error_document{
      key = "error.html"
    }
    index_document{
      suffix = "index.html"
    }
   routing_rule {
    condition {
      key_prefix_equals = "pages/"
    }
    redirect {
      replace_key_prefix_with = "web_pages/"
    }
  }
  }
}
```


# Adding website configuration with routing_rules configured.
Routing rules can also be configured using the `routing_rules` argument as shown in the example:

## Example usage

```terraform
// To redirect all the incoming redirect requests

resource ibm_cos_bucket_website_configuration "website" {
  bucket_crn = "bucket_crn"
  bucket_location = data.ibm_cos_bucket.cos_bucket_website_configuration.regional_location
 website_configuration {
    error_document{
      key = "error.html"
    }
    index_document{
      suffix = "index.html"
    }
    routing_rules = <<EOF
    [{
        "Condition": {
            "KeyPrefixEquals": "pages/"
        },
        "Redirect": {
            "ReplaceKeyPrefixWith": "web_pages/"
        }
    }]
    EOF
  }
}
```


## Argument reference
Review the argument references that you can specify for your resource. 
- `bucket_crn` - (Required, Forces new resource, String) The CRN of the COS bucket.
- `bucket_location` - (Required, Forces new resource, String) The location of the COS bucket.
- `endpoint_type`- (Optional, String) The type of the endpoint either `public` or `private` or `direct` to be used for buckets. Default value is `public`.
- `website_configuration`- (Required, List) Nested block have the following structure:
  
  Nested scheme for `website_configuration`:
  - `error_document`- (Optional , Conflicts with `redirect_all_requests_to`) When a static website bucket error occurs, an HTML page of the error provided by the configured `error_document` will be returned.
  
  Nested scheme for `error_document`:
  - `key`- (Required,String) Object key name to use when a 4XX class error occurs.

  - `index_document`- (Optional , Required if `redirect_all_requests_to` is not specified) This is the home or default page of the website.
  
  Nested scheme for `index_document`:
  - `suffix`- (Required,String)- This suffix is the document that will be returned for requests made to the root of the website .For example, if the suffix is `index.html` and you make a request to `mywebsitedomain.s3-web.us-east.cloud-object-storage.appdomain.cloud`, the data that is returned will be for `mywebsitedomain.s3-web.us-east.cloud-object-storage.appdomain.cloud/index.html`.
  **Note:**
    The suffix must not be empty and must not include a slash character.

  - `redirec_all_requests_to`- (Optional , Required if `index_document` is not specified) Specifies the redirect behavior for every request to the bucket's website endpoint.Conflicts with `error_document`, `index_document`, and `routing_rule`.
  
  Nested scheme for `redirec_all_requests_to`:
   - `host_name`- (Required,String) Name of the host where requests are redirected.
   - `protocol`- (Optional,String) Protocol to use when redirecting requests. The default is the protocol that is used in the original request. Valid values: http, https.

  - `routing_rule`- (Optional , Conflicts with `redirect_all_requests_to` is not specified) List of rules that define when a redirect is applied and the redirect behavior.


  Nested scheme for `routing_rule`:
   - `condition`- (Optional) Configuration block for a condition to be satisfied for the redirect behaviour to be applied.


  Nested scheme for `condition`:

   - `http_error_code_returned_equals` - (Optional, Required if key_prefix_equals is not specified) HTTP error code when the redirect is applied. If specified with `key_prefix_equals`, then both must be true for the redirect to be applied.
   - `key_prefix_equals` - (Optional, Required if http_error_code_returned_equals is not specified) Object key name prefix when the redirect is applied. If specified with `http_error_code_returned_equals`, then both must be true for the redirect to be applied.
  
   - `redirect`- (Required) Configuration block for redirect behaviour.
  Nested scheme for `redirect`:
   - `host_name` - (Optional) Host name to use in the redirect request.
   - `http_redirect_code` - (Optional) HTTP redirect code to use on the response.
   - `protocol` - (Optional) Protocol to use when redirecting requests. The default is the protocol that is used in the original request. Valid values: http, https.
   - `replace_key_prefix_with` - (Optional, Conflicts with `replace_key_with`) Object key prefix to use in the redirect request. For example, to redirect requests for all pages with prefix docs/ (objects in the docs/ folder) to documents/, you can set a condition block with key_prefix_equals set to docs/ and in the redirect set replace_key_prefix_with to /documents.
  - `replace_key_with` - (Optional, Conflicts with `replace_key_prefix_with`) Specific object key to use in the redirect request. For example, redirect request to error.html.

  - `routing_rules` - (Optional, Conflicts with `routing_rule` and `redirect_all_requests_to`) JSON array containing routing rules describing redirect behavior and when redirects are applied. Use this parameter when your routing rules contain empty String values ("") as seen in the example above.
  
**Note:**
 There is a limitation of 50 routing rules per website configuration.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN of the bucket.
- `id` - (String) The ID of the bucket.
- `website_endpoint` - (String) Endpoint of the website.

## Import IBM COS Bucket
The `ibm_cos_bucket_website_configuration` resource can be imported by using the `id`. The ID is formed from the `CRN` (Cloud Resource Name). The `CRN` and bucket location can be found on the portal.

id = `$CRN:meta:$bucketlocation:$endpointtype`

**Syntax**

```
$ terraform import ibm_cos_bucket_website_configuration.website  `$CRN:meta:$bucketlocation:public`

```

**Example**

```

$ terraform import ibm_cos_bucket_website_configuration.website crn:v1:bluemix:public:cloud-object-storage:global:a/ee858e45752d4696b2d082bcf2357559:84aaaaa4-3a22-477b-8635-75501eac96f7:bucket:bucketname:meta:us-south:public

```