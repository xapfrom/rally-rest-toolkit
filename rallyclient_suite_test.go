/**
* Copyright 2014 Comcast Cable Communications Management, LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */
 
package rallyresttoolkit_test

import (
	"io"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRallyConnector(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RallyConnector Suite")
}

type mockDoer struct {
	req *http.Request
	res *http.Response
	err error
}

func (s *mockDoer) Do(req *http.Request) (*http.Response, error) {
	return s.res, s.err
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }
