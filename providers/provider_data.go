package providers

import (
	"net/url"
	"time"
)

type ProviderData struct {
	ProviderName      string
	ClientID          string
	ClientSecret      string
	LoginURL          *url.URL
	RedeemURL         *url.URL
	ProfileURL        *url.URL
	ProtectedResource *url.URL
	ValidateURL       *url.URL
	Scope             string
	ApprovalPrompt    string
	MaxAge            time.Duration
}

func (p *ProviderData) Data() *ProviderData { return p }
