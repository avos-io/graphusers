package users

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarsItemEventsItemInstancesItemSnoozeReminderPostRequestBody 
type ItemCalendarsItemEventsItemInstancesItemSnoozeReminderPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The NewReminderTime property
    newReminderTime i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.DateTimeTimeZoneable
}
// NewItemCalendarsItemEventsItemInstancesItemSnoozeReminderPostRequestBody instantiates a new ItemCalendarsItemEventsItemInstancesItemSnoozeReminderPostRequestBody and sets the default values.
func NewItemCalendarsItemEventsItemInstancesItemSnoozeReminderPostRequestBody()(*ItemCalendarsItemEventsItemInstancesItemSnoozeReminderPostRequestBody) {
    m := &ItemCalendarsItemEventsItemInstancesItemSnoozeReminderPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCalendarsItemEventsItemInstancesItemSnoozeReminderPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarsItemEventsItemInstancesItemSnoozeReminderPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarsItemEventsItemInstancesItemSnoozeReminderPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCalendarsItemEventsItemInstancesItemSnoozeReminderPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCalendarsItemEventsItemInstancesItemSnoozeReminderPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["newReminderTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewReminderTime(val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.DateTimeTimeZoneable))
        }
        return nil
    }
    return res
}
// GetNewReminderTime gets the newReminderTime property value. The NewReminderTime property
func (m *ItemCalendarsItemEventsItemInstancesItemSnoozeReminderPostRequestBody) GetNewReminderTime()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.DateTimeTimeZoneable) {
    return m.newReminderTime
}
// Serialize serializes information the current object
func (m *ItemCalendarsItemEventsItemInstancesItemSnoozeReminderPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("newReminderTime", m.GetNewReminderTime())
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
func (m *ItemCalendarsItemEventsItemInstancesItemSnoozeReminderPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNewReminderTime sets the newReminderTime property value. The NewReminderTime property
func (m *ItemCalendarsItemEventsItemInstancesItemSnoozeReminderPostRequestBody) SetNewReminderTime(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.DateTimeTimeZoneable)() {
    m.newReminderTime = value
}
// ItemCalendarsItemEventsItemInstancesItemSnoozeReminderPostRequestBodyable 
type ItemCalendarsItemEventsItemInstancesItemSnoozeReminderPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNewReminderTime()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.DateTimeTimeZoneable)
    SetNewReminderTime(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.DateTimeTimeZoneable)()
}
