package users

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemContactFoldersItemChildFoldersDeltaResponse 
type ItemContactFoldersItemChildFoldersDeltaResponse struct {
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.BaseDeltaFunctionResponse
    // The value property
    value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ContactFolderable
}
// NewItemContactFoldersItemChildFoldersDeltaResponse instantiates a new ItemContactFoldersItemChildFoldersDeltaResponse and sets the default values.
func NewItemContactFoldersItemChildFoldersDeltaResponse()(*ItemContactFoldersItemChildFoldersDeltaResponse) {
    m := &ItemContactFoldersItemChildFoldersDeltaResponse{
        BaseDeltaFunctionResponse: *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.NewBaseDeltaFunctionResponse(),
    }
    return m
}
// CreateItemContactFoldersItemChildFoldersDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemContactFoldersItemChildFoldersDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemContactFoldersItemChildFoldersDeltaResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemContactFoldersItemChildFoldersDeltaResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseDeltaFunctionResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateContactFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ContactFolderable, len(val))
            for i, v := range val {
                res[i] = v.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ContactFolderable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ItemContactFoldersItemChildFoldersDeltaResponse) GetValue()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ContactFolderable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ItemContactFoldersItemChildFoldersDeltaResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseDeltaFunctionResponse.Serialize(writer)
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
func (m *ItemContactFoldersItemChildFoldersDeltaResponse) SetValue(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ContactFolderable)() {
    m.value = value
}
