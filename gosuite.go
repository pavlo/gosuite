package gosuite

import (
	"reflect"
	"strings"
	"testing"
)

const SuiteTestMethodPrefix = "GST" // stands for GoSuiteTest

type TestSuite interface {
	// SetUpSuite is called once before the very first test in suite runs
	SetUpSuite(t *testing.T)

	// TearDownSuite is called once after thevery last test in suite runs
	TearDownSuite()

	// SetUp is called before each test method
	SetUp()

	// TearDown is called after each test method
	TearDown()
}

func Run(t *testing.T, suite TestSuite) {
	suite.SetUpSuite(t)
	defer suite.TearDownSuite()

	suiteType := reflect.TypeOf(suite)
	for i := 0; i < suiteType.NumMethod(); i++ {
		m := suiteType.Method(i)
		if strings.HasPrefix(m.Name, SuiteTestMethodPrefix) {
			t.Run(m.Name, func(t *testing.T) {
				suite.SetUp()
				defer suite.TearDown()

				in := []reflect.Value{reflect.ValueOf(suite), reflect.ValueOf(t)}
				m.Func.Call(in)
			})
		}
	}
}
