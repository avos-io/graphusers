package users

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemChatsItemHideForUserPostRequestBodyable 
type ItemChatsItemHideForUserPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetUser()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TeamworkUserIdentityable)
    SetUser(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TeamworkUserIdentityable)()
}
