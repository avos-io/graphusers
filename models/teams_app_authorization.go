package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsAppAuthorization 
type TeamsAppAuthorization struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Set of permissions required by the teamsApp.
    requiredPermissionSet TeamsAppPermissionSetable
}
// NewTeamsAppAuthorization instantiates a new teamsAppAuthorization and sets the default values.
func NewTeamsAppAuthorization()(*TeamsAppAuthorization) {
    m := &TeamsAppAuthorization{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamsAppAuthorizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsAppAuthorizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsAppAuthorization(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamsAppAuthorization) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsAppAuthorization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["requiredPermissionSet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsAppPermissionSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredPermissionSet(val.(TeamsAppPermissionSetable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamsAppAuthorization) GetOdataType()(*string) {
    return m.odataType
}
// GetRequiredPermissionSet gets the requiredPermissionSet property value. Set of permissions required by the teamsApp.
func (m *TeamsAppAuthorization) GetRequiredPermissionSet()(TeamsAppPermissionSetable) {
    return m.requiredPermissionSet
}
// Serialize serializes information the current object
func (m *TeamsAppAuthorization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("requiredPermissionSet", m.GetRequiredPermissionSet())
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
func (m *TeamsAppAuthorization) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamsAppAuthorization) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRequiredPermissionSet sets the requiredPermissionSet property value. Set of permissions required by the teamsApp.
func (m *TeamsAppAuthorization) SetRequiredPermissionSet(value TeamsAppPermissionSetable)() {
    m.requiredPermissionSet = value
}
// TeamsAppAuthorizationable 
type TeamsAppAuthorizationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetRequiredPermissionSet()(TeamsAppPermissionSetable)
    SetOdataType(value *string)()
    SetRequiredPermissionSet(value TeamsAppPermissionSetable)()
}
