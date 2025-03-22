package activator

import (
	"context"
	"log"
	"monitoringo/src/models"
	"net/http"
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
			NewRequest, err := http.NewRequest(config.Method, config.Path, http.NoBody)
			if err != nil {
				log.Println(err)
			}

			resp, err := http.DefaultClient.Do(NewRequest)
			if err != nil {
				log.Printf("[%d] Failed to send HTTP request: %v", index, err)
				continue
			}
			defer resp.Body.Close()

			if resp.StatusCode == config.Response.ResultCode {
				log.Println("beeeeello")
			}

		}
	}
}
