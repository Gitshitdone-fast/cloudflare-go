// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// StorageKvService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewStorageKvService] method instead.
type StorageKvService struct {
	Options    []option.RequestOption
	Namespaces *StorageKvNamespaceService
}

// NewStorageKvService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStorageKvService(opts ...option.RequestOption) (r *StorageKvService) {
	r = &StorageKvService{}
	r.Options = opts
	r.Namespaces = NewStorageKvNamespaceService(opts...)
	return
}
