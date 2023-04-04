package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppHostedMediaConfig 
type AppHostedMediaConfig struct {
    MediaConfig
    // The media configuration blob generated by smart media agent.
    blob *string
}
// NewAppHostedMediaConfig instantiates a new AppHostedMediaConfig and sets the default values.
func NewAppHostedMediaConfig()(*AppHostedMediaConfig) {
    m := &AppHostedMediaConfig{
        MediaConfig: *NewMediaConfig(),
    }
    odataTypeValue := "#microsoft.graph.appHostedMediaConfig"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAppHostedMediaConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppHostedMediaConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppHostedMediaConfig(), nil
}
// GetBlob gets the blob property value. The media configuration blob generated by smart media agent.
func (m *AppHostedMediaConfig) GetBlob()(*string) {
    return m.blob
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppHostedMediaConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MediaConfig.GetFieldDeserializers()
    res["blob"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlob(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AppHostedMediaConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MediaConfig.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("blob", m.GetBlob())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBlob sets the blob property value. The media configuration blob generated by smart media agent.
func (m *AppHostedMediaConfig) SetBlob(value *string)() {
    m.blob = value
}
// AppHostedMediaConfigable 
type AppHostedMediaConfigable interface {
    MediaConfigable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBlob()(*string)
    SetBlob(value *string)()
}
