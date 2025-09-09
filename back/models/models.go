package models()

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CrosspromoFile struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	FileName       string             `bson:"fileName"`
	FileSize       int64              `bson:"fileSize"`
	URL            string             `bson:"url"`
	Localization   string             `bson:"localization"`
	Platform       string             `bson:"platform"`
	AdType         string             `bson:"adType"`
	StoreCode      string             `bson:"storeCode"`
	AppID          string             `bson:"appId"`
	MimeType       string             `bson:"mimeType"`
	Width          int                `bson:"width"`
	Height         int                `bson:"height"`
}

type AssetGroup struct {
	Localization string                  `bson:"localization" json:"localization"`
	Platforms    map[string]PlatformData `bson:"platforms" json:"platforms"`
}

type PlatformData struct {
	Banner       []AssetInfo `bson:"banner" json:"banner"`
	Interstitial []AssetInfo `bson:"interstitial" json:"interstitial"`
	Rewarded     []AssetInfo `bson:"rewarded" json:"rewarded"`
}

type AssetInfo struct {
	FileName string  `bson:"fileName" json:"fileName"`
	FileSize float64 `bson:"fileSize" json:"fileSize"` 
	URL      string  `bson:"url" json:"url"`
	MimeType string  `bson:"mimeType" json:"mimeType"`
}

type MongoDBRepository struct {
	client *mongo.Client
}

func NewMongoDBRepository(client *mongo.Client) *MongoDBRepository {
	return &MongoDBRepository{client: client}
}

func (r *MongoDBRepository) GetAssetsByFilters(ctx context.Context, storeCode, platform, appID string) ([]AssetGroup, error) {
	collection := r.client.Database("crosspromo").Collection("appDynamicData")
	
	filter := bson.M{}
	if storeCode != "" {
		filter["storeCode"] = storeCode
	}
	if platform != "" {
		filter["platform"] = platform
	}
	if appID != "" {
		filter["appId"] = appID
	}

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var assets []CrosspromoFile
	if err = cursor.All(ctx, &assets); err != nil {
		return nil, err
	}

	return r.groupAssets(assets), nil
}

func (r *MongoDBRepository) groupAssets(assets []CrosspromoFile) []AssetGroup {
	groups := make(map[string]*AssetGroup)

	for _, asset := range assets {
		if _, exists := groups[asset.Localization]; !exists {
			groups[asset.Localization] = &AssetGroup{
				Localization: asset.Localization,
				Platforms:    make(map[string]PlatformData),
			}
		}

		group := groups[asset.Localization]
		if _, exists := group.Platforms[asset.Platform]; !exists {
			group.Platforms[asset.Platform] = PlatformData{
				Banner:       []AssetInfo{},
				Interstitial: []AssetInfo{},
				Rewarded:     []AssetInfo{},
			}
		}

		platformData := group.Platforms[asset.Platform]
		assetInfo := AssetInfo{
			FileName: asset.FileName,
			FileSize: float64(asset.FileSize) / 1024 / 1024, 
			URL:      asset.URL,
			MimeType: asset.MimeType,
		}

		switch asset.AdType {
		case "banner":
			platformData.Banner = append(platformData.Banner, assetInfo)
		case "interstitial":
			platformData.Interstitial = append(platformData.Interstitial, assetInfo)
		case "rewarded":
			platformData.Rewarded = append(platformData.Rewarded, assetInfo)
		}

		group.Platforms[asset.Platform] = platformData
	}

	result := make([]AssetGroup, 0, len(groups))
	for _, group := range groups {
		result = append(result, *group)
	}

	return result
}