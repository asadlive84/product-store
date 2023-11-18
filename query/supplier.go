package query

func (q *Query) CreateSupplier(s Supplier) (string, error) {
	const query = `
		INSERT INTO suppliers (name, email, phone, status_id, is_verified_supplier, created_at)
		VALUES (:name, :email, :phone, :status_id, :is_verified_supplier, :created_at)
		RETURNING id
	`

	stmt, err := q.DB.PrepareNamed(query)
	if err != nil {
		return "", err
	}

	var supplierID string
	err = stmt.Get(&supplierID, s)

	if err != nil {
		return "", err
	}

	return supplierID, nil
}

func (q *Query) GetSupplierByID(supplierID int) (Supplier, error) {
	const query = `
		SELECT * FROM suppliers WHERE id = $1 AND status_id=1
	`

	var supplier Supplier
	err := q.DB.Get(&supplier, query, supplierID)
	if err != nil {
		return Supplier{}, err
	}

	return supplier, nil
}

func (q *Query) UpdateSupplier(s Supplier) error {
	const query = `
		UPDATE suppliers
		SET name = :name, email = :email, phone = :phone, status_id = :status_id,
			is_verified_supplier = :is_verified_supplier
		WHERE id = :id
	`

	_, err := q.DB.NamedExec(query, s)
	if err != nil {
		return err
	}

	return nil
}

func (q *Query) DeleteSupplier(supplierID int) error {
	const query = `
		DELETE FROM suppliers WHERE id = $1
	`

	_, err := q.DB.Exec(query, supplierID)
	if err != nil {
		return err
	}

	return nil
}
