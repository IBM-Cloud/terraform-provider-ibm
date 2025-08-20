package flex

import (
	"errors"
	"fmt"

	v "github.com/IBM-Cloud/terraform-provider-ibm/version"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	sdkDiag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// TerraformProblem provides a type that holds standardized information
// suitable to problems that occur in the Terraform Provider code.
type TerraformProblem struct {
	*core.IBMProblem

	Resource  string
	Operation string
}

// Compile-time check that the TerraformProblem struct implements
// the Terraform Plugin Framework's Diagnostic interface.
var _ diag.Diagnostic = &TerraformProblem{}

// GetID returns a hash value computed from stable fields in the
// TerraformProblem instance, including Resource and Operation.
func (e *TerraformProblem) GetID() string {
	return core.CreateIDHash("terraform", e.GetBaseSignature(), e.Resource, e.Operation)
}

// GetConsoleMessage returns the fields of the problem that
// are relevant to a user, formatted as a YAML string.
func (e *TerraformProblem) GetConsoleMessage() string {
	return core.ComputeConsoleMessage(e)
}

// GetConsoleMessage returns the fields of the problem that
// are relevant to a developer, formatted as a YAML string.
func (e *TerraformProblem) GetDebugMessage() string {
	return core.ComputeDebugMessage(e)
}

func (e *TerraformProblem) GetConsoleOrderedMaps() *core.OrderedMaps {
	orderedMaps := core.NewOrderedMaps()

	orderedMaps.Add("id", e.GetID())
	orderedMaps.Add("summary", e.IBMProblem.Summary)
	orderedMaps.Add("severity", e.IBMProblem.Severity)
	orderedMaps.Add("resource", e.Resource)
	orderedMaps.Add("operation", e.Operation)
	orderedMaps.Add("component", e.IBMProblem.Component)

	return orderedMaps
}

func (e *TerraformProblem) GetDebugOrderedMaps() *core.OrderedMaps {
	orderedMaps := e.GetConsoleOrderedMaps()

	var orderableCausedBy core.OrderableProblem
	if errors.As(e.GetCausedBy(), &orderableCausedBy) {
		orderedMaps.Add("caused_by", orderableCausedBy.GetDebugOrderedMaps().GetMaps())
	}

	return orderedMaps
}

// GetDiag returns a new Diagnostics object using the console
// message as the summary. It is used to create a Diagnostics
// object from a TerraformProblem in the resource/data source code.
func (e *TerraformProblem) GetDiag() sdkDiag.Diagnostics {
	return sdkDiag.Errorf("%s", e.GetConsoleMessage())
}

// Implement the Terraform Plugin Framework `Diagnostic` interface.
func (e *TerraformProblem) Severity() diag.Severity {
	if e.IBMProblem.Severity == core.WarningSeverity {
		return diag.SeverityWarning
	}

	return diag.SeverityError
}

func (e *TerraformProblem) Summary() string {
	return e.IBMProblem.Summary
}

func (e *TerraformProblem) Detail() string {
	return e.GetConsoleMessage()
}

func (e *TerraformProblem) Equal(d diag.Diagnostic) bool {
	if tfErr, ok := d.(*TerraformProblem); ok {
		return tfErr.GetID() == e.GetID()
	}
	return false
}

// TerraformErrorf creates and returns a new instance of `TerraformProblem`
// with "error" level severity and a blank discriminator - the "caused by"
// error is used to ensure uniqueness. This is a convenience function to
// use when creating a new TerraformProblem instance from an error that
// came from the SDK.
func TerraformErrorf(err error, summary, resource, operation string) *TerraformProblem {
	return DiscriminatedTerraformErrorf(err, summary, resource, operation, "")
}

// DiscriminatedTerraformErrorf creates and returns a new instance
// of `TerraformProblem` with "error" level severity that contains
// a discriminator used to make the instance unique relative to
// other problem scenarios in the same resource/operation.
func DiscriminatedTerraformErrorf(err error, summary, resource, operation, discriminator string) *TerraformProblem {
	return &TerraformProblem{
		IBMProblem: core.IBMErrorf(err, getComponentInfo(), summary, discriminator),
		Resource:   resource,
		Operation:  operation,
	}
}

// FromDiagnostic accepts an instance of the Diagnostic interface and converts
// it to a new TerraformProblem instance, preserving the information stored in
// the diagnostic.
func FromDiagnostic(d diag.Diagnostic, resource, operation, discriminator string) *TerraformProblem {
	// Since the TerraformProblem struct implements the Diagnositc interface,
	// check if this is already a TerraformProblem instance. If it is, return
	// it to preserve its details.
	if asProb, ok := d.(*TerraformProblem); ok {
		return asProb
	}

	// Combine the "summary" and "details" to preserve the original diagnostic information.
	msg := fmt.Sprintf("%s: %s", d.Summary(), d.Detail())
	prob := &TerraformProblem{
		IBMProblem: core.IBMErrorf(nil, getComponentInfo(), msg, discriminator),
		Resource:   resource,
		Operation:  operation,
	}

	// IBM problems are errors by default but the diagnostic instances may have
	// warning severity - ensure the severity is preserved.
	if d.Severity() == diag.SeverityWarning {
		prob.IBMProblem.Severity = core.WarningSeverity
	}

	return prob
}

func getComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent("github.com/IBM-Cloud/terraform-provider-ibm", v.Version)
}

