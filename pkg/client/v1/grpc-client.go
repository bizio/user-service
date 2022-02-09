package client

import (
	"context"
	"errors"
	"log"
	"time"

	v1 "github.com/bizio/user-service/pkg/api/v1"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UpdateUserNotification(refId string, status v1.UserNotificationStatus, reason string, sentAt time.Time, price float32) error {

	log.Printf("[GRPC-Client::UpdateUserNotification::debug] %v")
	userService, err := GetUserServiceClient()
	if err != nil {
		log.Printf("[UpdateUserNotification::error] %s", err)
		return err
	}

	sentAtProto, err := ptypes.TimestampProto(sentAt)
	if err != nil {
		log.Printf("[UpdateUserNotification::error] %s", err)
		return err
	}
	req := &v1.UserNotificationUpdateRequest{RefId: refId, Status: status, Reason: reason, SentAt: sentAtProto, Price: price}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("user", "api"))
	defer cancel()
	res, err := userService.UpdateUserNotification(ctx, req)
	if err != nil {
		log.Printf("[UpdateUserNotification::error] %s", err)
		return err
	}

	log.Printf("[UpdateUserNotification::res] %v", res)

	return nil

}

func GetUserServiceClient() (v1.UserServiceClient, error) {

	// use k8s internal dns
	address := "user-service.wa.internal:22002"
	// Set up a connection to the server.

	log.Printf("Connecting to user service on %s", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("Cannot connect to user service: %v", err)
		return nil, errors.New("Cannot connect to user service")
	}

	return v1.NewUserServiceClient(conn), err
}
