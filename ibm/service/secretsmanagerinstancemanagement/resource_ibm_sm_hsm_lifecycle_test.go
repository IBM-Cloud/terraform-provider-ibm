// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanagerinstancemanagement_test

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

// The HSM lifecycle is a chain of top-level tests, one per step, that run in
// source order and share one destination. If a step fails, the later steps are
// skipped and the last test still deletes the destination.
//
// Required environment variables (the tests skip when any is missing):
//
//	SECRETS_MANAGER_DEDICATED_INSTANCE_ID (or SECRETS_MANAGER_INSTANCE_ID)
//	SM_HSM_PPS_GATEWAY_CRN       CRN of the Private Path service gateway
//	SM_HSM_SERVER_CA_PEM_PATH    Server CA certificate PEM file
//	SM_HSM_CLIENT_CA_CERT_PATH   local CA certificate used to sign the client CSR
//	SM_HSM_CLIENT_CA_KEY_PATH    local CA private key used to sign the client CSR
//	SM_HSM_MEMBER_SERIALS        Member serial numbers (comma-separated, e.g. "1323465009870,1323465009871")
//
// Optional: SM_HSM_APPROVAL_TIMEOUT_MINUTES (default 30). While the path step
// waits, approve the VPE connection in the IBM Cloud console (VPC
// Infrastructure > Private Path service gateways).
var (
	hsmLifecycle       *hsmFlow
	hsmLifecycleBroken bool
)

func newHsmFlow(t *testing.T) *hsmFlow {
	instanceID := acc.SecretsManagerDedicatedInstanceID
	if instanceID == "" {
		instanceID = acc.SecretsManagerInstanceID
	}
	f := &hsmFlow{
		instanceID:   instanceID,
		crn:          acc.SecretsManagerHsmPpsGatewayCrn,
		serverCAPath: acc.SecretsManagerHsmServerCaPemPath,
		caCertPath:   acc.SecretsManagerHsmClientCaCertPath,
		caKeyPath:    acc.SecretsManagerHsmClientCaKeyPath,
		suffix:       acctest.RandIntRange(10, 100),
		timeout:      30 * time.Minute,
	}
	if f.instanceID == "" || f.crn == "" || f.serverCAPath == "" || f.caCertPath == "" || f.caKeyPath == "" || acc.SecretsManagerHsmMemberSerials == "" {
		t.Skip("set the Secrets Manager instance ID and SM_HSM_PPS_GATEWAY_CRN, SM_HSM_SERVER_CA_PEM_PATH, SM_HSM_CLIENT_CA_CERT_PATH, SM_HSM_CLIENT_CA_KEY_PATH, SM_HSM_MEMBER_SERIALS to run the HSM lifecycle tests")
	}
	if v, err := strconv.Atoi(os.Getenv("SM_HSM_APPROVAL_TIMEOUT_MINUTES")); err == nil && v > 0 {
		f.timeout = time.Duration(v) * time.Minute
	}
	acc.TestAccPreCheck(t)
	client, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		t.Fatal(err)
	}
	f.client = client
	f.meta = acc.TestAccProvider.Meta()
	f.ctx = context.Background()
	return f
}

func hsmLifecycleStep(t *testing.T, run func(f *hsmFlow, t *testing.T)) {
	if hsmLifecycleBroken {
		t.Skip("a previous lifecycle step failed")
	}
	if hsmLifecycle == nil {
		hsmLifecycle = newHsmFlow(t)
	}
	defer func() {
		if t.Failed() {
			hsmLifecycleBroken = true
		}
	}()
	run(hsmLifecycle, t)
}

