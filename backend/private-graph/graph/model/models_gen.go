// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Account struct {
	ID                   int        `json:"id"`
	Name                 string     `json:"name"`
	SessionCountCur      int        `json:"session_count_cur"`
	SessionCountPrev     int        `json:"session_count_prev"`
	SessionCountPrevPrev int        `json:"session_count_prev_prev"`
	SessionLimit         int        `json:"session_limit"`
	PaidPrev             int        `json:"paid_prev"`
	PaidPrevPrev         int        `json:"paid_prev_prev"`
	Email                string     `json:"email"`
	SubscriptionStart    *time.Time `json:"subscription_start"`
	PlanTier             string     `json:"plan_tier"`
	StripeCustomerID     string     `json:"stripe_customer_id"`
	MemberCount          int        `json:"member_count"`
	MemberLimit          int        `json:"member_limit"`
}

type AccountDetails struct {
	ID                   int           `json:"id"`
	Name                 string        `json:"name"`
	SessionCountPerMonth []*NamedCount `json:"session_count_per_month"`
	SessionCountPerDay   []*NamedCount `json:"session_count_per_day"`
	StripeCustomerID     string        `json:"stripe_customer_id"`
}

type AdminAboutYouDetails struct {
	FirstName          string  `json:"first_name"`
	LastName           string  `json:"last_name"`
	UserDefinedRole    string  `json:"user_defined_role"`
	UserDefinedPersona string  `json:"user_defined_persona"`
	Referral           string  `json:"referral"`
	Phone              *string `json:"phone"`
}

type AverageSessionLength struct {
	Length float64 `json:"length"`
}

type BillingDetails struct {
	Plan               *Plan `json:"plan"`
	Meter              int64 `json:"meter"`
	MembersMeter       int64 `json:"membersMeter"`
	SessionsOutOfQuota int64 `json:"sessionsOutOfQuota"`
}

type DashboardParamsInput struct {
	DateRange         *DateRangeInput `json:"date_range"`
	ResolutionMinutes *int            `json:"resolution_minutes"`
	Timezone          *string         `json:"timezone"`
}

type DashboardPayload struct {
	Date string  `json:"date"`
	Avg  float64 `json:"avg"`
	P50  float64 `json:"p50"`
	P75  float64 `json:"p75"`
	P90  float64 `json:"p90"`
	P99  float64 `json:"p99"`
}

