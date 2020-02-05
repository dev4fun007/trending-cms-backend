package replacerutil

import (
	"strings"
	"trending-cms-backend/model"
)

const (
	androidPlatform = `<span class="icons8-google-play"></span>`
	applePlatform   = `<span class="icons8-apple-app-store"></span>`
)

func InsertFeedCardData(startIndex, endIndex int, feedItemList []model.ArticleFeedItem, feedCardTemplate string) string {

	var sb strings.Builder
	sb.Grow(endIndex - startIndex)
	for _, _ = range feedItemList {
		article := feedItemList[startIndex]

		//Copy of feed card template
		cardTemplate := feedCardTemplate

		//Replace path
		cardTemplate = strings.Replace(cardTemplate, "{{article.path}}", article.Path, 1)

		//Replace title
		cardTemplate = strings.Replace(cardTemplate, "{{article.title}}", article.Title, 1)

		//Replace timestamp
		cardTemplate = strings.Replace(cardTemplate, "{{article.datetime}}", article.Timestamp, 2)

		//Replace image and imageAlt
		cardTemplate = strings.Replace(cardTemplate, "{{article.image}}", article.Image, 1)
		cardTemplate = strings.Replace(cardTemplate, "{{article.imageAlt}}", article.ImageAlt, 1)

		sb.WriteString(cardTemplate)
		sb.WriteString("\n")

		//Insert Data Till Here
		startIndex++
		if startIndex == endIndex {
			break
		}
	}

	return sb.String()
}

func InsertFeedLineData(startIndex, endIndex int, feedItemList []model.ArticleFeedItem, feedLineTemplate string) string {

	var sb strings.Builder
	sb.Grow(endIndex - startIndex)
	for _, _ = range feedItemList {
		article := feedItemList[startIndex]

		//Copy of feed card template
		lineTemplate := feedLineTemplate

		//Replace path
		lineTemplate = strings.Replace(lineTemplate, "{{article.path}}", article.Path, 1)

		//Replace title
		lineTemplate = strings.Replace(lineTemplate, "{{article.title}}", article.Title, 1)

		//Replace timestamp
		lineTemplate = strings.Replace(lineTemplate, "{{article.datetime}}", article.Timestamp, 2)

		//Replace image and imageAlt
		lineTemplate = strings.Replace(lineTemplate, "{{article.image}}", article.Image, 1)
		lineTemplate = strings.Replace(lineTemplate, "{{article.imageAlt}}", article.ImageAlt, 1)

		sb.WriteString(lineTemplate)
		sb.WriteString("\n")

		//Insert Data Till Here
		startIndex++
		if startIndex == endIndex {
			break
		}
	}

	return sb.String()

}

func InsertFeedBannerData(startIndex, endIndex int, bannerItemList []model.ArticleBannerItem, feedBannerTemplate string) string {

	var sb strings.Builder
	sb.Grow(endIndex - startIndex)
	for _, _ = range bannerItemList {
		banner := bannerItemList[startIndex]

		//Copy of feed card template
		bannerTemplate := feedBannerTemplate

		//Replace path
		bannerTemplate = strings.Replace(bannerTemplate, "{{banner.feedLink}}", banner.Path, 1)

		//Replace caption
		bannerTemplate = strings.Replace(bannerTemplate, "{{banner.caption}}", banner.Caption, 1)

		//Replace image and imageAlt
		bannerTemplate = strings.Replace(bannerTemplate, "{{banner.image}}", banner.Image, 1)
		bannerTemplate = strings.Replace(bannerTemplate, "{{banner.imageAlt}}", banner.ImageAlt, 1)

		sb.WriteString(bannerTemplate)
		sb.WriteString("\n")

		//Insert Data Till Here
		startIndex++
		if startIndex == endIndex {
			break
		}
	}

	return sb.String()

}

func InsertFeedYoutubeData(startIndex, endIndex int, youtubeItemList []model.YoutubeFeedItem, feedYoutubeTemplate string) string {

	var sb strings.Builder
	sb.Grow(endIndex - startIndex)
	for _, _ = range youtubeItemList {
		video := youtubeItemList[startIndex]

		//Copy of feed card template
		youtubeTemplate := feedYoutubeTemplate

		//Replace path
		youtubeTemplate = strings.Replace(youtubeTemplate, "{{youtube.videoUrl}}", video.VideoUrl, 1)

		//Replace caption
		youtubeTemplate = strings.Replace(youtubeTemplate, "{{youtube.caption}}", video.Caption, 1)

		sb.WriteString(youtubeTemplate)
		sb.WriteString("\n")

		//Insert Data Till Here
		startIndex++
		if startIndex == endIndex {
			break
		}
	}

	return sb.String()

}

func InsertFeedAppData(startIndex, endIndex int, appItemList []model.AppFeedItem, feedAppTemplate string) string {

	var sb strings.Builder
	sb.Grow(endIndex - startIndex)
	for _, _ = range appItemList {
		app := appItemList[startIndex]

		//Copy of feed card template
		appTemplate := feedAppTemplate

		appTemplate = strings.Replace(appTemplate, "{{app.storeUrl}}", app.StoreUrl, 1)
		appTemplate = strings.Replace(appTemplate, "{{app.iconUrl}}", app.IconUrl, 1)
		appTemplate = strings.Replace(appTemplate, "{{app.name}}", app.Name, 1)
		appTemplate = strings.Replace(appTemplate, "{{app.price}}", app.Price, 1)
		appTemplate = strings.Replace(appTemplate, "{{app.rating}}", app.Rating, 1)
		if strings.ToLower(app.Platform) == "android" {
			appTemplate = strings.Replace(appTemplate, "{{app.platform}}", androidPlatform, 1)
		} else if strings.ToLower(app.Platform) == "apple" {
			appTemplate = strings.Replace(appTemplate, "{{app.platform}}", applePlatform, 1)
		}

		sb.WriteString(appTemplate)
		sb.WriteString("\n")

		//Insert Data Till Here
		startIndex++
		if startIndex == endIndex {
			break
		}
	}

	return sb.String()

}
