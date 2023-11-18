package query

import (
	// "fmt"
	// "log"
	"strconv"
	"time"
	// "github.com/jmoiron/sqlx"
)

func InitializeDatabase(q Query) error {

	totalProduct, err := q.GetProductTotal()
	if err != nil {
		return err
	}

	if totalProduct >= 1 {
		return nil
	}

	// Insert brands
	brands := []Brand{
		{ID: 1, Name: "HP", StatusID: 1, CreatedAt: time.Now()},
		{ID: 2, Name: "Lenovo", StatusID: 1, CreatedAt: time.Now()},
		{ID: 3, Name: "Asus", StatusID: 1, CreatedAt: time.Now()},
		{ID: 2, Name: "Dell", StatusID: 1, CreatedAt: time.Now()},
		{ID: 2, Name: "Sony", StatusID: 1, CreatedAt: time.Now()},
		{ID: 2, Name: "Samsung", StatusID: 1, CreatedAt: time.Now()},
		// Add more brands as needed
	}

	for i, brand := range brands {
		brand.ID = i
		_, err := q.CreateBrand(brand)
		if err != nil {
			return err
		}
	}

	// Insert brands
	categorys := []Category{
		{ID: 1, Name: "Laptop", StatusID: 1, CreatedAt: time.Now()},
		{ID: 2, Name: "Notebook", StatusID: 1, CreatedAt: time.Now()},
		{ID: 2, Name: "Mobile", StatusID: 1, CreatedAt: time.Now()},
		// Add more brands as needed
	}

	for i, category := range categorys {
		category.ID = i
		_, err := q.CreateCategory(category)
		if err != nil {
			return err
		}
	}

	// Insert suppliers
	suppliers := []Supplier{
		{ID: 1, Name: "Ab Intetnationl", Email: "supplier1@example.com", Phone: "123-456-7890", StatusID: 1, IsVerifiedSupplier: true, CreatedAt: time.Now()},
		{ID: 2, Name: "Dhaka Supplier", Email: "supplier2@example.com", Phone: "987-654-3210", StatusID: 1, IsVerifiedSupplier: false, CreatedAt: time.Now()},
		{ID: 3, Name: "Khulna Supplier", Email: "supplierKhulna@example.com", Phone: "100-654-3210", StatusID: 1, IsVerifiedSupplier: false, CreatedAt: time.Now()},
		// Add more suppliers as needed
	}

	for i, supplier := range suppliers {
		supplier.ID = i
		_, err := q.CreateSupplier(supplier)
		if err != nil {
			return err
		}
	}

	// Insert products
	products := []Product{
		{ID: 1, Name: "Zbook 14ug6", Description: "Description1", BrandID: 1, CategoryID: 1, SupplierID: 1, UnitPrice: 100.0, DiscountPrice: 10.0, Tags: "Tag1", StatusID: 1, CreatedAt: time.Now(), StockQty: 29},
		{ID: 2, Name: "Zbook 15ug6", Description: "Description1", BrandID: 1, CategoryID: 1, SupplierID: 1, UnitPrice: 80.0, DiscountPrice: 10.0, Tags: "Tag1", StatusID: 1, CreatedAt: time.Now(), StockQty: 53},
		{ID: 3, Name: "Lenovo ThinkPad", Description: "Description1", BrandID: 1, CategoryID: 1, SupplierID: 2, UnitPrice: 200.0, DiscountPrice: 10.0, Tags: "Tag1", StatusID: 1, CreatedAt: time.Now(), StockQty: 67},
		{ID: 4, Name: "HP Envy", Description: "Description1", BrandID: 1, CategoryID: 1, SupplierID: 3, UnitPrice: 150.0, DiscountPrice: 15.0, Tags: "Tag1", StatusID: 1, CreatedAt: time.Now(), StockQty: 45},
		{ID: 5, Name: "Dell XPS 13", Description: "Description1", BrandID: 1, CategoryID: 1, SupplierID: 2, UnitPrice: 180.0, DiscountPrice: 12.0, Tags: "Tag1", StatusID: 1, CreatedAt: time.Now(), StockQty: 82},
		{ID: 6, Name: "MacBook Pro", Description: "Description1", BrandID: 1, CategoryID: 1, SupplierID: 1, UnitPrice: 220.0, DiscountPrice: 20.0, Tags: "Tag1", StatusID: 1, CreatedAt: time.Now(), StockQty: 60},
		{ID: 7, Name: "Surface Laptop 4", Description: "Description1", BrandID: 1, CategoryID: 1, SupplierID: 3, UnitPrice: 250.0, DiscountPrice: 18.0, Tags: "Tag1", StatusID: 1, CreatedAt: time.Now(), StockQty: 75},
		{ID: 8, Name: "Acer Predator", Description: "Description1", BrandID: 1, CategoryID: 1, SupplierID: 2, UnitPrice: 170.0, DiscountPrice: 22.0, Tags: "Tag1", StatusID: 1, CreatedAt: time.Now(), StockQty: 90},
		{ID: 9, Name: "Lenovo Yoga", Description: "Description1", BrandID: 1, CategoryID: 1, SupplierID: 1, UnitPrice: 190.0, DiscountPrice: 15.0, Tags: "Tag1", StatusID: 1, CreatedAt: time.Now(), StockQty: 50},
		{ID: 10, Name: "Asus ZenBook", Description: "Description1", BrandID: 1, CategoryID: 1, SupplierID: 3, UnitPrice: 160.0, DiscountPrice: 25.0, Tags: "Tag1", StatusID: 1, CreatedAt: time.Now(), StockQty: 110},
		{ID: 11, Name: "LG Gram", Description: "Description1", BrandID: 1, CategoryID: 1, SupplierID: 2, UnitPrice: 130.0, DiscountPrice: 10.0, Tags: "Tag1", StatusID: 1, CreatedAt: time.Now(), StockQty: 35},
		{ID: 12, Name: "Microsoft Surface Pro", Description: "Description1", BrandID: 1, CategoryID: 1, SupplierID: 1, UnitPrice: 200.0, DiscountPrice: 15.0, Tags: "Tag1", StatusID: 1, CreatedAt: time.Now(), StockQty: 65},
		{ID: 13, Name: "Razer Blade", Description: "Description1", BrandID: 1, CategoryID: 1, SupplierID: 3, UnitPrice: 300.0, DiscountPrice: 30.0, Tags: "Tag1", StatusID: 1, CreatedAt: time.Now(), StockQty: 40},
		{ID: 14, Name: "Samsung Galaxy Book", Description: "Description1", BrandID: 1, CategoryID: 1, SupplierID: 2, UnitPrice: 180.0, DiscountPrice: 18.0, Tags: "Tag1", StatusID: 1, CreatedAt: time.Now(), StockQty: 88},
		{ID: 15, Name: "HP Spectre x360", Description: "Description1", BrandID: 1, CategoryID: 1, SupplierID: 1, UnitPrice: 240.0, DiscountPrice: 20.0, Tags: "Tag1", StatusID: 1, CreatedAt: time.Now(), StockQty: 53},
		{ID: 16, Name: "Dell Inspiron", Description: "Description1", BrandID: 1, CategoryID: 1, SupplierID: 3, UnitPrice: 120.0, DiscountPrice: 10.0, Tags: "Tag1", StatusID: 1, CreatedAt: time.Now(), StockQty: 75},
		{ID: 17, Name: "Acer Aspire", Description: "Description1", BrandID: 1, CategoryID: 1, SupplierID: 2, UnitPrice: 100.0, DiscountPrice: 12.0, Tags: "Tag1", StatusID: 1, CreatedAt: time.Now(), StockQty: 45},
		{ID: 18, Name: "Lenovo Legion", Description: "Description1", BrandID: 1, CategoryID: 1, SupplierID: 1, UnitPrice: 220.0, DiscountPrice: 18.0, Tags: "Tag1", StatusID: 1, CreatedAt: time.Now(), StockQty: 90},
		{ID: 19, Name: "Asus ROG Strix", Description: "Description1", BrandID: 1, CategoryID: 1, SupplierID: 3, UnitPrice: 280.0, DiscountPrice: 25.0, Tags: "Tag1", StatusID: 1, CreatedAt: time.Now(), StockQty: 60},
		{ID: 20, Name: "Microsoft Surface Laptop", Description: "Description1", BrandID: 1, CategoryID: 1, SupplierID: 2, UnitPrice: 200.0, DiscountPrice: 15.0, Tags: "Tag1", StatusID: 1, CreatedAt: time.Now(), StockQty: 55},
	}

	for i, product := range products {
		product.ID = i
		pid, err := q.CreateProduct(product)
		if err != nil {
			return err
		}

		if pid != "" {
			pId, err := strconv.Atoi(pid)
			if err != nil {
				return err
			}
			if pId > 0 {
				_, err = q.CreateProductStock(ProductStock{
					ProductID: pId,
					StockQty:  product.StockQty,
				})

				if err != nil {
					return err
				}
			}

		}
	}

	return nil
}
