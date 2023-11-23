package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "mydb"
)

func SeedDatabase() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		host, port, user, password)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("----------<><>-----------", err)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbname))
	if err != nil {
		log.Fatal("----------", err)
	}

	sqlStatement := fmt.Sprintf("CREATE DATABASE %s", dbname)
	fmt.Println(sqlStatement)
	_, err = db.Exec(sqlStatement)
	if err != nil {
		log.Fatal("--->", err)
	}

	connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTables(db)

	fmt.Println("---------------------------------------------")
	seedData(db)
}

func createTables(db *sqlx.DB) {
	_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS brands (
				id SERIAL PRIMARY KEY,
				name VARCHAR(255),
				status_id INT,
				created_at BIGINT
			)
		`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS products (
				id SERIAL PRIMARY KEY,
				name VARCHAR(255),
				description TEXT,
				specification TEXT,
				brand_id INT,
				category_id INT,
				supplier_id INT,
				unit_price FLOAT,
				discount_price FLOAT,
				tags TEXT ARRAY,
				status_id INT,
				stock INT,
				created_at BIGINT
			)
		`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS suppliers (
				id SERIAL PRIMARY KEY,
				name VARCHAR(255),
				email VARCHAR(255),
				phone VARCHAR(20),
				status_id INT,
				is_verified_supplier BOOLEAN,
				created_at BIGINT
			)
		`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tables created successfully")
}

func seedData(db *sqlx.DB) {
	_, err := db.Exec(`
			INSERT INTO brands (name, status_id, created_at) VALUES
			('Brand1', 1, $1),
			('Brand2', 1, $1),
			('Brand3', 1, $1)
		`, time.Now().Unix())
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
			INSERT INTO products (name, description, specification, brand_id, category_id, supplier_id, unit_price, discount_price, tags, status_id, stock, created_at) VALUES
			('Product1', 'Description1', 'Spec1', 1, 1, 1, 100.0, 90.0, '{"tag1", "tag2"}', 1, 100, $1),
			('Product2', 'Description2', 'Spec2', 2, 1, 2, 150.0, 120.0, '{"tag3", "tag4"}', 1, 50, $1),
			('Product3', 'Description3', 'Spec3', 3, 2, 3, 200.0, 180.0, '{"tag5", "tag6"}', 1, 75, $1)
		`, time.Now().Unix())
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
			INSERT INTO suppliers (name, email, phone, status_id, is_verified_supplier, created_at) VALUES
			('Supplier1', 'supplier1@example.com', '1234567890', 1, true, $1),
			('Supplier2', 'supplier2@example.com', '9876543210', 1, true, $1),
			('Supplier3', 'supplier3@example.com', '5555555555', 1, true, $1)
		`, time.Now().Unix())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data seeded successfully")
}
