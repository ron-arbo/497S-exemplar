# Sample Microservice

## Group E - Ron Arbo, TJ Goldblatt, Joe Pasquale, Nick Thurai, Juelin Liu

This is a very simple example of a possible microservice to be used in a larger application. It uses HTTP to respond to certain requests in a particular URL form. The code is very simple. The main function starts up an HTTP server, and registers the blank endpoint "/" with the function "MathServer". Math Server then parses the given URL and, if given in the right form, it will perform some math operation on the arguments in that URL and print an output to the user's browser. When used in a larger application, it's obviously not ideal to print directly to some browser screen, but rather send results back to another service. However, for the purpose of demonstration, the service will display the results on the broswer. This service can obviously by optimized and refactored to be more concise and error-proof, but we've built it to be simple to read/follow and demonstrate communication through HTTP.

### API - URL Formatting
The URL you're sending to the service should follow the following format:  
`
<hostname>/<operation>/arg1/arg2/.../argn  
`

Where "operation" has the following possible values:
- add
- subtract
- multiply
- divide

And where each "arg" must be some number

The service will then display the result of the equation in which that operation is applied to each argument in consecutive order. So, for example, locally I write:  
`
localhost:8080/subtract/4/38/23
`

The service will evaluate the equation: 4-38-23 and then display the result to the browser, performing the first operation it sees first.
