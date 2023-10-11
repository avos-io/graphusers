package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// KubernetesServiceAccountEvidence 
type KubernetesServiceAccountEvidence struct {
    AlertEvidence
    // The service account name.
    name *string
    // The service account namespace.
    namespace KubernetesNamespaceEvidenceable
}
// NewKubernetesServiceAccountEvidence instantiates a new kubernetesServiceAccountEvidence and sets the default values.
func NewKubernetesServiceAccountEvidence()(*KubernetesServiceAccountEvidence) {
    m := &KubernetesServiceAccountEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.kubernetesServiceAccountEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateKubernetesServiceAccountEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateKubernetesServiceAccountEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKubernetesServiceAccountEvidence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *KubernetesServiceAccountEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["namespace"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateKubernetesNamespaceEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNamespace(val.(KubernetesNamespaceEvidenceable))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The service account name.
func (m *KubernetesServiceAccountEvidence) GetName()(*string) {
    return m.name
}
// GetNamespace gets the namespace property value. The service account namespace.
func (m *KubernetesServiceAccountEvidence) GetNamespace()(KubernetesNamespaceEvidenceable) {
    return m.namespace
}
// Serialize serializes information the current object
func (m *KubernetesServiceAccountEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("namespace", m.GetNamespace())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. The service account name.
func (m *KubernetesServiceAccountEvidence) SetName(value *string)() {
    m.name = value
}
// SetNamespace sets the namespace property value. The service account namespace.
func (m *KubernetesServiceAccountEvidence) SetNamespace(value KubernetesNamespaceEvidenceable)() {
    m.namespace = value
}
// KubernetesServiceAccountEvidenceable 
type KubernetesServiceAccountEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetNamespace()(KubernetesNamespaceEvidenceable)
    SetName(value *string)()
    SetNamespace(value KubernetesNamespaceEvidenceable)()
}
