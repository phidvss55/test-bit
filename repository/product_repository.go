package repository

import (
	"context"
	"database/sql"

	"go-circleci/types"
)

// ProductRepository defines the interface for product data access operations
type ProductRepository interface {
	GetAll(ctx context.Context) ([]*types.Product, error)
	GetByID(ctx context.Context, id int) (*types.Product, error)
	Create(ctx context.Context, product *types.Product) error
	Update(ctx context.Context, product *types.Product) error
	Delete(ctx context.Context, id int) error
}

// SQLiteProductRepository implements ProductRepository using SQLite
type SQLiteProductRepository struct {
	db *sql.DB
}

// NewSQLiteProductRepository creates a new SQLite product repository
func NewSQLiteProductRepository(db *sql.DB) *SQLiteProductRepository {
	return &SQLiteProductRepository{db: db}
}

// GetAll retrieves all products from the database
func (r *SQLiteProductRepository) GetAll(ctx context.Context) ([]*types.Product, error) {
	query := `SELECT id, name, description, price, stock FROM products`
	
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*types.Product
	for rows.Next() {
		product := &types.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// GetByID retrieves a single product by its ID
func (r *SQLiteProductRepository) GetByID(ctx context.Context, id int) (*types.Product, error) {
	query := `SELECT id, name, description, price, stock FROM products WHERE id = ?`
	
	product := &types.Product{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
	)
	
	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}
	if err != nil {
		return nil, err
	}

	return product, nil
}

// Create inserts a new product into the database and sets its generated ID
func (r *SQLiteProductRepository) Create(ctx context.Context, product *types.Product) error {
	query := `INSERT INTO products (name, description, price, stock) VALUES (?, ?, ?, ?)`
	
	result, err := r.db.ExecContext(ctx, query, product.Name, product.Description, product.Price, product.Stock)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	product.ID = int(id)
	return nil
}

// Update modifies an existing product in the database
func (r *SQLiteProductRepository) Update(ctx context.Context, product *types.Product) error {
	query := `UPDATE products SET name = ?, description = ?, price = ?, stock = ? WHERE id = ?`
	
	result, err := r.db.ExecContext(ctx, query, product.Name, product.Description, product.Price, product.Stock, product.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Delete removes a product from the database by its ID
func (r *SQLiteProductRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM products WHERE id = ?`
	
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
