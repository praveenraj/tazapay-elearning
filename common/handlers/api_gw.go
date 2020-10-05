package handlers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog/log"
	"tazapay.com/elearning/common/utils"

	"tazapay.com/elearning/common/constants"
	"tazapay.com/elearning/common/initializer"
	"tazapay.com/elearning/common/request"
	"tazapay.com/elearning/common/responses"
)

// API handler for AWS API-Gateway request
func API(event interface{}, controller func(*request.Payload) *responses.APIResponse) (interface{}, error) {
	// convert lambda req to api request
	var api events.APIGatewayV2HTTPRequest
	err := utils.JSONMarshalAndUnmarshal(event, &api)
	if err != nil {
		log.Error().Msgf("error converting lambda req to api req: %v", err)
		return responses.APIInternalError(), nil
	}

	// convert api request into common payload
	payload, err := request.CreatePayload(api)
	if err != nil {
		log.Error().Msgf("error converting request to payload: %v", err)
		return responses.APIInternalError(), nil
	}

	// log the method, route path, query params, path params, body
	logAPIRequest(payload)

	// initialize the application set up
	err = initializer.AppSetup()
	if err != nil {
		log.Error().Msgf("error during app setup: %v", err)
		return responses.APIInternalError(), nil
	}

	// controller call to process the payload
	return controller(payload), nil
}

// logAPIRequest add API request data to logging
func logAPIRequest(payload *request.Payload) {
	log.Info().Msgf("api: %v %v", payload.GetAPI().HTTPMethod, payload.GetAPI().Resource)
	if len(payload.GetAPI().QueryStringParameters) > constants.IntZero {
		log.Debug().Msgf("queryParams: %v", payload.GetAPI().QueryStringParameters)
	}
	if len(payload.GetAPI().PathParameters) > constants.IntZero {
		log.Debug().Msgf("pathParams: %v", payload.GetAPI().PathParameters)
	}
	if !payload.GetAPI().IsBase64Encoded && payload.GetAPI().Body != constants.Empty {
		log.Debug().Msgf("body: %v", payload.GetAPI().Body)
	}
}