type DateRangeInput struct {
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

type EnhancedUserDetailsResult struct {
	ID      *int          `json:"id"`
	Name    *string       `json:"name"`
	Avatar  *string       `json:"avatar"`
	Bio     *string       `json:"bio"`
	Socials []*SocialLink `json:"socials"`
	Email   *string       `json:"email"`
}

type ErrorDistributionItem struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

type ErrorMetadata struct {
	ErrorID         int        `json:"error_id"`
	SessionID       int        `json:"session_id"`
	SessionSecureID string     `json:"session_secure_id"`
	Environment     *string    `json:"environment"`
	Timestamp       *time.Time `json:"timestamp"`
	Os              *string    `json:"os"`
	Browser         *string    `json:"browser"`
	VisitedURL      *string    `json:"visited_url"`
	Fingerprint     string     `json:"fingerprint"`
	Identifier      *string    `json:"identifier"`
	UserProperties  *string    `json:"user_properties"`
	RequestID       *string    `json:"request_id"`
}

type ErrorSearchParamsInput struct {
	DateRange  *DateRangeInput `json:"date_range"`
	Os         *string         `json:"os"`
	Browser    *string         `json:"browser"`
	VisitedURL *string         `json:"visited_url"`
	State      *ErrorState     `json:"state"`
	Event      *string         `json:"event"`
	Type       *string         `json:"type"`
	Query      *string         `json:"query"`
}

type ErrorTrace struct {
	FileName     *string `json:"fileName"`
	LineNumber   *int    `json:"lineNumber"`
	FunctionName *string `json:"functionName"`
	ColumnNumber *int    `json:"columnNumber"`
	Error        *string `json:"error"`
	LineContent  *string `json:"lineContent"`
	LinesBefore  *string `json:"linesBefore"`
	LinesAfter   *string `json:"linesAfter"`
}

type Invoice struct {
	AmountDue    *int64     `json:"amountDue"`
	AmountPaid   *int64     `json:"amountPaid"`
	AttemptCount *int64     `json:"attemptCount"`
	Date         *time.Time `json:"date"`
	URL          *string    `json:"url"`
	Status       *string    `json:"status"`
}

type LengthRangeInput struct {
	Min *float64 `json:"min"`
	Max *float64 `json:"max"`
}

type LinearTeam struct {
	TeamID string `json:"team_id"`
	Name   string `json:"name"`
	Key    string `json:"key"`
}

type MetricPreview struct {
	Date  time.Time `json:"date"`
	Value float64   `json:"value"`
}

type NamedCount struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type NewUsersCount struct {
	Count int64 `json:"count"`
}

type Plan struct {
	Type         PlanType             `json:"type"`
	Interval     SubscriptionInterval `json:"interval"`
	Quota        int                  `json:"quota"`
	MembersLimit int                  `json:"membersLimit"`
}

type RageClickEventForProject struct {
	Identifier      string `json:"identifier"`
	SessionSecureID string `json:"session_secure_id"`
	TotalClicks     int    `json:"total_clicks"`
	UserProperties  string `json:"user_properties"`
}

type ReferrerTablePayload struct {
	Host    string  `json:"host"`
	Count   int     `json:"count"`
	Percent float64 `json:"percent"`
}

type SanitizedAdmin struct {
	ID       int     `json:"id"`
	Name     *string `json:"name"`
	Email    string  `json:"email"`
	PhotoURL *string `json:"photo_url"`
}

type SanitizedAdminInput struct {
	ID    int     `json:"id"`
	Name  *string `json:"name"`
	Email string  `json:"email"`
}

type SanitizedSlackChannel struct {
	WebhookChannel   *string `json:"webhook_channel"`
	WebhookChannelID *string `json:"webhook_channel_id"`
}

type SanitizedSlackChannelInput struct {
	WebhookChannelName *string `json:"webhook_channel_name"`
	WebhookChannelID   *string `json:"webhook_channel_id"`
}

type SearchParamsInput struct {
	UserProperties          []*UserPropertyInput `json:"user_properties"`
	ExcludedProperties      []*UserPropertyInput `json:"excluded_properties"`
	TrackProperties         []*UserPropertyInput `json:"track_properties"`
	ExcludedTrackProperties []*UserPropertyInput `json:"excluded_track_properties"`
	Environments            []*string            `json:"environments"`
	AppVersions             []*string            `json:"app_versions"`
	DateRange               *DateRangeInput      `json:"date_range"`
	LengthRange             *LengthRangeInput    `json:"length_range"`
	Os                      *string              `json:"os"`
	Browser                 *string              `json:"browser"`
	DeviceID                *string              `json:"device_id"`
	VisitedURL              *string              `json:"visited_url"`
	Referrer                *string              `json:"referrer"`
	Identified              *bool                `json:"identified"`
	HideViewed              *bool                `json:"hide_viewed"`
	FirstTime               *bool                `json:"first_time"`
	ShowLiveSessions        *bool                `json:"show_live_sessions"`
	Query                   *string              `json:"query"`
}

type SessionCommentTagInput struct {
	ID   *int   `json:"id"`
	Name string `json:"name"`
}

type SlackSyncResponse struct {
	Success               bool `json:"success"`
	NewChannelsAddedCount int  `json:"newChannelsAddedCount"`
}

type SocialLink struct {
	Type SocialType `json:"type"`
	Link *string    `json:"link"`
}

type SubscriptionDetails struct {
	BaseAmount      int64    `json:"baseAmount"`
	DiscountPercent float64  `json:"discountPercent"`
	DiscountAmount  int64    `json:"discountAmount"`
	LastInvoice     *Invoice `json:"lastInvoice"`
}

type TopUsersPayload struct {
	ID                   int     `json:"id"`
	Identifier           string  `json:"identifier"`
	TotalActiveTime      int     `json:"total_active_time"`
	ActiveTimePercentage float64 `json:"active_time_percentage"`
	UserProperties       string  `json:"user_properties"`
}

type TrackPropertyInput struct {
	ID    *int   `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type User struct {
	ID int `json:"id"`
}

type UserFingerprintCount struct {
	Count int64 `json:"count"`
}

type UserPropertyInput struct {
	ID    *int   `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ErrorState string

const (
	ErrorStateOpen     ErrorState = "OPEN"
	ErrorStateResolved ErrorState = "RESOLVED"
	ErrorStateIgnored  ErrorState = "IGNORED"
)

var AllErrorState = []ErrorState{
	ErrorStateOpen,
	ErrorStateResolved,
	ErrorStateIgnored,
}

func (e ErrorState) IsValid() bool {
	switch e {
	case ErrorStateOpen, ErrorStateResolved, ErrorStateIgnored:
		return true
	}
	return false
}

func (e ErrorState) String() string {
	return string(e)
}

func (e *ErrorState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ErrorState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ErrorState", str)
	}
	return nil
}

func (e ErrorState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type IntegrationType string

const (
	IntegrationTypeSlack  IntegrationType = "Slack"
	IntegrationTypeLinear IntegrationType = "Linear"
	IntegrationTypeZapier IntegrationType = "Zapier"
)

var AllIntegrationType = []IntegrationType{
	IntegrationTypeSlack,
	IntegrationTypeLinear,
	IntegrationTypeZapier,
}

func (e IntegrationType) IsValid() bool {
	switch e {
	case IntegrationTypeSlack, IntegrationTypeLinear, IntegrationTypeZapier:
		return true
	}
	return false
}

func (e IntegrationType) String() string {
	return string(e)
}

func (e *IntegrationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IntegrationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IntegrationType", str)
	}
	return nil
}

func (e IntegrationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MetricType string

const (
	MetricTypeWebVital MetricType = "WebVital"
	MetricTypeDevice   MetricType = "Device"
	MetricTypeBackend  MetricType = "Backend"
)

var AllMetricType = []MetricType{
	MetricTypeWebVital,
	MetricTypeDevice,
	MetricTypeBackend,
}

func (e MetricType) IsValid() bool {
	switch e {
	case MetricTypeWebVital, MetricTypeDevice, MetricTypeBackend:
		return true
	}
	return false
}

func (e MetricType) String() string {
	return string(e)
}

func (e *MetricType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MetricType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MetricType", str)
	}
	return nil
}

func (e MetricType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PlanType string

const (
	PlanTypeFree       PlanType = "Free"
	PlanTypeBasic      PlanType = "Basic"
	PlanTypeStartup    PlanType = "Startup"
	PlanTypeEnterprise PlanType = "Enterprise"
)

var AllPlanType = []PlanType{
	PlanTypeFree,
	PlanTypeBasic,
	PlanTypeStartup,
	PlanTypeEnterprise,
}

func (e PlanType) IsValid() bool {
	switch e {
	case PlanTypeFree, PlanTypeBasic, PlanTypeStartup, PlanTypeEnterprise:
		return true
	}
	return false
}

func (e PlanType) String() string {
	return string(e)
}

func (e *PlanType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PlanType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PlanType", str)
	}
	return nil
}

func (e PlanType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SessionCommentType string

const (
	SessionCommentTypeAdmin    SessionCommentType = "Admin"
	SessionCommentTypeFeedback SessionCommentType = "FEEDBACK"
)

var AllSessionCommentType = []SessionCommentType{
	SessionCommentTypeAdmin,
	SessionCommentTypeFeedback,
}

func (e SessionCommentType) IsValid() bool {
	switch e {
	case SessionCommentTypeAdmin, SessionCommentTypeFeedback:
		return true
	}
	return false
}

func (e SessionCommentType) String() string {
	return string(e)
}

func (e *SessionCommentType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SessionCommentType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SessionCommentType", str)
	}
	return nil
}

func (e SessionCommentType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SessionLifecycle string

const (
	SessionLifecycleAll       SessionLifecycle = "All"
	SessionLifecycleLive      SessionLifecycle = "Live"
	SessionLifecycleCompleted SessionLifecycle = "Completed"
)

var AllSessionLifecycle = []SessionLifecycle{
	SessionLifecycleAll,
	SessionLifecycleLive,
	SessionLifecycleCompleted,
}

func (e SessionLifecycle) IsValid() bool {
	switch e {
	case SessionLifecycleAll, SessionLifecycleLive, SessionLifecycleCompleted:
		return true
	}
	return false
}

func (e SessionLifecycle) String() string {
	return string(e)
}

func (e *SessionLifecycle) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SessionLifecycle(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SessionLifecycle", str)
	}
	return nil
}

func (e SessionLifecycle) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SocialType string

const (
	SocialTypeGithub   SocialType = "Github"
	SocialTypeLinkedIn SocialType = "LinkedIn"
	SocialTypeTwitter  SocialType = "Twitter"
	SocialTypeFacebook SocialType = "Facebook"
	SocialTypeSite     SocialType = "Site"
)

var AllSocialType = []SocialType{
	SocialTypeGithub,
	SocialTypeLinkedIn,
	SocialTypeTwitter,
	SocialTypeFacebook,
	SocialTypeSite,
}

func (e SocialType) IsValid() bool {
	switch e {
	case SocialTypeGithub, SocialTypeLinkedIn, SocialTypeTwitter, SocialTypeFacebook, SocialTypeSite:
		return true
	}
	return false
}

func (e SocialType) String() string {
	return string(e)
}

func (e *SocialType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SocialType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SocialType", str)
	}
	return nil
}

func (e SocialType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SubscriptionInterval string

const (
	SubscriptionIntervalMonthly SubscriptionInterval = "Monthly"
	SubscriptionIntervalAnnual  SubscriptionInterval = "Annual"
)

var AllSubscriptionInterval = []SubscriptionInterval{
	SubscriptionIntervalMonthly,
	SubscriptionIntervalAnnual,
}

func (e SubscriptionInterval) IsValid() bool {
	switch e {
	case SubscriptionIntervalMonthly, SubscriptionIntervalAnnual:
		return true
	}
	return false
}

func (e SubscriptionInterval) String() string {
	return string(e)
}

func (e *SubscriptionInterval) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SubscriptionInterval(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SubscriptionInterval", str)
	}
	return nil
}

func (e SubscriptionInterval) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
