package users

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMessagesItemForwardPostRequestBody 
type ItemMessagesItemForwardPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Comment property
    comment *string
    // The Message property
    message i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Messageable
    // The ToRecipients property
    toRecipients []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Recipientable
}
// NewItemMessagesItemForwardPostRequestBody instantiates a new ItemMessagesItemForwardPostRequestBody and sets the default values.
func NewItemMessagesItemForwardPostRequestBody()(*ItemMessagesItemForwardPostRequestBody) {
    m := &ItemMessagesItemForwardPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemMessagesItemForwardPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMessagesItemForwardPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMessagesItemForwardPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemMessagesItemForwardPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the comment property value. The Comment property
func (m *ItemMessagesItemForwardPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMessagesItemForwardPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
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
    res["toRecipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Recipientable, len(val))
            for i, v := range val {
                res[i] = v.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Recipientable)
            }
            m.SetToRecipients(res)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The Message property
func (m *ItemMessagesItemForwardPostRequestBody) GetMessage()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Messageable) {
    return m.message
}
// GetToRecipients gets the toRecipients property value. The ToRecipients property
func (m *ItemMessagesItemForwardPostRequestBody) GetToRecipients()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Recipientable) {
    return m.toRecipients
}
// Serialize serializes information the current object
func (m *ItemMessagesItemForwardPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    if m.GetToRecipients() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetToRecipients()))
        for i, v := range m.GetToRecipients() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("toRecipients", cast)
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
func (m *ItemMessagesItemForwardPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the comment property value. The Comment property
func (m *ItemMessagesItemForwardPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
// SetMessage sets the message property value. The Message property
func (m *ItemMessagesItemForwardPostRequestBody) SetMessage(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Messageable)() {
    m.message = value
}
// SetToRecipients sets the toRecipients property value. The ToRecipients property
func (m *ItemMessagesItemForwardPostRequestBody) SetToRecipients(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Recipientable)() {
    m.toRecipients = value
}
// ItemMessagesItemForwardPostRequestBodyable 
type ItemMessagesItemForwardPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComment()(*string)
    GetMessage()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Messageable)
    GetToRecipients()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Recipientable)
    SetComment(value *string)()
    SetMessage(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Messageable)()
    SetToRecipients(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Recipientable)()
}
