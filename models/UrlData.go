package models

type UrlData struct {
	Url      string `json:"url"`
	Meta     `json:"meta"`
	Elements []Element `json:"elements,omitempty"`
}

type Meta struct {
	Status        int    `json:"status"`
	ContentType   string `json:"content-type,omitempty"`
	ContentLength int    `json:"content-length,omitempty"`
}

type Element struct {
	TagName string `json:"tag-name"`
	Count   int    `json:"count"`
}
