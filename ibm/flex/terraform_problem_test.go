package flex

import (
	"errors"
	"fmt"
	"testing"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

const (
	MODULE_NAME  = "github.com/IBM-Cloud/terraform-provider-ibm"
	MOCK_VERSION = "1.63.0"
)

func TestTerraformProblemEmbedsIBMProblem(t *testing.T) {
	terraformProb := &TerraformProblem{}

	// Check that the methods defined by IBMProblem are supported here.
	assert.NotNil(t, terraformProb.Error)
	assert.NotNil(t, terraformProb.GetBaseSignature)
	assert.NotNil(t, terraformProb.GetCausedBy)
	assert.NotNil(t, terraformProb.Unwrap)
}

func TestTerraformProblemGetConsoleMessage(t *testing.T) {
	terraformProb := getPopulatedTerraformProblem()
	message := terraformProb.GetConsoleMessage()
	expected := `---
id: terraform-98c0e1fd
summary: Create failed.
severity: error
resource: ibm_some_resource
operation: create
component:
  name: github.com/IBM-Cloud/terraform-provider-ibm
  version: 1.63.0
---
`
	assert.Equal(t, expected, message)
}

func TestTerraformProblemGetDebugMessage(t *testing.T) {
	terraformProb := getPopulatedTerraformProblem()
	message := terraformProb.GetDebugMessage()
	expected := `---
id: terraform-98c0e1fd
summary: Create failed.
severity: error
resource: ibm_some_resource
operation: create
component:
  name: github.com/IBM-Cloud/terraform-provider-ibm
  version: 1.63.0
---
`
	assert.Equal(t, expected, message)
}

func TestTerraformProblemGetID(t *testing.T) {
	terraformProb := getPopulatedTerraformProblem()
	assert.Equal(t, "terraform-98c0e1fd", terraformProb.GetID())
}

func TestTerraformProblemGetConsoleOrderedMaps(t *testing.T) {
	terraformProb := getPopulatedTerraformProblem()
	orderedMaps := terraformProb.GetConsoleOrderedMaps()
	assert.NotNil(t, orderedMaps)

	maps := orderedMaps.GetMaps()
	assert.NotNil(t, maps)
	assert.Len(t, maps, 6)

	assert.Equal(t, "id", maps[0].Key)
	assert.Equal(t, "terraform-98c0e1fd", maps[0].Value)

	assert.Equal(t, "summary", maps[1].Key)
	assert.Equal(t, "Create failed.", maps[1].Value)

	assert.Equal(t, "severity", maps[2].Key)
	assert.Equal(t, core.ErrorSeverity, maps[2].Value)

	assert.Equal(t, "resource", maps[3].Key)
	assert.Equal(t, "ibm_some_resource", maps[3].Value)

	assert.Equal(t, "operation", maps[4].Key)
	assert.Equal(t, "create", maps[4].Value)

	assert.Equal(t, "component", maps[5].Key)
	assert.Equal(t, MODULE_NAME, maps[5].Value.(*core.ProblemComponent).Name)
	assert.Equal(t, MOCK_VERSION, maps[5].Value.(*core.ProblemComponent).Version)
}

func TestTerraformProblemGetDebugOrderedMaps(t *testing.T) {
	terraformProb := getPopulatedTerraformProblem()
	orderedMaps := terraformProb.GetDebugOrderedMaps()
	assert.NotNil(t, orderedMaps)

	maps := orderedMaps.GetMaps()
	assert.NotNil(t, maps)
	assert.Len(t, maps, 6)

	assert.Equal(t, "id", maps[0].Key)
	assert.Equal(t, "terraform-98c0e1fd", maps[0].Value)

	assert.Equal(t, "summary", maps[1].Key)
	assert.Equal(t, "Create failed.", maps[1].Value)

	assert.Equal(t, "severity", maps[2].Key)
	assert.Equal(t, core.ErrorSeverity, maps[2].Value)

	assert.Equal(t, "resource", maps[3].Key)
	assert.Equal(t, "ibm_some_resource", maps[3].Value)

	assert.Equal(t, "operation", maps[4].Key)
	assert.Equal(t, "create", maps[4].Value)

	assert.Equal(t, "component", maps[5].Key)
	assert.Equal(t, MODULE_NAME, maps[5].Value.(*core.ProblemComponent).Name)
	assert.Equal(t, MOCK_VERSION, maps[5].Value.(*core.ProblemComponent).Version)
}

func TestTerraformProblemGetDiag(t *testing.T) {
	terraformProb := getPopulatedTerraformProblem()
	diagnostics := terraformProb.GetDiag()

	assert.True(t, diagnostics.HasError())
	assert.Len(t, diagnostics, 1)

	diagnostic := diagnostics[0]
	assert.Nil(t, diagnostic.Validate())
	assert.Equal(t, terraformProb.GetConsoleMessage(), diagnostic.Summary)
}

func TestTerraformProblemSeverity(t *testing.T) {
	terraformProb := getPopulatedTerraformProblem()
	assert.Equal(t, terraformProb.Severity(), diag.SeverityError)
}

func TestTerraformProblemSummary(t *testing.T) {
	terraformProb := getPopulatedTerraformProblem()
	assert.Equal(t, terraformProb.Summary(), terraformProb.IBMProblem.Summary)
}

func TestTerraformProblemDetail(t *testing.T) {
	terraformProb := getPopulatedTerraformProblem()
	assert.Equal(t, terraformProb.Detail(), terraformProb.GetConsoleMessage())
}

func TestTerraformProblemEqual(t *testing.T) {
	terraformProb := getPopulatedTerraformProblem()

	terraformProbCopy := getPopulatedTerraformProblem()
	assert.True(t, terraformProb.Equal(terraformProbCopy))

	terraformProbCopy.Resource = "changed"
	assert.False(t, terraformProb.Equal(terraformProbCopy))
}

func TestAddCreateProblems(t *testing.T) {
	resp := &resource.CreateResponse{}
	var diags diag.Diagnostics

	diags.AddError("my summary", "my extra details")

	AddCreateProblems(resp, diags, "my_resource", "my_discriminator")
	assert.True(t, resp.Diagnostics.HasError())
	assert.Equal(t, 1, resp.Diagnostics.ErrorsCount())
	asDiag := resp.Diagnostics.Errors()[0]
	assert.Equal(t, "my summary: my extra details", asDiag.Summary())
	assert.Equal(t, diag.SeverityError, asDiag.Severity())

	asProblem, ok := asDiag.(*TerraformProblem)
	assert.True(t, ok)
	assert.Equal(t, asProblem.GetDebugMessage(), asDiag.Detail())
	assert.Equal(t, core.ErrorSeverity, asProblem.IBMProblem.Severity)
	assert.Equal(t, "create", asProblem.Operation)
	assert.Equal(t, "my_resource", asProblem.Resource)
	assert.Equal(t, "terraform-9beeb6f3", asProblem.GetID())
}

func TestAddUpdateProblems(t *testing.T) {
	resp := &resource.UpdateResponse{}
	var diags diag.Diagnostics

	diags.AddError("my summary", "my extra details")

	AddUpdateProblems(resp, diags, "my_resource", "my_discriminator")
	assert.True(t, resp.Diagnostics.HasError())
	assert.Equal(t, 1, resp.Diagnostics.ErrorsCount())
	asDiag := resp.Diagnostics.Errors()[0]
	assert.Equal(t, "my summary: my extra details", asDiag.Summary())
	assert.Equal(t, diag.SeverityError, asDiag.Severity())

	asProblem, ok := asDiag.(*TerraformProblem)
	assert.True(t, ok)
	assert.Equal(t, asProblem.GetDebugMessage(), asDiag.Detail())
	assert.Equal(t, core.ErrorSeverity, asProblem.IBMProblem.Severity)
	assert.Equal(t, "update", asProblem.Operation)
	assert.Equal(t, "my_resource", asProblem.Resource)
	assert.Equal(t, "terraform-32a9172a", asProblem.GetID())
}

func TestAddReadProblems(t *testing.T) {
	resp := &resource.ReadResponse{}
	var diags diag.Diagnostics

	diags.AddError("my summary", "my extra details")

	AddReadProblems(resp, diags, "my_resource", "my_discriminator")
	assert.True(t, resp.Diagnostics.HasError())
	assert.Equal(t, 1, resp.Diagnostics.ErrorsCount())
	asDiag := resp.Diagnostics.Errors()[0]
	assert.Equal(t, "my summary: my extra details", asDiag.Summary())
	assert.Equal(t, diag.SeverityError, asDiag.Severity())

	asProblem, ok := asDiag.(*TerraformProblem)
	assert.True(t, ok)
	assert.Equal(t, asProblem.GetDebugMessage(), asDiag.Detail())
	assert.Equal(t, core.ErrorSeverity, asProblem.IBMProblem.Severity)
	assert.Equal(t, "read", asProblem.Operation)
	assert.Equal(t, "my_resource", asProblem.Resource)
	assert.Equal(t, "terraform-e4739647", asProblem.GetID())
}

func TestAddDeleteProblems(t *testing.T) {
	resp := &resource.DeleteResponse{}
	var diags diag.Diagnostics

	diags.AddError("my summary", "my extra details")

	AddDeleteProblems(resp, diags, "my_resource", "my_discriminator")
	assert.True(t, resp.Diagnostics.HasError())
	assert.Equal(t, 1, resp.Diagnostics.ErrorsCount())
	asDiag := resp.Diagnostics.Errors()[0]
	assert.Equal(t, "my summary: my extra details", asDiag.Summary())
	assert.Equal(t, diag.SeverityError, asDiag.Severity())

	asProblem, ok := asDiag.(*TerraformProblem)
	assert.True(t, ok)
	assert.Equal(t, asProblem.GetDebugMessage(), asDiag.Detail())
	assert.Equal(t, core.ErrorSeverity, asProblem.IBMProblem.Severity)
	assert.Equal(t, "delete", asProblem.Operation)
	assert.Equal(t, "my_resource", asProblem.Resource)
	assert.Equal(t, "terraform-fbc62906", asProblem.GetID())
}

func TestAddDataSourceReadProblems(t *testing.T) {
	resp := &datasource.ReadResponse{}
	var diags diag.Diagnostics

	diags.AddError("my summary", "my extra details")

	AddDataSourceReadProblems(resp, diags, "my_resource", "my_discriminator")
	assert.True(t, resp.Diagnostics.HasError())
	assert.Equal(t, 1, resp.Diagnostics.ErrorsCount())
	asDiag := resp.Diagnostics.Errors()[0]
	assert.Equal(t, "my summary: my extra details", asDiag.Summary())
	assert.Equal(t, diag.SeverityError, asDiag.Severity())

	asProblem, ok := asDiag.(*TerraformProblem)
	assert.True(t, ok)
	assert.Equal(t, asProblem.GetDebugMessage(), asDiag.Detail())
	assert.Equal(t, core.ErrorSeverity, asProblem.IBMProblem.Severity)
	assert.Equal(t, "read", asProblem.Operation)
	assert.Equal(t, "(Data) my_resource", asProblem.Resource)
	assert.Equal(t, "terraform-502ffcc1", asProblem.GetID())
}

func TestAddProblems(t *testing.T) {
	var targetDiags, sourceDiags diag.Diagnostics

	targetDiags.Append(getPopulatedTerraformProblem())
	sourceDiags.AddError("summary", "details")

	addProblems(&targetDiags, &sourceDiags, "my_resource", "create", "my_discriminator")

	assert.True(t, targetDiags.HasError())
	assert.Equal(t, 2, targetDiags.ErrorsCount())
}

func TestAddProblemsWithNils(t *testing.T) {
	var targetDiags, sourceDiags diag.Diagnostics

	assert.Nil(t, targetDiags)
	assert.Nil(t, sourceDiags)

	addProblems(&targetDiags, &sourceDiags, "my_resource", "create", "my_discriminator")

	assert.False(t, targetDiags.HasError())
	assert.Equal(t, 0, targetDiags.ErrorsCount())
}

func TestFromDiagnostic(t *testing.T) {
	var d diag.Diagnostic

	// Error
	d = diag.NewErrorDiagnostic("summary", "details")
	prob := FromDiagnostic(d, "resource", "operation", "discriminator")
	assert.Equal(t, "summary: details", prob.IBMProblem.Summary)
	assert.Equal(t, core.ErrorSeverity, prob.IBMProblem.Severity)
	assert.Equal(t, "operation", prob.Operation)
	assert.Equal(t, "resource", prob.Resource)
	assert.Equal(t, "terraform-01ca8439", prob.GetID())

	// Warning
	d = diag.NewWarningDiagnostic("summary", "details")
	prob = FromDiagnostic(d, "resource", "operation", "discriminator")
	assert.Equal(t, "summary: details", prob.IBMProblem.Summary)
	assert.Equal(t, core.WarningSeverity, prob.IBMProblem.Severity)
	assert.Equal(t, "operation", prob.Operation)
	assert.Equal(t, "resource", prob.Resource)
	assert.Equal(t, "terraform-140a3d88", prob.GetID())

	// Existing TerraformProblem - ignores the given values
	d = getPopulatedTerraformProblem()
	prob = FromDiagnostic(d, "resource", "operation", "discriminator")
	assert.Equal(t, "Create failed.", prob.IBMProblem.Summary)
	assert.Equal(t, core.ErrorSeverity, prob.IBMProblem.Severity)
	assert.Equal(t, "create", prob.Operation)
	assert.Equal(t, "ibm_some_resource", prob.Resource)
	assert.Equal(t, "terraform-98c0e1fd", prob.GetID())
}

func TestTerraformErrorf(t *testing.T) {
	causedBy := &core.SDKProblem{}
	summary := "Update failed."
	resourceName := "ibm_some_resource"
	operation := "update"

	terraformProb := TerraformErrorf(causedBy, summary, resourceName, operation)
	assert.NotNil(t, terraformProb)
	assert.Equal(t, summary, terraformProb.IBMProblem.Summary)
	assert.Equal(t, getComponentInfo(), terraformProb.IBMProblem.Component)
	assert.Equal(t, core.ErrorSeverity, terraformProb.IBMProblem.Severity)
}

func TestFmtErrorfWithProblem(t *testing.T) {
	msg := "Request failed."
	sdkProb := &core.SDKProblem{
		IBMProblem: &core.IBMProblem{
			Summary: msg,
		},
	}

	err := FmtErrorf("Operation failed: %s", sdkProb)

	var errAsProb *core.SDKProblem
	assert.True(t, errors.As(err, &errAsProb))
	assert.NotNil(t, errAsProb)
	assert.Equal(t, msg, errAsProb.Summary)
}

func TestFmtErrorfWithNoError(t *testing.T) {
	msg := "Bad input"
	err := FmtErrorf("Operation failed: %s", msg)
	assert.Equal(t, "Operation failed: Bad input", err.Error())
	assert.Equal(t, fmt.Errorf("Operation failed: %s", msg).Error(), err.Error())

	var errAsProb core.Problem
	assert.False(t, errors.As(err, &errAsProb))
	assert.Nil(t, errAsProb)
}

func TestFmtErrorfWithProblemInServiceErrorResponse(t *testing.T) {
	msg := "Request failed."
	sdkProb := &core.SDKProblem{
		IBMProblem: &core.IBMProblem{
			Summary: msg,
		},
	}

	ser := &ServiceErrorResponse{
		Error: sdkProb,
	}

	err := FmtErrorf("Operation failed: %s", ser)

	var errAsProb *core.SDKProblem
	assert.True(t, errors.As(err, &errAsProb))
	assert.NotNil(t, errAsProb)
	assert.Equal(t, msg, errAsProb.Summary)
}

func TestDiscriminatedTerraformErrorf(t *testing.T) {
	summary := "Update failed."
	resourceName := "ibm_some_resource"
	operation := "update"
	discriminator := "failed-to-read-input"

	terraformProb := DiscriminatedTerraformErrorf(nil, summary, resourceName, operation, discriminator)
	assert.NotNil(t, terraformProb)
	assert.Equal(t, summary, terraformProb.IBMProblem.Summary)
	assert.Equal(t, getComponentInfo(), terraformProb.IBMProblem.Component)
	assert.Equal(t, core.ErrorSeverity, terraformProb.IBMProblem.Severity)

	// The discriminator field is private, so it can't be checked, so make
	// sure the hash is unique - that is the purpose of the discriminator.
	terraformProbNoDisc := TerraformErrorf(nil, summary, resourceName, operation)
	assert.NotEqual(t, terraformProbNoDisc.GetID(), terraformProb.GetID())
}

func TestGetComponentInfo(t *testing.T) {
	component := getComponentInfo()
	assert.NotNil(t, component)
	assert.Equal(t, MODULE_NAME, component.Name)
	assert.Equal(t, "0.0.1", component.Version)
}

func getPopulatedTerraformProblem() *TerraformProblem {
	return &TerraformProblem{
		IBMProblem: &core.IBMProblem{
			Summary:   "Create failed.",
			Component: core.NewProblemComponent(MODULE_NAME, MOCK_VERSION),
			Severity:  core.ErrorSeverity,
		},
		Resource:  "ibm_some_resource",
		Operation: "create",
	}
}
