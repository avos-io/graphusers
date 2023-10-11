package externalconnectors

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExternalActivityResult 
type ExternalActivityResult struct {
    ExternalActivity
    // Error information that explains the failure to process an external activity.
    error i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.PublicErrorable
}
// NewExternalActivityResult instantiates a new externalActivityResult and sets the default values.
func NewExternalActivityResult()(*ExternalActivityResult) {
    m := &ExternalActivityResult{
        ExternalActivity: *NewExternalActivity(),
    }
    return m
}
// CreateExternalActivityResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternalActivityResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalActivityResult(), nil
}
// GetError gets the error property value. Error information that explains the failure to process an external activity.
func (m *ExternalActivityResult) GetError()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.PublicErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalActivityResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ExternalActivity.GetFieldDeserializers()
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreatePublicErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.PublicErrorable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ExternalActivityResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ExternalActivity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetError sets the error property value. Error information that explains the failure to process an external activity.
func (m *ExternalActivityResult) SetError(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.PublicErrorable)() {
    m.error = value
}
// ExternalActivityResultable 
type ExternalActivityResultable interface {
    ExternalActivityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetError()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.PublicErrorable)
    SetError(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.PublicErrorable)()
}
