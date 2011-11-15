package gohub_test

import (
	"gohub"
	"launchpad.net/gocheck"
)

var _ = gocheck.Suite(&S{})

type S struct {
	HTTPSuite
	g *gohub.GoHub
}

func (s *S) SetUpSuite(c *gocheck.C) {
	s.HTTPSuite.SetUpSuite(c)
	s.g = gohub.New("foouser", "foopass", "http://localhost:4444")
}

func (s *S) TestPullRequests(c *gocheck.C) {
	testServer.PrepareResponse(200, nil, TestPullRequestsOK)

	resp, err := s.g.PullRequests("foo", "bar")
	req := testServer.WaitRequest()

	c.Assert(err, gocheck.IsNil)
	c.Assert(req.Method, gocheck.Equals, "GET")
	c.Assert(len(resp), gocheck.Not(gocheck.Equals), 0)
}

func (s *S) TestPullRequest(c *gocheck.C) {
	testServer.PrepareResponse(200, nil, TestPullRequestOK)

	resp, err := s.g.PullRequest("foo", "bar" , 29)
	req := testServer.WaitRequest()

	c.Assert(err, gocheck.IsNil)
	c.Assert(req.Method, gocheck.Equals, "GET")
	c.Assert(resp.Number, gocheck.Equals, 29)
}

func (s *S) TestMergeOK(c *gocheck.C) {
	testServer.FlushRequests()
	testServer.PrepareResponse(200, nil, TestPullRequestOK)
	resp, err := s.g.PullRequest("foo", "bar", 29)

	req := testServer.WaitRequest()

	testServer.PrepareResponse(200, nil, TestMergeSuccessOK)
	mr, err := resp.Merge()
	req = testServer.WaitRequest()

	c.Assert(err, gocheck.IsNil)
	c.Assert(req.Method, gocheck.Equals, "PUT")

	c.Assert(mr.Sha, gocheck.Equals, gohub.NullableString("6dcb09b5b57875f334f61aebed695e2e4193db5e"))
	c.Assert(mr.Merged, gocheck.Equals, true)
	c.Assert(mr.Message, gocheck.Equals, "PullRequestsuccessfullymerged")
}

func (s *S) TestMergeFail(c *gocheck.C) {
	testServer.FlushRequests()
	testServer.PrepareResponse(200, nil, TestPullRequestOK)
	resp, err := s.g.PullRequest("foo", "bar", 29)

	req := testServer.WaitRequest()

	testServer.PrepareResponse(200, nil, TestMergeFailureOK)
	mr, err := resp.Merge()
	req = testServer.WaitRequest()

	c.Assert(err, gocheck.Not(gocheck.IsNil))
	c.Assert(req.Method, gocheck.Equals, "PUT")

	c.Assert(mr, gocheck.IsNil)
}
