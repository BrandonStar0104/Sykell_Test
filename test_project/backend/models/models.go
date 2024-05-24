package models

type ProcessRequest struct {
	URL string `json:"url"`
}

type ProcessResult struct {
	URL               string         `json:"url"`
	HTMLVersion       string         `json:"html_version"`
	PageTitle         string         `json:"page_title"`
	HeadingsCount     map[string]int `json:"headings_count"`
	InternalLinks     int            `json:"internal_links"`
	ExternalLinks     int            `json:"external_links"`
	InaccessibleLinks int            `json:"inaccessible_links"`
	HasLoginForm      bool           `json:"has_login_form"`
}