func TestAccIbmSmHsmLifecycleStep01DestinationCreate(t *testing.T) {
	hsmLifecycleStep(t, (*hsmFlow).step01DestinationCreate)
}
func TestAccIbmSmHsmLifecycleStep02HaGroupCreate(t *testing.T) {
	hsmLifecycleStep(t, (*hsmFlow).step02HaGroupCreate)
}
func TestAccIbmSmHsmLifecycleStep03PathCreate(t *testing.T) {
	hsmLifecycleStep(t, (*hsmFlow).step03PathCreate)
}
func TestAccIbmSmHsmLifecycleStep04ServerCaCreate(t *testing.T) {
	hsmLifecycleStep(t, (*hsmFlow).step04ServerCaCreate)
}
func TestAccIbmSmHsmLifecycleStep05ServerCaList(t *testing.T) {
	hsmLifecycleStep(t, (*hsmFlow).step05ServerCaList)
}
func TestAccIbmSmHsmLifecycleStep06ServerCaActivate(t *testing.T) {
	hsmLifecycleStep(t, (*hsmFlow).step06ServerCaActivate)
}
func TestAccIbmSmHsmLifecycleStep07PathBuild(t *testing.T) {
	hsmLifecycleStep(t, (*hsmFlow).step07PathBuild)
}
func TestAccIbmSmHsmLifecycleStep08PathHealthCheckWaitCompleted(t *testing.T) {
	hsmLifecycleStep(t, (*hsmFlow).step08PathHealthCheck)
}
func TestAccIbmSmHsmLifecycleStep09ClientCertificateCreate(t *testing.T) {
	hsmLifecycleStep(t, (*hsmFlow).step09ClientCertificateCreate)
}
func TestAccIbmSmHsmLifecycleStep10ClientCertificateActivate(t *testing.T) {
	hsmLifecycleStep(t, (*hsmFlow).step10ClientCertificateActivate)
}
func TestAccIbmSmHsmLifecycleStep11DestinationApply(t *testing.T) {
	hsmLifecycleStep(t, (*hsmFlow).step11DestinationApply)
}
func TestAccIbmSmHsmLifecycleStep12DestinationHealth(t *testing.T) {
	hsmLifecycleStep(t, (*hsmFlow).step12DestinationHealth)
}
func TestAccIbmSmHsmLifecycleStep13DestinationDataSources(t *testing.T) {
	hsmLifecycleStep(t, (*hsmFlow).step13DestinationDataSources)
}
func TestAccIbmSmHsmLifecycleStep14HaGroupDataSources(t *testing.T) {
	hsmLifecycleStep(t, (*hsmFlow).step14HaGroupDataSources)
}
func TestAccIbmSmHsmLifecycleStep15PathDataSources(t *testing.T) {
	hsmLifecycleStep(t, (*hsmFlow).step15PathDataSources)
}
func TestAccIbmSmHsmLifecycleStep16ServerCaDataSource(t *testing.T) {
	hsmLifecycleStep(t, (*hsmFlow).step16ServerCaDataSource)
}
func TestAccIbmSmHsmLifecycleStep17ClientCertificateDataSources(t *testing.T) {
	hsmLifecycleStep(t, (*hsmFlow).step17ClientCertificateDataSources)
}

// Step 18 always runs, even after a failed step, so the destination is not left behind.
func TestAccIbmSmHsmLifecycleStep18DestinationDelete(t *testing.T) {
	f := hsmLifecycle
	if f == nil || f.dest == nil {
		t.Skip("no lifecycle destination was created")
	}
	r := acc.TestAccProvider.ResourcesMap["ibm_sm_hsm_destination"]
	if diags := r.DeleteContext(f.ctx, f.dest, f.meta); diags.HasError() {
		t.Fatalf("delete destination %s: %s", f.dest.Id(), diagString(diags))
	}
	_, resp, err := f.client.GetInstanceHsmDestination(f.client.NewGetInstanceHsmDestinationOptions(f.instanceID, f.destinationID()))
	if err == nil || resp == nil || resp.StatusCode != 404 {
		t.Fatalf("destination %s still exists after delete", f.destinationID())
	}
}

type hsmFlow struct {
	ctx    context.Context
	meta   interface{}
	client *secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2

	instanceID, crn, serverCAPath, caCertPath, caKeyPath string
	suffix                                               int
	timeout                                              time.Duration

	dest, ha, path, serverCA, build, cert *schema.ResourceData
}

func diagString(diags diag.Diagnostics) string {
	var out []string
	for _, d := range diags {
		out = append(out, d.Summary+" "+d.Detail)
	}
	return strings.Join(out, "; ")
}

// create runs the provider's create function for a resource, like terraform apply does.
func (f *hsmFlow) create(t *testing.T, resourceType string, raw map[string]interface{}) *schema.ResourceData {
	t.Helper()
	r := acc.TestAccProvider.ResourcesMap[resourceType]
	if r == nil {
		t.Fatalf("resource %s is not registered in the provider", resourceType)
	}
	d := schema.TestResourceDataRaw(t, r.Schema, raw)
	defer func() {
		if p := recover(); p != nil {
			t.Fatalf("%s create panicked: %v\n%s", resourceType, p, debug.Stack())
		}
	}()
	if diags := r.CreateContext(f.ctx, d, f.meta); diags.HasError() {
		t.Fatalf("%s create failed: %s", resourceType, diagString(diags))
	}
	if d.Id() == "" {
		t.Fatalf("%s create returned an empty id", resourceType)
	}
	return d
}

