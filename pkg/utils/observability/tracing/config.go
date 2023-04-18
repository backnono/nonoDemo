package tracing

type Config struct {
	ServiceName  string  `yaml:"service_name" mapstructure:"service_name"`
	SamplingRate float64 `yaml:"sampling_rate" mapstructure:"sampling_rate"`
	AgentHost    string  `yaml:"agent_host" mapstructure:"agent_host"`
	AgentPort    string  `yaml:"agent_port" mapstructure:"agent_port"`
}
