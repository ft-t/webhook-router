package structs

type PathRule struct {
	Path      string     `json:"path"`
	Endpoints []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	Url      string `json:"url"`
	Priority int    `json:"priority"`
}
