# Missing ARI Functions

## Summary
- **Total Missing Functions: 45**
- **Complete APIs:** Applications, Bridges, Playbacks
- **Partial APIs:** Channels (48% complete), Asterisk (27% complete), Endpoints (29% complete)
- **Missing APIs:** Device States, Events, Mailboxes, Recordings, Sounds

## Asterisk API (11 missing)

### Config Dynamic
- `AsteriskConfigDynamicGet` - GET `/asterisk/config/dynamic/{configClass}/{objectType}/{id}`
- `AsteriskConfigDynamicUpdate` - PUT `/asterisk/config/dynamic/{configClass}/{objectType}/{id}`
- `AsteriskConfigDynamicDelete` - DELETE `/asterisk/config/dynamic/{configClass}/{objectType}/{id}`

### Modules
- `AsteriskModulesList` - GET `/asterisk/modules`
- `AsteriskModulesGet` - GET `/asterisk/modules/{moduleName}`
- `AsteriskModulesLoad` - POST `/asterisk/modules/{moduleName}`
- `AsteriskModulesReload` - PUT `/asterisk/modules/{moduleName}`
- `AsteriskModulesUnload` - DELETE `/asterisk/modules/{moduleName}`

### Logging
- `AsteriskLoggingGet` - GET `/asterisk/logging`
- `AsteriskLoggingAdd` - POST `/asterisk/logging/{logChannelName}`
- `AsteriskLoggingRotate` - PUT `/asterisk/logging/{logChannelName}/rotate`
- `AsteriskLoggingDelete` - DELETE `/asterisk/logging/{logChannelName}`

## Channels API (16 missing)

- `ChannelContinue` - POST `/channels/{channelId}/continue`
- `ChannelRedirect` - POST `/channels/{channelId}/redirect`
- `ChannelDtmf` - POST `/channels/{channelId}/dtmf`
- `ChannelHold` - POST `/channels/{channelId}/hold`
- `ChannelUnhold` - DELETE `/channels/{channelId}/hold`
- `ChannelMoh` - POST `/channels/{channelId}/moh`
- `ChannelMohStop` - DELETE `/channels/{channelId}/moh`
- `ChannelSilence` - POST `/channels/{channelId}/silence`
- `ChannelSilenceStop` - DELETE `/channels/{channelId}/silence`
- `ChannelVariableGet` - GET `/channels/{channelId}/variable`
- `ChannelVariableSet` - POST `/channels/{channelId}/variable`
- `ChannelSnoop` - POST `/channels/{channelId}/snoop`
- `ChannelDial` - POST `/channels/{channelId}/dial`
- `ChannelRtpStatistics` - GET `/channels/{channelId}/rtp_statistics`
- `ChannelExternalMedia` - POST `/channels/externalMedia`
- `ChannelTransferProgress` - POST `/channels/{channelId}/transfer_progress`

## Device States API (4 missing)

- `DeviceStatesList` - GET `/deviceStates`
- `DeviceStatesGet` - GET `/deviceStates/{deviceName}`
- `DeviceStatesUpdate` - PUT `/deviceStates/{deviceName}`
- `DeviceStatesDelete` - DELETE `/deviceStates/{deviceName}`

## Endpoints API (5 missing)

- `EndpointsListByTech` - GET `/endpoints/{tech}`
- `EndpointsGet` - GET `/endpoints/{tech}/{resource}`
- `EndpointSendMessage` - PUT `/endpoints/{tech}/{resource}/sendMessage`
- `EndpointsRefer` - POST `/endpoints/refer`
- `EndpointRefer` - POST `/endpoints/{tech}/{resource}/refer`

## Events API (2 missing)

- `EventsWebSocket` - GET `/events` (WebSocket connection)
- `EventsUserEvent` - POST `/events/user/{eventName}`

## Mailboxes API (4 missing)

- `MailboxesList` - GET `/mailboxes`
- `MailboxesGet` - GET `/mailboxes/{mailboxName}`
- `MailboxesUpdate` - PUT `/mailboxes/{mailboxName}`
- `MailboxesDelete` - DELETE `/mailboxes/{mailboxName}`

## Recordings API (12 missing)

### Stored Recordings
- `RecordingsStoredList` - GET `/recordings/stored`
- `RecordingsStoredGet` - GET `/recordings/stored/{recordingName}`
- `RecordingsStoredDelete` - DELETE `/recordings/stored/{recordingName}`
- `RecordingsStoredFile` - GET `/recordings/stored/{recordingName}/file`
- `RecordingsStoredCopy` - POST `/recordings/stored/{recordingName}/copy`

### Live Recordings
- `RecordingsLiveGet` - GET `/recordings/live/{recordingName}`
- `RecordingsLiveDelete` - DELETE `/recordings/live/{recordingName}`
- `RecordingsLiveStop` - POST `/recordings/live/{recordingName}/stop`
- `RecordingsLivePause` - POST `/recordings/live/{recordingName}/pause`
- `RecordingsLiveUnpause` - DELETE `/recordings/live/{recordingName}/pause`
- `RecordingsLiveMute` - POST `/recordings/live/{recordingName}/mute`
- `RecordingsLiveUnmute` - DELETE `/recordings/live/{recordingName}/mute`

## Sounds API (2 missing)

- `SoundsList` - GET `/sounds`
- `SoundsGet` - GET `/sounds/{soundId}`