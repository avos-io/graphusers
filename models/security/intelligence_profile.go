package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IntelligenceProfile 
type IntelligenceProfile struct {
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Entity
    // A list of commonly-known aliases for the threat intelligence included in the intelligenceProfile.
    aliases []string
    // The country/region of origin for the given actor or threat associated with this intelligenceProfile.
    countriesOrRegionsOfOrigin []IntelligenceProfileCountryOrRegionOfOriginable
    // The description property
    description FormattedContentable
    // The date and time when this intelligenceProfile was first active. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    firstActiveDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Includes an assemblage of high-fidelity network indicators of compromise.
    indicators []IntelligenceProfileIndicatorable
    // The kind property
    kind *IntelligenceProfileKind
    // The summary property
    summary FormattedContentable
    // Known targets related to this intelligenceProfile.
    targets []string
    // The title of this intelligenceProfile.
    title *string
    // Formatted information featuring a description of the distinctive tactics, techniques, and procedures (TTP) of the group, followed by a list of all known custom, commodity, and publicly available implants used by the group.
    tradecraft FormattedContentable
}
// NewIntelligenceProfile instantiates a new intelligenceProfile and sets the default values.
func NewIntelligenceProfile()(*IntelligenceProfile) {
    m := &IntelligenceProfile{
        Entity: *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.NewEntity(),
    }
    return m
}
// CreateIntelligenceProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIntelligenceProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntelligenceProfile(), nil
}
// GetAliases gets the aliases property value. A list of commonly-known aliases for the threat intelligence included in the intelligenceProfile.
func (m *IntelligenceProfile) GetAliases()([]string) {
    return m.aliases
}
// GetCountriesOrRegionsOfOrigin gets the countriesOrRegionsOfOrigin property value. The country/region of origin for the given actor or threat associated with this intelligenceProfile.
func (m *IntelligenceProfile) GetCountriesOrRegionsOfOrigin()([]IntelligenceProfileCountryOrRegionOfOriginable) {
    return m.countriesOrRegionsOfOrigin
}
// GetDescription gets the description property value. The description property
func (m *IntelligenceProfile) GetDescription()(FormattedContentable) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IntelligenceProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["aliases"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAliases(res)
        }
        return nil
    }
    res["countriesOrRegionsOfOrigin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIntelligenceProfileCountryOrRegionOfOriginFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IntelligenceProfileCountryOrRegionOfOriginable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IntelligenceProfileCountryOrRegionOfOriginable)
                }
            }
            m.SetCountriesOrRegionsOfOrigin(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFormattedContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val.(FormattedContentable))
        }
        return nil
    }
    res["firstActiveDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstActiveDateTime(val)
        }
        return nil
    }
    res["indicators"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIntelligenceProfileIndicatorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IntelligenceProfileIndicatorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IntelligenceProfileIndicatorable)
                }
            }
            m.SetIndicators(res)
        }
        return nil
    }
    res["kind"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIntelligenceProfileKind)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKind(val.(*IntelligenceProfileKind))
        }
        return nil
    }
    res["summary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFormattedContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSummary(val.(FormattedContentable))
        }
        return nil
    }
    res["targets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetTargets(res)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["tradecraft"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFormattedContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTradecraft(val.(FormattedContentable))
        }
        return nil
    }
    return res
}
// GetFirstActiveDateTime gets the firstActiveDateTime property value. The date and time when this intelligenceProfile was first active. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *IntelligenceProfile) GetFirstActiveDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.firstActiveDateTime
}
// GetIndicators gets the indicators property value. Includes an assemblage of high-fidelity network indicators of compromise.
func (m *IntelligenceProfile) GetIndicators()([]IntelligenceProfileIndicatorable) {
    return m.indicators
}
// GetKind gets the kind property value. The kind property
func (m *IntelligenceProfile) GetKind()(*IntelligenceProfileKind) {
    return m.kind
}
// GetSummary gets the summary property value. The summary property
func (m *IntelligenceProfile) GetSummary()(FormattedContentable) {
    return m.summary
}
// GetTargets gets the targets property value. Known targets related to this intelligenceProfile.
func (m *IntelligenceProfile) GetTargets()([]string) {
    return m.targets
}
// GetTitle gets the title property value. The title of this intelligenceProfile.
func (m *IntelligenceProfile) GetTitle()(*string) {
    return m.title
}
// GetTradecraft gets the tradecraft property value. Formatted information featuring a description of the distinctive tactics, techniques, and procedures (TTP) of the group, followed by a list of all known custom, commodity, and publicly available implants used by the group.
func (m *IntelligenceProfile) GetTradecraft()(FormattedContentable) {
    return m.tradecraft
}
// Serialize serializes information the current object
func (m *IntelligenceProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAliases() != nil {
        err = writer.WriteCollectionOfStringValues("aliases", m.GetAliases())
        if err != nil {
            return err
        }
    }
    if m.GetCountriesOrRegionsOfOrigin() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCountriesOrRegionsOfOrigin()))
        for i, v := range m.GetCountriesOrRegionsOfOrigin() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("countriesOrRegionsOfOrigin", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("firstActiveDateTime", m.GetFirstActiveDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetIndicators() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIndicators()))
        for i, v := range m.GetIndicators() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("indicators", cast)
        if err != nil {
            return err
        }
    }
    if m.GetKind() != nil {
        cast := (*m.GetKind()).String()
        err = writer.WriteStringValue("kind", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("summary", m.GetSummary())
        if err != nil {
            return err
        }
    }
    if m.GetTargets() != nil {
        err = writer.WriteCollectionOfStringValues("targets", m.GetTargets())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("tradecraft", m.GetTradecraft())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAliases sets the aliases property value. A list of commonly-known aliases for the threat intelligence included in the intelligenceProfile.
func (m *IntelligenceProfile) SetAliases(value []string)() {
    m.aliases = value
}
// SetCountriesOrRegionsOfOrigin sets the countriesOrRegionsOfOrigin property value. The country/region of origin for the given actor or threat associated with this intelligenceProfile.
func (m *IntelligenceProfile) SetCountriesOrRegionsOfOrigin(value []IntelligenceProfileCountryOrRegionOfOriginable)() {
    m.countriesOrRegionsOfOrigin = value
}
// SetDescription sets the description property value. The description property
func (m *IntelligenceProfile) SetDescription(value FormattedContentable)() {
    m.description = value
}
// SetFirstActiveDateTime sets the firstActiveDateTime property value. The date and time when this intelligenceProfile was first active. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *IntelligenceProfile) SetFirstActiveDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.firstActiveDateTime = value
}
// SetIndicators sets the indicators property value. Includes an assemblage of high-fidelity network indicators of compromise.
func (m *IntelligenceProfile) SetIndicators(value []IntelligenceProfileIndicatorable)() {
    m.indicators = value
}
// SetKind sets the kind property value. The kind property
func (m *IntelligenceProfile) SetKind(value *IntelligenceProfileKind)() {
    m.kind = value
}
// SetSummary sets the summary property value. The summary property
func (m *IntelligenceProfile) SetSummary(value FormattedContentable)() {
    m.summary = value
}
// SetTargets sets the targets property value. Known targets related to this intelligenceProfile.
func (m *IntelligenceProfile) SetTargets(value []string)() {
    m.targets = value
}
// SetTitle sets the title property value. The title of this intelligenceProfile.
func (m *IntelligenceProfile) SetTitle(value *string)() {
    m.title = value
}
// SetTradecraft sets the tradecraft property value. Formatted information featuring a description of the distinctive tactics, techniques, and procedures (TTP) of the group, followed by a list of all known custom, commodity, and publicly available implants used by the group.
func (m *IntelligenceProfile) SetTradecraft(value FormattedContentable)() {
    m.tradecraft = value
}
// IntelligenceProfileable 
type IntelligenceProfileable interface {
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAliases()([]string)
    GetCountriesOrRegionsOfOrigin()([]IntelligenceProfileCountryOrRegionOfOriginable)
    GetDescription()(FormattedContentable)
    GetFirstActiveDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIndicators()([]IntelligenceProfileIndicatorable)
    GetKind()(*IntelligenceProfileKind)
    GetSummary()(FormattedContentable)
    GetTargets()([]string)
    GetTitle()(*string)
    GetTradecraft()(FormattedContentable)
    SetAliases(value []string)()
    SetCountriesOrRegionsOfOrigin(value []IntelligenceProfileCountryOrRegionOfOriginable)()
    SetDescription(value FormattedContentable)()
    SetFirstActiveDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIndicators(value []IntelligenceProfileIndicatorable)()
    SetKind(value *IntelligenceProfileKind)()
    SetSummary(value FormattedContentable)()
    SetTargets(value []string)()
    SetTitle(value *string)()
    SetTradecraft(value FormattedContentable)()
}
