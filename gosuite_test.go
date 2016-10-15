package gosuite

import (
	"gopkg.in/tylerb/is.v1"
	"testing"
)

func TestIt(t *testing.T) {
	s := &Suite{}
	Run(t, s)

	s.Equal(1, s.setUpSuiteCalledTimes)
	s.Equal(1, s.tearDownSuiteCalledTimes)
	s.Equal(2, s.setUpCalledTimes)
	s.Equal(2, s.tearDownUpCalledTimes)
}

type Suite struct {
	*is.Is
	setUpSuiteCalledTimes    int
	tearDownSuiteCalledTimes int
	setUpCalledTimes         int
	tearDownUpCalledTimes    int
}

func (s *Suite) SetUpSuite(t *testing.T) {
	s.Is = is.New(t)
	s.setUpSuiteCalledTimes++
}

func (s *Suite) TearDownSuite() {
	s.tearDownSuiteCalledTimes++
}

func (s *Suite) SetUp() {
	s.setUpCalledTimes++
}

func (s *Suite) TearDown() {
	s.tearDownUpCalledTimes++
}

func (s *Suite) GSTFirstTestMethod(t *testing.T) {
	s.Equal(1, s.setUpSuiteCalledTimes)
	s.Equal(0, s.tearDownSuiteCalledTimes)
	s.Equal(1, s.setUpCalledTimes)
	s.Equal(0, s.tearDownUpCalledTimes)
}

func (s *Suite) GSTSecondTestMethod(t *testing.T) {
	s.Equal(1, s.setUpSuiteCalledTimes)
	s.Equal(0, s.tearDownSuiteCalledTimes)
	s.Equal(2, s.setUpCalledTimes)
	s.Equal(1, s.tearDownUpCalledTimes)
}

func (s *Suite) TestFooMethod(t *testing.T) {
	t.Fatal("Should not be called as it does not start with GST prefix!")
}
