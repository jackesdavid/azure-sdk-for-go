//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpolicyinsights

const (
	moduleName    = "armpolicyinsights"
	moduleVersion = "v0.1.0"
)

// ComplianceState - The compliance state that should be set on the resource.
type ComplianceState string

const (
	// ComplianceStateCompliant - The resource is in compliance with the policy.
	ComplianceStateCompliant ComplianceState = "Compliant"
	// ComplianceStateNonCompliant - The resource is not in compliance with the policy.
	ComplianceStateNonCompliant ComplianceState = "NonCompliant"
	// ComplianceStateUnknown - The compliance state of the resource is not known.
	ComplianceStateUnknown ComplianceState = "Unknown"
)

// PossibleComplianceStateValues returns the possible values for the ComplianceState const type.
func PossibleComplianceStateValues() []ComplianceState {
	return []ComplianceState{
		ComplianceStateCompliant,
		ComplianceStateNonCompliant,
		ComplianceStateUnknown,
	}
}

// ToPtr returns a *ComplianceState pointing to the current value.
func (c ComplianceState) ToPtr() *ComplianceState {
	return &c
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

type Enum0 string

const (
	Enum0MicrosoftManagement Enum0 = "Microsoft.Management"
)

// PossibleEnum0Values returns the possible values for the Enum0 const type.
func PossibleEnum0Values() []Enum0 {
	return []Enum0{
		Enum0MicrosoftManagement,
	}
}

// ToPtr returns a *Enum0 pointing to the current value.
func (c Enum0) ToPtr() *Enum0 {
	return &c
}

type Enum1 string

const (
	Enum1Default Enum1 = "default"
)

// PossibleEnum1Values returns the possible values for the Enum1 const type.
func PossibleEnum1Values() []Enum1 {
	return []Enum1{
		Enum1Default,
	}
}

// ToPtr returns a *Enum1 pointing to the current value.
func (c Enum1) ToPtr() *Enum1 {
	return &c
}

type Enum4 string

const (
	Enum4MicrosoftAuthorization Enum4 = "Microsoft.Authorization"
)

// PossibleEnum4Values returns the possible values for the Enum4 const type.
func PossibleEnum4Values() []Enum4 {
	return []Enum4{
		Enum4MicrosoftAuthorization,
	}
}

// ToPtr returns a *Enum4 pointing to the current value.
func (c Enum4) ToPtr() *Enum4 {
	return &c
}

type Enum6 string

const (
	Enum6Latest Enum6 = "latest"
)

// PossibleEnum6Values returns the possible values for the Enum6 const type.
func PossibleEnum6Values() []Enum6 {
	return []Enum6{
		Enum6Latest,
	}
}

// ToPtr returns a *Enum6 pointing to the current value.
func (c Enum6) ToPtr() *Enum6 {
	return &c
}

// FieldRestrictionResult - The type of restriction that is imposed on the field.
type FieldRestrictionResult string

const (
	// FieldRestrictionResultDeny - The field and/or values will be denied by policy.
	FieldRestrictionResultDeny FieldRestrictionResult = "Deny"
	// FieldRestrictionResultRemoved - The field will be removed by policy.
	FieldRestrictionResultRemoved FieldRestrictionResult = "Removed"
	// FieldRestrictionResultRequired - The field and/or values are required by policy.
	FieldRestrictionResultRequired FieldRestrictionResult = "Required"
)

// PossibleFieldRestrictionResultValues returns the possible values for the FieldRestrictionResult const type.
func PossibleFieldRestrictionResultValues() []FieldRestrictionResult {
	return []FieldRestrictionResult{
		FieldRestrictionResultDeny,
		FieldRestrictionResultRemoved,
		FieldRestrictionResultRequired,
	}
}

// ToPtr returns a *FieldRestrictionResult pointing to the current value.
func (c FieldRestrictionResult) ToPtr() *FieldRestrictionResult {
	return &c
}

type PolicyStatesResource string

const (
	PolicyStatesResourceDefault PolicyStatesResource = "default"
	PolicyStatesResourceLatest  PolicyStatesResource = "latest"
)

// PossiblePolicyStatesResourceValues returns the possible values for the PolicyStatesResource const type.
func PossiblePolicyStatesResourceValues() []PolicyStatesResource {
	return []PolicyStatesResource{
		PolicyStatesResourceDefault,
		PolicyStatesResourceLatest,
	}
}

// ToPtr returns a *PolicyStatesResource pointing to the current value.
func (c PolicyStatesResource) ToPtr() *PolicyStatesResource {
	return &c
}

// ResourceDiscoveryMode - The way resources to remediate are discovered. Defaults to ExistingNonCompliant if not specified.
type ResourceDiscoveryMode string

const (
	// ResourceDiscoveryModeExistingNonCompliant - Remediate resources that are already known to be non-compliant.
	ResourceDiscoveryModeExistingNonCompliant ResourceDiscoveryMode = "ExistingNonCompliant"
	// ResourceDiscoveryModeReEvaluateCompliance - Re-evaluate the compliance state of resources and then remediate the resources
	// found to be non-compliant.
	ResourceDiscoveryModeReEvaluateCompliance ResourceDiscoveryMode = "ReEvaluateCompliance"
)

// PossibleResourceDiscoveryModeValues returns the possible values for the ResourceDiscoveryMode const type.
func PossibleResourceDiscoveryModeValues() []ResourceDiscoveryMode {
	return []ResourceDiscoveryMode{
		ResourceDiscoveryModeExistingNonCompliant,
		ResourceDiscoveryModeReEvaluateCompliance,
	}
}

// ToPtr returns a *ResourceDiscoveryMode pointing to the current value.
func (c ResourceDiscoveryMode) ToPtr() *ResourceDiscoveryMode {
	return &c
}
