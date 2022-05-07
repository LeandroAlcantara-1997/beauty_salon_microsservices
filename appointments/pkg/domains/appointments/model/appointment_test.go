package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var fakeAppointment = Appointment{
	UserID:          1,
	SalonID:         1,
	AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
}

var fakeAppointmentWithoutUserID = Appointment{
	SalonID:         1,
	AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
}

var fakeAppointmentWithoutSalonID = Appointment{
	UserID:          1,
	AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
}

func TestNewAppointment(t *testing.T) {
	type args struct {
		appointment UpsertAppointment
	}
	tests := []struct {
		name string
		args args
		want Appointment
	}{
		{
			name: "success, created appointment with all values",
			args: args{
				UpsertAppointment{
					UserID:          1,
					SalonID:         1,
					AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
				},
			},
			want: fakeAppointment,
		},
		{
			name: "success, created appointment without userid",
			args: args{
				UpsertAppointment{
					SalonID:         1,
					AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
				},
			},
			want: fakeAppointmentWithoutUserID,
		},
		{
			name: "success, created appointment without salonid",
			args: args{
				UpsertAppointment{
					UserID:          1,
					AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
				},
			},
			want: fakeAppointmentWithoutSalonID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAppointment(tt.args.appointment)
			assert.Equal(t, tt.want, got)
		})
	}
}
