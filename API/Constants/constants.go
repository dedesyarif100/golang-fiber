package Constants

import (
	"fmt"
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) constantsHandler(c *fiber.Ctx) error {
	fmt.Println("MASUK CONSTANTS HANDLER")

	// HTTP methods were copied from net/http.
		fmt.Println(fiber.MethodGet)
		fmt.Println(fiber.MethodHead)
		fmt.Println(fiber.MethodPost)
		fmt.Println(fiber.MethodPut)
		fmt.Println(fiber.MethodPatch)
		fmt.Println(fiber.MethodDelete)
		fmt.Println(fiber.MethodConnect)
		fmt.Println(fiber.MethodOptions)
		fmt.Println(fiber.MethodTrace)
		println()

	// MIME types that are commonly used
		fmt.Println(fiber.MIMETextXML)
		fmt.Println(fiber.MIMETextHTML)
		fmt.Println(fiber.MIMETextPlain)
		fmt.Println(fiber.MIMEApplicationXML)
		fmt.Println(fiber.MIMEApplicationJSON)
		fmt.Println(fiber.MIMEApplicationJavaScript)
		fmt.Println(fiber.MIMEApplicationForm)
		fmt.Println(fiber.MIMEOctetStream)
		fmt.Println(fiber.MIMEMultipartForm)
		println()

		fmt.Println(fiber.MIMETextXMLCharsetUTF8)
		fmt.Println(fiber.MIMETextHTMLCharsetUTF8)
		fmt.Println(fiber.MIMETextPlainCharsetUTF8)
		fmt.Println(fiber.MIMEApplicationXMLCharsetUTF8)
		fmt.Println(fiber.MIMEApplicationJSONCharsetUTF8)
		fmt.Println(fiber.MIMEApplicationJavaScriptCharsetUTF8)
		println()

	// HTTP status codes were copied from net/http.
		fmt.Println(fiber.StatusContinue)
		fmt.Println(fiber.StatusSwitchingProtocols)
		fmt.Println(fiber.StatusProcessing)
		fmt.Println(fiber.StatusEarlyHints)
		fmt.Println(fiber.StatusOK)
		fmt.Println(fiber.StatusCreated)
		fmt.Println(fiber.StatusAccepted)
		fmt.Println(fiber.StatusNoContent)
		fmt.Println(fiber.StatusResetContent)
		fmt.Println(fiber.StatusPartialContent)
		fmt.Println(fiber.StatusMultiStatus)
		fmt.Println(fiber.StatusAlreadyReported)
		fmt.Println(fiber.StatusIMUsed)
		fmt.Println(fiber.StatusMultipleChoices)
		fmt.Println(fiber.StatusMovedPermanently)
		fmt.Println(fiber.StatusFound)
		fmt.Println(fiber.StatusSeeOther)
		fmt.Println(fiber.StatusNotModified)
		fmt.Println(fiber.StatusUseProxy)
		fmt.Println(fiber.StatusTemporaryRedirect)
		fmt.Println(fiber.StatusPermanentRedirect)
		fmt.Println(fiber.StatusBadRequest)
		fmt.Println(fiber.StatusUnauthorized)
		fmt.Println(fiber.StatusPaymentRequired)
		fmt.Println(fiber.StatusForbidden)
		fmt.Println(fiber.StatusNotFound)
		fmt.Println(fiber.StatusMethodNotAllowed)
		fmt.Println(fiber.StatusNotAcceptable)
		fmt.Println(fiber.StatusProxyAuthRequired)
		fmt.Println(fiber.StatusRequestTimeout)
		fmt.Println(fiber.StatusConflict)
		fmt.Println(fiber.StatusGone)
		fmt.Println(fiber.StatusLengthRequired)
		fmt.Println(fiber.StatusPreconditionFailed)
		fmt.Println(fiber.StatusRequestEntityTooLarge)
		fmt.Println(fiber.StatusRequestURITooLong)
		fmt.Println(fiber.StatusUnsupportedMediaType)
		fmt.Println(fiber.StatusRequestedRangeNotSatisfiable)
		fmt.Println(fiber.StatusExpectationFailed)
		fmt.Println(fiber.StatusTeapot)
		fmt.Println(fiber.StatusMisdirectedRequest)
		fmt.Println(fiber.StatusUnprocessableEntity)
		fmt.Println(fiber.StatusLocked)
		fmt.Println(fiber.StatusFailedDependency)
		fmt.Println(fiber.StatusTooEarly)
		fmt.Println(fiber.StatusUpgradeRequired)
		fmt.Println(fiber.StatusPreconditionRequired)
		fmt.Println(fiber.StatusTooManyRequests)
		fmt.Println(fiber.StatusRequestHeaderFieldsTooLarge)
		fmt.Println(fiber.StatusUnavailableForLegalReasons)
		fmt.Println(fiber.StatusInternalServerError)
		fmt.Println(fiber.StatusNotImplemented)
		fmt.Println(fiber.StatusBadGateway)
		fmt.Println(fiber.StatusServiceUnavailable)
		fmt.Println(fiber.StatusGatewayTimeout)
		fmt.Println(fiber.StatusHTTPVersionNotSupported)
		fmt.Println(fiber.StatusVariantAlsoNegotiates)
		fmt.Println(fiber.StatusInsufficientStorage)
		fmt.Println(fiber.StatusLoopDetected)
		fmt.Println(fiber.StatusNotExtended)
		fmt.Println(fiber.StatusNetworkAuthenticationRequired)
		println()

	// Errors
		fmt.Println(fiber.NewError(fiber.StatusContinue))
		fmt.Println(fiber.NewError(fiber.StatusSwitchingProtocols))
		fmt.Println(fiber.NewError(fiber.StatusProcessing))
		fmt.Println(fiber.NewError(fiber.StatusEarlyHints))
		fmt.Println(fiber.NewError(fiber.StatusOK))
		fmt.Println(fiber.NewError(fiber.StatusCreated))
		fmt.Println(fiber.NewError(fiber.StatusAccepted))
		fmt.Println(fiber.NewError(fiber.StatusNoContent))
		fmt.Println(fiber.NewError(fiber.StatusResetContent))
		fmt.Println(fiber.NewError(fiber.StatusPartialContent))
		fmt.Println(fiber.NewError(fiber.StatusMultiStatus))
		fmt.Println(fiber.NewError(fiber.StatusAlreadyReported))
		fmt.Println(fiber.NewError(fiber.StatusIMUsed))
		fmt.Println(fiber.NewError(fiber.StatusMultipleChoices))
		fmt.Println(fiber.NewError(fiber.StatusMovedPermanently))
		fmt.Println(fiber.NewError(fiber.StatusFound))
		fmt.Println(fiber.NewError(fiber.StatusSeeOther))
		fmt.Println(fiber.NewError(fiber.StatusNotModified))
		fmt.Println(fiber.NewError(fiber.StatusUseProxy))
		fmt.Println(fiber.NewError(fiber.StatusTemporaryRedirect))
		fmt.Println(fiber.NewError(fiber.StatusPermanentRedirect))
		fmt.Println(fiber.NewError(fiber.StatusBadRequest))
		fmt.Println(fiber.NewError(fiber.StatusUnauthorized))
		fmt.Println(fiber.NewError(fiber.StatusPaymentRequired))
		fmt.Println(fiber.NewError(fiber.StatusForbidden))
		fmt.Println(fiber.NewError(fiber.StatusNotFound))
		fmt.Println(fiber.NewError(fiber.StatusMethodNotAllowed))
		fmt.Println(fiber.NewError(fiber.StatusNotAcceptable))
		fmt.Println(fiber.NewError(fiber.StatusProxyAuthRequired))
		fmt.Println(fiber.NewError(fiber.StatusRequestTimeout))
		fmt.Println(fiber.NewError(fiber.StatusConflict))
		fmt.Println(fiber.NewError(fiber.StatusGone))
		fmt.Println(fiber.NewError(fiber.StatusLengthRequired))
		fmt.Println(fiber.NewError(fiber.StatusPreconditionFailed))
		fmt.Println(fiber.NewError(fiber.StatusRequestEntityTooLarge))
		fmt.Println(fiber.NewError(fiber.StatusRequestURITooLong))
		fmt.Println(fiber.NewError(fiber.StatusUnsupportedMediaType))
		fmt.Println(fiber.NewError(fiber.StatusRequestedRangeNotSatisfiable))
		fmt.Println(fiber.NewError(fiber.StatusExpectationFailed))
		fmt.Println(fiber.NewError(fiber.StatusTeapot))
		fmt.Println(fiber.NewError(fiber.StatusMisdirectedRequest))
		fmt.Println(fiber.NewError(fiber.StatusUnprocessableEntity))
		fmt.Println(fiber.NewError(fiber.StatusLocked))
		fmt.Println(fiber.NewError(fiber.StatusFailedDependency))
		fmt.Println(fiber.NewError(fiber.StatusTooEarly))
		fmt.Println(fiber.NewError(fiber.StatusUpgradeRequired))
		fmt.Println(fiber.NewError(fiber.StatusPreconditionRequired))
		fmt.Println(fiber.NewError(fiber.StatusTooManyRequests))
		fmt.Println(fiber.NewError(fiber.StatusRequestHeaderFieldsTooLarge))
		fmt.Println(fiber.NewError(fiber.StatusUnavailableForLegalReasons))
		fmt.Println(fiber.NewError(fiber.StatusInternalServerError))
		fmt.Println(fiber.NewError(fiber.StatusNotImplemented))
		fmt.Println(fiber.NewError(fiber.StatusBadGateway))
		fmt.Println(fiber.NewError(fiber.StatusServiceUnavailable))
		fmt.Println(fiber.NewError(fiber.StatusGatewayTimeout))
		fmt.Println(fiber.NewError(fiber.StatusHTTPVersionNotSupported))
		fmt.Println(fiber.NewError(fiber.StatusVariantAlsoNegotiates))
		fmt.Println(fiber.NewError(fiber.StatusInsufficientStorage))
		fmt.Println(fiber.NewError(fiber.StatusLoopDetected))
		fmt.Println(fiber.NewError(fiber.StatusNotExtended))
		fmt.Println(fiber.NewError(fiber.StatusNetworkAuthenticationRequired))
		println()

	// HTTP Headers were copied from net/http.
		fmt.Println(fiber.HeaderAuthorization)
		fmt.Println(fiber.HeaderProxyAuthenticate)
		fmt.Println(fiber.HeaderProxyAuthorization)
		fmt.Println(fiber.HeaderWWWAuthenticate)
		fmt.Println(fiber.HeaderAge)
		fmt.Println(fiber.HeaderCacheControl)
		fmt.Println(fiber.HeaderClearSiteData)
		fmt.Println(fiber.HeaderExpires)
		fmt.Println(fiber.HeaderPragma)
		fmt.Println(fiber.HeaderWarning)
		fmt.Println(fiber.HeaderAcceptCH)
		fmt.Println(fiber.HeaderAcceptCHLifetime)
		fmt.Println(fiber.HeaderContentDPR)
		fmt.Println(fiber.HeaderDPR)
		fmt.Println(fiber.HeaderEarlyData)
		fmt.Println(fiber.HeaderSaveData)
		fmt.Println(fiber.HeaderViewportWidth)
		fmt.Println(fiber.HeaderWidth)
		fmt.Println(fiber.HeaderETag)
		fmt.Println(fiber.HeaderIfMatch)
		fmt.Println(fiber.HeaderIfModifiedSince)
		fmt.Println(fiber.HeaderIfNoneMatch)
		fmt.Println(fiber.HeaderIfUnmodifiedSince)
		fmt.Println(fiber.HeaderLastModified)
		fmt.Println(fiber.HeaderVary)
		fmt.Println(fiber.HeaderConnection)
		fmt.Println(fiber.HeaderKeepAlive)
		fmt.Println(fiber.HeaderAccept)
		fmt.Println(fiber.HeaderAcceptCharset)
		fmt.Println(fiber.HeaderAcceptEncoding)
		fmt.Println(fiber.HeaderAcceptLanguage)
		fmt.Println(fiber.HeaderCookie)
		fmt.Println(fiber.HeaderExpect)
		fmt.Println(fiber.HeaderMaxForwards)
		fmt.Println(fiber.HeaderSetCookie)
		fmt.Println(fiber.HeaderAccessControlAllowCredentials)
		fmt.Println(fiber.HeaderAccessControlAllowHeaders)
		fmt.Println(fiber.HeaderAccessControlAllowMethods)
		fmt.Println(fiber.HeaderAccessControlAllowOrigin)
		fmt.Println(fiber.HeaderAccessControlExposeHeaders)
		fmt.Println(fiber.HeaderAccessControlMaxAge)
		fmt.Println(fiber.HeaderAccessControlRequestHeaders)
		fmt.Println(fiber.HeaderAccessControlRequestMethod)
		fmt.Println(fiber.HeaderOrigin)
		fmt.Println(fiber.HeaderTimingAllowOrigin)
		fmt.Println(fiber.HeaderXPermittedCrossDomainPolicies)
		fmt.Println(fiber.HeaderDNT)
		fmt.Println(fiber.HeaderTk)
		fmt.Println(fiber.HeaderContentDisposition)
		fmt.Println(fiber.HeaderContentEncoding)
		fmt.Println(fiber.HeaderContentLanguage)
		fmt.Println(fiber.HeaderContentLength)
		fmt.Println(fiber.HeaderContentLocation)
		fmt.Println(fiber.HeaderContentType)
		fmt.Println(fiber.HeaderForwarded)
		fmt.Println(fiber.HeaderVia)
		fmt.Println(fiber.HeaderXForwardedFor)
		fmt.Println(fiber.HeaderXForwardedHost)
		fmt.Println(fiber.HeaderXForwardedProto)
		fmt.Println(fiber.HeaderXForwardedProtocol)
		fmt.Println(fiber.HeaderXForwardedSsl)
		fmt.Println(fiber.HeaderXUrlScheme)
		fmt.Println(fiber.HeaderLocation)
		fmt.Println(fiber.HeaderFrom)
		fmt.Println(fiber.HeaderHost)
		fmt.Println(fiber.HeaderReferer)
		fmt.Println(fiber.HeaderReferrerPolicy)
		fmt.Println(fiber.HeaderUserAgent)
		fmt.Println(fiber.HeaderAllow)
		fmt.Println(fiber.HeaderServer)
		fmt.Println(fiber.HeaderAcceptRanges)
		fmt.Println(fiber.HeaderContentRange)
		fmt.Println(fiber.HeaderIfRange)
		fmt.Println(fiber.HeaderRange)
		fmt.Println(fiber.HeaderContentSecurityPolicy)
		fmt.Println(fiber.HeaderContentSecurityPolicyReportOnly)
		fmt.Println(fiber.HeaderCrossOriginResourcePolicy)
		fmt.Println(fiber.HeaderExpectCT)
		fmt.Println(fiber.HeaderFeaturePolicy)
		fmt.Println(fiber.HeaderPublicKeyPins)
		fmt.Println(fiber.HeaderPublicKeyPinsReportOnly)
		fmt.Println(fiber.HeaderStrictTransportSecurity)
		fmt.Println(fiber.HeaderUpgradeInsecureRequests)
		fmt.Println(fiber.HeaderXContentTypeOptions)
		fmt.Println(fiber.HeaderXDownloadOptions)
		fmt.Println(fiber.HeaderXFrameOptions)
		fmt.Println(fiber.HeaderXPoweredBy)
		fmt.Println(fiber.HeaderXXSSProtection)
		fmt.Println(fiber.HeaderLastEventID)
		fmt.Println(fiber.HeaderNEL)
		fmt.Println(fiber.HeaderPingFrom)
		fmt.Println(fiber.HeaderPingTo)
		fmt.Println(fiber.HeaderReportTo)
		fmt.Println(fiber.HeaderTE)
		fmt.Println(fiber.HeaderTrailer)
		fmt.Println(fiber.HeaderTransferEncoding)
		fmt.Println(fiber.HeaderSecWebSocketAccept)
		fmt.Println(fiber.HeaderSecWebSocketExtensions)
		fmt.Println(fiber.HeaderSecWebSocketKey)
		fmt.Println(fiber.HeaderSecWebSocketProtocol)
		fmt.Println(fiber.HeaderSecWebSocketVersion)
		fmt.Println(fiber.HeaderAcceptPatch)
		fmt.Println(fiber.HeaderAcceptPushPolicy)
		fmt.Println(fiber.HeaderAcceptSignature)
		fmt.Println(fiber.HeaderAltSvc)
		fmt.Println(fiber.HeaderDate)
		fmt.Println(fiber.HeaderIndex)
		fmt.Println(fiber.HeaderLargeAllocation)
		fmt.Println(fiber.HeaderLink)
		fmt.Println(fiber.HeaderPushPolicy)
		fmt.Println(fiber.HeaderRetryAfter)
		fmt.Println(fiber.HeaderServerTiming)
		fmt.Println(fiber.HeaderSignature)
		fmt.Println(fiber.HeaderSignedHeaders)
		fmt.Println(fiber.HeaderSourceMap)
		fmt.Println(fiber.HeaderUpgrade)
		fmt.Println(fiber.HeaderXDNSPrefetchControl)
		fmt.Println(fiber.HeaderXPingback)
		fmt.Println(fiber.HeaderXRequestID)
		fmt.Println(fiber.HeaderXRequestedWith)
		fmt.Println(fiber.HeaderXRobotsTag)
		fmt.Println(fiber.HeaderXUACompatible)
	return nil
}

func (r *Repository) SetupRoutesConstants(app *fiber.App) {
	api := app.Group("/api", r.constantsHandler)
	constants := api.Group("/constants", r.constantsHandler)
	constants.Get("/user", r.constantsHandler)
}
