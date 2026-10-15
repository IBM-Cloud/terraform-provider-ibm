// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/rhaii/rhaiiv1"
	"github.com/IBM/platform-services-go-sdk/globalcatalogv1"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

const (
	// rhaiiServiceName is the global catalog name (and ID) of the Red Hat AI Inference service.
	rhaiiServiceName = "instructlab"
	// rhaiiDefaultPlan is the only plan that the service offers today.
	rhaiiDefaultPlan = "instructlab-pricing-plan"
	// rhaiiDefaultLocation is the only region that the service is deployed to today.
	rhaiiDefaultLocation = "us-east"

	// rhaiiEndpointEnv overrides the RHAII API URL, as an environment variable
	// or as a key of the endpoints file.
	rhaiiEndpointEnv = "IBMCLOUD_RHAII_API_ENDPOINT"

	// Tag types of the Global Tagging API.
	rhaiiUserTagType   = "user"
	rhaiiAccessTagType = "access"

	// Resource controller instance states.
	rhaiiStateActive             = "active"
	rhaiiStateFailed             = "failed"
	rhaiiStateRemoved            = "removed"
	rhaiiStatePendingReclamation = "pending_reclamation"
	rhaiiStateInProgress         = "in progress"
	rhaiiStateSucceeded          = "succeeded"

	// Synthetic states used by the wait functions.
	rhaiiWaitCreating = "creating"
	rhaiiWaitUpdating = "updating"
	rhaiiWaitDeleting = "deleting"
	rhaiiWaitDone     = "done"
)

// rhaiiProjectEndpoint builds the public Red Hat AI Inference API base URL for a project.
func rhaiiProjectEndpoint(location, projectID string) string {
	serviceURL, err := rhaiiv1.GetServiceURLForRegion(location)
	if err != nil {
		// Region not known to rhaiiv1 yet: use the host pattern of the known regions.
		serviceURL = fmt.Sprintf("https://%s.rhai.ibm.com/v1", location)
	}
	return rhaiiv1.ProjectServiceURL(serviceURL, projectID)
}

// rhaiiProjectPrivateEndpoint builds the private (service endpoint) Red Hat AI
// Inference API base URL for a project.
func rhaiiProjectPrivateEndpoint(location, projectID string) string {
	serviceURL, err := rhaiiv1.GetServiceURLForRegion("private." + location)
	if err != nil {
		// Region not known to rhaiiv1 yet: use the host pattern of the known regions.
		serviceURL = fmt.Sprintf("https://private.%s.rhai.ibm.com/v1", location)
	}
	return rhaiiv1.ProjectServiceURL(serviceURL, projectID)
}

// rhaiiResolvePlanID finds the catalog ID of a plan of the RHAII service by plan
// name. Only names are matched, like SERVICE_PLAN_NAME in
// "ibmcloud resource service-instance-create".
func rhaiiResolvePlanID(ctx context.Context, gcClient *globalcatalogv1.GlobalCatalogV1, planName string) (string, error) {
	plans, _, err := gcClient.GetChildObjectsWithContext(ctx, &globalcatalogv1.GetChildObjectsOptions{
		ID:   ptr(rhaiiServiceName),
		Kind: ptr("plan"),
	})
	if err != nil {
		return "", fmt.Errorf("error retrieving plans of service %q from the global catalog: %s", rhaiiServiceName, err)
	}

	available := []string{}
	for _, p := range plans.Resources {
		if p.ID == nil || p.Name == nil {
			continue
		}
		if *p.Name == planName {
			return *p.ID, nil
		}
		if *p.ID == planName {
			return "", fmt.Errorf("plan_name %q is a plan ID; set the plan name %q instead", planName, *p.Name)
		}
		available = append(available, *p.Name)
	}
	sort.Strings(available)
	return "", fmt.Errorf("plan %q not found for service %q. Valid plan names are: %q", planName, rhaiiServiceName, available)
}

// rhaiiResolveTargetCRN finds the deployment CRN of a plan in a location. This
// is the "target" of the resource controller create request.
func rhaiiResolveTargetCRN(ctx context.Context, gcClient *globalcatalogv1.GlobalCatalogV1, planID, location string) (string, error) {
	deployments, _, err := gcClient.GetChildObjectsWithContext(ctx, &globalcatalogv1.GetChildObjectsOptions{
		ID:      ptr(planID),
		Kind:    ptr("deployment"),
		Include: ptr("metadata"),
	})
	if err != nil {
		return "", fmt.Errorf("error retrieving deployments of plan %q from the global catalog: %s", planID, err)
	}

	locations := map[string]bool{}
	for _, dep := range deployments.Resources {
		if dep.Metadata == nil || dep.Metadata.Deployment == nil || dep.Metadata.Deployment.Location == nil || dep.CatalogCRN == nil {
			continue
		}
		if dep.Metadata.RcCompatible != nil && !*dep.Metadata.RcCompatible {
			continue
		}
		depLocation := *dep.Metadata.Deployment.Location
		if depLocation == location {
			return *dep.CatalogCRN, nil
		}
		locations[depLocation] = true
	}

	available := make([]string, 0, len(locations))
	for l := range locations {
		available = append(available, l)
	}
	sort.Strings(available)
	return "", fmt.Errorf("service %q is not available in location %q. Valid locations are: %q", rhaiiServiceName, location, available)
}

// rhaiiPlanName returns the catalog name of a plan.
func rhaiiPlanName(ctx context.Context, gcClient *globalcatalogv1.GlobalCatalogV1, planID string) (string, error) {
	entry, _, err := gcClient.GetCatalogEntryWithContext(ctx, &globalcatalogv1.GetCatalogEntryOptions{
		ID: ptr(planID),
	})
	if err != nil {
		return "", fmt.Errorf("error retrieving plan %q from the global catalog: %s", planID, err)
	}
	if entry.Name == nil {
		return "", fmt.Errorf("plan %q has no name in the global catalog", planID)
	}
	return *entry.Name, nil
}