// read runs the provider's data source read function.
func (f *hsmFlow) read(t *testing.T, dataSource string, raw map[string]interface{}) *schema.ResourceData {
	t.Helper()
	ds := acc.TestAccProvider.DataSourcesMap[dataSource]
	if ds == nil {
		t.Fatalf("data source %s is not registered in the provider", dataSource)
	}
	d := schema.TestResourceDataRaw(t, ds.Schema, raw)
	if diags := ds.ReadContext(f.ctx, d, f.meta); diags.HasError() {
		t.Fatalf("%s read failed: %s", dataSource, diagString(diags))
	}
	return d
}

func lastIDPart(id string) string { return id[strings.LastIndex(id, "/")+1:] }

func (f *hsmFlow) destinationID() string { return f.dest.Get("destination_id").(string) }
func (f *hsmFlow) pathID() string        { return f.path.Get("path_id").(string) }

func (f *hsmFlow) step01DestinationCreate(t *testing.T) {
	f.dest = f.create(t, "ibm_sm_hsm_destination", map[string]interface{}{
		"instance_id": f.instanceID,
		"name":        fmt.Sprintf("tf_lifecycle_dest_%d", f.suffix),
		"vendor":      "thales_luna",
	})
	if f.destinationID() == "" {
		t.Fatal("destination_id is empty")
	}
}

func (f *hsmFlow) step02HaGroupCreate(t *testing.T) {
	// TestResourceDataRaw cannot convert []string; it needs []interface{}.
	memberSerials := []interface{}{}
	for _, s := range hsmMemberSerials() {
		memberSerials = append(memberSerials, s)
	}
	f.ha = f.create(t, "ibm_sm_hsm_ha_group", map[string]interface{}{
		"instance_id":    f.instanceID,
		"destination_id": f.destinationID(),
		"name":           fmt.Sprintf("tf_lifecycle_ha_%d", f.suffix),
		"group_label":    hsmGroupLabel(),
		"group_serial":   hsmGroupSerial(),
		"member_serials": memberSerials,
	})
}

func (f *hsmFlow) step03PathCreate(t *testing.T) {
	f.path = f.create(t, "ibm_sm_hsm_path", map[string]interface{}{
		"instance_id":    f.instanceID,
		"destination_id": f.destinationID(),
		"name":           fmt.Sprintf("tf_lifecycle_path_%d", f.suffix),
		"crn":            f.crn,
	})
	if f.pathID() == "" {
		t.Fatal("path_id is empty")
	}
}

func (f *hsmFlow) step04ServerCaCreate(t *testing.T) {
	pemBytes, err := os.ReadFile(f.serverCAPath)
	if err != nil {
		t.Fatal(err)
	}
	f.serverCA = f.create(t, "ibm_sm_hsm_server_ca", map[string]interface{}{
		"instance_id":    f.instanceID,
		"destination_id": f.destinationID(),
		"path_id":        f.pathID(),
		"name":           fmt.Sprintf("tf_lifecycle_ca_%d", f.suffix),
		"certificate":    string(pemBytes),
	})
}

func (f *hsmFlow) step05ServerCaList(t *testing.T) {
	d := f.read(t, "ibm_sm_hsm_server_cas", map[string]interface{}{
		"instance_id":    f.instanceID,
		"destination_id": f.destinationID(),
		"path_id":        f.pathID(),
	})
	name := f.serverCA.Get("name").(string)
	for _, item := range d.Get("server_cas").([]interface{}) {
		if m, ok := item.(map[string]interface{}); ok && m["name"] == name {
			return
		}
	}
	t.Fatalf("server CA %q not found in the list of server CAs", name)
}

func (f *hsmFlow) step06ServerCaActivate(t *testing.T) {
	caID := lastIDPart(f.serverCA.Id())
	opts := f.client.NewUpdateInstanceHsmDestinationPathServerCaOptions(f.instanceID, f.destinationID(), f.pathID(), caID, map[string]interface{}{"status": "active"})
	if _, _, err := f.client.UpdateInstanceHsmDestinationPathServerCa(opts); err != nil {
		t.Fatalf("activate server CA: %s", err)
	}
	get := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathServerCaOptions{}
	get.SetInstanceID(f.instanceID)
	get.SetDestinationID(f.destinationID())
	get.SetPathID(f.pathID())
	get.SetServerCaID(caID)
	ca, _, err := f.client.GetInstanceHsmDestinationPathServerCa(get)
	if err != nil {
		t.Fatal(err)
	}
	if ca.Status == nil || *ca.Status != "active" {
		t.Fatalf("server CA status is %v, want active", ca.Status)
	}
}

