package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsAppDefinitionCollectionResponse
type TeamsAppDefinitionCollectionResponse struct {
	BaseCollectionPaginationCountResponse
	// The value property
	value []TeamsAppDefinitionable
}

// NewTeamsAppDefinitionCollectionResponse instantiates a new TeamsAppDefinitionCollectionResponse and sets the default values.
func NewTeamsAppDefinitionCollectionResponse() *TeamsAppDefinitionCollectionResponse {
	m := &TeamsAppDefinitionCollectionResponse{
		BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
	}
	return m
}

// CreateTeamsAppDefinitionCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsAppDefinitionCollectionResponseFromDiscriminatorValue(
	parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode,
) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTeamsAppDefinitionCollectionResponse(), nil
}

// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsAppDefinitionCollectionResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateTeamsAppDefinitionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]TeamsAppDefinitionable, len(val))
			for i, v := range val {
				res[i] = v.(TeamsAppDefinitionable)
			}
			m.SetValue(res)
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
func (m *TeamsAppDefinitionCollectionResponse) GetValue() []TeamsAppDefinitionable {
	return m.value
}

// Serialize serializes information the current object
func (m *TeamsAppDefinitionCollectionResponse) Serialize(
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
func (m *TeamsAppDefinitionCollectionResponse) SetValue(value []TeamsAppDefinitionable) {
	m.value = value
}

// TeamsAppDefinitionCollectionResponseable
type TeamsAppDefinitionCollectionResponseable interface {
	BaseCollectionPaginationCountResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetValue() []TeamsAppDefinitionable
	SetValue(value []TeamsAppDefinitionable)
}
