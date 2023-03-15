package users

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody 
type ItemTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The attachmentInfo property
    attachmentInfo i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AttachmentInfoable
}
// NewItemTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody instantiates a new ItemTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody and sets the default values.
func NewItemTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody()(*ItemTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody) {
    m := &ItemTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAttachmentInfo gets the attachmentInfo property value. The attachmentInfo property
func (m *ItemTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody) GetAttachmentInfo()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AttachmentInfoable) {
    return m.attachmentInfo
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attachmentInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateAttachmentInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttachmentInfo(val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AttachmentInfoable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attachmentInfo", m.GetAttachmentInfo())
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
func (m *ItemTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAttachmentInfo sets the attachmentInfo property value. The attachmentInfo property
func (m *ItemTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody) SetAttachmentInfo(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AttachmentInfoable)() {
    m.attachmentInfo = value
}
// ItemTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBodyable 
type ItemTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttachmentInfo()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AttachmentInfoable)
    SetAttachmentInfo(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AttachmentInfoable)()
}
