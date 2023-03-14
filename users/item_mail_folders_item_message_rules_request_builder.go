package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// ItemMailFoldersItemMessageRulesRequestBuilder provides operations to manage the messageRules property of the microsoft.graph.mailFolder entity.
type ItemMailFoldersItemMessageRulesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMailFoldersItemMessageRulesRequestBuilderGetQueryParameters get all the messageRule objects defined for the user's inbox.
type ItemMailFoldersItemMessageRulesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemMailFoldersItemMessageRulesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemMessageRulesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMailFoldersItemMessageRulesRequestBuilderGetQueryParameters
}
// ItemMailFoldersItemMessageRulesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemMessageRulesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMailFoldersItemMessageRulesRequestBuilderInternal instantiates a new MessageRulesRequestBuilder and sets the default values.
func NewItemMailFoldersItemMessageRulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemMessageRulesRequestBuilder) {
    m := &ItemMailFoldersItemMessageRulesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messageRules{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMailFoldersItemMessageRulesRequestBuilder instantiates a new MessageRulesRequestBuilder and sets the default values.
func NewItemMailFoldersItemMessageRulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemMessageRulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailFoldersItemMessageRulesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemMailFoldersItemMessageRulesRequestBuilder) Count()(*ItemMailFoldersItemMessageRulesCountRequestBuilder) {
    return NewItemMailFoldersItemMessageRulesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get get all the messageRule objects defined for the user's inbox.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/mailfolder-list-messagerules?view=graph-rest-1.0
func (m *ItemMailFoldersItemMessageRulesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMailFoldersItemMessageRulesRequestBuilderGetRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MessageRuleCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateMessageRuleCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MessageRuleCollectionResponseable), nil
}
// Post create a messageRule object by specifying a set of conditions and actions.  Outlook carries out those actions if an incoming message in the user's Inbox meets the specified conditions.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/mailfolder-post-messagerules?view=graph-rest-1.0
func (m *ItemMailFoldersItemMessageRulesRequestBuilder) Post(ctx context.Context, body i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MessageRuleable, requestConfiguration *ItemMailFoldersItemMessageRulesRequestBuilderPostRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MessageRuleable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateMessageRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MessageRuleable), nil
}
// ToGetRequestInformation get all the messageRule objects defined for the user's inbox.
func (m *ItemMailFoldersItemMessageRulesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMailFoldersItemMessageRulesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a messageRule object by specifying a set of conditions and actions.  Outlook carries out those actions if an incoming message in the user's Inbox meets the specified conditions.
func (m *ItemMailFoldersItemMessageRulesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MessageRuleable, requestConfiguration *ItemMailFoldersItemMessageRulesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
