// Copyright 2021 Nitric Technologies Pty Ltd.
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

//go:build no_runtime_type_checking

package topic

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Topic) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_Topic) validateAddProviderParameters(provider interface{}) error {
	return nil
}

func (t *jsiiProxy_Topic) validateGetStringParameters(output *string) error {
	return nil
}

func (t *jsiiProxy_Topic) validateInterpolationForOutputParameters(moduleOutput *string) error {
	return nil
}

func (t *jsiiProxy_Topic) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateTopic_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTopic_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Topic) validateSetLocationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Topic) validateSetResourceGroupNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Topic) validateSetStackIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Topic) validateSetSubscribersParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Topic) validateSetTopicNameParameters(val *string) error {
	return nil
}

func validateNewTopicParameters(scope constructs.Construct, id *string, config *TopicConfig) error {
	return nil
}

