package tito

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

)

// parse errors
var (
	ErrEventNotSpecifiedToParse  = errors.New("no Event specified to parse")
	ErrInvalidHTTPMethod         = errors.New("invalid HTTP Method")
	ErrMissingWebhookName        = errors.New("missing X-Webhook-Name")
	ErrMissingSignatureHeader    = errors.New("missing Tito-Signature in header")	
	ErrEventNotFound             = errors.New("event not defined to be parsed")
	ErrParsingPayload            = errors.New("error parsing payload")
	ErrHMACVerificationFailed    = errors.New("HMAC verification failed")
)

//"Name", "Payload"
// Webhooks allow data to be passed to an external system in real time. Webhooks will POST a JSON payload to the endpoint you specify, with a X-Webhook-Name header.
// go get gopkg.in/yaml.v2

type Event string

// Tito hook types
const (
	CheckinCreated        Event = "checkin.created"
	TicketCreated         Event = "ticket.created"
	TicketCompleted       Event = "ticket.completed"
	TicketReassigned      Event = "ticket.reassigned"
	TicketUpdated         Event = "ticket.updated"
	TicketUnsnoozed       Event = "ticket.unsnoozed"
	TicketUnvoided        Event = "ticket.unvoided"
	TicketVoided          Event = "ticket.voided"
	RegistrationStarted   Event = "registration.started"
	RegistrationFilling   Event = "registration.filling"
	RegistrationUpdated   Event = "registration.updated"
	RegistrationFinished  Event = "registration.finished"
	RegistrationCompleted Event = "registration.completed"
	RegistrationCancelled Event = "registration.cancelled"
)



// Option is a configuration option for the webhook
type Option func(*Webhook) error

// Options is a namespace var for configuration options
var Options = WebhookOptions{}

// WebhookOptions is a namespace for configuration option methods
type WebhookOptions struct{}

// Secret registers the GitHub secret
func (WebhookOptions) Secret(secret string) Option {
	return func(hook *Webhook) error {
		hook.secret = secret
		return nil
	}
}

// Webhook instance contains all methods needed to process events
type Webhook struct {
	secret string
}

// New creates and returns a WebHook instance denoted by the Provider type
func New(options ...Option) (*Webhook, error) {
	hook := new(Webhook)
	for _, opt := range options {
		if err := opt(hook); err != nil {
			return nil, errors.New("Error applying Option")
		}
	}
	return hook, nil
}

// Parse verifies and parses the events specified and returns the payload object or an error
func (hook Webhook) Parse(r *http.Request, events ...Event) (interface{}, error) {
	defer func() {
		_, _ = io.Copy(ioutil.Discard, r.Body)
		_ = r.Body.Close()
	}()

	if len(events) == 0 {
		return nil, ErrEventNotSpecifiedToParse
	}
	if r.Method != http.MethodPost {
		return nil, ErrInvalidHTTPMethod
	}

	event := r.Header.Get("X-Webhook-Name")
	if event == "" {
		return nil, ErrMissingWebhookName
	}
	titoEvent := Event(event)

	var found bool
	for _, evt := range events {
		if evt == titoEvent {
			found = true
			break
		}
	}
	// event not defined to be parsed
	if !found {
		return nil, ErrEventNotFound
	}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, ErrParsingPayload
	}

	// If we have a Secret set, we should check the MAC
	if len(hook.secret) > 0 {
		signature := r.Header.Get("Tito-Signature")
		if len(signature) == 0 {
			return nil, ErrMissingSignatureHeader
		}
		mac := hmac.New(sha256.New, []byte(hook.secret))
		_, _ = mac.Write(payload)
		expectedMAC := base64.StdEncoding.EncodeToString(mac.Sum(nil))

		if !hmac.Equal([]byte(signature[5:]), []byte(expectedMAC)) {
			return nil, ErrHMACVerificationFailed
		}
	}

	switch titoEvent {
	case CheckinCreated:
		var pl CheckinCreatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case TicketCreated:
		var pl TicketCreatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case TicketCompleted:
		var pl TicketCompletedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case TicketReassigned:
		var pl TicketReassignedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case TicketUpdated:
		var pl TicketUpdatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case TicketUnsnoozed:
		var pl TicketUnsnoozedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case TicketUnvoided:
		var pl TicketUnvoidedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case TicketVoided:
		var pl TicketVoidedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case RegistrationStarted:
		var pl RegistrationStartedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case RegistrationFilling:
		var pl RegistrationFillingPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case RegistrationUpdated:
		var pl RegistrationUpdatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case RegistrationFinished:
		var pl RegistrationFinishedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case RegistrationCompleted:
		var pl RegistrationCompletedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case RegistrationCancelled:
		var pl RegistrationCancelledPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	default:
		return nil, fmt.Errorf("unknown event %s", titoEvent)
	}
}

