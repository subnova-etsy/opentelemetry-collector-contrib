// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package receivercreator

import (
	"reflect"

	otelconfig "github.com/open-telemetry/opentelemetry-collector/config"
	"github.com/open-telemetry/opentelemetry-collector/config/configmodels"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

const (
	// receiversConfigKey is the config key name used to specify the subreceivers.
	receiversConfigKey = "receivers"
	// endpointConfigKey is the key name mapping to ReceiverSettings.Endpoint.
	endpointConfigKey = "endpoint"
	// configKey is the key name in a subreceiver.
	configKey = "config"
)

// receiverConfig describes a receiver instance with a default config.
type receiverConfig struct {
	// fullName is the full subreceiver name (ie <receiver type>/<id>).
	fullName string
	// typeStr is set based on the configured receiver name.
	typeStr configmodels.Type
	// config is the map configured by the user in the config file. It is the contents of the map from
	// the "config" section. The keys and values are arbitrarily configured by the user.
	config userConfigMap
}

// userConfigMap is an arbitrary map of string keys to arbitrary values as specified by the user
type userConfigMap map[string]interface{}

// receiverTemplate is the configuration of a single subreceiver.
type receiverTemplate struct {
	receiverConfig

	// Rule is the discovery rule that when matched will create a receiver instance
	// based on receiverTemplate.
	Rule string `mapstructure:"rule"`
}

// newReceiverTemplate creates a receiverTemplate instance from the full name of a subreceiver
// and its arbitrary config map values.
func newReceiverTemplate(name string, config userConfigMap) (receiverTemplate, error) {
	typeStr, fullName, err := otelconfig.DecodeTypeAndName(name)
	if err != nil {
		return receiverTemplate{}, err
	}

	return receiverTemplate{
		receiverConfig: receiverConfig{
			typeStr:  configmodels.Type(typeStr),
			fullName: fullName,
			config:   config,
		},
	}, nil
}

// Config defines configuration for receiver_creator.
type Config struct {
	configmodels.ReceiverSettings `mapstructure:",squash"`
	receiverTemplates             map[string]receiverTemplate
	// WatchObservers are the extensions to listen to endpoints from.
	WatchObservers []configmodels.Type `mapstructure:"watch_observers"`
}

// Copied from the Viper but changed to use the same delimiter.
// See https://github.com/spf13/viper/issues/871
func viperSub(v *viper.Viper, key string) *viper.Viper {
	subv := otelconfig.NewViper()
	data := v.Get(key)
	if data == nil {
		return subv
	}

	if reflect.TypeOf(data).Kind() == reflect.Map {
		subv.MergeConfigMap(cast.ToStringMap(data))
		return subv
	}
	return subv
}
