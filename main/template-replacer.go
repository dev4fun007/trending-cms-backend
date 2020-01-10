package main

import "regexp"

const (
	feedGroupPattern = `[{feed-group} [a-z]+]`
	feedCardPattern = `[{feed-card} range [0-9]+ to [0-9]+]`
	feedLinePattern = `[{feed-line} range [0-9]+ to [0-9]+]`
	feedBannerPattern = `[{feed-banner} range [0-9]+ to [0-9]+]`
	feedYoutubePattern = `[{feed-youtube} range [0-9]+ to [0-9]+]`
	feedMemePattern = `[{feed-meme} range [0-9]+ to [0-9]+]`
	feedAppsPattern = `[{feed-apps} range [0-9]+ to [0-9]+]`
)


var (
	feedGroupRegex *regexp.Regexp
	feedCardRegex *regexp.Regexp
	feedLineRegex *regexp.Regexp
	feedBannerRegex *regexp.Regexp
	feedYoutubeRegex *regexp.Regexp
	feedMemeRegex *regexp.Regexp
	feedAppsRegex *regexp.Regexp
)


func init() {
	feedGroupRegex = regexp.MustCompile(feedGroupPattern)
	feedCardRegex = regexp.MustCompile(feedCardPattern)
	feedLineRegex = regexp.MustCompile(feedLinePattern)
	feedBannerRegex = regexp.MustCompile(feedBannerPattern)
	feedYoutubeRegex = regexp.MustCompile(feedYoutubePattern)
	feedMemeRegex = regexp.MustCompile(feedMemePattern)
	feedAppsRegex = regexp.MustCompile(feedAppsPattern)
}


func ReplaceTemplate(data []byte) string {

	s := string(data)

	//1st - Insert QOD
	s = ReplaceQod(s)

	//2nd - Start replacing Feed Group Collections




	return s

}


