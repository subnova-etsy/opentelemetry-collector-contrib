// Copyright 2019 OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package honeycombexporter

import (
	"github.com/open-telemetry/opentelemetry-collector/component"
	"github.com/open-telemetry/opentelemetry-collector/config/configerror"
	"github.com/open-telemetry/opentelemetry-collector/config/configmodels"
	"go.uber.org/zap"
)

const (
	// The value of "type" key in configuration.
	typeStr = "honeycomb"
)

// Factory is the factory for the Honeycomb exporter.
type Factory struct{}

// Type gets the type of the Exporter config created by this factory.
func (f *Factory) Type() configmodels.Type {
	return configmodels.Type(typeStr)
}

// CreateDefaultConfig creates the default configuration for the exporter.
func (f *Factory) CreateDefaultConfig() configmodels.Exporter {
	return &Config{
		ExporterSettings: configmodels.ExporterSettings{
			TypeVal: configmodels.Type(typeStr),
			NameVal: typeStr,
		},
		APIKey:  "",
		Dataset: "",
		APIURL:  "https://api.honeycomb.io",
		Debug:   false,
	}
}

// CreateTraceExporter creates a trace exporter based on this config.
func (f *Factory) CreateTraceExporter(logger *zap.Logger, cfg configmodels.Exporter) (component.TraceExporterOld, error) {
	eCfg := cfg.(*Config)
	return newHoneycombTraceExporter(eCfg)
}

// CreateMetricsExporter always returns nil.
func (f *Factory) CreateMetricsExporter(logger *zap.Logger,
	cfg configmodels.Exporter) (component.MetricsExporterOld, error) {
	return nil, configerror.ErrDataTypeIsNotSupported
}
