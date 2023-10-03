package identitygovernance

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomTaskExtensionCallbackConfiguration 
type CustomTaskExtensionCallbackConfiguration struct {
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CustomExtensionCallbackConfiguration
    // The authorizedApps property
    authorizedApps []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Applicationable
}
// NewCustomTaskExtensionCallbackConfiguration instantiates a new customTaskExtensionCallbackConfiguration and sets the default values.
func NewCustomTaskExtensionCallbackConfiguration()(*CustomTaskExtensionCallbackConfiguration) {
    m := &CustomTaskExtensionCallbackConfiguration{
        CustomExtensionCallbackConfiguration: *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.NewCustomExtensionCallbackConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.customTaskExtensionCallbackConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCustomTaskExtensionCallbackConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomTaskExtensionCallbackConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomTaskExtensionCallbackConfiguration(), nil
}
// GetAuthorizedApps gets the authorizedApps property value. The authorizedApps property
func (m *CustomTaskExtensionCallbackConfiguration) GetAuthorizedApps()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Applicationable) {
    return m.authorizedApps
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomTaskExtensionCallbackConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomExtensionCallbackConfiguration.GetFieldDeserializers()
    res["authorizedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Applicationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Applicationable)
                }
            }
            m.SetAuthorizedApps(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CustomTaskExtensionCallbackConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomExtensionCallbackConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthorizedApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthorizedApps()))
        for i, v := range m.GetAuthorizedApps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("authorizedApps", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthorizedApps sets the authorizedApps property value. The authorizedApps property
func (m *CustomTaskExtensionCallbackConfiguration) SetAuthorizedApps(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Applicationable)() {
    m.authorizedApps = value
}
// CustomTaskExtensionCallbackConfigurationable 
type CustomTaskExtensionCallbackConfigurationable interface {
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CustomExtensionCallbackConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorizedApps()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Applicationable)
    SetAuthorizedApps(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Applicationable)()
}
