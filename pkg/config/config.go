/*
*

	@author: junwang
	@since: 2023/8/15
	@desc: //TODO

*
*/
package myconfig

type Config struct {
	LoadBalance `yaml:"loadBalance"`
}

type (
	LoadBalance struct {
		Specify    bool   `yaml:"specify" env:"SPECIFY"`
		Channel    string `env-required:"true" yaml:"channel" env:"CHANNEL"`
		SelectMode `env-required:"true" yaml:"selectMode" env:"SELECT_MODE"`
	}
	SelectMode struct {
		Strategy int       `yaml:"strategy" env:"STRATEGY"`
		Weight   []*Weight `env-required:"true" yaml:"weight" env:"WEIGHT"`
	}
	Weight struct {
		Chan  string `env-required:"true" yaml:"chan" env:"CHAN"`
		Value string `env-required:"true" yaml:"value" env:"VALUE"`
	}
)
