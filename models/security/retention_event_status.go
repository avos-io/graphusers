package security

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RetentionEventStatus 
type RetentionEventStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The error if the status isn't successful.
    error i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.PublicErrorable
    // The OdataType property
    odataType *string
    // The status of the distribution. The possible values are: pending, error, success, notAvaliable.
    status *EventStatusType
}
// NewRetentionEventStatus instantiates a new retentionEventStatus and sets the default values.
func NewRetentionEventStatus()(*RetentionEventStatus) {
    m := &RetentionEventStatus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRetentionEventStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRetentionEventStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRetentionEventStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RetentionEventStatus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetError gets the error property value. The error if the status isn't successful.
func (m *RetentionEventStatus) GetError()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.PublicErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RetentionEventStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEventStatusType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*EventStatusType))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RetentionEventStatus) GetOdataType()(*string) {
    return m.odataType
}
// GetStatus gets the status property value. The status of the distribution. The possible values are: pending, error, success, notAvaliable.
func (m *RetentionEventStatus) GetStatus()(*EventStatusType) {
    return m.status
}
// Serialize serializes information the current object
func (m *RetentionEventStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *RetentionEventStatus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetError sets the error property value. The error if the status isn't successful.
func (m *RetentionEventStatus) SetError(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.PublicErrorable)() {
    m.error = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RetentionEventStatus) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStatus sets the status property value. The status of the distribution. The possible values are: pending, error, success, notAvaliable.
func (m *RetentionEventStatus) SetStatus(value *EventStatusType)() {
    m.status = value
}
// RetentionEventStatusable 
type RetentionEventStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetError()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.PublicErrorable)
    GetOdataType()(*string)
    GetStatus()(*EventStatusType)
    SetError(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.PublicErrorable)()
    SetOdataType(value *string)()
    SetStatus(value *EventStatusType)()
}
