# Money Transfer API

This Go package provides a simple API for transferring money between accounts.

## Functions

### LoadAccountsFromJSON(filepath string)

This function loads account data from a JSON file and stores it in memory. The JSON file should contain an array of accounts, each with an `ID` and `Balance`.

### ListAccounts(w http.ResponseWriter, r *http.Request)

This is an HTTP handler function that lists all accounts. It responds to `GET` requests at the `/accounts` endpoint.

### Transfer(w http.ResponseWriter, r *http.Request)

This is an HTTP handler function that transfers money from one account to another. It responds to `POST` requests at the `/transfer` endpoint. The request body should be a JSON object with `from`, `to`, and `amount` fields.

## install
1. ensure that you have installed go programming language on your local machine before
2. clone this repository` git clone https://github.com/EslamYasser-Dev/api.git`
3. open terminal go to the repository directory called `moneytrans` or run treminal inside this directory.
4. use command `go run main.go` to run.

## Testing
You can test this API using any HTTP client like curl or Postman. Remember to run your server before testing.

you can change the port number from main `const port string = "whatever port you want to listen on"`

To list all accounts, send a `GET` request to `/accounts`.

To transfer money, send a `POST` request to `/transfer` with a JSON body like this:

sample/transfer
`{
    "FromAccount": "7a3cbfee-2e44-41bd-b89b-ce27b6349562",
    "ToAccount": "4cc8cf60-680d-4e84-9d02-3e4eb7b14be5",
    "Amount": "555.25"
}`
tested with postman

as following
!Alt picture loading
[url=https://ibb.co/ckvNsWY][img]https://i.ibb.co/7jzG03k/postman.png[/img][/url]

