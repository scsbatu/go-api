package middlewares

import (
	"github.com/scsbatu/go-api/utils/ierror"
	"net/http"
)

// Errors for HTTP requests
var (
	ErrStatusBadRequest                    = ierror.New(http.StatusBadRequest, 1001, http.StatusText(http.StatusBadRequest), http.StatusText(http.StatusBadRequest))                                                          // 400  RFC 7231, 6.5.1
	ErrStatusUnauthorized                  = ierror.New(http.StatusUnauthorized, 1002, http.StatusText(http.StatusUnauthorized), http.StatusText(http.StatusUnauthorized))                                                    // 401  RFC 7235, 3.1
	ErrStatusPaymentRequired               = ierror.New(http.StatusPaymentRequired, 1003, http.StatusText(http.StatusPaymentRequired), http.StatusText(http.StatusPaymentRequired))                                           // 402  RFC 7231, 6.5.2
	ErrStatusForbidden                     = ierror.New(http.StatusForbidden, 1004, http.StatusText(http.StatusForbidden), http.StatusText(http.StatusForbidden))                                                             // 403  RFC 7231, 6.5.3
	ErrStatusNotFound                      = ierror.New(http.StatusNotFound, 1005, http.StatusText(http.StatusNotFound), http.StatusText(http.StatusNotFound))                                                                // 404  RFC 7231, 6.5.4
	ErrStatusMethodNotAllowed              = ierror.New(http.StatusMethodNotAllowed, 1006, http.StatusText(http.StatusMethodNotAllowed), http.StatusText(http.StatusMethodNotAllowed))                                        // 405  RFC 7231, 6.5.5
	ErrStatusNotAcceptable                 = ierror.New(http.StatusNotAcceptable, 1007, http.StatusText(http.StatusNotAcceptable), http.StatusText(http.StatusNotAcceptable))                                                 // 406  RFC 7231, 6.5.6
	ErrStatusProxyAuthRequired             = ierror.New(http.StatusProxyAuthRequired, 1008, http.StatusText(http.StatusProxyAuthRequired), http.StatusText(http.StatusProxyAuthRequired))                                     // 407  RFC 7235, 3.2
	ErrStatusRequestTimeout                = ierror.New(http.StatusRequestTimeout, 1009, http.StatusText(http.StatusRequestTimeout), http.StatusText(http.StatusRequestTimeout))                                              // 408  RFC 7231, 6.5.7
	ErrStatusConflict                      = ierror.New(http.StatusConflict, 1010, http.StatusText(http.StatusConflict), http.StatusText(http.StatusConflict))                                                                // 409  RFC 7231, 6.5.8
	ErrStatusGone                          = ierror.New(http.StatusGone, 1011, http.StatusText(http.StatusGone), http.StatusText(http.StatusGone))                                                                            // 410  RFC 7231, 6.5.9
	ErrStatusLengthRequired                = ierror.New(http.StatusLengthRequired, 1012, http.StatusText(http.StatusLengthRequired), http.StatusText(http.StatusLengthRequired))                                              // 411  RFC 7231, 6.5.10
	ErrStatusPreconditionFailed            = ierror.New(http.StatusPreconditionFailed, 1013, http.StatusText(http.StatusPreconditionFailed), http.StatusText(http.StatusPreconditionFailed))                                  // 412  RFC 7232, 4.2
	ErrStatusRequestEntityTooLarge         = ierror.New(http.StatusRequestEntityTooLarge, 1014, http.StatusText(http.StatusRequestEntityTooLarge), http.StatusText(http.StatusRequestEntityTooLarge))                         // 413  RFC 7231, 6.5.11
	ErrStatusRequestURITooLong             = ierror.New(http.StatusRequestURITooLong, 1015, http.StatusText(http.StatusRequestURITooLong), http.StatusText(http.StatusRequestURITooLong))                                     // 414  RFC 7231, 6.5.12
	ErrStatusUnsupportedMediaType          = ierror.New(http.StatusUnsupportedMediaType, 1016, http.StatusText(http.StatusUnsupportedMediaType), http.StatusText(http.StatusUnsupportedMediaType))                            // 415  RFC 7231, 6.5.13
	ErrStatusRequestedRangeNotSatisfiable  = ierror.New(http.StatusRequestedRangeNotSatisfiable, 1017, http.StatusText(http.StatusRequestedRangeNotSatisfiable), http.StatusText(http.StatusRequestedRangeNotSatisfiable))    // 416  RFC 7233, 4.4
	ErrStatusExpectationFailed             = ierror.New(http.StatusExpectationFailed, 1018, http.StatusText(http.StatusExpectationFailed), http.StatusText(http.StatusExpectationFailed))                                     // 417  RFC 7231, 6.5.14
	ErrStatusTeapot                        = ierror.New(http.StatusTeapot, 1019, http.StatusText(http.StatusTeapot), http.StatusText(http.StatusTeapot))                                                                      // 418  RFC 7168, 2.3.3
	ErrStatusMisdirectedRequest            = ierror.New(http.StatusMisdirectedRequest, 1020, http.StatusText(http.StatusMisdirectedRequest), http.StatusText(http.StatusMisdirectedRequest))                                  // 421  RFC 7540, 9.1.2
	ErrStatusUnprocessableEntity           = ierror.New(http.StatusUnprocessableEntity, 1021, http.StatusText(http.StatusUnprocessableEntity), http.StatusText(http.StatusUnprocessableEntity))                               // 422  RFC 4918, 11.2
	ErrStatusLocked                        = ierror.New(http.StatusLocked, 1022, http.StatusText(http.StatusLocked), http.StatusText(http.StatusLocked))                                                                      // 423  RFC 4918, 11.3
	ErrStatusFailedDependency              = ierror.New(http.StatusFailedDependency, 1023, http.StatusText(http.StatusFailedDependency), http.StatusText(http.StatusFailedDependency))                                        // 424  RFC 4918, 11.4
	ErrStatusTooEarly                      = ierror.New(http.StatusTooEarly, 1024, http.StatusText(http.StatusTooEarly), http.StatusText(http.StatusTooEarly))                                                                // 425  RFC 8470, 5.2.
	ErrStatusUpgradeRequired               = ierror.New(http.StatusUpgradeRequired, 1025, http.StatusText(http.StatusUpgradeRequired), http.StatusText(http.StatusUpgradeRequired))                                           // 426  RFC 7231, 6.5.15
	ErrStatusPreconditionRequired          = ierror.New(http.StatusPreconditionRequired, 1026, http.StatusText(http.StatusPreconditionRequired), http.StatusText(http.StatusPreconditionRequired))                            // 428  RFC 6585, 3
	ErrStatusTooManyRequests               = ierror.New(http.StatusTooManyRequests, 1027, http.StatusText(http.StatusTooManyRequests), http.StatusText(http.StatusTooManyRequests))                                           // 429  RFC 6585, 4
	ErrStatusRequestHeaderFieldsTooLarge   = ierror.New(http.StatusRequestHeaderFieldsTooLarge, 1028, http.StatusText(http.StatusRequestHeaderFieldsTooLarge), http.StatusText(http.StatusRequestHeaderFieldsTooLarge))       // 431  RFC 6585, 5
	ErrStatusUnavailableForLegalReasons    = ierror.New(http.StatusUnavailableForLegalReasons, 1029, http.StatusText(http.StatusUnavailableForLegalReasons), http.StatusText(http.StatusUnavailableForLegalReasons))          // 451  RFC 7725, 3
	ErrStatusInternalServerError           = ierror.New(http.StatusInternalServerError, 1030, http.StatusText(http.StatusInternalServerError), http.StatusText(http.StatusInternalServerError))                               // 500  RFC 7231, 6.6.1
	ErrStatusNotImplemented                = ierror.New(http.StatusNotImplemented, 1031, http.StatusText(http.StatusNotImplemented), http.StatusText(http.StatusNotImplemented))                                              // 501  RFC 7231, 6.6.2
	ErrStatusBadGateway                    = ierror.New(http.StatusBadGateway, 1032, http.StatusText(http.StatusBadGateway), http.StatusText(http.StatusBadGateway))                                                          // 502  RFC 7231, 6.6.3
	ErrStatusServiceUnavailable            = ierror.New(http.StatusServiceUnavailable, 1033, http.StatusText(http.StatusServiceUnavailable), http.StatusText(http.StatusServiceUnavailable))                                  // 503  RFC 7231, 6.6.4
	ErrStatusGatewayTimeout                = ierror.New(http.StatusGatewayTimeout, 1034, http.StatusText(http.StatusGatewayTimeout), http.StatusText(http.StatusGatewayTimeout))                                              // 504  RFC 7231, 6.6.5
	ErrStatusHTTPVersionNotSupported       = ierror.New(http.StatusHTTPVersionNotSupported, 1035, http.StatusText(http.StatusHTTPVersionNotSupported), http.StatusText(http.StatusHTTPVersionNotSupported))                   // 505  RFC 7231, 6.6.6
	ErrStatusVariantAlsoNegotiates         = ierror.New(http.StatusVariantAlsoNegotiates, 1036, http.StatusText(http.StatusVariantAlsoNegotiates), http.StatusText(http.StatusVariantAlsoNegotiates))                         // 506  RFC 2295, 8.1
	ErrStatusInsufficientStorage           = ierror.New(http.StatusInsufficientStorage, 1037, http.StatusText(http.StatusInsufficientStorage), http.StatusText(http.StatusInsufficientStorage))                               // 507  RFC 4918, 11.5
	ErrStatusLoopDetected                  = ierror.New(http.StatusLoopDetected, 1038, http.StatusText(http.StatusLoopDetected), http.StatusText(http.StatusLoopDetected))                                                    // 508  RFC 5842, 7.2
	ErrStatusNotExtended                   = ierror.New(http.StatusNotExtended, 1039, http.StatusText(http.StatusNotExtended), http.StatusText(http.StatusNotExtended))                                                       // 510  RFC 2774, 7
	ErrStatusNetworkAuthenticationRequired = ierror.New(http.StatusNetworkAuthenticationRequired, 1040, http.StatusText(http.StatusNetworkAuthenticationRequired), http.StatusText(http.StatusNetworkAuthenticationRequired)) // 511  RFC 6585, 6

	ErrTypeMismatch      = ierror.New(http.StatusBadRequest, 1041, http.StatusText(http.StatusBadRequest), "Type Mismatch")
	ErrParametersMissing = ierror.New(http.StatusBadRequest, 1042, http.StatusText(http.StatusBadRequest), "Parameters missing:")
)
