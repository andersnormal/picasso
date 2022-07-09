package action

import (
	"time"

	"github.com/sethvargo/go-githubactions"
)

type Config struct {
	Role          string
	LeaseDuration time.Duration
}

func NewFromInputs(action *githubactions.Action) (*Config, error) {
	lease := action.GetInput("lease-duration")
	d, err := time.ParseDuration(lease)
	if err != nil {
		return nil, err
	}

	c := Config{
		Role:          action.GetInput("role"),
		LeaseDuration: d,
	}
	return &c, nil
}
