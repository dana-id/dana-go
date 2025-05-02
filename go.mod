module github.com/dana-id/dana-go-api-client

go 1.18

require (
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
	github.com/stretchr/testify v1.10.0
	gopkg.in/validator.v2 v2.0.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/dana-id/dana-go-api-client/payment_gateway => ./payment_gateway

replace github.com/dana-id/dana-go-api-client/utils => ../utils
