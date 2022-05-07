package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var fakeAppointmentReponse = AppointmentResponse{
	UserID:          1,
	SalonID:         1,
	AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
}

var fakeAppointmentResponsessWithoutUserID = AppointmentResponse{
	SalonID:         1,
	AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
}

var fakeUpsertAppointmentResponseWithoutSalonID = AppointmentResponse{
	UserID:          1,
	AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
}

func TestNewAppointmentResponse(t *testing.T) {
	type args struct {
		appointment Appointment
	}
	tests := []struct {
		name string
		args args
		want AppointmentResponse
	}{
		{
			name: "success, created appointmentResponse with all values",
			args: args{
				Appointment{
					UserID:          1,
					SalonID:         1,
					AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
				},
			},
			want: fakeAppointmentReponse,
		},
		{
			name: "success, created appointmentReponse without userid",
			args: args{
				Appointment{
					SalonID:         1,
					AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
				},
			},
			want: fakeAppointmentResponsessWithoutUserID,
		},
		{
			name: "success, created appointmentResponse without salonid",
			args: args{
				Appointment{
					UserID:          1,
					AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
				},
			},
			want: fakeUpsertAppointmentResponseWithoutSalonID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAppointmentResponse(tt.args.appointment)
			assert.Equal(t, tt.want, got)
		})
	}
}
