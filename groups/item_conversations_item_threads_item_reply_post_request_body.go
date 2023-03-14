package groups

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemConversationsItemThreadsItemReplyPostRequestBody 
type ItemConversationsItemThreadsItemReplyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Post property
    post i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Postable
}
// NewItemConversationsItemThreadsItemReplyPostRequestBody instantiates a new ItemConversationsItemThreadsItemReplyPostRequestBody and sets the default values.
func NewItemConversationsItemThreadsItemReplyPostRequestBody()(*ItemConversationsItemThreadsItemReplyPostRequestBody) {
    m := &ItemConversationsItemThreadsItemReplyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemConversationsItemThreadsItemReplyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemConversationsItemThreadsItemReplyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemConversationsItemThreadsItemReplyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemConversationsItemThreadsItemReplyPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemConversationsItemThreadsItemReplyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["post"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreatePostFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPost(val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Postable))
        }
        return nil
    }
    return res
}
// GetPost gets the post property value. The Post property
func (m *ItemConversationsItemThreadsItemReplyPostRequestBody) GetPost()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Postable) {
    return m.post
}
// Serialize serializes information the current object
func (m *ItemConversationsItemThreadsItemReplyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("post", m.GetPost())
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
func (m *ItemConversationsItemThreadsItemReplyPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPost sets the post property value. The Post property
func (m *ItemConversationsItemThreadsItemReplyPostRequestBody) SetPost(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Postable)() {
    m.post = value
}
// ItemConversationsItemThreadsItemReplyPostRequestBodyable 
type ItemConversationsItemThreadsItemReplyPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPost()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Postable)
    SetPost(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Postable)()
}
