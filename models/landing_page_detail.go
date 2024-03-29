package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LandingPageDetail 
type LandingPageDetail struct {
    Entity
    // The content property
    content *string
    // The isDefaultLangauge property
    isDefaultLangauge *bool
    // The language property
    language *string
}
// NewLandingPageDetail instantiates a new landingPageDetail and sets the default values.
func NewLandingPageDetail()(*LandingPageDetail) {
    m := &LandingPageDetail{
        Entity: *NewEntity(),
    }
    return m
}
// CreateLandingPageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLandingPageDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLandingPageDetail(), nil
}
// GetContent gets the content property value. The content property
func (m *LandingPageDetail) GetContent()(*string) {
    return m.content
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LandingPageDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["isDefaultLangauge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefaultLangauge(val)
        }
        return nil
    }
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
        }
        return nil
    }
    return res
}
// GetIsDefaultLangauge gets the isDefaultLangauge property value. The isDefaultLangauge property
func (m *LandingPageDetail) GetIsDefaultLangauge()(*bool) {
    return m.isDefaultLangauge
}
// GetLanguage gets the language property value. The language property
func (m *LandingPageDetail) GetLanguage()(*string) {
    return m.language
}
// Serialize serializes information the current object
func (m *LandingPageDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefaultLangauge", m.GetIsDefaultLangauge())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("language", m.GetLanguage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The content property
func (m *LandingPageDetail) SetContent(value *string)() {
    m.content = value
}
// SetIsDefaultLangauge sets the isDefaultLangauge property value. The isDefaultLangauge property
func (m *LandingPageDetail) SetIsDefaultLangauge(value *bool)() {
    m.isDefaultLangauge = value
}
// SetLanguage sets the language property value. The language property
func (m *LandingPageDetail) SetLanguage(value *string)() {
    m.language = value
}
// LandingPageDetailable 
type LandingPageDetailable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()(*string)
    GetIsDefaultLangauge()(*bool)
    GetLanguage()(*string)
    SetContent(value *string)()
    SetIsDefaultLangauge(value *bool)()
    SetLanguage(value *string)()
}
