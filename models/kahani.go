package models

type Kahani struct {
	Name       string `json:"name,omitempty"`
	Title       string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	Created_on int64 `json:"created_on,omitempty"`
	Updated_on int64 `json:"created_on,omitempty"`
	Audio      string `json:"audio,omitempty"`
	Banner     string `json:"banner,omitempty"`
	Cate       string `json:"cate,omitempty"`
	Slug       string `json:"slug,omitempty"`
	Category   string `json:"category,omitempty"`
	AuthorData string `json:"authorData,omitempty"`
	Author     string `json:"author,omitempty"`
	Thumbnail  string `json:"thumbnail,omitempty"`
	Excerpt    string `json:"excerpt,omitempty"`
	Voice      string `json:"voice,omitempty"`
	Producer   string `json:"producer,omitempty"`
	Meme       string `json:"meme,omitempty"`
	Date       int64 `json:"date,omitempty"`
	Played     string `json:"played,omitempty"`
	Liked      string `json:"liked,omitempty"`
}
