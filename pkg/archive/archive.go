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

package archive

import (
	"archive/zip"
	"os"
	"sync"
)

var zipMux sync.Mutex

// Writer is an archive writer interface
type Writer interface {
	Add(string, []byte) error
	Close() error
}

// FileWriter is a file-based implementation of archiveWriter
type FileWriter struct {
	file *os.File
	zw   *zip.Writer
}

// NewWriter creates a new file-based archive writer
func NewWriter(name string) (Writer, error) {
	f, err := os.Create(name)
	if err != nil {
		return FileWriter{}, err
	}
	zw := zip.NewWriter(f)
	return FileWriter{
		file: f,
		zw:   zw,
	}, nil
}

// Close closes the zip writer and file
func (a FileWriter) Close() error {
	err := a.zw.Close()
	if err != nil {
		return err
	}
	return a.file.Close()
}

// Add adds a file and content to the zip archive
func (a FileWriter) Add(name string, content []byte) error {
	zipMux.Lock()
	defer zipMux.Unlock()
	f, err := a.zw.Create(name)
	if err != nil {
		return nil
	}
	_, err = f.Write(content)
	return err
}
