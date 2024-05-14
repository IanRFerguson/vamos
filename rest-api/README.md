# REST API

Based on [this DEV walkthrough](https://dev.to/aurelievache/learning-go-by-examples-part-2-create-an-http-rest-api-server-in-go-1cdm)

This is a simple REST API that you can run locally:

* The endpoints and defined in `package/swagger/swagger.yml`
* That source code was generated with `swagger` - see `package/swagger/server/restapi/operations`
* Those endpoints are handled in `internal/main.go`

Everything to run this API is handled by the `Makefile`

```
# Validate the YAML file
make swagger.validate

# Generate the underlying source code
make swagger.gen

# Install the dependencies
make deps

# Run the API
make run

# Compile the source code in an executable binary
make build
```