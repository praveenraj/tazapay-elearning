package responses

// Response common response structure for the application
type Response struct {
	HTTPStatusCode int    `json:"http_status_code"`
	Status         string `json:"status"`
	MetaData
}

// MetaData meta values for the common response object
type MetaData struct {
	Message string            `json:"message,omitempty"`
	Data    *Data             `json:"data,omitempty"`
	Error   *Err              `json:"error,omitempty"`
	Cookies []string          `json:"cookies,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}

// Data response's data structure
type Data struct {
	Value interface{} `json:"value,omitempty"`
}

// Err response's error structure
type Err struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Type    string `json:"type,omitempty"`
	Remarks string `json:"remarks,omitempty"`
}

// APIResponse custom response structure for AWS API-Gateway
type APIResponse struct {
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
	Cookies         []string          `json:"cookies"`
	IsBase64Encoded bool              `json:"isBase64Encoded,omitempty"`
}

// APIResponseBody response body structure for AWS API-Gateway
type APIResponseBody struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}
