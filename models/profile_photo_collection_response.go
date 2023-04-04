package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProfilePhotoCollectionResponse 
type ProfilePhotoCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []ProfilePhotoable
}
// NewProfilePhotoCollectionResponse instantiates a new ProfilePhotoCollectionResponse and sets the default values.
func NewProfilePhotoCollectionResponse()(*ProfilePhotoCollectionResponse) {
    m := &ProfilePhotoCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateProfilePhotoCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProfilePhotoCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProfilePhotoCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProfilePhotoCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProfilePhotoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProfilePhotoable, len(val))
            for i, v := range val {
                res[i] = v.(ProfilePhotoable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ProfilePhotoCollectionResponse) GetValue()([]ProfilePhotoable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ProfilePhotoCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ProfilePhotoCollectionResponse) SetValue(value []ProfilePhotoable)() {
    m.value = value
}
// ProfilePhotoCollectionResponseable 
type ProfilePhotoCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]ProfilePhotoable)
    SetValue(value []ProfilePhotoable)()
}
