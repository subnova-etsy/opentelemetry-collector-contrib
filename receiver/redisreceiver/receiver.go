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

package redisreceiver

import (
	"context"

	"github.com/go-redis/redis/v7"
	"github.com/open-telemetry/opentelemetry-collector/component"
	"github.com/open-telemetry/opentelemetry-collector/consumer"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver/interval"
)

type redisReceiver struct {
	logger         *zap.Logger
	config         *config
	consumer       consumer.MetricsConsumerOld
	intervalRunner *interval.Runner
}

func newRedisReceiver(
	logger *zap.Logger,
	config *config,
	consumer consumer.MetricsConsumerOld,
) *redisReceiver {
	return &redisReceiver{
		logger:   logger,
		config:   config,
		consumer: consumer,
	}
}

// Set up and kick off the interval runner.
func (r redisReceiver) Start(ctx context.Context, host component.Host) error {
	client := newRedisClient(&redis.Options{
		Addr:     r.config.Endpoint,
		Password: r.config.Password,
	})
	redisRunnable := newRedisRunnable(ctx, client, r.consumer, r.logger)
	r.intervalRunner = interval.NewRunner(r.config.CollectionInterval, redisRunnable)

	go func() {
		if err := r.intervalRunner.Start(); err != nil {
			host.ReportFatalError(err)
		}
	}()

	return nil
}

func (r redisReceiver) Shutdown(ctx context.Context) error {
	r.intervalRunner.Stop()
	return nil
}
