package log

//go:generate mockgen -destination log_mock.go -package=log -source=log.go
type AppointmentLogI interface {
	Log(event interface{}) error
	LogWithTime(event interface{}) error
}
