# Another-Backend-Api

This is a simple project demonstrating some concepts for building web apis in go.
This is an easy expandable REST api. The design allows for having another apis alongside, like graphql or rpc.

## Setting up

We use env file for some enviroment configurations, this allow us to run the api in different enviroments like containers easily.

- Copy the env_sample file as a new .env
    - `$ cp .env_sample .env`

## Running

**Must have docker installed.**

Using compose to automatically build our api and start a mysql instance just run: `$ docker compose up`

Otherwise, to run the api locally you can use the Makefile comands

To create the sample database: `$ make db`

To start the db: `$ make start-db`

To start the api: `$ make run`

## Endpoints

### base api endpoint: `/api`

This endpoint uses basic auth. The username and password can be provided in the .env file.

### products endpoint: `/api/product`

Consists of a CRUD for products accepting methods GET, PUT, POST, DELETE. Examples cam be seen in [Making Requests](#making-requests).



## Resources

A sample product consists of the fields:
```
Product {
    id integer
    name string
    value float
}
```

## Making requests

You can use curl or any other software for making http requests


To create a product:

```
curl --request POST \
  --url http://localhost:8080/api/product/ \
  --header 'Authorization: Basic QURNOkFCQw==' \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "rocket", "value": 123.45
}'
```

To get a product:

```
curl --request GET \
  --url http://localhost:8080/api/product/1 \
  --header 'Authorization: Basic QURNOkFCQw=='
```

To update a product:

```
curl --request PUT \
  --url http://localhost:8080/api/product/1 \
  --header 'Authorization: Basic QURNOkFCQw==' \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "velotrol", "value": 123.45
}'
```

To delete a product:
```
curl --request DELETE \
  --url http://localhost:8080/api/product/1 \
  --header 'Authorization: Basic QURNOkFCQw=='
```
