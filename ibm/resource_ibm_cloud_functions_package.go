package ibm

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/apache/incubator-openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIBMCloudFunctionsPackage() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMCloudFunctionsPackageCreate,
		Read:     resourceIBMCloudFunctionsPackageRead,
		Update:   resourceIBMCloudFunctionsPackageUpdate,
		Delete:   resourceIBMCloudFunctionsPackageDelete,
		Exists:   resourceIBMCloudFunctionsPackageExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Name of package.",
				ValidateFunc: validateCloudFunctionsName,
			},
			"publish": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Package visibilty.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Semantic version of the item.",
			},
			"user_defined_annotations": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Annotation values in KEY VALUE format.",
				Default:          "[]",
				ValidateFunc:     validateJSONString,
				DiffSuppressFunc: suppressEquivalentJSON,
				StateFunc: func(v interface{}) string {
					json, _ := normalizeJSONString(v)
					return json
				},
			},
			"user_defined_parameters": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Parameters values in KEY VALUE format. Parameter bindings included in the context passed to the package.",
				ValidateFunc:     validateJSONString,
				Default:          "[]",
				DiffSuppressFunc: suppressEquivalentJSON,
				StateFunc: func(v interface{}) string {
					json, _ := normalizeJSONString(v)
					return json
				},
			},
			"annotations": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "All annotations set on package by user and those set by the IBM Cloud Function backend/API.",
			},
			"parameters": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "All parameters set on package by user and those set by the IBM Cloud Function backend/API.",
			},
			"bind_package_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Description:  "Name of package to be binded.",
				ValidateFunc: validateBindedPackageName,
				DiffSuppressFunc: func(k, o, n string, d *schema.ResourceData) bool {
					if o == "" {
						return false
					}
					if strings.HasPrefix(n, "/_") {
						temp := strings.Replace(n, "/_", "/"+os.Getenv("CLOUD_FUNCTIONS_NAMESPACE"), 1)
						if strings.Compare(temp, o) == 0 {
							return true
						}
					}
					return false
				},
			},
		},
	}
}

func resourceIBMCloudFunctionsPackageCreate(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).CloudFunctionsClient()
	if err != nil {
		return err
	}
	packageService := wskClient.Packages

	name := d.Get("name").(string)

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(name); err != nil {
		return NewQualifiedNameError(name, err)
	}

	payload := whisk.Package{
		Name:      qualifiedName.GetEntityName(),
		Namespace: qualifiedName.GetNamespace(),
	}

	userDefinedAnnotations := d.Get("user_defined_annotations").(string)
	payload.Annotations, err = expandAnnotations(userDefinedAnnotations)
	if err != nil {
		return err
	}

	userDefinedParameters := d.Get("user_defined_parameters").(string)
	payload.Parameters, err = expandParameters(userDefinedParameters)
	if err != nil {
		return err
	}

	if publish, ok := d.GetOk("publish"); ok {
		p := publish.(bool)
		payload.Publish = &p
	}

	if v, ok := d.GetOk("bind_package_name"); ok {
		var BindingQualifiedName = new(QualifiedName)
		if BindingQualifiedName, err = NewQualifiedName(v.(string)); err != nil {
			return NewQualifiedNameError(v.(string), err)
		}
		BindingPayload := whisk.Binding{
			Name:      BindingQualifiedName.GetEntityName(),
			Namespace: BindingQualifiedName.GetNamespace(),
		}
		payload.Binding = &BindingPayload
	}
	log.Println("[INFO] Creating IBM CLoud Functions package")
	result, _, err := packageService.Insert(&payload, false)
	if err != nil {
		return fmt.Errorf("Error creating IBM CLoud Functions package: %s", err)
	}

	d.SetId(result.Name)

	return resourceIBMCloudFunctionsPackageRead(d, meta)
}

