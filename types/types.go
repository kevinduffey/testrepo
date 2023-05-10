package types

// Describe whether all repositories have been selected or there's a selection involved
type RepositoryNonCharSelection string

// The level of permission to grant the access token for deployments and deployment statuses.
type Deployments string

// The level of permission to grant the access token to view and manage security events like code scanning alerts.
type SecurityNonCharEvents string

// The level of permission to grant the access token to search repositories, list collaborators, and access repository metadata.
type Metadata string

// The level of permission to grant the access token to manage team discussions and related comments.
type TeamNonCharDiscussions string

// The level of permission to grant the access token to manage the post-receive hooks for a repository.
type RepositoryNonCharHooks string

// The level of permission to grant the access token to view and manage secret scanning alerts.
type SecretNonCharScanningNonCharAlerts string

// The level of permission to grant the access token to manage the post-receive hooks for an organization.
type OrganizationNonCharHooks string

// The level of permission to grant the access token to manage organization secrets.
type OrganizationNonCharSecrets string

// The level of permission to grant the access token to view and manage GitHub Actions self-hosted runners available to an organization.
type OrganizationNonCharSelfNonCharHostedNonCharRunners string

// The level of permission to grant the access token to manage repository projects, columns, and cards.
type RepositoryNonCharProjects string

// The level of permission to grant the access token to manage repository secrets.
type Secrets string

// The level of permission to grant the access token to manage organization projects and projects beta (where available).
type OrganizationNonCharProjects string

// The level of permission to grant the access token for GitHub Actions workflows, workflow runs, and artifacts.
type Actions string

// The level of permission to grant the access token for checks on code.
type Checks string

// The level of permission to grant the access token for repository contents, commits, branches, downloads, releases, and merges.
type Contents string

// The level of permission to grant the access token to manage just a single file.
type SingleNonCharFile string

// The level of permission to grant the access token for organization packages published to GitHub Packages.
type OrganizationNonCharPackages string

// The level of permission to grant the access token to update GitHub Actions workflow files.
type Workflows string

// The level of permission to grant the access token for packages published to GitHub Packages.
type Packages string

// The level of permission to grant the access token for managing repository environments.
type Environments string

// The level of permission to grant the access token for issues and related comments, assignees, labels, and milestones.
type Issues string

// The level of permission to grant the access token for repository creation, deletion, settings, teams, and collaborators creation.
type Administration string

// The level of permission to grant the access token for viewing an organization's plan.
type OrganizationNonCharPlan string

// The level of permission to grant the access token to manage access to an organization.
type OrganizationNonCharAdministration string

// The level of permission to grant the access token for organization teams and members.
type Members string

// The level of permission to grant the access token to retrieve Pages statuses, configuration, and builds, as well as create new builds.
type Pages string

// The level of permission to grant the access token to view and manage users blocked by the organization.
type OrganizationNonCharUserNonCharBlocking string

// The level of permission to grant the access token for pull requests and related comments, assignees, labels, milestones, and merges.
type PullNonCharRequests string

// The level of permission to grant the access token for commit statuses.
type Statuses string

// The level of permission to grant the access token to manage Dependabot alerts.
type VulnerabilityNonCharAlerts string

const (
	RepositoryNonCharSelectionAll      RepositoryNonCharSelection = "all"
	RepositoryNonCharSelectionSelected RepositoryNonCharSelection = "selected"
)

const (
	DeploymentsRead  Deployments = "read"
	DeploymentsWrite Deployments = "write"
)

const (
	SecurityNonCharEventsRead  SecurityNonCharEvents = "read"
	SecurityNonCharEventsWrite SecurityNonCharEvents = "write"
)

const (
	MetadataRead  Metadata = "read"
	MetadataWrite Metadata = "write"
)

const (
	TeamNonCharDiscussionsRead  TeamNonCharDiscussions = "read"
	TeamNonCharDiscussionsWrite TeamNonCharDiscussions = "write"
)

const (
	RepositoryNonCharHooksRead  RepositoryNonCharHooks = "read"
	RepositoryNonCharHooksWrite RepositoryNonCharHooks = "write"
)

