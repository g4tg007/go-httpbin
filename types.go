package httpbin

type ipResponse struct {
	Origin string `json:"origin"`
}

type errorResponse struct {
	Error errObj `json:"error"`
}

type errObj struct {
	Message string `json:"message"`
}

type userAgentResponse struct {
	UA string `json:"user-agent"`
}

type headersResponse struct {
	Headers map[string]string `json:"headers"`
}

type getResponse struct {
	headersResponse
	ipResponse
	URL  string                 `json:"url"`
	Args map[string]interface{} `json:"args"`
}