package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// ItemTodoListsItemTasksTodoTaskItemRequestBuilder provides operations to manage the tasks property of the microsoft.graph.todoTaskList entity.
type ItemTodoListsItemTasksTodoTaskItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTodoListsItemTasksTodoTaskItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTodoListsItemTasksTodoTaskItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTodoListsItemTasksTodoTaskItemRequestBuilderGetQueryParameters the tasks in this task list. Read-only. Nullable.
type ItemTodoListsItemTasksTodoTaskItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTodoListsItemTasksTodoTaskItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTodoListsItemTasksTodoTaskItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTodoListsItemTasksTodoTaskItemRequestBuilderGetQueryParameters
}
// ItemTodoListsItemTasksTodoTaskItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTodoListsItemTasksTodoTaskItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.todoTask entity.
func (m *ItemTodoListsItemTasksTodoTaskItemRequestBuilder) Attachments()(*ItemTodoListsItemTasksItemAttachmentsRequestBuilder) {
    return NewItemTodoListsItemTasksItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.todoTask entity.
func (m *ItemTodoListsItemTasksTodoTaskItemRequestBuilder) AttachmentsById(id string)(*ItemTodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachmentBase%2Did"] = id
    }
    return NewItemTodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// AttachmentSessions provides operations to manage the attachmentSessions property of the microsoft.graph.todoTask entity.
func (m *ItemTodoListsItemTasksTodoTaskItemRequestBuilder) AttachmentSessions()(*ItemTodoListsItemTasksItemAttachmentSessionsRequestBuilder) {
    return NewItemTodoListsItemTasksItemAttachmentSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentSessionsById provides operations to manage the attachmentSessions property of the microsoft.graph.todoTask entity.
func (m *ItemTodoListsItemTasksTodoTaskItemRequestBuilder) AttachmentSessionsById(id string)(*ItemTodoListsItemTasksItemAttachmentSessionsAttachmentSessionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachmentSession%2Did"] = id
    }
    return NewItemTodoListsItemTasksItemAttachmentSessionsAttachmentSessionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ChecklistItems provides operations to manage the checklistItems property of the microsoft.graph.todoTask entity.
func (m *ItemTodoListsItemTasksTodoTaskItemRequestBuilder) ChecklistItems()(*ItemTodoListsItemTasksItemChecklistItemsRequestBuilder) {
    return NewItemTodoListsItemTasksItemChecklistItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ChecklistItemsById provides operations to manage the checklistItems property of the microsoft.graph.todoTask entity.
func (m *ItemTodoListsItemTasksTodoTaskItemRequestBuilder) ChecklistItemsById(id string)(*ItemTodoListsItemTasksItemChecklistItemsChecklistItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["checklistItem%2Did"] = id
    }
    return NewItemTodoListsItemTasksItemChecklistItemsChecklistItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemTodoListsItemTasksTodoTaskItemRequestBuilderInternal instantiates a new TodoTaskItemRequestBuilder and sets the default values.
func NewItemTodoListsItemTasksTodoTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTodoListsItemTasksTodoTaskItemRequestBuilder) {
    m := &ItemTodoListsItemTasksTodoTaskItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/todo/lists/{todoTaskList%2Did}/tasks/{todoTask%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemTodoListsItemTasksTodoTaskItemRequestBuilder instantiates a new TodoTaskItemRequestBuilder and sets the default values.
func NewItemTodoListsItemTasksTodoTaskItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTodoListsItemTasksTodoTaskItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTodoListsItemTasksTodoTaskItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property tasks for users
func (m *ItemTodoListsItemTasksTodoTaskItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTodoListsItemTasksTodoTaskItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Extensions provides operations to manage the extensions property of the microsoft.graph.todoTask entity.
func (m *ItemTodoListsItemTasksTodoTaskItemRequestBuilder) Extensions()(*ItemTodoListsItemTasksItemExtensionsRequestBuilder) {
    return NewItemTodoListsItemTasksItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.todoTask entity.
func (m *ItemTodoListsItemTasksTodoTaskItemRequestBuilder) ExtensionsById(id string)(*ItemTodoListsItemTasksItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewItemTodoListsItemTasksItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the tasks in this task list. Read-only. Nullable.
func (m *ItemTodoListsItemTasksTodoTaskItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTodoListsItemTasksTodoTaskItemRequestBuilderGetRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TodoTaskable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateTodoTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TodoTaskable), nil
}
// LinkedResources provides operations to manage the linkedResources property of the microsoft.graph.todoTask entity.
func (m *ItemTodoListsItemTasksTodoTaskItemRequestBuilder) LinkedResources()(*ItemTodoListsItemTasksItemLinkedResourcesRequestBuilder) {
    return NewItemTodoListsItemTasksItemLinkedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// LinkedResourcesById provides operations to manage the linkedResources property of the microsoft.graph.todoTask entity.
func (m *ItemTodoListsItemTasksTodoTaskItemRequestBuilder) LinkedResourcesById(id string)(*ItemTodoListsItemTasksItemLinkedResourcesLinkedResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["linkedResource%2Did"] = id
    }
    return NewItemTodoListsItemTasksItemLinkedResourcesLinkedResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property tasks in users
func (m *ItemTodoListsItemTasksTodoTaskItemRequestBuilder) Patch(ctx context.Context, body i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TodoTaskable, requestConfiguration *ItemTodoListsItemTasksTodoTaskItemRequestBuilderPatchRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TodoTaskable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateTodoTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TodoTaskable), nil
}
// ToDeleteRequestInformation delete navigation property tasks for users
func (m *ItemTodoListsItemTasksTodoTaskItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTodoListsItemTasksTodoTaskItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the tasks in this task list. Read-only. Nullable.
func (m *ItemTodoListsItemTasksTodoTaskItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTodoListsItemTasksTodoTaskItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property tasks in users
func (m *ItemTodoListsItemTasksTodoTaskItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TodoTaskable, requestConfiguration *ItemTodoListsItemTasksTodoTaskItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
