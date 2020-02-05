package crud

import (
	"trending-cms-backend/model"
	"trending-cms-backend/util"
)

var (
	feedMap         map[string][]model.ArticleFeedItem
	categoryInfoMap map[string]model.Category
)

func init() {
	categoryInfoMap = make(map[string]model.Category)
	populateCategoryMap()
	feedMap = make(map[string][]model.ArticleFeedItem)
	populateFeedMap()
}

func populateCategoryMap() {
	categoryInfoMap[util.TECHNOLOGY] = model.Category{
		Name:         "Technology",
		CategoryLink: "#",
		MoreLink:     "#",
	}
	categoryInfoMap[util.SMARTPHONES] = model.Category{
		Name:         "Smartphones",
		CategoryLink: "#",
		MoreLink:     "#",
	}
	categoryInfoMap[util.ENTERTAINMENT] = model.Category{
		Name:         "Entertainment",
		CategoryLink: "#",
		MoreLink:     "#",
	}

	categoryInfoMap[util.APPS] = model.Category{
		Name:         "Apps",
		CategoryLink: "#",
		MoreLink:     "#",
	}
	categoryInfoMap[util.ANDROID] = model.Category{
		Name:         "Android",
		CategoryLink: "#",
		MoreLink:     "#",
	}
	categoryInfoMap[util.APPLE] = model.Category{
		Name:         "Apple",
		CategoryLink: "#",
		MoreLink:     "#",
	}

	categoryInfoMap[util.GAMING] = model.Category{
		Name:         "Gaming",
		CategoryLink: "#",
		MoreLink:     "#",
	}
	categoryInfoMap[util.SCIENCE] = model.Category{
		Name:         "Science",
		CategoryLink: "#",
		MoreLink:     "#",
	}
	categoryInfoMap[util.HowTo] = model.Category{
		Name:         "How To",
		CategoryLink: "#",
		MoreLink:     "#",
	}
}

func GetCategoryInfoByCategoryName(category string) model.Category {
	return categoryInfoMap[category]
}

func populateFeedMap() {
	feedMap[util.TECHNOLOGY] = PopulateArticleFeedTechnology()
	feedMap[util.SMARTPHONES] = []model.ArticleFeedItem{}
	feedMap[util.ENTERTAINMENT] = []model.ArticleFeedItem{}

	feedMap[util.APPS] = []model.ArticleFeedItem{}
	feedMap[util.ANDROID] = []model.ArticleFeedItem{}
	feedMap[util.APPLE] = []model.ArticleFeedItem{}

	feedMap[util.GAMING] = []model.ArticleFeedItem{}
	feedMap[util.SCIENCE] = []model.ArticleFeedItem{}
	feedMap[util.HowTo] = []model.ArticleFeedItem{}
}

func GetArticleFeedItemListByCategory(category string) []model.ArticleFeedItem {
	return feedMap[category]
}

func PopulateArticleFeedTechnology() []model.ArticleFeedItem {
	list := make([]model.ArticleFeedItem, 0, 10)
	example := model.ArticleFeedItem{
		Id:        "ABC",
		Category:  "Technology",
		Title:     "Title",
		SubTitle:  "SubTitle",
		Author:    "Author",
		Timestamp: "12 Jan, 2020",
		Tags:      nil,
		Image:     "IMAGE",
		ImageAlt:  "ALT",
		Thumbnail: "Thumbnail",
		Path:      "#",
	}
	list = append(list, example)
	list = append(list, example)
	list = append(list, example)
	list = append(list, example)
	list = append(list, example)
	list = append(list, example)
	list = append(list, example)
	list = append(list, example)
	return list
}
