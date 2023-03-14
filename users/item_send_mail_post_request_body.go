package users

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSendMailPostRequestBody 
type ItemSendMailPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Message property
    message i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Messageable
    // The SaveToSentItems property
    saveToSentItems *bool
}
// NewItemSendMailPostRequestBody instantiates a new ItemSendMailPostRequestBody and sets the default values.
func NewItemSendMailPostRequestBody()(*ItemSendMailPostRequestBody) {
    m := &ItemSendMailPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemSendMailPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSendMailPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSendMailPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSendMailPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSendMailPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Messageable))
        }
        return nil
    }
    res["saveToSentItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSaveToSentItems(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The Message property
func (m *ItemSendMailPostRequestBody) GetMessage()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Messageable) {
    return m.message
}
// GetSaveToSentItems gets the saveToSentItems property value. The SaveToSentItems property
func (m *ItemSendMailPostRequestBody) GetSaveToSentItems()(*bool) {
    return m.saveToSentItems
}
// Serialize serializes information the current object
func (m *ItemSendMailPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("saveToSentItems", m.GetSaveToSentItems())
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
func (m *ItemSendMailPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMessage sets the message property value. The Message property
func (m *ItemSendMailPostRequestBody) SetMessage(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Messageable)() {
    m.message = value
}
// SetSaveToSentItems sets the saveToSentItems property value. The SaveToSentItems property
func (m *ItemSendMailPostRequestBody) SetSaveToSentItems(value *bool)() {
    m.saveToSentItems = value
}
// ItemSendMailPostRequestBodyable 
type ItemSendMailPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMessage()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Messageable)
    GetSaveToSentItems()(*bool)
    SetMessage(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Messageable)()
    SetSaveToSentItems(value *bool)()
}