func (f *hsmFlow) step07PathBuild(t *testing.T) {
	f.build = f.create(t, "ibm_sm_hsm_path_build", map[string]interface{}{
		"instance_id":    f.instanceID,
		"destination_id": f.destinationID(),
		"path_id":        f.pathID(),
	})
}

// step08PathHealthCheck polls the path health check until provisioning_state is
// "completed". The next steps need a completed path, so the VPE connection must
// be approved in the IBM Cloud console while this step waits.
func (f *hsmFlow) step08PathHealthCheck(t *testing.T) {
	accountID := f.dest.Get("account_id")
	deadline := time.Now().Add(f.timeout)
	for {
		result, _, err := f.client.GetPathVpeStatus(f.client.NewGetPathVpeStatusOptions(f.instanceID, f.destinationID(), f.pathID()))
		if err != nil {
			t.Fatalf("path health check: %s", err)
		}
		state := *result.ProvisioningState
		switch state {
		case secretsmanagerinstancemanagementv2.Path_ProvisioningState_Completed:
			return
		case secretsmanagerinstancemanagementv2.Path_ProvisioningState_ApprovalDeclined,
			secretsmanagerinstancemanagementv2.Path_ProvisioningState_CrnAlreadyInUse,
			secretsmanagerinstancemanagementv2.Path_ProvisioningState_VpeProvisioningFailed:
			t.Fatalf("path %s ended in provisioning_state %q; reset and rebuild the path", f.pathID(), state)
		}
		if time.Now().After(deadline) {
			t.Fatalf("path %s did not reach provisioning_state completed within %s (last state %q). In the IBM Cloud console (VPC Infrastructure > Private Path service gateways): allow account %v in Account policies, then Permit the pending endpoint gateway connection.", f.pathID(), f.timeout, state, accountID)
		}
		time.Sleep(30 * time.Second)
	}
}

func (f *hsmFlow) step09ClientCertificateCreate(t *testing.T) {
	f.cert = f.create(t, "ibm_sm_hsm_client_certificate", map[string]interface{}{
		"instance_id":    f.instanceID,
		"destination_id": f.destinationID(),
		"name":           fmt.Sprintf("tf_lifecycle_cert_%d", f.suffix),
	})
	if f.cert.Get("csr").(string) == "" {
		t.Fatal("client certificate has no csr")
	}
}

func (f *hsmFlow) step10ClientCertificateActivate(t *testing.T) {
	signed, err := testAccHsmLifecycleSignCSR(f.cert.Get("csr").(string), f.caCertPath, f.caKeyPath)
	if err != nil {
		t.Fatal(err)
	}
	certID := lastIDPart(f.cert.Id())
	opts := f.client.NewUpdateInstanceHsmDestinationClientCertificateOptions(f.instanceID, f.destinationID(), certID, map[string]interface{}{
		"signed_cert": signed,
		"status":      "active",
	})
	if _, _, err = f.client.UpdateInstanceHsmDestinationClientCertificate(opts); err != nil {
		t.Fatalf("activate client certificate: %s", err)
	}
	get := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationClientCertificateOptions{}
	get.SetInstanceID(f.instanceID)
	get.SetDestinationID(f.destinationID())
	get.SetCertID(certID)
	cert, _, err := f.client.GetInstanceHsmDestinationClientCertificate(get)
	if err != nil {
		t.Fatal(err)
	}
	if cert.ProvisioningState == nil || *cert.ProvisioningState != secretsmanagerinstancemanagementv2.ClientCertificate_ProvisioningState_Active {
		t.Fatalf("client certificate provisioning_state is %v, want active", cert.ProvisioningState)
	}
}

