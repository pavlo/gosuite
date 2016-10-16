package gosuite

import (
	"reflect"
	"strings"
	"testing"
)

// SuiteTestMethodPrefix specifies the prefix each test case method in a suite should have. This is the marker, the `Run` method will use to figure out whether to run this particular case method or not
const suiteTestMethodPrefix = "Test"

// TestSuite an interface that describes a test suite
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

// Run - runs the suite:
// 1. Calls `SetUpSuite`
// 2. Seeks for any methods that have `Test` prefix, for each of them it:
// 2a) Calls `SetUp`
// 2b) Calls the method itself
// 2c) Calls `TearDown`
// 3. Calls `TearDownSuite`
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
