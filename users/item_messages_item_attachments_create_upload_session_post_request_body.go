package users

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMessagesItemAttachmentsCreateUploadSessionPostRequestBody 
type ItemMessagesItemAttachmentsCreateUploadSessionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The AttachmentItem property
    attachmentItem i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AttachmentItemable
}
// NewItemMessagesItemAttachmentsCreateUploadSessionPostRequestBody instantiates a new ItemMessagesItemAttachmentsCreateUploadSessionPostRequestBody and sets the default values.
func NewItemMessagesItemAttachmentsCreateUploadSessionPostRequestBody()(*ItemMessagesItemAttachmentsCreateUploadSessionPostRequestBody) {
    m := &ItemMessagesItemAttachmentsCreateUploadSessionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemMessagesItemAttachmentsCreateUploadSessionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMessagesItemAttachmentsCreateUploadSessionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMessagesItemAttachmentsCreateUploadSessionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemMessagesItemAttachmentsCreateUploadSessionPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAttachmentItem gets the attachmentItem property value. The AttachmentItem property
func (m *ItemMessagesItemAttachmentsCreateUploadSessionPostRequestBody) GetAttachmentItem()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AttachmentItemable) {
    return m.attachmentItem
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMessagesItemAttachmentsCreateUploadSessionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attachmentItem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateAttachmentItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttachmentItem(val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AttachmentItemable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemMessagesItemAttachmentsCreateUploadSessionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attachmentItem", m.GetAttachmentItem())
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
func (m *ItemMessagesItemAttachmentsCreateUploadSessionPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAttachmentItem sets the attachmentItem property value. The AttachmentItem property
func (m *ItemMessagesItemAttachmentsCreateUploadSessionPostRequestBody) SetAttachmentItem(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AttachmentItemable)() {
    m.attachmentItem = value
}
