package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var fakeAppointmentReponse = AppResponse{
	UserID:          1,
	SalonID:         1,
	AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
}

var fakeAppResponsessWithoutUserID = AppResponse{
	SalonID:         1,
	AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
}

var fakeUpsertAppResponseWithoutSalonID = AppResponse{
	UserID:          1,
	AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
}

func TestNewAppResponse(t *testing.T) {
	type args struct {
		appointment Appointment
	}
	tests := []struct {
		name string
		args args
		want AppResponse
	}{
		{
			name: "success, created appResponse with all values",
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
			want: fakeAppResponsessWithoutUserID,
		},
		{
			name: "success, created appResponse without salonid",
			args: args{
				Appointment{
					UserID:          1,
					AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
				},
			},
			want: fakeUpsertAppResponseWithoutSalonID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAppResponse(tt.args.appointment)
			assert.Equal(t, tt.want, got)
		})
	}
}
