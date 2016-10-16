package gosuite

import (
	"testing"

	"gopkg.in/tylerb/is.v1"
)

func TestIt(t *testing.T) {
	s := &Suite{Is: is.New(t)}
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

func (s *Suite) SetUpSuite() {
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

func (s *Suite) TestFirstTestMethod(t *testing.T) {
	s.Equal(1, s.setUpSuiteCalledTimes)
	s.Equal(0, s.tearDownSuiteCalledTimes)
	s.Equal(1, s.setUpCalledTimes)
	s.Equal(0, s.tearDownUpCalledTimes)
}

func (s *Suite) TestSecondTestMethod(t *testing.T) {
	s.Equal(1, s.setUpSuiteCalledTimes)
	s.Equal(0, s.tearDownSuiteCalledTimes)
	s.Equal(2, s.setUpCalledTimes)
	s.Equal(1, s.tearDownUpCalledTimes)
}
