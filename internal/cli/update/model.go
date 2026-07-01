package update

type Release struct {
	TagName string  `json:"tag_name"`
	HtmlUrl string  `json:"html_url"`
	Assets  []Asset `json:"assets"`
}

type Asset struct {
	Name string `json:"name"`
	URL  string `json:"browser_download_url"`
}
