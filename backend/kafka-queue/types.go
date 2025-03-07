package kafka_queue

import (
	"github.com/leonelquinteros/hubspot"
	"math"
	"time"

	"github.com/highlight-run/highlight/backend/clickhouse"
	"github.com/highlight-run/highlight/backend/model"

	customModels "github.com/highlight-run/highlight/backend/public-graph/graph/model"
	"github.com/segmentio/kafka-go"
)

type PayloadType = int

const (
	PushPayload                      PayloadType = iota
	InitializeSession                PayloadType = iota
	IdentifySession                  PayloadType = iota
	AddTrackProperties               PayloadType = iota // Deprecated: track events are now processed in pushPayload
	AddSessionProperties             PayloadType = iota
	PushBackendPayload               PayloadType = iota
	PushMetrics                      PayloadType = iota
	MarkBackendSetup                 PayloadType = iota
	AddSessionFeedback               PayloadType = iota
	PushLogs                         PayloadType = iota
	HubSpotCreateContactForAdmin     PayloadType = iota
	HubSpotCreateCompanyForWorkspace PayloadType = iota
	HubSpotUpdateContactProperty     PayloadType = iota
	HubSpotUpdateCompanyProperty     PayloadType = iota
	HealthCheck                      PayloadType = math.MaxInt
)

type PushPayloadArgs struct {
	SessionSecureID    string
	Events             customModels.ReplayEventsInput
	Messages           string
	Resources          string
	Errors             []*customModels.ErrorObjectInput
	IsBeacon           *bool
	HasSessionUnloaded *bool
	HighlightLogs      *string
	PayloadID          *int
}

type InitializeSessionArgs struct {
	SessionSecureID                string
	CreatedAt                      time.Time
	ProjectVerboseID               string
	EnableStrictPrivacy            bool
	EnableRecordingNetworkContents bool
	ClientVersion                  string
	FirstloadVersion               string
	ClientConfig                   string
	Environment                    string
	AppVersion                     *string
	Fingerprint                    string
	UserAgent                      string
	AcceptLanguage                 string
	IP                             string
	ClientID                       string
	NetworkRecordingDomains        []string
	DisableSessionRecording        *bool
}

type IdentifySessionArgs struct {
	SessionSecureID string
	UserIdentifier  string
	UserObject      interface{}
}
type AddTrackPropertiesArgs struct {
	SessionSecureID  string
	PropertiesObject interface{}
}

type AddSessionPropertiesArgs struct {
	SessionSecureID  string
	PropertiesObject interface{}
}
type PushBackendPayloadArgs struct {
	ProjectVerboseID *string
	SessionSecureID  *string
	Errors           []*customModels.BackendErrorObjectInput
}

type PushMetricsArgs struct {
	SessionSecureID string
	SessionID       int
	ProjectID       int
	Metrics         []*customModels.MetricInput
}

type MarkBackendSetupArgs struct {
	ProjectVerboseID *string
	SessionSecureID  *string
	ProjectID        int
	Type             model.MarkBackendSetupType
}

type AddSessionFeedbackArgs struct {
	SessionSecureID string
	UserName        *string
	UserEmail       *string
	Verbatim        string
	Timestamp       time.Time
}

type PushLogsArgs struct {
	LogRows []*clickhouse.LogRow
}

type HubSpotCreateContactForAdminArgs struct {
	AdminID            int
	Email              string
	UserDefinedRole    string
	UserDefinedPersona string
	First              string
	Last               string
	Phone              string
	Referral           string
}

type HubSpotCreateCompanyForWorkspaceArgs struct {
	WorkspaceID int
	AdminEmail  string
	Name        string
}

type HubSpotUpdateContactPropertyArgs struct {
	AdminID    int
	Properties []hubspot.Property
}

type HubSpotUpdateCompanyPropertyArgs struct {
	WorkspaceID int
	Properties  []hubspot.Property
}

type Message struct {
	Type                             PayloadType
	Failures                         int
	MaxRetries                       int
	KafkaMessage                     *kafka.Message
	PushPayload                      *PushPayloadArgs
	InitializeSession                *InitializeSessionArgs
	IdentifySession                  *IdentifySessionArgs
	AddTrackProperties               *AddTrackPropertiesArgs
	AddSessionProperties             *AddSessionPropertiesArgs
	PushBackendPayload               *PushBackendPayloadArgs
	PushMetrics                      *PushMetricsArgs
	MarkBackendSetup                 *MarkBackendSetupArgs
	AddSessionFeedback               *AddSessionFeedbackArgs
	PushLogs                         *PushLogsArgs
	HubSpotCreateContactForAdmin     *HubSpotCreateContactForAdminArgs
	HubSpotCreateCompanyForWorkspace *HubSpotCreateCompanyForWorkspaceArgs
	HubSpotUpdateContactProperty     *HubSpotUpdateContactPropertyArgs
	HubSpotUpdateCompanyProperty     *HubSpotUpdateCompanyPropertyArgs
}

type PartitionMessage struct {
	Message   *Message
	Partition int32
}
