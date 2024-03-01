package flex

import (
	"errors"
	"fmt"
	v "github.com/IBM-Cloud/terraform-provider-ibm/version"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// TerraformProblem provides a type that holds standardized information
// suitable to problems that occur in the Terraform Provider code.
type TerraformProblem struct {
	*core.IBMProblem

	Resource  string
	Operation string
}

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
	orderedMaps.Add("summary", e.Summary)
	orderedMaps.Add("severity", e.Severity)
	orderedMaps.Add("resource", e.Resource)
	orderedMaps.Add("operation", e.Operation)
	orderedMaps.Add("component", e.Component)

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
func (e *TerraformProblem) GetDiag() diag.Diagnostics {
	return diag.Errorf("%s", e.GetConsoleMessage())
}

// TerraformErrorf creates and returns a new instance
// of `TerraformProblem` with "error" level severity.
func TerraformErrorf(err error, summary, resource, operation string) *TerraformProblem {
	return &TerraformProblem{
		IBMProblem: core.IBMErrorf(err, getComponentInfo(), summary, ""),
		Resource:   resource,
		Operation:  operation,
	}
}

func getComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent("github.com/IBM-Cloud/terraform-provider-ibm", v.Version)
}

// FmtErrorf wraps `fmt.Errorf(format string, a ...interface{}) error`
// and checks for the instance of a "Problem" type. If it finds one,
// the Problem instance needs to be returned instead of wrapped by
// `fmt.Errorf`.
func FmtErrorf(format string, a ...interface{}) error {
	for _, arg := range a {
		// Look for an error instance among the arguments.
		var err error

		if errArg, ok := arg.(error); ok {
			err = errArg
		} else if ser, ok := arg.(*ServiceErrorResponse); ok {
			// Deal with the "ServiceErrorResponse" type, which
			// wraps errors in some of the handwritten code.
			err = ser.Error
		}

		if err != nil {
			var problem core.Problem
			if errors.As(err, &problem) {
				return problem
			}
		}
	}

	return fmt.Errorf(format, a...)
}
