/*
gosuite package provides a simple and tiny tool that brings the support of test suites to Go 1.7 Subtests addition to "testing".

A test suite is an abstraction that allows you to group certain test cases together as well as allowing you to perform setup/teardown
logic for each of test cases as well as the setup/teardown stuff for the suite itself.

This is useful, for instance, in cases where you need to set up database schema before your suite as well as truncate the database
tables before each test case so each of them is run against an empty database.
*/
package gosuite

import (
	"reflect"
	"strings"
	"testing"
)

const suiteTestMethodPrefix = "Test"

// TestSuite is an interface where you define suite and test case preparation and tear down logic.
type TestSuite interface {
	// SetUpSuite is called once before the very first test in suite runs
	SetUpSuite()

	// TearDownSuite is called once after thevery last test in suite runs
	TearDownSuite()

	// SetUp is called before each test method
	SetUp()

	// TearDown is called after each test method
	TearDown()
}

/*
Run sets up the suite, runs its test cases and tears it down:
    1. Calls `suite.SetUpSuite`
    2. Seeks for any methods that have `Test` prefix, for each of them it:
      a. Calls `SetUp`
      b. Calls the test method itself
      c. Calls `TearDown`
    3. Calls `suite.TearDownSuite`
*/
func Run(t *testing.T, suite TestSuite) {
	suite.SetUpSuite()
	defer suite.TearDownSuite()

	suiteType := reflect.TypeOf(suite)
	for i := 0; i < suiteType.NumMethod(); i++ {
		m := suiteType.Method(i)
		if strings.HasPrefix(m.Name, suiteTestMethodPrefix) {
			t.Run(m.Name, func(t *testing.T) {
				suite.SetUp()
				defer suite.TearDown()

				in := []reflect.Value{reflect.ValueOf(suite), reflect.ValueOf(t)}
				m.Func.Call(in)
			})
		}
	}
}
