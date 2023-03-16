package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OmaSettingString
type OmaSettingString struct {
	OmaSetting
	// Value.
	value *string
}

// NewOmaSettingString instantiates a new OmaSettingString and sets the default values.
func NewOmaSettingString() *OmaSettingString {
	m := &OmaSettingString{
		OmaSetting: *NewOmaSetting(),
	}
	odataTypeValue := "#microsoft.graph.omaSettingString"
	m.SetOdataType(&odataTypeValue)
	return m
}

// CreateOmaSettingStringFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOmaSettingStringFromDiscriminatorValue(
	parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode,
) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOmaSettingString(), nil
}

// GetFieldDeserializers the deserialization information for the current model
func (m *OmaSettingString) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.OmaSetting.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val)
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. Value.
func (m *OmaSettingString) GetValue() *string {
	return m.value
}

// Serialize serializes information the current object
func (m *OmaSettingString) Serialize(
	writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter,
) error {
	err := m.OmaSetting.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteStringValue("value", m.GetValue())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetValue sets the value property value. Value.
func (m *OmaSettingString) SetValue(value *string) {
	m.value = value
}

// OmaSettingStringable
type OmaSettingStringable interface {
	OmaSettingable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetValue() *string
	SetValue(value *string)
}
