package test

import (
	"testing"

	"github.com/Azure/terratest-terraform-fluent/check"
	"github.com/Azure/terratest-terraform-fluent/setuptest"
	"github.com/stretchr/testify/require"
)

const (
	moduleDir = "../"
)

func TestSimple(t *testing.T) {
	t.Parallel()

	// Setup the test
	test, err := setuptest.Dirs(moduleDir, "").WithVars(nil).InitPlanShow(t)
	require.NoError(t, err)
	defer test.Cleanup()

	// Check the results
	check.InPlan(test.Plan).NumberOfResourcesEquals(1).ErrorIsNil(t)
	// How can we check that the input attribute of the terraform_data.example resource is equal to "example"?
	// hint: see the github.com/Azure/terratest-terraform-fluent README...

	t.Fail() // remove this once you have done!
}

func TestCondition(t *testing.T) {
	t.Parallel()

	// Define module variables
	// See https://go.dev/tour/methods/14 for a description of the empty interface.
	// Search the internet for 'golang empty interface' for more examples.
	vars := map[string]interface{}{
		"example_condition": true,
	}

	// Setup two tests!
	testPresent, err := setuptest.Dirs(moduleDir, "").WithVars(vars).InitPlanShow(t)
	require.NoError(t, err)
	testNotPresent, err := setuptest.Dirs(moduleDir, "").WithVars(nil).InitPlanShow(t)
	require.NoError(t, err)
	defer testPresent.Cleanup()
	defer testNotPresent.Cleanup()

	// Check that the resource is present in the test_present plan and that the input attribute has the value "example_condition"...

	// Now check that the resource is NOT present in the test_notpresent plan...

	t.Fail() // remove this once you have done!
}

func TestForEach(t *testing.T) {
	t.Parallel()

	// Define module variables.
	// See https://go.dev/tour/methods/14 for a description of the empty interface.
	// Search the internet for 'golang empty interface' for more examples.
	vars := map[string]interface{}{
		"example_for_each": map[string]string{
			"instance_1": "data_1",
			"instance_2": "data_2",
		},
	}

	// Setup the test
	test, err := setuptest.Dirs(moduleDir, "").WithVars(vars).InitPlanShow(t)
	require.NoError(t, err)
	defer test.Cleanup()

	// Check the results
	check.InPlan(test.Plan).NumberOfResourcesEquals(3).ErrorIsNil(t)

	// How can we check the values of the terraform_data.example_for_each resources?
	// hint: use a loop...

	t.Fail() // remove this once you have done!
}
