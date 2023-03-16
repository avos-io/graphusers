package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationExcelResource
type EducationExcelResource struct {
	EducationResource
	// Pointer to the Excel file object.
	fileUrl *string
}

// NewEducationExcelResource instantiates a new EducationExcelResource and sets the default values.
func NewEducationExcelResource() *EducationExcelResource {
	m := &EducationExcelResource{
		EducationResource: *NewEducationResource(),
	}
	odataTypeValue := "#microsoft.graph.educationExcelResource"
	m.SetOdataType(&odataTypeValue)
	return m
}

// CreateEducationExcelResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationExcelResourceFromDiscriminatorValue(
	parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode,
) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEducationExcelResource(), nil
}

// GetFieldDeserializers the deserialization information for the current model
func (m *EducationExcelResource) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.EducationResource.GetFieldDeserializers()
	res["fileUrl"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFileUrl(val)
		}
		return nil
	}
	return res
}

// GetFileUrl gets the fileUrl property value. Pointer to the Excel file object.
func (m *EducationExcelResource) GetFileUrl() *string {
	return m.fileUrl
}

// Serialize serializes information the current object
func (m *EducationExcelResource) Serialize(
	writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter,
) error {
	err := m.EducationResource.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteStringValue("fileUrl", m.GetFileUrl())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetFileUrl sets the fileUrl property value. Pointer to the Excel file object.
func (m *EducationExcelResource) SetFileUrl(value *string) {
	m.fileUrl = value
}

// EducationExcelResourceable
type EducationExcelResourceable interface {
	EducationResourceable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetFileUrl() *string
	SetFileUrl(value *string)
}
