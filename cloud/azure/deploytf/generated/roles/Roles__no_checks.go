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

package roles

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Roles) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (r *jsiiProxy_Roles) validateAddProviderParameters(provider interface{}) error {
	return nil
}

func (r *jsiiProxy_Roles) validateGetStringParameters(output *string) error {
	return nil
}

func (r *jsiiProxy_Roles) validateInterpolationForOutputParameters(moduleOutput *string) error {
	return nil
}

func (r *jsiiProxy_Roles) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateRoles_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRoles_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Roles) validateSetResourceGroupNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Roles) validateSetStackIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Roles) validateSetSubscriptionIdParameters(val *string) error {
	return nil
}

func validateNewRolesParameters(scope constructs.Construct, id *string, config *RolesConfig) error {
	return nil
}

