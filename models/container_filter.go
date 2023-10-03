package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContainerFilter 
type ContainerFilter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The includedContainers property
    includedContainers []string
    // The OdataType property
    odataType *string
}
// NewContainerFilter instantiates a new containerFilter and sets the default values.
func NewContainerFilter()(*ContainerFilter) {
    m := &ContainerFilter{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateContainerFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContainerFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContainerFilter(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContainerFilter) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContainerFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["includedContainers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetIncludedContainers(res)
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
    return res
}
// GetIncludedContainers gets the includedContainers property value. The includedContainers property
func (m *ContainerFilter) GetIncludedContainers()([]string) {
    return m.includedContainers
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ContainerFilter) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ContainerFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIncludedContainers() != nil {
        err := writer.WriteCollectionOfStringValues("includedContainers", m.GetIncludedContainers())
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContainerFilter) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIncludedContainers sets the includedContainers property value. The includedContainers property
func (m *ContainerFilter) SetIncludedContainers(value []string)() {
    m.includedContainers = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ContainerFilter) SetOdataType(value *string)() {
    m.odataType = value
}
// ContainerFilterable 
type ContainerFilterable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIncludedContainers()([]string)
    GetOdataType()(*string)
    SetIncludedContainers(value []string)()
    SetOdataType(value *string)()
}
