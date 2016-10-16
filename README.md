# Go Suite

[![Build Status](https://semaphoreci.com/api/v1/pavlikus/gosuite/branches/master/badge.svg)](https://semaphoreci.com/pavlikus/gosuite) [![codecov](https://codecov.io/gh/pavlo/gosuite/branch/master/graph/badge.svg)](https://codecov.io/gh/pavlo/gosuite)


The support for test suites for Golang 1.7 and later.
 
Golang 1.7 featured [Subtests](https://golang.org/pkg/testing/) that allowed you to group tests in order to share common setup and teardown logic. While that was a great addition to the `testing` package, it was a bit clunky syntactically. The GoSuite package leverages Golang's 1.7 Subtests feature, defines a simple `TestSuite` interface and runs test cases inside of them keeping setup/teardown logic for the whole suite and for single cases in place.

## Quick Start

To start with, create a struct with the four methods implemented:

```go
type MyTestSuite struct {
    // DB connection
    // etc
}

// SetUpSuite is called once before the very first test in suite runs
func (s *MyTestSuite) SetUpSuite() {
}

// TearDownSuite is called once after thevery last test in suite runs
func (s *MyTestSuite) TearDownSuite() {
}

// SetUp is called before each test method
func (s *MyTestSuite) SetUp() {
}

// TearDown is called after each test method
func (s *MyTestSuite) TearDown() {
}
```
 
Then add one or more test methods to it, prefixing them with `Test` prefix:

```go
func (s *MyTestSuite) TestMyFirstTestCase(t *testing.T) {
    if !someJob {
        t.Fail("Unexpected failure!")
    }
}

```

Almost done! The only piece that remains is to run the suite! You do this by calling the `Run` method. Note, the enclosing `TestIt` method is a normal testing method you usually write in Go, nothing fancy at all!

```go
func TestIt(t *testing.T) {
	Run(t, &MyTestSuite{})
}
```

## Complete Example

The complete example is shown to help you to see the whole thing on the same page. Note, it leverages the [Is](https://github.com/tylerb/is) package for assertions... the package is great though indeed it is not required to use with Go Suite. *The example however demonstrates a slick technique making the assertion methods available on the suite itself!* 

```go

import gosuite

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

func TestIt(t *testing.T) {
    s := &Suite{Is: is.New(s.t)}
	gosuite.Run(t, s)
	
	s.Equal(1, s.setUpSuiteCalledTimes)
	s.Equal(1, s.tearDownSuiteCalledTimes)
	s.Equal(2, s.setUpCalledTimes)
	s.Equal(2, s.tearDownUpCalledTimes)
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

```

Running it with `go test -v` would emit this:

```
> go test -v

=== RUN   TestIt
=== RUN   TestIt/TestFirstTestMethod
=== RUN   TestIt/TestSecondTestMethod
--- PASS: TestIt (0.00s)
    --- PASS: TestIt/TestFirstTestMethod (0.00s)
    --- PASS: TestIt/TestSecondTestMethod (0.00s)
PASS
ok  	github.com/pavlo/gosuite	0.009s
Success: Tests passed.
```


## License

`Go Suite` is released under the [MIT License](http://www.opensource.org/licenses/MIT).