// waitHealthSettled waits until the destination health is no longer in_progress/evaluating.
func (f *hsmFlow) waitHealthSettled(t *testing.T) string {
	deadline := time.Now().Add(10 * time.Minute)
	for {
		snapshot, _, err := f.client.GetHsmDestinationHealthSnapshot(f.client.NewGetHsmDestinationHealthSnapshotOptions(f.instanceID, f.destinationID()))
		if err != nil {
			t.Fatalf("destination health: %s", err)
		}
		state := *snapshot.HealthState
		if state != secretsmanagerinstancemanagementv2.DestinationHealthSnapshot_HealthState_InProgress &&
			state != secretsmanagerinstancemanagementv2.DestinationHealthSnapshot_HealthState_Evaluating {
			return state
		}
		if time.Now().After(deadline) {
			t.Fatalf("destination %s still %q after 10 minutes", f.destinationID(), state)
		}
		time.Sleep(15 * time.Second)
	}
}

func (f *hsmFlow) step11DestinationApply(t *testing.T) {
	if _, _, err := f.client.ApplyHsmDestination(f.client.NewApplyHsmDestinationOptions(f.instanceID, f.destinationID())); err != nil {
		t.Fatalf("apply destination: %s", err)
	}
	f.waitHealthSettled(t)
}

// step12DestinationHealth checks that the health snapshot was evaluated after apply.
// The HA group uses placeholder serials, so no real HSM answers and the result is
// normally "unhealthy". Any evaluated state is accepted; only a snapshot that is
// still pending, unknown or not_configured (prerequisites missing) fails the step.
func (f *hsmFlow) step12DestinationHealth(t *testing.T) {
	d := f.read(t, "ibm_sm_hsm_destination_health", map[string]interface{}{
		"instance_id":    f.instanceID,
		"destination_id": f.destinationID(),
	})
	switch state := d.Get("health_state").(string); state {
	case secretsmanagerinstancemanagementv2.DestinationHealthSnapshot_HealthState_Healthy,
		secretsmanagerinstancemanagementv2.DestinationHealthSnapshot_HealthState_Degraded,
		secretsmanagerinstancemanagementv2.DestinationHealthSnapshot_HealthState_Unhealthy:
	default:
		t.Fatalf("destination health_state is %q, want healthy, degraded or unhealthy", state)
	}
	checkSet(t, d, "health_as_of")
}

// checkAttr fails the step when a data source attribute is not the expected value.
func checkAttr(t *testing.T, d *schema.ResourceData, key string, want interface{}) {
	t.Helper()
	if got := d.Get(key); got != want {
		t.Fatalf("%s is %v, want %v", key, got, want)
	}
}

// checkSet fails the step when a data source attribute is empty.
func checkSet(t *testing.T, d *schema.ResourceData, keys ...string) {
	t.Helper()
	for _, key := range keys {
		if v := d.Get(key); v == nil || v == "" {
			t.Fatalf("%s is not set", key)
		}
	}
}

// checkListHas fails the step when no item of the list attribute has the given id.
func checkListHas(t *testing.T, d *schema.ResourceData, listKey, id string) {
	t.Helper()
	for _, item := range d.Get(listKey).([]interface{}) {
		if m, ok := item.(map[string]interface{}); ok && m["id"] == id {
			return
		}
	}
	t.Fatalf("id %q not found in %s", id, listKey)
}

func (f *hsmFlow) step13DestinationDataSources(t *testing.T) {
	one := f.read(t, "ibm_sm_hsm_destination", map[string]interface{}{
		"instance_id":    f.instanceID,
		"destination_id": f.destinationID(),
	})
	checkAttr(t, one, "name", f.dest.Get("name"))
	checkAttr(t, one, "vendor", "thales_luna")
	checkSet(t, one, "provisioning_state", "health_state", "created_at", "created_by")

	list := f.read(t, "ibm_sm_hsm_destinations", map[string]interface{}{"instance_id": f.instanceID})
	checkListHas(t, list, "hsm_destinations", f.destinationID())
}

func (f *hsmFlow) step14HaGroupDataSources(t *testing.T) {
	haID := f.ha.Get("ha_group_id").(string)
	one := f.read(t, "ibm_sm_hsm_ha_group", map[string]interface{}{
		"instance_id":    f.instanceID,
		"destination_id": f.destinationID(),
		"ha_group_id":    haID,
	})
	checkAttr(t, one, "name", f.ha.Get("name"))
	checkAttr(t, one, "group_label", hsmGroupLabel())
	checkAttr(t, one, "group_serial", hsmGroupSerial())
	checkAttr(t, one, "member_serials.#", len(hsmMemberSerials()))
	checkSet(t, one, "provisioning_state", "health_state", "created_at", "created_by")

	list := f.read(t, "ibm_sm_hsm_ha_groups", map[string]interface{}{
		"instance_id":    f.instanceID,
		"destination_id": f.destinationID(),
	})
	checkListHas(t, list, "ha_groups", haID)
}

