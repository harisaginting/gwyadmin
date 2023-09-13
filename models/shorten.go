package models

type Shorten struct {
	ID            int    `json:"id"`
	Shortcode     string `json:"shortcode"`
	Url           string `json:"url"`
	RedirectCount int64  `json:"redirectCount"`
	LastSeenDate  string `json:"lastSeenDate"`
	StartDate     string `json:"startDate"`
}
