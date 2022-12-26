package userStorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/user/userModel"
)

func (db *mongoStore) UpdateResult(ctx context.Context, conditions interface{}, data *userModel.User) error {
	collection := db.db.Database("ManagerStudent").Collection("User")
	tmp, _ := bson.Marshal(data)
	var target bson.D

	_ = bson.Unmarshal(tmp, &target)
	update := bson.M{
		"$set": target,
	}
	_, err := collection.UpdateOne(ctx, conditions, update)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't update to DB, something DB is error")
		return solveError.ErrDB(err)
	}
	managerLog.InfoLogger.Println("update result success, storage return record and nil error")
	return nil
}
