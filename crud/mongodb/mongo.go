package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"strings"
	"time"
	"trending-cms-backend/crud"
	"trending-cms-backend/model"
	"trending-cms-backend/util"
)

type Mongo struct {
	client *mongo.Client
}

func NewMongoClient() crud.CRUD {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Unable to connect to mongodb instance: %s", err)
	}
	return &Mongo{client: client}
}

func (m *Mongo) Connect() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := m.client.Ping(ctx, readpref.Primary())
	return err
}

func (m *Mongo) GetCategoryInfoByCategoryName(category string) model.Category {

	var cat model.Category
	collection := m.client.Database(util.DatabaseName).Collection(util.CategoryInfoCollection)
	res := collection.FindOne(context.Background(), bson.D{{"Name", strings.ToLower(category)}})
	err := res.Decode(&cat)
	if err != nil {
		log.Printf("Error Decoding to CategoryInfo: %s", err)
	}
	return cat

}

/*
	Category specifies the article category
	Limit == 0 indicates fetch all
*/
func (m *Mongo) GetArticleFeedItemListByCategory(category string, limit int64) []model.ArticleFeedItem {

	articles := make([]model.ArticleFeedItem, 0, 10)
	collection := m.client.Database(util.DatabaseName).Collection(category)

	findOptions := options.FindOptions{
		Limit: &limit,
		Sort:  bson.D{{"Timestamp", -1}},
	}

	cursor, err := collection.Find(context.Background(), bson.D{}, &findOptions)
	if err != nil {
		return articles
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var article model.ArticleFeedItem
		err = cursor.Decode(&article)
		if err != nil {
			log.Printf("Cannot Decode to ArticleFeedItem: %s", err)
			continue
		}
		articles = append(articles, article)
	}

	return articles

}

func (m *Mongo) GetArticleFeedBannerList() []model.ArticleBannerItem {

	banners := make([]model.ArticleBannerItem, 0, 3)
	collection := m.client.Database(util.DatabaseName).Collection(util.BannerItemCollection)

	limit := int64(3)
	findOptions := options.FindOptions{
		Limit: &limit,
		Sort:  bson.D{{"Timestamp", -1}},
	}

	cursor, err := collection.Find(context.Background(), bson.D{}, &findOptions)
	if err != nil {
		return banners
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var banner model.ArticleBannerItem
		err = cursor.Decode(&banner)
		if err != nil {
			log.Printf("Cannot Decode to BannerItem: %s", err)
			continue
		}
		banners = append(banners, banner)
	}

	return banners

}

func (m *Mongo) GetFeedYoutubeList() []model.YoutubeFeedItem {

	videos := make([]model.YoutubeFeedItem, 0, 3)
	collection := m.client.Database(util.DatabaseName).Collection(util.YoutubeVideosCollection)

	limit := int64(3)
	findOptions := options.FindOptions{
		Limit: &limit,
		Sort:  bson.D{{"Timestamp", -1}},
	}

	cursor, err := collection.Find(context.Background(), bson.D{}, &findOptions)
	if err != nil {
		return videos
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var video model.YoutubeFeedItem
		err = cursor.Decode(&video)
		if err != nil {
			log.Printf("Cannot Decode to YoutubeVideoItem: %s", err)
			continue
		}
		videos = append(videos, video)
	}

	return videos

}

func (m *Mongo) GetFeedAppList() []model.AppFeedItem {

	apps := make([]model.AppFeedItem, 0, 8)
	collection := m.client.Database(util.DatabaseName).Collection(util.FreeAppsCollection)

	//Only 8 free apps for today
	limit := int64(8)
	findOptions := options.FindOptions{
		Limit: &limit,
		Sort:  bson.D{{"Timestamp", -1}},
	}

	cursor, err := collection.Find(context.Background(), bson.D{}, &findOptions)
	if err != nil {
		return apps
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var app model.AppFeedItem
		err = cursor.Decode(&app)
		if err != nil {
			log.Printf("Cannot Decode to FreeAppItem: %s", err)
			continue
		}
		apps = append(apps, app)
	}

	return apps

}
