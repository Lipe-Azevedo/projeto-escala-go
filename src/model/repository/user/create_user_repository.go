package user

import (
	"context"
	"fmt"
	"os"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	userconv "github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity/converter/user" // IMPORT MODIFICADO
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(
	userDomain domain.UserDomainInterface,
) (domain.UserDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init CreateUser repository",
		zap.String("journey", "createUser"),
		zap.String("email", userDomain.GetEmail()))

	collectionName := os.Getenv(MONGODB_USERS_COLLECTION_ENV_KEY)
	if collectionName == "" {
		errMsg := fmt.Sprintf("Environment variable %s not set for users collection name", MONGODB_USERS_COLLECTION_ENV_KEY)
		logger.Error(errMsg, nil, zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError("database configuration error: users collection name not set")
	}
	collection := ur.databaseConnection.Collection(collectionName)

	value := userconv.ConvertDomainToEntity(userDomain) // USO MODIFICADO

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		if writeException, ok := err.(mongo.WriteException); ok {
			for _, writeError := range writeException.WriteErrors {
				if writeError.Code == 11000 {
					errorMessage := fmt.Sprintf("User with email %s already exists or another unique constraint violated.", value.Email)
					logger.Error(errorMessage, err,
						zap.String("journey", "createUser"),
						zap.String("email", value.Email))
					return nil, rest_err.NewConflictError(errorMessage)
				}
			}
		}
		logger.Error(
			"Error trying to create user in repository",
			err,
			zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)
	logger.Info(
		"CreateUser repository executed successfully",
		zap.String("userId", value.ID.Hex()),
		zap.String("journey", "createUser"),
	)

	return userconv.ConvertEntityToDomain(*value), nil // USO MODIFICADO
}
