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

	Id       string   `json:"id"`
	Category string   `json:"category"`
	Caption  string   `json:"caption"`
	Image    string   `json:"image"`
	Tags     []string `json:"tags"`
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
	Thumbnail string   `json:"thumbnail"`
}