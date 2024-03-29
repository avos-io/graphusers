package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PositiveReinforcementNotification 
type PositiveReinforcementNotification struct {
    BaseEndUserNotification
    // The deliveryPreference property
    deliveryPreference *NotificationDeliveryPreference
}
// NewPositiveReinforcementNotification instantiates a new positiveReinforcementNotification and sets the default values.
func NewPositiveReinforcementNotification()(*PositiveReinforcementNotification) {
    m := &PositiveReinforcementNotification{
        BaseEndUserNotification: *NewBaseEndUserNotification(),
    }
    odataTypeValue := "#microsoft.graph.positiveReinforcementNotification"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePositiveReinforcementNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePositiveReinforcementNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPositiveReinforcementNotification(), nil
}
// GetDeliveryPreference gets the deliveryPreference property value. The deliveryPreference property
func (m *PositiveReinforcementNotification) GetDeliveryPreference()(*NotificationDeliveryPreference) {
    return m.deliveryPreference
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PositiveReinforcementNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseEndUserNotification.GetFieldDeserializers()
    res["deliveryPreference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNotificationDeliveryPreference)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveryPreference(val.(*NotificationDeliveryPreference))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PositiveReinforcementNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseEndUserNotification.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeliveryPreference() != nil {
        cast := (*m.GetDeliveryPreference()).String()
        err = writer.WriteStringValue("deliveryPreference", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeliveryPreference sets the deliveryPreference property value. The deliveryPreference property
func (m *PositiveReinforcementNotification) SetDeliveryPreference(value *NotificationDeliveryPreference)() {
    m.deliveryPreference = value
}
// PositiveReinforcementNotificationable 
type PositiveReinforcementNotificationable interface {
    BaseEndUserNotificationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeliveryPreference()(*NotificationDeliveryPreference)
    SetDeliveryPreference(value *NotificationDeliveryPreference)()
}
