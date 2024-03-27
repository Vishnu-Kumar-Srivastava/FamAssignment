package daos

import (
	"context"
	"ytvideofetcher/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type YtVideoDAO struct{}

func NewYtVideoDAO() IYtVideoDAO {
	return &YtVideoDAO{}
}

var videosCollection *mongo.Collection = OpenCollection(Client, "Videos")

func (dao *YtVideoDAO) UpsertVideos(ctx context.Context, response *models.Response) error {
	var operations []mongo.WriteModel

	for _, video := range response.Items {
		filter := bson.M{"videoId": video.ID.VideoID, "etag":video.Etag}
		update := bson.M{"$set": video}
		model := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true)
		operations = append(operations, model)

	}
	_, err := videosCollection.BulkWrite(context.Background(), operations)

	if err != nil {
		return err
	}

	return nil

}
func (dao *YtVideoDAO) GetVideos(ctx context.Context) ([]*models.Video, error) {
	var videos []*models.Video
	cursor, err := videosCollection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &videos); err != nil {
		return nil, err
	}
	return videos, nil
}
