package repository

import (
	"context"
	"errors"
	"time"

	"projects/grafit_info/config"
	"projects/grafit_info/internal/database/mongodb"
	"projects/grafit_info/internal/database/mongodb/models"

	"go.mongodb.org/mongo-driver/bson"
)

type CommonInfoRepo struct {
	mongoClient *mongodb.Client
	cfg         *config.Config
}

type CommonInfo interface {
	Find() (*models.CommonInfo, error)
	FindByType(req *models.FindReq) (*models.CommonInfo, error)
	Delete(req *models.FindReq) error
	Update(*models.FindReq, *models.CommonInfo) error
	Create(*models.CommonInfo) error
}

func NewCommonInfoRepo(cfg *config.Config, mongoClient *mongodb.Client) CommonInfo {
	return CommonInfoRepo{
		mongoClient: mongoClient,
		cfg:         cfg,
	}
}

func (c CommonInfoRepo) Find() (*models.CommonInfo, error) {
	var info *models.CommonInfo
	timeout := time.Duration(c.cfg.MongoDB.Timeout)

	collection := c.mongoClient.Client.Database("grafit").Collection("common_info", nil)
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

func (c CommonInfoRepo) FindByID(req *models.FindReq) (*models.CommonInfo, error) {
	var info *models.CommonInfo
	timeout := time.Duration(c.cfg.MongoDB.Timeout)

	collection := c.mongoClient.Client.Database("grafit").Collection("common_info", nil)
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

func (c CommonInfoRepo) FindByType(req *models.FindReq) (*models.CommonInfo, error) {
	var info *models.CommonInfo
	timeout := time.Duration(c.cfg.MongoDB.Timeout)

	collection := c.mongoClient.Client.Database("grafit").Collection("common_info", nil)
	filter := bson.D{
		{"type", req.Type},
	}

	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()
	err := collection.FindOne(timeoutCtx, filter).Decode(&info)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (c CommonInfoRepo) Update(req *models.FindReq, NewDoc *models.CommonInfo) error {
	timeout := time.Duration(c.cfg.MongoDB.Timeout)

	collection := c.mongoClient.Client.Database("grafit").Collection("common_info", nil)
	filter := bson.D{
		{"_id", req.ID},
	}

	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()
	res, err := collection.UpdateOne(timeoutCtx, filter, NewDoc)
	if res == nil {
		return errors.New("info not found")
	}
	if err != nil {
		return err
	}

	return nil
}

func (c CommonInfoRepo) Delete(req *models.FindReq) error {
	timeout := time.Duration(c.cfg.MongoDB.Timeout)

	collection := c.mongoClient.Client.Database("grafit").Collection("common_info", nil)
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

func (c CommonInfoRepo) Create(info *models.CommonInfo) error {
	timeout := time.Duration(c.cfg.MongoDB.Timeout)

	collection := c.mongoClient.Client.Database("grafit").Collection("common_info", nil)

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
