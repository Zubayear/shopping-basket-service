package publisher

import (
	"ShoppingBasket/model"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"go-micro.dev/v4/broker"
	log "go-micro.dev/v4/logger"
	"time"
)

func PublishMessage(msg model.BasketCheckoutMessage, topic string) error {
	log.Infof("Publishing message: %v", msg)
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	m := &broker.Message{
		Header: map[string]string{
			"id":   id.String(),
			"time": time.Now().String(),
		},
		Body: body,
	}
	err = broker.Publish(topic, m)
	if err != nil {
		return fmt.Errorf("failed publishing message: %w", err)
	}
	return nil
}
