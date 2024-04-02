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

func NewClientRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewClientRepository(db, domain.CollectionClient)
	tc := &controller.ClientController{
		ClientUsecase: usecase.NewClientUsecase(tr, timeout),
	}
	group.GET("/client", tc.Fetch)
	group.POST("/client", tc.Create)
}
