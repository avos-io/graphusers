package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceAnnouncementAttachment provides operations to manage the collection of user entities.
type ServiceAnnouncementAttachment struct {
    Entity
    // The attachment content.
    content []byte
    // The contentType property
    contentType *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name property
    name *string
    // The size property
    size *int32
}
// NewServiceAnnouncementAttachment instantiates a new serviceAnnouncementAttachment and sets the default values.
func NewServiceAnnouncementAttachment()(*ServiceAnnouncementAttachment) {
    m := &ServiceAnnouncementAttachment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateServiceAnnouncementAttachmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceAnnouncementAttachmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceAnnouncementAttachment(), nil
}
// GetContent gets the content property value. The attachment content.
func (m *ServiceAnnouncementAttachment) GetContent()([]byte) {
    return m.content
}
// GetContentType gets the contentType property value. The contentType property
func (m *ServiceAnnouncementAttachment) GetContentType()(*string) {
    return m.contentType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceAnnouncementAttachment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetContent)
    res["contentType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetContentType)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["size"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSize)
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *ServiceAnnouncementAttachment) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetName gets the name property value. The name property
func (m *ServiceAnnouncementAttachment) GetName()(*string) {
    return m.name
}
// GetSize gets the size property value. The size property
func (m *ServiceAnnouncementAttachment) GetSize()(*int32) {
    return m.size
}
// Serialize serializes information the current object
func (m *ServiceAnnouncementAttachment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentType", m.GetContentType())
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
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The attachment content.
func (m *ServiceAnnouncementAttachment) SetContent(value []byte)() {
    m.content = value
}
// SetContentType sets the contentType property value. The contentType property
func (m *ServiceAnnouncementAttachment) SetContentType(value *string)() {
    m.contentType = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *ServiceAnnouncementAttachment) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetName sets the name property value. The name property
func (m *ServiceAnnouncementAttachment) SetName(value *string)() {
    m.name = value
}
// SetSize sets the size property value. The size property
func (m *ServiceAnnouncementAttachment) SetSize(value *int32)() {
    m.size = value
}
