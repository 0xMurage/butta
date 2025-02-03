package http

import "net/http"

// ResponseStatusCode ensure status codes can be used as type
type ResponseStatusCode int

const (
	StatusContinue           ResponseStatusCode = http.StatusContinue           // RFC 9110, 15.2.1
	StatusSwitchingProtocols ResponseStatusCode = http.StatusSwitchingProtocols // RFC 9110, 15.2.2
	StatusProcessing         ResponseStatusCode = http.StatusProcessing         // RFC 2518, 10.1
	StatusEarlyHints         ResponseStatusCode = http.StatusEarlyHints         // RFC 8297

	StatusOK                   ResponseStatusCode = http.StatusOK                   // RFC 9110, 15.3.1
	StatusCreated              ResponseStatusCode = http.StatusCreated              // RFC 9110, 15.3.2
	StatusAccepted             ResponseStatusCode = http.StatusAccepted             // RFC 9110, 15.3.3
	StatusNonAuthoritativeInfo ResponseStatusCode = http.StatusNonAuthoritativeInfo // RFC 9110, 15.3.4
	StatusNoContent            ResponseStatusCode = http.StatusNoContent            // RFC 9110, 15.3.5
	StatusResetContent         ResponseStatusCode = http.StatusResetContent         // RFC 9110, 15.3.6
	StatusPartialContent       ResponseStatusCode = http.StatusPartialContent       // RFC 9110, 15.3.7
	StatusMultiStatus          ResponseStatusCode = http.StatusMultiStatus          // RFC 4918, 11.1
	StatusAlreadyReported      ResponseStatusCode = http.StatusAlreadyReported      // RFC 5842, 7.1
	StatusIMUsed               ResponseStatusCode = http.StatusIMUsed               // RFC 3229, 10.4.1

	StatusMultipleChoices   ResponseStatusCode = http.StatusMultipleChoices   // RFC 9110, 15.4.1
	StatusMovedPermanently  ResponseStatusCode = http.StatusMovedPermanently  // RFC 9110, 15.4.2
	StatusFound             ResponseStatusCode = http.StatusFound             // RFC 9110, 15.4.3
	StatusSeeOther          ResponseStatusCode = http.StatusSeeOther          // RFC 9110, 15.4.4
	StatusNotModified       ResponseStatusCode = http.StatusNotModified       // RFC 9110, 15.4.5
	StatusUseProxy          ResponseStatusCode = http.StatusUseProxy          // RFC 9110, 15.4.6
	_                       ResponseStatusCode = 306                          // RFC 9110, 15.4.7 (Unused)
	StatusTemporaryRedirect ResponseStatusCode = http.StatusTemporaryRedirect // RFC 9110, 15.4.8
	StatusPermanentRedirect ResponseStatusCode = http.StatusPermanentRedirect // RFC 9110, 15.4.9

	StatusBadRequest                   ResponseStatusCode = http.StatusBadRequest                   // RFC 9110, 15.5.1
	StatusUnauthorized                 ResponseStatusCode = http.StatusUnauthorized                 // RFC 9110, 15.5.2
	StatusPaymentRequired              ResponseStatusCode = http.StatusPaymentRequired              // RFC 9110, 15.5.3
	StatusForbidden                    ResponseStatusCode = http.StatusForbidden                    // RFC 9110, 15.5.4
	StatusNotFound                     ResponseStatusCode = http.StatusNotFound                     // RFC 9110, 15.5.5
	StatusMethodNotAllowed             ResponseStatusCode = http.StatusMethodNotAllowed             // RFC 9110, 15.5.6
	StatusNotAcceptable                ResponseStatusCode = http.StatusNotAcceptable                // RFC 9110, 15.5.7
	StatusProxyAuthRequired            ResponseStatusCode = http.StatusProxyAuthRequired            // RFC 9110, 15.5.8
	StatusRequestTimeout               ResponseStatusCode = http.StatusRequestTimeout               // RFC 9110, 15.5.9
	StatusConflict                     ResponseStatusCode = http.StatusConflict                     // RFC 9110, 15.5.
	StatusGone                         ResponseStatusCode = http.StatusGone                         // RFC 9110, 15.5.11
	StatusLengthRequired               ResponseStatusCode = http.StatusLengthRequired               // RFC 9110, 15.5.12
	StatusPreconditionFailed           ResponseStatusCode = http.StatusPreconditionFailed           // RFC 9110, 15.5.13
	StatusRequestEntityTooLarge        ResponseStatusCode = http.StatusRequestEntityTooLarge        // RFC 9110, 15.5.14
	StatusRequestURITooLong            ResponseStatusCode = http.StatusRequestURITooLong            // RFC 9110, 15.5.15
	StatusUnsupportedMediaType         ResponseStatusCode = http.StatusUnsupportedMediaType         // RFC 9110, 15.5.16
	StatusRequestedRangeNotSatisfiable ResponseStatusCode = http.StatusRequestedRangeNotSatisfiable // RFC 9110, 15.5.17
	StatusExpectationFailed            ResponseStatusCode = http.StatusExpectationFailed            // RFC 9110, 15.5.18
	StatusTeapot                       ResponseStatusCode = http.StatusTeapot                       // RFC 9110, 15.5.19 (Unused)
	StatusMisdirectedRequest           ResponseStatusCode = http.StatusMisdirectedRequest           // RFC 9110, 15.5.20
	StatusUnprocessableEntity          ResponseStatusCode = http.StatusUnprocessableEntity          // RFC 9110, 15.5.21
	StatusLocked                       ResponseStatusCode = http.StatusLocked                       // RFC 4918, 11.3
	StatusFailedDependency             ResponseStatusCode = http.StatusFailedDependency             // RFC 4918, 11.4
	StatusTooEarly                     ResponseStatusCode = http.StatusTooEarly                     // RFC 8470, 5.2.
	StatusUpgradeRequired              ResponseStatusCode = http.StatusUpgradeRequired              // RFC 9110, 15.5.22
	StatusPreconditionRequired         ResponseStatusCode = http.StatusPreconditionRequired         // RFC 6585, 3
	StatusTooManyRequests              ResponseStatusCode = http.StatusTooManyRequests              // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  ResponseStatusCode = http.StatusRequestHeaderFieldsTooLarge  // RFC 6585, 5
	StatusUnavailableForLegalReasons   ResponseStatusCode = http.StatusUnavailableForLegalReasons   // RFC 7725, 3

	StatusInternalServerError           ResponseStatusCode = http.StatusInternalServerError           // RFC 9110, 15.6.1
	StatusNotImplemented                ResponseStatusCode = http.StatusNotImplemented                // RFC 9110, 15.6.2
	StatusBadGateway                    ResponseStatusCode = http.StatusBadGateway                    // RFC 9110, 15.6.3
	StatusServiceUnavailable            ResponseStatusCode = http.StatusServiceUnavailable            // RFC 9110, 15.6.4
	StatusGatewayTimeout                ResponseStatusCode = http.StatusGatewayTimeout                // RFC 9110, 15.6.5
	StatusHTTPVersionNotSupported       ResponseStatusCode = http.StatusHTTPVersionNotSupported       // RFC 9110, 15.6.6
	StatusVariantAlsoNegotiates         ResponseStatusCode = http.StatusVariantAlsoNegotiates         // RFC 2295, 8.1
	StatusInsufficientStorage           ResponseStatusCode = http.StatusInsufficientStorage           // RFC 4918, 11.5
	StatusLoopDetected                  ResponseStatusCode = http.StatusLoopDetected                  // RFC 5842, 7.2
	StatusNotExtended                   ResponseStatusCode = http.StatusNotExtended                   // RFC 2774, 7
	StatusNetworkAuthenticationRequired ResponseStatusCode = http.StatusNetworkAuthenticationRequired // RFC 6585, 6
)
