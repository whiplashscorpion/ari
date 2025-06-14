package ari

import ()

// The base model for all websocket events
type Event struct {
	AsteriskId  string `json:"asterisk_id"`
	Type        string `json:"type"`
	Application string `json:"application"`
	Timestamp   string `json:"timestamp"`
}

type ApplicationMoveFailed struct {
	Event
	Args        []string `json:"args"`
	Channel     Channel  `json:"channel"`
	Destination string   `json:"destination"`
}

type ApplicationReplaced struct {
	Event
}

// Bridges
type BridgeAttendedTransfer struct {
	Event
	DestinationApplication     string  `json:"destination_application"`
	DestinationBridge          string  `json:"destination_bridge"`
	DestinationLinkFirstLeg    Channel `json:"destination_link_first_leg"`
	DestinationLinkSecondLeg   Channel `json:"destination_link_second_leg"`
	DestinationThreewayBridge  Bridge  `json:"destination_threeway_bridge"`
	DestinationThreewayChannel Channel `json:"destination_threeway_channel"`
	DestinationType            string  `json:"destination_type"`
	IsExternal                 bool    `json:"is_external"`
	ReplaceChannel             Channel `json:"replace_channel"`
	Result                     string  `json:"result"`
	TransferTarget             Channel `json:"transfer_target"`
	Transferee                 Channel `json:"transferee"`
	TransfererFirstLeg         Channel `json:"transferer_first_leg"`
	TransfererFirstLegBridge   Bridge  `json:"transferer_first_leg_bridge"`
	TransfererSecondLeg        Channel `json:"transferer_second_leg"`
	TransfererSecondLegBridge  Bridge  `json:"transferer_second_leg_bridge"`
}

type BridgeBlindTransfer struct {
	Event
	Bridge         Bridge  `json:"bridge"`
	Channel        Channel `json:"channel"`
	Context        string  `json:"context"`
	Exten          string  `json:"exten"`
	IsExternal     bool    `json:"is_external"`
	ReplaceChannel Channel `json:"replace_channel"`
	Result         string  `json:"result"`
	Transferee     Channel `json:"transferee"`
}

type BridgeCreated struct {
	Event
	Bridge Bridge `json:"bridge"`
}

type BridgeDestroyed struct {
	Event
	Bridge Bridge `json:"bridge"`
}

type BridgeMerged struct {
	Event
	Bridge     Bridge `json:"bridge"`
	BridgeFrom Bridge `json:"bridge_from"`
}

type BridgeVideoSourceChanged struct {
	Event
	Bridge           Bridge `json:"bridge"`
	OldVideoSourceId string `json:"old_video_source_id"`
}

// Channels
type ChannelCallerId struct {
	Event
	CallerPresentation    int     `json:"caller_presentation"`
	CallerPresentationTxt string  `json:"caller_presentation_txt"`
	Channel               Channel `json:"channel"`
}

type ChannelConnectedLine struct {
	Event
	Channel Channel `json:"channel"`
}

type ChannelCreated struct {
	Event
	Channel Channel `json:"channel"`
}

type ChannelDestroyed struct {
	Event
	Cause    int     `json:"cause"`
	CauseTxt string  `json:"cause_txt"`
	Channel  Channel `json:"channel"`
}

type ChannelDialplan struct {
	Event
	Channel         Channel `json:"channel"`
	DialplanApp     string  `json:"dialplan_app"`
	DialplanAppData string  `json:"diaplan_app_data"`
}

type ChannelDtmfReceived struct {
	Event
	Channel    Channel `json:"channel"`
	Digit      string  `json:"digit"`
	DurationMs int     `json:"duration_ms"`
}

type ChannelEnteredBridge struct {
	Event
	Bridge  Bridge  `json:"bridge"`
	Channel Channel `json:"channel"`
}

type ChannelHangupRequest struct {
	Event
	Cause   int     `json:"cause"`
	Channel Channel `json:"channel"`
	Soft    bool    `json:"soft"`
}

type ChannelHold struct {
	Event
	Channel    Channel `json:"channel"`
	MusicClass string  `json:"musicclass"`
}

