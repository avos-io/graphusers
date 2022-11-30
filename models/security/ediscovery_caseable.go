package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryCaseable 
type EdiscoveryCaseable interface {
    Case_escapedable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClosedBy()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.IdentitySetable)
    GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustodians()([]EdiscoveryCustodianable)
    GetExternalId()(*string)
    GetNoncustodialDataSources()([]EdiscoveryNoncustodialDataSourceable)
    GetOperations()([]CaseOperationable)
    GetReviewSets()([]EdiscoveryReviewSetable)
    GetSearches()([]EdiscoverySearchable)
    GetSettings()(EdiscoveryCaseSettingsable)
    GetTags()([]EdiscoveryReviewTagable)
    SetClosedBy(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.IdentitySetable)()
    SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustodians(value []EdiscoveryCustodianable)()
    SetExternalId(value *string)()
    SetNoncustodialDataSources(value []EdiscoveryNoncustodialDataSourceable)()
    SetOperations(value []CaseOperationable)()
    SetReviewSets(value []EdiscoveryReviewSetable)()
    SetSearches(value []EdiscoverySearchable)()
    SetSettings(value EdiscoveryCaseSettingsable)()
    SetTags(value []EdiscoveryReviewTagable)()
}
