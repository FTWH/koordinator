//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Koordinator Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta2

import (
	unsafe "unsafe"

	config "github.com/koordinator-sh/koordinator/apis/scheduling/config"
	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apisconfig "k8s.io/kubernetes/pkg/scheduler/apis/config"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*LoadAwareSchedulingArgs)(nil), (*config.LoadAwareSchedulingArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_LoadAwareSchedulingArgs_To_config_LoadAwareSchedulingArgs(a.(*LoadAwareSchedulingArgs), b.(*config.LoadAwareSchedulingArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.LoadAwareSchedulingArgs)(nil), (*LoadAwareSchedulingArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_LoadAwareSchedulingArgs_To_v1beta2_LoadAwareSchedulingArgs(a.(*config.LoadAwareSchedulingArgs), b.(*LoadAwareSchedulingArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NodeNUMAResourceArgs)(nil), (*config.NodeNUMAResourceArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_NodeNUMAResourceArgs_To_config_NodeNUMAResourceArgs(a.(*NodeNUMAResourceArgs), b.(*config.NodeNUMAResourceArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.NodeNUMAResourceArgs)(nil), (*NodeNUMAResourceArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_NodeNUMAResourceArgs_To_v1beta2_NodeNUMAResourceArgs(a.(*config.NodeNUMAResourceArgs), b.(*NodeNUMAResourceArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ScoringStrategy)(nil), (*config.ScoringStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_ScoringStrategy_To_config_ScoringStrategy(a.(*ScoringStrategy), b.(*config.ScoringStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ScoringStrategy)(nil), (*ScoringStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ScoringStrategy_To_v1beta2_ScoringStrategy(a.(*config.ScoringStrategy), b.(*ScoringStrategy), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta2_LoadAwareSchedulingArgs_To_config_LoadAwareSchedulingArgs(in *LoadAwareSchedulingArgs, out *config.LoadAwareSchedulingArgs, s conversion.Scope) error {
	out.FilterExpiredNodeMetrics = (*bool)(unsafe.Pointer(in.FilterExpiredNodeMetrics))
	out.NodeMetricExpirationSeconds = (*int64)(unsafe.Pointer(in.NodeMetricExpirationSeconds))
	out.ResourceWeights = *(*map[v1.ResourceName]int64)(unsafe.Pointer(&in.ResourceWeights))
	out.UsageThresholds = *(*map[v1.ResourceName]int64)(unsafe.Pointer(&in.UsageThresholds))
	out.EstimatedScalingFactors = *(*map[v1.ResourceName]int64)(unsafe.Pointer(&in.EstimatedScalingFactors))
	return nil
}

// Convert_v1beta2_LoadAwareSchedulingArgs_To_config_LoadAwareSchedulingArgs is an autogenerated conversion function.
func Convert_v1beta2_LoadAwareSchedulingArgs_To_config_LoadAwareSchedulingArgs(in *LoadAwareSchedulingArgs, out *config.LoadAwareSchedulingArgs, s conversion.Scope) error {
	return autoConvert_v1beta2_LoadAwareSchedulingArgs_To_config_LoadAwareSchedulingArgs(in, out, s)
}

func autoConvert_config_LoadAwareSchedulingArgs_To_v1beta2_LoadAwareSchedulingArgs(in *config.LoadAwareSchedulingArgs, out *LoadAwareSchedulingArgs, s conversion.Scope) error {
	out.FilterExpiredNodeMetrics = (*bool)(unsafe.Pointer(in.FilterExpiredNodeMetrics))
	out.NodeMetricExpirationSeconds = (*int64)(unsafe.Pointer(in.NodeMetricExpirationSeconds))
	out.ResourceWeights = *(*map[v1.ResourceName]int64)(unsafe.Pointer(&in.ResourceWeights))
	out.UsageThresholds = *(*map[v1.ResourceName]int64)(unsafe.Pointer(&in.UsageThresholds))
	out.EstimatedScalingFactors = *(*map[v1.ResourceName]int64)(unsafe.Pointer(&in.EstimatedScalingFactors))
	return nil
}

// Convert_config_LoadAwareSchedulingArgs_To_v1beta2_LoadAwareSchedulingArgs is an autogenerated conversion function.
func Convert_config_LoadAwareSchedulingArgs_To_v1beta2_LoadAwareSchedulingArgs(in *config.LoadAwareSchedulingArgs, out *LoadAwareSchedulingArgs, s conversion.Scope) error {
	return autoConvert_config_LoadAwareSchedulingArgs_To_v1beta2_LoadAwareSchedulingArgs(in, out, s)
}

func autoConvert_v1beta2_NodeNUMAResourceArgs_To_config_NodeNUMAResourceArgs(in *NodeNUMAResourceArgs, out *config.NodeNUMAResourceArgs, s conversion.Scope) error {
	out.DefaultCPUBindPolicy = config.CPUBindPolicy(in.DefaultCPUBindPolicy)
	out.NUMAAllocateStrategy = config.NUMAAllocateStrategy(in.NUMAAllocateStrategy)
	out.ScoringStrategy = (*config.ScoringStrategy)(unsafe.Pointer(in.ScoringStrategy))
	return nil
}

// Convert_v1beta2_NodeNUMAResourceArgs_To_config_NodeNUMAResourceArgs is an autogenerated conversion function.
func Convert_v1beta2_NodeNUMAResourceArgs_To_config_NodeNUMAResourceArgs(in *NodeNUMAResourceArgs, out *config.NodeNUMAResourceArgs, s conversion.Scope) error {
	return autoConvert_v1beta2_NodeNUMAResourceArgs_To_config_NodeNUMAResourceArgs(in, out, s)
}

func autoConvert_config_NodeNUMAResourceArgs_To_v1beta2_NodeNUMAResourceArgs(in *config.NodeNUMAResourceArgs, out *NodeNUMAResourceArgs, s conversion.Scope) error {
	out.DefaultCPUBindPolicy = CPUBindPolicy(in.DefaultCPUBindPolicy)
	out.NUMAAllocateStrategy = NUMAAllocateStrategy(in.NUMAAllocateStrategy)
	out.ScoringStrategy = (*ScoringStrategy)(unsafe.Pointer(in.ScoringStrategy))
	return nil
}

// Convert_config_NodeNUMAResourceArgs_To_v1beta2_NodeNUMAResourceArgs is an autogenerated conversion function.
func Convert_config_NodeNUMAResourceArgs_To_v1beta2_NodeNUMAResourceArgs(in *config.NodeNUMAResourceArgs, out *NodeNUMAResourceArgs, s conversion.Scope) error {
	return autoConvert_config_NodeNUMAResourceArgs_To_v1beta2_NodeNUMAResourceArgs(in, out, s)
}

func autoConvert_v1beta2_ScoringStrategy_To_config_ScoringStrategy(in *ScoringStrategy, out *config.ScoringStrategy, s conversion.Scope) error {
	out.Type = config.ScoringStrategyType(in.Type)
	out.Resources = *(*[]apisconfig.ResourceSpec)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_v1beta2_ScoringStrategy_To_config_ScoringStrategy is an autogenerated conversion function.
func Convert_v1beta2_ScoringStrategy_To_config_ScoringStrategy(in *ScoringStrategy, out *config.ScoringStrategy, s conversion.Scope) error {
	return autoConvert_v1beta2_ScoringStrategy_To_config_ScoringStrategy(in, out, s)
}

func autoConvert_config_ScoringStrategy_To_v1beta2_ScoringStrategy(in *config.ScoringStrategy, out *ScoringStrategy, s conversion.Scope) error {
	out.Type = ScoringStrategyType(in.Type)
	out.Resources = *(*[]apisconfig.ResourceSpec)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_config_ScoringStrategy_To_v1beta2_ScoringStrategy is an autogenerated conversion function.
func Convert_config_ScoringStrategy_To_v1beta2_ScoringStrategy(in *config.ScoringStrategy, out *ScoringStrategy, s conversion.Scope) error {
	return autoConvert_config_ScoringStrategy_To_v1beta2_ScoringStrategy(in, out, s)
}