const (
	SecretNonCharScanningNonCharAlertsRead  SecretNonCharScanningNonCharAlerts = "read"
	SecretNonCharScanningNonCharAlertsWrite SecretNonCharScanningNonCharAlerts = "write"
)

const (
	OrganizationNonCharHooksRead  OrganizationNonCharHooks = "read"
	OrganizationNonCharHooksWrite OrganizationNonCharHooks = "write"
)

const (
	OrganizationNonCharSecretsRead  OrganizationNonCharSecrets = "read"
	OrganizationNonCharSecretsWrite OrganizationNonCharSecrets = "write"
)

const (
	OrganizationNonCharSelfNonCharHostedNonCharRunnersRead  OrganizationNonCharSelfNonCharHostedNonCharRunners = "read"
	OrganizationNonCharSelfNonCharHostedNonCharRunnersWrite OrganizationNonCharSelfNonCharHostedNonCharRunners = "write"
)

const (
	RepositoryNonCharProjectsRead  RepositoryNonCharProjects = "read"
	RepositoryNonCharProjectsWrite RepositoryNonCharProjects = "write"
	RepositoryNonCharProjectsAdmin RepositoryNonCharProjects = "admin"
)

const (
	SecretsRead  Secrets = "read"
	SecretsWrite Secrets = "write"
)

const (
	OrganizationNonCharProjectsRead  OrganizationNonCharProjects = "read"
	OrganizationNonCharProjectsWrite OrganizationNonCharProjects = "write"
	OrganizationNonCharProjectsAdmin OrganizationNonCharProjects = "admin"
)

const (
	ActionsRead  Actions = "read"
	ActionsWrite Actions = "write"
)

const (
	ChecksRead  Checks = "read"
	ChecksWrite Checks = "write"
)

const (
	ContentsRead  Contents = "read"
	ContentsWrite Contents = "write"
)

const (
	SingleNonCharFileRead  SingleNonCharFile = "read"
	SingleNonCharFileWrite SingleNonCharFile = "write"
)

const (
	OrganizationNonCharPackagesRead  OrganizationNonCharPackages = "read"
	OrganizationNonCharPackagesWrite OrganizationNonCharPackages = "write"
)

const (
	WorkflowsWrite Workflows = "write"
)

const (
	PackagesRead  Packages = "read"
	PackagesWrite Packages = "write"
)

const (
	EnvironmentsRead  Environments = "read"
	EnvironmentsWrite Environments = "write"
)

const (
	IssuesRead  Issues = "read"
	IssuesWrite Issues = "write"
)

const (
	AdministrationRead  Administration = "read"
	AdministrationWrite Administration = "write"
)

const (
	OrganizationNonCharPlanRead OrganizationNonCharPlan = "read"
)

const (
	OrganizationNonCharAdministrationRead  OrganizationNonCharAdministration = "read"
	OrganizationNonCharAdministrationWrite OrganizationNonCharAdministration = "write"
)

const (
	MembersRead  Members = "read"
	MembersWrite Members = "write"
)

const (
	PagesRead  Pages = "read"
	PagesWrite Pages = "write"
)

const (
	OrganizationNonCharUserNonCharBlockingRead  OrganizationNonCharUserNonCharBlocking = "read"
	OrganizationNonCharUserNonCharBlockingWrite OrganizationNonCharUserNonCharBlocking = "write"
)

const (
	PullNonCharRequestsRead  PullNonCharRequests = "read"
	PullNonCharRequestsWrite PullNonCharRequests = "write"
)

const (
	StatusesRead  Statuses = "read"
	StatusesWrite Statuses = "write"
)

const (
	VulnerabilityNonCharAlertsRead  VulnerabilityNonCharAlerts = "read"
	VulnerabilityNonCharAlertsWrite VulnerabilityNonCharAlerts = "write"
)

