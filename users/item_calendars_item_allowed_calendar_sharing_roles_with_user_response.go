package users

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse 
type ItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse struct {
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.BaseCollectionPaginationCountResponse
    // The value property
    value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CalendarRoleType
}
// NewItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse instantiates a new ItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse and sets the default values.
func NewItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse()(*ItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) {
    m := &ItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse{
        BaseCollectionPaginationCountResponse: *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateItemCalendarsItemAllowedCalendarSharingRolesWithUserResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarsItemAllowedCalendarSharingRolesWithUserResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ParseCalendarRoleType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CalendarRoleType, len(val))
            for i, v := range val {
                res[i] = *(v.(*i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CalendarRoleType))
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) GetValue()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CalendarRoleType) {
    return m.value
}
// Serialize serializes information the current object
func (m *ItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        err = writer.WriteCollectionOfStringValues("value", i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.SerializeCalendarRoleType(m.GetValue()))
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *ItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) SetValue(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CalendarRoleType)() {
    m.value = value
}
