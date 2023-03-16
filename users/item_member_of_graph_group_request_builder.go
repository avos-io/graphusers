package users

import (
	"context"

	i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
	i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMemberOfGraphGroupRequestBuilder casts the previous resource to group.
type ItemMemberOfGraphGroupRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// ItemMemberOfGraphGroupRequestBuilderGetQueryParameters get the items of type microsoft.graph.group in the microsoft.graph.directoryObject collection
type ItemMemberOfGraphGroupRequestBuilderGetQueryParameters struct {
	// Include count of items
	Count *bool `uriparametername:"%24count"`
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
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

// ItemMemberOfGraphGroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMemberOfGraphGroupRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemMemberOfGraphGroupRequestBuilderGetQueryParameters
}

// NewItemMemberOfGraphGroupRequestBuilderInternal instantiates a new GraphGroupRequestBuilder and sets the default values.
func NewItemMemberOfGraphGroupRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter,
) *ItemMemberOfGraphGroupRequestBuilder {
	m := &ItemMemberOfGraphGroupRequestBuilder{}
	m.urlTemplate = "{+baseurl}/users/{user%2Did}/memberOf/graph.group{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}"
	urlTplParams := make(map[string]string)
	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}
	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter
	return m
}

// NewItemMemberOfGraphGroupRequestBuilder instantiates a new GraphGroupRequestBuilder and sets the default values.
func NewItemMemberOfGraphGroupRequestBuilder(
	rawUrl string,
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter,
) *ItemMemberOfGraphGroupRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemMemberOfGraphGroupRequestBuilderInternal(urlParams, requestAdapter)
}

// Count provides operations to count the resources in the collection.
func (m *ItemMemberOfGraphGroupRequestBuilder) Count() *ItemMemberOfGraphGroupCountRequestBuilder {
	return NewItemMemberOfGraphGroupCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Get get the items of type microsoft.graph.group in the microsoft.graph.directoryObject collection
func (m *ItemMemberOfGraphGroupRequestBuilder) Get(
	ctx context.Context,
	requestConfiguration *ItemMemberOfGraphGroupRequestBuilderGetRequestConfiguration,
) (i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.GroupCollectionResponseable, error) {
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
		i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateGroupCollectionResponseFromDiscriminatorValue,
		errorMapping,
	)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.GroupCollectionResponseable), nil
}

// ToGetRequestInformation get the items of type microsoft.graph.group in the microsoft.graph.directoryObject collection
func (m *ItemMemberOfGraphGroupRequestBuilder) ToGetRequestInformation(
	ctx context.Context,
	requestConfiguration *ItemMemberOfGraphGroupRequestBuilderGetRequestConfiguration,
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
