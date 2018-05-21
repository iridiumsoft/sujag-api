package models

type Post struct {
	Title       string  `json:"title,omitempty"`
	Excerpt     string  `json:"excerpt,omitempty"`
	Content     string  `json:"content,omitempty"`
	Thumbnail   string  `json:"thumbnail,omitempty"`
	Meme        string  `json:"meme,omitempty"`
	District    string  `json:"district,omitempty"`
	Slug        string  `json:"slug,omitempty"`
	Author      string  `json:"author,omitempty"`
	Audio       string  `json:"audio,omitempty"`
	HasAudio    string  `json:"hasAudio,omitempty"`
	PublishedBy string  `json:"published_by,omitempty"`
	User        string  `json:"user,omitempty"`
	Category    string  `json:"category,omitempty"`
	Baylag      string  `json:"baylag,omitempty"`
	Status      int     `json:"status,omitempty"`
	Type        string  `json:"type,omitempty"`
	PublishedOn float64 `json:"published_on,omitempty"`
	CreatedOn   int64   `json:"created_on,omitempty"`
	UpdatedOn   int64   `json:"updated_on,omitempty"`
	Date        string  `json:"date,omitempty"`
	Liked       float64 `json:"liked,omitempty"`
}
