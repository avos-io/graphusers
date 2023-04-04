package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedAppStatus represents app protection and configuration status for the organization.
type ManagedAppStatus struct {
    Entity
    // Friendly name of the status report.
    displayName *string
    // Version of the entity.
    version *string
}
// NewManagedAppStatus instantiates a new managedAppStatus and sets the default values.
func NewManagedAppStatus()(*ManagedAppStatus) {
    m := &ManagedAppStatus{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedAppStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedAppStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.managedAppStatusRaw":
                        return NewManagedAppStatusRaw(), nil
                }
            }
        }
    }
    return NewManagedAppStatus(), nil
}
// GetDisplayName gets the displayName property value. Friendly name of the status report.
func (m *ManagedAppStatus) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAppStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetVersion gets the version property value. Version of the entity.
func (m *ManagedAppStatus) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *ManagedAppStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. Friendly name of the status report.
func (m *ManagedAppStatus) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetVersion sets the version property value. Version of the entity.
func (m *ManagedAppStatus) SetVersion(value *string)() {
    m.version = value
}
// ManagedAppStatusable 
type ManagedAppStatusable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetVersion()(*string)
    SetDisplayName(value *string)()
    SetVersion(value *string)()
}
