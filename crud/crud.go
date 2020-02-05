package crud

import "trending-cms-backend/model"

type CRUD interface {
	Connect() error
	GetCategoryInfoByCategoryName(category string) model.Category
	GetArticleFeedItemListByCategory(category string, limit int64) []model.ArticleFeedItem
	GetArticleFeedBannerList() []model.ArticleBannerItem
	GetFeedYoutubeList() []model.YoutubeFeedItem
	GetFeedAppList() []model.AppFeedItem
}
