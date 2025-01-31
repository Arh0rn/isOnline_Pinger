package mongo

import (
	"context"
	"fmt"
	"github.com/Arh0rn/isOnline_Pinger/config"
	"github.com/Arh0rn/isOnline_Pinger/models"
	"github.com/Arh0rn/isOnline_Pinger/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type MDB struct {
	client    *mongo.Client
	urls      *mongo.Collection
	params    *mongo.Collection
	idCounter *mongo.Collection
}

func NewMongoDB() storage.DB {
	return &MDB{}
}

var ctx = context.TODO()

func (mdb *MDB) ConnectDB(conf config.Config) error {
	uri := "mongodb://"
	dbName := os.Getenv("MONGO_DB_NAME")
	if os.Getenv("MONGO_USER") != "" && os.Getenv("MONGO_USER") != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
			os.Getenv("MONGO_USER"),
			os.Getenv("MONGO_PASSWORD"),
			os.Getenv("MONGO_HOST"),
			os.Getenv("MONGO_PORT"),
			dbName,
		)
	} else {
		// In case no username/password is needed, just construct URI without authentication.
		uri = fmt.Sprintf("mongodb://%s:%v/%s",
			os.Getenv("MONGO_HOST"),
			os.Getenv("MONGO_PORT"),
			dbName,
		)
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	mdb.client = client
	mdb.urls = client.Database(dbName).Collection("urls")
	mdb.params = client.Database(dbName).Collection("parameters")
	mdb.idCounter = client.Database(dbName).Collection("idCounter")
	return nil
}

func (mdb *MDB) CloseDB() error {
	err := mdb.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (mdb *MDB) GetUrls() ([]models.Url, error) {
	cursor, err := mdb.urls.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var urls []models.Url
	for cursor.Next(ctx) {
		var url models.Url
		if err := cursor.Decode(&url); err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}
	return urls, nil
}

func (mdb *MDB) AddUrl(url string) error {
	var c = struct {
		ID      int `bson:"id"` // Field 'id' from MongoDB maps to 'ID' in the struct
		Counter int `bson:"counter"`
	}{}

	err := mdb.idCounter.FindOne(ctx, bson.M{"id": 1}).Decode(&c)
	if err != nil {
		return err
	}

	_, err = mdb.urls.InsertOne(ctx, models.Url{URL: url, ID: c.Counter})
	if err != nil {
		return err
	}
	_, err = mdb.idCounter.UpdateOne(ctx, bson.M{"id": 1}, bson.M{"$inc": bson.M{"counter": 1}})
	if err != nil {
		return err
	}
	return nil
}

func (mdb *MDB) DeleteUrl(id int) error {
	_, err := mdb.urls.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}

func (mdb *MDB) GetParameters() (models.Parameters, error) {
	var p models.Parameters
	err := mdb.params.FindOne(ctx, bson.M{"id": 1}).Decode(&p)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (mdb *MDB) SetParameters(p models.Parameters) error {
	_, err := mdb.params.UpdateOne(
		ctx,
		bson.M{"id": 1},
		bson.M{"$set": bson.M{"timeout": p.Timeout, "interval": p.Interval, "workers": p.Workers}},
	)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	storage.RegisterDB("mongodb", NewMongoDB)
}
