package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i1c6ae1291e6b78186b51838b365482c969557ded4a1e546b08b36a68fcb4c237 "github.com/avos-io/graphusers/models/termstore"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// ItemSitesItemTermStoreRequestBuilder provides operations to manage the termStore property of the microsoft.graph.site entity.
type ItemSitesItemTermStoreRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemTermStoreRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoreRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemTermStoreRequestBuilderGetQueryParameters read the properties and relationships of a store object.
type ItemSitesItemTermStoreRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemTermStoreRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoreRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemTermStoreRequestBuilderGetQueryParameters
}
// ItemSitesItemTermStoreRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoreRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemTermStoreRequestBuilderInternal instantiates a new TermStoreRequestBuilder and sets the default values.
func NewItemSitesItemTermStoreRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermStoreRequestBuilder) {
    m := &ItemSitesItemTermStoreRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStore{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemTermStoreRequestBuilder instantiates a new TermStoreRequestBuilder and sets the default values.
func NewItemSitesItemTermStoreRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermStoreRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemTermStoreRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property termStore for groups
func (m *ItemSitesItemTermStoreRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemTermStoreRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read the properties and relationships of a store object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/termstore-store-get?view=graph-rest-1.0
func (m *ItemSitesItemTermStoreRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemTermStoreRequestBuilderGetRequestConfiguration)(i1c6ae1291e6b78186b51838b365482c969557ded4a1e546b08b36a68fcb4c237.Storeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i1c6ae1291e6b78186b51838b365482c969557ded4a1e546b08b36a68fcb4c237.CreateStoreFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i1c6ae1291e6b78186b51838b365482c969557ded4a1e546b08b36a68fcb4c237.Storeable), nil
}
// Groups provides operations to manage the groups property of the microsoft.graph.termStore.store entity.
func (m *ItemSitesItemTermStoreRequestBuilder) Groups()(*ItemSitesItemTermStoreGroupsRequestBuilder) {
    return NewItemSitesItemTermStoreGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GroupsById provides operations to manage the groups property of the microsoft.graph.termStore.store entity.
func (m *ItemSitesItemTermStoreRequestBuilder) GroupsById(id string)(*ItemSitesItemTermStoreGroupsGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group%2Did1"] = id
    }
    return NewItemSitesItemTermStoreGroupsGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the properties of a store object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/termstore-store-update?view=graph-rest-1.0
func (m *ItemSitesItemTermStoreRequestBuilder) Patch(ctx context.Context, body i1c6ae1291e6b78186b51838b365482c969557ded4a1e546b08b36a68fcb4c237.Storeable, requestConfiguration *ItemSitesItemTermStoreRequestBuilderPatchRequestConfiguration)(i1c6ae1291e6b78186b51838b365482c969557ded4a1e546b08b36a68fcb4c237.Storeable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i1c6ae1291e6b78186b51838b365482c969557ded4a1e546b08b36a68fcb4c237.CreateStoreFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i1c6ae1291e6b78186b51838b365482c969557ded4a1e546b08b36a68fcb4c237.Storeable), nil
}
// Sets provides operations to manage the sets property of the microsoft.graph.termStore.store entity.
func (m *ItemSitesItemTermStoreRequestBuilder) Sets()(*ItemSitesItemTermStoreSetsRequestBuilder) {
    return NewItemSitesItemTermStoreSetsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SetsById provides operations to manage the sets property of the microsoft.graph.termStore.store entity.
func (m *ItemSitesItemTermStoreRequestBuilder) SetsById(id string)(*ItemSitesItemTermStoreSetsSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["set%2Did"] = id
    }
    return NewItemSitesItemTermStoreSetsSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property termStore for groups
func (m *ItemSitesItemTermStoreRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermStoreRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a store object.
func (m *ItemSitesItemTermStoreRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermStoreRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a store object.
func (m *ItemSitesItemTermStoreRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i1c6ae1291e6b78186b51838b365482c969557ded4a1e546b08b36a68fcb4c237.Storeable, requestConfiguration *ItemSitesItemTermStoreRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
