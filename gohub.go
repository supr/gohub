package gohub

import (
	"json"
	"fmt"
	"io/ioutil"
	"http"
	"time"
	"os"
	"url"
)

const GH_API_ROOT = "https://api.github.com"

type GoHub struct {
	user     string
	password string
	client   *http.Client
	apiHost  string
}

type Timestamp time.Time
type NullableString string

type Link struct {
	Href string `json:"href"`
}

type Links struct {
	Self           Link `json:"self"`
	Comments       Link `json:"comments"`
	ReviewComments Link `json:"comments"`
	Html           Link `json:"html"`
}

type PullRequests struct {
	Url       NullableString `json:"url"`
	HtmlUrl   NullableString `json:"html_url"`
	DiffUrl   NullableString `json:"diff_url"`
	PatchUrl  NullableString `json:"patch_url"`
	IssueUrl  NullableString `json:"issue_url"`
	Number    int            `json:"number"`
	State     NullableString `json:"state"`
	Title     NullableString `json:"title"`
	Body      NullableString `json:"body"`
	CreatedAt Timestamp      `json:"created_at"`
	UpdatedAt Timestamp      `json:"updated_at"`
	Links     Links          `json:"_links"`
	g         *GoHub
}

type PullRequest struct {
	Url       NullableString    `json:"url"`
	HtmlUrl   NullableString    `json:"html_url"`
	DiffUrl   NullableString    `json:"diff_url"`
	PatchUrl  NullableString    `json:"patch_url"`
	IssueUrl  NullableString    `json:"issue_url"`
	Number    int               `json:"number"`
	State     NullableString    `json:"state"`
	Title     NullableString    `json:"title"`
	Body      NullableString    `json:"body"`
	CreatedAt Timestamp         `json:"created_at"`
	UpdatedAt Timestamp         `json:"updated_at"`
	Head      PullRequestMarker `json:"head"`
	Base      PullRequestMarker `json:"base"`
	Links     Links             `json:"_links"`
	g         *GoHub
}

type PullRequestMarker struct {
	Label NullableString `json:"label"`
	Ref   NullableString `json:"ref"`
	Sha   NullableString `json:"sha"`
	Repo  Repository     `json:"repo"`
	User  User           `json:"user"`
}

type Repository struct {
	Url          NullableString `json:"url"`
	HtmlUrl      NullableString `json:"html_url"`
	CloneUrl     NullableString `json:"clone_url"`
	GitUrl       NullableString `json:"git_url"`
	SshUrl       NullableString `json:"ssh_url"`
	SvnUrl       NullableString `json:"svn_url"`
	Owner        User           `json:"owner"`
	Name         NullableString `json:"name"`
	Description  NullableString `json:"description"`
	Homepage     NullableString `json:"homepage"`
	Language     NullableString `json:"language"`
	Private      bool           `json:"private"`
	Fork         bool           `json:"fork"`
	Forks        int            `json:"forks"`
	Watchers     int            `json:"watchers"`
	Size         int            `json:"size"`
	MasterBranch NullableString `json:"master_branch"`
	OpenIssues   int            `json:"open_issues"`
	PushedAt     Timestamp      `json:"pushed_at"`
	CreatedAt    Timestamp      `json:"created_at"`
}

type User struct {
	Login     string `json:"login"`
	Id        int    `json:"id"`
	AvatarUrl string `json:"avatar_url"`
	Url       string `json:"url"`
}

type PullRequestMergeResponse struct {
	Sha     NullableString `json:"sha"`
	Merged  bool           `json:"merged"`
	Message string         `json:"message"`
}

type Comment struct {
	CreatedAt Timestamp `json:"created_at"`
	UpdatedAt Timestamp `json:"updated_at"`
	User      User      `json:"user"`
	CommitId  string    `json:"commit_id"`
	Position  int       `json:"position"`
	Body      string    `json:"body"`
}

func New(user, password, api_root string) *GoHub {
	if api_root == "" {
		return &GoHub{user, password, &http.Client{}, GH_API_ROOT}
	}

	return &GoHub{user, password, &http.Client{}, api_root}
}

func (g *GoHub) makeAuthRequest(method, url_ string) (*http.Request, os.Error) {

	req, err := http.NewRequest(method, url_, nil)
	if err != nil {
		return nil, err
	}

	if g.user != "" && g.password != "" {
		req.SetBasicAuth(g.user, g.password)
	}

	return req, nil
}