type SimpleUser struct {
	Name              string `json:"name" yaml:"name"`
	Login             string `json:"login" yaml:"login"`
	Id                int    `json:"id" yaml:"id"`
	StarredUrl        string `json:"starred_url" yaml:"starred_url"`
	HtmlUrl           string `json:"html_url" yaml:"html_url"`
	SiteAdmin         bool   `json:"site_admin" yaml:"site_admin"`
	Url               string `json:"url" yaml:"url"`
	OrganizationsUrl  string `json:"organizations_url" yaml:"organizations_url"`
	GistsUrl          string `json:"gists_url" yaml:"gists_url"`
	FollowingUrl      string `json:"following_url" yaml:"following_url"`
	ReceivedEventsUrl string `json:"received_events_url" yaml:"received_events_url"`
	StarredAt         string `json:"starred_at" yaml:"starred_at"`
	SubscriptionsUrl  string `json:"subscriptions_url" yaml:"subscriptions_url"`
	Type              string `json:"type" yaml:"type"`
	GravatarId        string `json:"gravatar_id" yaml:"gravatar_id"`
	EventsUrl         string `json:"events_url" yaml:"events_url"`
	AvatarUrl         string `json:"avatar_url" yaml:"avatar_url"`
	NodeId            string `json:"node_id" yaml:"node_id"`
	FollowersUrl      string `json:"followers_url" yaml:"followers_url"`
	Email             string `json:"email" yaml:"email"`
	ReposUrl          string `json:"repos_url" yaml:"repos_url"`
}

type BasicError struct {
	DocumentationUrl string `json:"documentation_url" yaml:"documentation_url"`
	Status           string `json:"status" yaml:"status"`
	Url              string `json:"url" yaml:"url"`
	Message          string `json:"message" yaml:"message"`
}

type Authorization struct {
	TokenLastEight string `json:"token_last_eight" yaml:"token_last_eight"`
	Scopes         []scopes

	Url          string       `json:"url" yaml:"url"`
	CreatedAt    string       `json:"created_at" yaml:"created_at"`
	Note         string       `json:"note" yaml:"note"`
	Installation Installation `json:"installation" yaml:"installation"`
	App          App          `json:"app" yaml:"app"`
	User         User         `json:"user" yaml:"user"`
	Token        string       `json:"token" yaml:"token"`
	Fingerprint  string       `json:"fingerprint" yaml:"fingerprint"`
	HashedToken  string       `json:"hashed_token" yaml:"hashed_token"`
	UpdatedAt    string       `json:"updated_at" yaml:"updated_at"`
	Id           int          `json:"id" yaml:"id"`
	ExpiresAt    string       `json:"expires_at" yaml:"expires_at"`
	NoteUrl      string       `json:"note_url" yaml:"note_url"`
}

type Installation struct{}

type App struct{}

type User struct{}

type NullableSimpleUser struct {
	OrganizationsUrl  string `json:"organizations_url" yaml:"organizations_url"`
	Name              string `json:"name" yaml:"name"`
	HtmlUrl           string `json:"html_url" yaml:"html_url"`
	NodeId            string `json:"node_id" yaml:"node_id"`
	Url               string `json:"url" yaml:"url"`
	Email             string `json:"email" yaml:"email"`
	AvatarUrl         string `json:"avatar_url" yaml:"avatar_url"`
	ReceivedEventsUrl string `json:"received_events_url" yaml:"received_events_url"`
	Login             string `json:"login" yaml:"login"`
	GistsUrl          string `json:"gists_url" yaml:"gists_url"`
	Type              string `json:"type" yaml:"type"`
	StarredAt         string `json:"starred_at" yaml:"starred_at"`
	StarredUrl        string `json:"starred_url" yaml:"starred_url"`
	ReposUrl          string `json:"repos_url" yaml:"repos_url"`
	EventsUrl         string `json:"events_url" yaml:"events_url"`
	Id                int    `json:"id" yaml:"id"`
	FollowersUrl      string `json:"followers_url" yaml:"followers_url"`
	FollowingUrl      string `json:"following_url" yaml:"following_url"`
	GravatarId        string `json:"gravatar_id" yaml:"gravatar_id"`
	SubscriptionsUrl  string `json:"subscriptions_url" yaml:"subscriptions_url"`
	SiteAdmin         bool   `json:"site_admin" yaml:"site_admin"`
}

type NullableScopedInstallation struct {
	SingleFilePaths []single_file_paths

	RepositoriesUrl        string      `json:"repositories_url" yaml:"repositories_url"`
	Account                Account     `json:"account" yaml:"account"`
	Permissions            Permissions `json:"permissions" yaml:"permissions"`
	RepositorySelection    string      `json:"repository_selection" yaml:"repository_selection"`
	SingleFileName         string      `json:"single_file_name" yaml:"single_file_name"`
	HasMultipleSingleFiles bool        `json:"has_multiple_single_files" yaml:"has_multiple_single_files"`
}

