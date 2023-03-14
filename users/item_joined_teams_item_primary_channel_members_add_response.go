package users

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemJoinedTeamsItemPrimaryChannelMembersAddResponse 
type ItemJoinedTeamsItemPrimaryChannelMembersAddResponse struct {
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.BaseCollectionPaginationCountResponse
    // The value property
    value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ActionResultPartable
}
// NewItemJoinedTeamsItemPrimaryChannelMembersAddResponse instantiates a new ItemJoinedTeamsItemPrimaryChannelMembersAddResponse and sets the default values.
func NewItemJoinedTeamsItemPrimaryChannelMembersAddResponse()(*ItemJoinedTeamsItemPrimaryChannelMembersAddResponse) {
    m := &ItemJoinedTeamsItemPrimaryChannelMembersAddResponse{
        BaseCollectionPaginationCountResponse: *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateItemJoinedTeamsItemPrimaryChannelMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemJoinedTeamsItemPrimaryChannelMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedTeamsItemPrimaryChannelMembersAddResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemJoinedTeamsItemPrimaryChannelMembersAddResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateActionResultPartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ActionResultPartable, len(val))
            for i, v := range val {
                res[i] = v.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ActionResultPartable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ItemJoinedTeamsItemPrimaryChannelMembersAddResponse) GetValue()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ActionResultPartable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ItemJoinedTeamsItemPrimaryChannelMembersAddResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *ItemJoinedTeamsItemPrimaryChannelMembersAddResponse) SetValue(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ActionResultPartable)() {
    m.value = value
}
