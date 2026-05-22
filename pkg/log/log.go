// Package logger provides a simple wrapper around zerolog.

// SPDX-License-Identifier: Apache-2.0

// Copyright 2026 Cisco Systems, Inc. and their affiliates

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package log

import (
	"io"
	"os"
	"runtime"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

const logFile = "aci-vetr.log"

// Logger aliases the zerolog.Logger
type Logger = zerolog.Logger

var (
	logger = New()

	// Convenience shortcuts for logging levels
	Debug = logger.Debug
	Info  = logger.Info
	Warn  = logger.Warn
	Error = logger.Error
	Fatal = logger.Fatal
	Panic = logger.Panic
	With  = logger.With
)

// New creates a new multi-level logger
func New() Logger {
	if testing.Testing() {
		return zerolog.Nop()
	}
	// If filename is specified, open file and assume file logging
	file, err := os.Create(logFile)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create log file")
	}

	// Levels default to zero, i.e. debug
	return zerolog.New(

		zerolog.ConsoleWriter{
			Out:     io.MultiWriter(os.Stderr, file),
			NoColor: runtime.GOOS == "windows",
		}).With().Timestamp().Logger()
}

func init() {
	// defaults to Info level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.DurationFieldInteger = true
}

// SetLevel sets the global log level.
func SetLevel(level zerolog.Level) {
	zerolog.SetGlobalLevel(level)
}

// WithFabric returns a logger with fabric context.
func WithFabric(fabricName string) Logger {
	return logger.With().Str("fabric", fabricName).Logger()
}
