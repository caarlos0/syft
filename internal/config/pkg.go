package config

import (
	"github.com/anchore/syft/syft/pkg/cataloger/packages"
	"github.com/spf13/viper"
)

type pkg struct {
	Cataloger               catalogerOptions `yaml:"cataloger" json:"cataloger" mapstructure:"cataloger"`
	SearchUnindexedArchives bool             `yaml:"search-unindexed-archives" json:"search-unindexed-archives" mapstructure:"search-unindexed-archives"`
	SearchIndexedArchives   bool             `yaml:"search-indexed-archives" json:"search-indexed-archives" mapstructure:"search-indexed-archives"`
}

func (cfg pkg) loadDefaultValues(v *viper.Viper) {
	cfg.Cataloger.loadDefaultValues(v)
	c := packages.DefaultSearchConfig()
	v.SetDefault("package.search-unindexed-archives", c.IncludeUnindexedArchives)
	v.SetDefault("package.search-indexed-archives", c.IncludeIndexedArchives)
}

func (cfg *pkg) parseConfigValues() error {
	return cfg.Cataloger.parseConfigValues()
}

func (cfg pkg) ToConfig() packages.SearchConfig {
	return packages.SearchConfig{
		IncludeIndexedArchives:   cfg.SearchIndexedArchives,
		IncludeUnindexedArchives: cfg.SearchUnindexedArchives,
	}
}
