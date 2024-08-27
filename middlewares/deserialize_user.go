package middlewares

import (
	"api-curut-in/structs"
	"encoding/json"
	"net/http"
	"strings"

	"api-curut-in/config"
	"api-curut-in/repository"
	"api-curut-in/utils"

	"github.com/gin-gonic/gin"
)

func DeserializeUser(userRepository repository.UsersRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && (fields[0] == "Bearer") {
			token = fields[1]
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, structs.ResponseModel{
				Response:   http.StatusUnauthorized,
				Error:      "Unauthorized",
				AppID:      "api.curut.id",
				Controller: "-",
				Action:     "-",
				Result:     nil,
			})
			return
		}

		config, _ := config.LoadConfig(".")
		sub, err := utils.ValidateToken(token, config.SecretKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, structs.ResponseModel{
				Response:   http.StatusUnauthorized,
				Error:      "Unauthorized",
				AppID:      "api.curut.id",
				Controller: "-",
				Action:     "-",
				Result:     nil,
			})
			return
		}

		id := sub.(string)
		result, err := userRepository.FindById(id)

		//got role from user
		idRole, _ := json.Marshal(result.RolesID)
		role := string(idRole)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, structs.ResponseModel{
				Response:   http.StatusForbidden,
				Error:      "You are not authorized to access this route",
				AppID:      "api.curut.id",
				Controller: "-",
				Action:     "-",
				Result:     nil,
			})
			return
		}

		ctx.Set("currentUser", result.Username)
		ctx.Set("userId", id)
		ctx.Set("role", role)
		ctx.Next()

	}
}

func AuthorizationAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, _ := ctx.Get("role")
		// jika role bukan 1 maka akan di abort
		if role != "1" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, structs.ResponseModel{
				Response:   http.StatusForbidden,
				Error:      "You are not authorized to access this route",
				AppID:      "api.curut.id",
				Controller: "-",
				Action:     "-",
				Result:     nil,
			})
			return
		}
		// log.Println(role)
		ctx.Next()
	}
}
