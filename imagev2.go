// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// ImageV2Service contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewImageV2Service] method instead.
type ImageV2Service struct {
	Options       []option.RequestOption
	DirectUploads *ImageV2DirectUploadService
}

// NewImageV2Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewImageV2Service(opts ...option.RequestOption) (r *ImageV2Service) {
	r = &ImageV2Service{}
	r.Options = opts
	r.DirectUploads = NewImageV2DirectUploadService(opts...)
	return
}

// List up to 10000 images with one request. Use the optional parameters below to
// get a specific range of images. Endpoint returns continuation_token if more
// images are present.
func (r *ImageV2Service) List(ctx context.Context, accountID string, query ImageV2ListParams, opts ...option.RequestOption) (res *ImageV2ListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ImagesImagesListResponseV2
	path := fmt.Sprintf("accounts/%s/images/v2", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ImagesImagesListResponseV2 struct {
	Errors   []ImagesImagesListResponseV2Error   `json:"errors,required"`
	Messages []ImagesImagesListResponseV2Message `json:"messages,required"`
	Result   ImagesImagesListResponseV2Result    `json:"result,required"`
	// Whether the API call was successful
	Success ImagesImagesListResponseV2Success `json:"success,required"`
	JSON    imagesImagesListResponseV2JSON    `json:"-"`
}

// imagesImagesListResponseV2JSON contains the JSON metadata for the struct
// [ImagesImagesListResponseV2]
type imagesImagesListResponseV2JSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImagesImagesListResponseV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImagesImagesListResponseV2Error struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    imagesImagesListResponseV2ErrorJSON `json:"-"`
}

// imagesImagesListResponseV2ErrorJSON contains the JSON metadata for the struct
// [ImagesImagesListResponseV2Error]
type imagesImagesListResponseV2ErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImagesImagesListResponseV2Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImagesImagesListResponseV2Message struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    imagesImagesListResponseV2MessageJSON `json:"-"`
}

// imagesImagesListResponseV2MessageJSON contains the JSON metadata for the struct
// [ImagesImagesListResponseV2Message]
type imagesImagesListResponseV2MessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImagesImagesListResponseV2Message) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImagesImagesListResponseV2Result struct {
	// Continuation token to fetch next page. Passed as a query param when requesting
	// List V2 api endpoint.
	ContinuationToken string                                  `json:"continuation_token,nullable"`
	Images            []ImagesImagesListResponseV2ResultImage `json:"images"`
	JSON              imagesImagesListResponseV2ResultJSON    `json:"-"`
}

// imagesImagesListResponseV2ResultJSON contains the JSON metadata for the struct
// [ImagesImagesListResponseV2Result]
type imagesImagesListResponseV2ResultJSON struct {
	ContinuationToken apijson.Field
	Images            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ImagesImagesListResponseV2Result) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImagesImagesListResponseV2ResultImage struct {
	// Image unique identifier.
	ID string `json:"id"`
	// Image file name.
	Filename string `json:"filename"`
	// User modifiable key-value store. Can be used for keeping references to another
	// system of record for managing images. Metadata must not exceed 1024 bytes.
	Meta interface{} `json:"meta"`
	// Indicates whether the image can be a accessed only using it's UID. If set to
	// true, a signed token needs to be generated with a signing key to view the image.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// When the media item was uploaded.
	Uploaded time.Time `json:"uploaded" format:"date-time"`
	// Object specifying available variants for an image.
	Variants []ImagesImagesListResponseV2ResultImagesVariant `json:"variants" format:"uri"`
	JSON     imagesImagesListResponseV2ResultImageJSON       `json:"-"`
}

// imagesImagesListResponseV2ResultImageJSON contains the JSON metadata for the
// struct [ImagesImagesListResponseV2ResultImage]
type imagesImagesListResponseV2ResultImageJSON struct {
	ID                apijson.Field
	Filename          apijson.Field
	Meta              apijson.Field
	RequireSignedURLs apijson.Field
	Uploaded          apijson.Field
	Variants          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ImagesImagesListResponseV2ResultImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// URI to thumbnail variant for an image.
//
// Union satisfied by [shared.UnionString], [shared.UnionString] or
// [shared.UnionString].
type ImagesImagesListResponseV2ResultImagesVariant interface {
	ImplementsImagesImagesListResponseV2ResultImagesVariant()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ImagesImagesListResponseV2ResultImagesVariant)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type ImagesImagesListResponseV2Success bool

const (
	ImagesImagesListResponseV2SuccessTrue ImagesImagesListResponseV2Success = true
)

type ImageV2ListResponse struct {
	// Continuation token to fetch next page. Passed as a query param when requesting
	// List V2 api endpoint.
	ContinuationToken string                     `json:"continuation_token,nullable"`
	Images            []ImageV2ListResponseImage `json:"images"`
	JSON              imageV2ListResponseJSON    `json:"-"`
}

// imageV2ListResponseJSON contains the JSON metadata for the struct
// [ImageV2ListResponse]
type imageV2ListResponseJSON struct {
	ContinuationToken apijson.Field
	Images            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ImageV2ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV2ListResponseImage struct {
	// Image unique identifier.
	ID string `json:"id"`
	// Image file name.
	Filename string `json:"filename"`
	// User modifiable key-value store. Can be used for keeping references to another
	// system of record for managing images. Metadata must not exceed 1024 bytes.
	Meta interface{} `json:"meta"`
	// Indicates whether the image can be a accessed only using it's UID. If set to
	// true, a signed token needs to be generated with a signing key to view the image.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// When the media item was uploaded.
	Uploaded time.Time `json:"uploaded" format:"date-time"`
	// Object specifying available variants for an image.
	Variants []ImageV2ListResponseImagesVariant `json:"variants" format:"uri"`
	JSON     imageV2ListResponseImageJSON       `json:"-"`
}

// imageV2ListResponseImageJSON contains the JSON metadata for the struct
// [ImageV2ListResponseImage]
type imageV2ListResponseImageJSON struct {
	ID                apijson.Field
	Filename          apijson.Field
	Meta              apijson.Field
	RequireSignedURLs apijson.Field
	Uploaded          apijson.Field
	Variants          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ImageV2ListResponseImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// URI to thumbnail variant for an image.
//
// Union satisfied by [shared.UnionString], [shared.UnionString] or
// [shared.UnionString].
type ImageV2ListResponseImagesVariant interface {
	ImplementsImageV2ListResponseImagesVariant()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ImageV2ListResponseImagesVariant)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ImageV2ListParams struct {
	// Continuation token for a next page. List images V2 returns continuation_token
	ContinuationToken param.Field[string] `query:"continuation_token"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Sorting order by upload time.
	SortOrder param.Field[ImageV2ListParamsSortOrder] `query:"sort_order"`
}

// URLQuery serializes [ImageV2ListParams]'s query parameters as `url.Values`.
func (r ImageV2ListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sorting order by upload time.
type ImageV2ListParamsSortOrder string

const (
	ImageV2ListParamsSortOrderAsc  ImageV2ListParamsSortOrder = "asc"
	ImageV2ListParamsSortOrderDesc ImageV2ListParamsSortOrder = "desc"
)
