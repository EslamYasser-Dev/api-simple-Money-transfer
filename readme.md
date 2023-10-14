# Money Transfer API
it is just a simple API for transferring money between accounts.

## Functions

### LoadAccountsFromJSON(filepath string)

This function loads account data from a JSON file and stores it in memory. The JSON file should contain an array of accounts, each with an `ID` and `Balance`.

### ListAccounts(w http.ResponseWriter, r *http.Request)

This is an HTTP handler function that lists all accounts. It responds to `GET` requests at the `/accounts` endpoint.

### Transfer(w http.ResponseWriter, r *http.Request)

This is an HTTP handler function that transfers money from one account to another. It responds to `POST` requests at the `/transfer` endpoint. The request body should be a JSON object with `from`, `to`, and `amount` fields.

## Usage

To use this package, import it in your Go application:
1. ensure that you have installed go language on your local machine before.
2. clone this repository `git clone https://github.com/EslamYasser-Dev/api.git`
3. open terminal go to the repository directory called `moneytrans` or open the terminal inside this directory/folder.
4. use command `go run main.go` to run.

## Testing
You can test this API using any HTTP client like curl or Postman. Remember to run your server before testing.

To list all accounts, send a `GET` request to /accounts.

To transfer money, send a `POST `request to /transfer with a JSON body like this:

`{
    "from": "account1",   // this must be and id of the sender account
    "to": "account2",       // this must be id for the receiver accout 
    "amount": "50"          
}`

as following
!alt image loading
url : `https://ibb.co/ckvNsWY`