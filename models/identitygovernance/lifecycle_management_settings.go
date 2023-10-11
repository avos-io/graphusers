package identitygovernance

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LifecycleManagementSettings 
type LifecycleManagementSettings struct {
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Entity
    // The emailSettings property
    emailSettings i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.EmailSettingsable
    // The interval in hours at which all workflows running in the tenant should be scheduled for execution. This interval has a minimum value of 1 and a maximum value of 24. The default value is 3 hours.
    workflowScheduleIntervalInHours *int32
}
// NewLifecycleManagementSettings instantiates a new lifecycleManagementSettings and sets the default values.
func NewLifecycleManagementSettings()(*LifecycleManagementSettings) {
    m := &LifecycleManagementSettings{
        Entity: *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.NewEntity(),
    }
    return m
}
// CreateLifecycleManagementSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLifecycleManagementSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLifecycleManagementSettings(), nil
}
// GetEmailSettings gets the emailSettings property value. The emailSettings property
func (m *LifecycleManagementSettings) GetEmailSettings()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.EmailSettingsable) {
    return m.emailSettings
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LifecycleManagementSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["emailSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateEmailSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailSettings(val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.EmailSettingsable))
        }
        return nil
    }
    res["workflowScheduleIntervalInHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkflowScheduleIntervalInHours(val)
        }
        return nil
    }
    return res
}
// GetWorkflowScheduleIntervalInHours gets the workflowScheduleIntervalInHours property value. The interval in hours at which all workflows running in the tenant should be scheduled for execution. This interval has a minimum value of 1 and a maximum value of 24. The default value is 3 hours.
func (m *LifecycleManagementSettings) GetWorkflowScheduleIntervalInHours()(*int32) {
    return m.workflowScheduleIntervalInHours
}
// Serialize serializes information the current object
func (m *LifecycleManagementSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("emailSettings", m.GetEmailSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workflowScheduleIntervalInHours", m.GetWorkflowScheduleIntervalInHours())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEmailSettings sets the emailSettings property value. The emailSettings property
func (m *LifecycleManagementSettings) SetEmailSettings(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.EmailSettingsable)() {
    m.emailSettings = value
}
// SetWorkflowScheduleIntervalInHours sets the workflowScheduleIntervalInHours property value. The interval in hours at which all workflows running in the tenant should be scheduled for execution. This interval has a minimum value of 1 and a maximum value of 24. The default value is 3 hours.
func (m *LifecycleManagementSettings) SetWorkflowScheduleIntervalInHours(value *int32)() {
    m.workflowScheduleIntervalInHours = value
}
// LifecycleManagementSettingsable 
type LifecycleManagementSettingsable interface {
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmailSettings()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.EmailSettingsable)
    GetWorkflowScheduleIntervalInHours()(*int32)
    SetEmailSettings(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.EmailSettingsable)()
    SetWorkflowScheduleIntervalInHours(value *int32)()
}
