package tool

import (
	"github.com/pinpoint-apm/pinpoint-go-agent"
	"go.uber.org/zap"
)

type PinPointConfig struct {
	ServerURL   string
	CounterRate int
	AgentPort   int
	StatPort    int
	SpanPort    int
}

func (c *PinPointConfig) AppendFieldMap(fMap map[string]string) {
	fMap["pinpoint.url"] = "PinPoint.ServerURL"
	fMap["pinpoint.sampling.counterRate"] = "PinPoint.CounterRate"
	fMap["pinpoint.agentPort"] = "PinPoint.AgentPort"
	fMap["pinpoint.statPort"] = "PinPoint.StatPort"
	fMap["pinpoint.spanPort"] = "PinPoint.SpanPort"
}

func (c *PinPointConfig) Print() {
	zap.L().Info("------------ pinpoint ------------")
	zap.L().Info("-- ", zap.String("collector.url", c.ServerURL))
	zap.L().Info("-- ", zap.Int("collector.agentPort", c.AgentPort))
	zap.L().Info("-- ", zap.Int("collector.statPort", c.StatPort))
	zap.L().Info("-- ", zap.Int("collector.spanPort", c.SpanPort))
	zap.L().Info("-- ", zap.Int("sampling.counterRate", c.CounterRate))
}

func (c *PinPointConfig) InitPinPoint(prj string, host string) pinpoint.Agent {
	opts := []pinpoint.ConfigOption{
		pinpoint.WithAppName(prj),
		pinpoint.WithAgentId(host),
		pinpoint.WithCollectorHost(c.ServerURL),
		pinpoint.WithCollectorAgentPort(c.AgentPort),
		pinpoint.WithCollectorStatPort(c.StatPort),
		pinpoint.WithCollectorSpanPort(c.SpanPort),
		pinpoint.WithSamplingCounterRate(c.CounterRate),
		pinpoint.WithLogLevel("Error"),
	}
	cfg, _ := pinpoint.NewConfig(opts...)
	agent, err := pinpoint.NewAgent(cfg)
	if err != nil {
		zap.L().Warn("pinpoint agent start: ", zap.String("failed", err.Error()))
		return nil
	}
	zap.L().Info("pinpoint connected: ", zap.String("to", c.ServerURL))
	return agent
}
