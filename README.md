# ProductStore project


### how to run this project?

```
First you would go to you postgres database and set up your database, 
must be postgres installed in your computer 
``````

```golang 1.19+, postgres```

```I will  be added docker asap```

### 
Clone this project
Now please setup with you postgres and this application
please go to this project and open this folder
###
```
config/env/dev.env
```
open this ```dev.env``` file and set up this config with postgres database like username, dbname etc

```
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=product
POSTGRES_PORT=5432
POSTGRES_HOST=localhost
PORT=8080
```

Now run this command

```
go mod tidy
```
then 

```
go mod vendor

```

then run this project
```
go run main.go

```

Default port is ```8080```


You have to clone this project



### Create Product
### Endpoint: 
 ```
/api/create/product
```
### HTTP Method: POST
### URL:
 
```
 http://localhost:8080/api/create/product
```
###  Request Body:
```
{
  "name": "lenovo thinkpad x280l",
  "description": "Powerful laptop with high-performance specifications",
  "specifications": "Processor: Intel Core i7, RAM: 16GB, Storage: 512GB SSD",
  "brand_id": 1,  // Replace with the actual brand ID
  "category_id": 1,  // Replace with the actual category ID
  "supplier_id": 1,  // Replace with the actual supplier ID
  "unit_price": 1200.00,
  "discount_price": 100.00,
  "tags": "electronics, laptop, high-performance",
  "status_id": 1,  // Active status
  "stock_quantity":19 // for product stock
}

```

### Response:

###  Status Code: 
``` 
200 OK
 ```

###  Response Body:
```
{
    "message": "product created"
}
```

### GET Product

###  Endpoint: 

```/api/get/product
```

###  HTTP Method: GET

###  URL: 
```
http://localhost:8080/api/get/product
```
Request Body:
```
{
    "id":1
}
```

###  Response:

###  Status Code: 200 OK

###  Response Body:
```
{
    "id": 1,
    "name": "Zbook 14ug6",
    "description": "Description1",
    "stock_quantity": 0,
    "brand_name": "",
    "categories_name": "",
    "supplier_name": "",
    "brand_id": 1,
    "category_id": 1,
    "supplier_id": 1,
    "unit_price": 100,
    "discount_price": 10,
    "is_verified_supplier": false,
    "tags": "Tag1",
    "status_id": 1,
    "created_at": "2023-11-17T18:29:05.730715Z"
}
```

###  UPDATE Product

###  Endpoint: /api/update/product

###  HTTP Method: POST

###  URL: ### 
```

http://localhost:8080/api/update/product
```

```
Request Body:
```

```
{
    "id": 1,
    "name": "Zbook 14ug6 update",
    "description": "Description1 update",
    "brand_id": 1,
    "category_id": 1,
    "supplier_id": 1,
    "tags": "Tag1",
    "status_id": 1,
    "stock_quantity":10000
}
```

###  Response:

###  Status Code: 200 OK

###  Response Body:
```
{
    "message": "product updated"
}
```

###  DELETE Product

###  Endpoint: ###

```
/api/delete/product
```

###  HTTP Method: GET

###  URL:
###  
```
 http://localhost:8080/api/delete/product
```


###  Request Body:
```
{
    "id":1,
    
}
```

###  Response:

###  Status Code: 200 OK

###  Response Body:
```
{
    "message": "product deleted"
}
```


###  Filter product

###  Endpoint: 
### 
```
 /api/filter/product
