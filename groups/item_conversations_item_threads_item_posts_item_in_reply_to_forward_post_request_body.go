package groups

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBody 
type ItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Comment property
    comment *string
    // The ToRecipients property
    toRecipients []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Recipientable
}
// NewItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBody instantiates a new ItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBody and sets the default values.
func NewItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBody()(*ItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBody) {
    m := &ItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the comment property value. The Comment property
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetToRecipients gets the toRecipients property value. The ToRecipients property
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBody) GetToRecipients()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Recipientable) {
    return m.toRecipients
}
// Serialize serializes information the current object
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
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
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the comment property value. The Comment property
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
// SetToRecipients sets the toRecipients property value. The ToRecipients property
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBody) SetToRecipients(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Recipientable)() {
    m.toRecipients = value
}
// ItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBodyable 
type ItemConversationsItemThreadsItemPostsItemInReplyToForwardPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComment()(*string)
    GetToRecipients()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Recipientable)
    SetComment(value *string)()
    SetToRecipients(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Recipientable)()
}
