package entities

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewAppointment(t *testing.T) {
	t.Run("should be able to create a new appointment", func(t *testing.T) {
		appointment, _ := NewAppointment("111", time.Now())

		require.Equal(t, appointment.ProviderId, appointment.ProviderId)
	})

	t.Run("should be not able to create a new appointment", func(t *testing.T) {
		_, err := NewAppointment("", time.Now())

		require.NotNil(t, err)
	})
}
