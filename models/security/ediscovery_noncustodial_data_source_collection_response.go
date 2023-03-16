package security

import (
	i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryNoncustodialDataSourceCollectionResponse
type EdiscoveryNoncustodialDataSourceCollectionResponse struct {
	i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.BaseCollectionPaginationCountResponse
	// The value property
	value []EdiscoveryNoncustodialDataSourceable
}

// NewEdiscoveryNoncustodialDataSourceCollectionResponse instantiates a new EdiscoveryNoncustodialDataSourceCollectionResponse and sets the default values.
func NewEdiscoveryNoncustodialDataSourceCollectionResponse() *EdiscoveryNoncustodialDataSourceCollectionResponse {
	m := &EdiscoveryNoncustodialDataSourceCollectionResponse{
		BaseCollectionPaginationCountResponse: *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.NewBaseCollectionPaginationCountResponse(),
	}
	return m
}

// CreateEdiscoveryNoncustodialDataSourceCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryNoncustodialDataSourceCollectionResponseFromDiscriminatorValue(
	parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode,
) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEdiscoveryNoncustodialDataSourceCollectionResponse(), nil
}

// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryNoncustodialDataSourceCollectionResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(
			CreateEdiscoveryNoncustodialDataSourceFromDiscriminatorValue,
		)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]EdiscoveryNoncustodialDataSourceable, len(val))
			for i, v := range val {
				res[i] = v.(EdiscoveryNoncustodialDataSourceable)
			}
			m.SetValue(res)
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
func (m *EdiscoveryNoncustodialDataSourceCollectionResponse) GetValue() []EdiscoveryNoncustodialDataSourceable {
	return m.value
}

// Serialize serializes information the current object
func (m *EdiscoveryNoncustodialDataSourceCollectionResponse) Serialize(
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
func (m *EdiscoveryNoncustodialDataSourceCollectionResponse) SetValue(
	value []EdiscoveryNoncustodialDataSourceable,
) {
	m.value = value
}

// EdiscoveryNoncustodialDataSourceCollectionResponseable
type EdiscoveryNoncustodialDataSourceCollectionResponseable interface {
	i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.BaseCollectionPaginationCountResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetValue() []EdiscoveryNoncustodialDataSourceable
	SetValue(value []EdiscoveryNoncustodialDataSourceable)
}
