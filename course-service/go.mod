module tazapay.com/elearning/svc/course

go 1.15

replace tazapay.com/elearning/common v1.0.0 => ../common

require (
	github.com/aws/aws-lambda-go v1.19.1
	github.com/aws/aws-sdk-go v1.35.0
	github.com/jinzhu/gorm v1.9.16
	github.com/rs/zerolog v1.20.0
	github.com/spf13/cast v1.3.1
	github.com/stretchr/testify v1.6.1
	tazapay.com/elearning/common v1.0.0
)
