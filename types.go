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

type cookiesResponse struct {
	Cookies map[string]string `json:"cookies"`
}

type getResponse struct {
	headersResponse
	ipResponse
	URL  string                 `json:"url"`
	Args map[string]interface{} `json:"args"`
}

type gzipResponse struct {
	headersResponse
	ipResponse
	Gzipped bool `json:"gzipped"`
}

type deflateResponse struct {
	headersResponse
	ipResponse
	Deflated bool `json:"deflated"`
}

type basicAuthResponse struct {
	Authenticated bool   `json:"authenticated"`
	User          string `json:"string"`
}
