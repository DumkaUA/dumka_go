package middleware

import (
	"github.com/DumkaUA/dumka_go/src/model"
	"github.com/DumkaUA/dumka_go/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var q struct {
			Token string `form:"token" binding:"required"`
		}

		if err := c.ShouldBind(&q); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"ok":         false,
				"error":      err.Error(),
				"error_code": -2,
			})
			return
		}

		db := util.GetDB(c)

		var user model.User
		db.Where(model.User{Token: q.Token}).First(&user)

		if user.Id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"ok":         false,
				"error":      "USER WITH SUCH TOKEN DOES NOT EXIST",
				"error_code": -2,
			})
			return
		}

		c.Set("USER", &user)

		c.Next()
	}
}
