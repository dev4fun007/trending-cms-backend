package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
	"trending-cms-backend/crud"
	"trending-cms-backend/crud/mongodb"
	"trending-cms-backend/main/replacerutil"
	"trending-cms-backend/model"
)

const (
	feedGroupPattern   = `\[\{feed-group\} [a-z-]+\]`
	feedCardPattern    = `\[\{feed-card\} range [0-9]+ to [0-9]+\]`
	feedLinePattern    = `\[\{feed-line\} range [0-9]+ to [0-9]+\]`
	feedBannerPattern  = `\[\{feed-banner\} range [0-9]+ to [0-9]+\]`
	feedYoutubePattern = `\[\{feed-youtube\} range [0-9]+ to [0-9]+\]`
	feedMemePattern    = `\[\{feed-meme\} range [0-9]+ to [0-9]+\]`
	feedAppsPattern    = `\[\{feed-apps\} range [0-9]+ to [0-9]+\]`

	digitPattern = `[\d]+`
)

var (
	feedGroupRegex   *regexp.Regexp
	feedCardRegex    *regexp.Regexp
	feedLineRegex    *regexp.Regexp
	feedBannerRegex  *regexp.Regexp
	feedYoutubeRegex *regexp.Regexp
	feedMemeRegex    *regexp.Regexp
	feedAppsRegex    *regexp.Regexp

	digitRegex *regexp.Regexp

	feedGroupTemplate   string
	feedCardTemplate    string
	feedLineTemplate    string
	feedBannerTemplate  string
	feedYoutubeTemplate string
	feedAppTemplate     string

	dbClient crud.CRUD
)

func init() {
	feedGroupRegex = regexp.MustCompile(feedGroupPattern)
	feedCardRegex = regexp.MustCompile(feedCardPattern)
	feedLineRegex = regexp.MustCompile(feedLinePattern)
	feedBannerRegex = regexp.MustCompile(feedBannerPattern)
	feedYoutubeRegex = regexp.MustCompile(feedYoutubePattern)
	feedMemeRegex = regexp.MustCompile(feedMemePattern)
	feedAppsRegex = regexp.MustCompile(feedAppsPattern)

	digitRegex = regexp.MustCompile(digitPattern)

	data, _ := ioutil.ReadFile("asset/templates/feed-card.html")
	feedCardTemplate = string(data)

	data, _ = ioutil.ReadFile("asset/templates/feed-line.html")
	feedLineTemplate = string(data)

	data, _ = ioutil.ReadFile("asset/templates/feed-group.html")
	feedGroupTemplate = string(data)

	data, _ = ioutil.ReadFile("asset/templates/feed-banner.html")
	feedBannerTemplate = string(data)

	data, _ = ioutil.ReadFile("asset/templates/feed-youtube.html")
	feedYoutubeTemplate = string(data)

	data, _ = ioutil.ReadFile("asset/templates/feed-apps.html")
	feedAppTemplate = string(data)

	dbClient = mongodb.NewMongoClient()
	err := dbClient.Connect()
	if err != nil {
		log.Fatalf("Error Connecting to Database %s", err)
	}
}

func ReplaceTemplate(data []byte) string {

	indexHtml := string(data)

	/* 0. Replace Banner Article Template */
	banners := dbClient.GetArticleFeedBannerList()
	feedBannerString, formattedBannerString := getFormattedFeedBanner(indexHtml, banners)
	indexHtml = strings.Replace(indexHtml, feedBannerString, formattedBannerString, 1)

	/* 1. Replace Quote of the day */
	//s = ReplaceQod(s)

	/* 2. Replace Groups */
	//Each group will be replaced by feedGroupTemplate
	//But first replace other templates in the feedGroupTemplate
	feedGroups := feedGroupRegex.FindAllString(indexHtml, -1)
	for index, _ := range feedGroups {

		group := feedGroups[index]
		fmt.Println(group)
		//Make a copy of feedGroupTemplate for replacing with actual values
		groupTemplate := feedGroupTemplate

		split := strings.Split(group, " ")
		if len(split) == 2 {
			category := strings.TrimRight(split[1], "]")

			feedItemList := dbClient.GetArticleFeedItemListByCategory(category, 6)
			//No data to populate - continue the loop
			if len(feedItemList) == 0 {
				continue
			}

			/* 0. Replace Category Info */
			groupTemplate = getFormattedGroupTemplate(groupTemplate, crud.GetCategoryInfoByCategoryName(category))

			/* 1. Replace Feed Card Items */
			feedCardString, formattedCardString := getFormattedFeedCard(feedItemList)
			groupTemplate = strings.Replace(groupTemplate, feedCardString, formattedCardString, 1)

			/* 2. Replace Feed Line Items */
			feedLineString, formattedLineString := getFormattedFeedLine(feedItemList)
			groupTemplate = strings.Replace(groupTemplate, feedLineString, formattedLineString, 1)

			//Insert the group in indexHtml file
			indexHtml = strings.Replace(indexHtml, group, groupTemplate, 1)

		} else {
			log.Fatal("Invalid FeedGroup Template Syntax, Check index.html")
		}
	}

	/* 3. Replace Youtube Feed */
	youtubeVideos := dbClient.GetFeedYoutubeList()
	feedYoutubeString, formattedYoutubeString := getFormattedFeedYoutube(indexHtml, youtubeVideos)
	indexHtml = strings.Replace(indexHtml, feedYoutubeString, formattedYoutubeString, 1)

	/* 4. Replace Free Apps Feed*/
	freeApps := dbClient.GetFeedAppList()
	freeAppsString, formattedAppsString := getFormattedFeedApps(indexHtml, freeApps)
	indexHtml = strings.Replace(indexHtml, freeAppsString, formattedAppsString, 1)

	return indexHtml

}

