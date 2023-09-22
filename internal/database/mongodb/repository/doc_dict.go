package repository

import (
	"context"
	"errors"
	"projects/grafit_info/config"
	"projects/grafit_info/internal/database/mongodb"
	"projects/grafit_info/internal/database/mongodb/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type DocumentDictRepo struct {
	mongoClient *mongodb.Client
	cfg         *config.Config
}

type DocumentDict interface {
	Find() ([]*models.DocumentDict, error)
	FindByName(req *models.FindReq) (*models.DocumentDict, error)
}

func NewDocumentDictRepo(cfg *config.Config, mongoClient *mongodb.Client) DocumentDict {
	return DocumentDictRepo{
		mongoClient: mongoClient,
		cfg:         cfg,
	}
}

func (d DocumentDictRepo) Find() ([]*models.DocumentDict, error) {
	var docs []*models.DocumentDict
	timeout := time.Duration(d.cfg.MongoDB.Timeout)

	collection := d.mongoClient.Client.Database("grafit").Collection("document_dict", nil)

	filter := bson.D{}

	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()
	cur, err := collection.Find(timeoutCtx, filter)
	if cur == nil {
		return nil, errors.New("Get document dictionary result is empty")
	}
	if err != nil {
		return nil, err
	}

	timeoutCtx, cancelFunc = context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()
	if err := cur.All(timeoutCtx, &docs); err != nil {
		return nil, err
	}

	return docs, nil
}

func (d DocumentDictRepo) FindByName(req *models.FindReq) (*models.DocumentDict, error) {
	var doc *models.DocumentDict
	timeout := time.Duration(d.cfg.MongoDB.Timeout)

	collection := d.mongoClient.Client.Database("grafit").Collection("document_dict", nil)
	filter := bson.D{
		{"name", req.Name},
	}

	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()
	err := collection.FindOne(timeoutCtx, filter).Decode(&doc)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