func (g *GoHub) makeGetRequest(url_ string) ([]byte, os.Error) {
	req, err := g.makeAuthRequest("GET", url_)
	if err != nil {
		return nil, err
	}

	resp, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}

	outbuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return outbuf, err
}

func (g *GoHub) makePutRequest(url_ string) ([]byte, os.Error) {
	req, err := g.makeAuthRequest("PUT", url_)
	if err != nil {
		return nil, err
	}

	resp, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}

	outbuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return outbuf, err
}

func (g *GoHub) makePostRequest(url_ string, params url.Values) ([]byte, os.Error) {
	req, err := g.makeAuthRequest("POST", url_)
	if err != nil {
		return nil, err
	}

	req.Form = params
	resp, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}

	outbuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return outbuf, err
}

func (g *GoHub) PullRequest(user, repo string, id int) (*PullRequest, os.Error) {
	url_ := fmt.Sprintf("%v/repos/%v/%v/pulls/%v", g.apiHost, user, repo, id)
	out, err := g.makeGetRequest(url_)

	if err != nil {
		return nil, err
	}

	var pr PullRequest
	err = json.Unmarshal(out, &pr)
	if err != nil {
		return nil, err
	}

	pr.g = g
	return &pr, nil
}

func (g *GoHub) PullRequests(user, repo string) ([]PullRequests, os.Error) {
	url_ := fmt.Sprintf("%v/repos/%v/%v/pulls", g.apiHost, user, repo)
	out, err := g.makeGetRequest(url_)

	if err != nil {
		return nil, err
	}

	var prs []PullRequests
	err = json.Unmarshal(out, &prs)
	if err != nil {
		return nil, err
	}

	for _, v := range prs {
		v.g = g
	}

	return prs, nil
}

func (p *PullRequest) Merge() (*PullRequestMergeResponse, os.Error) {
	url_ := fmt.Sprintf("%v/repos/%v/%v/pulls/%v/merge", p.g.apiHost, p.Head.Repo.Name, p.Head.Repo.Owner.Login, p.Number)
	out, err := p.g.makePutRequest(url_)

	if err != nil {
		return nil, err
	}

	var pullRequestMergeResponse PullRequestMergeResponse
	err = json.Unmarshal(out, &pullRequestMergeResponse)
	if err != nil {
		return nil, err
	}

	if !pullRequestMergeResponse.Merged {
		return nil, os.NewError(pullRequestMergeResponse.Message)
	}

	return &pullRequestMergeResponse, nil
}

func (p *PullRequest) Comments() ([]Comment, os.Error) {
	url_ := fmt.Sprintf("%v/repos/%v/%v/pulls/%v/comments", p.g.apiHost, p.Head.Repo.Name, p.Head.Repo.Owner.Login, p.Number)
	out, err := p.g.makeGetRequest(url_)

	if err != nil {
		return nil, err
	}

	var comments []Comment
	err = json.Unmarshal(out, &comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (p *PullRequest) IssueComments() ([]Comment, os.Error) {
	url_ := fmt.Sprintf("%v/repos/%v/%v/issues/%v/comments", p.g.apiHost, p.Head.Repo.Name, p.Head.Repo.Owner.Login, p.Number)
	out, err := p.g.makeGetRequest(url_)

	if err != nil {
		return nil, err
	}

	var comments []Comment
	err = json.Unmarshal(out, &comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (p *PullRequest) NewIssueComment(body string) (*Comment, os.Error) {
	url_ := fmt.Sprintf("%v/repos/%v/%v/issues/%v/comments", p.g.apiHost, p.Head.Repo.Name, p.Head.Repo.Owner.Login, p.Number)
	params := url.Values{}
	params.Add("body", body)
	out, err := p.g.makePostRequest(url_, params)
	if err != nil {
		return nil, err
	}

	var comment Comment
	err = json.Unmarshal(out, &comment)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (ts *Timestamp) UnmarshalJSON(data []byte) os.Error {
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		//fail Gracefully
		return os.NewError("cannot un-marshal non-string to timestamp")
	}
	data = data[1 : len(data)-1]
	t, err := time.Parse(time.RFC3339, string(data))
	if err != nil {
		return err
	}

	*ts = Timestamp(*t)
	return nil
}

func (s *NullableString) UnmarshalJSON(data []byte) os.Error {
	if len(data) < 2 || string(data) == "null" {
		*s = ""
		return nil
	}
	var out string
	err := json.Unmarshal(data, &out)
	if err != nil {
		return err
	}

	*s = NullableString(out)
	return nil
}
