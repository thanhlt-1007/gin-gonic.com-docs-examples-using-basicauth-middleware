package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

var users = map[string]any {
    "foo": map[string]string {
        "email": "foo@example.com",
        "name": "FOO",
    },
}

func getAdminUserHandler(context *gin.Context) {
    user := context.MustGet(gin.AuthUserKey).(string)
    userData, ok := users[user]
    if ok {
        context.JSON(
            http.StatusOK,
            gin.H {
                "user": user,
                "data": userData,
            },
        )
    } else {
        context.JSON(
            http.StatusUnauthorized,
            gin.H {
                "error": "Unauthorized",
            },
        )
    }
}

func main() {
    engine := gin.Default()

    adminGroup := engine.Group("/admin", gin.BasicAuth(gin.Accounts{
        "foo": "bar",
    }))

    adminGroup.GET("/user", getAdminUserHandler)
    engine.Run()
}
