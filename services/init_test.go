package services

import (
	"testing"

	"github.com/cloudfoundry/cf-acceptance-tests/helpers"
)

func TestServices(t *testing.T) {
	helpers.GinkgoBootstrap(t, "Services")
}