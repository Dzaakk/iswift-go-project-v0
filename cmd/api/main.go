package main

import (
	admin "iswift-go-project/internal/admin/injector"
	oauth "iswift-go-project/internal/oauth/injector"
	profile "iswift-go-project/internal/profile/injector"
	register "iswift-go-project/internal/register/injector"
	mysql "iswift-go-project/pkg/db/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	db := mysql.DB()

	r := gin.Default()

	register.InitializedService(db).Route(&r.RouterGroup)
	oauth.InitializedService(db).Route(&r.RouterGroup)
	profile.InitializedService(db).Route(&r.RouterGroup)
	admin.InitializedService(db).Route(&r.RouterGroup)
	
	r.Run()

}
