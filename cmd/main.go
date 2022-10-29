package main

import (
	"context"

	"os"
	"time"

	"github.com/evansopilo/visuai/pkg/blob"
	"github.com/evansopilo/visuai/pkg/data"
	"github.com/evansopilo/visuai/pkg/log"
	"github.com/evansopilo/visuai/pkg/secret"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type App struct {
	Models    data.Models
	BlobModel interface {
		UploadBytesToBlob(data []byte, metadata map[string]string) (string, error)
	}
	Logger interface {
		Trace(args string, fields map[string]interface{})

		Debug(args string, fields map[string]interface{})

		Info(args string, fields map[string]interface{})

		Warn(args string, fields map[string]interface{})

		Error(args string, fields map[string]interface{})

		Fatal(msg string, fields map[string]interface{})

		Panic(args string, fields map[string]interface{})
	}
}

func main() {

	logger := log.New("json", os.Stdout, -1)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	logger.Info("connect to azure key vault with the uri obtained from the application env variables", nil)

	secret, err := secret.New(os.Getenv("azure_vault_uri"))
	if err != nil {
		logger.Fatal("failed to connect to azure key vault with the uri obtained from the application env variables", nil)
	}

	logger.Info("connected to azure key vault with the uri obtained from the application env variables", nil)

	logger.Info("get to get mongodb database uri provided in azure key vault", nil)

	mongoUri, err := secret.GetSecret(ctx, "mongo_uri", "")
	if err != nil {
		logger.Fatal("failed to get mongodb database uri provided in azure key vault", nil)
	}

	logger.Info("connect to mongodb database with the uri secret obtained from azure key vault", nil)

	opts := options.Client().ApplyURI(*mongoUri)

	client, err := mongo.NewClient(opts)
	if err != nil {
		logger.Fatal("failed to get mongodb database uri provided in azure key vault", nil)
	}

	if err := client.Connect(ctx); err != nil {
		logger.Fatal("failed to get mongodb database uri provided in azure key vault", nil)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Fatal("failed to get mongodb database uri provided in azure key vault", nil)
	}

	logger.Info("connected to mongodb database with the uri secret obtained from azure key vault", nil)

	logger.Info("get blob endpoint of azure blob provided in azure key vault", nil)

	blobEndpoint, err := secret.GetSecret(ctx, "blob_endpoint", "")
	if err != nil {
		logger.Fatal("failed get blob endpoint of azure blob provided in azure key vault", nil)
	}

	logger.Info("get blob container of azure blob provided in azure key vault", nil)

	blobContainer, err := secret.GetSecret(ctx, "blob_container", "")
	if err != nil {
		logger.Info("failed to get blob container of azure blob provided in azure key vault", nil)
	}

	logger.Info("get azure key of azure blob provided in azure key vault", nil)

	blobAzrKey, err := secret.GetSecret(ctx, "blob_azr_key", "")
	if err != nil {
		logger.Fatal("failed to get azure key of azure blob provided in azure key vault", nil)
	}

	logger.Info("get azure account name of azure blob provided in azure key vault", nil)

	accountName, err := secret.GetSecret(ctx, "account_name", "")
	if err != nil {
		logger.Info("failed get azure account name of azure blob provided in azure key vault", nil)
	}

	app := App{
		Models: data.Models{
			Post: data.NewPostModel(client),
		},
		BlobModel: blob.New(*blobEndpoint, *blobContainer, *blobAzrKey, *accountName),
		Logger:    logger,
	}

	logger.Info("start application server to listen to port: 8080", nil)
	if err := app.Router().Listen(":8080"); err != nil {
		logger.Info("failed to start application server to listen to port: 8080", nil)
	}
}
