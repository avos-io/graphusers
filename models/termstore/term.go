package termstore

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Term 
type Term struct {
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Entity
    // Children of current term.
    children []Termable
    // Date and time of term creation. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Description about term that is dependent on the languageTag.
    descriptions []LocalizedDescriptionable
    // Label metadata for a term.
    labels []LocalizedLabelable
    // Last date and time of term modification. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Collection of properties on the term.
    properties []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.KeyValueable
    // To indicate which terms are related to the current term as either pinned or reused.
    relations []Relationable
    // The [set] in which the term is created.
    set Setable
}
// NewTerm instantiates a new term and sets the default values.
func NewTerm()(*Term) {
    m := &Term{
        Entity: *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.NewEntity(),
    }
    return m
}
// CreateTermFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTermFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTerm(), nil
}
// GetChildren gets the children property value. Children of current term.
func (m *Term) GetChildren()([]Termable) {
    return m.children
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time of term creation. Read-only.
func (m *Term) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescriptions gets the descriptions property value. Description about term that is dependent on the languageTag.
func (m *Term) GetDescriptions()([]LocalizedDescriptionable) {
    return m.descriptions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Term) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["children"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTermFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Termable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Termable)
                }
            }
            m.SetChildren(res)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["descriptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLocalizedDescriptionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LocalizedDescriptionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LocalizedDescriptionable)
                }
            }
            m.SetDescriptions(res)
        }
        return nil
    }
    res["labels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLocalizedLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LocalizedLabelable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LocalizedLabelable)
                }
            }
            m.SetLabels(res)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["properties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateKeyValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.KeyValueable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.KeyValueable)
                }
            }
            m.SetProperties(res)
        }
        return nil
    }
    res["relations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRelationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Relationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Relationable)
                }
            }
            m.SetRelations(res)
        }
        return nil
    }
    res["set"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSet(val.(Setable))
        }
        return nil
    }
    return res
}
// GetLabels gets the labels property value. Label metadata for a term.
func (m *Term) GetLabels()([]LocalizedLabelable) {
    return m.labels
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last date and time of term modification. Read-only.
func (m *Term) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetProperties gets the properties property value. Collection of properties on the term.
func (m *Term) GetProperties()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.KeyValueable) {
    return m.properties
}
// GetRelations gets the relations property value. To indicate which terms are related to the current term as either pinned or reused.
func (m *Term) GetRelations()([]Relationable) {
    return m.relations
}
// GetSet gets the set property value. The [set] in which the term is created.
func (m *Term) GetSet()(Setable) {
    return m.set
}
// Serialize serializes information the current object
func (m *Term) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChildren() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChildren()))
        for i, v := range m.GetChildren() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("children", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetDescriptions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDescriptions()))
        for i, v := range m.GetDescriptions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("descriptions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLabels()))
        for i, v := range m.GetLabels() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("labels", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProperties()))
        for i, v := range m.GetProperties() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("properties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRelations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRelations()))
        for i, v := range m.GetRelations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("relations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("set", m.GetSet())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChildren sets the children property value. Children of current term.
func (m *Term) SetChildren(value []Termable)() {
    m.children = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time of term creation. Read-only.
func (m *Term) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescriptions sets the descriptions property value. Description about term that is dependent on the languageTag.
func (m *Term) SetDescriptions(value []LocalizedDescriptionable)() {
    m.descriptions = value
}
// SetLabels sets the labels property value. Label metadata for a term.
func (m *Term) SetLabels(value []LocalizedLabelable)() {
    m.labels = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last date and time of term modification. Read-only.
func (m *Term) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetProperties sets the properties property value. Collection of properties on the term.
func (m *Term) SetProperties(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.KeyValueable)() {
    m.properties = value
}
// SetRelations sets the relations property value. To indicate which terms are related to the current term as either pinned or reused.
func (m *Term) SetRelations(value []Relationable)() {
    m.relations = value
}
// SetSet sets the set property value. The [set] in which the term is created.
func (m *Term) SetSet(value Setable)() {
    m.set = value
}
// Termable 
type Termable interface {
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChildren()([]Termable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescriptions()([]LocalizedDescriptionable)
    GetLabels()([]LocalizedLabelable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetProperties()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.KeyValueable)
    GetRelations()([]Relationable)
    GetSet()(Setable)
    SetChildren(value []Termable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescriptions(value []LocalizedDescriptionable)()
    SetLabels(value []LocalizedLabelable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetProperties(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.KeyValueable)()
    SetRelations(value []Relationable)()
    SetSet(value Setable)()
}
