package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosHomeScreenAppCollectionResponse 
type IosHomeScreenAppCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []IosHomeScreenAppable
}
// NewIosHomeScreenAppCollectionResponse instantiates a new IosHomeScreenAppCollectionResponse and sets the default values.
func NewIosHomeScreenAppCollectionResponse()(*IosHomeScreenAppCollectionResponse) {
    m := &IosHomeScreenAppCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateIosHomeScreenAppCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosHomeScreenAppCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosHomeScreenAppCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosHomeScreenAppCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosHomeScreenAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosHomeScreenAppable, len(val))
            for i, v := range val {
                res[i] = v.(IosHomeScreenAppable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *IosHomeScreenAppCollectionResponse) GetValue()([]IosHomeScreenAppable) {
    return m.value
}
// Serialize serializes information the current object
func (m *IosHomeScreenAppCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *IosHomeScreenAppCollectionResponse) SetValue(value []IosHomeScreenAppable)() {
    m.value = value
}
// IosHomeScreenAppCollectionResponseable 
type IosHomeScreenAppCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]IosHomeScreenAppable)
    SetValue(value []IosHomeScreenAppable)()
}
