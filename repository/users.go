// package repository

// import (
// 	"context"

// 	"github.com/NachooNazar/prueba-tecnica-backend/models"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// const UsersCollections = "users"

// type UserRepository interface{}

// type userRepository struct {
// 	c *mongo.Collection
// }

// func NewUserRepository(client *mongo.Client) UserRepository {
// 	return &userRepository{(*mongo.Collection)(client.Database("gomongo").Collection("users"))}
// }

// func (r *userRepository) Save(user *models.User) error {
// 	res, err := r.c.InsertOne(context.TODO(), user)
// 	if err != nil {
// 		return err
// 	}
// 	return res
// }

// func (r *userRepository) getAll(user *models.User) error {

// }