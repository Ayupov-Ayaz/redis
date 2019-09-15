package usecase

import (
	"context"
	"github.com/ayupov-ayaz/redis/mocks"
	"github.com/ayupov-ayaz/redis/models"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestBaseUserUsecase(t *testing.T){
	Convey("test user base_usecase", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repositoryMock := mocks.NewMockUserRepository(ctrl)
		usecase := NewBaseUserUsecase(repositoryMock, 20 * time.Second)

		Convey("Create user", func() {
			first := repositoryMock.EXPECT().GetByEmail(gomock.Any(), gomock.Any()).Return(nil, nil)
			repositoryMock.EXPECT().Create(gomock.Any(),gomock.Any()).Return(nil).After(first)
			pass := "qwerty"
			u :=  &models.User{
				Name:     "u1",
				Email:    "u1@gmail.com",
				Password: pass,
			}
			err := usecase.Create(context.TODO(), u)
			So(err, ShouldBeNil)
			So(pass, ShouldNotEqual, u.Password)
		})

		Convey("Get existing user", func() {
			repositoryMock.EXPECT().Get(gomock.Any(), uint64(1)).Times(1).Return(&models.User{}, nil)
			user, err := usecase.Get(context.TODO(), 1)
			So(err, ShouldBeNil)
			So(user, ShouldNotBeNil)
		})

		Convey("Get not existing user", func() {
			repositoryMock.EXPECT().Get(gomock.Any(), uint64(1)).Times(1).Return(nil, models.UserNotFound)
			user, err := usecase.Get(context.TODO(), 1)
			So(err, ShouldNotBeNil)
			So(user, ShouldBeNil)
			So(err, ShouldEqual, models.UserNotFound)
		})

		Convey("Update existing user", func() {
			user := &models.User{
				ID:       1,
				Name:     "u1",
				Email:    "u1@gmail.com",
				Password: "pass",
			}
			repositoryMock.EXPECT().Update(gomock.Any(), gomock.Any()).Times(1).Return(nil)
			err := usecase.Update(context.TODO(), user)
			So(err, ShouldBeNil)
		})

		Convey("Delete user", func() {
			repositoryMock.EXPECT().Delete(gomock.Any(), uint64(1)).Times(1).Return(nil)
			err := usecase.Delete(context.TODO(), 1)
			So(err, ShouldBeNil)
		})
	})
}