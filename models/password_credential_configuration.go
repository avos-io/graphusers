package models

import (
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"

	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PasswordCredentialConfiguration
type PasswordCredentialConfiguration struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The maxLifetime property
	maxLifetime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
	// The OdataType property
	odataType *string
	// The restrictForAppsCreatedAfterDateTime property
	restrictForAppsCreatedAfterDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The restrictionType property
	restrictionType *AppCredentialRestrictionType
}

// NewPasswordCredentialConfiguration instantiates a new passwordCredentialConfiguration and sets the default values.
func NewPasswordCredentialConfiguration() *PasswordCredentialConfiguration {
	m := &PasswordCredentialConfiguration{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreatePasswordCredentialConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePasswordCredentialConfigurationFromDiscriminatorValue(
	parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode,
) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPasswordCredentialConfiguration(), nil
}

// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PasswordCredentialConfiguration) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
func (m *PasswordCredentialConfiguration) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(
		map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error,
	)
	res["maxLifetime"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetISODurationValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMaxLifetime(val)
		}
		return nil
	}
	res["@odata.type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOdataType(val)
		}
		return nil
	}
	res["restrictForAppsCreatedAfterDateTime"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRestrictForAppsCreatedAfterDateTime(val)
		}
		return nil
	}
	res["restrictionType"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseAppCredentialRestrictionType)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRestrictionType(val.(*AppCredentialRestrictionType))
		}
		return nil
	}
	return res
}

// GetMaxLifetime gets the maxLifetime property value. The maxLifetime property
func (m *PasswordCredentialConfiguration) GetMaxLifetime() *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration {
	return m.maxLifetime
}

// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PasswordCredentialConfiguration) GetOdataType() *string {
	return m.odataType
}

// GetRestrictForAppsCreatedAfterDateTime gets the restrictForAppsCreatedAfterDateTime property value. The restrictForAppsCreatedAfterDateTime property
func (m *PasswordCredentialConfiguration) GetRestrictForAppsCreatedAfterDateTime() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.restrictForAppsCreatedAfterDateTime
}

// GetRestrictionType gets the restrictionType property value. The restrictionType property
func (m *PasswordCredentialConfiguration) GetRestrictionType() *AppCredentialRestrictionType {
	return m.restrictionType
}

// Serialize serializes information the current object
func (m *PasswordCredentialConfiguration) Serialize(
	writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter,
) error {
	{
		err := writer.WriteISODurationValue("maxLifetime", m.GetMaxLifetime())
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
		err := writer.WriteTimeValue(
			"restrictForAppsCreatedAfterDateTime",
			m.GetRestrictForAppsCreatedAfterDateTime(),
		)
		if err != nil {
			return err
		}
	}
	if m.GetRestrictionType() != nil {
		cast := (*m.GetRestrictionType()).String()
		err := writer.WriteStringValue("restrictionType", &cast)
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
func (m *PasswordCredentialConfiguration) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetMaxLifetime sets the maxLifetime property value. The maxLifetime property
func (m *PasswordCredentialConfiguration) SetMaxLifetime(
	value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration,
) {
	m.maxLifetime = value
}

// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PasswordCredentialConfiguration) SetOdataType(value *string) {
	m.odataType = value
}

// SetRestrictForAppsCreatedAfterDateTime sets the restrictForAppsCreatedAfterDateTime property value. The restrictForAppsCreatedAfterDateTime property
func (m *PasswordCredentialConfiguration) SetRestrictForAppsCreatedAfterDateTime(
	value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time,
) {
	m.restrictForAppsCreatedAfterDateTime = value
}

// SetRestrictionType sets the restrictionType property value. The restrictionType property
func (m *PasswordCredentialConfiguration) SetRestrictionType(value *AppCredentialRestrictionType) {
	m.restrictionType = value
}

// PasswordCredentialConfigurationable
type PasswordCredentialConfigurationable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetMaxLifetime() *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
	GetOdataType() *string
	GetRestrictForAppsCreatedAfterDateTime() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetRestrictionType() *AppCredentialRestrictionType
	SetMaxLifetime(
		value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration,
	)
	SetOdataType(value *string)
	SetRestrictForAppsCreatedAfterDateTime(
		value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time,
	)
	SetRestrictionType(value *AppCredentialRestrictionType)
}
