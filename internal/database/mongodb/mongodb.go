package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"projects/grafit_info/config"
)

type Client struct {
	cfg    *config.Config
	log    *zap.Logger
	client *mongo.Client
}

func NewMongoDB(cfg *config.Config, log *zap.Logger) *Client {
	return &Client{
		cfg:    cfg,
		log:    log,
		client: nil,
	}
}

func (m *Client) Connect() error {
	var err error
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d/%s?authSource=admin",
		m.cfg.MongoDB.Username,
		m.cfg.MongoDB.Password,
		m.cfg.MongoDB.Host,
		m.cfg.MongoDB.Port,
		m.cfg.MongoDB.Database,
	)

	opts := options.Client().ApplyURI(uri)
	m.client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	var result bson.M
	if err := m.client.Database("admin").RunCommand(context.TODO(), bson.M{"ping": 1}).Decode(&result); err != nil {
		return err
	}

	return nil
}

func (m *Client) Release() {
	if m.client != nil {
		m.client.Disconnect(context.TODO())
	}
}
