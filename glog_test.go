package glog_test

import (
	"testing"

	"github.com/ggdcode/glog"
)

func TestGLog(t *testing.T) {
	out(&glog.Options{})
	out(&glog.Options{
		Format:     glog.FormatText,
		TimeFormat: "2006-01-02 15:04:05",
	})
	out(&glog.Options{
		Output:     "file",
		TimeFormat: "epoch_time",
		Level:      "info",

		FileName: "1.log",
		Size:     1024,
	})
}

func TestDefault(t *testing.T) {
	out2(glog.GetDefaultLog())

	l := glog.New(&glog.Options{
		Format: glog.FormatText,
	})

	glog.SetDefaultLog(l.WithField("app_name", "glog"))
	out2(glog.GetDefaultLog())
}

func out(opt *glog.Options) {
	l := glog.New(opt)
	out2(l)
}

func out2(l glog.Logger) {
	l.Debug("debug")
	l.Info("info")
	l.Warn("warn")
	l.Error("error")

	l0 := l.WithField("wf0", "va0")
	l1 := l.WithField("wf1", "va1")

	l1.Info("l1.Info")
	l0.Info("l0.Info")
	l.Info("l.Info")

	l.WithFields("p0", "v0", "p1", 1, "p2", 2, 3).
		Error("ok")
}
