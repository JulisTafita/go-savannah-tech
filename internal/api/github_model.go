//go:generate mockgen -source=repository.go -destination=mock/repository.go
package api

import (
	"time"
)

type Repository struct {
	ID                       int                  `json:"id"`
	NodeID                   string               `json:"node_id"`
	Name                     string               `json:"name"`
	FullName                 string               `json:"full_name"`
	Private                  bool                 `json:"private"`
	Owner                    RepositoryOwner      `json:"owner"`
	HTMLURL                  string               `json:"html_url"`
	Description              string               `json:"description"`
	Fork                     bool                 `json:"fork"`
	URL                      string               `json:"url"`
	ForksURL                 string               `json:"forks_url"`
	KeysURL                  string               `json:"keys_url"`
	CollaboratorsURL         string               `json:"collaborators_url"`
	TeamsURL                 string               `json:"teams_url"`
	HooksURL                 string               `json:"hooks_url"`
	IssueEventsURL           string               `json:"issue_events_url"`
	EventsURL                string               `json:"events_url"`
	AssigneesURL             string               `json:"assignees_url"`
	BranchesURL              string               `json:"branches_url"`
	TagsURL                  string               `json:"tags_url"`
	BlobsURL                 string               `json:"blobs_url"`
	GitTagsURL               string               `json:"git_tags_url"`
	GitRefsURL               string               `json:"git_refs_url"`
	TreesURL                 string               `json:"trees_url"`
	StatusesURL              string               `json:"statuses_url"`
	LanguagesURL             string               `json:"languages_url"`
	StargazersURL            string               `json:"stargazers_url"`
	ContributorsURL          string               `json:"contributors_url"`
	SubscribersURL           string               `json:"subscribers_url"`
	SubscriptionURL          string               `json:"subscription_url"`
	CommitsURL               string               `json:"commits_url"`
	GitCommitsURL            string               `json:"git_commits_url"`
	CommentsURL              string               `json:"comments_url"`
	IssueCommentURL          string               `json:"issue_comment_url"`
	ContentsURL              string               `json:"contents_url"`
	CompareURL               string               `json:"compare_url"`
	MergesURL                string               `json:"merges_url"`
	ArchiveURL               string               `json:"archive_url"`
	DownloadsURL             string               `json:"downloads_url"`
	IssuesURL                string               `json:"issues_url"`
	PullsURL                 string               `json:"pulls_url"`
	MilestonesURL            string               `json:"milestones_url"`
	NotificationsURL         string               `json:"notifications_url"`
	LabelsURL                string               `json:"labels_url"`
	ReleasesURL              string               `json:"releases_url"`
	DeploymentsURL           string               `json:"deployments_url"`
	CreatedAt                time.Time            `json:"created_at"`
	UpdatedAt                time.Time            `json:"updated_at"`
	PushedAt                 time.Time            `json:"pushed_at"`
	GitURL                   string               `json:"git_url"`
	SSHURL                   string               `json:"ssh_url"`
	CloneURL                 string               `json:"clone_url"`
	SvnURL                   string               `json:"svn_url"`
	Homepage                 any                  `json:"homepage"`
	Size                     int                  `json:"size"`
	StargazersCount          int                  `json:"stargazers_count"`
	WatchersCount            int                  `json:"watchers_count"`
	Language                 string               `json:"language"`
	HasIssues                bool                 `json:"has_issues"`
	HasProjects              bool                 `json:"has_projects"`
	HasDownloads             bool                 `json:"has_downloads"`
	HasWiki                  bool                 `json:"has_wiki"`
	HasPages                 bool                 `json:"has_pages"`
	HasDiscussions           bool                 `json:"has_discussions"`
	ForksCount               int                  `json:"forks_count"`
	MirrorURL                any                  `json:"mirror_url"`
	Archived                 bool                 `json:"archived"`
	Disabled                 bool                 `json:"disabled"`
	OpenIssuesCount          int                  `json:"open_issues_count"`
	License                  any                  `json:"license"`
	AllowForking             bool                 `json:"allow_forking"`
	IsTemplate               bool                 `json:"is_template"`
	WebCommitSignoffRequired bool                 `json:"web_commit_signoff_required"`
	Topics                   []any                `json:"topics"`
	Visibility               string               `json:"visibility"`
	Forks                    int                  `json:"forks"`
	OpenIssues               int                  `json:"open_issues"`
	Watchers                 int                  `json:"watchers"`
	DefaultBranch            string               `json:"default_branch"`
	Permissions              RepositoryPermission `json:"permissions"`
}

