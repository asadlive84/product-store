package query

func (q *Query) CreateProductStock(ps ProductStock) (string, error) {
	const query = `
		INSERT INTO product_stocks (product_id, stock_quantity, updated_at)
		VALUES (:product_id, :stock_quantity, :updated_at)
		RETURNING id
	`

	stmt, err := q.DB.PrepareNamed(query)
	if err != nil {
		return "", err
	}

	var productStockID string
	err = stmt.Get(&productStockID, ps)

	if err != nil {
		return "", err
	}

	return productStockID, nil
}

func (q *Query) GetProductStockByID(productStockID int) (ProductStock, error) {
	const query = `
		SELECT * FROM product_stocks WHERE id = $1
	`

	var productStock ProductStock
	err := q.DB.Get(&productStock, query, productStockID)
	if err != nil {
		return ProductStock{}, err
	}

	return productStock, nil
}

func (q *Query) UpdateProductStock(ps ProductStock) error {
	const query = `
		UPDATE product_stocks
		SET product_id = :product_id, stock_quantity = :stock_quantity, updated_at = :updated_at
		WHERE id = :id
	`

	_, err := q.DB.NamedExec(query, ps)
	if err != nil {
		return err
	}

	return nil
}


func (q *Query) UpdateProductStockByPId(ps ProductStock) error {
	const query = `
		UPDATE product_stocks
		SET stock_quantity = :stock_quantity, updated_at = :updated_at
		WHERE product_id = :product_id
	`

	_, err := q.DB.NamedExec(query, ps)
	if err != nil {
		return err
	}

	return nil
}

func (q *Query) DeleteProductStock(productStockID int) error {
	const query = `
		DELETE FROM product_stocks WHERE id = $1
	`

	_, err := q.DB.Exec(query, productStockID)
	if err != nil {
		return err
	}

	return nil
}
