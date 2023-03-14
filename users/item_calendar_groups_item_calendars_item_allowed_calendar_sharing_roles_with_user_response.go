package users

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse 
type ItemCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse struct {
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.BaseCollectionPaginationCountResponse
    // The value property
    value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CalendarRoleType
}
// NewItemCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse instantiates a new ItemCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse()(*ItemCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) {
    m := &ItemCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse{
        BaseCollectionPaginationCountResponse: *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateItemCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) GetValue()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CalendarRoleType) {
    return m.value
}
// Serialize serializes information the current object
func (m *ItemCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) SetValue(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CalendarRoleType)() {
    m.value = value
}
