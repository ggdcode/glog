package glog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type glog struct {
	l *zap.Logger
	f []zap.Field
}

func New(opt *Options) Logger {
	l := zap.New(
		zapcore.NewCore(opt.getEncoder(), opt.getWriteSyncer(), opt.getLevel()),
		opt.getAddCaller()...,
	)

	return &glog{
		l: l,
	}
}

func (g *glog) AddCallerSkip(n int) Logger {
	l := g.l.WithOptions(zap.AddCallerSkip(n))
	return &glog{l: l, f: g.f}
}

func (g *glog) WithField(key string, val interface{}) Logger {
	f := append(g.f, zap.Any(key, val))
	return &glog{l: g.l, f: f}
}

func (g *glog) WithFields(key string, val interface{}, kvs ...interface{}) Logger {
	f := append(g.f, zap.Any(key, val))

	for i := 0; i < len(kvs)-1; i += 2 {
		key, ok := kvs[i].(string)
		if ok {
			f = append(f, zap.Any(key, kvs[i+1]))
		}
	}

	return &glog{l: g.l, f: f}
}

func (g *glog) Data(val interface{}) Logger     { return g.WithField("data", val) }
func (g *glog) Action(val string) Logger        { return g.WithField("action", val) }
func (g *glog) Request(val interface{}) Logger  { return g.WithField("request", val) }
func (g *glog) Response(val interface{}) Logger { return g.WithField("response", val) }

func (g *glog) Debug(msg string) { g.l.Debug(msg, g.f...) }
func (g *glog) Info(msg string)  { g.l.Info(msg, g.f...) }
func (g *glog) Warn(msg string)  { g.l.Warn(msg, g.f...) }
func (g *glog) Error(msg string) { g.l.Error(msg, g.f...) }
func (g *glog) Fatal(msg string) { g.l.Fatal(msg, g.f...) }
