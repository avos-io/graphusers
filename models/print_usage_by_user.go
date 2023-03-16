package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintUsageByUser
type PrintUsageByUser struct {
	PrintUsage
	// The UPN of the user represented by these statistics.
	userPrincipalName *string
}

// NewPrintUsageByUser instantiates a new PrintUsageByUser and sets the default values.
func NewPrintUsageByUser() *PrintUsageByUser {
	m := &PrintUsageByUser{
		PrintUsage: *NewPrintUsage(),
	}
	odataTypeValue := "#microsoft.graph.printUsageByUser"
	m.SetOdataType(&odataTypeValue)
	return m
}

// CreatePrintUsageByUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintUsageByUserFromDiscriminatorValue(
	parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode,
) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPrintUsageByUser(), nil
}

// GetFieldDeserializers the deserialization information for the current model
func (m *PrintUsageByUser) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.PrintUsage.GetFieldDeserializers()
	res["userPrincipalName"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUserPrincipalName(val)
		}
		return nil
	}
	return res
}

// GetUserPrincipalName gets the userPrincipalName property value. The UPN of the user represented by these statistics.
func (m *PrintUsageByUser) GetUserPrincipalName() *string {
	return m.userPrincipalName
}

// Serialize serializes information the current object
func (m *PrintUsageByUser) Serialize(
	writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter,
) error {
	err := m.PrintUsage.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetUserPrincipalName sets the userPrincipalName property value. The UPN of the user represented by these statistics.
func (m *PrintUsageByUser) SetUserPrincipalName(value *string) {
	m.userPrincipalName = value
}

// PrintUsageByUserable
type PrintUsageByUserable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	PrintUsageable
	GetUserPrincipalName() *string
	SetUserPrincipalName(value *string)
}
