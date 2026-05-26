package action

const githubApiLatestReleaseURL = "https://api.github.com/repos/mithrandie/csvq/releases/latest"
const githubApiLatestPreReleaseURL = "https://api.github.com/repos/mithrandie/csvq/releases?per_page=1"
const preReleaseIdentifier = "pr"

type GithubRelease struct {
	HTMLURL     string               `json:"html_url"`
	TagName     string               `json:"tag_name"`
	PublishedAt string               `json:"published_at"`
	Assets      []GithubReleaseAsset `json:"assets"`
}

type GithubReleaseAsset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

var CurrentVersion = &Version{}

type Version struct {
	Major      int
	Minor      int
	Patch      int
	PreRelease int
}

func (v *Version) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (v *Version) IsLaterThan(v2 *Version) bool { _ = "STUB: not implemented"; return false }

func (v *Version) String() string { _ = "STUB: not implemented"; return "" }

func ParseVersion(s string) (*Version, error) { _ = "STUB: not implemented"; return nil, nil }

type GithubClient interface {
	GetLatestRelease() (*GithubRelease, error)
	GetLatestReleaseIncludingPreRelease() (*GithubRelease, error)
}

type Client struct{}

func NewClient() GithubClient { _ = "STUB: not implemented"; return *new(GithubClient) }

func (c Client) GetLatestRelease() (*GithubRelease, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) GetLatestReleaseIncludingPreRelease() (*GithubRelease, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PickVersionNumber(s string) string { _ = "STUB: not implemented"; return "" }

func CheckUpdate(includePreRelaese bool) error { _ = "STUB: not implemented"; return nil }

func CheckForUpdates(includePreRelease bool, client GithubClient, goos string, goarch string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
