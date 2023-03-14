package users

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamworkSendActivityNotificationPostRequestBody 
type ItemTeamworkSendActivityNotificationPostRequestBody struct {
    // The activityType property
    activityType *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The chainId property
    chainId *int64
    // The previewText property
    previewText i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ItemBodyable
    // The templateParameters property
    templateParameters []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.KeyValuePairable
    // The topic property
    topic i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TeamworkActivityTopicable
}
// NewItemTeamworkSendActivityNotificationPostRequestBody instantiates a new ItemTeamworkSendActivityNotificationPostRequestBody and sets the default values.
func NewItemTeamworkSendActivityNotificationPostRequestBody()(*ItemTeamworkSendActivityNotificationPostRequestBody) {
    m := &ItemTeamworkSendActivityNotificationPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemTeamworkSendActivityNotificationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamworkSendActivityNotificationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamworkSendActivityNotificationPostRequestBody(), nil
}
// GetActivityType gets the activityType property value. The activityType property
func (m *ItemTeamworkSendActivityNotificationPostRequestBody) GetActivityType()(*string) {
    return m.activityType
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemTeamworkSendActivityNotificationPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChainId gets the chainId property value. The chainId property
func (m *ItemTeamworkSendActivityNotificationPostRequestBody) GetChainId()(*int64) {
    return m.chainId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemTeamworkSendActivityNotificationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["activityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityType(val)
        }
        return nil
    }
    res["chainId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChainId(val)
        }
        return nil
    }
    res["previewText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewText(val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ItemBodyable))
        }
        return nil
    }
    res["templateParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.KeyValuePairable)
            }
            m.SetTemplateParameters(res)
        }
        return nil
    }
    res["topic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateTeamworkActivityTopicFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopic(val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TeamworkActivityTopicable))
        }
        return nil
    }
    return res
}
// GetPreviewText gets the previewText property value. The previewText property
func (m *ItemTeamworkSendActivityNotificationPostRequestBody) GetPreviewText()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ItemBodyable) {
    return m.previewText
}
// GetTemplateParameters gets the templateParameters property value. The templateParameters property
func (m *ItemTeamworkSendActivityNotificationPostRequestBody) GetTemplateParameters()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.KeyValuePairable) {
    return m.templateParameters
}
// GetTopic gets the topic property value. The topic property
func (m *ItemTeamworkSendActivityNotificationPostRequestBody) GetTopic()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TeamworkActivityTopicable) {
    return m.topic
}
// Serialize serializes information the current object
func (m *ItemTeamworkSendActivityNotificationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("activityType", m.GetActivityType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("chainId", m.GetChainId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("previewText", m.GetPreviewText())
        if err != nil {
            return err
        }
    }
    if m.GetTemplateParameters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTemplateParameters()))
        for i, v := range m.GetTemplateParameters() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("templateParameters", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("topic", m.GetTopic())
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
// SetActivityType sets the activityType property value. The activityType property
func (m *ItemTeamworkSendActivityNotificationPostRequestBody) SetActivityType(value *string)() {
    m.activityType = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemTeamworkSendActivityNotificationPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChainId sets the chainId property value. The chainId property
func (m *ItemTeamworkSendActivityNotificationPostRequestBody) SetChainId(value *int64)() {
    m.chainId = value
}
// SetPreviewText sets the previewText property value. The previewText property
func (m *ItemTeamworkSendActivityNotificationPostRequestBody) SetPreviewText(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ItemBodyable)() {
    m.previewText = value
}
// SetTemplateParameters sets the templateParameters property value. The templateParameters property
func (m *ItemTeamworkSendActivityNotificationPostRequestBody) SetTemplateParameters(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.KeyValuePairable)() {
    m.templateParameters = value
}
// SetTopic sets the topic property value. The topic property
func (m *ItemTeamworkSendActivityNotificationPostRequestBody) SetTopic(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TeamworkActivityTopicable)() {
    m.topic = value
}
// ItemTeamworkSendActivityNotificationPostRequestBodyable 
type ItemTeamworkSendActivityNotificationPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivityType()(*string)
    GetChainId()(*int64)
    GetPreviewText()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ItemBodyable)
    GetTemplateParameters()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.KeyValuePairable)
    GetTopic()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TeamworkActivityTopicable)
    SetActivityType(value *string)()
    SetChainId(value *int64)()
    SetPreviewText(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ItemBodyable)()
    SetTemplateParameters(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.KeyValuePairable)()
    SetTopic(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TeamworkActivityTopicable)()
}
