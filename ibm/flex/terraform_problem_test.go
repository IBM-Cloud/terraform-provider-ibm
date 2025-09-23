package flex

import (
	"errors"
	"fmt"
	"testing"

	"github.com/IBM/go-sdk-core/v5/core"
	v "github.com/Mavrickk3/terraform-provider-ibm/version"
	"github.com/stretchr/testify/assert"
)

const (
	MODULE_NAME  = "github.com/Mavrickk3/terraform-provider-ibm"
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
  name: github.com/Mavrickk3/terraform-provider-ibm
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
  name: github.com/Mavrickk3/terraform-provider-ibm
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

func TestTerraformErrorf(t *testing.T) {
	causedBy := &core.SDKProblem{}
	summary := "Update failed."
	resourceName := "ibm_some_resource"
	operation := "update"

	terraformProb := TerraformErrorf(causedBy, summary, resourceName, operation)
	assert.NotNil(t, terraformProb)
	assert.Equal(t, summary, terraformProb.Summary)
	assert.Equal(t, getComponentInfo(), terraformProb.Component)
	assert.Equal(t, core.ErrorSeverity, terraformProb.Severity)
}

func TestFmtErrorfWithProblem(t *testing.T) {
	msg := "Request failed."
	sdkProb := &core.SDKProblem{
		IBMProblem: &core.IBMProblem{
			Summary: msg,
		},
	}

	err := FmtErrorf("Operation failed: %s", sdkProb)

	var tfErr *TerraformProblem
	assert.ErrorAs(t, err, &tfErr)
	assert.NotNil(t, tfErr)
	assert.Equal(t, "Operation failed: Request failed.", tfErr.Summary)
	assert.ErrorIs(t, tfErr.GetCausedBy(), sdkProb)
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

	var tfErr *TerraformProblem
	assert.ErrorAs(t, err, &tfErr)
	assert.NotNil(t, tfErr)
	assert.Equal(t, fmt.Sprintf("Operation failed: %s", ser), tfErr.Summary)
	assert.ErrorIs(t, tfErr.GetCausedBy(), sdkProb)
}

func TestDiscriminatedTerraformErrorf(t *testing.T) {
	summary := "Update failed."
	resourceName := "ibm_some_resource"
	operation := "update"
	discriminator := "failed-to-read-input"

	terraformProb := DiscriminatedTerraformErrorf(nil, summary, resourceName, operation, discriminator)
	assert.NotNil(t, terraformProb)
	assert.Equal(t, summary, terraformProb.Summary)
	assert.Equal(t, getComponentInfo(), terraformProb.Component)
	assert.Equal(t, core.ErrorSeverity, terraformProb.Severity)

	// The discriminator field is private, so it can't be checked, so make
	// sure the hash is unique - that is the purpose of the discriminator.
	terraformProbNoDisc := TerraformErrorf(nil, summary, resourceName, operation)
	assert.NotEqual(t, terraformProbNoDisc.GetID(), terraformProb.GetID())
}

func TestGetComponentInfo(t *testing.T) {
	component := getComponentInfo()
	assert.NotNil(t, component)
	assert.Equal(t, MODULE_NAME, component.Name)
	assert.Equal(t, v.Version, component.Version)
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
