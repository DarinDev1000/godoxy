package docker_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yusing/go-proxy/internal/docker"
)

func TestExpandWildcard(t *testing.T) {
	labels := map[string]string{
		"proxy.a.host":                "localhost",
		"proxy.b.port":                "4444",
		"proxy.b.scheme":              "http",
		"proxy.*.port":                "5555",
		"proxy.*.healthcheck.disable": "true",
	}

	docker.ExpandWildcard(labels)

	require.Equal(t, map[string]string{
		"proxy.a.host":                "localhost",
		"proxy.a.port":                "5555",
		"proxy.a.healthcheck.disable": "true",
		"proxy.b.scheme":              "http",
		"proxy.b.port":                "5555",
		"proxy.b.healthcheck.disable": "true",
	}, labels)
}

func BenchmarkParseLabels(b *testing.B) {
	for b.Loop() {
		_, _ = docker.ParseLabels(map[string]string{
			"proxy.a.host":   "localhost",
			"proxy.b.port":   "4444",
			"proxy.*.scheme": "http",
			"proxy.*.middlewares.request.hide_headers": "X-Header1,X-Header2",
		})
	}
}
