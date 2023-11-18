package query

func (q *Query) CreateBrand(b Brand) (string, error) {
	const query = `
		INSERT INTO brands (name, status_id, created_at)
		VALUES (:name, :status_id, :created_at)
		RETURNING id
	`

	stmt, err := q.DB.PrepareNamed(query)
	if err != nil {
		return "", err
	}

	var brandID string
	err = stmt.Get(&brandID, b)

	if err != nil {
		return "", err
	}

	return brandID, nil
}

func (q *Query) GetBrandByID(brandID int) (Brand, error) {
	const query = `
		SELECT * FROM brands WHERE id = $1 AND status_id=1
	`

	var brand Brand
	err := q.DB.Get(&brand, query, brandID)
	if err != nil {
		return Brand{}, err
	}

	return brand, nil
}

func (q *Query) UpdateBrand(b Brand) error {
	const query = `
		UPDATE brands
		SET name = :name, status_id = :status_id
		WHERE id = :id
	`

	_, err := q.DB.NamedExec(query, b)
	if err != nil {
		return err
	}

	return nil
}

func (q *Query) DeleteBrand(brandID int) error {
	const query = `
		DELETE FROM brands WHERE id = $1
	`

	_, err := q.DB.Exec(query, brandID)
	if err != nil {
		return err
	}

	return nil
}
