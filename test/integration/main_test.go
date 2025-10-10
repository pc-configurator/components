//go:build integration

package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
}

func Test_Integration(t *testing.T) {
	suite.Run(t, &Suite{})
}

// Before all tests
func (s *Suite) SetupSuite() {
	fmt.Println("hello world")
}

// After all tests
func (s *Suite) TearDownSuite() {}

// Before each test
func (s *Suite) SetupTest() {}

// After each test
func (s *Suite) TearDownTest() {}
