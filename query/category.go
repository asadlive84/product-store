package query

func (q *Query) CreateCategory(c Category) (string, error) {
	const query = `
		INSERT INTO categories (name, parent_id, sequence, status_id, created_at)
		VALUES (:name, :parent_id, :sequence, :status_id, :created_at)
		RETURNING id
	`

	stmt, err := q.DB.PrepareNamed(query)
	if err != nil {
		return "", err
	}

	var categoryID string
	err = stmt.Get(&categoryID, c)

	if err != nil {
		return "", err
	}

	return categoryID, nil
}

func (q *Query) GetCategoryByID(categoryID int) (Category, error) {
	const query = `
		SELECT * FROM categories WHERE id = $1 AND status_id=1
	`

	var category Category
	err := q.DB.Get(&category, query, categoryID)
	if err != nil {
		return Category{}, err
	}

	return category, nil
}

func (q *Query) UpdateCategory(c Category) error {
	const query = `
		UPDATE categories
		SET name = :name, parent_id = :parent_id, sequence = :sequence, status_id = :status_id
		WHERE id = :id
	`

	_, err := q.DB.NamedExec(query, c)
	if err != nil {
		return err
	}

	return nil
}

func (q *Query) DeleteCategory(categoryID int) error {
	const query = `
		DELETE FROM categories WHERE id = $1
	`

	_, err := q.DB.Exec(query, categoryID)
	if err != nil {
		return err
	}

	return nil
}

func (q *Query) TreeCategories() ([]TreeCategory, error) {
	const query = `
		SELECT * FROM categories WHERE status_id=1
	`

	var category []TreeCategory
	err := q.DB.Select(&category, query)
	if err != nil {
		return category, err
	}

	return category, nil
}
