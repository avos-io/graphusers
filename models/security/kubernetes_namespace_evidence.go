package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// KubernetesNamespaceEvidence 
type KubernetesNamespaceEvidence struct {
    AlertEvidence
    // The namespace cluster.
    cluster KubernetesClusterEvidenceable
    // The labels for the Kubernetes pod.
    labels Dictionaryable
    // The namespace name.
    name *string
}
// NewKubernetesNamespaceEvidence instantiates a new kubernetesNamespaceEvidence and sets the default values.
func NewKubernetesNamespaceEvidence()(*KubernetesNamespaceEvidence) {
    m := &KubernetesNamespaceEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.kubernetesNamespaceEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateKubernetesNamespaceEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateKubernetesNamespaceEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKubernetesNamespaceEvidence(), nil
}
// GetCluster gets the cluster property value. The namespace cluster.
func (m *KubernetesNamespaceEvidence) GetCluster()(KubernetesClusterEvidenceable) {
    return m.cluster
}
// GetFieldDeserializers the deserialization information for the current model
func (m *KubernetesNamespaceEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["cluster"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateKubernetesClusterEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCluster(val.(KubernetesClusterEvidenceable))
        }
        return nil
    }
    res["labels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDictionaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabels(val.(Dictionaryable))
        }
        return nil
    }
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
    return res
}
// GetLabels gets the labels property value. The labels for the Kubernetes pod.
func (m *KubernetesNamespaceEvidence) GetLabels()(Dictionaryable) {
    return m.labels
}
// GetName gets the name property value. The namespace name.
func (m *KubernetesNamespaceEvidence) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *KubernetesNamespaceEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("cluster", m.GetCluster())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("labels", m.GetLabels())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCluster sets the cluster property value. The namespace cluster.
func (m *KubernetesNamespaceEvidence) SetCluster(value KubernetesClusterEvidenceable)() {
    m.cluster = value
}
// SetLabels sets the labels property value. The labels for the Kubernetes pod.
func (m *KubernetesNamespaceEvidence) SetLabels(value Dictionaryable)() {
    m.labels = value
}
// SetName sets the name property value. The namespace name.
func (m *KubernetesNamespaceEvidence) SetName(value *string)() {
    m.name = value
}
// KubernetesNamespaceEvidenceable 
type KubernetesNamespaceEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCluster()(KubernetesClusterEvidenceable)
    GetLabels()(Dictionaryable)
    GetName()(*string)
    SetCluster(value KubernetesClusterEvidenceable)()
    SetLabels(value Dictionaryable)()
    SetName(value *string)()
}
