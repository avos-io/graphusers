package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftCustomTrainingSetting 
type MicrosoftCustomTrainingSetting struct {
    TrainingSetting
    // The completionDateTime property
    completionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The trainingAssignmentMappings property
    trainingAssignmentMappings []MicrosoftTrainingAssignmentMappingable
    // The trainingCompletionDuration property
    trainingCompletionDuration *TrainingCompletionDuration
}
// NewMicrosoftCustomTrainingSetting instantiates a new microsoftCustomTrainingSetting and sets the default values.
func NewMicrosoftCustomTrainingSetting()(*MicrosoftCustomTrainingSetting) {
    m := &MicrosoftCustomTrainingSetting{
        TrainingSetting: *NewTrainingSetting(),
    }
    odataTypeValue := "#microsoft.graph.microsoftCustomTrainingSetting"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMicrosoftCustomTrainingSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftCustomTrainingSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftCustomTrainingSetting(), nil
}
// GetCompletionDateTime gets the completionDateTime property value. The completionDateTime property
func (m *MicrosoftCustomTrainingSetting) GetCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completionDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftCustomTrainingSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["trainingAssignmentMappings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMicrosoftTrainingAssignmentMappingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MicrosoftTrainingAssignmentMappingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MicrosoftTrainingAssignmentMappingable)
                }
            }
            m.SetTrainingAssignmentMappings(res)
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
// GetTrainingAssignmentMappings gets the trainingAssignmentMappings property value. The trainingAssignmentMappings property
func (m *MicrosoftCustomTrainingSetting) GetTrainingAssignmentMappings()([]MicrosoftTrainingAssignmentMappingable) {
    return m.trainingAssignmentMappings
}
// GetTrainingCompletionDuration gets the trainingCompletionDuration property value. The trainingCompletionDuration property
func (m *MicrosoftCustomTrainingSetting) GetTrainingCompletionDuration()(*TrainingCompletionDuration) {
    return m.trainingCompletionDuration
}
// Serialize serializes information the current object
func (m *MicrosoftCustomTrainingSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetTrainingAssignmentMappings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTrainingAssignmentMappings()))
        for i, v := range m.GetTrainingAssignmentMappings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("trainingAssignmentMappings", cast)
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
func (m *MicrosoftCustomTrainingSetting) SetCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completionDateTime = value
}
// SetTrainingAssignmentMappings sets the trainingAssignmentMappings property value. The trainingAssignmentMappings property
func (m *MicrosoftCustomTrainingSetting) SetTrainingAssignmentMappings(value []MicrosoftTrainingAssignmentMappingable)() {
    m.trainingAssignmentMappings = value
}
// SetTrainingCompletionDuration sets the trainingCompletionDuration property value. The trainingCompletionDuration property
func (m *MicrosoftCustomTrainingSetting) SetTrainingCompletionDuration(value *TrainingCompletionDuration)() {
    m.trainingCompletionDuration = value
}
// MicrosoftCustomTrainingSettingable 
type MicrosoftCustomTrainingSettingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TrainingSettingable
    GetCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTrainingAssignmentMappings()([]MicrosoftTrainingAssignmentMappingable)
    GetTrainingCompletionDuration()(*TrainingCompletionDuration)
    SetCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTrainingAssignmentMappings(value []MicrosoftTrainingAssignmentMappingable)()
    SetTrainingCompletionDuration(value *TrainingCompletionDuration)()
}
