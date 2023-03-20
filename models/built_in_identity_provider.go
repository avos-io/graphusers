package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BuiltInIdentityProvider 
type BuiltInIdentityProvider struct {
    IdentityProviderBase
    // The identity provider type. For a B2B scenario, possible values: AADSignup, MicrosoftAccount, EmailOTP. Required.
    identityProviderType *string
}
// NewBuiltInIdentityProvider instantiates a new BuiltInIdentityProvider and sets the default values.
func NewBuiltInIdentityProvider()(*BuiltInIdentityProvider) {
    m := &BuiltInIdentityProvider{
        IdentityProviderBase: *NewIdentityProviderBase(),
    }
    odataTypeValue := "#microsoft.graph.builtInIdentityProvider"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateBuiltInIdentityProviderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBuiltInIdentityProviderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBuiltInIdentityProvider(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BuiltInIdentityProvider) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityProviderBase.GetFieldDeserializers()
    res["identityProviderType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityProviderType(val)
        }
        return nil
    }
    return res
}
// GetIdentityProviderType gets the identityProviderType property value. The identity provider type. For a B2B scenario, possible values: AADSignup, MicrosoftAccount, EmailOTP. Required.
func (m *BuiltInIdentityProvider) GetIdentityProviderType()(*string) {
    return m.identityProviderType
}
// Serialize serializes information the current object
func (m *BuiltInIdentityProvider) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityProviderBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("identityProviderType", m.GetIdentityProviderType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIdentityProviderType sets the identityProviderType property value. The identity provider type. For a B2B scenario, possible values: AADSignup, MicrosoftAccount, EmailOTP. Required.
func (m *BuiltInIdentityProvider) SetIdentityProviderType(value *string)() {
    m.identityProviderType = value
}
// BuiltInIdentityProviderable 
type BuiltInIdentityProviderable interface {
    IdentityProviderBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIdentityProviderType()(*string)
    SetIdentityProviderType(value *string)()
}
