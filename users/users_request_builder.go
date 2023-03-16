package users

import (
	"context"

	i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
	i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UsersRequestBuilder provides operations to manage the collection of user entities.
type UsersRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// UsersRequestBuilderGetQueryParameters retrieve a list of user objects.
type UsersRequestBuilderGetQueryParameters struct {
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
	// Show only the first n items
	Top *int32 `uriparametername:"%24top"`
}

// UsersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *UsersRequestBuilderGetQueryParameters
}

// UsersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewUsersRequestBuilderInternal instantiates a new UsersRequestBuilder and sets the default values.
func NewUsersRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter,
) *UsersRequestBuilder {
	m := &UsersRequestBuilder{}
	m.urlTemplate = "{+baseurl}/users{?%24top,%24search,%24filter,%24count,%24orderby,%24select,%24expand}"
	urlTplParams := make(map[string]string)
	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}
	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter
	return m
}

// NewUsersRequestBuilder instantiates a new UsersRequestBuilder and sets the default values.
func NewUsersRequestBuilder(
	rawUrl string,
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter,
) *UsersRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewUsersRequestBuilderInternal(urlParams, requestAdapter)
}

// Count provides operations to count the resources in the collection.
func (m *UsersRequestBuilder) Count() *CountRequestBuilder {
	return NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Delta provides operations to call the delta method.
func (m *UsersRequestBuilder) Delta() *DeltaRequestBuilder {
	return NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Get retrieve a list of user objects.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/user-list?view=graph-rest-1.0
func (m *UsersRequestBuilder) Get(
	ctx context.Context,
	requestConfiguration *UsersRequestBuilderGetRequestConfiguration,
) (i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.UserCollectionResponseable, error) {
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
		i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateUserCollectionResponseFromDiscriminatorValue,
		errorMapping,
	)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.UserCollectionResponseable), nil
}

// GetAvailableExtensionProperties provides operations to call the getAvailableExtensionProperties method.
func (m *UsersRequestBuilder) GetAvailableExtensionProperties() *GetAvailableExtensionPropertiesRequestBuilder {
	return NewGetAvailableExtensionPropertiesRequestBuilderInternal(
		m.pathParameters,
		m.requestAdapter,
	)
}

// GetByIds provides operations to call the getByIds method.
func (m *UsersRequestBuilder) GetByIds() *GetByIdsRequestBuilder {
	return NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Post create a new user.The request body contains the user to create. At a minimum, you must specify the required properties for the user. You can optionally specify any other writable properties.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/user-post-users?view=graph-rest-1.0
func (m *UsersRequestBuilder) Post(
	ctx context.Context,
	body i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable,
	requestConfiguration *UsersRequestBuilderPostRequestConfiguration,
) (i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable, error) {
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
		i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateUserFromDiscriminatorValue,
		errorMapping,
	)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable), nil
}

// ToGetRequestInformation retrieve a list of user objects.
func (m *UsersRequestBuilder) ToGetRequestInformation(
	ctx context.Context,
	requestConfiguration *UsersRequestBuilderGetRequestConfiguration,
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

// ToPostRequestInformation create a new user.The request body contains the user to create. At a minimum, you must specify the required properties for the user. You can optionally specify any other writable properties.
func (m *UsersRequestBuilder) ToPostRequestInformation(
	ctx context.Context,
	body i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable,
	requestConfiguration *UsersRequestBuilderPostRequestConfiguration,
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

// ValidateProperties provides operations to call the validateProperties method.
func (m *UsersRequestBuilder) ValidateProperties() *ValidatePropertiesRequestBuilder {
	return NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
