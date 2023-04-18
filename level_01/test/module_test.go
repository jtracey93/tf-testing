package test

import (
	"fmt"
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

	check.InPlan(test.Plan).That("terraform_data.example").Exists().ErrorIsNil(t)
	check.InPlan(test.Plan).That("terraform_data.example").Key("input").HasValue("example").ErrorIsNil(t)
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
	check.InPlan(testPresent.Plan).That("terraform_data.example_condition[0]").Exists().ErrorIsNil(t)
	check.InPlan(testPresent.Plan).That("terraform_data.example_condition[0]").Key("input").HasValue("example_condition").ErrorIsNil(t)

	// // Now check that the resource is NOT present in the test_notpresent plan...
	check.InPlan(testNotPresent.Plan).That("terraform_data.example_condition[0]").DoesNotExist().ErrorIsNil(t)
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

	for key, value := range vars["example_for_each"].(map[string]string) {
		instanceName := fmt.Sprintf("terraform_data.example_for_each[\"%s\"]", key)

		check.InPlan(test.Plan).That(instanceName).Exists().ErrorIsNil(t)
		check.InPlan(test.Plan).That(instanceName).Key("input").HasValue(value).ErrorIsNil(t)
	}
}