// isRhaiiInstance reports whether a resource instance is an RHAII project.
func isRhaiiInstance(instance *rc.ResourceInstance) bool {
	return instance != nil && instance.ResourceID != nil && *instance.ResourceID == rhaiiServiceName
}

// isRhaiiInstanceGone reports whether a resource instance is deleted or waiting for reclamation.
func isRhaiiInstanceGone(instance *rc.ResourceInstance) bool {
	if instance == nil || instance.State == nil {
		return false
	}
	return *instance.State == rhaiiStateRemoved || *instance.State == rhaiiStatePendingReclamation
}

// waitForRhaiiProject polls a resource instance until refresh reports rhaiiWaitDone.
func waitForRhaiiProject(ctx context.Context, rcClient *rc.ResourceControllerV2, id string, pending string, timeout time.Duration, refresh func(*rc.ResourceInstance) (string, error)) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{pending},
		Target:  []string{rhaiiWaitDone},
		Refresh: func() (interface{}, string, error) {
			instance, resp, err := rcClient.GetResourceInstanceWithContext(ctx, &rc.GetResourceInstanceOptions{ID: &id})
			if err != nil {
				if resp != nil && resp.StatusCode == 404 && pending == rhaiiWaitDeleting {
					return struct{}{}, rhaiiWaitDone, nil
				}
				return nil, "", fmt.Errorf("error getting resource instance %s: %s", id, err)
			}
			state, err := refresh(instance)
			return instance, state, err
		},
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	_, err := stateConf.WaitForStateContext(ctx)
	return err
}

func rhaiiCreateRefresh(instance *rc.ResourceInstance) (string, error) {
	switch state := strDeref(instance.State); state {
	case rhaiiStateActive:
		return rhaiiWaitDone, nil
	case rhaiiStateFailed, rhaiiStateRemoved, rhaiiStatePendingReclamation:
		return "", fmt.Errorf("project %s ended in state %q", strDeref(instance.ID), state)
	default:
		return rhaiiWaitCreating, nil
	}
}

func rhaiiUpdateRefresh(instance *rc.ResourceInstance) (string, error) {
	if instance.LastOperation == nil {
		return rhaiiWaitDone, nil
	}
	switch state := strDeref(instance.LastOperation.State); state {
	case rhaiiStateInProgress:
		return rhaiiWaitUpdating, nil
	case rhaiiStateFailed:
		return "", fmt.Errorf("update of project %s failed: %s", strDeref(instance.ID), strDeref(instance.LastOperation.Description))
	default:
		return rhaiiWaitDone, nil
	}
}

func rhaiiDeleteRefresh(instance *rc.ResourceInstance) (string, error) {
	if isRhaiiInstanceGone(instance) {
		return rhaiiWaitDone, nil
	}
	if strDeref(instance.State) == rhaiiStateFailed {
		return "", fmt.Errorf("deletion of project %s failed", strDeref(instance.ID))
	}
	return rhaiiWaitDeleting, nil
}

// rhaiiNextStart extracts the "start" token from a resource controller next_url.
func rhaiiNextStart(nextURL *string) (string, error) {
	if nextURL == nil || *nextURL == "" {
		return "", nil
	}
	u, err := url.Parse(*nextURL)
	if err != nil {
		return "", err
	}
	return u.Query().Get("start"), nil
}

// rhaiiToJSON marshals free form API data into a JSON string attribute.
func rhaiiToJSON(v map[string]interface{}) (string, error) {
	if len(v) == 0 {
		return "", nil
	}
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// rhaiiClientForLocation returns an RHAII API client. When location is empty,
// the client from the provider configuration is used as is. When location is
// set, the URL is built for that region with the same rules as the provider
// configuration (ibm/conns/config.go), in this order:
//  1. the IBMCLOUD_RHAII_API_ENDPOINT environment variable
//  2. the IBMCLOUD_RHAII_API_ENDPOINT key of the endpoints file (not for public-and-private)
//  3. the public or private endpoint of the region, based on the provider visibility
func rhaiiClientForLocation(meta interface{}, location string) (*rhaiiv1.RhaiiV1, error) {
	client, err := meta.(conns.ClientSession).RhaiiV1()
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(location) == "" {
		return client, nil
	}

	sess, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		return nil, err
	}
	serviceURL, err := rhaiiServiceURLForLocation(location, sess.Config.Visibility, sess.Config.EndpointsFile)
	if err != nil {
		return nil, err
	}
	if err := client.SetServiceURL(serviceURL); err != nil {
		return nil, err
	}
	return client, nil
}

// rhaiiServiceURLForLocation resolves the RHAII API URL of a region. See rhaiiClientForLocation.
func rhaiiServiceURLForLocation(location, visibility, endpointsFile string) (string, error) {
	if v := os.Getenv(rhaiiEndpointEnv); v != "" {
		return v, nil
	}
	serviceURL, urlErr := rhaiiv1.GetServiceURLForVisibility(location, visibility)
	if visibility != "public-and-private" {
		// FileFallBack also reads IBMCLOUD_ENDPOINTS_FILE_PATH and returns serviceURL when no entry matches.
		serviceURL = conns.FileFallBack(endpointsFile, visibility, rhaiiEndpointEnv, location, serviceURL)
	}
	if serviceURL == "" {
		return "", fmt.Errorf("no %s endpoint of the Red Hat AI Inference API for location %q: %v", visibility, location, urlErr)
	}
	return serviceURL, nil
}

func ptr(s string) *string { return &s }

func strDeref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
