package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// ItemTodoListsItemTasksItemAttachmentSessionsRequestBuilder provides operations to manage the attachmentSessions property of the microsoft.graph.todoTask entity.
type ItemTodoListsItemTasksItemAttachmentSessionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTodoListsItemTasksItemAttachmentSessionsRequestBuilderGetQueryParameters get attachmentSessions from users
type ItemTodoListsItemTasksItemAttachmentSessionsRequestBuilderGetQueryParameters struct {
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
// ItemTodoListsItemTasksItemAttachmentSessionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTodoListsItemTasksItemAttachmentSessionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTodoListsItemTasksItemAttachmentSessionsRequestBuilderGetQueryParameters
}
// NewItemTodoListsItemTasksItemAttachmentSessionsRequestBuilderInternal instantiates a new AttachmentSessionsRequestBuilder and sets the default values.
func NewItemTodoListsItemTasksItemAttachmentSessionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTodoListsItemTasksItemAttachmentSessionsRequestBuilder) {
    m := &ItemTodoListsItemTasksItemAttachmentSessionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/todo/lists/{todoTaskList%2Did}/tasks/{todoTask%2Did}/attachmentSessions{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemTodoListsItemTasksItemAttachmentSessionsRequestBuilder instantiates a new AttachmentSessionsRequestBuilder and sets the default values.
func NewItemTodoListsItemTasksItemAttachmentSessionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTodoListsItemTasksItemAttachmentSessionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTodoListsItemTasksItemAttachmentSessionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemTodoListsItemTasksItemAttachmentSessionsRequestBuilder) Count()(*ItemTodoListsItemTasksItemAttachmentSessionsCountRequestBuilder) {
    return NewItemTodoListsItemTasksItemAttachmentSessionsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get get attachmentSessions from users
func (m *ItemTodoListsItemTasksItemAttachmentSessionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTodoListsItemTasksItemAttachmentSessionsRequestBuilderGetRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AttachmentSessionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateAttachmentSessionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AttachmentSessionCollectionResponseable), nil
}
// ToGetRequestInformation get attachmentSessions from users
func (m *ItemTodoListsItemTasksItemAttachmentSessionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTodoListsItemTasksItemAttachmentSessionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
