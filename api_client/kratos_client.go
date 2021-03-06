package api_client

import (
	"net/url"
	"sync"

	kratos "github.com/ory/kratos-client-go/client"
)

var (
	publicClientOnce     sync.Once
	publicClientInstance *kratos.OryKratos
)

func InitPublicClient(url url.URL) *kratos.OryKratos {

	publicClientOnce.Do(func() { // <-- atomic, does not allow repeating

		publicClientInstance = kratos.NewHTTPClientWithConfig(
			nil,
			&kratos.TransportConfig{
				Schemes:  []string{url.Scheme},
				Host:     url.Host,
				BasePath: url.Path})
	})

	return publicClientInstance
}

func PublicClient() *kratos.OryKratos {
	return publicClientInstance
}

var (
	adminClientOnce     sync.Once
	adminClientInstance *kratos.OryKratos
)

func InitAdminClient(url url.URL) *kratos.OryKratos {

	adminClientOnce.Do(func() { // <-- atomic, does not allow repeating

		adminClientInstance = kratos.NewHTTPClientWithConfig(
			nil,
			&kratos.TransportConfig{
				Schemes:  []string{url.Scheme},
				Host:     url.Host,
				BasePath: url.Path})
	})

	return adminClientInstance
}

func AdminClient() *kratos.OryKratos {
	return adminClientInstance
}

var (
	whoamiClientOnce     sync.Once
	whoamiClientInstance *kratos.OryKratos
)
