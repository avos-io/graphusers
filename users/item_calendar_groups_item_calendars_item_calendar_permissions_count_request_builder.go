package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// ItemCalendarGroupsItemCalendarsItemCalendarPermissionsCountRequestBuilder provides operations to count the resources in the collection.
type ItemCalendarGroupsItemCalendarsItemCalendarPermissionsCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCalendarGroupsItemCalendarsItemCalendarPermissionsCountRequestBuilderGetQueryParameters get the number of the resource
type ItemCalendarGroupsItemCalendarsItemCalendarPermissionsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
}
// ItemCalendarGroupsItemCalendarsItemCalendarPermissionsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarGroupsItemCalendarsItemCalendarPermissionsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarGroupsItemCalendarsItemCalendarPermissionsCountRequestBuilderGetQueryParameters
}
// NewItemCalendarGroupsItemCalendarsItemCalendarPermissionsCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarPermissionsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemCalendarPermissionsCountRequestBuilder) {
    m := &ItemCalendarGroupsItemCalendarsItemCalendarPermissionsCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarPermissions/$count{?%24filter}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemCalendarGroupsItemCalendarsItemCalendarPermissionsCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarPermissionsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemCalendarPermissionsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarGroupsItemCalendarsItemCalendarPermissionsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *ItemCalendarGroupsItemCalendarsItemCalendarPermissionsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemCalendarPermissionsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *ItemCalendarGroupsItemCalendarsItemCalendarPermissionsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemCalendarPermissionsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "text/plain")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