type Account struct{}

type Permissions struct{}

type AppPermissions struct {
	Deployments                   string `json:"deployments" yaml:"deployments"`
	SecurityEvents                string `json:"security_events" yaml:"security_events"`
	Metadata                      string `json:"metadata" yaml:"metadata"`
	TeamDiscussions               string `json:"team_discussions" yaml:"team_discussions"`
	RepositoryHooks               string `json:"repository_hooks" yaml:"repository_hooks"`
	SecretScanningAlerts          string `json:"secret_scanning_alerts" yaml:"secret_scanning_alerts"`
	OrganizationHooks             string `json:"organization_hooks" yaml:"organization_hooks"`
	OrganizationSecrets           string `json:"organization_secrets" yaml:"organization_secrets"`
	OrganizationSelfHostedRunners string `json:"organization_self_hosted_runners" yaml:"organization_self_hosted_runners"`
	RepositoryProjects            string `json:"repository_projects" yaml:"repository_projects"`
	Secrets                       string `json:"secrets" yaml:"secrets"`
	OrganizationProjects          string `json:"organization_projects" yaml:"organization_projects"`
	Actions                       string `json:"actions" yaml:"actions"`
	Checks                        string `json:"checks" yaml:"checks"`
	Contents                      string `json:"contents" yaml:"contents"`
	SingleFile                    string `json:"single_file" yaml:"single_file"`
	OrganizationPackages          string `json:"organization_packages" yaml:"organization_packages"`
	Workflows                     string `json:"workflows" yaml:"workflows"`
	Packages                      string `json:"packages" yaml:"packages"`
	Environments                  string `json:"environments" yaml:"environments"`
	Issues                        string `json:"issues" yaml:"issues"`
	Administration                string `json:"administration" yaml:"administration"`
	OrganizationPlan              string `json:"organization_plan" yaml:"organization_plan"`
	OrganizationAdministration    string `json:"organization_administration" yaml:"organization_administration"`
	Members                       string `json:"members" yaml:"members"`
	Pages                         string `json:"pages" yaml:"pages"`
	OrganizationUserBlocking      string `json:"organization_user_blocking" yaml:"organization_user_blocking"`
	PullRequests                  string `json:"pull_requests" yaml:"pull_requests"`
	Statuses                      string `json:"statuses" yaml:"statuses"`
	VulnerabilityAlerts           string `json:"vulnerability_alerts" yaml:"vulnerability_alerts"`
}

type OauthAuthorizationsCreateAuthorizationInlineReqJson struct {
	ClientId     string `json:"client_id" yaml:"client_id"`
	ClientSecret string `json:"client_secret" yaml:"client_secret"`
	Fingerprint  string `json:"fingerprint" yaml:"fingerprint"`
	Scopes       []scopes

	Note    string `json:"note" yaml:"note"`
	NoteUrl string `json:"note_url" yaml:"note_url"`
}

type OauthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintInlineReqJson struct {
	ClientSecret string `json:"client_secret" yaml:"client_secret"`
	Scopes       []scopes

	Note    string `json:"note" yaml:"note"`
	NoteUrl string `json:"note_url" yaml:"note_url"`
}

type OauthAuthorizationsUpdateAuthorizationInlineReqJson struct {
	RemoveScopes []remove_scopes

	Note        string `json:"note" yaml:"note"`
	NoteUrl     string `json:"note_url" yaml:"note_url"`
	Fingerprint string `json:"fingerprint" yaml:"fingerprint"`
	Scopes      []scopes

	AddScopes []add_scopes
}

type OauthAuthorizationsGetOrCreateAuthorizationForAppInlineReqJson struct {
	Fingerprint  string `json:"fingerprint" yaml:"fingerprint"`
	ClientSecret string `json:"client_secret" yaml:"client_secret"`
	Scopes       []scopes

	Note    string `json:"note" yaml:"note"`
	NoteUrl string `json:"note_url" yaml:"note_url"`
}
