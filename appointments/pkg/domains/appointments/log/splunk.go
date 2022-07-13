package log

import (
	"time"

	"github.com/ZachtimusPrime/Go-Splunk-HTTP/splunk/v2"
)

type SplunkLog struct {
	Splunk     *splunk.Client
	Source     string
	SourceType string
	Index      string
}

func NewSplunkLog(s *splunk.Client, src, srcType, index string) *SplunkLog {
	return &SplunkLog{
		Splunk: s,
	}
}

func (d *SplunkLog) Log(data interface{}) error {
	event := d.Splunk.NewEvent(data, d.Source, d.SourceType, d.Index)
	if err := d.Splunk.Log(event); err != nil {
		return err
	}
	return nil
}

func (d *SplunkLog) LogWithTime(data interface{}) error {
	event := d.Splunk.NewEvent(data, d.Source, d.SourceType, d.Index)
	if err := d.Splunk.LogWithTime(time.Now(), event); err != nil {
		return err
	}
	return nil
}
