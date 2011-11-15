package gohub_test

const TestPullRequestsOK = `
[{"_links":{"html":{"href":"https://github.com/github/developer.github.com/pull/29"},"review_comments":{"href":"https://api.github.com/repos/github/developer.github.com/pulls/29/comments"},"comments":{"href":"https://api.github.com/repos/github/developer.github.com/issues/29/comments"},"self":{"href":"https://api.github.com/repos/github/developer.github.com/pulls/29"}},"title":"Update libraries information with Ruby wrapper for GitHub v3 API","user":{"gravatar_id":"fb249d5564a0f8c5a0b69e0d71d58949","url":"https://api.github.com/users/peter-murach","avatar_url":"https://secure.gravatar.com/avatar/fb249d5564a0f8c5a0b69e0d71d58949?d=https://a248.e.akamai.net/assets.github.com%2Fimages%2Fgravatars%2Fgravatar-140.png","login":"peter-murach","id":444312},"url":"https://api.github.com/repos/github/developer.github.com/pulls/29","issue_url":"https://github.com/github/developer.github.com/issues/29","merged_at":null,"created_at":"2011-10-22T22:07:23Z","state":"open","html_url":"https://github.com/github/developer.github.com/pull/29","closed_at":null,"body":"","updated_at":"2011-11-02T17:39:41Z","number":29,"patch_url":"https://github.com/github/developer.github.com/pull/29.patch","id":425184,"diff_url":"https://github.com/github/developer.github.com/pull/29.diff"}]`

const TestPullRequestOK = `{"merged_by":null,"title":"Update libraries information with Ruby wrapper for GitHub v3 API","patch_url":"https://github.com/github/developer.github.com/pull/29.patch","user":{"gravatar_id":"fb249d5564a0f8c5a0b69e0d71d58949","url":"https://api.github.com/users/peter-murach","login":"peter-murach","avatar_url":"https://secure.gravatar.com/avatar/fb249d5564a0f8c5a0b69e0d71d58949?d=https://a248.e.akamai.net/assets.github.com%2Fimages%2Fgravatars%2Fgravatar-140.png","id":444312},"diff_url":"https://github.com/github/developer.github.com/pull/29.diff","url":"https://api.github.com/repos/github/developer.github.com/pulls/29","issue_url":"https://github.com/github/developer.github.com/issues/29","merged_at":null,"created_at":"2011-10-22T22:07:23Z","_links":{"html":{"href":"https://github.com/github/developer.github.com/pull/29"},"comments":{"href":"https://api.github.com/repos/github/developer.github.com/issues/29/comments"},"review_comments":{"href":"https://api.github.com/repos/github/developer.github.com/pulls/29/comments"},"self":{"href":"https://api.github.com/repos/github/developer.github.com/pulls/29"}},"state":"open","html_url":"https://github.com/github/developer.github.com/pull/29","comments":0,"deletions":2,"closed_at":null,"body":"","additions":16,"changed_files":3,"updated_at":"2011-11-02T17:39:41Z","merged":false,"review_comments":0,"base":{"user":{"gravatar_id":"61024896f291303615bcd4f7a0dcfb74","url":"https://api.github.com/users/github","login":"github","avatar_url":"https://secure.gravatar.com/avatar/61024896f291303615bcd4f7a0dcfb74?d=https://a248.e.akamai.net/assets.github.com%2Fimages%2Fgravatars%2Fgravatar-orgs.png","id":9919},"ref":"master","sha":"63502d1b4232156551d33ad8a91150cfdd1874fb","label":"github:master","repo":{"watchers":106,"clone_url":"https://github.com/github/developer.github.com.git","forks":47,"url":"https://api.github.com/repos/github/developer.github.com","description":"GitHub API documentation","ssh_url":"git@github.com:github/developer.github.com.git","created_at":"2011-04-26T19:20:56Z","html_url":"https://github.com/github/developer.github.com","open_issues":5,"svn_url":"https://svn.github.com/github/developer.github.com","fork":false,"git_url":"git://github.com/github/developer.github.com.git","homepage":"http://developer.github.com","updated_at":"2011-11-09T15:05:20Z","private":false,"size":316,"language":"Ruby","pushed_at":"2011-11-09T15:05:20Z","owner":{"gravatar_id":"61024896f291303615bcd4f7a0dcfb74","url":"https://api.github.com/users/github","login":"github","avatar_url":"https://secure.gravatar.com/avatar/61024896f291303615bcd4f7a0dcfb74?d=https://a248.e.akamai.net/assets.github.com%2Fimages%2Fgravatars%2Fgravatar-orgs.png","id":9919},"name":"developer.github.com","master_branch":null,"id":1666784}},"number":29,"mergeable":false,"commits":2,"head":{"user":{"gravatar_id":"fb249d5564a0f8c5a0b69e0d71d58949","url":"https://api.github.com/users/peter-murach","login":"peter-murach","avatar_url":"https://secure.gravatar.com/avatar/fb249d5564a0f8c5a0b69e0d71d58949?d=https://a248.e.akamai.net/assets.github.com%2Fimages%2Fgravatars%2Fgravatar-140.png","id":444312},"ref":"master","sha":"12d12dcbe4b12bad6b7fe8cfcaf415ea07867932","label":"peter-murach:master","repo":{"watchers":1,"clone_url":"https://github.com/peter-murach/developer.github.com.git","forks":0,"url":"https://api.github.com/repos/peter-murach/developer.github.com","description":"GitHub API documentation","ssh_url":"git@github.com:peter-murach/developer.github.com.git","created_at":"2011-10-07T13:51:31Z","html_url":"https://github.com/peter-murach/developer.github.com","open_issues":0,"svn_url":"https://svn.github.com/peter-murach/developer.github.com","fork":true,"git_url":"git://github.com/peter-murach/developer.github.com.git","homepage":"http://developer.github.com","updated_at":"2011-10-22T21:55:29Z","private":false,"size":116,"language":"Ruby","pushed_at":"2011-10-22T21:55:25Z","owner":{"gravatar_id":"fb249d5564a0f8c5a0b69e0d71d58949","url":"https://api.github.com/users/peter-murach","login":"peter-murach","avatar_url":"https://secure.gravatar.com/avatar/fb249d5564a0f8c5a0b69e0d71d58949?d=https://a248.e.akamai.net/assets.github.com%2Fimages%2Fgravatars%2Fgravatar-140.png","id":444312},"name":"developer.github.com","master_branch":null,"id":2532691}},"id":425184}`

const TestMergeSuccessOK = `
{"sha":"6dcb09b5b57875f334f61aebed695e2e4193db5e","merged":true,"message":"PullRequestsuccessfullymerged"}
`

const TestMergeFailureOK = `
{"sha":null,"merged":false,"message":"failure"}
`

const TestCommentsOK = `
[{"updated_at":"2011-10-18T13:36:48Z","user":{"avatar_url":"https://secure.gravatar.com/avatar/b2ebc1aa12b24ade90b519c1ac059b63?d=https://a248.e.akamai.net/assets.github.com%2Fimages%2Fgravatars%2Fgravatar-140.png","login":"tillsc","url":"https://api.github.com/users/tillsc","gravatar_id":"b2ebc1aa12b24ade90b519c1ac059b63","id":311544},"url":"https://api.github.com/repos/rails/rails/issues/comments/2441740","created_at":"2011-10-18T13:36:48Z","body":"+1","id":2441740}]
`
