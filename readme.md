# Money Transfer API

This Go package provides a simple API for transferring money between accounts.

## Functions

### LoadAccountsFromJSON(filepath string)

This function loads account data from a JSON file and stores it in memory. The JSON file should contain an array of accounts, each with an `ID` and `Balance`.

### ListAccounts(w http.ResponseWriter, r *http.Request)

This is an HTTP handler function that lists all accounts. It responds to `GET` requests at the `/accounts` endpoint.

### Transfer(w http.ResponseWriter, r *http.Request)

This is an HTTP handler function that transfers money from one account to another. It responds to `POST` requests at the `/transfer` endpoint. The request body should be a JSON object with `from`, `to`, and `amount` fields.

## Usage

To use this package, import it in your Go application:
1. ensure that you have installed go on your local machine before
2. clone this repository git clone https://github.com/EslamYasser-Dev/api.git
3. open terminal go to the repository directory called moneytrans.
4. use command go run main.go to run.

## Testing
You can test this API using any HTTP client like curl or Postman. Remember to run your server before testing.

To list all accounts, send a GET request to /accounts.

To transfer money, send a POST request to /transfer with a JSON body like this:

JSON
This code is AI-generated. Review and use carefully. Visit our FAQ for more information.

{
    "from": "account1",
    "to": "account2",
    "amount": 50
}