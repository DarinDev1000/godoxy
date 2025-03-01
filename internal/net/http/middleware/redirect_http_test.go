package middleware

import (
	"net/http"
	"testing"

	"github.com/yusing/go-proxy/internal/net/types"
	. "github.com/yusing/go-proxy/internal/utils/testing"
)

func TestRedirectToHTTPs(t *testing.T) {
	result, err := newMiddlewareTest(RedirectHTTP, &testArgs{
		reqURL: types.MustParseURL("http://example.com"),
	})
	ExpectNoError(t, err)
	ExpectEqual(t, result.ResponseStatus, http.StatusMovedPermanently)
	ExpectEqual(t, result.ResponseHeaders.Get("Location"), "https://example.com")
}

func TestNoRedirect(t *testing.T) {
	result, err := newMiddlewareTest(RedirectHTTP, &testArgs{
		reqURL: types.MustParseURL("https://example.com"),
	})
	ExpectNoError(t, err)
	ExpectEqual(t, result.ResponseStatus, http.StatusOK)
}
