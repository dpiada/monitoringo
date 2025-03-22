package activator

import (
	"context"
	"log"
	"monitoringo/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Activator(client *mongo.Client) {

	log.Println("Start activator")
	// Get congifurations
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := client.Database("configurations").Collection("health_config").Find(ctx, bson.M{})
	if err != nil {
		log.Print(err)
		return
	}
	defer cursor.Close(ctx)

	var healthConfigs []models.HealthCfgResponse
	if err := cursor.All(ctx, &healthConfigs); err != nil {
		log.Print(err)
		return
	}

	for index, config := range healthConfigs {
		log.Print(index)
		log.Print(config)

		if !config.IsActive {
			continue
		}

		now := time.Now()
		lastUpdate := config.LastUpdate
		diff := now.Sub(lastUpdate)
		if diff.Seconds() > float64(config.Interval) {
			filter := bson.D{{"_id", id}} // Filter to match the document
			update := bson.D{{"$set", bson.D{{"fieldName", newValue}}}}

			result, err := client.Database("configurations").
				Collection("health_config").
				UpdateOne(ctx, filter, update)
		}
	}
}