```



###  HTTP Method: GET

###  URL: 
###  

```
localhost:8080/api/filter/products
```
Request Body:
```
{
   "product_name":"zbook",
  "min_price":0,
  "max_price":10,
  "brand_ids":[1,2],
  "category_id":1,
  "supplier_id":1,
  "is_verified_supplier":false,
  "page": 1,
  "page_size":12,
  "sort_by": "unit_price ASC"
}
```


###  Response:

###  Status Code: 200 OK

###  Response Body:

```
{
    "Message": "Filter successfully",
    "count": 12,
    "data": [
        {
            "id": 1,
            "name": "",
            "product_name": "Zbook 14ug6 update",
            "description": "Description1 update",
            "stock_quantity": 10000,
            "brand_name": "hapnew",
            "categories_name": "laptop update NEW",
            "supplier_name": "",
            "brand_id": 0,
            "category_id": 1,
            "supplier_id": 0,
            "unit_price": 0,
            "discount_price": 0,
            "is_verified_supplier": false,
            "tags": "Tag1",
            "status_id": 1,
            "created_at": "2023-11-17T18:29:05.730715Z"
        },
        {
            "id": 2,
            "name": "",
            "product_name": "Zbook 15ug6",
            "description": "Description1",
            "stock_quantity": 53,
            "brand_name": "hapnew",
            "categories_name": "laptop update NEW",
            "supplier_name": "",
            "brand_id": 0,
            "category_id": 1,
            "supplier_id": 0,
            "unit_price": 80,
            "discount_price": 10,
            "is_verified_supplier": false,
            "tags": "Tag1",
            "status_id": 1,
            "created_at": "2023-11-17T18:29:05.730715Z"
        },
        {
            "id": 17,
            "name": "",
            "product_name": "Acer Aspire",
            "description": "Description1",
            "stock_quantity": 45,
            "brand_name": "hapnew",
            "categories_name": "laptop update NEW",
            "supplier_name": "Dhaka Supplier",
            "brand_id": 0,
            "category_id": 1,
            "supplier_id": 0,
            "unit_price": 100,
            "discount_price": 12,
            "is_verified_supplier": false,
            "tags": "Tag1",
            "status_id": 1,
            "created_at": "2023-11-17T18:29:05.730716Z"
        },
        {
            "id": 16,
            "name": "",
            "product_name": "Dell Inspiron",
            "description": "Description1",
            "stock_quantity": 75,
            "brand_name": "hapnew",
            "categories_name": "laptop update NEW",
            "supplier_name": "Khulna Supplier",
            "brand_id": 0,
            "category_id": 1,
            "supplier_id": 0,
            "unit_price": 120,
            "discount_price": 10,
            "is_verified_supplier": false,
            "tags": "Tag1",
            "status_id": 1,
            "created_at": "2023-11-17T18:29:05.730716Z"
        },
        {
            "id": 11,
            "name": "",
            "product_name": "LG Gram",
            "description": "Description1",
            "stock_quantity": 35,
            "brand_name": "hapnew",
            "categories_name": "laptop update NEW",
            "supplier_name": "Dhaka Supplier",
            "brand_id": 0,
            "category_id": 1,
            "supplier_id": 0,
            "unit_price": 130,
            "discount_price": 10,
            "is_verified_supplier": false,
            "tags": "Tag1",
            "status_id": 1,
            "created_at": "2023-11-17T18:29:05.730716Z"
        },
        {
            "id": 4,
            "name": "",
            "product_name": "HP Envy",
            "description": "Description1",
            "stock_quantity": 45,
            "brand_name": "hapnew",
            "categories_name": "laptop update NEW",
            "supplier_name": "Khulna Supplier",
            "brand_id": 0,
            "category_id": 1,
            "supplier_id": 0,
            "unit_price": 150,
            "discount_price": 15,
            "is_verified_supplier": false,
            "tags": "Tag1",
            "status_id": 1,
            "created_at": "2023-11-17T18:29:05.730715Z"
        },
        {
            "id": 10,
            "name": "",
            "product_name": "Asus ZenBook",
            "description": "Description1",
            "stock_quantity": 110,
            "brand_name": "hapnew",
            "categories_name": "laptop update NEW",
            "supplier_name": "Khulna Supplier",
            "brand_id": 0,
            "category_id": 1,
            "supplier_id": 0,
            "unit_price": 160,
            "discount_price": 25,
            "is_verified_supplier": false,
            "tags": "Tag1",
            "status_id": 1,
            "created_at": "2023-11-17T18:29:05.730716Z"
        },
        {
            "id": 8,
            "name": "",
            "product_name": "Acer Predator",
            "description": "Description1",
            "stock_quantity": 90,
            "brand_name": "hapnew",
            "categories_name": "laptop update NEW",
            "supplier_name": "Dhaka Supplier",
            "brand_id": 0,
            "category_id": 1,
            "supplier_id": 0,
            "unit_price": 170,
            "discount_price": 22,
            "is_verified_supplier": false,
            "tags": "Tag1",
            "status_id": 1,
            "created_at": "2023-11-17T18:29:05.730715Z"
        },
        {
            "id": 14,
            "name": "",
            "product_name": "Samsung Galaxy Book",
            "description": "Description1",
            "stock_quantity": 88,
            "brand_name": "hapnew",
            "categories_name": "laptop update NEW",
            "supplier_name": "Dhaka Supplier",
            "brand_id": 0,
            "category_id": 1,
            "supplier_id": 0,
            "unit_price": 180,
            "discount_price": 18,
            "is_verified_supplier": false,
            "tags": "Tag1",
            "status_id": 1,
            "created_at": "2023-11-17T18:29:05.730716Z"
        },
        {
            "id": 5,
            "name": "",
            "product_name": "Dell XPS 13",
            "description": "Description1",
            "stock_quantity": 82,
            "brand_name": "hapnew",
            "categories_name": "laptop update NEW",
            "supplier_name": "Dhaka Supplier",
            "brand_id": 0,
            "category_id": 1,
            "supplier_id": 0,
            "unit_price": 180,
            "discount_price": 12,
            "is_verified_supplier": false,
            "tags": "Tag1",
            "status_id": 1,
            "created_at": "2023-11-17T18:29:05.730715Z"
        },
        {
            "id": 9,
            "name": "",
            "product_name": "Lenovo Yoga",
            "description": "Description1",
            "stock_quantity": 50,
            "brand_name": "hapnew",
            "categories_name": "laptop update NEW",
            "supplier_name": "",
            "brand_id": 0,
            "category_id": 1,
            "supplier_id": 0,
            "unit_price": 190,
            "discount_price": 15,
            "is_verified_supplier": false,
            "tags": "Tag1",
            "status_id": 1,
            "created_at": "2023-11-17T18:29:05.730716Z"
        },
        {
            "id": 3,
            "name": "",
            "product_name": "Lenovo ThinkPad",
            "description": "Description1",
            "stock_quantity": 67,
            "brand_name": "hapnew",
            "categories_name": "laptop update NEW",
            "supplier_name": "Dhaka Supplier",
            "brand_id": 0,
            "category_id": 1,
            "supplier_id": 0,
            "unit_price": 200,
            "discount_price": 10,
            "is_verified_supplier": false,
            "tags": "Tag1",
            "status_id": 1,
            "created_at": "2023-11-17T18:29:05.730715Z"
        }
    ]
}

