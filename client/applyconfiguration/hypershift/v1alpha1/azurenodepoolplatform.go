/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// AzureNodePoolPlatformApplyConfiguration represents an declarative configuration of the AzureNodePoolPlatform type for use
// with apply.
type AzureNodePoolPlatformApplyConfiguration struct {
	VMSize                 *string                        `json:"vmsize,omitempty"`
	ImageID                *string                        `json:"imageID,omitempty"`
	DiskSizeGB             *int32                         `json:"diskSizeGB,omitempty"`
	DiskStorageAccountType *string                        `json:"diskStorageAccountType,omitempty"`
	AvailabilityZone       *string                        `json:"availabilityZone,omitempty"`
	DiskEncryptionSetID    *string                        `json:"diskEncryptionSetID,omitempty"`
	EnableEphemeralOSDisk  *bool                          `json:"enableEphemeralOSDisk,omitempty"`
	SubnetID               *string                        `json:"subnetID,omitempty"`
	Diagnostics            *DiagnosticsApplyConfiguration `json:"diagnostics,omitempty"`
	MachineIdentityID      *string                        `json:"machineIdentityID,omitempty"`
}

// AzureNodePoolPlatformApplyConfiguration constructs an declarative configuration of the AzureNodePoolPlatform type for use with
// apply.
func AzureNodePoolPlatform() *AzureNodePoolPlatformApplyConfiguration {
	return &AzureNodePoolPlatformApplyConfiguration{}
}

// WithVMSize sets the VMSize field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VMSize field is set to the value of the last call.
func (b *AzureNodePoolPlatformApplyConfiguration) WithVMSize(value string) *AzureNodePoolPlatformApplyConfiguration {
	b.VMSize = &value
	return b
}

// WithImageID sets the ImageID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImageID field is set to the value of the last call.
func (b *AzureNodePoolPlatformApplyConfiguration) WithImageID(value string) *AzureNodePoolPlatformApplyConfiguration {
	b.ImageID = &value
	return b
}

// WithDiskSizeGB sets the DiskSizeGB field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DiskSizeGB field is set to the value of the last call.
func (b *AzureNodePoolPlatformApplyConfiguration) WithDiskSizeGB(value int32) *AzureNodePoolPlatformApplyConfiguration {
	b.DiskSizeGB = &value
	return b
}

// WithDiskStorageAccountType sets the DiskStorageAccountType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DiskStorageAccountType field is set to the value of the last call.
func (b *AzureNodePoolPlatformApplyConfiguration) WithDiskStorageAccountType(value string) *AzureNodePoolPlatformApplyConfiguration {
	b.DiskStorageAccountType = &value
	return b
}

// WithAvailabilityZone sets the AvailabilityZone field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AvailabilityZone field is set to the value of the last call.
func (b *AzureNodePoolPlatformApplyConfiguration) WithAvailabilityZone(value string) *AzureNodePoolPlatformApplyConfiguration {
	b.AvailabilityZone = &value
	return b
}

// WithDiskEncryptionSetID sets the DiskEncryptionSetID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DiskEncryptionSetID field is set to the value of the last call.
func (b *AzureNodePoolPlatformApplyConfiguration) WithDiskEncryptionSetID(value string) *AzureNodePoolPlatformApplyConfiguration {
	b.DiskEncryptionSetID = &value
	return b
}

// WithEnableEphemeralOSDisk sets the EnableEphemeralOSDisk field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableEphemeralOSDisk field is set to the value of the last call.
func (b *AzureNodePoolPlatformApplyConfiguration) WithEnableEphemeralOSDisk(value bool) *AzureNodePoolPlatformApplyConfiguration {
	b.EnableEphemeralOSDisk = &value
	return b
}

// WithSubnetID sets the SubnetID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SubnetID field is set to the value of the last call.
func (b *AzureNodePoolPlatformApplyConfiguration) WithSubnetID(value string) *AzureNodePoolPlatformApplyConfiguration {
	b.SubnetID = &value
	return b
}

// WithDiagnostics sets the Diagnostics field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Diagnostics field is set to the value of the last call.
func (b *AzureNodePoolPlatformApplyConfiguration) WithDiagnostics(value *DiagnosticsApplyConfiguration) *AzureNodePoolPlatformApplyConfiguration {
	b.Diagnostics = value
	return b
}

// WithMachineIdentityID sets the MachineIdentityID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MachineIdentityID field is set to the value of the last call.
func (b *AzureNodePoolPlatformApplyConfiguration) WithMachineIdentityID(value string) *AzureNodePoolPlatformApplyConfiguration {
	b.MachineIdentityID = &value
	return b
}