type RepositoryOwner struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type RepositoryPermission struct {
	Admin    bool `json:"admin"`
	Maintain bool `json:"maintain"`
	Push     bool `json:"push"`
	Triage   bool `json:"triage"`
	Pull     bool `json:"pull"`
}

type RepositorySearchResult struct {
	TotalCount        int          `json:"total_count"`
	IncompleteResults bool         `json:"incomplete_results"`
	Items             []Repository `json:"items"`
}

type RepositoryCommit struct {
	Sha         string                   `json:"sha"`
	NodeID      string                   `json:"node_id"`
	Commit      Commit                   `json:"commit"`
	URL         string                   `json:"url"`
	HTMLURL     string                   `json:"html_url"`
	CommentsURL string                   `json:"comments_url"`
	Author      RepositoryOwner          `json:"author"`
	Committer   RepositoryCommitter      `json:"committer"`
	Parents     []RepositoryCommitParent `json:"parents"`
}

type Commit struct {
	Author       CommitAuthor       `json:"author"`
	Committer    CommitCommitter    `json:"committer"`
	Message      string             `json:"message"`
	Tree         CommitTree         `json:"tree"`
	URL          string             `json:"url"`
	CommentCount int                `json:"comment_count"`
	Verification CommitVerification `json:"verification"`
}

type CommitTree struct {
	Sha string `json:"sha"`
	URL string `json:"url"`
}

type CommitVerification struct {
	Verified  bool   `json:"verified"`
	Reason    string `json:"reason"`
	Signature any    `json:"signature"`
	Payload   any    `json:"payload"`
}

type CommitCommitter struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Date  time.Time `json:"date"`
}

type CommitAuthor struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Date  time.Time `json:"date"`
}

type RepositoryCommitter struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type RepositoryCommitParent struct {
	Sha     string `json:"sha"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}

type Profile struct {
	Login                   string      `json:"login"`
	ID                      int         `json:"id"`
	NodeID                  string      `json:"node_id"`
	AvatarURL               string      `json:"avatar_url"`
	GravatarID              string      `json:"gravatar_id"`
	URL                     string      `json:"url"`
	HTMLURL                 string      `json:"html_url"`
	FollowersURL            string      `json:"followers_url"`
	FollowingURL            string      `json:"following_url"`
	GistsURL                string      `json:"gists_url"`
	StarredURL              string      `json:"starred_url"`
	SubscriptionsURL        string      `json:"subscriptions_url"`
	OrganizationsURL        string      `json:"organizations_url"`
	ReposURL                string      `json:"repos_url"`
	EventsURL               string      `json:"events_url"`
	ReceivedEventsURL       string      `json:"received_events_url"`
	Type                    string      `json:"type"`
	SiteAdmin               bool        `json:"site_admin"`
	Name                    string      `json:"name"`
	Company                 any         `json:"company"`
	Blog                    string      `json:"blog"`
	Location                string      `json:"location"`
	Email                   any         `json:"email"`
	Hireable                any         `json:"hireable"`
	Bio                     string      `json:"bio"`
	TwitterUsername         any         `json:"twitter_username"`
	PublicRepos             int         `json:"public_repos"`
	PublicGists             int         `json:"public_gists"`
	Followers               int         `json:"followers"`
	Following               int         `json:"following"`
	CreatedAt               time.Time   `json:"created_at"`
	UpdatedAt               time.Time   `json:"updated_at"`
	PrivateGists            int         `json:"private_gists"`
	TotalPrivateRepos       int         `json:"total_private_repos"`
	OwnedPrivateRepos       int         `json:"owned_private_repos"`
	DiskUsage               int         `json:"disk_usage"`
	Collaborators           int         `json:"collaborators"`
	TwoFactorAuthentication bool        `json:"two_factor_authentication"`
	Plan                    ProfilePlan `json:"plan"`
}

type ProfilePlan struct {
	Name          string `json:"name"`
	Space         int    `json:"space"`
	Collaborators int    `json:"collaborators"`
	PrivateRepos  int    `json:"private_repos"`
}
