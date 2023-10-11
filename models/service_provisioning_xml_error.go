package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceProvisioningXmlError 
type ServiceProvisioningXmlError struct {
    ServiceProvisioningError
    // The errorDetail property
    errorDetail *string
}
// NewServiceProvisioningXmlError instantiates a new serviceProvisioningXmlError and sets the default values.
func NewServiceProvisioningXmlError()(*ServiceProvisioningXmlError) {
    m := &ServiceProvisioningXmlError{
        ServiceProvisioningError: *NewServiceProvisioningError(),
    }
    odataTypeValue := "#microsoft.graph.serviceProvisioningXmlError"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateServiceProvisioningXmlErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceProvisioningXmlErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceProvisioningXmlError(), nil
}
// GetErrorDetail gets the errorDetail property value. The errorDetail property
func (m *ServiceProvisioningXmlError) GetErrorDetail()(*string) {
    return m.errorDetail
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceProvisioningXmlError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ServiceProvisioningError.GetFieldDeserializers()
    res["errorDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDetail(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ServiceProvisioningXmlError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ServiceProvisioningError.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("errorDetail", m.GetErrorDetail())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetErrorDetail sets the errorDetail property value. The errorDetail property
func (m *ServiceProvisioningXmlError) SetErrorDetail(value *string)() {
    m.errorDetail = value
}
// ServiceProvisioningXmlErrorable 
type ServiceProvisioningXmlErrorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ServiceProvisioningErrorable
    GetErrorDetail()(*string)
    SetErrorDetail(value *string)()
}
