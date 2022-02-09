package cloudpubsub

import (
	"context"
	"log"
	"sync"
	"time"

	"cloud.google.com/go/pubsub"
)

const ReadTopic = "alerts"
const WriteTopic = "user-notifications"
const PullSubscription = "user-alerts"

type Queue interface {
	GetMessages(messages chan *pubsub.Message) error
	GetTopic(topic string) (*pubsub.Topic, error)
}

type MessageQueue struct {
	client        *pubsub.Client
	ReadTopic     *pubsub.Topic
	WriteTopic    *pubsub.Topic
	UserAlertsSub *pubsub.Subscription
}

func NewMessageQueue(projectId string) *MessageQueue {

	var err error
	var sub *pubsub.Subscription

	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Cannot create pubsub client: '%s'", err)
	}

	readTopic, err := getTopic(ReadTopic, pubsubClient)
	if err != nil {
		log.Fatal("Error checking read topic: %s", err)
	}

	writeTopic, err := createTopicIfNotExists(WriteTopic, pubsubClient)
	if err != nil {
		log.Fatal("Error checking read topic: %s", err)
	}

	// [START pubsub_create_pull_subscription]
	sub, err = pubsubClient.CreateSubscription(ctx, PullSubscription, pubsub.SubscriptionConfig{
		Topic:       readTopic,
		AckDeadline: 20 * time.Second,
	})
	if err != nil {
		sub = pubsubClient.Subscription(PullSubscription)
		log.Printf("Cannot create subscription: %s", err)
	}

	return &MessageQueue{
		pubsubClient,
		readTopic,
		writeTopic,
		sub,
	}

}

func (q *MessageQueue) GetMessages(messages chan *pubsub.Message) error {

	var mu sync.Mutex
	var err error

	ctx, cancel := context.WithCancel(context.Background())
	err = q.UserAlertsSub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		mu.Lock()
		defer mu.Unlock()
		log.Printf("Got message: %q\n", string(msg.Data))
		messages <- msg
	})

	if err != nil {
		cancel()
	}

	return err
}

func (q *MessageQueue) GetTopic(topic string) (*pubsub.Topic, error) {
	return getTopic(topic, q.client)
}

func getTopic(topic string, c *pubsub.Client) (*pubsub.Topic, error) {

	ctx := context.Background()
	t := c.Topic(topic)
	ok, err := t.Exists(ctx)
	if err != nil {
		return nil, err
	}

	if ok {
		return t, nil
	}

	return nil, nil

}

func createTopicIfNotExists(topic string, c *pubsub.Client) (*pubsub.Topic, error) {

	var t *pubsub.Topic
	var err error

	ctx := context.Background()

	t, err = getTopic(topic, c)
	if err != nil {
		log.Printf("Failed to get the topic: %v", err)
		return nil, err
	}

	// Create a topic to subscribe to.
	if t == nil {
		t, err = c.CreateTopic(ctx, topic)
		if err != nil {
			log.Printf("Failed to create the topic: %v", err)
			return nil, err
		}
	}

	return t, nil
}
