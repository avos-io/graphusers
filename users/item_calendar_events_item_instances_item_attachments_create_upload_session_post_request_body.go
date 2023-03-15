package users

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody 
type ItemCalendarEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The AttachmentItem property
    attachmentItem i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AttachmentItemable
}
// NewItemCalendarEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody instantiates a new ItemCalendarEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody and sets the default values.
func NewItemCalendarEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody()(*ItemCalendarEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) {
    m := &ItemCalendarEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCalendarEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCalendarEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAttachmentItem gets the attachmentItem property value. The AttachmentItem property
func (m *ItemCalendarEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) GetAttachmentItem()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AttachmentItemable) {
    return m.attachmentItem
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCalendarEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemCalendarEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemCalendarEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAttachmentItem sets the attachmentItem property value. The AttachmentItem property
func (m *ItemCalendarEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) SetAttachmentItem(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AttachmentItemable)() {
    m.attachmentItem = value
}
// ItemCalendarEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBodyable 
type ItemCalendarEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttachmentItem()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AttachmentItemable)
    SetAttachmentItem(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AttachmentItemable)()
}
