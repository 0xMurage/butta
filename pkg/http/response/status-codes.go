package response

import "net/http"

// StatusCode ensure status codes can be used as type
type StatusCode int

const (
	StatusContinue           StatusCode = http.StatusContinue           // RFC 9110, 15.2.1
	StatusSwitchingProtocols StatusCode = http.StatusSwitchingProtocols // RFC 9110, 15.2.2
	StatusProcessing         StatusCode = http.StatusProcessing         // RFC 2518, 10.1
	StatusEarlyHints         StatusCode = http.StatusEarlyHints         // RFC 8297

	StatusOK                   StatusCode = http.StatusOK                   // RFC 9110, 15.3.1
	StatusCreated              StatusCode = http.StatusCreated              // RFC 9110, 15.3.2
	StatusAccepted             StatusCode = http.StatusAccepted             // RFC 9110, 15.3.3
	StatusNonAuthoritativeInfo StatusCode = http.StatusNonAuthoritativeInfo // RFC 9110, 15.3.4
	StatusNoContent            StatusCode = http.StatusNoContent            // RFC 9110, 15.3.5
	StatusResetContent         StatusCode = http.StatusResetContent         // RFC 9110, 15.3.6
	StatusPartialContent       StatusCode = http.StatusPartialContent       // RFC 9110, 15.3.7
	StatusMultiStatus          StatusCode = http.StatusMultiStatus          // RFC 4918, 11.1
	StatusAlreadyReported      StatusCode = http.StatusAlreadyReported      // RFC 5842, 7.1
	StatusIMUsed               StatusCode = http.StatusIMUsed               // RFC 3229, 10.4.1

	StatusMultipleChoices   StatusCode = http.StatusMultipleChoices   // RFC 9110, 15.4.1
	StatusMovedPermanently  StatusCode = http.StatusMovedPermanently  // RFC 9110, 15.4.2
	StatusFound             StatusCode = http.StatusFound             // RFC 9110, 15.4.3
	StatusSeeOther          StatusCode = http.StatusSeeOther          // RFC 9110, 15.4.4
	StatusNotModified       StatusCode = http.StatusNotModified       // RFC 9110, 15.4.5
	StatusUseProxy          StatusCode = http.StatusUseProxy          // RFC 9110, 15.4.6
	_                       StatusCode = 306                          // RFC 9110, 15.4.7 (Unused)
	StatusTemporaryRedirect StatusCode = http.StatusTemporaryRedirect // RFC 9110, 15.4.8
	StatusPermanentRedirect StatusCode = http.StatusPermanentRedirect // RFC 9110, 15.4.9

	StatusBadRequest                   StatusCode = http.StatusBadRequest                   // RFC 9110, 15.5.1
	StatusUnauthorized                 StatusCode = http.StatusUnauthorized                 // RFC 9110, 15.5.2
	StatusPaymentRequired              StatusCode = http.StatusPaymentRequired              // RFC 9110, 15.5.3
	StatusForbidden                    StatusCode = http.StatusForbidden                    // RFC 9110, 15.5.4
	StatusNotFound                     StatusCode = http.StatusNotFound                     // RFC 9110, 15.5.5
	StatusMethodNotAllowed             StatusCode = http.StatusMethodNotAllowed             // RFC 9110, 15.5.6
	StatusNotAcceptable                StatusCode = http.StatusNotAcceptable                // RFC 9110, 15.5.7
	StatusProxyAuthRequired            StatusCode = http.StatusProxyAuthRequired            // RFC 9110, 15.5.8
	StatusRequestTimeout               StatusCode = http.StatusRequestTimeout               // RFC 9110, 15.5.9
	StatusConflict                     StatusCode = http.StatusConflict                     // RFC 9110, 15.5.
	StatusGone                         StatusCode = http.StatusGone                         // RFC 9110, 15.5.11
	StatusLengthRequired               StatusCode = http.StatusLengthRequired               // RFC 9110, 15.5.12
	StatusPreconditionFailed           StatusCode = http.StatusPreconditionFailed           // RFC 9110, 15.5.13
	StatusRequestEntityTooLarge        StatusCode = http.StatusRequestEntityTooLarge        // RFC 9110, 15.5.14
	StatusRequestURITooLong            StatusCode = http.StatusRequestURITooLong            // RFC 9110, 15.5.15
	StatusUnsupportedMediaType         StatusCode = http.StatusUnsupportedMediaType         // RFC 9110, 15.5.16
	StatusRequestedRangeNotSatisfiable StatusCode = http.StatusRequestedRangeNotSatisfiable // RFC 9110, 15.5.17
	StatusExpectationFailed            StatusCode = http.StatusExpectationFailed            // RFC 9110, 15.5.18
	StatusTeapot                       StatusCode = http.StatusTeapot                       // RFC 9110, 15.5.19 (Unused)
	StatusMisdirectedRequest           StatusCode = http.StatusMisdirectedRequest           // RFC 9110, 15.5.20
	StatusUnprocessableEntity          StatusCode = http.StatusUnprocessableEntity          // RFC 9110, 15.5.21
	StatusLocked                       StatusCode = http.StatusLocked                       // RFC 4918, 11.3
	StatusFailedDependency             StatusCode = http.StatusFailedDependency             // RFC 4918, 11.4
	StatusTooEarly                     StatusCode = http.StatusTooEarly                     // RFC 8470, 5.2.
	StatusUpgradeRequired              StatusCode = http.StatusUpgradeRequired              // RFC 9110, 15.5.22
	StatusPreconditionRequired         StatusCode = http.StatusPreconditionRequired         // RFC 6585, 3
	StatusTooManyRequests              StatusCode = http.StatusTooManyRequests              // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  StatusCode = http.StatusRequestHeaderFieldsTooLarge  // RFC 6585, 5
	StatusUnavailableForLegalReasons   StatusCode = http.StatusUnavailableForLegalReasons   // RFC 7725, 3

	StatusInternalServerError           StatusCode = http.StatusInternalServerError           // RFC 9110, 15.6.1
	StatusNotImplemented                StatusCode = http.StatusNotImplemented                // RFC 9110, 15.6.2
	StatusBadGateway                    StatusCode = http.StatusBadGateway                    // RFC 9110, 15.6.3
	StatusServiceUnavailable            StatusCode = http.StatusServiceUnavailable            // RFC 9110, 15.6.4
	StatusGatewayTimeout                StatusCode = http.StatusGatewayTimeout                // RFC 9110, 15.6.5
	StatusHTTPVersionNotSupported       StatusCode = http.StatusHTTPVersionNotSupported       // RFC 9110, 15.6.6
	StatusVariantAlsoNegotiates         StatusCode = http.StatusVariantAlsoNegotiates         // RFC 2295, 8.1
	StatusInsufficientStorage           StatusCode = http.StatusInsufficientStorage           // RFC 4918, 11.5
	StatusLoopDetected                  StatusCode = http.StatusLoopDetected                  // RFC 5842, 7.2
	StatusNotExtended                   StatusCode = http.StatusNotExtended                   // RFC 2774, 7
	StatusNetworkAuthenticationRequired StatusCode = http.StatusNetworkAuthenticationRequired // RFC 6585, 6
)
