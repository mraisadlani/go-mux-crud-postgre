package impl

import (
	"database/sql"
	"errors"
	"github.com/vanilla/go-mux-postgre/api/entities"
	"time"
)

type ProductRepositoryImpl struct {
	db *sql.DB
}

func NewProductRepositoryImpl(db *sql.DB)  *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db}
}

func (r *ProductRepositoryImpl) FindAll() ([]entities.Product, error) {
	row, err := r.db.Query(`
		SELECT * FROM products ORDER BY id ASC
	`)
	defer row.Close()

	if err != nil {
		return nil, err
	}

	var products []entities.Product
	for row.Next() {
		var product entities.Product

		err := row.Scan(
			&product.ID,
			&product.Name,
			&product.Slug,
			&product.Description,
			&product.Image,
			&product.Price,
			&product.Weight,
			&product.Status,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepositoryImpl) FindById(uid uint64) (entities.Product, error) {
	row, err := r.db.Prepare(`
		SELECT * FROM products WHERE id=$1
	`)
	defer row.Close()

	if err != nil {
		return entities.Product{}, err
	}

	var product entities.Product
	err = row.QueryRow(uid).Scan(
		&product.ID,
		&product.Name,
		&product.Slug,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Weight,
		&product.Status,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return entities.Product{}, err
	}

	if product.ID.Valid == false {
		return entities.Product{}, errors.New("Product not found")
	}

	return product, nil
}

func (r *ProductRepositoryImpl) Save(prod entities.Product) (bool, error) {
	query, err := r.db.Prepare(`
		SELECT id FROM products WHERE name=$1
	`)
	defer query.Close()

	if err != nil {
		return false, err
	}

	var product entities.Product
	err = query.QueryRow(prod.Name.String).Scan(&product.ID)

	if product.ID.Valid == true {
		return false, errors.New("Product already exists")
	}

	row, err := r.db.Prepare(`
		INSERT INTO products (name, slug, description, image, price, weight, status, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`)
	defer row.Close()

	if err != nil {
		return false, err
	}

	_, err = row.Exec(
		prod.Name.String,
		prod.Slug.String,
		prod.Description.String,
		prod.Image.String,
		prod.Price.Float64,
		prod.Weight.Float64,
		prod.Status.Bool,
		time.Now(),
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *ProductRepositoryImpl) Update(uid uint64, prod entities.Product) (bool, error) {
	query, err := r.db.Prepare(`
		SELECT id FROM products WHERE id=$1
	`)
	defer query.Close()

	if err != nil {
		return false, err
	}

	var product entities.Product
	err = query.QueryRow(uid).Scan(&product.ID)

	if product.ID.Valid == false {
		return false, errors.New("Product not found")
	}

	row, err := r.db.Prepare(`
		UPDATE products SET name=$1, slug=$2, description=$3, image=$4, price=$5, weight=$6, status=$7, updated_at=$8 WHERE id=$9
	`)
	defer row.Close()

	if err != nil {
		return false, err
	}

	_, err = row.Exec(
		prod.Name.String,
		prod.Slug.String,
		prod.Description.String,
		prod.Image.String,
		prod.Price.Float64,
		prod.Weight.Float64,
		prod.Status.Bool,
		time.Now(),
		uid,
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *ProductRepositoryImpl) Delete(uid uint64) (bool, error) {
	row, err := r.db.Prepare(`
		DELETE FROM products WHERE id=$1
	`)
	defer row.Close()

	if err != nil {
		return false, err
	}

	_, err = row.Exec(uid)

	if err != nil {
		return false, err
	}

	return true, nil
}