```






###  create Product

###  Endpoint:
### 
 /api/create/brand

###  HTTP Method: POST

###  URL: 
```
http://localhost:8080/api/create/brand
```
Request Body:
```
{
    "name":"symphony",
    "status_id":1
}
```

###  Response:

###  Status Code: 200 OK

###  Response Body:
```
{
    "message": "brand created"
}

```

###  Get brand

###  Endpoint: /api/get/brand

###  HTTP Method: GET

###  URL: 
``` 
http://localhost:8080/api/get/brand
```
Request Body:
```
{
    "id":1
}
```

###  Response:

###  Status Code: 200 OK

###  Response Body:

```
{
    "id": 1,
    "name": "HP",
    "status_id": 1,
    "created_at": "2023-11-17T18:29:05.71639Z"
}
```

###  update brand

###  Endpoint: /api/update/brand
 
###  HTTP Method: POST

###  URL: 
### 
```
 http://localhost:8080/api/update/brand
```


### Request Body:
```
{
    "id":1,
    "name":"HP NEW"
}
```

###  Response:

###  Status Code: 200 OK

###  Response Body:

```
{
    "message": "brand updated"
}
```

###  delete

###  Endpoint: /api/delete/brand
 
### HTTP Method: GET
 
### URL: 
```
http://localhost:8080/api/delete/brand
```

###  Request Body:

```
{
    "id":1,
    
}
```


###  Response:

###  Status Code: 200 OK
 
###  Response Body:
```
{
    "message": "brand deleted"
}
```

###  create category
 
###  how to create category

###  Endpoint: /api/create/category

###  HTTP Method: POST

###  URL: 
```
http://localhost:8080/api/create/category
```

###  Request Body:

```
{
    "name":"symphony",
    "status_id":1,
    "sequence":1
}
```

###  Response:

###  Status Code: 200 OK

###  Response Body:
```
{
    "message": "category created"
}
```

###  Get category

###  Endpoint: /api/get/category

###  HTTP Method: GET

###  URL:
```
 http://localhost:8080/api/get/category
