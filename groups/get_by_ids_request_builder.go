package groups

import (
	"context"

	i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetByIdsRequestBuilder provides operations to call the getByIds method.
type GetByIdsRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// GetByIdsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetByIdsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewGetByIdsRequestBuilderInternal instantiates a new GetByIdsRequestBuilder and sets the default values.
func NewGetByIdsRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter,
) *GetByIdsRequestBuilder {
	m := &GetByIdsRequestBuilder{}
	m.urlTemplate = "{+baseurl}/groups/getByIds"
	urlTplParams := make(map[string]string)
	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}
	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter
	return m
}

// NewGetByIdsRequestBuilder instantiates a new GetByIdsRequestBuilder and sets the default values.
func NewGetByIdsRequestBuilder(
	rawUrl string,
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter,
) *GetByIdsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewGetByIdsRequestBuilderInternal(urlParams, requestAdapter)
}

// Post return the directory objects specified in a list of IDs. Some common uses for this function are to:
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/directoryobject-getbyids?view=graph-rest-1.0
func (m *GetByIdsRequestBuilder) Post(
	ctx context.Context,
	body GetByIdsPostRequestBodyable,
	requestConfiguration *GetByIdsRequestBuilderPostRequestConfiguration,
) (GetByIdsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
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
		CreateGetByIdsResponseFromDiscriminatorValue,
		errorMapping,
	)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(GetByIdsResponseable), nil
}

// ToPostRequestInformation return the directory objects specified in a list of IDs. Some common uses for this function are to:
func (m *GetByIdsRequestBuilder) ToPostRequestInformation(
	ctx context.Context,
	body GetByIdsPostRequestBodyable,
	requestConfiguration *GetByIdsRequestBuilderPostRequestConfiguration,
) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.urlTemplate
	requestInfo.PathParameters = m.pathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
	requestInfo.Headers.Add("Accept", "application/json")
	err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
	if err != nil {
		return nil, err
	}
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}
