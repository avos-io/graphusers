package users

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBodyable 
type ItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailabilityViewInterval()(*int32)
    GetEndTime()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.DateTimeTimeZoneable)
    GetSchedules()([]string)
    GetStartTime()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.DateTimeTimeZoneable)
    SetAvailabilityViewInterval(value *int32)()
    SetEndTime(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.DateTimeTimeZoneable)()
    SetSchedules(value []string)()
    SetStartTime(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.DateTimeTimeZoneable)()
}
