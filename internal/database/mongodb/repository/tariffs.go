package repository

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"projects/grafit_info/config"
	"projects/grafit_info/internal/database/mongodb"
	"projects/grafit_info/internal/database/mongodb/models"
)

type TariffRepo struct {
	mongoClient *mongodb.Client
	cfg         *config.Config
}

func NewTariffRepository(cfg *config.Config, mongoClient *mongodb.Client) *TariffRepo {
	return &TariffRepo{
		mongoClient: mongoClient,
		cfg:         cfg,
	}
}

func (t TariffRepo) Find() (models.Tariffs, error) {
	var tariffs models.Tariffs
	timeout := time.Duration(t.cfg.MongoDB.Timeout)

	collection := t.mongoClient.Client.Database("grafit").Collection("tariffs", nil)
	filter := bson.D{}

	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()
	cur, err := collection.Find(timeoutCtx, filter)
	if cur == nil {
		return nil, errors.New("getTariffs result is empty")
	}
	if err != nil {
		return nil, err
	}

	timeoutCtx, cancelFunc = context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()
	if err := cur.All(timeoutCtx, &tariffs); err != nil {
		return nil, err
	}

	return tariffs, nil
}

func (t TariffRepo) FindByID(req *models.TariffRequest) (*models.Tariff, error) {
	var tariff *models.Tariff
	timeout := time.Duration(t.cfg.MongoDB.Timeout)

	collection := t.mongoClient.Client.Database("grafit").Collection("tariffs", nil)
	filter := bson.D{
		{"_id", req.ID},
	}

	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()
	err := collection.FindOne(timeoutCtx, filter).Decode(&tariff)
	if err != nil {
		return nil, err
	}

	return tariff, nil
}

func (t TariffRepo) Update(req models.TariffRequest, newTariff *models.Tariff) error {
	timeout := time.Duration(t.cfg.MongoDB.Timeout)

	collection := t.mongoClient.Client.Database("grafit").Collection("tariffs", nil)
	filter := bson.D{
		{"_id", req.ID},
	}

	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()
	res, err := collection.UpdateOne(timeoutCtx, filter, newTariff)
	if res == nil {
		return errors.New("tariff not found")
	}
	if err != nil {
		return err
	}

	return nil
}

func (t TariffRepo) Delete(req models.TariffRequest) error {
	timeout := time.Duration(t.cfg.MongoDB.Timeout)

	collection := t.mongoClient.Client.Database("grafit").Collection("tariffs", nil)
	filter := bson.D{
		{"_id", req.ID},
	}

	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()
	res, err := collection.DeleteOne(timeoutCtx, filter)
	if res == nil {
		return errors.New("tariff not found")
	}
	if err != nil {
		return err
	}

	return nil
}

func (t TariffRepo) Create(tariff *models.Tariff) error {
	timeout := time.Duration(t.cfg.MongoDB.Timeout)

	collection := t.mongoClient.Client.Database("grafit").Collection("tariffs", nil)

	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()
	res, err := collection.InsertOne(timeoutCtx, tariff)
	if res == nil {
		return errors.New("tariff not created")
	}

	if err != nil {
		return err
	}

	return nil
}
