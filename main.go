package main

import (
	"kstyle-test/handler"
	"kstyle-test/repository"
	"kstyle-test/service"
	"kstyle-test/utils"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := utils.ConnectDB()
	if err != nil {
		panic(err)
	}
	memberRepo := repository.NewMemberRepository(db)
	memberService := service.NewMemberService(memberRepo)
	memberHandler := handler.NewMemberHandler(memberService)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	reviewRepo := repository.NewReviewRepository(db)
	reviewService := service.NewReviewService(reviewRepo, memberRepo, productRepo)
	reviewHandler := handler.NewReviewHandler(reviewService)

	likeRepo := repository.NewLikeRepository(db)
	likeService := service.NewLikeService(likeRepo, memberRepo, reviewRepo)
	likeHandler := handler.NewLikeHandler(likeService)

	// Routes
	e.POST("/v1/members", memberHandler.Create)
	e.GET("/v1/members", memberHandler.GetAll)
	e.GET("/v1/members/:id", memberHandler.GetByID)
	e.PUT("/v1/members/:id", memberHandler.Update)
	e.DELETE("/v1/members/:id", memberHandler.Delete)
	e.POST("/v1/likes", likeHandler.Like)
	e.POST("/v1/reviews", reviewHandler.Create)
	e.GET("/v1/products", productHandler.GetAll)
	e.GET("/v1/products/:id", productHandler.GetByID)
	e.POST("/v1/products", productHandler.Create)
	e.PUT("/v1/products/:id", productHandler.Update)
	e.DELETE("/v1/products/:id", productHandler.Delete)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
