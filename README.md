# Ecommerce

This project is a Ecommerce.

## Installation

Make sure you have Go installed. Run the following commands to start the project:

```bash
go mod tidy
go run main.go


Ensure the following environment variables are set in your .env file

HOST=localhost
PORT=3000

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_NAME=mydb
DB_USER=postgres
DB_PASS=mysecretpassword




For starting the project simply need to install Go. Then USE go mod tidy to install all depenencies. 
Connect the Database with given credentials below

After connecting the DB start the project and write BOTH for doing seeding and starting the project. 
Or write START to simply start the project.

The project will start

API_ENDPOINTS

Create Brand -- Method POST
http://localhost:3000/api/brand
BODY: {
  "name": "Apex",
  "status_id": 1
}

Get Brand -- Method GET
http://localhost:3000/api/brand/460c3893-d368-4cdb-b407-cd267c32e36d

Update Brand -- Method PATCH
http://localhost:3000/api/brand/460c3893-d368-4cdb-b407-cd267c32e36d
BODY: {
  "name": "Bata",
  "status_id": 1
}

Get ALL Brand -- Method GET
http://localhost:3000/api/brand?page=1&limit=10

Delete Brand -- Method DELETE
http://localhost:3000/api/brand/460c3893-d368-4cdb-b407-cd267c32e36d



/////////Supplier////////

Create Supplier -- Method POST
http://localhost:3000/api/supplier
BODY: {
  "name": "Tonmoy Baroi",
  "email": "tonmoy@gmail.com",
  "phone": "01626758447",
  "status_id": 1,
  "is_verified_supplier": true
}

Get Supplier -- Method GET
http://localhost:3000/api/supplier/fef438e9-2c04-4e12-961d-d35e2d75e5cd

Get ALL Supplier -- Method GET
http://localhost:3000/api/supplier/all

Update Supplier -- Method PATCH
http://localhost:3000/api/supplier/fef438e9-2c04-4e12-961d-d35e2d75e5cd
BODY: {
  "name": "Tonmoy Baroi",
  "email": "tonmoy@gmail.com",
  "phone": "01626758447",
  "status_id": 1,
  "is_verified_supplier": true
}


Delete Supplier -- Method DELETE
http://localhost:3000/api/supplier/fef438e9-2c04-4e12-961d-d35e2d75e5cd


/////Product/////

Create Product -- Method POST
http://localhost:3000/api/product
Body: {
  "name": "anything",
  "description": "no description",
  "brand_id": "460c3893-d368-4cdb-b407-cd267c32e36d",
  "category_id": "28228a5c-a926-4709-86fb-6fba77e99d75",
  "supplier_id": "fef438e9-2c04-4e12-961d-d35e2d75e5cd",
  "unit_price": 170.5,
  "discount_price": 29.5,
  "tags": ["fresh", "latest"],
  "status_id": 1,
  "stock_quantity": 50
}


Get Product -- Method GET
http://localhost:5000/api/product/b10c8de7-f6a8-4fc9-a5da-61420f3096dc

Get ALL Products -- Method GET
http://localhost:5000/api/product/all


Update Product -- Method PATCH
http://localhost:5000/api/product/b10c8de7-f6a8-4fc9-a5da-61420f3096dc
Body: {
  "name": "everything",
  "description": "add some description",
  "brand_id": "460c3893-d368-4cdb-b407-cd267c32e36d",
  "category_id": "28228a5c-a926-4709-86fb-6fba77e99d75",
  "supplier_id": "fef438e9-2c04-4e12-961d-d35e2d75e5cd",
  "unit_price": 170.5,
  "discount_price": 29.5,
  "tags": ["fresh", "latest"],
  "status_id": 1,
  "stock_quantity": 50
}

Delete Product -- Method DELETE
http://localhost:5000/api/product/b10c8de7-f6a8-4fc9-a5da-61420f3096dc


