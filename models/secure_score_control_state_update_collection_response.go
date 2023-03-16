package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecureScoreControlStateUpdateCollectionResponse
type SecureScoreControlStateUpdateCollectionResponse struct {
	BaseCollectionPaginationCountResponse
	// The value property
	value []SecureScoreControlStateUpdateable
}

// NewSecureScoreControlStateUpdateCollectionResponse instantiates a new SecureScoreControlStateUpdateCollectionResponse and sets the default values.
func NewSecureScoreControlStateUpdateCollectionResponse() *SecureScoreControlStateUpdateCollectionResponse {
	m := &SecureScoreControlStateUpdateCollectionResponse{
		BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
	}
	return m
}

// CreateSecureScoreControlStateUpdateCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecureScoreControlStateUpdateCollectionResponseFromDiscriminatorValue(
	parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode,
) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSecureScoreControlStateUpdateCollectionResponse(), nil
}

// GetFieldDeserializers the deserialization information for the current model
func (m *SecureScoreControlStateUpdateCollectionResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(
			CreateSecureScoreControlStateUpdateFromDiscriminatorValue,
		)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]SecureScoreControlStateUpdateable, len(val))
			for i, v := range val {
				res[i] = v.(SecureScoreControlStateUpdateable)
			}
			m.SetValue(res)
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
func (m *SecureScoreControlStateUpdateCollectionResponse) GetValue() []SecureScoreControlStateUpdateable {
	return m.value
}

// Serialize serializes information the current object
func (m *SecureScoreControlStateUpdateCollectionResponse) Serialize(
	writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter,
) error {
	err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
	if err != nil {
		return err
	}
	if m.GetValue() != nil {
		cast := make(
			[]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable,
			len(m.GetValue()),
		)
		for i, v := range m.GetValue() {
			cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
		}
		err = writer.WriteCollectionOfObjectValues("value", cast)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetValue sets the value property value. The value property
func (m *SecureScoreControlStateUpdateCollectionResponse) SetValue(
	value []SecureScoreControlStateUpdateable,
) {
	m.value = value
}

// SecureScoreControlStateUpdateCollectionResponseable
type SecureScoreControlStateUpdateCollectionResponseable interface {
	BaseCollectionPaginationCountResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetValue() []SecureScoreControlStateUpdateable
	SetValue(value []SecureScoreControlStateUpdateable)
}
