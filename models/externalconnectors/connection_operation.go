package externalconnectors

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConnectionOperation 
type ConnectionOperation struct {
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Entity
    // If status is failed, provides more information about the error that caused the failure.
    error i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.PublicErrorable
    // Indicates the status of the asynchronous operation. Possible values are: unspecified, inprogress, completed, failed, unknownFutureValue.
    status *ConnectionOperationStatus
}
// NewConnectionOperation instantiates a new ConnectionOperation and sets the default values.
func NewConnectionOperation()(*ConnectionOperation) {
    m := &ConnectionOperation{
        Entity: *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.NewEntity(),
    }
    return m
}
// CreateConnectionOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConnectionOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConnectionOperation(), nil
}
// GetError gets the error property value. If status is failed, provides more information about the error that caused the failure.
func (m *ConnectionOperation) GetError()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.PublicErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConnectionOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectionOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ConnectionOperationStatus))
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. Indicates the status of the asynchronous operation. Possible values are: unspecified, inprogress, completed, failed, unknownFutureValue.
func (m *ConnectionOperation) GetStatus()(*ConnectionOperationStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *ConnectionOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetError sets the error property value. If status is failed, provides more information about the error that caused the failure.
func (m *ConnectionOperation) SetError(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.PublicErrorable)() {
    m.error = value
}
// SetStatus sets the status property value. Indicates the status of the asynchronous operation. Possible values are: unspecified, inprogress, completed, failed, unknownFutureValue.
func (m *ConnectionOperation) SetStatus(value *ConnectionOperationStatus)() {
    m.status = value
}
// ConnectionOperationable 
type ConnectionOperationable interface {
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetError()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.PublicErrorable)
    GetStatus()(*ConnectionOperationStatus)
    SetError(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.PublicErrorable)()
    SetStatus(value *ConnectionOperationStatus)()
}
