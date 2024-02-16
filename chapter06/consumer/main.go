package main

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Request struct {
	URL string `json:"url"`
}

type Feed struct {
	Entries []Entry `xml:"entry"`
}

type Entry struct {
	Link struct {
		Href string `xml:"href,attr"`
	} `xml:"link"`
	Content string `xml:"content"`
	Title   string `xml:"title"`
}

func GetFeedEntries(feedURL string) ([]Entry, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteValue, _ := io.ReadAll(resp.Body)
	var feed Feed
	xml.Unmarshal(byteValue, &feed)

	imgRegex := regexp.MustCompile(`https://i\.redd\.it/\w+\.(jpeg|jpg|png|gif)`)

	for i, entry := range feed.Entries {
		matches := imgRegex.FindStringSubmatch(entry.Content)
		if len(matches) > 0 {
			feed.Entries[i].Content = matches[0]
		} else {
			feed.Entries[i].Content = "default_img_url"
		}
	}

	return feed.Entries, nil
}

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal("Error loading the .env file")
	}
	ctx := context.Background()
	mongoClient, _ := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	defer mongoClient.Disconnect(ctx)

	amqpConnection, err := amqp.Dial(os.Getenv("RABBITMQ_URI"))
	if err != nil {
		log.Fatal(err)
	}
	defer amqpConnection.Close()

	channelAmqp, _ := amqpConnection.Channel()
	defer channelAmqp.Close()

	forever := make(chan bool)

	msgs, err := channelAmqp.Consume(
		os.Getenv("RABBITMQ_QUEUE"),
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			var request Request
			json.Unmarshal(d.Body, &request)

			log.Println("RSS URL:", request.URL)

			entries, _ := GetFeedEntries(request.URL)

			fmt.Println(entries)

			collection := mongoClient.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")
			fmt.Println(len(entries))
			for _, entry := range entries {
				thumbnail := entry.Content
				if thumbnail == "" {
					thumbnail = "default_image_url"
				}
				collection.InsertOne(ctx, bson.M{
					"title":     entry.Title,
					"thumbnail": thumbnail,
					"url":       entry.Link.Href,
				})
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
