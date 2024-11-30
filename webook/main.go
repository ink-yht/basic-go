package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/ink-yht/basic-go/webook/internal/repository"
	"github.com/ink-yht/basic-go/webook/internal/repository/dao"
	"github.com/ink-yht/basic-go/webook/internal/service"
	"github.com/ink-yht/basic-go/webook/internal/web"
	"github.com/ink-yht/basic-go/webook/internal/web/middlewares"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

func main() {
	//db := initDB()
	//server := initWebService()
	//u := initUser(db)
	//u.RegisterRouter(server)
	server := gin.Default()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})
	server.Run(":8080")
}

func initWebService() *gin.Engine {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"x-jwt-token"},
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "your_company.com")
		},
		MaxAge: 12 * time.Hour,
	}))

	//store := cookie.NewStore([]byte("secret"))

	store, err := redis.NewStore(16, "tcp", "localhost:6379", "", []byte("wFLynr8VV9MZKfi2InewAGEsBcA8lgFC"), []byte("kcCaPDlxAs6ehUJvW0sUhgityvdm3hHA"))
	if err != nil {
		panic(err)
	}

	server.Use(sessions.Sessions("mysession", store))

	server.Use(middlewares.NewLoginJWTMiddlewareBuilder().IgnorePaths("/users/signup").IgnorePaths("/users/login").Build())

	return server
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDao(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		panic(err)
	}

	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}
