package users

import (
	"context"

	i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeltaRequestBuilder provides operations to call the delta method.
type DeltaRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// DeltaRequestBuilderGetQueryParameters invoke function delta
type DeltaRequestBuilderGetQueryParameters struct {
	// Include count of items
	Count *bool `uriparametername:"%24count"`
	// Filter items by property values
	Filter *string `uriparametername:"%24filter"`
	// Order items by property values
	Orderby []string `uriparametername:"%24orderby"`
	// Search items by search phrases
	Search *string `uriparametername:"%24search"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
	// Skip the first n items
	Skip *int32 `uriparametername:"%24skip"`
	// Show only the first n items
	Top *int32 `uriparametername:"%24top"`
}

// DeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeltaRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *DeltaRequestBuilderGetQueryParameters
}

// NewDeltaRequestBuilderInternal instantiates a new DeltaRequestBuilder and sets the default values.
func NewDeltaRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter,
) *DeltaRequestBuilder {
	m := &DeltaRequestBuilder{}
	m.urlTemplate = "{+baseurl}/users/delta(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}"
	urlTplParams := make(map[string]string)
	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}
	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter
	return m
}

// NewDeltaRequestBuilder instantiates a new DeltaRequestBuilder and sets the default values.
func NewDeltaRequestBuilder(
	rawUrl string,
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter,
) *DeltaRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewDeltaRequestBuilderInternal(urlParams, requestAdapter)
}

// Get invoke function delta
func (m *DeltaRequestBuilder) Get(
	ctx context.Context,
	requestConfiguration *DeltaRequestBuilderGetRequestConfiguration,
) (DeltaResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
		"5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.requestAdapter.Send(
		ctx,
		requestInfo,
		CreateDeltaResponseFromDiscriminatorValue,
		errorMapping,
	)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(DeltaResponseable), nil
}

// ToGetRequestInformation invoke function delta
func (m *DeltaRequestBuilder) ToGetRequestInformation(
	ctx context.Context,
	requestConfiguration *DeltaRequestBuilderGetRequestConfiguration,
) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.urlTemplate
	requestInfo.PathParameters = m.pathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
	requestInfo.Headers.Add("Accept", "application/json")
	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}
