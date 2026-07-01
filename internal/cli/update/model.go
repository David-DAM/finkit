package update

type Release struct {
	TagName string `json:"tag_name"`
	URL     string `json:"url"`
}
