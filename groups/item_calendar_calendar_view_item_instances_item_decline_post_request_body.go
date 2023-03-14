package groups

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBody 
type ItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Comment property
    comment *string
    // The ProposedNewTime property
    proposedNewTime i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TimeSlotable
    // The SendResponse property
    sendResponse *bool
}
// NewItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBody instantiates a new ItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBody and sets the default values.
func NewItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBody()(*ItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBody) {
    m := &ItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the comment property value. The Comment property
func (m *ItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["proposedNewTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateTimeSlotFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProposedNewTime(val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TimeSlotable))
        }
        return nil
    }
    res["sendResponse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSendResponse(val)
        }
        return nil
    }
    return res
}
// GetProposedNewTime gets the proposedNewTime property value. The ProposedNewTime property
func (m *ItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBody) GetProposedNewTime()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TimeSlotable) {
    return m.proposedNewTime
}
// GetSendResponse gets the sendResponse property value. The SendResponse property
func (m *ItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBody) GetSendResponse()(*bool) {
    return m.sendResponse
}
// Serialize serializes information the current object
func (m *ItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("proposedNewTime", m.GetProposedNewTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sendResponse", m.GetSendResponse())
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
func (m *ItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the comment property value. The Comment property
func (m *ItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBody) SetComment(value *string)() {
    m.comment = value
}
// SetProposedNewTime sets the proposedNewTime property value. The ProposedNewTime property
func (m *ItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBody) SetProposedNewTime(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TimeSlotable)() {
    m.proposedNewTime = value
}
// SetSendResponse sets the sendResponse property value. The SendResponse property
func (m *ItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBody) SetSendResponse(value *bool)() {
    m.sendResponse = value
}
// ItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBodyable 
type ItemCalendarCalendarViewItemInstancesItemDeclinePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComment()(*string)
    GetProposedNewTime()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TimeSlotable)
    GetSendResponse()(*bool)
    SetComment(value *string)()
    SetProposedNewTime(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TimeSlotable)()
    SetSendResponse(value *bool)()
}
