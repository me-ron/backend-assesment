package middleware

import (
	"loan_tracker/domain"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RoleBasedAuth(protected bool, repo domain.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, _ := c.Get("userId")
		ID, _ := userId.(string)

		user, _ := repo.GetUserDocumentByID(ID)
		bools,_ := repo.GetBools(ID)
		admin := bools.IsAdmin

		if !admin {
			if protected {
				c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "you need to be an admin"})
				c.Abort()
				return
			}

			path := c.Request.URL.Path
			idx := c.Param("id")
        	if strings.Contains(path, "user") && idx != "" && idx != user.ID {
				c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				c.Abort()
				return
			}else if strings.Contains(path, "user") && idx == ""{
				c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				c.Abort()
				return
			}
		}

		c.Next()
	}

}