package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftManagedTrainingSetting 
type MicrosoftManagedTrainingSetting struct {
    TrainingSetting
    // The completionDateTime property
    completionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The trainingCompletionDuration property
    trainingCompletionDuration *TrainingCompletionDuration
}
// NewMicrosoftManagedTrainingSetting instantiates a new microsoftManagedTrainingSetting and sets the default values.
func NewMicrosoftManagedTrainingSetting()(*MicrosoftManagedTrainingSetting) {
    m := &MicrosoftManagedTrainingSetting{
        TrainingSetting: *NewTrainingSetting(),
    }
    odataTypeValue := "#microsoft.graph.microsoftManagedTrainingSetting"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMicrosoftManagedTrainingSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftManagedTrainingSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftManagedTrainingSetting(), nil
}
// GetCompletionDateTime gets the completionDateTime property value. The completionDateTime property
func (m *MicrosoftManagedTrainingSetting) GetCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completionDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftManagedTrainingSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TrainingSetting.GetFieldDeserializers()
    res["completionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletionDateTime(val)
        }
        return nil
    }
    res["trainingCompletionDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTrainingCompletionDuration)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrainingCompletionDuration(val.(*TrainingCompletionDuration))
        }
        return nil
    }
    return res
}
// GetTrainingCompletionDuration gets the trainingCompletionDuration property value. The trainingCompletionDuration property
func (m *MicrosoftManagedTrainingSetting) GetTrainingCompletionDuration()(*TrainingCompletionDuration) {
    return m.trainingCompletionDuration
}
// Serialize serializes information the current object
func (m *MicrosoftManagedTrainingSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TrainingSetting.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("completionDateTime", m.GetCompletionDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetTrainingCompletionDuration() != nil {
        cast := (*m.GetTrainingCompletionDuration()).String()
        err = writer.WriteStringValue("trainingCompletionDuration", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompletionDateTime sets the completionDateTime property value. The completionDateTime property
func (m *MicrosoftManagedTrainingSetting) SetCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completionDateTime = value
}
// SetTrainingCompletionDuration sets the trainingCompletionDuration property value. The trainingCompletionDuration property
func (m *MicrosoftManagedTrainingSetting) SetTrainingCompletionDuration(value *TrainingCompletionDuration)() {
    m.trainingCompletionDuration = value
}
// MicrosoftManagedTrainingSettingable 
type MicrosoftManagedTrainingSettingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TrainingSettingable
    GetCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTrainingCompletionDuration()(*TrainingCompletionDuration)
    SetCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTrainingCompletionDuration(value *TrainingCompletionDuration)()
}
