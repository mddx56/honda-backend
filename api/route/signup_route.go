package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/waltherx/honda-backend/api/controller"
	"github.com/waltherx/honda-backend/bootstrap"
	"github.com/waltherx/honda-backend/domain"
	"github.com/waltherx/honda-backend/mongo"
	"github.com/waltherx/honda-backend/repository"
	"github.com/waltherx/honda-backend/usecase"
)

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
