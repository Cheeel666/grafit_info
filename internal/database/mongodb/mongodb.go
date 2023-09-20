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
	Client *mongo.Client
}

func NewMongoDB(cfg *config.Config, log *zap.Logger) *Client {
	return &Client{
		cfg:    cfg,
		log:    log,
		Client: nil,
	}
}

func (m *Client) Connect() error {
	var err error

	uri := fmt.Sprintf(
		"mongodb://%s:%d/%s?authSource=admin",
		m.cfg.MongoDB.Host,
		m.cfg.MongoDB.Port,
		m.cfg.MongoDB.Database,
	)

	if len(m.cfg.MongoDB.Username) > 0 && len(m.cfg.MongoDB.Password) > 0 {
		uri = fmt.Sprintf(
			"mongodb://%s:%s@%s:%d/%s?authSource=admin",
			m.cfg.MongoDB.Username,
			m.cfg.MongoDB.Password,
			m.cfg.MongoDB.Host,
			m.cfg.MongoDB.Port,
			m.cfg.MongoDB.Database,
		)

	}

	opts := options.Client().ApplyURI(uri)
	m.Client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	var result bson.M
	if err := m.Client.Database("grafit").RunCommand(context.TODO(), bson.M{"ping": 1}).Decode(&result); err != nil {
		return err
	}

	return nil
}

func (m *Client) Release() {
	if m.Client != nil {
		m.Client.Disconnect(context.TODO())
	}
}
