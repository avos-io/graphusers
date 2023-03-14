package groups

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemPermissionsItemGrantPostRequestBody 
type ItemSitesItemPermissionsItemGrantPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The recipients property
    recipients []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.DriveRecipientable
    // The roles property
    roles []string
}
// NewItemSitesItemPermissionsItemGrantPostRequestBody instantiates a new ItemSitesItemPermissionsItemGrantPostRequestBody and sets the default values.
func NewItemSitesItemPermissionsItemGrantPostRequestBody()(*ItemSitesItemPermissionsItemGrantPostRequestBody) {
    m := &ItemSitesItemPermissionsItemGrantPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemSitesItemPermissionsItemGrantPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemPermissionsItemGrantPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemPermissionsItemGrantPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemPermissionsItemGrantPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSitesItemPermissionsItemGrantPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["recipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateDriveRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.DriveRecipientable, len(val))
            for i, v := range val {
                res[i] = v.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.DriveRecipientable)
            }
            m.SetRecipients(res)
        }
        return nil
    }
    res["roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoles(res)
        }
        return nil
    }
    return res
}
// GetRecipients gets the recipients property value. The recipients property
func (m *ItemSitesItemPermissionsItemGrantPostRequestBody) GetRecipients()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.DriveRecipientable) {
    return m.recipients
}
// GetRoles gets the roles property value. The roles property
func (m *ItemSitesItemPermissionsItemGrantPostRequestBody) GetRoles()([]string) {
    return m.roles
}
// Serialize serializes information the current object
func (m *ItemSitesItemPermissionsItemGrantPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRecipients() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecipients()))
        for i, v := range m.GetRecipients() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("recipients", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoles() != nil {
        err := writer.WriteCollectionOfStringValues("roles", m.GetRoles())
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
func (m *ItemSitesItemPermissionsItemGrantPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRecipients sets the recipients property value. The recipients property
func (m *ItemSitesItemPermissionsItemGrantPostRequestBody) SetRecipients(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.DriveRecipientable)() {
    m.recipients = value
}
// SetRoles sets the roles property value. The roles property
func (m *ItemSitesItemPermissionsItemGrantPostRequestBody) SetRoles(value []string)() {
    m.roles = value
}
// ItemSitesItemPermissionsItemGrantPostRequestBodyable 
type ItemSitesItemPermissionsItemGrantPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRecipients()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.DriveRecipientable)
    GetRoles()([]string)
    SetRecipients(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.DriveRecipientable)()
    SetRoles(value []string)()
}
