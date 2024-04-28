# Golang-CRUD-api

> This GitHub repository provides a robust API implementation for customer, transaction and item The API includes endpoints for customer creation, transaction creation, item creation and transaction fetching.

## Built With
- Golang
- Gin Goonics framework
- Postgresql
- Gorm

## Getting Started

To set up a local version of this project, a collection of steps have been put together to help guide you from installations to usage. Simply follow the guide below and you'll be up and running in no time.

### Set up

- Install Golang, if you haven't already.
- Open Terminal
- Navigate to the preferred location/folder you want the app on your local machine. Use `cd <file-path>` for this.
- Run `git clone https://github.com/Rashad-Muntar/go-api.git` to clone the source folder.
- Now that you have a local copy of the project, navigate to the root of the project folder from your terminal.
- Run create a .env.local file and use this test postgresql `postgres://bamlmoap:GpN_sGGB2qcmfbaHdMLJwIEh-ACwxvWX@flora.db.elephantsql.com/bamlmoap` with a variable name of "DATABASE_URL"
- The app is connected with a Postgresql database but you can choose to add you own database enpoint.
- Run `go run server.go`  to get a server running on your local machine.

## Endpoints
_Base URL_: `http://localhost:8080`

|Description|Method|Endpoint|
|:---|:---|:---|
|Create a customer |POST|`/customer`|
|update a customer |PUT|`/customer`|
|delete a customer |DELETE|`/customer/:id`|
|Create a item |POST|`/item`|
|update a item |PUT|`/item`|
|delete a item |DELETE|`/item/:id`|
|Create a transaction |POST|`/transaction`|
|update a transaction |PUT|`/transaction`|
|delete a transaction |DELETE|`/trasaction/:id`|
|delete a transactions |GET|`/trasactions/` with parameter of `transaction_id, customer_name, item_name`|
  - example of transactions url `http://localhost:8080/transactions?transaction_id=123&customer_name=John&item_name=Computer`

### Usage

At this point, you now have everything you need to properly run the program (source code, Golang, Postgres, Gine, Gorm). If not, refer back to the setup section of this documentation.

👤 **Rashad Muntar**

- GitHub: [@Rashad-Muntar](https://github.com/Rashad-Muntar)
- Twitter - [@RashadToure](https://twitter.com/RashadToure)
- LinkedIn: [Rashad Muntar](https://www.linkedin.com/in/rashad-muntar/)
