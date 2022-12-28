package studentBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/student/studentModel"
)

type GetStudentStore interface {
	FindStudent(ctx context.Context, conditions interface{}, location string) (*studentModel.Student, error)
}

type getStudentBiz struct {
	store GetStudentStore
}

func NewGetStudent(store GetStudentStore) *getStudentBiz {
	return &getStudentBiz{store: store}
}

func (biz *getStudentBiz) GetStudent(ctx context.Context, filter interface{}) (*studentModel.Student, error) {
	data, err := biz.store.FindStudent(ctx, bson.M{"id": filter}, studentModel.StudentCollection)
	if err != nil {
		managerLog.ErrorLogger.Println("Some thing error in storage user, may be from database")
		return nil, solveError.ErrDB(err)
	}

	managerLog.InfoLogger.Println("Get student ok")
	return data, nil
}
