package Services

import (
	"Api/Data/Request"
	"github.com/gin-gonic/gin"
)

type IUserService interface {
	LoginWithUser(email string, password string) (int, gin.H)
	RegisterUser(req Request.RegisterRequest) (int, gin.H)
}
