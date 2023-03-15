package graphusers

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ia61ecdcc3fa2ab2cf62124fbb2909069ac54c44228e3555ad521c3217c49f1b1 "github.com/avos-io/graphusers/users"
    iec5f2438c41fe292e9cdda62cf2ddf63215212b46956129fc2cc02257c7cec8e "github.com/avos-io/graphusers/groups"
)

// Client the main entry point of the SDK, exposes the configuration and the fluent API.
type Client struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewClient instantiates a new Client and sets the default values.
func NewClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Client) {
    m := &Client{
    }
    m.pathParameters = make(map[string]string);
    m.urlTemplate = "{+baseurl}";
    m.requestAdapter = requestAdapter
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.requestAdapter.GetBaseUrl() == "" {
        m.requestAdapter.SetBaseUrl("https://graph.microsoft.com/v1.0")
    }
    m.pathParameters["baseurl"] = m.requestAdapter.GetBaseUrl()
    return m
}
// Groups provides operations to manage the collection of group entities.
func (m *Client) Groups()(*iec5f2438c41fe292e9cdda62cf2ddf63215212b46956129fc2cc02257c7cec8e.GroupsRequestBuilder) {
    return iec5f2438c41fe292e9cdda62cf2ddf63215212b46956129fc2cc02257c7cec8e.NewGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GroupsById provides operations to manage the collection of group entities.
func (m *Client) GroupsById(id string)(*iec5f2438c41fe292e9cdda62cf2ddf63215212b46956129fc2cc02257c7cec8e.GroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group%2Did"] = id
    }
    return iec5f2438c41fe292e9cdda62cf2ddf63215212b46956129fc2cc02257c7cec8e.NewGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Users provides operations to manage the collection of user entities.
func (m *Client) Users()(*ia61ecdcc3fa2ab2cf62124fbb2909069ac54c44228e3555ad521c3217c49f1b1.UsersRequestBuilder) {
    return ia61ecdcc3fa2ab2cf62124fbb2909069ac54c44228e3555ad521c3217c49f1b1.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UsersById provides operations to manage the collection of user entities.
func (m *Client) UsersById(id string)(*ia61ecdcc3fa2ab2cf62124fbb2909069ac54c44228e3555ad521c3217c49f1b1.UserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return ia61ecdcc3fa2ab2cf62124fbb2909069ac54c44228e3555ad521c3217c49f1b1.NewUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