```

###  Request Body:
```
{
    "id":1
}
```

###  Response:

###  Status Code: 200 OK

###  Response Body:
```
{
    "id": 1,
    "name": "Laptop",
    "parent_id": 0,
    "sequence": 0,
    "status_id": 1,
    "created_at": "2023-11-17T18:29:05.723644Z"
}
```


###  update category

###  Endpoint: /api/update/category

###  HTTP Method: POST

###  URL: 
```
http://localhost:8080/api/update/category
```

###  Request Body:
```
{
    "id":1,
    "name":"laptop update NEW"
}
```

###  Response:

###  Status Code: 200 OK

###  Response Body:
{
    "message": "category updated"
}


###  delete category

###  Endpoint: /api/delete/category

###  HTTP Method: GET

###  URL:
```
 http://localhost:8080/api/delete/category
```

###  Request Body:
```
{
    "id":1,
    
}
```

###  Response:

###  Status Code: 200 OK

###  Response Body:
```
{
    "message": "category deleted"
}
```
If this category is using reference with nay product

###  Response:

###  Status Code: 400 bad request

###  Response Body:
```
{
    "error": "pq: update or delete on table \"categories\" violates foreign key constraint \"products_category_id_fkey\" on table \"products\""
}
```

###  create supplier

###  how to create supplier

###  Endpoint: /api/create/supplier

###  HTTP Method: POST

###  URL:
```
 http://localhost:8080/api/create/supplier
```

### Request Body:

```
{
    "name":"Khulna traders",
    "email":"khulna@me.com",
    "phone":"01737786084",
    "status_id":1,
    "is_verified_supplier":true
}
```

###  Response:

###  Status Code: 200 OK

###  Response Body:
```
{
    "message": "supplier created"
}
```


###  get supplier

###  Endpoint: /api/get/supplier

###  HTTP Method: GET

###  URL: 
```
http://localhost:8080/api/get/supplier
```

###  Request Body:

```
{
    "id":1
}
```

###  Response:

###  Status Code: 200 OK

###  Response Body:
```
{
    "id": 1,
    "name": "Ab Intetnationl",
    "email": "supplier1@example.com",
    "phone": "123-456-7890",
    "status_id": 1,
    "is_verified_supplier": true,
    "created_at": "2023-11-17T18:29:05.727213Z"
}
```

###  update supplier

###  Endpoint: /api/update/supplier

###  HTTP Method: POST

###  URL: 
```
http://localhost:8080/api/update/supplier
```

###  Request Body:
```
{
    "id":1,
    "name":"Ab Intetnationl 1 update",
    "email": "supplier1update@example.com"
}
```


###  Response:

###  Status Code: 200 OK

###  Response Body:

```
{
    "message": "supplier updated"
}
```


###  delete supplier

###  Endpoint: /api/delete/supplier

###  HTTP Method: GET

###  URL: 

```
http://localhost:8080/api/delete/supplier
```

###  Request Body:
```
{
    "id":1,
    
}
```

###  Response:

###  Status Code: 200 OK

###  Response Body:
```
{
    "message": "supplier deleted"
}
```

If this supplier is using reference with nay product

###  Response:

###  Status Code: 400 bad request

###  Response Body:
```
{
    "error": "pq: update or delete on table \"suppliers\" violates foreign key constraint \"products_supplier_id_fkey\" on table \"products\""
}
```