package users

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemChatsItemUnhideForUserPostRequestBody 
type ItemChatsItemUnhideForUserPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The user property
    user i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TeamworkUserIdentityable
}
// NewItemChatsItemUnhideForUserPostRequestBody instantiates a new ItemChatsItemUnhideForUserPostRequestBody and sets the default values.
func NewItemChatsItemUnhideForUserPostRequestBody()(*ItemChatsItemUnhideForUserPostRequestBody) {
    m := &ItemChatsItemUnhideForUserPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemChatsItemUnhideForUserPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemChatsItemUnhideForUserPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChatsItemUnhideForUserPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemChatsItemUnhideForUserPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemChatsItemUnhideForUserPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateTeamworkUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TeamworkUserIdentityable))
        }
        return nil
    }
    return res
}
// GetUser gets the user property value. The user property
func (m *ItemChatsItemUnhideForUserPostRequestBody) GetUser()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TeamworkUserIdentityable) {
    return m.user
}
// Serialize serializes information the current object
func (m *ItemChatsItemUnhideForUserPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("user", m.GetUser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemChatsItemUnhideForUserPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetUser sets the user property value. The user property
func (m *ItemChatsItemUnhideForUserPostRequestBody) SetUser(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TeamworkUserIdentityable)() {
    m.user = value
}
// ItemChatsItemUnhideForUserPostRequestBodyable 
type ItemChatsItemUnhideForUserPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetUser()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TeamworkUserIdentityable)
    SetUser(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TeamworkUserIdentityable)()
}
