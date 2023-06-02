package main

import (
	"testing"

	"github.com/jarcoal/httpmock"
	. "gopkg.in/check.v1"
)


func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) SetUpSuite(c *C) {
  httpmock.Activate()
}

func (s *MySuite) TearDownSuite(c *C) {
  httpmock.DeactivateAndReset()
}
