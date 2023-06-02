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


func (s *MySuite) TestFindMinServer(c *C) {
  serversPool = nil
  serversPool = []*Server{
    {URL: "url_of_server_1", ConnCnt: 10, Healthy: true},
    {URL: "url_of_server_2", ConnCnt: 5, Healthy: true},
    {URL: "url_of_server_3", ConnCnt: 30, Healthy: true},
  }
  c.Assert(minServerIndex(), Equals, 1)
  serversPool = []*Server{
    {URL: "url_of_server_1", ConnCnt: 10, Healthy: false},
    {URL: "url_of_server_2", ConnCnt: 20, Healthy: false},
    {URL: "url_of_server_3", ConnCnt: 30, Healthy: false},
  }
  c.Assert(minServerIndex(), Equals, -1)

  serversPool = []*Server{
    {URL: "url_of_server_1", ConnCnt: 10, Healthy: true},
    {URL: "url_of_server_2", ConnCnt: 20, Healthy: true},
    {URL: "url_of_server_3", ConnCnt: 30, Healthy: true},
  }
  c.Assert(minServerIndex(), Equals, 0)
}