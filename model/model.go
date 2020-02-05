package model

type ArticleDetailed struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	SubTitle    string   `json:"subTitle"`
	Timestamp   string   `json:"timestamp"`
	CoverImgUrl string   `json:"coverImgUrl"`
	Tldr        []string `json:"tldr"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
	Content     string   `json:"content"`
}

type ArticleBannerItem struct {
	Id        string   `json:"id"`
	Category  string   `json:"category"`
	Caption   string   `json:"caption"`
	Image     string   `json:"image"`
	ImageAlt  string   `json:"imageAlt"`
	Path      string   `json:"path"`
	Tags      []string `json:"tags"`
	Timestamp string   `json:"timestamp"`
}

type ArticleFeedItem struct {
	Id        string   `json:"id"`
	Category  string   `json:"category"`
	Title     string   `json:"title"`
	SubTitle  string   `json:"subTitle"`
	Author    string   `json:"author"`
	Timestamp string   `json:"timestamp"`
	Tags      []string `json:"tags"`
	Image     string   `json:"image"`
	ImageAlt  string   `json:"imageAlt"`
	Thumbnail string   `json:"thumbnail"`
	Path      string   `json:"path"`
}

type Category struct {
	Name         string `json:"name"`
	CategoryLink string `json:"categoryLink"`
	MoreLink     string `json:"moreLink"`
}

type YoutubeFeedItem struct {
	Caption   string `json:"caption"`
	VideoUrl  string `json:"videoUrl"`
	Timestamp string `json:"timestamp"`
}

type AppFeedItem struct {
	Platform  string `json:"platform"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	Rating    string `json:"rating"`
	StoreUrl  string `json:"storeUrl"`
	IconUrl   string `json:"iconUrl"`
	Timestamp string `json:"timestamp"`
}
