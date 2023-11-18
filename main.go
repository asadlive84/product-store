package main

import (
	"fmt"

	c "github.com/asadlive84/productstore/config"
	"github.com/asadlive84/productstore/db"
	"github.com/asadlive84/productstore/query"
	"github.com/asadlive84/productstore/view"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {
	config, err := c.LoadConfig()
	if err != nil {
		fmt.Printf("error load config %+v", err)
		return
	}

	dConfig := db.DBConfig{
		POSTGRES_USER:     config.POSTGRES_USER,
		POSTGRES_PASSWORD: config.POSTGRES_PASSWORD,
		POSTGRES_DB:       config.POSTGRES_DB,
		POSTGRES_PORT:     config.POSTGRES_PORT,
		POSTGRES_HOST:     config.POSTGRES_HOST,
		PORT:              config.PORT,
	}

	databaseConfig, err := db.NewDatabaseConfig(dConfig)

	if err != nil {
		fmt.Printf("error load config %+v", err)
		return
	}

	database, err := db.DbInit(databaseConfig)
	if err != nil {
		fmt.Printf("error load config %+v", err)
		return
	}

	db, err := sqlx.Connect("postgres", databaseConfig)
	if err != nil {
		fmt.Printf("error load config %+v", err)
		return
	}

	q := &query.Query{
		DB: db,
	}

	if err = query.InitializeDatabase(query.Query{
		DB: db,
	}); err != nil {
		fmt.Println("####################InitializeDatabase##################################")
		fmt.Printf("InitializeDatabase err: %+v", err)
		fmt.Println("######################################################")
		return

	}
	fmt.Println("#######################fake data seed success###############################")

	s := view.Server{
		SqlDB: database.SqlDB,
		Q:     q,
	}

	router := gin.Default()

	router.GET("/", s.Home)

	// Routes
	router.POST("/api/create/brand", s.CreateBrand)
	router.GET("/api/get/brand", s.GetBrandByID)
	router.POST("/api/update/brand", s.UpdateBrand)
	router.GET("/api/delete/brand", s.DeleteBrand)

	router.POST("/api/create/category", s.CreateCategory)
	router.GET("/api/get/category", s.GetCategoryByID)
	router.POST("/api/update/category", s.UpdateCategory)
	router.GET("/api/delete/category", s.DeleteCategory)

	router.POST("/api/create/supplier", s.CreateSupplier)
	router.GET("/api/get/supplier", s.GetSupplierByID)
	router.POST("/api/update/supplier", s.UpdateSupplier)
	router.GET("/api/delete/supplier", s.DeleteSupplier)

	router.POST("/api/create/product", s.CreateProduct)
	router.GET("/api/get/product", s.GetProductByID)
	router.POST("/api/update/product", s.UpdateProduct)
	router.GET("/api/delete/product", s.DeleteProduct)
	router.POST("/api/create/productstock", s.CreateProductStock)

	router.GET("/api/filter/products", s.GetProductFilter)
	router.GET("/api/tree/category", s.TreeCategories)

	// Run the server on port 8080
	router.Run(":" + config.PORT)
}
