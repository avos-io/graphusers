package users

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemFindMeetingTimesPostRequestBodyable 
type ItemFindMeetingTimesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttendees()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AttendeeBaseable)
    GetIsOrganizerOptional()(*bool)
    GetLocationConstraint()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.LocationConstraintable)
    GetMaxCandidates()(*int32)
    GetMeetingDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetMinimumAttendeePercentage()(*float64)
    GetReturnSuggestionReasons()(*bool)
    GetTimeConstraint()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TimeConstraintable)
    SetAttendees(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.AttendeeBaseable)()
    SetIsOrganizerOptional(value *bool)()
    SetLocationConstraint(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.LocationConstraintable)()
    SetMaxCandidates(value *int32)()
    SetMeetingDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetMinimumAttendeePercentage(value *float64)()
    SetReturnSuggestionReasons(value *bool)()
    SetTimeConstraint(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TimeConstraintable)()
}
