package security

import (
	i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryReviewSetCollectionResponse
type EdiscoveryReviewSetCollectionResponse struct {
	i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.BaseCollectionPaginationCountResponse
	// The value property
	value []EdiscoveryReviewSetable
}

// NewEdiscoveryReviewSetCollectionResponse instantiates a new EdiscoveryReviewSetCollectionResponse and sets the default values.
func NewEdiscoveryReviewSetCollectionResponse() *EdiscoveryReviewSetCollectionResponse {
	m := &EdiscoveryReviewSetCollectionResponse{
		BaseCollectionPaginationCountResponse: *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.NewBaseCollectionPaginationCountResponse(),
	}
	return m
}

// CreateEdiscoveryReviewSetCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryReviewSetCollectionResponseFromDiscriminatorValue(
	parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode,
) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEdiscoveryReviewSetCollectionResponse(), nil
}

// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryReviewSetCollectionResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateEdiscoveryReviewSetFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]EdiscoveryReviewSetable, len(val))
			for i, v := range val {
				res[i] = v.(EdiscoveryReviewSetable)
			}
			m.SetValue(res)
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
func (m *EdiscoveryReviewSetCollectionResponse) GetValue() []EdiscoveryReviewSetable {
	return m.value
}

// Serialize serializes information the current object
func (m *EdiscoveryReviewSetCollectionResponse) Serialize(
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
func (m *EdiscoveryReviewSetCollectionResponse) SetValue(value []EdiscoveryReviewSetable) {
	m.value = value
}

// EdiscoveryReviewSetCollectionResponseable
type EdiscoveryReviewSetCollectionResponseable interface {
	i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.BaseCollectionPaginationCountResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetValue() []EdiscoveryReviewSetable
	SetValue(value []EdiscoveryReviewSetable)
}