// AddCreateProblems converts Diagnostics instances to TerraformProblem instances and
// appends them to the given CreateResponse instance.
func AddCreateProblems(resp *resource.CreateResponse, d diag.Diagnostics, resource, discriminator string) {
	addProblems(&resp.Diagnostics, &d, resource, "create", discriminator)
}

// AddUpdateProblems converts Diagnostics instances to TerraformProblem instances and
// appends them to the given UpdateResponse instance.
func AddUpdateProblems(resp *resource.UpdateResponse, d diag.Diagnostics, resource, discriminator string) {
	addProblems(&resp.Diagnostics, &d, resource, "update", discriminator)
}

// AddReadProblems converts Diagnostics instances to TerraformProblem instances and
// appends them to the given ReadResponse instance.
func AddReadProblems(resp *resource.ReadResponse, d diag.Diagnostics, resource, discriminator string) {
	addProblems(&resp.Diagnostics, &d, resource, "read", discriminator)
}

// AddDeleteProblems converts Diagnostics instances to TerraformProblem instances and
// appends them to the given DeleteResponse instance.
func AddDeleteProblems(resp *resource.DeleteResponse, d diag.Diagnostics, resource, discriminator string) {
	addProblems(&resp.Diagnostics, &d, resource, "delete", discriminator)
}

// AddDataSourceReadProblems converts Diagnostics instances to TerraformProblem instances and
// appends them to the given ReadResponse instance.
func AddDataSourceReadProblems(resp *datasource.ReadResponse, d diag.Diagnostics, datasource, discriminator string) {
	addProblems(&resp.Diagnostics, &d, "(Data) "+datasource, "read", discriminator)
}

func addProblems(target, source *diag.Diagnostics, resource, operation, discriminator string) {
	for _, diagnostic := range *source {
		target.Append(FromDiagnostic(diagnostic, resource, operation, discriminator))
	}
}

// FmtErrorf wraps `fmt.Errorf(format string, a ...interface{}) error`
// and attempts to return a TerraformProblem instance instead of a
// plain error instance, if an error object is found among the arguments
func FmtErrorf(format string, a ...interface{}) error {
	intendedError := fmt.Errorf(format, a...)

	var err error
	for _, arg := range a {
		// Look for an error instance among the arguments.

		if errArg, ok := arg.(error); ok {
			err = errArg
		} else if ser, ok := arg.(*ServiceErrorResponse); ok {
			// Deal with the "ServiceErrorResponse" type, which
			// wraps errors in some of the handwritten code.
			err = ser.Error
		}

		if err != nil {
			var tfError *TerraformProblem
			if !errors.As(err, &tfError) {
				tfError = TerraformErrorf(err, err.Error(), "", "")
			}

			tfError.IBMProblem.Summary = intendedError.Error()
			return tfError
		}
	}

	return intendedError
}
