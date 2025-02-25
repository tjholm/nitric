// Copyright 2021 Nitric Pty Ltd.
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

package secret_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nitrictech/nitric/pkg/plugins/secret"
)

var _ = Describe("Unimplemented Secret Plugin Tests", func() {
	uisp := &secret.UnimplementedSecretPlugin{}

	Context("Put", func() {
		When("Calling Put on UnimplementedSecretPlugin", func() {
			_, err := uisp.Put(nil, nil)

			It("should return an unimplemented error", func() {
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("UNIMPLEMENTED"))
			})
		})
	})

	Context("Access", func() {
		When("Calling Access on UnimplementedSecretPlugin", func() {
			_, err := uisp.Access(nil)

			It("should return an unimplemented error", func() {
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("UNIMPLEMENTED"))
			})
		})
	})
})
