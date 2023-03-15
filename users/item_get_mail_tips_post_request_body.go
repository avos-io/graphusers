package users

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemGetMailTipsPostRequestBody 
type ItemGetMailTipsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The EmailAddresses property
    emailAddresses []string
    // The MailTipsOptions property
    mailTipsOptions *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MailTipsType
}
// NewItemGetMailTipsPostRequestBody instantiates a new ItemGetMailTipsPostRequestBody and sets the default values.
func NewItemGetMailTipsPostRequestBody()(*ItemGetMailTipsPostRequestBody) {
    m := &ItemGetMailTipsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemGetMailTipsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemGetMailTipsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGetMailTipsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemGetMailTipsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmailAddresses gets the emailAddresses property value. The EmailAddresses property
func (m *ItemGetMailTipsPostRequestBody) GetEmailAddresses()([]string) {
    return m.emailAddresses
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemGetMailTipsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["emailAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEmailAddresses(res)
        }
        return nil
    }
    res["mailTipsOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ParseMailTipsType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailTipsOptions(val.(*i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MailTipsType))
        }
        return nil
    }
    return res
}
// GetMailTipsOptions gets the mailTipsOptions property value. The MailTipsOptions property
func (m *ItemGetMailTipsPostRequestBody) GetMailTipsOptions()(*i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MailTipsType) {
    return m.mailTipsOptions
}
// Serialize serializes information the current object
func (m *ItemGetMailTipsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEmailAddresses() != nil {
        err := writer.WriteCollectionOfStringValues("emailAddresses", m.GetEmailAddresses())
        if err != nil {
            return err
        }
    }
    if m.GetMailTipsOptions() != nil {
        cast := (*m.GetMailTipsOptions()).String()
        err := writer.WriteStringValue("mailTipsOptions", &cast)
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
func (m *ItemGetMailTipsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmailAddresses sets the emailAddresses property value. The EmailAddresses property
func (m *ItemGetMailTipsPostRequestBody) SetEmailAddresses(value []string)() {
    m.emailAddresses = value
}
// SetMailTipsOptions sets the mailTipsOptions property value. The MailTipsOptions property
func (m *ItemGetMailTipsPostRequestBody) SetMailTipsOptions(value *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MailTipsType)() {
    m.mailTipsOptions = value
}
// ItemGetMailTipsPostRequestBodyable 
type ItemGetMailTipsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmailAddresses()([]string)
    GetMailTipsOptions()(*i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MailTipsType)
    SetEmailAddresses(value []string)()
    SetMailTipsOptions(value *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MailTipsType)()
}
