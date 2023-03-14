package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i1c6ae1291e6b78186b51838b365482c969557ded4a1e546b08b36a68fcb4c237 "github.com/avos-io/graphusers/models/termstore"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilder provides operations to manage the terms property of the microsoft.graph.termStore.set entity.
type ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilderGetQueryParameters all the terms under the set.
type ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilderGetQueryParameters
}
// ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Children provides operations to manage the children property of the microsoft.graph.termStore.term entity.
func (m *ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilder) Children()(*ItemSitesItemTermStoresItemSetsItemTermsItemChildrenRequestBuilder) {
    return NewItemSitesItemTermStoresItemSetsItemTermsItemChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ChildrenById provides operations to manage the children property of the microsoft.graph.termStore.term entity.
func (m *ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilder) ChildrenById(id string)(*ItemSitesItemTermStoresItemSetsItemTermsItemChildrenTermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term%2Did1"] = id
    }
    return NewItemSitesItemTermStoresItemSetsItemTermsItemChildrenTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilderInternal instantiates a new TermItemRequestBuilder and sets the default values.
func NewItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilder) {
    m := &ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStores/{store%2Did}/sets/{set%2Did}/terms/{term%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilder instantiates a new TermItemRequestBuilder and sets the default values.
func NewItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property terms for groups
func (m *ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get all the terms under the set.
func (m *ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilderGetRequestConfiguration)(i1c6ae1291e6b78186b51838b365482c969557ded4a1e546b08b36a68fcb4c237.Termable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i1c6ae1291e6b78186b51838b365482c969557ded4a1e546b08b36a68fcb4c237.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i1c6ae1291e6b78186b51838b365482c969557ded4a1e546b08b36a68fcb4c237.Termable), nil
}
// Patch update the navigation property terms in groups
func (m *ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilder) Patch(ctx context.Context, body i1c6ae1291e6b78186b51838b365482c969557ded4a1e546b08b36a68fcb4c237.Termable, requestConfiguration *ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilderPatchRequestConfiguration)(i1c6ae1291e6b78186b51838b365482c969557ded4a1e546b08b36a68fcb4c237.Termable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i1c6ae1291e6b78186b51838b365482c969557ded4a1e546b08b36a68fcb4c237.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i1c6ae1291e6b78186b51838b365482c969557ded4a1e546b08b36a68fcb4c237.Termable), nil
}
// Relations provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
func (m *ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilder) Relations()(*ItemSitesItemTermStoresItemSetsItemTermsItemRelationsRequestBuilder) {
    return NewItemSitesItemTermStoresItemSetsItemTermsItemRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RelationsById provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
func (m *ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilder) RelationsById(id string)(*ItemSitesItemTermStoresItemSetsItemTermsItemRelationsRelationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation%2Did"] = id
    }
    return NewItemSitesItemTermStoresItemSetsItemTermsItemRelationsRelationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Set provides operations to manage the set property of the microsoft.graph.termStore.term entity.
func (m *ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilder) Set()(*ItemSitesItemTermStoresItemSetsItemTermsItemSetRequestBuilder) {
    return NewItemSitesItemTermStoresItemSetsItemTermsItemSetRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property terms for groups
func (m *ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation all the terms under the set.
func (m *ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property terms in groups
func (m *ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i1c6ae1291e6b78186b51838b365482c969557ded4a1e546b08b36a68fcb4c237.Termable, requestConfiguration *ItemSitesItemTermStoresItemSetsItemTermsTermItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