func resourceIBMCloudFunctionsPackageRead(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).CloudFunctionsClient()
	if err != nil {
		return err
	}
	packageService := wskClient.Packages
	id := d.Id()

	pkg, _, err := packageService.Get(id)
	if err != nil {
		return fmt.Errorf("Error retrieving IBM Cloud Functions package %s : %s", id, err)
	}

	d.SetId(pkg.Name)
	d.Set("name", pkg.Name)
	d.Set("publish", pkg.Publish)
	d.Set("version", pkg.Version)
	annotations, err := flattenAnnotations(pkg.Annotations)
	if err != nil {
		return err
	}
	d.Set("annotations", annotations)
	parameters, err := flattenParameters(pkg.Parameters)
	if err != nil {
		return err
	}
	d.Set("parameters", parameters)
	if isEmpty(*pkg.Binding) {

		d.Set("user_defined_annotations", annotations)
		d.Set("user_defined_parameters", parameters)

	} else {
		d.Set("bind_package_name", fmt.Sprintf("/%s/%s", pkg.Binding.Namespace, pkg.Binding.Name))

		c, err := whisk.NewClient(http.DefaultClient, &whisk.Config{
			Namespace: pkg.Binding.Namespace,
			AuthToken: wskClient.AuthToken,
			Host:      wskClient.Host,
		})
		bindedPkg, _, err := c.Packages.Get(pkg.Binding.Name)

		if err != nil {
			return fmt.Errorf("Error retrieving Binded IBM Cloud Functions package %s : %s", pkg.Binding.Name, err)
		}

		userAnnotations, err := flattenAnnotations(filterInheritedAnnotations(bindedPkg.Annotations, pkg.Annotations))
		if err != nil {
			return err
		}
		d.Set("user_defined_annotations", userAnnotations)

		userParameters, err := flattenParameters(filterInheritedParameters(bindedPkg.Parameters, pkg.Parameters))
		if err != nil {
			return err
		}
		d.Set("user_defined_parameters", userParameters)
	}
	return nil
}

func resourceIBMCloudFunctionsPackageUpdate(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).CloudFunctionsClient()
	if err != nil {
		return err
	}
	packageService := wskClient.Packages

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(d.Get("name").(string)); err != nil {
		return NewQualifiedNameError(d.Get("name").(string), err)
	}

	payload := whisk.Package{
		Name:      qualifiedName.GetEntityName(),
		Namespace: qualifiedName.GetNamespace(),
	}
	ischanged := false
	if d.HasChange("publish") {
		p := d.Get("publish").(bool)
		payload.Publish = &p
		ischanged = true
	}

	if d.HasChange("user_defined_parameters") {
		var err error
		payload.Parameters, err = expandParameters(d.Get("user_defined_parameters").(string))
		if err != nil {
			return err
		}
		ischanged = true
	}

	if d.HasChange("user_defined_annotations") {
		var err error
		payload.Annotations, err = expandAnnotations(d.Get("user_defined_annotations").(string))
		if err != nil {
			return err
		}
		ischanged = true
	}

	if ischanged {
		log.Println("[INFO] Update IBM Cloud Functions Package")
		_, _, err = packageService.Insert(&payload, true)
		if err != nil {
			return fmt.Errorf("Error updating IBM Cloud Functions Package: %s", err)
		}
	}

	return resourceIBMCloudFunctionsPackageRead(d, meta)
}

func resourceIBMCloudFunctionsPackageDelete(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).CloudFunctionsClient()
	if err != nil {
		return err
	}
	packageService := wskClient.Packages
	id := d.Id()

	_, err = packageService.Delete(id)
	if err != nil {
		return fmt.Errorf("Error deleting IBM Cloud Functions Package: %s", err)
	}

	d.SetId("")
	return nil
}

func resourceIBMCloudFunctionsPackageExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	wskClient, err := meta.(ClientSession).CloudFunctionsClient()
	if err != nil {
		return false, err
	}
	packageService := wskClient.Packages
	id := d.Id()

	pkg, resp, err := packageService.Get(id)
	if err != nil {
		if resp.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error communicating with IBM Cloud Functions Client : %s", err)
	}
	return pkg.Name == id, nil
}
