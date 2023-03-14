package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// ItemOwnedObjectsDirectoryObjectItemRequestBuilder provides operations to manage the ownedObjects property of the microsoft.graph.user entity.
type ItemOwnedObjectsDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemOwnedObjectsDirectoryObjectItemRequestBuilderGetQueryParameters directory objects that are owned by the user. Read-only. Nullable. Supports $expand.
type ItemOwnedObjectsDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemOwnedObjectsDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOwnedObjectsDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOwnedObjectsDirectoryObjectItemRequestBuilderGetQueryParameters
}
// NewItemOwnedObjectsDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemOwnedObjectsDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOwnedObjectsDirectoryObjectItemRequestBuilder) {
    m := &ItemOwnedObjectsDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/ownedObjects/{directoryObject%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemOwnedObjectsDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemOwnedObjectsDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOwnedObjectsDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOwnedObjectsDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get directory objects that are owned by the user. Read-only. Nullable. Supports $expand.
func (m *ItemOwnedObjectsDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOwnedObjectsDirectoryObjectItemRequestBuilderGetRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.DirectoryObjectable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.DirectoryObjectable), nil
}
// GraphApplication casts the previous resource to application.
func (m *ItemOwnedObjectsDirectoryObjectItemRequestBuilder) GraphApplication()(*ItemOwnedObjectsItemGraphApplicationRequestBuilder) {
    return NewItemOwnedObjectsItemGraphApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GraphGroup casts the previous resource to group.
func (m *ItemOwnedObjectsDirectoryObjectItemRequestBuilder) GraphGroup()(*ItemOwnedObjectsItemGraphGroupRequestBuilder) {
    return NewItemOwnedObjectsItemGraphGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
func (m *ItemOwnedObjectsDirectoryObjectItemRequestBuilder) GraphServicePrincipal()(*ItemOwnedObjectsItemGraphServicePrincipalRequestBuilder) {
    return NewItemOwnedObjectsItemGraphServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation directory objects that are owned by the user. Read-only. Nullable. Supports $expand.
func (m *ItemOwnedObjectsDirectoryObjectItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOwnedObjectsDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
