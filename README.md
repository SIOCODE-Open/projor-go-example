# ProJor Go Example

This repository contains the source code for the ProJor Go Example application. The repository uses [ProJor](https://docs.siocode.hu/projor), the model-based code generator to maintain most of the source code. See in the `.projor/` directory to examine the schema, model, and templates used to generate the code.

This example is documented [here](https://docs.siocode.hu/projor/full-examples/go).

## Building

* First, you must build the backend by running `go build`
* Then, you must build the frontend by running `npm install` and `npm run build`

## Running

* First, start the backend by executing the built binary file (called `projor-go-example` or `projor-go-example.exe` on Windows)
* Then, start the frontend by running `npm run dev`
* Open your browser and navigate to `http://localhost:3000`
* Use [Postman](https://www.postman.com/) to test the backend API (running on `http://localhost:8080`)

## Documentation

### User Account

The user accounts in the system

Fields:

* `id` : `string` - Unique identifier of the UserAccount
* `username` : `String` - The username of the user account
* `email` : `String` - The email address of the user account
* `flags` : `String` - The flags of the user account

### Product

The products in the system

Fields:

* `id` : `string` - Unique identifier of the Product
* `name` : `String` - The name of the product
* `price` : `Float32` - The price of the product
* `description` : `String` - The description of the product

