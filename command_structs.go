package ari

// https://docs.asterisk.org/Asterisk_20_Documentation/API_Documentation/Asterisk_REST_Interface/Asterisk_REST_Data_Models/

type Application struct {
	BridgeIds     []string `json:"bridge_ids"`
	ChannelIds    []string `json:"channel_ids"`
	DeviceNames   []string `json:"device_names"`
	EndpointIds   []string `json:"endpoint_ids"`
	EventsAllowed []struct {
		Type string `json:"type"`
	} `json:"events_allowed"`
	EventsDisallowed []struct {
		Type string `json:"type"`
	} `json:"events_disallowed"`
	Name string `json:"name"`
}

type AsteriskInfo struct {
	Build  BuildInfo  `json:"build"`
	Config ConfigInfo `json:"config"`
	Status StatusInfo `json:"status"`
	System SystemInfo `json:"system"`
}

type AsteriskPing struct {
	AsteriskId string `json:"asterisk_id"`
	Ping       string `json:"ping"`
	Timestamp  string `json:"timestamp"`
}
type Bridge struct {
	BridgeClass   string   `json:"bridge_class"`
	BridgeType    string   `json:"bridge_type"`
	Channels      []string `json:"channels"`
	CreationTime  string   `json:"creationtime"`
	Creator       string   `json:"creator"`
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	Technology    string   `json:"technology"`
	VideoMode     string   `json:"video_mode"`
	VideoSourceId string   `json:"video_source_id"`
}

type BuildInfo struct {
	Date    string `json:"date"`
	Kernel  string `json:"kernel"`
	Machine string `json:"machine"`
	Options string `json:"options"`
	Os      string `json:"os"`
	User    string `json:"user"`
}

type CallerID struct {
	Name   string `json:"name"`
	Number string `json:"number"`
}

type Channel struct {
	AccountCode  string            `json:"accountcode"`
	Caller       CallerID          `json:"caller"`
	CallerRdnis  string            `json:"caller_rdnis"`
	ChannelVars  map[string]string `json:"channelvars"` //unclear in documentation
	Connected    CallerID          `json:"connected"`
	CreationTime string            `json:"creationtime"`
	Dialplan     DialplanCEP       `json:"dialplan"`
	Id           string            `json:"id"`
	Language     string            `json:"language"`
	Name         string            `json:"name"`
	ProtocolId   string            `json:"protocol_id"`
	State        string            `json:"state"`
	TenantId     string            `json:"tenantid"`
}

type ConfigInfo struct {
	DefaultLanguage string `json:"default_language"`
	MaxChannels     int    `json:"max_channels"`
	Max_load        int    `json:"max_load"`
	Max_open_files  int    `json:"max_open_files"`
	Name            string `json:"name"`
	Setid           SetId  `json:"setid"`
}

type ConfigTuple struct {
	Attribute string `json:"attribute"`
	Value     string `json:"value"`
}

type DeviceState struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

// Not used anywhere, but it's in the docs...
type Dialed struct{}

type DialplanCEP struct {
	AppData  string `json:"app_data"`
	AppName  string `json:"app_name"`
	Context  string `json:"context"`
	Exten    string `json:"exten"`
	Priority uint64 `json:"priority"`
}
type Endpoint struct {
	ChannelIds []string `json:"channel_ids"`
	Resource   string   `json:"resource"`
	State      string   `json:"state"`
	Technology string   `json:"technology"`
}
type FormatLangPair struct {
	Format   string `json:"format"`
	Language string `json:"language"`
}
type LiveRecording struct {
	Cause           string `json:"cause"`
	Duration        string `json:"duration"`
	Format          string `json:"format"`
	Name            string `json:"name"`
	SilenceDuration int    `json:"silence_duration"`
	State           string `json:"state"`
	TalkingDuration int    `json:"talking_duration"`
	TargetUri       string `json:"target_uri"`
}
type LogChannel struct {
	Channel       string `json:"channel"`
	Configuration string `json:"configuration"`
	Status        string `json:"status"`
	Type          string `json:"type"`
}
type Mailbox struct {
	Name        string `json:"name"`
	NewMessages int    `json:"new_messages"`
	OldMessages int    `json:"old_messages"`
}

type MissingParams struct {
	AsteriskId string   `json:"asterisk_id"`
	Type       string   `json:"type"`
	Params     []string `json:"params"`
	/*
		The ARI base model "Message" has not been included in this spec.
		MissingParams is the only model that uses the Message base,
		so the fields have been included here directly.
	*/
}

type Module struct {
	Description   string `json:"description"`
	Name          string `json:"name"`
	Status        string `json:"status"`
	Support_level string `json:"support_level"`
	UseCount      int    `json:"use_count"`
}

type Peer struct {
	Address    string
	Cause      string
	PeerStatus string
	Port       string
	Time       string
}

type Playback struct {
	Id           string `json:"id"`
	Language     string `json:"language"`
	MediaUri     string `json:"media_uri"`
	NextMediaUri string `json:"next_media_uri"`
	State        string `json:"state"`
	TargetUri    string `json:"target_uri"`
}
type RTPstat struct {
	Channel_uniqueid       string
	Local_maxjitter        float64
	Local_maxrxploss       float64
	Local_minjitter        float64
	Local_minrxploss       float64
	Local_normdevjitter    float64
	Local_normdevrxploss   float64
	Local_ssrc             int
	Local_stdevjitter      float64
	Local_stdevrxploss     float64
	Maxrtt                 float64
	Minrtt                 float64
	Normdevrtt             float64
	Remote_maxjitter       float64
	Remote_maxrxploss      float64
	Remote_minjitter       float64
	Remote_minrxploss      float64
	Remote_normdevjitter   float64
	Remmote_normdevrxploss float64
	Remote_ssrc            int
	Remote_stdevjitter     float64
	Remote_stdevrxploss    float64
	Rtt                    float64
	Rxcount                int
	Rxjitter               float64
	Rxoctetcount           int
	Rxploss                int
	Stdevrtt               float64
	Txcount                int
	Txjitter               float64
	Txoctetcount           int
	Txploss                int
}

type SetId struct {
	Group string `json:"group"`
	User  string `json:"user"`
}

type Sound struct {
	Formats []FormatLangPair `json:"formats"`
	Id      string           `json:"id"`
	Text    string           `json:"text"`
}

type StatusInfo struct {
	LastReloadTime string `json:"last_reload_time"`
	StartupTime    string `json:"startup_time"`
}

type StoredRecording struct {
	Format string `json:"format"`
	Name   string `json:"name"`
}

type SystemInfo struct {
	EntityId string `json:"entity_id"`
	Version  string `json:"version"`
}

type TextMessage struct {
	Body      string                `json:"body"`
	From      string                `json:"from"`
	To        string                `json:"to"`
	Variables []TextMessageVariable `json:"variables"` // inferred, documentation not specific
}

type TextMessageVariable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Variable struct {
	// this might be [2]string, for a dynamic key/value scheme
	// but how can that be transformed to a json object?
	Value string `json:"value"`
}
