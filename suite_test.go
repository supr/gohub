package gohub_test

import (
	"http"
	"url"
	. "launchpad.net/gocheck"
	"fmt"
	"os"
	"testing"
	"time"
)

func Test(t *testing.T) {
	TestingT(t)
}

type HTTPSuite struct {}

var testServer = NewTestHTTPServer("http://localhost:4444", 5e9)

func (s *HTTPSuite) SetUpSuite(c *C) {
	testServer.Start()
}

func (s *HTTPSuite) TearDownTest(c *C) {
	testServer.FlushRequests()
}

type TestHTTPServer struct {
	URL string
	Timeout int64
	started bool
	request chan *http.Request
	response chan *testResponse
	pending chan bool
}

type testResponse struct {
	Status int
	Headers map[string]string
	Body string
}

func NewTestHTTPServer(url string, timeout int64) *TestHTTPServer {
	return &TestHTTPServer{URL: url, Timeout: timeout}
}

func (s *TestHTTPServer) Start() {
	if s.started {
		return
	}
	s.started = true

	s.request = make(chan *http.Request, 64)
	s.response = make(chan *testResponse, 64)
	s.pending = make(chan bool, 64)

	url,_ := url.Parse(s.URL)
	go http.ListenAndServe(url.Host, s)

	s.PrepareResponse(202, nil, "Nothing.")
	fmt.Fprintf(os.Stderr, "\nWaiting for the fake server to be up...")
	for {
		resp, err := http.Get(s.URL)
		if err == nil && resp.StatusCode == 202 {
			break
		}

		time.Sleep(1e8)
	}

	fmt.Fprintf(os.Stderr, "Done\n")
	s.WaitRequest()
}

func (s *TestHTTPServer) FlushRequests() {
	for {
		select {
			case <-s.request:
			default:
				return
		}
	}
}

func (s *TestHTTPServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.request <- req
	var resp *testResponse
	select {
		case resp = <-s.response:
		case <-time.After(s.Timeout):
			fmt.Fprintf(os.Stderr, "ERROR: Timeout waiting for the test to provide any response\n")
			resp = &testResponse{500, nil, ""}
	}

	if resp.Headers != nil {
		h := w.Header()
		for k,v := range resp.Headers {
			h.Set(k, v)
		}
	}

	if resp.Status != 0 {
		w.WriteHeader(resp.Status)
	}

	w.Write([]byte(resp.Body))
}

func (s *TestHTTPServer) WaitRequest() *http.Request {
	select {
		case req := <-s.request:
			req.ParseForm()
			return req
		case <-time.After(s.Timeout):
			panic("timeout waiting for request")
	}
	panic("unreached")
}

func (s *TestHTTPServer) PrepareResponse(status int, headers map[string]string, body string) {
	s.response <- &testResponse{status, headers, body}
}
