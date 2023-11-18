package query

import (
	"fmt"
	"strconv"
	"strings"
)

func (q *Query) CreateProduct(p Product) (string, error) {
	const query = `
		INSERT INTO products (name, description, brand_id, category_id, supplier_id, unit_price, discount_price, tags, status_id, created_at)
		VALUES (:name, :description, :brand_id, :category_id, :supplier_id, :unit_price, :discount_price, :tags, :status_id, :created_at)
		RETURNING id
	`

	stmt, err := q.DB.PrepareNamed(query)
	if err != nil {
		return "", err
	}

	var productID string
	err = stmt.Get(&productID, p)

	if err != nil {
		return "", err
	}

	return productID, nil
}

func (q *Query) GetProductByID(productID int) (Product, error) {
	const query = `
		SELECT * FROM products WHERE id = $1 AND status_id=1
	`

	var product Product
	err := q.DB.Get(&product, query, productID)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (q *Query) GetProductTotal() (int, error) {

	var productCount struct {
		Count int `db:"count"`
	}

	const query = `
		SELECT count(*) as count FROM products
	`

	err := q.DB.Get(&productCount.Count, query)
	if err != nil {
		return productCount.Count, err
	}

	return productCount.Count, nil
}

func (q *Query) UpdateProduct(p Product) error {
	const query = `
		UPDATE products
		SET name = :name, description = :description, brand_id = :brand_id, category_id = :category_id, 
		    supplier_id = :supplier_id, unit_price = :unit_price, discount_price = :discount_price,
			tags = :tags, status_id = :status_id
		WHERE id = :id
	`

	_, err := q.DB.NamedExec(query, p)
	if err != nil {
		return err
	}

	return nil
}

func (q *Query) DeleteProduct(productID int) error {
	const query = `
		DELETE FROM products WHERE id = $1
	`

	_, err := q.DB.Exec(query, productID)
	if err != nil {
		return err
	}

	return nil
}



func (q *Query) GetProductList(filter ProductListFilter) ([]Product, error) {

	query := `
		SELECT 
		products.id,
		products.name as product_name,
		products.unit_price,
		products.discount_price,
		products.description,
		brands.name as brand_name,
		categories.name as categories_name,
		suppliers.name as supplier_name ,
		suppliers.is_verified_supplier,
		products.category_id,
		ps.stock_quantity,
		products.status_id,
		products.tags,
		products.created_at
		
		FROM products
		left join product_stocks ps on ps.product_id= products.id
		left join brands on brands.id=products.brand_id
		left join categories on categories.id=products.category_id
		left join suppliers on suppliers.id=products.supplier_id
		WHERE products.status_id = 1
			AND ps.stock_quantity > 0
	`

	conditions := []string{}
	args := map[string]interface{}{}

	if filter.ProductName != "" {
		conditions = append(conditions, "products.name ILIKE :product_name")
		args["product_name"] = "%" + filter.ProductName + "%"
	}

	if filter.MinPrice > 0 {
		conditions = append(conditions, "unit_price >= :min_price")
		args["min_price"] = filter.MinPrice
	}

	if filter.MaxPrice > 0 {
		conditions = append(conditions, "unit_price <= :max_price")
		args["max_price"] = filter.MaxPrice
	}

	if len(filter.BrandIDs) > 0 {
		brandCondition := fmt.Sprintf("brand_id IN (%s)", intsToString(filter.BrandIDs))
		conditions = append(conditions, brandCondition)
	}

	if filter.CategoryID > 0 {
		conditions = append(conditions, "category_id = :category_id")
		args["category_id"] = filter.CategoryID
	}

	if filter.SupplierID > 0 {
		conditions = append(conditions, "supplier_id = :supplier_id")
		args["supplier_id"] = filter.SupplierID
	}

	if filter.IsVerifiedSupplier {
		conditions = append(conditions, "is_verified_supplier = :is_verified_supplier")
		args["is_verified_supplier"] = filter.IsVerifiedSupplier
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY " + filter.SortBy
	query += " LIMIT :page_size OFFSET :offset"

	offset := (filter.Page - 1) * filter.PageSize
	args["page_size"] = filter.PageSize
	args["offset"] = offset

	fmt.Println("######################################################")
	fmt.Printf("query: %+v\n", query)
	fmt.Printf("args: %+v\n", args)
	fmt.Println("######################################################")

	var products []Product

	stmt, err := q.DB.PrepareNamed(query)
	if err != nil {
		return nil, err
	}

	if err := stmt.Select(&products, args); err != nil {
		return nil, err
	}

	return products, nil
}

func intsToString(ints []int) string {
	strs := make([]string, len(ints))
	for i, v := range ints {
		strs[i] = strconv.Itoa(v)
	}
	return strings.Join(strs, ",")
}
