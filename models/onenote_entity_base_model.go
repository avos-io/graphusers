package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnenoteEntityBaseModel provides operations to manage the collection of user entities.
type OnenoteEntityBaseModel struct {
    Entity
    // The endpoint where you can get details about the page. Read-only.
    self *string
}
// NewOnenoteEntityBaseModel instantiates a new onenoteEntityBaseModel and sets the default values.
func NewOnenoteEntityBaseModel()(*OnenoteEntityBaseModel) {
    m := &OnenoteEntityBaseModel{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOnenoteEntityBaseModelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenoteEntityBaseModelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.notebook":
                        return NewNotebook(), nil
                    case "#microsoft.graph.onenoteEntityHierarchyModel":
                        return NewOnenoteEntityHierarchyModel(), nil
                    case "#microsoft.graph.onenoteEntitySchemaObjectModel":
                        return NewOnenoteEntitySchemaObjectModel(), nil
                    case "#microsoft.graph.onenotePage":
                        return NewOnenotePage(), nil
                    case "#microsoft.graph.onenoteResource":
                        return NewOnenoteResource(), nil
                    case "#microsoft.graph.onenoteSection":
                        return NewOnenoteSection(), nil
                    case "#microsoft.graph.sectionGroup":
                        return NewSectionGroup(), nil
                }
            }
        }
    }
    return NewOnenoteEntityBaseModel(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenoteEntityBaseModel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["self"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSelf)
    return res
}
// GetSelf gets the self property value. The endpoint where you can get details about the page. Read-only.
func (m *OnenoteEntityBaseModel) GetSelf()(*string) {
    return m.self
}
// Serialize serializes information the current object
func (m *OnenoteEntityBaseModel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("self", m.GetSelf())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSelf sets the self property value. The endpoint where you can get details about the page. Read-only.
func (m *OnenoteEntityBaseModel) SetSelf(value *string)() {
    m.self = value
}
