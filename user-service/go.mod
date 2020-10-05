module tazapay.com/elearning/svc/user

go 1.15

replace tazapay.com/elearning/common v1.0.0 => ../common

require (
	github.com/aws/aws-lambda-go v1.19.1
	github.com/rs/zerolog v1.20.0
	github.com/stretchr/testify v1.6.1
	tazapay.com/elearning/common v1.0.0
)
