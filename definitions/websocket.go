package definitions

type WebsocketAction string

const (
	WebsocketActionConnect WebsocketAction = "connect"
)

type EmbedWebsocketRequest struct {
	Action WebsocketAction `json:"action"`
	Token  *string         `json:"token,omitempty"`
}
