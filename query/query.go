package query

import (
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jmoiron/sqlx"
)

type Query struct {
	DB *sqlx.DB
}

type Brand struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	StatusID  int       `db:"status_id" json:"status_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func (b Brand) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Name, validation.Required, validation.Length(3, 100)),
		validation.Field(&b.StatusID, validation.Min(0), validation.Max(1)),
	)
}

type Category struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	ParentID  int       `db:"parent_id" json:"parent_id"`
	Sequence  int       `db:"sequence" json:"sequence"`
	StatusID  int       `db:"status_id" json:"status_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func (c Category) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required, validation.Length(3, 100)),
		validation.Field(&c.ParentID, validation.In(1)),
		validation.Field(&c.Sequence, validation.Required),
		validation.Field(&c.StatusID, validation.Min(0), validation.Max(1)),
	)
}

type Supplier struct {
	ID                 int       `db:"id" json:"id"`
	Name               string    `db:"name" json:"name"`
	Email              string    `db:"email" json:"email"`
	Phone              string    `db:"phone" json:"phone"`
	StatusID           int       `db:"status_id" json:"status_id"`
	IsVerifiedSupplier bool      `db:"is_verified_supplier" json:"is_verified_supplier"`
	CreatedAt          time.Time `db:"created_at" json:"created_at"`
}

func (s Supplier) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Name, validation.Required, validation.Length(3, 100)),
		validation.Field(&s.Email, validation.Length(3, 100), validation.Match(regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`))),
		validation.Field(&s.Phone, validation.Length(3, 100)),
		validation.Field(&s.StatusID, validation.Min(0), validation.Max(1)),
		validation.Field(&s.IsVerifiedSupplier, validation.In(true, false)),
	)
}

type ProductFilterResponse struct {
	Message string    `json:"Message"`
	Count   int       `json:"count"`
	Data    []Product `json:"data"`
}

type Product struct {
	ID                 int       `db:"id" json:"id"`
	Name               string    `db:"name" json:"name"`
	ProductName        string    `db:"product_name" json:"product_name"`
	Description        string    `db:"description" json:"description"`
	StockQty           int       `db:"stock_quantity" json:"stock_quantity"`
	BrandName          string    `db:"brand_name" json:"brand_name"`
	CategoriesName     string    `db:"categories_name" json:"categories_name"`
	SupplierName       string    `db:"supplier_name" json:"supplier_name"`
	BrandID            int       `db:"brand_id" json:"brand_id"`
	CategoryID         int       `db:"category_id" json:"category_id"`
	SupplierID         int       `db:"supplier_id" json:"supplier_id"`
	UnitPrice          float64   `db:"unit_price" json:"unit_price"`
	DiscountPrice      float64   `db:"discount_price" json:"discount_price"`
	IsVerifiedSupplier bool      `db:"is_verified_supplier" json:"is_verified_supplier"`
	Tags               string    `db:"tags" json:"tags"`
	StatusID           int       `db:"status_id" json:"status_id"`
	CreatedAt          time.Time `db:"created_at" json:"created_at"`
}

func (p Product) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Length(3, 100)),
		validation.Field(&p.Description, validation.Required, validation.Length(3, 500)),
		validation.Field(&p.StockQty, validation.Required, validation.Min(1)),
		validation.Field(&p.CategoryID, validation.Required, validation.Min(1)),
		validation.Field(&p.BrandID, validation.Required, validation.Min(1)),
		validation.Field(&p.SupplierID, validation.Required, validation.Min(1)),
		validation.Field(&p.UnitPrice, validation.Required),
	)
}

type ProductStock struct {
	ID        int       `db:"id" json:"id"`
	ProductID int       `db:"product_id" json:"product_id"`
	StockQty  int       `db:"stock_quantity" json:"stock_quantity"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (c ProductStock) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.ProductID, validation.Required),
		validation.Field(&c.StockQty, validation.Required),
	)
}

type ProductListFilter struct {
	ProductName        string  `json:"product_name"`
	MinPrice           float64 `json:"min_price"`
	MaxPrice           float64 `json:"max_price"`
	BrandIDs           []int   `json:"brand_ids"`
	CategoryID         int     `json:"category_id"`
	SupplierID         int     `json:"supplier_id"`
	IsVerifiedSupplier bool    `json:"is_verified_supplier"`
	Page               int     `json:"page"`
	PageSize           int     `json:"page_size"`
	SortBy             string  `json:"sort_by"`
}

func (p ProductListFilter) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.ProductName),
		validation.Field(&p.MinPrice),
		validation.Field(&p.MaxPrice),
		validation.Field(&p.BrandIDs),
		validation.Field(&p.CategoryID),
		validation.Field(&p.SupplierID),
		validation.Field(&p.IsVerifiedSupplier, validation.In(true, false)),
		validation.Field(&p.Page, validation.Required),
		validation.Field(&p.PageSize, validation.Required),
		validation.Field(&p.SortBy),
	)
}

type TreeCategory struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	ParentID  int        `json:"parent_id"`
	Sequence  int        `json:"sequence"`
	StatusID  int        `json:"status_id"`
	CreatedAt string     `json:"created_at"`
	Children  []Category `json:"children,omitempty"`
}

type DataQuery interface {
	// Brand CRUD
	CreateBrand(Brand) (string, error)
	GetBrandByID(brandID int) (Brand, error)
	UpdateBrand(Brand) error
	DeleteBrand(brandID int) error

	// Category CRUD
	CreateCategory(Category) (string, error)
	GetCategoryByID(categoryID int) (Category, error)
	UpdateCategory(Category) error
	TreeCategories() ([]TreeCategory, error)
	DeleteCategory(categoryID int) error

	// Supplier CRUD
	CreateSupplier(Supplier) (string, error)
	GetSupplierByID(supplierID int) (Supplier, error)
	UpdateSupplier(Supplier) error
	DeleteSupplier(supplierID int) error

	// Product CRUD
	CreateProduct(Product) (string, error)
	GetProductByID(productID int) (Product, error)
	UpdateProduct(Product) error
	DeleteProduct(productID int) error

	GetProductList(filter ProductListFilter) ([]Product, error)

	// ProductStock CRUD
	CreateProductStock(ProductStock) (string, error)
	GetProductStockByID(productStockID int) (ProductStock, error)
	UpdateProductStock(ProductStock) error
	UpdateProductStockByPId(ProductStock) error
	DeleteProductStock(productStockID int) error
}
