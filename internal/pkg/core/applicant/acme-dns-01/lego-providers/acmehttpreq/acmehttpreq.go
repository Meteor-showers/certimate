﻿package acmehttpreq

import (
	"errors"
	"net/url"
	"time"

	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/providers/dns/httpreq"
)

type ACMEHttpReqApplicantConfig struct {
	Endpoint           string `json:"endpoint"`
	Mode               string `json:"mode"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	PropagationTimeout int32  `json:"propagationTimeout,omitempty"`
}

func NewChallengeProvider(config *ACMEHttpReqApplicantConfig) (challenge.Provider, error) {
	if config == nil {
		return nil, errors.New("config is nil")
	}

	endpoint, _ := url.Parse(config.Endpoint)
	providerConfig := httpreq.NewDefaultConfig()
	providerConfig.Endpoint = endpoint
	providerConfig.Mode = config.Mode
	providerConfig.Username = config.Username
	providerConfig.Password = config.Password
	if config.PropagationTimeout != 0 {
		providerConfig.PropagationTimeout = time.Duration(config.PropagationTimeout) * time.Second
	}

	provider, err := httpreq.NewDNSProviderConfig(providerConfig)
	if err != nil {
		return nil, err
	}

	return provider, nil
}
