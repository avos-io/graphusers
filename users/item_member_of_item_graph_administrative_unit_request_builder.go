package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// ItemMemberOfItemGraphAdministrativeUnitRequestBuilder casts the previous resource to administrativeUnit.
type ItemMemberOfItemGraphAdministrativeUnitRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMemberOfItemGraphAdministrativeUnitRequestBuilderGetQueryParameters get the item of type microsoft.graph.directoryObject as microsoft.graph.administrativeUnit
type ItemMemberOfItemGraphAdministrativeUnitRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMemberOfItemGraphAdministrativeUnitRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMemberOfItemGraphAdministrativeUnitRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMemberOfItemGraphAdministrativeUnitRequestBuilderGetQueryParameters
}
// NewItemMemberOfItemGraphAdministrativeUnitRequestBuilderInternal instantiates a new GraphAdministrativeUnitRequestBuilder and sets the default values.
func NewItemMemberOfItemGraphAdministrativeUnitRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMemberOfItemGraphAdministrativeUnitRequestBuilder) {
    m := &ItemMemberOfItemGraphAdministrativeUnitRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/memberOf/{directoryObject%2Did}/graph.administrativeUnit{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemMemberOfItemGraphAdministrativeUnitRequestBuilder instantiates a new GraphAdministrativeUnitRequestBuilder and sets the default values.
func NewItemMemberOfItemGraphAdministrativeUnitRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMemberOfItemGraphAdministrativeUnitRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMemberOfItemGraphAdministrativeUnitRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.administrativeUnit
func (m *ItemMemberOfItemGraphAdministrativeUnitRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMemberOfItemGraphAdministrativeUnitRequestBuilderGetRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AdministrativeUnitable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateAdministrativeUnitFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AdministrativeUnitable), nil
}
// ToGetRequestInformation get the item of type microsoft.graph.directoryObject as microsoft.graph.administrativeUnit
func (m *ItemMemberOfItemGraphAdministrativeUnitRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMemberOfItemGraphAdministrativeUnitRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemMemberOfItemGraphAdministrativeUnitRequestBuilder) WithUrl(rawUrl string)(*ItemMemberOfItemGraphAdministrativeUnitRequestBuilder) {
    return NewItemMemberOfItemGraphAdministrativeUnitRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
