package users

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOnlineMeetingsCreateOrGetPostRequestBody 
type ItemOnlineMeetingsCreateOrGetPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The chatInfo property
    chatInfo i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ChatInfoable
    // The endDateTime property
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The externalId property
    externalId *string
    // The participants property
    participants i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MeetingParticipantsable
    // The startDateTime property
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The subject property
    subject *string
}
// NewItemOnlineMeetingsCreateOrGetPostRequestBody instantiates a new ItemOnlineMeetingsCreateOrGetPostRequestBody and sets the default values.
func NewItemOnlineMeetingsCreateOrGetPostRequestBody()(*ItemOnlineMeetingsCreateOrGetPostRequestBody) {
    m := &ItemOnlineMeetingsCreateOrGetPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemOnlineMeetingsCreateOrGetPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOnlineMeetingsCreateOrGetPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOnlineMeetingsCreateOrGetPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemOnlineMeetingsCreateOrGetPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChatInfo gets the chatInfo property value. The chatInfo property
func (m *ItemOnlineMeetingsCreateOrGetPostRequestBody) GetChatInfo()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ChatInfoable) {
    return m.chatInfo
}
// GetEndDateTime gets the endDateTime property value. The endDateTime property
func (m *ItemOnlineMeetingsCreateOrGetPostRequestBody) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetExternalId gets the externalId property value. The externalId property
func (m *ItemOnlineMeetingsCreateOrGetPostRequestBody) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemOnlineMeetingsCreateOrGetPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["chatInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateChatInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChatInfo(val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ChatInfoable))
        }
        return nil
    }
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["participants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateMeetingParticipantsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipants(val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MeetingParticipantsable))
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    return res
}
// GetParticipants gets the participants property value. The participants property
func (m *ItemOnlineMeetingsCreateOrGetPostRequestBody) GetParticipants()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MeetingParticipantsable) {
    return m.participants
}
// GetStartDateTime gets the startDateTime property value. The startDateTime property
func (m *ItemOnlineMeetingsCreateOrGetPostRequestBody) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetSubject gets the subject property value. The subject property
func (m *ItemOnlineMeetingsCreateOrGetPostRequestBody) GetSubject()(*string) {
    return m.subject
}
// Serialize serializes information the current object
func (m *ItemOnlineMeetingsCreateOrGetPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("chatInfo", m.GetChatInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("participants", m.GetParticipants())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subject", m.GetSubject())
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
func (m *ItemOnlineMeetingsCreateOrGetPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChatInfo sets the chatInfo property value. The chatInfo property
func (m *ItemOnlineMeetingsCreateOrGetPostRequestBody) SetChatInfo(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ChatInfoable)() {
    m.chatInfo = value
}
// SetEndDateTime sets the endDateTime property value. The endDateTime property
func (m *ItemOnlineMeetingsCreateOrGetPostRequestBody) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetExternalId sets the externalId property value. The externalId property
func (m *ItemOnlineMeetingsCreateOrGetPostRequestBody) SetExternalId(value *string)() {
    m.externalId = value
}
// SetParticipants sets the participants property value. The participants property
func (m *ItemOnlineMeetingsCreateOrGetPostRequestBody) SetParticipants(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MeetingParticipantsable)() {
    m.participants = value
}
// SetStartDateTime sets the startDateTime property value. The startDateTime property
func (m *ItemOnlineMeetingsCreateOrGetPostRequestBody) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetSubject sets the subject property value. The subject property
func (m *ItemOnlineMeetingsCreateOrGetPostRequestBody) SetSubject(value *string)() {
    m.subject = value
}
// ItemOnlineMeetingsCreateOrGetPostRequestBodyable 
type ItemOnlineMeetingsCreateOrGetPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChatInfo()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ChatInfoable)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExternalId()(*string)
    GetParticipants()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MeetingParticipantsable)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSubject()(*string)
    SetChatInfo(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ChatInfoable)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExternalId(value *string)()
    SetParticipants(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MeetingParticipantsable)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSubject(value *string)()
}
