package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RedirectUriSettings 
type RedirectUriSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The index property
    index *int32
    // The OdataType property
    odataType *string
    // The uri property
    uri *string
}
// NewRedirectUriSettings instantiates a new redirectUriSettings and sets the default values.
func NewRedirectUriSettings()(*RedirectUriSettings) {
    m := &RedirectUriSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRedirectUriSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRedirectUriSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRedirectUriSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RedirectUriSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RedirectUriSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["index"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetIndex)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["uri"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUri)
    return res
}
// GetIndex gets the index property value. The index property
func (m *RedirectUriSettings) GetIndex()(*int32) {
    return m.index
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RedirectUriSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetUri gets the uri property value. The uri property
func (m *RedirectUriSettings) GetUri()(*string) {
    return m.uri
}
// Serialize serializes information the current object
func (m *RedirectUriSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("index", m.GetIndex())
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
        err := writer.WriteStringValue("uri", m.GetUri())
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
func (m *RedirectUriSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIndex sets the index property value. The index property
func (m *RedirectUriSettings) SetIndex(value *int32)() {
    m.index = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RedirectUriSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUri sets the uri property value. The uri property
func (m *RedirectUriSettings) SetUri(value *string)() {
    m.uri = value
}
