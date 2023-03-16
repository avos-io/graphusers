package security

import (
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"

	i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DataSource
type DataSource struct {
	i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Entity
	// The user who created the dataSource.
	createdBy i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.IdentitySetable
	// The date and time the dataSource was created.
	createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The display name of the dataSource. This will be the name of the SharePoint site.
	displayName *string
	// The hold status of the dataSource.The possible values are: notApplied, applied, applying, removing, partial
	holdStatus *DataSourceHoldStatus
}

// NewDataSource instantiates a new dataSource and sets the default values.
func NewDataSource() *DataSource {
	m := &DataSource{
		Entity: *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.NewEntity(),
	}
	return m
}

// CreateDataSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDataSourceFromDiscriminatorValue(
	parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode,
) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	if parseNode != nil {
		mappingValueNode, err := parseNode.GetChildNode("@odata.type")
		if err != nil {
			return nil, err
		}
		if mappingValueNode != nil {
			mappingValue, err := mappingValueNode.GetStringValue()
			if err != nil {
				return nil, err
			}
			if mappingValue != nil {
				switch *mappingValue {
				case "#microsoft.graph.security.siteSource":
					return NewSiteSource(), nil
				case "#microsoft.graph.security.unifiedGroupSource":
					return NewUnifiedGroupSource(), nil
				case "#microsoft.graph.security.userSource":
					return NewUserSource(), nil
				}
			}
		}
	}
	return NewDataSource(), nil
}

// GetCreatedBy gets the createdBy property value. The user who created the dataSource.
func (m *DataSource) GetCreatedBy() i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.IdentitySetable {
	return m.createdBy
}

// GetCreatedDateTime gets the createdDateTime property value. The date and time the dataSource was created.
func (m *DataSource) GetCreatedDateTime() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.createdDateTime
}

// GetDisplayName gets the displayName property value. The display name of the dataSource. This will be the name of the SharePoint site.
func (m *DataSource) GetDisplayName() *string {
	return m.displayName
}

// GetFieldDeserializers the deserialization information for the current model
func (m *DataSource) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Entity.GetFieldDeserializers()
	res["createdBy"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(
			i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateIdentitySetFromDiscriminatorValue,
		)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCreatedBy(
				val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.IdentitySetable),
			)
		}
		return nil
	}
	res["createdDateTime"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCreatedDateTime(val)
		}
		return nil
	}
	res["displayName"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDisplayName(val)
		}
		return nil
	}
	res["holdStatus"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseDataSourceHoldStatus)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetHoldStatus(val.(*DataSourceHoldStatus))
		}
		return nil
	}
	return res
}

// GetHoldStatus gets the holdStatus property value. The hold status of the dataSource.The possible values are: notApplied, applied, applying, removing, partial
func (m *DataSource) GetHoldStatus() *DataSourceHoldStatus {
	return m.holdStatus
}

// Serialize serializes information the current object
func (m *DataSource) Serialize(
	writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter,
) error {
	err := m.Entity.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteStringValue("displayName", m.GetDisplayName())
		if err != nil {
			return err
		}
	}
	if m.GetHoldStatus() != nil {
		cast := (*m.GetHoldStatus()).String()
		err = writer.WriteStringValue("holdStatus", &cast)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetCreatedBy sets the createdBy property value. The user who created the dataSource.
func (m *DataSource) SetCreatedBy(
	value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.IdentitySetable,
) {
	m.createdBy = value
}

// SetCreatedDateTime sets the createdDateTime property value. The date and time the dataSource was created.
func (m *DataSource) SetCreatedDateTime(
	value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time,
) {
	m.createdDateTime = value
}

// SetDisplayName sets the displayName property value. The display name of the dataSource. This will be the name of the SharePoint site.
func (m *DataSource) SetDisplayName(value *string) {
	m.displayName = value
}

// SetHoldStatus sets the holdStatus property value. The hold status of the dataSource.The possible values are: notApplied, applied, applying, removing, partial
func (m *DataSource) SetHoldStatus(value *DataSourceHoldStatus) {
	m.holdStatus = value
}

// DataSourceable
type DataSourceable interface {
	i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Entityable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCreatedBy() i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.IdentitySetable
	GetCreatedDateTime() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetDisplayName() *string
	GetHoldStatus() *DataSourceHoldStatus
	SetCreatedBy(
		value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.IdentitySetable,
	)
	SetCreatedDateTime(
		value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time,
	)
	SetDisplayName(value *string)
	SetHoldStatus(value *DataSourceHoldStatus)
}
