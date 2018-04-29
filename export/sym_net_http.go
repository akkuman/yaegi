package export

// Generated by 'go run gen.go net/http'. Do not edit!

import "net/http"

var sym_http = &map[string]interface{}{
	"CanonicalHeaderKey":                  http.CanonicalHeaderKey,
	"Client":                              new(http.Client),
	"CloseNotifier":                       new(http.CloseNotifier),
	"ConnState":                           new(http.ConnState),
	"Cookie":                              new(http.Cookie),
	"CookieJar":                           new(http.CookieJar),
	"DefaultClient":                       http.DefaultClient,
	"DefaultMaxHeaderBytes":               http.DefaultMaxHeaderBytes,
	"DefaultMaxIdleConnsPerHost":          http.DefaultMaxIdleConnsPerHost,
	"DefaultServeMux":                     http.DefaultServeMux,
	"DefaultTransport":                    http.DefaultTransport,
	"DetectContentType":                   http.DetectContentType,
	"Dir":                                 new(http.Dir),
	"ErrAbortHandler":                     http.ErrAbortHandler,
	"ErrBodyNotAllowed":                   http.ErrBodyNotAllowed,
	"ErrBodyReadAfterClose":               http.ErrBodyReadAfterClose,
	"ErrContentLength":                    http.ErrContentLength,
	"ErrHandlerTimeout":                   http.ErrHandlerTimeout,
	"ErrHeaderTooLong":                    http.ErrHeaderTooLong,
	"ErrHijacked":                         http.ErrHijacked,
	"ErrLineTooLong":                      http.ErrLineTooLong,
	"ErrMissingBoundary":                  http.ErrMissingBoundary,
	"ErrMissingContentLength":             http.ErrMissingContentLength,
	"ErrMissingFile":                      http.ErrMissingFile,
	"ErrNoCookie":                         http.ErrNoCookie,
	"ErrNoLocation":                       http.ErrNoLocation,
	"ErrNotMultipart":                     http.ErrNotMultipart,
	"ErrNotSupported":                     http.ErrNotSupported,
	"ErrServerClosed":                     http.ErrServerClosed,
	"ErrShortBody":                        http.ErrShortBody,
	"ErrSkipAltProtocol":                  http.ErrSkipAltProtocol,
	"ErrUnexpectedTrailer":                http.ErrUnexpectedTrailer,
	"ErrUseLastResponse":                  http.ErrUseLastResponse,
	"ErrWriteAfterFlush":                  http.ErrWriteAfterFlush,
	"Error":                               http.Error,
	"File":                                new(http.File),
	"FileServer":                          http.FileServer,
	"FileSystem":                          new(http.FileSystem),
	"Flusher":                             new(http.Flusher),
	"Get":                                 http.Get,
	"Handle":                              http.Handle,
	"HandleFunc":                          http.HandleFunc,
	"Handler":                             new(http.Handler),
	"HandlerFunc":                         new(http.HandlerFunc),
	"Head":                                http.Head,
	"Header":                              new(http.Header),
	"Hijacker":                            new(http.Hijacker),
	"ListenAndServe":                      http.ListenAndServe,
	"ListenAndServeTLS":                   http.ListenAndServeTLS,
	"LocalAddrContextKey":                 http.LocalAddrContextKey,
	"MaxBytesReader":                      http.MaxBytesReader,
	"MethodConnect":                       http.MethodConnect,
	"MethodDelete":                        http.MethodDelete,
	"MethodGet":                           http.MethodGet,
	"MethodHead":                          http.MethodHead,
	"MethodOptions":                       http.MethodOptions,
	"MethodPatch":                         http.MethodPatch,
	"MethodPost":                          http.MethodPost,
	"MethodPut":                           http.MethodPut,
	"MethodTrace":                         http.MethodTrace,
	"NewFileTransport":                    http.NewFileTransport,
	"NewRequest":                          http.NewRequest,
	"NewServeMux":                         http.NewServeMux,
	"NoBody":                              http.NoBody,
	"NotFound":                            http.NotFound,
	"NotFoundHandler":                     http.NotFoundHandler,
	"ParseHTTPVersion":                    http.ParseHTTPVersion,
	"ParseTime":                           http.ParseTime,
	"Post":                                http.Post,
	"PostForm":                            http.PostForm,
	"ProtocolError":                       new(http.ProtocolError),
	"ProxyFromEnvironment":                http.ProxyFromEnvironment,
	"ProxyURL":                            http.ProxyURL,
	"PushOptions":                         new(http.PushOptions),
	"Pusher":                              new(http.Pusher),
	"ReadRequest":                         http.ReadRequest,
	"ReadResponse":                        http.ReadResponse,
	"Redirect":                            http.Redirect,
	"RedirectHandler":                     http.RedirectHandler,
	"Request":                             new(http.Request),
	"Response":                            new(http.Response),
	"ResponseWriter":                      new(http.ResponseWriter),
	"RoundTripper":                        new(http.RoundTripper),
	"Serve":                               http.Serve,
	"ServeContent":                        http.ServeContent,
	"ServeFile":                           http.ServeFile,
	"ServeMux":                            new(http.ServeMux),
	"ServeTLS":                            http.ServeTLS,
	"Server":                              new(http.Server),
	"ServerContextKey":                    http.ServerContextKey,
	"SetCookie":                           http.SetCookie,
	"StateActive":                         http.StateActive,
	"StateClosed":                         http.StateClosed,
	"StateHijacked":                       http.StateHijacked,
	"StateIdle":                           http.StateIdle,
	"StateNew":                            http.StateNew,
	"StatusAccepted":                      http.StatusAccepted,
	"StatusAlreadyReported":               http.StatusAlreadyReported,
	"StatusBadGateway":                    http.StatusBadGateway,
	"StatusBadRequest":                    http.StatusBadRequest,
	"StatusConflict":                      http.StatusConflict,
	"StatusContinue":                      http.StatusContinue,
	"StatusCreated":                       http.StatusCreated,
	"StatusExpectationFailed":             http.StatusExpectationFailed,
	"StatusFailedDependency":              http.StatusFailedDependency,
	"StatusForbidden":                     http.StatusForbidden,
	"StatusFound":                         http.StatusFound,
	"StatusGatewayTimeout":                http.StatusGatewayTimeout,
	"StatusGone":                          http.StatusGone,
	"StatusHTTPVersionNotSupported":       http.StatusHTTPVersionNotSupported,
	"StatusIMUsed":                        http.StatusIMUsed,
	"StatusInsufficientStorage":           http.StatusInsufficientStorage,
	"StatusInternalServerError":           http.StatusInternalServerError,
	"StatusLengthRequired":                http.StatusLengthRequired,
	"StatusLocked":                        http.StatusLocked,
	"StatusLoopDetected":                  http.StatusLoopDetected,
	"StatusMethodNotAllowed":              http.StatusMethodNotAllowed,
	"StatusMovedPermanently":              http.StatusMovedPermanently,
	"StatusMultiStatus":                   http.StatusMultiStatus,
	"StatusMultipleChoices":               http.StatusMultipleChoices,
	"StatusNetworkAuthenticationRequired": http.StatusNetworkAuthenticationRequired,
	"StatusNoContent":                     http.StatusNoContent,
	"StatusNonAuthoritativeInfo":          http.StatusNonAuthoritativeInfo,
	"StatusNotAcceptable":                 http.StatusNotAcceptable,
	"StatusNotExtended":                   http.StatusNotExtended,
	"StatusNotFound":                      http.StatusNotFound,
	"StatusNotImplemented":                http.StatusNotImplemented,
	"StatusNotModified":                   http.StatusNotModified,
	"StatusOK":                            http.StatusOK,
	"StatusPartialContent":                http.StatusPartialContent,
	"StatusPaymentRequired":               http.StatusPaymentRequired,
	"StatusPermanentRedirect":             http.StatusPermanentRedirect,
	"StatusPreconditionFailed":            http.StatusPreconditionFailed,
	"StatusPreconditionRequired":          http.StatusPreconditionRequired,
	"StatusProcessing":                    http.StatusProcessing,
	"StatusProxyAuthRequired":             http.StatusProxyAuthRequired,
	"StatusRequestEntityTooLarge":         http.StatusRequestEntityTooLarge,
	"StatusRequestHeaderFieldsTooLarge":   http.StatusRequestHeaderFieldsTooLarge,
	"StatusRequestTimeout":                http.StatusRequestTimeout,
	"StatusRequestURITooLong":             http.StatusRequestURITooLong,
	"StatusRequestedRangeNotSatisfiable":  http.StatusRequestedRangeNotSatisfiable,
	"StatusResetContent":                  http.StatusResetContent,
	"StatusSeeOther":                      http.StatusSeeOther,
	"StatusServiceUnavailable":            http.StatusServiceUnavailable,
	"StatusSwitchingProtocols":            http.StatusSwitchingProtocols,
	"StatusTeapot":                        http.StatusTeapot,
	"StatusTemporaryRedirect":             http.StatusTemporaryRedirect,
	"StatusText":                          http.StatusText,
	"StatusTooManyRequests":               http.StatusTooManyRequests,
	"StatusUnauthorized":                  http.StatusUnauthorized,
	"StatusUnavailableForLegalReasons":    http.StatusUnavailableForLegalReasons,
	"StatusUnprocessableEntity":           http.StatusUnprocessableEntity,
	"StatusUnsupportedMediaType":          http.StatusUnsupportedMediaType,
	"StatusUpgradeRequired":               http.StatusUpgradeRequired,
	"StatusUseProxy":                      http.StatusUseProxy,
	"StatusVariantAlsoNegotiates":         http.StatusVariantAlsoNegotiates,
	"StripPrefix":                         http.StripPrefix,
	"TimeFormat":                          http.TimeFormat,
	"TimeoutHandler":                      http.TimeoutHandler,
	"TrailerPrefix":                       http.TrailerPrefix,
	"Transport":                           new(http.Transport),
}
