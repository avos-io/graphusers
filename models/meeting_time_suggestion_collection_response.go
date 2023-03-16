package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingTimeSuggestionCollectionResponse
type MeetingTimeSuggestionCollectionResponse struct {
	BaseCollectionPaginationCountResponse
	// The value property
	value []MeetingTimeSuggestionable
}

// NewMeetingTimeSuggestionCollectionResponse instantiates a new MeetingTimeSuggestionCollectionResponse and sets the default values.
func NewMeetingTimeSuggestionCollectionResponse() *MeetingTimeSuggestionCollectionResponse {
	m := &MeetingTimeSuggestionCollectionResponse{
		BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
	}
	return m
}

// CreateMeetingTimeSuggestionCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingTimeSuggestionCollectionResponseFromDiscriminatorValue(
	parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode,
) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewMeetingTimeSuggestionCollectionResponse(), nil
}

// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingTimeSuggestionCollectionResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateMeetingTimeSuggestionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]MeetingTimeSuggestionable, len(val))
			for i, v := range val {
				res[i] = v.(MeetingTimeSuggestionable)
			}
			m.SetValue(res)
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
func (m *MeetingTimeSuggestionCollectionResponse) GetValue() []MeetingTimeSuggestionable {
	return m.value
}

// Serialize serializes information the current object
func (m *MeetingTimeSuggestionCollectionResponse) Serialize(
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
func (m *MeetingTimeSuggestionCollectionResponse) SetValue(value []MeetingTimeSuggestionable) {
	m.value = value
}

// MeetingTimeSuggestionCollectionResponseable
type MeetingTimeSuggestionCollectionResponseable interface {
	BaseCollectionPaginationCountResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetValue() []MeetingTimeSuggestionable
	SetValue(value []MeetingTimeSuggestionable)
}
