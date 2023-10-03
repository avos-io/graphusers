package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomTaskExtension 
type CustomTaskExtension struct {
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CustomCalloutExtension
    // The callback configuration for a custom task extension.
    callbackConfiguration i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CustomExtensionCallbackConfigurationable
    // The unique identifier of the Azure AD user that created the custom task extension.Supports $filter(eq, ne) and $expand.
    createdBy i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable
    // When the custom task extension was created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The unique identifier of the Azure AD user that modified the custom task extension last.Supports $filter(eq, ne) and $expand.
    lastModifiedBy i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable
    // When the custom extension was last modified.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewCustomTaskExtension instantiates a new customTaskExtension and sets the default values.
func NewCustomTaskExtension()(*CustomTaskExtension) {
    m := &CustomTaskExtension{
        CustomCalloutExtension: *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.NewCustomCalloutExtension(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.customTaskExtension"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCustomTaskExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomTaskExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomTaskExtension(), nil
}
// GetCallbackConfiguration gets the callbackConfiguration property value. The callback configuration for a custom task extension.
func (m *CustomTaskExtension) GetCallbackConfiguration()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CustomExtensionCallbackConfigurationable) {
    return m.callbackConfiguration
}
// GetCreatedBy gets the createdBy property value. The unique identifier of the Azure AD user that created the custom task extension.Supports $filter(eq, ne) and $expand.
func (m *CustomTaskExtension) GetCreatedBy()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. When the custom task extension was created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *CustomTaskExtension) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomTaskExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomCalloutExtension.GetFieldDeserializers()
    res["callbackConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateCustomExtensionCallbackConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallbackConfiguration(val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CustomExtensionCallbackConfigurationable))
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The unique identifier of the Azure AD user that modified the custom task extension last.Supports $filter(eq, ne) and $expand.
func (m *CustomTaskExtension) GetLastModifiedBy()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. When the custom extension was last modified.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *CustomTaskExtension) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// Serialize serializes information the current object
func (m *CustomTaskExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomCalloutExtension.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("callbackConfiguration", m.GetCallbackConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCallbackConfiguration sets the callbackConfiguration property value. The callback configuration for a custom task extension.
func (m *CustomTaskExtension) SetCallbackConfiguration(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CustomExtensionCallbackConfigurationable)() {
    m.callbackConfiguration = value
}
// SetCreatedBy sets the createdBy property value. The unique identifier of the Azure AD user that created the custom task extension.Supports $filter(eq, ne) and $expand.
func (m *CustomTaskExtension) SetCreatedBy(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. When the custom task extension was created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *CustomTaskExtension) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The unique identifier of the Azure AD user that modified the custom task extension last.Supports $filter(eq, ne) and $expand.
func (m *CustomTaskExtension) SetLastModifiedBy(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. When the custom extension was last modified.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *CustomTaskExtension) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// CustomTaskExtensionable 
type CustomTaskExtensionable interface {
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CustomCalloutExtensionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCallbackConfiguration()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CustomExtensionCallbackConfigurationable)
    GetCreatedBy()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedBy()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCallbackConfiguration(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CustomExtensionCallbackConfigurationable)()
    SetCreatedBy(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedBy(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
