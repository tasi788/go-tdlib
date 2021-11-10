package tdlib

// PageBlockRelatedArticle Contains information about a related article
type PageBlockRelatedArticle struct {
	tdCommon
	Url         string `json:"url"`          // Related article URL
	Title       string `json:"title"`        // Article title; may be empty
	Description string `json:"description"`  // Article description; may be empty
	Photo       *Photo `json:"photo"`        // Article photo; may be null
	Author      string `json:"author"`       // Article author; may be empty
	PublishDate int32  `json:"publish_date"` // Point in time (Unix timestamp) when the article was published; 0 if unknown
}

// MessageType return the string telegram-type of PageBlockRelatedArticle
func (pageBlockRelatedArticle *PageBlockRelatedArticle) MessageType() string {
	return "pageBlockRelatedArticle"
}

// NewPageBlockRelatedArticle creates a new PageBlockRelatedArticle
//
// @param url Related article URL
// @param title Article title; may be empty
// @param description Article description; may be empty
// @param photo Article photo; may be null
// @param author Article author; may be empty
// @param publishDate Point in time (Unix timestamp) when the article was published; 0 if unknown
func NewPageBlockRelatedArticle(url string, title string, description string, photo *Photo, author string, publishDate int32) *PageBlockRelatedArticle {
	pageBlockRelatedArticleTemp := PageBlockRelatedArticle{
		tdCommon:    tdCommon{Type: "pageBlockRelatedArticle"},
		Url:         url,
		Title:       title,
		Description: description,
		Photo:       photo,
		Author:      author,
		PublishDate: publishDate,
	}

	return &pageBlockRelatedArticleTemp
}