type ChannelLeftBridge struct {
	Event
	Bridge  Bridge  `json:"bridge"`
	Channel Channel `json:"channel"`
}

type ChannelStateChange struct {
	Event
	Channel Channel `json:"channel"`
}

type ChannelTalkingFinished struct {
	Event
	Channel  Channel `json:"channel"`
	Duration int     `json:"duration"`
}

type ChannelTalkingStarted struct {
	Event
	Channel Channel `json:"channel"`
}

type ChannelTransfer struct {
	Event
	ReferTo    ReferTo    `json:"refer_to"`
	ReferredBy ReferredBy `json:"referred_by"`
	State      string     `json:"state"`
}

// Refer types are not events
// Refer types only used by ChannelTransfer, so defined here
type ReferTo struct {
	Bridge               Bridge              `json:"bridge"`
	ConnectedChannel     Channel             `json:"connected_channel"`
	DestinationChannel   Channel             `json:"destination_channel"`
	RequestedDestination RequiredDestination `json:"requested_destination"`
}

type ReferredBy struct {
	Bridge           Bridge  `json:"bridge"`
	ConnectedChannel Channel `json:"channel"`
	SourceChannel    Channel `json:"source_channel"`
}

type RequiredDestination struct {
	AdditionalProtocolParams []struct {
		ParameterName  string `json:"parameter_name"`
		ParameterValue string `json:"parameter_value"`
	} `json:"additional_protocol_params"`
	Destination string `json:"destination"`
	ProtocolId  string `json:"protocol_id"`
}

// end Refer

type ChannelUnhold struct {
	Event
	Channel Channel `json:"channel"`
}

type ChannelUserEvent struct {
	Event
	Bridge    Bridge   `json:"bridge"`
	Channel   Channel  `json:"channel"`
	Endpoint  Endpoint `json:"endpoint"`
	EventName string   `json:"eventname"`
	UserEvent any      `json:"userevent"` // unknown type
}

type ChannelVarset struct {
	Event
	Channel  Channel `json:"channel"`
	Variable string  `json:"variable"`
	Value    string  `json:"value"`
}

// Contacts
type ContactStatusChange struct {
	Event
	ContactInfo ContactInfo `json:"contact_info"`
	Endpoint    Endpoint    `json:"endpoint"`
}

// Not an event, only used in ContactStatusChange
type ContactInfo struct {
	Aor           string `json:"aor"`
	ContactStatus string `json:"contact_status"`
	RoundtripUsec string `json:"roundtrip_usec"`
	Uri           string `json:"uri"`
}

/*
Devices
*/
type DeviceStateChanged struct {
	Event
	DeviceState DeviceState `json:"device_state"`
}

/*
Dial
*/
type Dial struct {
	Event
	Caller     Channel `json:"caller"`
	Dialstatus string  `json:"dialstatus"`
	Dialstring string  `json:"dialstring"`
	Forward    string  `json:"forward"`
	Forwarded  Channel `json:"forwarded"`
	Peer       Channel `json:"peer"`
}

/*
Endpoints
*/
type EndpointStateChange struct {
	Event
	Endpoint Endpoint `json:"endpoint"`
}

/*
PeerStatus
*/
type PeerStatusChange struct {
	Event
	Endpoint Endpoint
	Peer     Peer
}

// Playbacks
// Includes PlaybackStarted, PlaybackContinuing, PlaybackFinished
type PlaybackEvent struct {
	Event
	Playback Playback
}

// Recordings
type RecordingStarted struct {
	Event
	Recording LiveRecording
}
type RecordingFinished struct {
	Event
	Recording LiveRecording
}
type RecordingFailed struct {
	Event
	Recording LiveRecording
}

// Stasis
type StasisStart struct {
	Event
	Args           []string `json:"args"`
	Channel        Channel  `json:"channel"`
	ReplaceChannel Channel  `json:"replace_channel"`
}

type StasisEnd struct {
	Event
	Channel Channel `json:"channel"`
}

// TextMessage
type TextMessageReceived struct {
	Event
	Endpoint Endpoint    `json:"endpoint"`
	Message  TextMessage `json:"message"`
}
