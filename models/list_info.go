package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ListInfo 
type ListInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // If true, indicates that content types are enabled for this list.
    contentTypesEnabled *bool
    // If true, indicates that the list is not normally visible in the SharePoint user experience.
    hidden *bool
    // The OdataType property
    odataType *string
    // An enumerated value that represents the base list template used in creating the list. Possible values include documentLibrary, genericList, task, survey, announcements, contacts, and more.
    template *string
}
// NewListInfo instantiates a new listInfo and sets the default values.
func NewListInfo()(*ListInfo) {
    m := &ListInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateListInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateListInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewListInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ListInfo) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetContentTypesEnabled gets the contentTypesEnabled property value. If true, indicates that content types are enabled for this list.
func (m *ListInfo) GetContentTypesEnabled()(*bool) {
    return m.contentTypesEnabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ListInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentTypesEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetContentTypesEnabled)
    res["hidden"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetHidden)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["template"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTemplate)
    return res
}
// GetHidden gets the hidden property value. If true, indicates that the list is not normally visible in the SharePoint user experience.
func (m *ListInfo) GetHidden()(*bool) {
    return m.hidden
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ListInfo) GetOdataType()(*string) {
    return m.odataType
}
// GetTemplate gets the template property value. An enumerated value that represents the base list template used in creating the list. Possible values include documentLibrary, genericList, task, survey, announcements, contacts, and more.
func (m *ListInfo) GetTemplate()(*string) {
    return m.template
}
// Serialize serializes information the current object
func (m *ListInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("contentTypesEnabled", m.GetContentTypesEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hidden", m.GetHidden())
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
        err := writer.WriteStringValue("template", m.GetTemplate())
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
func (m *ListInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetContentTypesEnabled sets the contentTypesEnabled property value. If true, indicates that content types are enabled for this list.
func (m *ListInfo) SetContentTypesEnabled(value *bool)() {
    m.contentTypesEnabled = value
}
// SetHidden sets the hidden property value. If true, indicates that the list is not normally visible in the SharePoint user experience.
func (m *ListInfo) SetHidden(value *bool)() {
    m.hidden = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ListInfo) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTemplate sets the template property value. An enumerated value that represents the base list template used in creating the list. Possible values include documentLibrary, genericList, task, survey, announcements, contacts, and more.
func (m *ListInfo) SetTemplate(value *string)() {
    m.template = value
}
