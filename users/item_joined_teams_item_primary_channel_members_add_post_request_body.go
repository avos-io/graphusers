package users

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemJoinedTeamsItemPrimaryChannelMembersAddPostRequestBody 
type ItemJoinedTeamsItemPrimaryChannelMembersAddPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The values property
    values []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ConversationMemberable
}
// NewItemJoinedTeamsItemPrimaryChannelMembersAddPostRequestBody instantiates a new ItemJoinedTeamsItemPrimaryChannelMembersAddPostRequestBody and sets the default values.
func NewItemJoinedTeamsItemPrimaryChannelMembersAddPostRequestBody()(*ItemJoinedTeamsItemPrimaryChannelMembersAddPostRequestBody) {
    m := &ItemJoinedTeamsItemPrimaryChannelMembersAddPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemJoinedTeamsItemPrimaryChannelMembersAddPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemJoinedTeamsItemPrimaryChannelMembersAddPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedTeamsItemPrimaryChannelMembersAddPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemJoinedTeamsItemPrimaryChannelMembersAddPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemJoinedTeamsItemPrimaryChannelMembersAddPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateConversationMemberFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ConversationMemberable, len(val))
            for i, v := range val {
                res[i] = v.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ConversationMemberable)
            }
            m.SetValues(res)
        }
        return nil
    }
    return res
}
// GetValues gets the values property value. The values property
func (m *ItemJoinedTeamsItemPrimaryChannelMembersAddPostRequestBody) GetValues()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ConversationMemberable) {
    return m.values
}
// Serialize serializes information the current object
func (m *ItemJoinedTeamsItemPrimaryChannelMembersAddPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetValues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValues()))
        for i, v := range m.GetValues() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("values", cast)
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
func (m *ItemJoinedTeamsItemPrimaryChannelMembersAddPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetValues sets the values property value. The values property
func (m *ItemJoinedTeamsItemPrimaryChannelMembersAddPostRequestBody) SetValues(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ConversationMemberable)() {
    m.values = value
}
// ItemJoinedTeamsItemPrimaryChannelMembersAddPostRequestBodyable 
type ItemJoinedTeamsItemPrimaryChannelMembersAddPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValues()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ConversationMemberable)
    SetValues(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ConversationMemberable)()
}
