package dto

// ShortenRequest holds the original URL
type ShortenRequest struct {
	URL string `json:"url" binding:"required,url"`
}

// ShortenResponse returns the generated code
type ShortenResponse struct {
	Code string `json:"code"`
}
