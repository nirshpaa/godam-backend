package schema

import (
	"database/sql"
	"fmt"
)

// seeds is a string constant containing all the queries needed to get the
// db seeded to a useful state for development.

const clearUsers string = `DELETE FROM users`
const clearAccess string = `DELETE FROM access`
const clearRoles string = `DELETE FROM roles`
const clearAccessRoles string = `DELETE FROM access_roles`
const clearRolesUsers string = `DELETE FROM roles_users`
const clearCompanies string = `DELETE FROM companies`
const clearCategories string = `DELETE FROM categories`
const clearBranches string = `DELETE FROM branches`
const clearShelves string = `DELETE FROM shelves`
const clearSuppliers string = `DELETE FROM suppliers`
const clearBrands string = `DELETE FROM brands`
const clearProductCategories string = `DELETE FROM product_categories`
const clearProducts string = `DELETE FROM products`

const seedUsers string = `
INSERT INTO users (id, username, password, email, is_active, company_id, branch_id) VALUES 
(1, 'admin', '$2y$10$ekouPwVdtMEy5AFbogzfSeRloxHzUwEAsM7SyNJXnso/F9ds/XUYy', 'admin@admin.com', 1, 1, null),
(2, 'nishanpandit', '$2a$10$gT5pAqbiLxXTElwluvNJuef0jlOlHyt4q9ApC7jyhMb49OvIHeKgO', 'nishan@gmail.com', 1, 1, 1);
`

const seedAccess string = `INSERT INTO access (id, name, alias, created) VALUES (1, 'root', 'root', NOW())`

const seedRoles string = `INSERT INTO roles (id, name, company_id, created) VALUES (1, 'superadmin', 1, NOW())`

const seedAccessRoles string = `INSERT INTO access_roles (access_id, role_id) VALUES (1, 1)`

const seedRolesUsers string = `INSERT INTO roles_users (role_id, user_id) VALUES
(1, 1),
(1, 2)`

const seedCompanies string = `INSERT INTO companies (id, code, name) VALUES (1, "DM", "Dummy")`

const seedCategories string = `INSERT INTO categories (id, name) VALUES (1, "Accesories")`

const seedBranches string = `INSERT INTO branches (id, company_id, code, name, type, address ) VALUES (1, 1, "123", "Toko Bagus", "s", "jalan jalan")`

const seedShelves string = `INSERT INTO shelves (id, branch_id, code, capacity) VALUES (1, 1, "SHV-01", 1000)`

const seedSuppliers string = `INSERT INTO suppliers (id, company_id, code, name, address ) VALUES (1, 1, "SUP_01", "Supplier Test", "jalan supplier")`

const seedBrands string = `INSERT INTO brands (id, company_id, code, name) VALUES (1, 1, "BRAND-01", "Brand Test")`

const seedProductCategories string = `INSERT INTO product_categories (id, company_id, category_id, name) VALUES(1, 1, 1, "Furniture")`

const seedProducts string = `INSERT INTO products (id, company_id, brand_id, product_category_id, code, name, sale_price, minimum_stock, image_url, barcode_value, image_recognition_data) VALUES
(1, 1, 1, 1, "PROD-01", "Product Satu", "1000", "25", "", "", ""),
(2, 1, 1, 1, "PROD-02", "Product Dua", "500", "1000", "", "", "")`

// Seed runs the set of seed-data queries against db. The queries are run in a
// transaction and rolled back if any fail.
func Seed(db *sql.DB) error {
	// Clear previous data
	clearQueries := []string{
		// Start by clearing the child tables (those with foreign keys pointing to others)
		clearShelves,           // Delete shelves
		clearRolesUsers,        // Delete roles_users
		clearAccessRoles,       // Delete access_roles
		clearUsers,             // Delete users (dependent on roles_users)
		clearBranches,          // Delete branches first since it references companies
		clearProducts,          // Delete products
		clearProductCategories, // Delete product_categories
		clearBrands,            // Delete brands
		clearSuppliers,         // Delete suppliers
		clearCategories,        // Delete categories
		clearRoles,             // Delete roles
		clearCompanies,         // Finally delete companies
		clearAccess,            // Delete access
	}

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Clear previous data
	for _, clear := range clearQueries {
		_, err = tx.Exec(clear)
		if err != nil {
			tx.Rollback()
			fmt.Println("error clearing data:", err)
			return err
		}
	}

	// Insert seed data
	seeds := []string{
		seedCompanies,
		seedBranches,
		seedUsers,
		seedAccess,
		seedRoles,
		seedAccessRoles,
		seedRolesUsers,
		seedCategories,
		seedShelves,
		seedSuppliers,
		seedBrands,
		seedProductCategories,
		seedProducts,
	}

	for _, seed := range seeds {
		_, err = tx.Exec(seed)
		if err != nil {
			tx.Rollback()
			fmt.Println("error executing seed:", err)
			return err
		}
	}

	// Commit transaction
	return tx.Commit()
}
