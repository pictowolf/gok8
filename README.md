# GOK8
This is a golang project to make a OpenAPIv3 REST API into kubernetes, mainly for the use with the Infinity Servers Panel.

# Usage
go-swagger3 --module-path . --output .\swaggerui\swagger.yml --schema-without-pkg --generate-yaml true
go run ./main.go
go build .
