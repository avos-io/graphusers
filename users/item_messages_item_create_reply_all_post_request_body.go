package users

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMessagesItemCreateReplyAllPostRequestBody 
type ItemMessagesItemCreateReplyAllPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Comment property
    comment *string
    // The Message property
    message i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Messageable
}
// NewItemMessagesItemCreateReplyAllPostRequestBody instantiates a new ItemMessagesItemCreateReplyAllPostRequestBody and sets the default values.
func NewItemMessagesItemCreateReplyAllPostRequestBody()(*ItemMessagesItemCreateReplyAllPostRequestBody) {
    m := &ItemMessagesItemCreateReplyAllPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemMessagesItemCreateReplyAllPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMessagesItemCreateReplyAllPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMessagesItemCreateReplyAllPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemMessagesItemCreateReplyAllPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the comment property value. The Comment property
func (m *ItemMessagesItemCreateReplyAllPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMessagesItemCreateReplyAllPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetMessage gets the message property value. The Message property
func (m *ItemMessagesItemCreateReplyAllPostRequestBody) GetMessage()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Messageable) {
    return m.message
}
// Serialize serializes information the current object
func (m *ItemMessagesItemCreateReplyAllPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemMessagesItemCreateReplyAllPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the comment property value. The Comment property
func (m *ItemMessagesItemCreateReplyAllPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
// SetMessage sets the message property value. The Message property
func (m *ItemMessagesItemCreateReplyAllPostRequestBody) SetMessage(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Messageable)() {
    m.message = value
}
// ItemMessagesItemCreateReplyAllPostRequestBodyable 
type ItemMessagesItemCreateReplyAllPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComment()(*string)
    GetMessage()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Messageable)
    SetComment(value *string)()
    SetMessage(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Messageable)()
}
