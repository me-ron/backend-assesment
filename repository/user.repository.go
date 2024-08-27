package repository

import (
	"context"
	"errors"
	"loan_tracker/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepo struct {
	Collection domain.CollectionInterface
}

func (repo *UserRepo) EnsureIndexes() error {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1}, 
		Options: options.Index().SetUnique(true),
	}
	
	_, err := repo.Collection.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}

func NewUserRepo(coll domain.CollectionInterface) (*UserRepo, error) {
	repo := &UserRepo{
		Collection : coll,
	}

	// Ensure indexes are created
	if err := repo.EnsureIndexes(); err != nil {
		return nil, err
	}

	return repo, nil
}

func (repo *UserRepo) SaveUser(user *domain.UserInfo) error {
	_, err := repo.Collection.InsertOne(context.TODO(), user)
	return err
}

func (repo *UserRepo) FindUserByEmail(email string) (*domain.UserInfo, error) {
	var user domain.UserInfo
	err := repo.Collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepo)GetUserDocumentByID(id string) (domain.UserResponse , error) {
	pid,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.UserResponse{},err
	}
	filter := bson.D{{Key: "_id" , Value: pid}}
	result := repo.Collection.FindOne(context.TODO(),filter)
	
	var user domain.UserResponse
	err = result.Decode(&user)
	if err != nil {
		return domain.UserResponse{},err
	}
	return user,nil
}

func (repo *UserRepo)GetUserDocuments() ([]domain.UserResponse , error) {
	filter := bson.D{{}}
	cursor,err := repo.Collection.Find(context.TODO() , filter)

	if err != nil {
		return []domain.UserResponse{},err
	}

	users := []domain.UserResponse{}
	for cursor.Next(context.TODO())  {
		var user domain.UserResponse
		err := cursor.Decode(&user)
		if err != nil {
			return []domain.UserResponse{},err
		}
		users = append(users, user)
	}

	return users , nil
}

func (repo *UserRepo)DeleteUserDocument(id string) (error) {
	pid,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id" , Value: pid}}
	_,err = repo.Collection.DeleteOne(context.TODO() ,filter)
	return err
}

func (repo *UserRepo)GetBools(id string) (domain.Bools, error){
	pid,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Bools{},err
	}
	filter := bson.D{{Key: "_id" , Value: pid}}
	result := repo.Collection.FindOne(context.TODO(),filter)
	
	var user domain.Bools
	err = result.Decode(&user)
	if err != nil {
		return domain.Bools{},err
	}

	return user, nil
}

func (repo *UserRepo) VerifyUser(id string) error {
	objID,_ := primitive.ObjectIDFromHex(id) 
	filter := bson.D{{Key: "_id" , Value: objID}}
	setter := bson.D{{Key:"$set" , Value: bson.D{{Key:"Verified" , Value: true}, {Key: "VerifiedCode", Value: ""}}}}

	_,err := repo.Collection.UpdateOne(context.TODO() , filter , setter)

	return err
}

func (repo *UserRepo) UpdatePassword(email string, password string) error {
	filter := bson.D{{Key: "email" , Value: email}}
	setter := bson.D{{Key:"$set" , Value: bson.D{{Key:"password" , Value: password}}}}

	_,err := repo.Collection.UpdateOne(context.TODO() , filter , setter)

	return err
}

