package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// ItemOnlineMeetingsCreateOrGetRequestBuilder provides operations to call the createOrGet method.
type ItemOnlineMeetingsCreateOrGetRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemOnlineMeetingsCreateOrGetRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlineMeetingsCreateOrGetRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnlineMeetingsCreateOrGetRequestBuilderInternal instantiates a new CreateOrGetRequestBuilder and sets the default values.
func NewItemOnlineMeetingsCreateOrGetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlineMeetingsCreateOrGetRequestBuilder) {
    m := &ItemOnlineMeetingsCreateOrGetRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/onlineMeetings/createOrGet";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemOnlineMeetingsCreateOrGetRequestBuilder instantiates a new CreateOrGetRequestBuilder and sets the default values.
func NewItemOnlineMeetingsCreateOrGetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlineMeetingsCreateOrGetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnlineMeetingsCreateOrGetRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create an onlineMeeting object with a custom specified external ID. If the external ID already exists, this API will return the onlineMeeting object with that external ID. 
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/onlinemeeting-createorget?view=graph-rest-1.0
func (m *ItemOnlineMeetingsCreateOrGetRequestBuilder) Post(ctx context.Context, body ItemOnlineMeetingsCreateOrGetPostRequestBodyable, requestConfiguration *ItemOnlineMeetingsCreateOrGetRequestBuilderPostRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.OnlineMeetingable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateOnlineMeetingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.OnlineMeetingable), nil
}
// ToPostRequestInformation create an onlineMeeting object with a custom specified external ID. If the external ID already exists, this API will return the onlineMeeting object with that external ID. 
func (m *ItemOnlineMeetingsCreateOrGetRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemOnlineMeetingsCreateOrGetPostRequestBodyable, requestConfiguration *ItemOnlineMeetingsCreateOrGetRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
