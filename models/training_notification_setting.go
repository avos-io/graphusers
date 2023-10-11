package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TrainingNotificationSetting 
type TrainingNotificationSetting struct {
    EndUserNotificationSetting
    // The trainingAssignment property
    trainingAssignment BaseEndUserNotificationable
    // The trainingReminder property
    trainingReminder TrainingReminderNotificationable
}
// NewTrainingNotificationSetting instantiates a new trainingNotificationSetting and sets the default values.
func NewTrainingNotificationSetting()(*TrainingNotificationSetting) {
    m := &TrainingNotificationSetting{
        EndUserNotificationSetting: *NewEndUserNotificationSetting(),
    }
    odataTypeValue := "#microsoft.graph.trainingNotificationSetting"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTrainingNotificationSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTrainingNotificationSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTrainingNotificationSetting(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TrainingNotificationSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EndUserNotificationSetting.GetFieldDeserializers()
    res["trainingAssignment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBaseEndUserNotificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrainingAssignment(val.(BaseEndUserNotificationable))
        }
        return nil
    }
    res["trainingReminder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTrainingReminderNotificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrainingReminder(val.(TrainingReminderNotificationable))
        }
        return nil
    }
    return res
}
// GetTrainingAssignment gets the trainingAssignment property value. The trainingAssignment property
func (m *TrainingNotificationSetting) GetTrainingAssignment()(BaseEndUserNotificationable) {
    return m.trainingAssignment
}
// GetTrainingReminder gets the trainingReminder property value. The trainingReminder property
func (m *TrainingNotificationSetting) GetTrainingReminder()(TrainingReminderNotificationable) {
    return m.trainingReminder
}
// Serialize serializes information the current object
func (m *TrainingNotificationSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EndUserNotificationSetting.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("trainingAssignment", m.GetTrainingAssignment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("trainingReminder", m.GetTrainingReminder())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTrainingAssignment sets the trainingAssignment property value. The trainingAssignment property
func (m *TrainingNotificationSetting) SetTrainingAssignment(value BaseEndUserNotificationable)() {
    m.trainingAssignment = value
}
// SetTrainingReminder sets the trainingReminder property value. The trainingReminder property
func (m *TrainingNotificationSetting) SetTrainingReminder(value TrainingReminderNotificationable)() {
    m.trainingReminder = value
}
// TrainingNotificationSettingable 
type TrainingNotificationSettingable interface {
    EndUserNotificationSettingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTrainingAssignment()(BaseEndUserNotificationable)
    GetTrainingReminder()(TrainingReminderNotificationable)
    SetTrainingAssignment(value BaseEndUserNotificationable)()
    SetTrainingReminder(value TrainingReminderNotificationable)()
}