func getFormattedGroupTemplate(groupTemplate string, category model.Category) string {

	groupTemplate = strings.Replace(groupTemplate, "{{feedgroup.categoryLink}}", category.CategoryLink, 1)
	groupTemplate = strings.Replace(groupTemplate, "{{feedgroup.name}}", category.Name, 1)
	groupTemplate = strings.Replace(groupTemplate, "{{feedgroup.moreLink}}", category.MoreLink, 1)

	return groupTemplate
}

func getFormattedFeedCard(feedItemList []model.ArticleFeedItem) (string, string) {
	var formattedCardString string

	//Get FeedCard Template
	feedCardString := feedCardRegex.FindString(feedGroupTemplate)
	//Get FeedCard Looping Range
	cardRange := digitRegex.FindAllString(feedCardString, 2)
	if len(cardRange) == 2 {
		start, _ := strconv.Atoi(cardRange[0])
		end, _ := strconv.Atoi(cardRange[1])
		//Replace the card template string with actual value
		formattedCardString = replacerutil.InsertFeedCardData(start, end, feedItemList, feedCardTemplate) //Actual value
	} else {
		log.Fatal("Invalid FeedCard Template Syntax")
	}

	return feedCardString, formattedCardString
}

func getFormattedFeedLine(feedItemList []model.ArticleFeedItem) (string, string) {
	var formattedLineString string

	//Get FeedLine Template
	feedLineString := feedLineRegex.FindString(feedGroupTemplate)
	//Get FeedLine Looping Range
	lineRange := digitRegex.FindAllString(feedLineString, 2)
	if len(lineRange) == 2 {
		start, _ := strconv.Atoi(lineRange[0])
		end, _ := strconv.Atoi(lineRange[1])
		formattedLineString = replacerutil.InsertFeedLineData(start, end, feedItemList, feedLineTemplate)
	} else {
		log.Fatal("Invalid FeedLine Template Syntax")
	}

	return feedLineString, formattedLineString
}

func getFormattedFeedBanner(indexHtml string, bannerItemList []model.ArticleBannerItem) (string, string) {
	var formattedBannerString string

	feedBannerString := feedBannerRegex.FindString(indexHtml) //we will get :- [{feed-banner} range 0 to 3]
	bannerTemplate := feedBannerTemplate

	bannerRange := digitRegex.FindAllString(feedBannerString, 2)
	if len(bannerRange) == 2 {
		start, _ := strconv.Atoi(bannerRange[0])
		end, _ := strconv.Atoi(bannerRange[1])
		formattedBannerString = replacerutil.InsertFeedBannerData(start, end, bannerItemList, bannerTemplate)
	} else {
		log.Fatal("Invalid FeedBanner Template Syntax")
	}

	return feedBannerString, formattedBannerString
}

func getFormattedFeedYoutube(indexHtml string, youtubeItemList []model.YoutubeFeedItem) (string, string) {
	var formattedYoutubeString string

	feedYoutubeString := feedYoutubeRegex.FindString(indexHtml) //we will get :- [{feed-youtube} range 0 to 3]
	youtubeTemplate := feedYoutubeTemplate

	youtubeRange := digitRegex.FindAllString(feedYoutubeString, 2)
	if len(youtubeRange) == 2 {
		start, _ := strconv.Atoi(youtubeRange[0])
		end, _ := strconv.Atoi(youtubeRange[1])
		formattedYoutubeString = replacerutil.InsertFeedYoutubeData(start, end, youtubeItemList, youtubeTemplate)
	} else {
		log.Fatal("Invalid FeedYoutube Template Syntax")
	}

	return feedYoutubeString, formattedYoutubeString
}

func getFormattedFeedApps(indexHtml string, appsItemList []model.AppFeedItem) (string, string) {
	var formattedAppsString string

	feedAppsString := feedAppsRegex.FindString(indexHtml) //we will get :- [{feed-youtube} range 0 to 3]
	appTemplate := feedAppTemplate

	appRange := digitRegex.FindAllString(feedAppsString, 2)
	if len(appRange) == 2 {
		start, _ := strconv.Atoi(appRange[0])
		end, _ := strconv.Atoi(appRange[1])
		formattedAppsString = replacerutil.InsertFeedAppData(start, end, appsItemList, appTemplate)
	} else {
		log.Fatal("Invalid FeedApp Template Syntax")
	}

	return feedAppsString, formattedAppsString
}
