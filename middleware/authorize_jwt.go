package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"jwt-request-demo/util"
	"log"
	"net/http"
)

/**
 * @Author wyf
 * @Date 2021/7/8 11:06
 **/

const BearerSchema = "Bearer "

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
				"code": http.StatusUnauthorized,
				"msg":  "No Authorization: header found",
			})
			return
		}

		tokenString := authHeader[len(BearerSchema):]

		if token, err := util.ValidateToken(tokenString); err != nil {
			log.Println("token", token)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
				"code": http.StatusUnauthorized,
				"msg":  "No Valid Token",
			})
		} else {
			if claims, ok := token.Claims.(jwt.MapClaims); !ok {
				ctx.AbortWithStatus(http.StatusUnauthorized)
			} else {
				if token.Valid {
					ctx.Set("username", claims["username"])
					ctx.Set("password", claims["password"])
				} else {
					ctx.AbortWithStatus(http.StatusUnauthorized)
				}
			}
		}
	}
}
