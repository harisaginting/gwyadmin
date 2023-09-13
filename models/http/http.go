package http

import model "github.com/harisaginting/gwyn/models"

type RequestCreate struct {
	URL       string `json:"url"`
	Shortcode string `json:"shortcode"`
}

type ResponseCreate struct {
	Shortcode string `json:"shortcode"`
}

type ResponseStatus struct {
	StartDate     string `json:"startDate"`
	LastSeenDate  string `json:"lastSeenDate,omitempty"`
	RedirectCount int64  `json:"redirectCount"`
}

type ResponseList struct {
	Items []model.Shorten `json:"items"`
	Total int             `json:"total"`
}
