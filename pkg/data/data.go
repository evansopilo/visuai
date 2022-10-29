package data

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrNoDocument     = errors.New("error no document")
	ErrCreateDocument = errors.New("error create document")
	ErrUpdateDocument = errors.New("error update document")
	ErrDeleteDocument = errors.New("error delete document")
)

type Post struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Title     string    `json:"title,omitempty" bson:"title,omitempty"`
	Desc      string    `json:"desc,omitempty" bson:"desc,omitempty"`
	PhotoURL  string    `json:"photo_url,omitempty" bson:"photo_url,omitempty"`
	DestURL   string    `json:"dest_url,omitempty" bson:"dest_url,omitempty"`
	Category  string    `json:"category,omitempty" bson:"category,omitempty"`
	GeoTag    GeoTag    `json:"geo_tag,omitempty" bson:"geo_tag,omitempty"`
	Tags      []string  `json:"tags,omitempty" bson:"tags,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type GeoTag struct {
	Type        string    `json:"type,omitempty" bson:"type,omitempty"`
	Coordinates []float64 `json:"coordinates,omitempty" bson:"coordinates,omitempty"`
}

type PostModel struct {
	client *mongo.Client
}

func NewPostModel(client *mongo.Client) *PostModel { return &PostModel{client: client} }

func (p PostModel) Create(ctx context.Context, post *Post) error {

	coll := p.client.Database("visuai").Collection("posts")

	result, err := coll.InsertOne(ctx, post)
	if err != nil {
		return err
	}

	if result.InsertedID.(string) != post.ID {
		return ErrCreateDocument
	}

	return nil
}

func (p PostModel) GetByID(ctx context.Context, id string) (*Post, error) {

	coll := p.client.Database("visuai").Collection("posts")

	var post Post

	if err := coll.FindOne(ctx, bson.M{"_id": id}).Decode(&post); err != nil {
		return nil, err
	}

	return &post, nil
}

func (p PostModel) GetByUserID(ctx context.Context, id string, skip, limit int64) (*[]Post, error) {

	opts := options.Find().SetSkip(skip).SetLimit(limit)

	coll := p.client.Database("visuai").Collection("posts")

	filterCursor, err := coll.Find(ctx, bson.M{"user_id": id}, opts)
	if err != nil {
		return nil, err
	}

	var posts []Post

	if err := filterCursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return &posts, nil
}

func (p PostModel) GetByCategory(ctx context.Context, category string, skip, limit int64) (*[]Post, error) {

	opts := options.Find().SetSkip(skip).SetLimit(limit)

	coll := p.client.Database("visuai").Collection("posts")

	filterCursor, err := coll.Find(ctx, bson.M{"category": category}, opts)
	if err != nil {
		return nil, err
	}

	var posts []Post

	if err := filterCursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return &posts, nil
}

func (p PostModel) GetByTags(ctx context.Context, tags []string, skip, limit int64) (*[]Post, error) {

	opts := options.Find().SetSkip(skip).SetLimit(limit)

	coll := p.client.Database("visuai").Collection("posts")

	filterCursor, err := coll.Find(ctx, bson.M{"tags": bson.M{"$all": tags}}, opts)
	if err != nil {
		return nil, err
	}

	var posts []Post

	if err := filterCursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return &posts, nil
}

func (p PostModel) Get(ctx context.Context, skip, limit int64) (*[]Post, error) {

	opts := options.Find().SetSkip(skip).SetLimit(limit)

	coll := p.client.Database("visuai").Collection("posts")

	filterCursor, err := coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}

	var posts []Post

	if err := filterCursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return &posts, nil
}

func (p PostModel) UpdateByID(ctx context.Context, id string, post *Post) error {

	coll := p.client.Database("visuai").Collection("posts")

	result, err := coll.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.D{{Key: "$set", Value: post}})
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return ErrNoDocument
	}

	return nil
}

func (p PostModel) DeleteByID(ctx context.Context, id string) error {

	coll := p.client.Database("visuai").Collection("posts")

	result, err := coll.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return ErrNoDocument
	}

	return nil
}

func (p PostModel) DeleteByUserID(ctx context.Context, id string) error {

	coll := p.client.Database("visuai").Collection("posts")

	result, err := coll.DeleteMany(ctx, bson.M{"user_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return ErrNoDocument
	}

	return nil
}
