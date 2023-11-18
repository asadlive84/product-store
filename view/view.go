package view

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/asadlive84/productstore/query"
	"github.com/gin-gonic/gin"
)

type Server struct {
	SqlDB *sql.DB
	Q     query.DataQuery
	// View  View
}

// type Server struct {
// 	SqlDB *sql.DB
// 	Q     query.DataQuery
// }

// type View interface {
// 	CreateBrand(c *gin.Context)
// 	GetBrandByID(c *gin.Context)
// 	GetBrandList(c *gin.Context)
// 	UpdateBrand(c *gin.Context)
// 	DeleteBrand(c *gin.Context)

// 	CreateCategory(c *gin.Context)
// 	GetCategoryByID(c *gin.Context)
// 	GetCategoryList(c *gin.Context)
// 	UpdateCategory(c *gin.Context)
// 	DeleteCategory(c *gin.Context)

// 	CreateSupplier(c *gin.Context)
// 	GetSupplierByID(c *gin.Context)
// 	GetSupplierList(c *gin.Context)
// 	UpdateSupplier(c *gin.Context)
// 	DeleteSupplier(c *gin.Context)

// 	CreateProduct(c *gin.Context)
// 	GetProductByID(c *gin.Context)
// 	GetProductList(c *gin.Context)
// 	UpdateProduct(c *gin.Context)
// 	DeleteProduct(c *gin.Context)
// }

func (s *Server) Home(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Hello, welcome to the Gin API!"})
}

func (s *Server) CreateProduct(c *gin.Context) {

	var product query.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := product.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation error": err.Error()})
		return
	}

	pid, err := s.Q.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var pStockdID string
	if pid != "" {
		pId, err := strconv.Atoi(pid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if pId > 0 {
			pStockdID, err = s.Q.CreateProductStock(query.ProductStock{
				ProductID: pId,
				StockQty:  product.StockQty,
			})

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

	}

	c.JSON(http.StatusOK, gin.H{"message": "product created", "productID": pid, "Stock": pStockdID})
}

func (s *Server) GetProductByID(c *gin.Context) {

	var product query.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := s.Q.GetProductByID(product.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)

}

func (s *Server) DeleteProduct(c *gin.Context) {

	var product query.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.Q.DeleteProduct(product.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})

}

func (s *Server) UpdateProduct(c *gin.Context) {

	var product query.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.Q.UpdateProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = s.Q.UpdateProductStockByPId(query.ProductStock{
		ProductID: product.ID,
		StockQty:  product.StockQty,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product updated"})

}

func (s *Server) CreateBrand(c *gin.Context) {
	var brand query.Brand

	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := brand.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation error": err.Error()})
		return
	}

	if _, err := s.Q.CreateBrand(brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "brand created"})
}

func (s *Server) GetBrandByID(c *gin.Context) {

	var brand query.Brand

	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := s.Q.GetBrandByID(brand.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)

}

func (s *Server) DeleteBrand(c *gin.Context) {

	var brand query.Brand

	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.Q.DeleteBrand(brand.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "brand deleted"})

}

func (s *Server) UpdateBrand(c *gin.Context) {

	var brand query.Brand

	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.Q.UpdateBrand(brand)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "brand updated"})

}

func (s *Server) CreateCategory(c *gin.Context) {
	var category query.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := category.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation error": err.Error()})
		return
	}

	if _, err := s.Q.CreateCategory(category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "category created"})
}

func (s *Server) GetCategoryByID(c *gin.Context) {

	var category query.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := s.Q.GetCategoryByID(category.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)

}

func (s *Server) DeleteCategory(c *gin.Context) {

	var category query.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.Q.DeleteCategory(category.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "category deleted"})

}

func (s *Server) UpdateCategory(c *gin.Context) {

	var category query.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.Q.UpdateCategory(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "category updated"})

}

func (s *Server) CreateSupplier(c *gin.Context) {
	var supplier query.Supplier

	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := supplier.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation error": err.Error()})
		return
	}

	if _, err := s.Q.CreateSupplier(supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "supplier created"})
}

func (s *Server) GetSupplierByID(c *gin.Context) {

	var supplier query.Supplier

	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := s.Q.GetSupplierByID(supplier.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)

}

func (s *Server) DeleteSupplier(c *gin.Context) {

	var supplier query.Supplier

	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.Q.DeleteSupplier(supplier.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "supplier deleted"})

}

func (s *Server) UpdateSupplier(c *gin.Context) {

	var supplier query.Supplier

	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.Q.UpdateSupplier(supplier)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Supplier updated"})

}

func (s *Server) CreateProductStock(c *gin.Context) {
	var productStock query.ProductStock

	if err := c.ShouldBindJSON(&productStock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := productStock.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation error": err.Error()})
		return
	}

	if _, err := s.Q.CreateProductStock(productStock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product stock created"})
}

func (s *Server) GetProductStockByID(c *gin.Context) {

	var productStock query.ProductStock

	if err := c.ShouldBindJSON(&productStock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := s.Q.GetProductStockByID(productStock.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)

}

func (s *Server) DeleteProductStock(c *gin.Context) {

	var productStock query.ProductStock

	if err := c.ShouldBindJSON(&productStock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.Q.DeleteProductStock(productStock.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "productStock deleted"})

}

func (s *Server) UpdateProductStock(c *gin.Context) {

	var product query.ProductStock

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.Q.UpdateProductStock(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "productStock updated"})

}

func (s *Server) GetProductFilter(c *gin.Context) {
	var filter query.ProductListFilter

	if err := c.BindJSON(&filter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := filter.Validate(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	products, err := s.Q.GetProductList(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, query.ProductFilterResponse{
		Message: "Filter successfully",
		Count:   len(products),
		Data:    products,
	})
}

func (s *Server) TreeCategories(c *gin.Context) {

	treeCat, err := s.Q.TreeCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, t := range treeCat {
		s.Q.GetCategoryByID(t.ParentID)

	}

}
