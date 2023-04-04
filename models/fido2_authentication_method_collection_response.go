package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Fido2AuthenticationMethodCollectionResponse 
type Fido2AuthenticationMethodCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []Fido2AuthenticationMethodable
}
// NewFido2AuthenticationMethodCollectionResponse instantiates a new Fido2AuthenticationMethodCollectionResponse and sets the default values.
func NewFido2AuthenticationMethodCollectionResponse()(*Fido2AuthenticationMethodCollectionResponse) {
    m := &Fido2AuthenticationMethodCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateFido2AuthenticationMethodCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFido2AuthenticationMethodCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFido2AuthenticationMethodCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Fido2AuthenticationMethodCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFido2AuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Fido2AuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(Fido2AuthenticationMethodable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *Fido2AuthenticationMethodCollectionResponse) GetValue()([]Fido2AuthenticationMethodable) {
    return m.value
}
// Serialize serializes information the current object
func (m *Fido2AuthenticationMethodCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *Fido2AuthenticationMethodCollectionResponse) SetValue(value []Fido2AuthenticationMethodable)() {
    m.value = value
}
// Fido2AuthenticationMethodCollectionResponseable 
type Fido2AuthenticationMethodCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]Fido2AuthenticationMethodable)
    SetValue(value []Fido2AuthenticationMethodable)()
}
