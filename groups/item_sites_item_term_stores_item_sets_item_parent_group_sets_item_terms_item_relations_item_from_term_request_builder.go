package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i1c6ae1291e6b78186b51838b365482c969557ded4a1e546b08b36a68fcb4c237 "github.com/avos-io/graphusers/models/termstore"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemTermsItemRelationsItemFromTermRequestBuilder provides operations to manage the fromTerm property of the microsoft.graph.termStore.relation entity.
type ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemTermsItemRelationsItemFromTermRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemTermsItemRelationsItemFromTermRequestBuilderGetQueryParameters the from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set].
type ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemTermsItemRelationsItemFromTermRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemTermsItemRelationsItemFromTermRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemTermsItemRelationsItemFromTermRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemTermsItemRelationsItemFromTermRequestBuilderGetQueryParameters
}
// NewItemSitesItemTermStoresItemSetsItemParentGroupSetsItemTermsItemRelationsItemFromTermRequestBuilderInternal instantiates a new FromTermRequestBuilder and sets the default values.
func NewItemSitesItemTermStoresItemSetsItemParentGroupSetsItemTermsItemRelationsItemFromTermRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemTermsItemRelationsItemFromTermRequestBuilder) {
    m := &ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemTermsItemRelationsItemFromTermRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStores/{store%2Did}/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/terms/{term%2Did}/relations/{relation%2Did}/fromTerm{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemTermStoresItemSetsItemParentGroupSetsItemTermsItemRelationsItemFromTermRequestBuilder instantiates a new FromTermRequestBuilder and sets the default values.
func NewItemSitesItemTermStoresItemSetsItemParentGroupSetsItemTermsItemRelationsItemFromTermRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemTermsItemRelationsItemFromTermRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemTermStoresItemSetsItemParentGroupSetsItemTermsItemRelationsItemFromTermRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set].
func (m *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemTermsItemRelationsItemFromTermRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemTermsItemRelationsItemFromTermRequestBuilderGetRequestConfiguration)(i1c6ae1291e6b78186b51838b365482c969557ded4a1e546b08b36a68fcb4c237.Termable, error) {
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
// ToGetRequestInformation the from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set].
func (m *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemTermsItemRelationsItemFromTermRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemTermsItemRelationsItemFromTermRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