func (f *hsmFlow) step15PathDataSources(t *testing.T) {
	one := f.read(t, "ibm_sm_hsm_path", map[string]interface{}{
		"instance_id":    f.instanceID,
		"destination_id": f.destinationID(),
		"path_id":        f.pathID(),
	})
	checkAttr(t, one, "name", f.path.Get("name"))
	checkAttr(t, one, "crn", f.crn)
	checkSet(t, one, "provisioning_state", "health_state", "created_at", "created_by")

	list := f.read(t, "ibm_sm_hsm_paths", map[string]interface{}{
		"instance_id":    f.instanceID,
		"destination_id": f.destinationID(),
	})
	checkListHas(t, list, "paths", f.pathID())
}

func (f *hsmFlow) step16ServerCaDataSource(t *testing.T) {
	caID := f.serverCA.Get("server_ca_id").(string)
	one := f.read(t, "ibm_sm_hsm_server_ca", map[string]interface{}{
		"instance_id":    f.instanceID,
		"destination_id": f.destinationID(),
		"path_id":        f.pathID(),
		"server_ca_id":   caID,
	})
	checkAttr(t, one, "name", f.serverCA.Get("name"))
	checkAttr(t, one, "status", "active")
	checkSet(t, one, "certificate", "expires_at", "health_state", "created_at", "created_by")
}

func (f *hsmFlow) step17ClientCertificateDataSources(t *testing.T) {
	certID := f.cert.Get("cert_id").(string)
	one := f.read(t, "ibm_sm_hsm_client_certificate", map[string]interface{}{
		"instance_id":    f.instanceID,
		"destination_id": f.destinationID(),
		"cert_id":        certID,
	})
	checkAttr(t, one, "name", f.cert.Get("name"))
	checkAttr(t, one, "provisioning_state", secretsmanagerinstancemanagementv2.ClientCertificate_ProvisioningState_Active)
	checkSet(t, one, "csr", "csr_cn", "signed_cert", "expires_at", "created_at", "created_by")

	list := f.read(t, "ibm_sm_hsm_client_certificates", map[string]interface{}{
		"instance_id":    f.instanceID,
		"destination_id": f.destinationID(),
	})
	checkListHas(t, list, "client_certificates", certID)
}

func testAccHsmLifecycleSignCSR(csrPEM, caCertPath, caKeyPath string) (string, error) {
	csrBlock, _ := pem.Decode([]byte(csrPEM))
	if csrBlock == nil {
		return "", fmt.Errorf("client certificate csr is not PEM")
	}
	csr, err := x509.ParseCertificateRequest(csrBlock.Bytes)
	if err != nil {
		return "", err
	}
	caCertBytes, err := os.ReadFile(caCertPath)
	if err != nil {
		return "", err
	}
	caKeyBytes, err := os.ReadFile(caKeyPath)
	if err != nil {
		return "", err
	}
	caCertBlock, _ := pem.Decode(caCertBytes)
	caKeyBlock, _ := pem.Decode(caKeyBytes)
	if caCertBlock == nil || caKeyBlock == nil {
		return "", fmt.Errorf("CA certificate or key is not PEM")
	}
	caCert, err := x509.ParseCertificate(caCertBlock.Bytes)
	if err != nil {
		return "", err
	}
	var caKey crypto.Signer
	if k, err := x509.ParsePKCS8PrivateKey(caKeyBlock.Bytes); err == nil {
		var ok bool
		if caKey, ok = k.(crypto.Signer); !ok {
			return "", fmt.Errorf("CA key cannot sign")
		}
	} else if k, err := x509.ParsePKCS1PrivateKey(caKeyBlock.Bytes); err == nil {
		caKey = k
	} else if k, err := x509.ParseECPrivateKey(caKeyBlock.Bytes); err == nil {
		caKey = k
	} else {
		return "", fmt.Errorf("unsupported CA key format")
	}

	serial, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 127))
	if err != nil {
		return "", err
	}
	template := &x509.Certificate{
		SerialNumber: serial,
		Subject:      pkix.Name{CommonName: csr.Subject.CommonName},
		NotBefore:    time.Now().Add(-time.Minute),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
	}
	der, err := x509.CreateCertificate(rand.Reader, template, caCert, csr.PublicKey, caKey)
	if err != nil {
		return "", err
	}
	return string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})), nil
}
