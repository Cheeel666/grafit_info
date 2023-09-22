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

type CommonInfoRepo struct {
	mongoClient *mongodb.Client
	cfg         *config.Config
}

func NewCommonInfoRepo(cfg *config.Config, mongoClient *mongodb.Client) *CommonInfoRepo {
	return &CommonInfoRepo{
		mongoClient: mongoClient,
		cfg:         cfg,
	}
}

func (r CommonInfoRepo) Find() (*models.CommonInfo, error) {
	var info *models.CommonInfo
	timeout := time.Duration(r.cfg.MongoDB.Timeout)

	collection := r.mongoClient.Client.Database("grafit").Collection("common_info", nil)
	filter := bson.D{}

	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()
	cur, err := collection.Find(timeoutCtx, filter)
	if cur == nil {
		return nil, errors.New("getInfo result is empty")
	}
	if err != nil {
		return nil, err
	}

	timeoutCtx, cancelFunc = context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()
	if err := cur.All(timeoutCtx, &info); err != nil {
		return nil, err
	}

	return info, nil
}

func (r CommonInfoRepo) FindByID(req *models.FindReq) (*models.CommonInfo, error) {
	var info *models.CommonInfo
	timeout := time.Duration(r.cfg.MongoDB.Timeout)

	collection := r.mongoClient.Client.Database("grafit").Collection("common_info", nil)
	filter := bson.D{
		{"_id", req.ID},
	}

	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()
	err := collection.FindOne(timeoutCtx, filter).Decode(&info)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (r CommonInfoRepo) FindByType(req *models.FindReq) (*models.CommonInfo, error) {
	var info *models.CommonInfo
	timeout := time.Duration(r.cfg.MongoDB.Timeout)

	collection := r.mongoClient.Client.Database("grafit").Collection("common_info", nil)
	filter := bson.D{
		{"_id", req.ID},
	}

	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()
	err := collection.FindOne(timeoutCtx, filter).Decode(&info)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (r CommonInfoRepo) Update(req models.FindReq, newTariff *models.CommonInfo) error {
	timeout := time.Duration(r.cfg.MongoDB.Timeout)

	collection := r.mongoClient.Client.Database("grafit").Collection("common_info", nil)
	filter := bson.D{
		{"_id", req.ID},
	}

	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()
	res, err := collection.UpdateOne(timeoutCtx, filter, newTariff)
	if res == nil {
		return errors.New("info not found")
	}
	if err != nil {
		return err
	}

	return nil
}

func (r CommonInfoRepo) Delete(req models.FindReq) error {
	timeout := time.Duration(r.cfg.MongoDB.Timeout)

	collection := r.mongoClient.Client.Database("grafit").Collection("common_info", nil)
	filter := bson.D{
		{"_id", req.ID},
	}

	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()
	res, err := collection.DeleteOne(timeoutCtx, filter)
	if res == nil {
		return errors.New("info not found")
	}
	if err != nil {
		return err
	}

	return nil
}

func (r CommonInfoRepo) Create(info *models.CommonInfo) error {
	timeout := time.Duration(r.cfg.MongoDB.Timeout)

	collection := r.mongoClient.Client.Database("grafit").Collection("common_info", nil)

	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()
	res, err := collection.InsertOne(timeoutCtx, info)
	if res == nil {
		return errors.New("info not created")
	}

	if err != nil {
		return err
	}

	return nil
}
