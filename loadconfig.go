// Copyright 2018 Saferwall. All rights reserved.
// Use of this source code is governed by Apache v2 license
// license that can be found in the LICENSE file.

package pe

import (
	"encoding/binary"
	"fmt"
	"reflect"
)

const (

	// The GuardFlags field contains a combination of one or more of the
	// following flags and subfields:

	// ImageGuardCfInstrumented indicates that the module performs control flow
	// integrity checks using system-supplied support.
	ImageGuardCfInstrumented = 0x00000100

	// ImageGuardCfWInstrumented indicates that the module performs control
	// flow and write integrity checks.
	ImageGuardCfWInstrumented = 0x00000200

	// ImageGuardCfFunctionTablePresent indicates that the module contains
	// valid control flow target metadata.
	ImageGuardCfFunctionTablePresent = 0x00000400

	// ImageGuardSecurityCookieUnused indicates that the module does not make
	// use of the /GS security cookie.
	ImageGuardSecurityCookieUnused = 0x00000800

	// ImageGuardProtectDelayloadIAT indicates that the module supports read
	// only delay load IAT.
	ImageGuardProtectDelayloadIAT = 0x00001000

	// ImageGuardDelayloadIATInItsOwnSection indicates that the Delayload
	// import table in its own .didat section (with nothing else in it) that
	// can be freely reprotected.
	ImageGuardDelayloadIATInItsOwnSection = 0x00002000

	// ImageGuardCfExportSuppressionInfoPresent indicates that the module
	// contains suppressed export information. This also infers that the
	// address taken IAT table is also present in the load config.
	ImageGuardCfExportSuppressionInfoPresent = 0x00004000

	// ImageGuardCfEnableExportSuppression indicates that the module enables
	// suppression of exports.
	ImageGuardCfEnableExportSuppression = 0x00008000

	// ImageGuardCfLongjumpTablePresent indicates that the module contains
	// longjmp target information.
	ImageGuardCfLongjumpTablePresent = 0x00010000

	// ImageGuardCfFunctionTableSizeMask indicates that the mask for the
	// subfield that contains the stride of Control Flow Guard function table
	// entries (that is, the additional count of bytes per table entry).
	ImageGuardCfFunctionTableSizeMask = 0xF0000000

	// ImageGuardCfFunctionTableSizeShift indicates the shift to right-justify
	// Guard CF function table stride.
	ImageGuardCfFunctionTableSizeShift = 28

	// ImageGuardFlagFIDSuppressed indicates that the call target is explicitly
	// suppressed (do not treat it as valid for purposes of CFG)
	ImageGuardFlagFIDSuppressed = 0x1

	// ImageGuardFlagExportSuppressed indicates that the call target is export
	// suppressed. See Export suppression for more details
	ImageGuardFlagExportSuppressed = 0x2

	ImageDynamicRelocationGuardRfPrologue = 0x00000001
	ImageDynamicRelocationGuardREpilogue  = 0x00000002
	ImageEnclaveLongIDLength              = 32
	ImageEnclaveShortIDLength             = 16

	// ImageEnclaveImportMatchNone indicates that none of the identifiers of the
	// image need to match the value in the import record.
	ImageEnclaveImportMatchNone = 0x00000000

	// ImageEnclaveImportMatchUniqueId indicates that the value of the enclave
	// unique identifier of the image must match the value in the import record.
	// Otherwise, loading of the image fails.
	ImageEnclaveImportMatchUniqueID = 0x00000001

	// ImageEnclaveImportMatchAuthorId indicates that the value of the enclave
	// author identifier of the image must match the value in the import record.
	// Otherwise, loading of the image fails. If this flag is set and the import
	// record indicates an author identifier of all zeros, the imported image
	// must be part of the Windows installation.
	ImageEnclaveImportMatchAuthorID = 0x00000002

	// ImageEnclaveImportMatchFamilyId indicates that the value of the enclave
	// family identifier of the image must match the value in the import record.
	// Otherwise, loading of the image fails.
	ImageEnclaveImportMatchFamilyID = 0x00000003

	// ImageEnclaveImportMatchImageId indicates that the value of the enclave
	// image identifier must match the value in the import record. Otherwise,
	// loading of the image fails.
	ImageEnclaveImportMatchImageID = 0x00000004
)

// https://www.virtualbox.org/svn/vbox/trunk/include/iprt/formats/pecoff.h

// ImageLoadConfigDirectory32v1 size is 0x40.
type ImageLoadConfigDirectory32v1 struct {
	Size                          uint32 `json:"size"`
	TimeDateStamp                 uint32 `json:"time_date_stamp"`
	MajorVersion                  uint16 `json:"major_version"`
	MinorVersion                  uint16 `json:"minor_version"`
	GlobalFlagsClear              uint32 `json:"global_flags_clear"`
	GlobalFlagsSet                uint32 `json:"global_flags_set"`
	CriticalSectionDefaultTimeout uint32 `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold    uint32 `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold    uint32 `json:"de_commit_total_free_threshold"`
	LockPrefixTable               uint32 `json:"lock_prefix_table"`
	MaximumAllocationSize         uint32 `json:"maximum_allocation_size"`
	VirtualMemoryThreshold        uint32 `json:"virtual_memory_threshold"`
	ProcessHeapFlags              uint32 `json:"process_heap_flags"`
	ProcessAffinityMask           uint32 `json:"process_affinity_mask"`
	CSDVersion                    uint16 `json:"csd_version"`
	DependentLoadFlags            uint16 `json:"dependent_load_flags"`
	EditList                      uint32 `json:"edit_list"`
	SecurityCookie                uint32 `json:"security_cookie"`
}

// ImageLoadConfigDirectory32v2 size is 0x48.
type ImageLoadConfigDirectory32v2 struct {
	Size                          uint32 `json:"size"`
	TimeDateStamp                 uint32 `json:"time_date_stamp"`
	MajorVersion                  uint16 `json:"major_version"`
	MinorVersion                  uint16 `json:"minor_version"`
	GlobalFlagsClear              uint32 `json:"global_flags_clear"`
	GlobalFlagsSet                uint32 `json:"global_flags_set"`
	CriticalSectionDefaultTimeout uint32 `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold    uint32 `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold    uint32 `json:"de_commit_total_free_threshold"`
	LockPrefixTable               uint32 `json:"lock_prefix_table"`
	MaximumAllocationSize         uint32 `json:"maximum_allocation_size"`
	VirtualMemoryThreshold        uint32 `json:"virtual_memory_threshold"`
	ProcessHeapFlags              uint32 `json:"process_heap_flags"`
	ProcessAffinityMask           uint32 `json:"process_affinity_mask"`
	CSDVersion                    uint16 `json:"csd_version"`
	DependentLoadFlags            uint16 `json:"dependent_load_flags"`
	EditList                      uint32 `json:"edit_list"`
	SecurityCookie                uint32 `json:"security_cookie"`
	SEHandlerTable                uint32 `json:"se_handler_table"`
	SEHandlerCount                uint32 `json:"se_handler_count"`
}

// ImageLoadConfigDirectory32v3 size is 0x5C.
type ImageLoadConfigDirectory32v3 struct {
	Size                           uint32 `json:"size"`
	TimeDateStamp                  uint32 `json:"time_date_stamp"`
	MajorVersion                   uint16 `json:"major_version"`
	MinorVersion                   uint16 `json:"minor_version"`
	GlobalFlagsClear               uint32 `json:"global_flags_clear"`
	GlobalFlagsSet                 uint32 `json:"global_flags_set"`
	CriticalSectionDefaultTimeout  uint32 `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold     uint32 `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold     uint32 `json:"de_commit_total_free_threshold"`
	LockPrefixTable                uint32 `json:"lock_prefix_table"`
	MaximumAllocationSize          uint32 `json:"maximum_allocation_size"`
	VirtualMemoryThreshold         uint32 `json:"virtual_memory_threshold"`
	ProcessHeapFlags               uint32 `json:"process_heap_flags"`
	ProcessAffinityMask            uint32 `json:"process_affinity_mask"`
	CSDVersion                     uint16 `json:"csd_version"`
	DependentLoadFlags             uint16 `json:"dependent_load_flags"`
	EditList                       uint32 `json:"edit_list"`
	SecurityCookie                 uint32 `json:"security_cookie"`
	SEHandlerTable                 uint32 `json:"se_handler_table"`
	SEHandlerCount                 uint32 `json:"se_handler_count"`
	GuardCFCheckFunctionPointer    uint32 `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer uint32 `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable           uint32 `json:"guard_cf_function_table"`
	GuardCFFunctionCount           uint32 `json:"guard_cf_function_count"`
	GuardFlags                     uint32 `json:"guard_flags"`
}

// ImageLoadConfigDirectory32v4 size is 0x6c, since Windows 10 (preview 9879)
type ImageLoadConfigDirectory32v4 struct {
	Size                           uint32                       `json:"size"`
	TimeDateStamp                  uint32                       `json:"time_date_stamp"`
	MajorVersion                   uint16                       `json:"major_version"`
	MinorVersion                   uint16                       `json:"minor_version"`
	GlobalFlagsClear               uint32                       `json:"global_flags_clear"`
	GlobalFlagsSet                 uint32                       `json:"global_flags_set"`
	CriticalSectionDefaultTimeout  uint32                       `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold     uint32                       `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold     uint32                       `json:"de_commit_total_free_threshold"`
	LockPrefixTable                uint32                       `json:"lock_prefix_table"`
	MaximumAllocationSize          uint32                       `json:"maximum_allocation_size"`
	VirtualMemoryThreshold         uint32                       `json:"virtual_memory_threshold"`
	ProcessHeapFlags               uint32                       `json:"process_heap_flags"`
	ProcessAffinityMask            uint32                       `json:"process_affinity_mask"`
	CSDVersion                     uint16                       `json:"csd_version"`
	DependentLoadFlags             uint16                       `json:"dependent_load_flags"`
	EditList                       uint32                       `json:"edit_list"`
	SecurityCookie                 uint32                       `json:"security_cookie"`
	SEHandlerTable                 uint32                       `json:"se_handler_table"`
	SEHandlerCount                 uint32                       `json:"se_handler_count"`
	GuardCFCheckFunctionPointer    uint32                       `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer uint32                       `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable           uint32                       `json:"guard_cf_function_table"`
	GuardCFFunctionCount           uint32                       `json:"guard_cf_function_count"`
	GuardFlags                     uint32                       `json:"guard_flags"`
	CodeIntegrity                  ImageLoadConfigCodeIntegrity `json:"code_integrity"`
}

// ImageLoadConfigDirectory32v5 size is 0x78, since Windows 10 build 14286 (or maybe earlier).
type ImageLoadConfigDirectory32v5 struct {
	Size                           uint32                       `json:"size"`
	TimeDateStamp                  uint32                       `json:"time_date_stamp"`
	MajorVersion                   uint16                       `json:"major_version"`
	MinorVersion                   uint16                       `json:"minor_version"`
	GlobalFlagsClear               uint32                       `json:"global_flags_clear"`
	GlobalFlagsSet                 uint32                       `json:"global_flags_set"`
	CriticalSectionDefaultTimeout  uint32                       `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold     uint32                       `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold     uint32                       `json:"de_commit_total_free_threshold"`
	LockPrefixTable                uint32                       `json:"lock_prefix_table"`
	MaximumAllocationSize          uint32                       `json:"maximum_allocation_size"`
	VirtualMemoryThreshold         uint32                       `json:"virtual_memory_threshold"`
	ProcessHeapFlags               uint32                       `json:"process_heap_flags"`
	ProcessAffinityMask            uint32                       `json:"process_affinity_mask"`
	CSDVersion                     uint16                       `json:"csd_version"`
	DependentLoadFlags             uint16                       `json:"dependent_load_flags"`
	EditList                       uint32                       `json:"edit_list"`
	SecurityCookie                 uint32                       `json:"security_cookie"`
	SEHandlerTable                 uint32                       `json:"se_handler_table"`
	SEHandlerCount                 uint32                       `json:"se_handler_count"`
	GuardCFCheckFunctionPointer    uint32                       `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer uint32                       `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable           uint32                       `json:"guard_cf_function_table"`
	GuardCFFunctionCount           uint32                       `json:"guard_cf_function_count"`
	GuardFlags                     uint32                       `json:"guard_flags"`
	CodeIntegrity                  ImageLoadConfigCodeIntegrity `json:"code_integrity"`
	GuardAddressTakenIatEntryTable uint32                       `json:"guard_address_taken_iat_entry_table"`
	GuardAddressTakenIatEntryCount uint32                       `json:"guard_address_taken_iat_entry_count"`
	GuardLongJumpTargetTable       uint32                       `json:"guard_long_jump_target_table"`
	GuardLongJumpTargetCount       uint32                       `json:"guard_long_jump_target_count"`
}

// ImageLoadConfigDirectory32v6 size is 0x80, since Windows 10 build 14383 (or maybe earlier).
type ImageLoadConfigDirectory32v6 struct {
	Size                           uint32                       `json:"size"`
	TimeDateStamp                  uint32                       `json:"time_date_stamp"`
	MajorVersion                   uint16                       `json:"major_version"`
	MinorVersion                   uint16                       `json:"minor_version"`
	GlobalFlagsClear               uint32                       `json:"global_flags_clear"`
	GlobalFlagsSet                 uint32                       `json:"global_flags_set"`
	CriticalSectionDefaultTimeout  uint32                       `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold     uint32                       `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold     uint32                       `json:"de_commit_total_free_threshold"`
	LockPrefixTable                uint32                       `json:"lock_prefix_table"`
	MaximumAllocationSize          uint32                       `json:"maximum_allocation_size"`
	VirtualMemoryThreshold         uint32                       `json:"virtual_memory_threshold"`
	ProcessHeapFlags               uint32                       `json:"process_heap_flags"`
	ProcessAffinityMask            uint32                       `json:"process_affinity_mask"`
	CSDVersion                     uint16                       `json:"csd_version"`
	DependentLoadFlags             uint16                       `json:"dependent_load_flags"`
	EditList                       uint32                       `json:"edit_list"`
	SecurityCookie                 uint32                       `json:"security_cookie"`
	SEHandlerTable                 uint32                       `json:"se_handler_table"`
	SEHandlerCount                 uint32                       `json:"se_handler_count"`
	GuardCFCheckFunctionPointer    uint32                       `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer uint32                       `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable           uint32                       `json:"guard_cf_function_table"`
	GuardCFFunctionCount           uint32                       `json:"guard_cf_function_count"`
	GuardFlags                     uint32                       `json:"guard_flags"`
	CodeIntegrity                  ImageLoadConfigCodeIntegrity `json:"code_integrity"`
	GuardAddressTakenIatEntryTable uint32                       `json:"guard_address_taken_iat_entry_table"`
	GuardAddressTakenIatEntryCount uint32                       `json:"guard_address_taken_iat_entry_count"`
	GuardLongJumpTargetTable       uint32                       `json:"guard_long_jump_target_table"`
	GuardLongJumpTargetCount       uint32                       `json:"guard_long_jump_target_count"`
	DynamicValueRelocTable         uint32                       `json:"dynamic_value_reloc_table"`
	HybridMetadataPointer          uint32                       `json:"hybrid_metadata_pointer"`
}

// ImageLoadConfigDirectory32v7 size is 0x90, since Windows 10 build 14901 (or maybe earlier)
type ImageLoadConfigDirectory32v7 struct {
	Size                                 uint32                       `json:"size"`
	TimeDateStamp                        uint32                       `json:"time_date_stamp"`
	MajorVersion                         uint16                       `json:"major_version"`
	MinorVersion                         uint16                       `json:"minor_version"`
	GlobalFlagsClear                     uint32                       `json:"global_flags_clear"`
	GlobalFlagsSet                       uint32                       `json:"global_flags_set"`
	CriticalSectionDefaultTimeout        uint32                       `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold           uint32                       `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold           uint32                       `json:"de_commit_total_free_threshold"`
	LockPrefixTable                      uint32                       `json:"lock_prefix_table"`
	MaximumAllocationSize                uint32                       `json:"maximum_allocation_size"`
	VirtualMemoryThreshold               uint32                       `json:"virtual_memory_threshold"`
	ProcessHeapFlags                     uint32                       `json:"process_heap_flags"`
	ProcessAffinityMask                  uint32                       `json:"process_affinity_mask"`
	CSDVersion                           uint16                       `json:"csd_version"`
	DependentLoadFlags                   uint16                       `json:"dependent_load_flags"`
	EditList                             uint32                       `json:"edit_list"`
	SecurityCookie                       uint32                       `json:"security_cookie"`
	SEHandlerTable                       uint32                       `json:"se_handler_table"`
	SEHandlerCount                       uint32                       `json:"se_handler_count"`
	GuardCFCheckFunctionPointer          uint32                       `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer       uint32                       `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable                 uint32                       `json:"guard_cf_function_table"`
	GuardCFFunctionCount                 uint32                       `json:"guard_cf_function_count"`
	GuardFlags                           uint32                       `json:"guard_flags"`
	CodeIntegrity                        ImageLoadConfigCodeIntegrity `json:"code_integrity"`
	GuardAddressTakenIatEntryTable       uint32                       `json:"guard_address_taken_iat_entry_table"`
	GuardAddressTakenIatEntryCount       uint32                       `json:"guard_address_taken_iat_entry_count"`
	GuardLongJumpTargetTable             uint32                       `json:"guard_long_jump_target_table"`
	GuardLongJumpTargetCount             uint32                       `json:"guard_long_jump_target_count"`
	DynamicValueRelocTable               uint32                       `json:"dynamic_value_reloc_table"`
	CHPEMetadataPointer                  uint32                       `json:"chpe_metadata_pointer"`
	GuardRFFailureRoutine                uint32                       `json:"guard_rf_failure_routine"`
	GuardRFFailureRoutineFunctionPointer uint32                       `json:"guard_rf_failure_routine_function_pointer"`
	DynamicValueRelocTableOffset         uint32                       `json:"dynamic_value_reloc_table_offset"`
	DynamicValueRelocTableSection        uint16                       `json:"dynamic_value_reloc_table_section"`
	Reserved2                            uint16                       `json:"reserved_2"`
}

// ImageLoadConfigDirectory32v8 size is 0x98, since Windows 10 build 15002 (or maybe earlier).
type ImageLoadConfigDirectory32v8 struct {
	Size                                     uint32                       `json:"size"`
	TimeDateStamp                            uint32                       `json:"time_date_stamp"`
	MajorVersion                             uint16                       `json:"major_version"`
	MinorVersion                             uint16                       `json:"minor_version"`
	GlobalFlagsClear                         uint32                       `json:"global_flags_clear"`
	GlobalFlagsSet                           uint32                       `json:"global_flags_set"`
	CriticalSectionDefaultTimeout            uint32                       `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold               uint32                       `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold               uint32                       `json:"de_commit_total_free_threshold"`
	LockPrefixTable                          uint32                       `json:"lock_prefix_table"`
	MaximumAllocationSize                    uint32                       `json:"maximum_allocation_size"`
	VirtualMemoryThreshold                   uint32                       `json:"virtual_memory_threshold"`
	ProcessHeapFlags                         uint32                       `json:"process_heap_flags"`
	ProcessAffinityMask                      uint32                       `json:"process_affinity_mask"`
	CSDVersion                               uint16                       `json:"csd_version"`
	DependentLoadFlags                       uint16                       `json:"dependent_load_flags"`
	EditList                                 uint32                       `json:"edit_list"`
	SecurityCookie                           uint32                       `json:"security_cookie"`
	SEHandlerTable                           uint32                       `json:"se_handler_table"`
	SEHandlerCount                           uint32                       `json:"se_handler_count"`
	GuardCFCheckFunctionPointer              uint32                       `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer           uint32                       `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable                     uint32                       `json:"guard_cf_function_table"`
	GuardCFFunctionCount                     uint32                       `json:"guard_cf_function_count"`
	GuardFlags                               uint32                       `json:"guard_flags"`
	CodeIntegrity                            ImageLoadConfigCodeIntegrity `json:"code_integrity"`
	GuardAddressTakenIatEntryTable           uint32                       `json:"guard_address_taken_iat_entry_table"`
	GuardAddressTakenIatEntryCount           uint32                       `json:"guard_address_taken_iat_entry_count"`
	GuardLongJumpTargetTable                 uint32                       `json:"guard_long_jump_target_table"`
	GuardLongJumpTargetCount                 uint32                       `json:"guard_long_jump_target_count"`
	DynamicValueRelocTable                   uint32                       `json:"dynamic_value_reloc_table"`
	CHPEMetadataPointer                      uint32                       `json:"chpe_metadata_pointer"`
	GuardRFFailureRoutine                    uint32                       `json:"guard_rf_failure_routine"`
	GuardRFFailureRoutineFunctionPointer     uint32                       `json:"guard_rf_failure_routine_function_pointer"`
	DynamicValueRelocTableOffset             uint32                       `json:"dynamic_value_reloc_table_offset"`
	DynamicValueRelocTableSection            uint16                       `json:"dynamic_value_reloc_table_section"`
	Reserved2                                uint16                       `json:"reserved_2"`
	GuardRFVerifyStackPointerFunctionPointer uint32                       `json:"guard_rf_verify_stack_pointer_function_pointer"`
	HotPatchTableOffset                      uint32                       `json:"hot_patch_table_offset"`
}

// ImageLoadConfigDirectory32v9 size is 0xA0, since Windows 10 build 16237 (or maybe earlier).
type ImageLoadConfigDirectory32v9 struct {
	Size                                     uint32                       `json:"size"`
	TimeDateStamp                            uint32                       `json:"time_date_stamp"`
	MajorVersion                             uint16                       `json:"major_version"`
	MinorVersion                             uint16                       `json:"minor_version"`
	GlobalFlagsClear                         uint32                       `json:"global_flags_clear"`
	GlobalFlagsSet                           uint32                       `json:"global_flags_set"`
	CriticalSectionDefaultTimeout            uint32                       `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold               uint32                       `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold               uint32                       `json:"de_commit_total_free_threshold"`
	LockPrefixTable                          uint32                       `json:"lock_prefix_table"`
	MaximumAllocationSize                    uint32                       `json:"maximum_allocation_size"`
	VirtualMemoryThreshold                   uint32                       `json:"virtual_memory_threshold"`
	ProcessHeapFlags                         uint32                       `json:"process_heap_flags"`
	ProcessAffinityMask                      uint32                       `json:"process_affinity_mask"`
	CSDVersion                               uint16                       `json:"csd_version"`
	DependentLoadFlags                       uint16                       `json:"dependent_load_flags"`
	EditList                                 uint32                       `json:"edit_list"`
	SecurityCookie                           uint32                       `json:"security_cookie"`
	SEHandlerTable                           uint32                       `json:"se_handler_table"`
	SEHandlerCount                           uint32                       `json:"se_handler_count"`
	GuardCFCheckFunctionPointer              uint32                       `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer           uint32                       `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable                     uint32                       `json:"guard_cf_function_table"`
	GuardCFFunctionCount                     uint32                       `json:"guard_cf_function_count"`
	GuardFlags                               uint32                       `json:"guard_flags"`
	CodeIntegrity                            ImageLoadConfigCodeIntegrity `json:"code_integrity"`
	GuardAddressTakenIatEntryTable           uint32                       `json:"guard_address_taken_iat_entry_table"`
	GuardAddressTakenIatEntryCount           uint32                       `json:"guard_address_taken_iat_entry_count"`
	GuardLongJumpTargetTable                 uint32                       `json:"guard_long_jump_target_table"`
	GuardLongJumpTargetCount                 uint32                       `json:"guard_long_jump_target_count"`
	DynamicValueRelocTable                   uint32                       `json:"dynamic_value_reloc_table"`
	CHPEMetadataPointer                      uint32                       `json:"chpe_metadata_pointer"`
	GuardRFFailureRoutine                    uint32                       `json:"guard_rf_failure_routine"`
	GuardRFFailureRoutineFunctionPointer     uint32                       `json:"guard_rf_failure_routine_function_pointer"`
	DynamicValueRelocTableOffset             uint32                       `json:"dynamic_value_reloc_table_offset"`
	DynamicValueRelocTableSection            uint16                       `json:"dynamic_value_reloc_table_section"`
	Reserved2                                uint16                       `json:"reserved_2"`
	GuardRFVerifyStackPointerFunctionPointer uint32                       `json:"guard_rf_verify_stack_pointer_function_pointer"`
	HotPatchTableOffset                      uint32                       `json:"hot_patch_table_offset"`
	Reserved3                                uint32                       `json:"reserved_3"`
	EnclaveConfigurationPointer              uint32                       `json:"enclave_configuration_pointer"`
}

// ImageLoadConfigDirectory32v10 size is 0xA4, since Windows 10 build 18362 (or maybe earlier).
type ImageLoadConfigDirectory32v10 struct {
	Size                                     uint32                       `json:"size"`
	TimeDateStamp                            uint32                       `json:"time_date_stamp"`
	MajorVersion                             uint16                       `json:"major_version"`
	MinorVersion                             uint16                       `json:"minor_version"`
	GlobalFlagsClear                         uint32                       `json:"global_flags_clear"`
	GlobalFlagsSet                           uint32                       `json:"global_flags_set"`
	CriticalSectionDefaultTimeout            uint32                       `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold               uint32                       `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold               uint32                       `json:"de_commit_total_free_threshold"`
	LockPrefixTable                          uint32                       `json:"lock_prefix_table"`
	MaximumAllocationSize                    uint32                       `json:"maximum_allocation_size"`
	VirtualMemoryThreshold                   uint32                       `json:"virtual_memory_threshold"`
	ProcessHeapFlags                         uint32                       `json:"process_heap_flags"`
	ProcessAffinityMask                      uint32                       `json:"process_affinity_mask"`
	CSDVersion                               uint16                       `json:"csd_version"`
	DependentLoadFlags                       uint16                       `json:"dependent_load_flags"`
	EditList                                 uint32                       `json:"edit_list"`
	SecurityCookie                           uint32                       `json:"security_cookie"`
	SEHandlerTable                           uint32                       `json:"se_handler_table"`
	SEHandlerCount                           uint32                       `json:"se_handler_count"`
	GuardCFCheckFunctionPointer              uint32                       `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer           uint32                       `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable                     uint32                       `json:"guard_cf_function_table"`
	GuardCFFunctionCount                     uint32                       `json:"guard_cf_function_count"`
	GuardFlags                               uint32                       `json:"guard_flags"`
	CodeIntegrity                            ImageLoadConfigCodeIntegrity `json:"code_integrity"`
	GuardAddressTakenIatEntryTable           uint32                       `json:"guard_address_taken_iat_entry_table"`
	GuardAddressTakenIatEntryCount           uint32                       `json:"guard_address_taken_iat_entry_count"`
	GuardLongJumpTargetTable                 uint32                       `json:"guard_long_jump_target_table"`
	GuardLongJumpTargetCount                 uint32                       `json:"guard_long_jump_target_count"`
	DynamicValueRelocTable                   uint32                       `json:"dynamic_value_reloc_table"`
	CHPEMetadataPointer                      uint32                       `json:"chpe_metadata_pointer"`
	GuardRFFailureRoutine                    uint32                       `json:"guard_rf_failure_routine"`
	GuardRFFailureRoutineFunctionPointer     uint32                       `json:"guard_rf_failure_routine_function_pointer"`
	DynamicValueRelocTableOffset             uint32                       `json:"dynamic_value_reloc_table_offset"`
	DynamicValueRelocTableSection            uint16                       `json:"dynamic_value_reloc_table_section"`
	Reserved2                                uint16                       `json:"reserved_2"`
	GuardRFVerifyStackPointerFunctionPointer uint32                       `json:"guard_rf_verify_stack_pointer_function_pointer"`
	HotPatchTableOffset                      uint32                       `json:"hot_patch_table_offset"`
	Reserved3                                uint32                       `json:"reserved_3"`
	EnclaveConfigurationPointer              uint32                       `json:"enclave_configuration_pointer"`
	VolatileMetadataPointer                  uint32                       `json:"volatile_metadata_pointer"`
}

// ImageLoadConfigDirectory32v11 size is 0xAC, since Windows 10 build 19564 (or maybe earlier).
type ImageLoadConfigDirectory32v11 struct {
	Size                                     uint32                       `json:"size"`
	TimeDateStamp                            uint32                       `json:"time_date_stamp"`
	MajorVersion                             uint16                       `json:"major_version"`
	MinorVersion                             uint16                       `json:"minor_version"`
	GlobalFlagsClear                         uint32                       `json:"global_flags_clear"`
	GlobalFlagsSet                           uint32                       `json:"global_flags_set"`
	CriticalSectionDefaultTimeout            uint32                       `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold               uint32                       `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold               uint32                       `json:"de_commit_total_free_threshold"`
	LockPrefixTable                          uint32                       `json:"lock_prefix_table"`
	MaximumAllocationSize                    uint32                       `json:"maximum_allocation_size"`
	VirtualMemoryThreshold                   uint32                       `json:"virtual_memory_threshold"`
	ProcessHeapFlags                         uint32                       `json:"process_heap_flags"`
	ProcessAffinityMask                      uint32                       `json:"process_affinity_mask"`
	CSDVersion                               uint16                       `json:"csd_version"`
	DependentLoadFlags                       uint16                       `json:"dependent_load_flags"`
	EditList                                 uint32                       `json:"edit_list"`
	SecurityCookie                           uint32                       `json:"security_cookie"`
	SEHandlerTable                           uint32                       `json:"se_handler_table"`
	SEHandlerCount                           uint32                       `json:"se_handler_count"`
	GuardCFCheckFunctionPointer              uint32                       `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer           uint32                       `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable                     uint32                       `json:"guard_cf_function_table"`
	GuardCFFunctionCount                     uint32                       `json:"guard_cf_function_count"`
	GuardFlags                               uint32                       `json:"guard_flags"`
	CodeIntegrity                            ImageLoadConfigCodeIntegrity `json:"code_integrity"`
	GuardAddressTakenIatEntryTable           uint32                       `json:"guard_address_taken_iat_entry_table"`
	GuardAddressTakenIatEntryCount           uint32                       `json:"guard_address_taken_iat_entry_count"`
	GuardLongJumpTargetTable                 uint32                       `json:"guard_long_jump_target_table"`
	GuardLongJumpTargetCount                 uint32                       `json:"guard_long_jump_target_count"`
	DynamicValueRelocTable                   uint32                       `json:"dynamic_value_reloc_table"`
	CHPEMetadataPointer                      uint32                       `json:"chpe_metadata_pointer"`
	GuardRFFailureRoutine                    uint32                       `json:"guard_rf_failure_routine"`
	GuardRFFailureRoutineFunctionPointer     uint32                       `json:"guard_rf_failure_routine_function_pointer"`
	DynamicValueRelocTableOffset             uint32                       `json:"dynamic_value_reloc_table_offset"`
	DynamicValueRelocTableSection            uint16                       `json:"dynamic_value_reloc_table_section"`
	Reserved2                                uint16                       `json:"reserved_2"`
	GuardRFVerifyStackPointerFunctionPointer uint32                       `json:"guard_rf_verify_stack_pointer_function_pointer"`
	HotPatchTableOffset                      uint32                       `json:"hot_patch_table_offset"`
	Reserved3                                uint32                       `json:"reserved_3"`
	EnclaveConfigurationPointer              uint32                       `json:"enclave_configuration_pointer"`
	VolatileMetadataPointer                  uint32                       `json:"volatile_metadata_pointer"`
	GuardEHContinuationTable                 uint32                       `json:"guard_eh_continuation_table"`
	GuardEHContinuationCount                 uint32                       `json:"guard_eh_continuation_count"`
}

// ImageLoadConfigDirectory32v12 size is 0xB8, since Visual C++ 2019 / RS5_IMAGE_LOAD_CONFIG_DIRECTORY32.
type ImageLoadConfigDirectory32v12 struct {
	Size                                     uint32                       `json:"size"`
	TimeDateStamp                            uint32                       `json:"time_date_stamp"`
	MajorVersion                             uint16                       `json:"major_version"`
	MinorVersion                             uint16                       `json:"minor_version"`
	GlobalFlagsClear                         uint32                       `json:"global_flags_clear"`
	GlobalFlagsSet                           uint32                       `json:"global_flags_set"`
	CriticalSectionDefaultTimeout            uint32                       `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold               uint32                       `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold               uint32                       `json:"de_commit_total_free_threshold"`
	LockPrefixTable                          uint32                       `json:"lock_prefix_table"`
	MaximumAllocationSize                    uint32                       `json:"maximum_allocation_size"`
	VirtualMemoryThreshold                   uint32                       `json:"virtual_memory_threshold"`
	ProcessHeapFlags                         uint32                       `json:"process_heap_flags"`
	ProcessAffinityMask                      uint32                       `json:"process_affinity_mask"`
	CSDVersion                               uint16                       `json:"csd_version"`
	DependentLoadFlags                       uint16                       `json:"dependent_load_flags"`
	EditList                                 uint32                       `json:"edit_list"`
	SecurityCookie                           uint32                       `json:"security_cookie"`
	SEHandlerTable                           uint32                       `json:"se_handler_table"`
	SEHandlerCount                           uint32                       `json:"se_handler_count"`
	GuardCFCheckFunctionPointer              uint32                       `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer           uint32                       `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable                     uint32                       `json:"guard_cf_function_table"`
	GuardCFFunctionCount                     uint32                       `json:"guard_cf_function_count"`
	GuardFlags                               uint32                       `json:"guard_flags"`
	CodeIntegrity                            ImageLoadConfigCodeIntegrity `json:"code_integrity"`
	GuardAddressTakenIatEntryTable           uint32                       `json:"guard_address_taken_iat_entry_table"`
	GuardAddressTakenIatEntryCount           uint32                       `json:"guard_address_taken_iat_entry_count"`
	GuardLongJumpTargetTable                 uint32                       `json:"guard_long_jump_target_table"`
	GuardLongJumpTargetCount                 uint32                       `json:"guard_long_jump_target_count"`
	DynamicValueRelocTable                   uint32                       `json:"dynamic_value_reloc_table"`
	CHPEMetadataPointer                      uint32                       `json:"chpe_metadata_pointer"`
	GuardRFFailureRoutine                    uint32                       `json:"guard_rf_failure_routine"`
	GuardRFFailureRoutineFunctionPointer     uint32                       `json:"guard_rf_failure_routine_function_pointer"`
	DynamicValueRelocTableOffset             uint32                       `json:"dynamic_value_reloc_table_offset"`
	DynamicValueRelocTableSection            uint16                       `json:"dynamic_value_reloc_table_section"`
	Reserved2                                uint16                       `json:"reserved_2"`
	GuardRFVerifyStackPointerFunctionPointer uint32                       `json:"guard_rf_verify_stack_pointer_function_pointer"`
	HotPatchTableOffset                      uint32                       `json:"hot_patch_table_offset"`
	Reserved3                                uint32                       `json:"reserved_3"`
	EnclaveConfigurationPointer              uint32                       `json:"enclave_configuration_pointer"`
	VolatileMetadataPointer                  uint32                       `json:"volatile_metadata_pointer"`
	GuardEHContinuationTable                 uint32                       `json:"guard_eh_continuation_table"`
	GuardEHContinuationCount                 uint32                       `json:"guard_eh_continuation_count"`
	GuardXFGCheckFunctionPointer             uint32                       `json:"guard_xfg_check_function_pointer"`
	GuardXFGDispatchFunctionPointer          uint32                       `json:"guard_xfg_dispatch_function_pointer"`
	GuardXFGTableDispatchFunctionPointer     uint32                       `json:"guard_xfg_table_dispatch_function_pointer"`
}

// ImageLoadConfigDirectory32 Contains the load configuration data of an image for x86 binaries.
type ImageLoadConfigDirectory32 struct {
	// The actual size of the structure inclusive. May differ from the size
	// given in the data directory for Windows XP and earlier compatibility.
	Size uint32 `json:"size"`

	// Date and time stamp value.
	TimeDateStamp uint32 `json:"time_date_stamp"`

	// Major version number.
	MajorVersion uint16 `json:"major_version"`

	// Minor version number.
	MinorVersion uint16 `json:"minor_version"`

	// The global loader flags to clear for this process as the loader starts
	// the process.
	GlobalFlagsClear uint32 `json:"global_flags_clear"`

	// The global loader flags to set for this process as the loader starts the
	// process.
	GlobalFlagsSet uint32 `json:"global_flags_set"`

	// The default timeout value to use for this process's critical sections
	// that are abandoned.
	CriticalSectionDefaultTimeout uint32 `json:"critical_section_default_timeout"`

	// Memory that must be freed before it is returned to the system, in bytes.
	DeCommitFreeBlockThreshold uint32 `json:"de_commit_free_block_threshold"`

	// Total amount of free memory, in bytes.
	DeCommitTotalFreeThreshold uint32 `json:"de_commit_total_free_threshold"`

	// [x86 only] The VA of a list of addresses where the LOCK prefix is used so
	// that they can be replaced with NOP on single processor machines.
	LockPrefixTable uint32 `json:"lock_prefix_table"`

	// Maximum allocation size, in bytes.
	MaximumAllocationSize uint32 `json:"maximum_allocation_size"`

	// Maximum virtual memory size, in bytes.
	VirtualMemoryThreshold uint32 `json:"virtual_memory_threshold"`

	// Process heap flags that correspond to the first argument of the HeapCreate
	// function. These flags apply to the process heap that is created during
	// process startup.
	ProcessHeapFlags uint32 `json:"process_heap_flags"`

	// Setting this field to a non-zero value is equivalent to calling
	// SetProcessAffinityMask with this value during process startup (.exe only)
	ProcessAffinityMask uint32 `json:"process_affinity_mask"`

	// The service pack version identifier.
	CSDVersion uint16 `json:"csd_version"`

	// Must be zero.
	DependentLoadFlags uint16 `json:"dependent_load_flags"`

	// Reserved for use by the system.
	EditList uint32 `json:"edit_list"`

	// A pointer to a cookie that is used by Visual C++ or GS implementation.
	SecurityCookie uint32 `json:"security_cookie"`

	// [x86 only] The VA of the sorted table of RVAs of each valid, unique SE
	// handler in the image.
	SEHandlerTable uint32 `json:"se_handler_table"`

	// [x86 only] The count of unique handlers in the table.
	SEHandlerCount uint32 `json:"se_handler_count"`

	// The VA where Control Flow Guard check-function pointer is stored.
	GuardCFCheckFunctionPointer uint32 `json:"guard_cf_check_function_pointer"`

	// The VA where Control Flow Guard dispatch-function pointer is stored.
	GuardCFDispatchFunctionPointer uint32 `json:"guard_cf_dispatch_function_pointer"`

	// The VA of the sorted table of RVAs of each Control Flow Guard function in
	// the image.
	GuardCFFunctionTable uint32 `json:"guard_cf_function_table"`

	// The count of unique RVAs in the above table.
	GuardCFFunctionCount uint32 `json:"guard_cf_function_count"`

	// Control Flow Guard related flags.
	GuardFlags uint32 `json:"guard_flags"`

	// Code integrity information.
	CodeIntegrity ImageLoadConfigCodeIntegrity `json:"code_integrity"`

	// The VA where Control Flow Guard address taken IAT table is stored.
	GuardAddressTakenIatEntryTable uint32 `json:"guard_address_taken_iat_entry_table"`

	// The count of unique RVAs in the above table.
	GuardAddressTakenIatEntryCount uint32 `json:"guard_address_taken_iat_entry_count"`

	// The VA where Control Flow Guard long jump target table is stored.
	GuardLongJumpTargetTable uint32 `json:"guard_long_jump_target_table"`

	// The count of unique RVAs in the above table.
	GuardLongJumpTargetCount uint32 `json:"guard_long_jump_target_count"`

	DynamicValueRelocTable uint32 `json:"dynamic_value_reloc_table"`

	// Not sure when this was renamed from HybridMetadataPointer.
	CHPEMetadataPointer uint32 `json:"chpe_metadata_pointer"`

	GuardRFFailureRoutine                    uint32 `json:"guard_rf_failure_routine"`
	GuardRFFailureRoutineFunctionPointer     uint32 `json:"guard_rf_failure_routine_function_pointer"`
	DynamicValueRelocTableOffset             uint32 `json:"dynamic_value_reloc_table_offset"`
	DynamicValueRelocTableSection            uint16 `json:"dynamic_value_reloc_table_section"`
	Reserved2                                uint16 `json:"reserved_2"`
	GuardRFVerifyStackPointerFunctionPointer uint32 `json:"guard_rf_verify_stack_pointer_function_pointer"`
	HotPatchTableOffset                      uint32 `json:"hot_patch_table_offset"`
	Reserved3                                uint32 `json:"reserved_3"`
	EnclaveConfigurationPointer              uint32 `json:"enclave_configuration_pointer"`
	VolatileMetadataPointer                  uint32 `json:"volatile_metadata_pointer"`
	GuardEHContinuationTable                 uint32 `json:"guard_eh_continuation_table"`
	GuardEHContinuationCount                 uint32 `json:"guard_eh_continuation_count"`
	GuardXFGCheckFunctionPointer             uint32 `json:"guard_xfg_check_function_pointer"`
	GuardXFGDispatchFunctionPointer          uint32 `json:"guard_xfg_dispatch_function_pointer"`
	GuardXFGTableDispatchFunctionPointer     uint32 `json:"guard_xfg_table_dispatch_function_pointer"`
}

// ImageLoadConfigDirectory64v2 is the first structure for x64.
// No _IMAGE_LOAD_CONFIG_DIRECTORY64_V1 exists
type ImageLoadConfigDirectory64v2 struct {
	Size                          uint32 `json:"size"`
	TimeDateStamp                 uint32 `json:"time_date_stamp"`
	MajorVersion                  uint16 `json:"major_version"`
	MinorVersion                  uint16 `json:"minor_version"`
	GlobalFlagsClear              uint32 `json:"global_flags_clear"`
	GlobalFlagsSet                uint32 `json:"global_flags_set"`
	CriticalSectionDefaultTimeout uint32 `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold    uint64 `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold    uint64 `json:"de_commit_total_free_threshold"`
	LockPrefixTable               uint64 `json:"lock_prefix_table"`
	MaximumAllocationSize         uint64 `json:"maximum_allocation_size"`
	VirtualMemoryThreshold        uint64 `json:"virtual_memory_threshold"`
	ProcessAffinityMask           uint64 `json:"process_affinity_mask"`
	ProcessHeapFlags              uint32 `json:"process_heap_flags"`
	CSDVersion                    uint16 `json:"csd_version"`
	DependentLoadFlags            uint16 `json:"dependent_load_flags"`
	EditList                      uint64 `json:"edit_list"`
	SecurityCookie                uint64 `json:"security_cookie"`
	SEHandlerTable                uint64 `json:"se_handler_table"`
	SEHandlerCount                uint64 `json:"se_handler_count"`
}

// ImageLoadConfigDirectory64v3 added #pragma pack(4).
type ImageLoadConfigDirectory64v3 struct {
	Size                           uint32 `json:"size"`
	TimeDateStamp                  uint32 `json:"time_date_stamp"`
	MajorVersion                   uint16 `json:"major_version"`
	MinorVersion                   uint16 `json:"minor_version"`
	GlobalFlagsClear               uint32 `json:"global_flags_clear"`
	GlobalFlagsSet                 uint32 `json:"global_flags_set"`
	CriticalSectionDefaultTimeout  uint32 `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold     uint64 `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold     uint64 `json:"de_commit_total_free_threshold"`
	LockPrefixTable                uint64 `json:"lock_prefix_table"`
	MaximumAllocationSize          uint64 `json:"maximum_allocation_size"`
	VirtualMemoryThreshold         uint64 `json:"virtual_memory_threshold"`
	ProcessAffinityMask            uint64 `json:"process_affinity_mask"`
	ProcessHeapFlags               uint32 `json:"process_heap_flags"`
	CSDVersion                     uint16 `json:"csd_version"`
	DependentLoadFlags             uint16 `json:"dependent_load_flags"`
	EditList                       uint64 `json:"edit_list"`
	SecurityCookie                 uint64 `json:"security_cookie"`
	SEHandlerTable                 uint64 `json:"se_handler_table"`
	SEHandlerCount                 uint64 `json:"se_handler_count"`
	GuardCFCheckFunctionPointer    uint64 `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer uint64 `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable           uint64 `json:"guard_cf_function_table"`
	GuardCFFunctionCount           uint64 `json:"guard_cf_function_count"`
	GuardFlags                     uint32 `json:"guard_flags"`
}

// ImageLoadConfigDirectory64v4 for binaries compiled since Windows 10 build 9879 (or maybe earlier).
type ImageLoadConfigDirectory64v4 struct {
	Size                           uint32                       `json:"size"`
	TimeDateStamp                  uint32                       `json:"time_date_stamp"`
	MajorVersion                   uint16                       `json:"major_version"`
	MinorVersion                   uint16                       `json:"minor_version"`
	GlobalFlagsClear               uint32                       `json:"global_flags_clear"`
	GlobalFlagsSet                 uint32                       `json:"global_flags_set"`
	CriticalSectionDefaultTimeout  uint32                       `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold     uint64                       `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold     uint64                       `json:"de_commit_total_free_threshold"`
	LockPrefixTable                uint64                       `json:"lock_prefix_table"`
	MaximumAllocationSize          uint64                       `json:"maximum_allocation_size"`
	VirtualMemoryThreshold         uint64                       `json:"virtual_memory_threshold"`
	ProcessAffinityMask            uint64                       `json:"process_affinity_mask"`
	ProcessHeapFlags               uint32                       `json:"process_heap_flags"`
	CSDVersion                     uint16                       `json:"csd_version"`
	DependentLoadFlags             uint16                       `json:"dependent_load_flags"`
	EditList                       uint64                       `json:"edit_list"`
	SecurityCookie                 uint64                       `json:"security_cookie"`
	SEHandlerTable                 uint64                       `json:"se_handler_table"`
	SEHandlerCount                 uint64                       `json:"se_handler_count"`
	GuardCFCheckFunctionPointer    uint64                       `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer uint64                       `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable           uint64                       `json:"guard_cf_function_table"`
	GuardCFFunctionCount           uint64                       `json:"guard_cf_function_count"`
	GuardFlags                     uint32                       `json:"guard_flags"`
	CodeIntegrity                  ImageLoadConfigCodeIntegrity `json:"code_integrity"`
}

// ImageLoadConfigDirectory64v5 for binaries compiled since Windows 10 build 14286 (or maybe earlier).
type ImageLoadConfigDirectory64v5 struct {
	Size                           uint32                       `json:"size"`
	TimeDateStamp                  uint32                       `json:"time_date_stamp"`
	MajorVersion                   uint16                       `json:"major_version"`
	MinorVersion                   uint16                       `json:"minor_version"`
	GlobalFlagsClear               uint32                       `json:"global_flags_clear"`
	GlobalFlagsSet                 uint32                       `json:"global_flags_set"`
	CriticalSectionDefaultTimeout  uint32                       `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold     uint64                       `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold     uint64                       `json:"de_commit_total_free_threshold"`
	LockPrefixTable                uint64                       `json:"lock_prefix_table"`
	MaximumAllocationSize          uint64                       `json:"maximum_allocation_size"`
	VirtualMemoryThreshold         uint64                       `json:"virtual_memory_threshold"`
	ProcessAffinityMask            uint64                       `json:"process_affinity_mask"`
	ProcessHeapFlags               uint32                       `json:"process_heap_flags"`
	CSDVersion                     uint16                       `json:"csd_version"`
	DependentLoadFlags             uint16                       `json:"dependent_load_flags"`
	EditList                       uint64                       `json:"edit_list"`
	SecurityCookie                 uint64                       `json:"security_cookie"`
	SEHandlerTable                 uint64                       `json:"se_handler_table"`
	SEHandlerCount                 uint64                       `json:"se_handler_count"`
	GuardCFCheckFunctionPointer    uint64                       `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer uint64                       `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable           uint64                       `json:"guard_cf_function_table"`
	GuardCFFunctionCount           uint64                       `json:"guard_cf_function_count"`
	GuardFlags                     uint32                       `json:"guard_flags"`
	CodeIntegrity                  ImageLoadConfigCodeIntegrity `json:"code_integrity"`
	GuardAddressTakenIatEntryTable uint64                       `json:"guard_address_taken_iat_entry_table"`
	GuardAddressTakenIatEntryCount uint64                       `json:"guard_address_taken_iat_entry_count"`
	GuardLongJumpTargetTable       uint64                       `json:"guard_long_jump_target_table"`
	GuardLongJumpTargetCount       uint64                       `json:"guard_long_jump_target_count"`
}

// ImageLoadConfigDirectory64v6 for binaries compiled since Windows 10 build 14393 (or maybe earlier).
type ImageLoadConfigDirectory64v6 struct {
	Size                           uint32                       `json:"size"`
	TimeDateStamp                  uint32                       `json:"time_date_stamp"`
	MajorVersion                   uint16                       `json:"major_version"`
	MinorVersion                   uint16                       `json:"minor_version"`
	GlobalFlagsClear               uint32                       `json:"global_flags_clear"`
	GlobalFlagsSet                 uint32                       `json:"global_flags_set"`
	CriticalSectionDefaultTimeout  uint32                       `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold     uint64                       `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold     uint64                       `json:"de_commit_total_free_threshold"`
	LockPrefixTable                uint64                       `json:"lock_prefix_table"`
	MaximumAllocationSize          uint64                       `json:"maximum_allocation_size"`
	VirtualMemoryThreshold         uint64                       `json:"virtual_memory_threshold"`
	ProcessAffinityMask            uint64                       `json:"process_affinity_mask"`
	ProcessHeapFlags               uint32                       `json:"process_heap_flags"`
	CSDVersion                     uint16                       `json:"csd_version"`
	DependentLoadFlags             uint16                       `json:"dependent_load_flags"`
	EditList                       uint64                       `json:"edit_list"`
	SecurityCookie                 uint64                       `json:"security_cookie"`
	SEHandlerTable                 uint64                       `json:"se_handler_table"`
	SEHandlerCount                 uint64                       `json:"se_handler_count"`
	GuardCFCheckFunctionPointer    uint64                       `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer uint64                       `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable           uint64                       `json:"guard_cf_function_table"`
	GuardCFFunctionCount           uint64                       `json:"guard_cf_function_count"`
	GuardFlags                     uint32                       `json:"guard_flags"`
	CodeIntegrity                  ImageLoadConfigCodeIntegrity `json:"code_integrity"`
	GuardAddressTakenIatEntryTable uint64                       `json:"guard_address_taken_iat_entry_table"`
	GuardAddressTakenIatEntryCount uint64                       `json:"guard_address_taken_iat_entry_count"`
	GuardLongJumpTargetTable       uint64                       `json:"guard_long_jump_target_table"`
	GuardLongJumpTargetCount       uint64                       `json:"guard_long_jump_target_count"`
	DynamicValueRelocTable         uint64                       `json:"dynamic_value_reloc_table"`
	HybridMetadataPointer          uint64                       `json:"hybrid_metadata_pointer"`
}

// ImageLoadConfigDirectory64v7 for binaries compiled since Windows 10 build 14901 (or maybe earlier).
type ImageLoadConfigDirectory64v7 struct {
	Size                                 uint32                       `json:"size"`
	TimeDateStamp                        uint32                       `json:"time_date_stamp"`
	MajorVersion                         uint16                       `json:"major_version"`
	MinorVersion                         uint16                       `json:"minor_version"`
	GlobalFlagsClear                     uint32                       `json:"global_flags_clear"`
	GlobalFlagsSet                       uint32                       `json:"global_flags_set"`
	CriticalSectionDefaultTimeout        uint32                       `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold           uint64                       `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold           uint64                       `json:"de_commit_total_free_threshold"`
	LockPrefixTable                      uint64                       `json:"lock_prefix_table"`
	MaximumAllocationSize                uint64                       `json:"maximum_allocation_size"`
	VirtualMemoryThreshold               uint64                       `json:"virtual_memory_threshold"`
	ProcessAffinityMask                  uint64                       `json:"process_affinity_mask"`
	ProcessHeapFlags                     uint32                       `json:"process_heap_flags"`
	CSDVersion                           uint16                       `json:"csd_version"`
	DependentLoadFlags                   uint16                       `json:"dependent_load_flags"`
	EditList                             uint64                       `json:"edit_list"`
	SecurityCookie                       uint64                       `json:"security_cookie"`
	SEHandlerTable                       uint64                       `json:"se_handler_table"`
	SEHandlerCount                       uint64                       `json:"se_handler_count"`
	GuardCFCheckFunctionPointer          uint64                       `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer       uint64                       `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable                 uint64                       `json:"guard_cf_function_table"`
	GuardCFFunctionCount                 uint64                       `json:"guard_cf_function_count"`
	GuardFlags                           uint32                       `json:"guard_flags"`
	CodeIntegrity                        ImageLoadConfigCodeIntegrity `json:"code_integrity"`
	GuardAddressTakenIatEntryTable       uint64                       `json:"guard_address_taken_iat_entry_table"`
	GuardAddressTakenIatEntryCount       uint64                       `json:"guard_address_taken_iat_entry_count"`
	GuardLongJumpTargetTable             uint64                       `json:"guard_long_jump_target_table"`
	GuardLongJumpTargetCount             uint64                       `json:"guard_long_jump_target_count"`
	DynamicValueRelocTable               uint64                       `json:"dynamic_value_reloc_table"`
	CHPEMetadataPointer                  uint64                       `json:"chpe_metadata_pointer"`
	GuardRFFailureRoutine                uint64                       `json:"guard_rf_failure_routine"`
	GuardRFFailureRoutineFunctionPointer uint64                       `json:"guard_rf_failure_routine_function_pointer"`
	DynamicValueRelocTableOffset         uint32                       `json:"dynamic_value_reloc_table_offset"`
	DynamicValueRelocTableSection        uint16                       `json:"dynamic_value_reloc_table_section"`
	Reserved2                            uint16                       `json:"reserved_2"`
}

// ImageLoadConfigDirectory64v8 for binaries compiled since Windows 10 build 15002 (or maybe earlier).
// #pragma pack(4) available here.
type ImageLoadConfigDirectory64v8 struct {
	Size                                     uint32                       `json:"size"`
	TimeDateStamp                            uint32                       `json:"time_date_stamp"`
	MajorVersion                             uint16                       `json:"major_version"`
	MinorVersion                             uint16                       `json:"minor_version"`
	GlobalFlagsClear                         uint32                       `json:"global_flags_clear"`
	GlobalFlagsSet                           uint32                       `json:"global_flags_set"`
	CriticalSectionDefaultTimeout            uint32                       `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold               uint64                       `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold               uint64                       `json:"de_commit_total_free_threshold"`
	LockPrefixTable                          uint64                       `json:"lock_prefix_table"`
	MaximumAllocationSize                    uint64                       `json:"maximum_allocation_size"`
	VirtualMemoryThreshold                   uint64                       `json:"virtual_memory_threshold"`
	ProcessAffinityMask                      uint64                       `json:"process_affinity_mask"`
	ProcessHeapFlags                         uint32                       `json:"process_heap_flags"`
	CSDVersion                               uint16                       `json:"csd_version"`
	DependentLoadFlags                       uint16                       `json:"dependent_load_flags"`
	EditList                                 uint64                       `json:"edit_list"`
	SecurityCookie                           uint64                       `json:"security_cookie"`
	SEHandlerTable                           uint64                       `json:"se_handler_table"`
	SEHandlerCount                           uint64                       `json:"se_handler_count"`
	GuardCFCheckFunctionPointer              uint64                       `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer           uint64                       `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable                     uint64                       `json:"guard_cf_function_table"`
	GuardCFFunctionCount                     uint64                       `json:"guard_cf_function_count"`
	GuardFlags                               uint32                       `json:"guard_flags"`
	CodeIntegrity                            ImageLoadConfigCodeIntegrity `json:"code_integrity"`
	GuardAddressTakenIatEntryTable           uint64                       `json:"guard_address_taken_iat_entry_table"`
	GuardAddressTakenIatEntryCount           uint64                       `json:"guard_address_taken_iat_entry_count"`
	GuardLongJumpTargetTable                 uint64                       `json:"guard_long_jump_target_table"`
	GuardLongJumpTargetCount                 uint64                       `json:"guard_long_jump_target_count"`
	DynamicValueRelocTable                   uint64                       `json:"dynamic_value_reloc_table"`
	CHPEMetadataPointer                      uint64                       `json:"chpe_metadata_pointer"`
	GuardRFFailureRoutine                    uint64                       `json:"guard_rf_failure_routine"`
	GuardRFFailureRoutineFunctionPointer     uint64                       `json:"guard_rf_failure_routine_function_pointer"`
	DynamicValueRelocTableOffset             uint32                       `json:"dynamic_value_reloc_table_offset"`
	DynamicValueRelocTableSection            uint16                       `json:"dynamic_value_reloc_table_section"`
	Reserved2                                uint16                       `json:"reserved_2"`
	GuardRFVerifyStackPointerFunctionPointer uint64                       `json:"guard_rf_verify_stack_pointer_function_pointer"`
	HotPatchTableOffset                      uint32                       `json:"hot_patch_table_offset"`
}

// ImageLoadConfigDirectory64v9 for binaries compiled since Windows 10 build 15002 (or maybe earlier).
// #pragma pack(4) was taken.
type ImageLoadConfigDirectory64v9 struct {
	Size                                     uint32                       `json:"size"`
	TimeDateStamp                            uint32                       `json:"time_date_stamp"`
	MajorVersion                             uint16                       `json:"major_version"`
	MinorVersion                             uint16                       `json:"minor_version"`
	GlobalFlagsClear                         uint32                       `json:"global_flags_clear"`
	GlobalFlagsSet                           uint32                       `json:"global_flags_set"`
	CriticalSectionDefaultTimeout            uint32                       `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold               uint64                       `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold               uint64                       `json:"de_commit_total_free_threshold"`
	LockPrefixTable                          uint64                       `json:"lock_prefix_table"`
	MaximumAllocationSize                    uint64                       `json:"maximum_allocation_size"`
	VirtualMemoryThreshold                   uint64                       `json:"virtual_memory_threshold"`
	ProcessAffinityMask                      uint64                       `json:"process_affinity_mask"`
	ProcessHeapFlags                         uint32                       `json:"process_heap_flags"`
	CSDVersion                               uint16                       `json:"csd_version"`
	DependentLoadFlags                       uint16                       `json:"dependent_load_flags"`
	EditList                                 uint64                       `json:"edit_list"`
	SecurityCookie                           uint64                       `json:"security_cookie"`
	SEHandlerTable                           uint64                       `json:"se_handler_table"`
	SEHandlerCount                           uint64                       `json:"se_handler_count"`
	GuardCFCheckFunctionPointer              uint64                       `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer           uint64                       `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable                     uint64                       `json:"guard_cf_function_table"`
	GuardCFFunctionCount                     uint64                       `json:"guard_cf_function_count"`
	GuardFlags                               uint32                       `json:"guard_flags"`
	CodeIntegrity                            ImageLoadConfigCodeIntegrity `json:"code_integrity"`
	GuardAddressTakenIatEntryTable           uint64                       `json:"guard_address_taken_iat_entry_table"`
	GuardAddressTakenIatEntryCount           uint64                       `json:"guard_address_taken_iat_entry_count"`
	GuardLongJumpTargetTable                 uint64                       `json:"guard_long_jump_target_table"`
	GuardLongJumpTargetCount                 uint64                       `json:"guard_long_jump_target_count"`
	DynamicValueRelocTable                   uint64                       `json:"dynamic_value_reloc_table"`
	CHPEMetadataPointer                      uint64                       `json:"chpe_metadata_pointer"`
	GuardRFFailureRoutine                    uint64                       `json:"guard_rf_failure_routine"`
	GuardRFFailureRoutineFunctionPointer     uint64                       `json:"guard_rf_failure_routine_function_pointer"`
	DynamicValueRelocTableOffset             uint32                       `json:"dynamic_value_reloc_table_offset"`
	DynamicValueRelocTableSection            uint16                       `json:"dynamic_value_reloc_table_section"`
	Reserved2                                uint16                       `json:"reserved_2"`
	GuardRFVerifyStackPointerFunctionPointer uint64                       `json:"guard_rf_verify_stack_pointer_function_pointer"`
	HotPatchTableOffset                      uint32                       `json:"hot_patch_table_offset"`
	Reserved3                                uint32                       `json:"reserved_3"`
	EnclaveConfigurationPointer              uint64                       `json:"enclave_configuration_pointer"`
}

// ImageLoadConfigDirectory64v10 for binaries compiled since Windows 10 build 18362 (or maybe earlier).
type ImageLoadConfigDirectory64v10 struct {
	Size                                     uint32                       `json:"size"`
	TimeDateStamp                            uint32                       `json:"time_date_stamp"`
	MajorVersion                             uint16                       `json:"major_version"`
	MinorVersion                             uint16                       `json:"minor_version"`
	GlobalFlagsClear                         uint32                       `json:"global_flags_clear"`
	GlobalFlagsSet                           uint32                       `json:"global_flags_set"`
	CriticalSectionDefaultTimeout            uint32                       `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold               uint64                       `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold               uint64                       `json:"de_commit_total_free_threshold"`
	LockPrefixTable                          uint64                       `json:"lock_prefix_table"`
	MaximumAllocationSize                    uint64                       `json:"maximum_allocation_size"`
	VirtualMemoryThreshold                   uint64                       `json:"virtual_memory_threshold"`
	ProcessAffinityMask                      uint64                       `json:"process_affinity_mask"`
	ProcessHeapFlags                         uint32                       `json:"process_heap_flags"`
	CSDVersion                               uint16                       `json:"csd_version"`
	DependentLoadFlags                       uint16                       `json:"dependent_load_flags"`
	EditList                                 uint64                       `json:"edit_list"`
	SecurityCookie                           uint64                       `json:"security_cookie"`
	SEHandlerTable                           uint64                       `json:"se_handler_table"`
	SEHandlerCount                           uint64                       `json:"se_handler_count"`
	GuardCFCheckFunctionPointer              uint64                       `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer           uint64                       `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable                     uint64                       `json:"guard_cf_function_table"`
	GuardCFFunctionCount                     uint64                       `json:"guard_cf_function_count"`
	GuardFlags                               uint32                       `json:"guard_flags"`
	CodeIntegrity                            ImageLoadConfigCodeIntegrity `json:"code_integrity"`
	GuardAddressTakenIatEntryTable           uint64                       `json:"guard_address_taken_iat_entry_table"`
	GuardAddressTakenIatEntryCount           uint64                       `json:"guard_address_taken_iat_entry_count"`
	GuardLongJumpTargetTable                 uint64                       `json:"guard_long_jump_target_table"`
	GuardLongJumpTargetCount                 uint64                       `json:"guard_long_jump_target_count"`
	DynamicValueRelocTable                   uint64                       `json:"dynamic_value_reloc_table"`
	CHPEMetadataPointer                      uint64                       `json:"chpe_metadata_pointer"`
	GuardRFFailureRoutine                    uint64                       `json:"guard_rf_failure_routine"`
	GuardRFFailureRoutineFunctionPointer     uint64                       `json:"guard_rf_failure_routine_function_pointer"`
	DynamicValueRelocTableOffset             uint32                       `json:"dynamic_value_reloc_table_offset"`
	DynamicValueRelocTableSection            uint16                       `json:"dynamic_value_reloc_table_section"`
	Reserved2                                uint16                       `json:"reserved_2"`
	GuardRFVerifyStackPointerFunctionPointer uint64                       `json:"guard_rf_verify_stack_pointer_function_pointer"`
	HotPatchTableOffset                      uint32                       `json:"hot_patch_table_offset"`
	Reserved3                                uint32                       `json:"reserved_3"`
	EnclaveConfigurationPointer              uint64                       `json:"enclave_configuration_pointer"`
	VolatileMetadataPointer                  uint64                       `json:"volatile_metadata_pointer"`
}

// ImageLoadConfigDirectory64v11 for binaries compiled since Windows 10 build 19534 (or maybe earlier).
type ImageLoadConfigDirectory64v11 struct {
	Size                                     uint32                       `json:"size"`
	TimeDateStamp                            uint32                       `json:"time_date_stamp"`
	MajorVersion                             uint16                       `json:"major_version"`
	MinorVersion                             uint16                       `json:"minor_version"`
	GlobalFlagsClear                         uint32                       `json:"global_flags_clear"`
	GlobalFlagsSet                           uint32                       `json:"global_flags_set"`
	CriticalSectionDefaultTimeout            uint32                       `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold               uint64                       `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold               uint64                       `json:"de_commit_total_free_threshold"`
	LockPrefixTable                          uint64                       `json:"lock_prefix_table"`
	MaximumAllocationSize                    uint64                       `json:"maximum_allocation_size"`
	VirtualMemoryThreshold                   uint64                       `json:"virtual_memory_threshold"`
	ProcessAffinityMask                      uint64                       `json:"process_affinity_mask"`
	ProcessHeapFlags                         uint32                       `json:"process_heap_flags"`
	CSDVersion                               uint16                       `json:"csd_version"`
	DependentLoadFlags                       uint16                       `json:"dependent_load_flags"`
	EditList                                 uint64                       `json:"edit_list"`
	SecurityCookie                           uint64                       `json:"security_cookie"`
	SEHandlerTable                           uint64                       `json:"se_handler_table"`
	SEHandlerCount                           uint64                       `json:"se_handler_count"`
	GuardCFCheckFunctionPointer              uint64                       `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer           uint64                       `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable                     uint64                       `json:"guard_cf_function_table"`
	GuardCFFunctionCount                     uint64                       `json:"guard_cf_function_count"`
	GuardFlags                               uint32                       `json:"guard_flags"`
	CodeIntegrity                            ImageLoadConfigCodeIntegrity `json:"code_integrity"`
	GuardAddressTakenIatEntryTable           uint64                       `json:"guard_address_taken_iat_entry_table"`
	GuardAddressTakenIatEntryCount           uint64                       `json:"guard_address_taken_iat_entry_count"`
	GuardLongJumpTargetTable                 uint64                       `json:"guard_long_jump_target_table"`
	GuardLongJumpTargetCount                 uint64                       `json:"guard_long_jump_target_count"`
	DynamicValueRelocTable                   uint64                       `json:"dynamic_value_reloc_table"`
	CHPEMetadataPointer                      uint64                       `json:"chpe_metadata_pointer"`
	GuardRFFailureRoutine                    uint64                       `json:"guard_rf_failure_routine"`
	GuardRFFailureRoutineFunctionPointer     uint64                       `json:"guard_rf_failure_routine_function_pointer"`
	DynamicValueRelocTableOffset             uint32                       `json:"dynamic_value_reloc_table_offset"`
	DynamicValueRelocTableSection            uint16                       `json:"dynamic_value_reloc_table_section"`
	Reserved2                                uint16                       `json:"reserved_2"`
	GuardRFVerifyStackPointerFunctionPointer uint64                       `json:"guard_rf_verify_stack_pointer_function_pointer"`
	HotPatchTableOffset                      uint32                       `json:"hot_patch_table_offset"`
	Reserved3                                uint32                       `json:"reserved_3"`
	EnclaveConfigurationPointer              uint64                       `json:"enclave_configuration_pointer"`
	VolatileMetadataPointer                  uint64                       `json:"volatile_metadata_pointer"`
	GuardEHContinuationTable                 uint64                       `json:"guard_eh_continuation_table"`
	GuardEHContinuationCount                 uint64                       `json:"guard_eh_continuation_count"`
}

// ImageLoadConfigDirectory64v12 for binaries compiled since Visual C++ 2019 / RS5_IMAGE_LOAD_CONFIG_DIRECTORY64.
type ImageLoadConfigDirectory64v12 struct {
	Size                                     uint32                       `json:"size"`
	TimeDateStamp                            uint32                       `json:"time_date_stamp"`
	MajorVersion                             uint16                       `json:"major_version"`
	MinorVersion                             uint16                       `json:"minor_version"`
	GlobalFlagsClear                         uint32                       `json:"global_flags_clear"`
	GlobalFlagsSet                           uint32                       `json:"global_flags_set"`
	CriticalSectionDefaultTimeout            uint32                       `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold               uint64                       `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold               uint64                       `json:"de_commit_total_free_threshold"`
	LockPrefixTable                          uint64                       `json:"lock_prefix_table"`
	MaximumAllocationSize                    uint64                       `json:"maximum_allocation_size"`
	VirtualMemoryThreshold                   uint64                       `json:"virtual_memory_threshold"`
	ProcessAffinityMask                      uint64                       `json:"process_affinity_mask"`
	ProcessHeapFlags                         uint32                       `json:"process_heap_flags"`
	CSDVersion                               uint16                       `json:"csd_version"`
	DependentLoadFlags                       uint16                       `json:"dependent_load_flags"`
	EditList                                 uint64                       `json:"edit_list"`
	SecurityCookie                           uint64                       `json:"security_cookie"`
	SEHandlerTable                           uint64                       `json:"se_handler_table"`
	SEHandlerCount                           uint64                       `json:"se_handler_count"`
	GuardCFCheckFunctionPointer              uint64                       `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer           uint64                       `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable                     uint64                       `json:"guard_cf_function_table"`
	GuardCFFunctionCount                     uint64                       `json:"guard_cf_function_count"`
	GuardFlags                               uint32                       `json:"guard_flags"`
	CodeIntegrity                            ImageLoadConfigCodeIntegrity `json:"code_integrity"`
	GuardAddressTakenIatEntryTable           uint64                       `json:"guard_address_taken_iat_entry_table"`
	GuardAddressTakenIatEntryCount           uint64                       `json:"guard_address_taken_iat_entry_count"`
	GuardLongJumpTargetTable                 uint64                       `json:"guard_long_jump_target_table"`
	GuardLongJumpTargetCount                 uint64                       `json:"guard_long_jump_target_count"`
	DynamicValueRelocTable                   uint64                       `json:"dynamic_value_reloc_table"`
	CHPEMetadataPointer                      uint64                       `json:"chpe_metadata_pointer"`
	GuardRFFailureRoutine                    uint64                       `json:"guard_rf_failure_routine"`
	GuardRFFailureRoutineFunctionPointer     uint64                       `json:"guard_rf_failure_routine_function_pointer"`
	DynamicValueRelocTableOffset             uint32                       `json:"dynamic_value_reloc_table_offset"`
	DynamicValueRelocTableSection            uint16                       `json:"dynamic_value_reloc_table_section"`
	Reserved2                                uint16                       `json:"reserved_2"`
	GuardRFVerifyStackPointerFunctionPointer uint64                       `json:"guard_rf_verify_stack_pointer_function_pointer"`
	HotPatchTableOffset                      uint32                       `json:"hot_patch_table_offset"`
	Reserved3                                uint32                       `json:"reserved_3"`
	EnclaveConfigurationPointer              uint64                       `json:"enclave_configuration_pointer"`
	VolatileMetadataPointer                  uint64                       `json:"volatile_metadata_pointer"`
	GuardEHContinuationTable                 uint64                       `json:"guard_eh_continuation_table"`
	GuardEHContinuationCount                 uint64                       `json:"guard_eh_continuation_count"`
	GuardXFGCheckFunctionPointer             uint64                       `json:"guard_xfg_check_function_pointer"`
	GuardXFGDispatchFunctionPointer          uint64                       `json:"guard_xfg_dispatch_function_pointer"`
	GuardXFGTableDispatchFunctionPointer     uint64                       `json:"guard_xfg_table_dispatch_function_pointer"`
}

// ImageLoadConfigDirectory64 Contains the load configuration data of an image for x64 binaries.
type ImageLoadConfigDirectory64 struct {
	Size                                     uint32                       `json:"size"`
	TimeDateStamp                            uint32                       `json:"time_date_stamp"`
	MajorVersion                             uint16                       `json:"major_version"`
	MinorVersion                             uint16                       `json:"minor_version"`
	GlobalFlagsClear                         uint32                       `json:"global_flags_clear"`
	GlobalFlagsSet                           uint32                       `json:"global_flags_set"`
	CriticalSectionDefaultTimeout            uint32                       `json:"critical_section_default_timeout"`
	DeCommitFreeBlockThreshold               uint64                       `json:"de_commit_free_block_threshold"`
	DeCommitTotalFreeThreshold               uint64                       `json:"de_commit_total_free_threshold"`
	LockPrefixTable                          uint64                       `json:"lock_prefix_table"`
	MaximumAllocationSize                    uint64                       `json:"maximum_allocation_size"`
	VirtualMemoryThreshold                   uint64                       `json:"virtual_memory_threshold"`
	ProcessAffinityMask                      uint64                       `json:"process_affinity_mask"`
	ProcessHeapFlags                         uint32                       `json:"process_heap_flags"`
	CSDVersion                               uint16                       `json:"csd_version"`
	DependentLoadFlags                       uint16                       `json:"dependent_load_flags"`
	EditList                                 uint64                       `json:"edit_list"`
	SecurityCookie                           uint64                       `json:"security_cookie"`
	SEHandlerTable                           uint64                       `json:"se_handler_table"`
	SEHandlerCount                           uint64                       `json:"se_handler_count"`
	GuardCFCheckFunctionPointer              uint64                       `json:"guard_cf_check_function_pointer"`
	GuardCFDispatchFunctionPointer           uint64                       `json:"guard_cf_dispatch_function_pointer"`
	GuardCFFunctionTable                     uint64                       `json:"guard_cf_function_table"`
	GuardCFFunctionCount                     uint64                       `json:"guard_cf_function_count"`
	GuardFlags                               uint32                       `json:"guard_flags"`
	CodeIntegrity                            ImageLoadConfigCodeIntegrity `json:"code_integrity"`
	GuardAddressTakenIatEntryTable           uint64                       `json:"guard_address_taken_iat_entry_table"`
	GuardAddressTakenIatEntryCount           uint64                       `json:"guard_address_taken_iat_entry_count"`
	GuardLongJumpTargetTable                 uint64                       `json:"guard_long_jump_target_table"`
	GuardLongJumpTargetCount                 uint64                       `json:"guard_long_jump_target_count"`
	DynamicValueRelocTable                   uint64                       `json:"dynamic_value_reloc_table"`
	CHPEMetadataPointer                      uint64                       `json:"chpe_metadata_pointer"`
	GuardRFFailureRoutine                    uint64                       `json:"guard_rf_failure_routine"`
	GuardRFFailureRoutineFunctionPointer     uint64                       `json:"guard_rf_failure_routine_function_pointer"`
	DynamicValueRelocTableOffset             uint32                       `json:"dynamic_value_reloc_table_offset"`
	DynamicValueRelocTableSection            uint16                       `json:"dynamic_value_reloc_table_section"`
	Reserved2                                uint16                       `json:"reserved_2"`
	GuardRFVerifyStackPointerFunctionPointer uint64                       `json:"guard_rf_verify_stack_pointer_function_pointer"`
	HotPatchTableOffset                      uint32                       `json:"hot_patch_table_offset"`
	Reserved3                                uint32                       `json:"reserved_3"`
	EnclaveConfigurationPointer              uint64                       `json:"enclave_configuration_pointer"`
	VolatileMetadataPointer                  uint64                       `json:"volatile_metadata_pointer"`
	GuardEHContinuationTable                 uint64                       `json:"guard_eh_continuation_table"`
	GuardEHContinuationCount                 uint64                       `json:"guard_eh_continuation_count"`
	GuardXFGCheckFunctionPointer             uint64                       `json:"guard_xfg_check_function_pointer"`
	GuardXFGDispatchFunctionPointer          uint64                       `json:"guard_xfg_dispatch_function_pointer"`
	GuardXFGTableDispatchFunctionPointer     uint64                       `json:"guard_xfg_table_dispatch_function_pointer"`
}

type ImageCHPEMetadataX86v1 struct {
	Version                                  uint32 `json:"version"`
	CHPECodeAddressRangeOffset               uint32 `json:"chpe_code_address_range_offset"`
	CHPECodeAddressRangeCount                uint32 `json:"chpe_code_address_range_count"`
	WowA64ExceptionHandlerFunctionPtr        uint32 `json:"wow_a_64_exception_handler_function_ptr"`
	WowA64DispatchCallFunctionPtr            uint32 `json:"wow_a_64_dispatch_call_function_ptr"`
	WowA64DispatchIndirectCallFunctionPtr    uint32 `json:"wow_a_64_dispatch_indirect_call_function_ptr"`
	WowA64DispatchIndirectCallCfgFunctionPtr uint32 `json:"wow_a_64_dispatch_indirect_call_cfg_function_ptr"`
	WowA64DispatchRetFunctionPtr             uint32 `json:"wow_a_64_dispatch_ret_function_ptr"`
	WowA64DispatchRetLeafFunctionPtr         uint32 `json:"wow_a_64_dispatch_ret_leaf_function_ptr"`
	WowA64DispatchJumpFunctionPtr            uint32 `json:"wow_a_64_dispatch_jump_function_ptr"`
}

type ImageCHPEMetadataX86v2 struct {
	Version                                  uint32 `json:"version"`
	CHPECodeAddressRangeOffset               uint32 `json:"chpe_code_address_range_offset"`
	CHPECodeAddressRangeCount                uint32 `json:"chpe_code_address_range_count"`
	WowA64ExceptionHandlerFunctionPtr        uint32 `json:"wow_a_64_exception_handler_function_ptr"`
	WowA64DispatchCallFunctionPtr            uint32 `json:"wow_a_64_dispatch_call_function_ptr"`
	WowA64DispatchIndirectCallFunctionPtr    uint32 `json:"wow_a_64_dispatch_indirect_call_function_ptr"`
	WowA64DispatchIndirectCallCfgFunctionPtr uint32 `json:"wow_a_64_dispatch_indirect_call_cfg_function_ptr"`
	WowA64DispatchRetFunctionPtr             uint32 `json:"wow_a_64_dispatch_ret_function_ptr"`
	WowA64DispatchRetLeafFunctionPtr         uint32 `json:"wow_a_64_dispatch_ret_leaf_function_ptr"`
	WowA64DispatchJumpFunctionPtr            uint32 `json:"wow_a_64_dispatch_jump_function_ptr"`
	CompilerIATPointer                       uint32 `json:"compiler_iat_pointer"` // Present if Version >= 2
}

type ImageCHPEMetadataX86v3 struct {
	Version                                  uint32 `json:"version"`
	CHPECodeAddressRangeOffset               uint32 `json:"chpe_code_address_range_offset"`
	CHPECodeAddressRangeCount                uint32 `json:"chpe_code_address_range_count"`
	WowA64ExceptionHandlerFunctionPtr        uint32 `json:"wow_a_64_exception_handler_function_ptr"`
	WowA64DispatchCallFunctionPtr            uint32 `json:"wow_a_64_dispatch_call_function_ptr"`
	WowA64DispatchIndirectCallFunctionPtr    uint32 `json:"wow_a_64_dispatch_indirect_call_function_ptr"`
	WowA64DispatchIndirectCallCfgFunctionPtr uint32 `json:"wow_a_64_dispatch_indirect_call_cfg_function_ptr"`
	WowA64DispatchRetFunctionPtr             uint32 `json:"wow_a_64_dispatch_ret_function_ptr"`
	WowA64DispatchRetLeafFunctionPtr         uint32 `json:"wow_a_64_dispatch_ret_leaf_function_ptr"`
	WowA64DispatchJumpFunctionPtr            uint32 `json:"wow_a_64_dispatch_jump_function_ptr"`
	CompilerIATPointer                       uint32 `json:"compiler_iat_pointer"`
	WowA64RDTSCFunctionPtr                   uint32 `json:"wow_a_64_rdtsc_function_ptr"` // Present if Version >= 3
}

type CodeRange struct {
	Begin   uint32 `json:"begin"`
	Length  uint32 `json:"length"`
	Machine uint8  `json:"machine"`
}

type CompilerIAT struct {
	RVA         uint32 `json:"rva"`
	Value       uint32 `json:"value"`
	Description string `json:"description"`
}
type HybridPE struct {
	CHPEMetadata interface{}   `json:"chpe_metadata"`
	CodeRanges   []CodeRange   `json:"code_ranges"`
	CompilerIAT  []CompilerIAT `json:"compiler_iat"`
}

type ImageDynamicRelocationTable struct {
	Version uint32 `json:"version"`
	Size    uint32 `json:"size"`
	//  IMAGE_DYNAMIC_RELOCATION DynamicRelocations[0];
}

type ImageDynamicRelocation32 struct {
	Symbol        uint32 `json:"symbol"`
	BaseRelocSize uint32 `json:"base_reloc_size"`
	//  IMAGE_BASE_RELOCATION BaseRelocations[0];
}

type ImageDynamicRelocation64 struct {
	Symbol        uint64 `json:"symbol"`
	BaseRelocSize uint32 `json:"base_reloc_size"`
	//  IMAGE_BASE_RELOCATION BaseRelocations[0];
}

type ImageDynamicRelocation32v2 struct {
	HeaderSize    uint32 `json:"header_size"`
	FixupInfoSize uint32 `json:"fixup_info_size"`
	Symbol        uint32 `json:"symbol"`
	SymbolGroup   uint32 `json:"symbol_group"`
	Flags         uint32 `json:"flags"`
	// ...     variable length header fields
	// UCHAR   FixupInfo[FixupInfoSize]
}

type ImageDynamicRelocation64v2 struct {
	HeaderSize    uint32 `json:"header_size"`
	FixupInfoSize uint32 `json:"fixup_info_size"`
	Symbol        uint64 `json:"symbol"`
	SymbolGroup   uint32 `json:"symbol_group"`
	Flags         uint32 `json:"flags"`
	// ...     variable length header fields
	// UCHAR   FixupInfo[FixupInfoSize]
}

type ImagePrologueDynamicRelocationHeader struct {
	PrologueByteCount uint8 `json:"prologue_byte_count"`
	// UCHAR   PrologueBytes[PrologueByteCount];
}

type ImageEpilogueDynamicRelocationHeader struct {
	EpilogueCount               uint32 `json:"epilogue_count"`
	EpilogueByteCount           uint8  `json:"epilogue_byte_count"`
	BranchDescriptorElementSize uint8  `json:"branch_descriptor_element_size"`
	BranchDescriptorCount       uint8  `json:"branch_descriptor_count"`
	// UCHAR   BranchDescriptors[...];
	// UCHAR   BranchDescriptorBitMap[...];
}

type CFGFunction struct {
	Target      uint32 `json:"target"`
	Flags       *uint8 `json:"flags"`
	Description string `json:"description"`
}

type CFGIATEntry struct {
	RVA         uint32 `json:"rva"`
	IATValue    uint32 `json:"iat_value"`
	INTValue    uint32 `json:"int_value"`
	Description string `json:"description"`
}

type TypeOffset struct {
	Value               uint16 `json:"value"`
	Type                uint8  `json:"type"`
	DynamicSymbolOffset uint16 `json:"dynamic_symbol_offset"`
}

type RelocBlock struct {
	ImgBaseReloc ImageBaseRelocation `json:"img_base_reloc"`
	TypeOffsets  []TypeOffset        `json:"type_offsets"`
}
type RelocEntry struct {
	// Could be ImageDynamicRelocation32{} or ImageDynamicRelocation64{}
	ImgDynReloc interface{}  `json:"img_dyn_reloc"`
	RelocBlocks []RelocBlock `json:"reloc_blocks"`
}

// DVRT Dynamic Value Relocation Table.
type DVRT struct {
	ImgDynRelocTable ImageDynamicRelocationTable `json:"img_dyn_reloc_table"`
	Entries          []RelocEntry                `json:"entries"`
}

type Enclave struct {

	// Points to either ImageEnclaveConfig32{} or ImageEnclaveConfig64{}.
	Config interface{} `json:"config"`

	Imports []ImageEnclaveImport `json:"imports"`
}

type RangeTableEntry struct {
	Rva  uint32 `json:"rva"`
	Size uint32 `json:"size"`
}

type VolatileMetadata struct {
	Struct         ImageVolatileMetadata `json:"struct"`
	AccessRVATable []uint32              `json:"access_rva_table"`
	InfoRangeTable []RangeTableEntry     `json:"info_range_table"`
}
type LoadConfig struct {
	LoadCfgStruct    interface{}       `json:"struct"`
	SEH              []uint32          `json:"seh"`
	GFIDS            []CFGFunction     `json:"gfids"`
	CFGIAT           []CFGIATEntry     `json:"cfgiat"`
	CFGLongJump      []uint32          `json:"cfg_long_jump"`
	CHPE             *HybridPE         `json:"chpe"`
	DVRT             *DVRT             `json:"dvrt"`
	Enclave          *Enclave          `json:"enclave"`
	VolatileMetadata *VolatileMetadata `json:"volatile_metadata"`
}

// ImageLoadConfigCodeIntegrity Code Integrity in loadconfig (CI).
type ImageLoadConfigCodeIntegrity struct {
	Flags         uint16 `json:"flags"`   // Flags to indicate if CI information is available, etc.
	Catalog       uint16 `json:"catalog"` // 0xFFFF means not available
	CatalogOffset uint32 `json:"catalog_offset"`
	Reserved      uint32 `json:"reserved"` // Additional bitmask to be defined later
}

type ImageEnclaveConfig32 struct {

	// The size of the IMAGE_ENCLAVE_CONFIG32 structure, in bytes.
	Size uint32 `json:"size"`

	// The minimum size of the IMAGE_ENCLAVE_CONFIG32 structure that the image
	// loader must be able to process in order for the enclave to be usable.
	// This member allows an enclave to inform an earlier version of the image
	// loader that the image loader can safely load the enclave and ignore optional
	// members added to IMAGE_ENCLAVE_CONFIG32 for later versions of the enclave.

	// If the size of IMAGE_ENCLAVE_CONFIG32 that the image loader can process is
	// less than MinimumRequiredConfigSize, the enclave cannot be run securely.
	// If MinimumRequiredConfigSize is zero, the minimum size of the
	// IMAGE_ENCLAVE_CONFIG32 structure that the image loader must be able to
	// process in order for the enclave to be usable is assumed to be the size
	// of the structure through and including the MinimumRequiredConfigSize member.
	MinimumRequiredConfigSize uint32 `json:"minimum_required_config_size"`

	// A flag that indicates whether the enclave permits debugging.
	PolicyFlags uint32 `json:"policy_flags"`

	// The number of images in the array of images that the ImportList member
	// points to.
	NumberOfImports uint32 `json:"number_of_imports"`

	// The relative virtual address of the array of images that the enclave
	// image may import, with identity information for each image.
	ImportList uint32 `json:"import_list"`

	// The size of each image in the array of images that the ImportList member
	// points to.
	ImportEntrySize uint32 `json:"import_entry_size"`

	// The family identifier that the author of the enclave assigned to the enclave.
	FamilyID [ImageEnclaveShortIDLength]uint8 `json:"family_id"`

	// The image identifier that the author of the enclave assigned to the enclave.
	ImageID [ImageEnclaveShortIDLength]uint8 `json:"image_id"`

	// The version number that the author of the enclave assigned to the enclave.
	ImageVersion uint32 `json:"image_version"`

	// The security version number that the author of the enclave assigned to
	// the enclave.
	SecurityVersion uint32 `json:"security_version"`

	// The expected virtual size of the private address range for the enclave,
	// in bytes.
	EnclaveSize uint32 `json:"enclave_size"`

	// The maximum number of threads that can be created within the enclave.
	NumberOfThreads uint32 `json:"number_of_threads"`

	// A flag that indicates whether the image is suitable for use as the
	// primary image in the enclave.
	EnclaveFlags uint32 `json:"enclave_flags"`
}

type ImageEnclaveConfig64 struct {

	// The size of the IMAGE_ENCLAVE_CONFIG32 structure, in bytes.
	Size uint32 `json:"size"`

	// The minimum size of the IMAGE_ENCLAVE_CONFIG32 structure that the image
	// loader must be able to process in order for the enclave to be usable.
	// This member allows an enclave to inform an earlier version of the image
	// loader that the image loader can safely load the enclave and ignore
	// optional members added to IMAGE_ENCLAVE_CONFIG32 for later versions of
	// the enclave.

	// If the size of IMAGE_ENCLAVE_CONFIG32 that the image loader can process
	// is less than MinimumRequiredConfigSize, the enclave cannot be run securely.
	// If MinimumRequiredConfigSize is zero, the minimum size of the
	// IMAGE_ENCLAVE_CONFIG32 structure that the image loader must be able to
	// process in order for the enclave to be usable is assumed to be the size
	// of the structure through and including the MinimumRequiredConfigSize member.
	MinimumRequiredConfigSize uint32 `json:"minimum_required_config_size"`

	// A flag that indicates whether the enclave permits debugging.
	PolicyFlags uint32 `json:"policy_flags"`

	// The number of images in the array of images that the ImportList member
	// points to.
	NumberOfImports uint32 `json:"number_of_imports"`

	// The relative virtual address of the array of images that the enclave
	// image may import, with identity information for each image.
	ImportList uint32 `json:"import_list"`

	// The size of each image in the array of images that the ImportList member
	// points to.
	ImportEntrySize uint32 `json:"import_entry_size"`

	// The family identifier that the author of the enclave assigned to the enclave.
	FamilyID [ImageEnclaveShortIDLength]uint8 `json:"family_id"`

	// The image identifier that the author of the enclave assigned to the enclave.
	ImageID [ImageEnclaveShortIDLength]uint8 `json:"image_id"`

	// The version number that the author of the enclave assigned to the enclave.
	ImageVersion uint32 `json:"image_version"`

	// The security version number that the author of the enclave assigned to the enclave.
	SecurityVersion uint32 `json:"security_version"`

	// The expected virtual size of the private address range for the enclave,in bytes.
	EnclaveSize uint64 `json:"enclave_size"`

	// The maximum number of threads that can be created within the enclave.
	NumberOfThreads uint32 `json:"number_of_threads"`

	// A flag that indicates whether the image is suitable for use as the primary
	// image in the enclave.
	EnclaveFlags uint32 `json:"enclave_flags"`
}

// ImageEnclaveImport defines a entry in the array of images that an enclave can import.
type ImageEnclaveImport struct {

	// The type of identifier of the image that must match the value in the import record.
	MatchType uint32 `json:"match_type"`

	// The minimum enclave security version that each image must have for the
	// image to be imported successfully. The image is rejected unless its
	// enclave security version is equal to or greater than the minimum value in
	// the import record. Set the value in the import record to zero to turn off
	// the security version check.
	MinimumSecurityVersion uint32 `json:"minimum_security_version"`

	// The unique identifier of the primary module for the enclave, if the
	// MatchType member is IMAGE_ENCLAVE_IMPORT_MATCH_UNIQUE_ID. Otherwise,
	// the author identifier of the primary module for the enclave..
	UniqueOrAuthorID [ImageEnclaveShortIDLength]uint8 `json:"unique_or_author_id"`

	// The family identifier of the primary module for the enclave.
	FamilyID [ImageEnclaveShortIDLength]uint8 `json:"family_id"`

	// The image identifier of the primary module for the enclave.
	ImageID [ImageEnclaveShortIDLength]uint8 `json:"image_id"`

	// The relative virtual address of a NULL-terminated string that contains
	// the same value found in the import directory for the image.
	ImportName uint32 `json:"import_name"`

	// Reserved.
	Reserved uint32 `json:"reserved"`
}

type ImageVolatileMetadata struct {
	Size                       uint32
	Version                    uint32
	VolatileAccessTable        uint32
	VolatileAccessTableSize    uint32
	VolatileInfoRangeTable     uint32
	VolatileInfoRangeTableSize uint32
}

// The load configuration structure (IMAGE_LOAD_CONFIG_DIRECTORY) was formerly
// used in very limited cases in the Windows NT operating system itself to
// describe various features too difficult or too large to describe in the file

// header or optional header of the image. Current versions of the Microsoft
// linker and Windows XP and later versions of Windows use a new version of this
// structure for 32-bit x86-based systems that include reserved SEH technology.
// The data directory entry for a pre-reserved SEH load configuration structure
// must specify a particular size of the load configuration structure because
// the operating system loader always expects it to be a certain value. In that
// regard, the size is really only a version check. For compatibility with
// Windows XP and earlier versions of Windows, the size must be 64 for x86 images.
func (pe *File) parseLoadConfigDirectory(rva, size uint32) error {

	// As the load config structure changes over time,
	// we first read it size to figure out which one we have to cast against.
	fileOffset := pe.GetOffsetFromRva(rva)
	structSize, err := pe.ReadUint32(fileOffset)
	if err != nil {
		return err
	}

	// Use this helper function to print struct size.
	// PrintLoadConfigStruct()
	var loadCfg interface{}

	if pe.Is32 {
		switch int(structSize) {
		case 0x40:
			loadCfgv1 := ImageLoadConfigDirectory32v1{}
			err = pe.structUnpack(&loadCfgv1, fileOffset, structSize)
			loadCfg = loadCfgv1
		case 0x48:
			loadCfgv2 := ImageLoadConfigDirectory32v2{}
			err = pe.structUnpack(&loadCfgv2, fileOffset, structSize)
			loadCfg = loadCfgv2
		case 0x5c:
			loadCfgv3 := ImageLoadConfigDirectory32v3{}
			err = pe.structUnpack(&loadCfgv3, fileOffset, structSize)
			loadCfg = loadCfgv3
		case 0x68:
			loadCfgv4 := ImageLoadConfigDirectory32v4{}
			err = pe.structUnpack(&loadCfgv4, fileOffset, structSize)
			loadCfg = loadCfgv4
		case 0x78:
			loadCfgv5 := ImageLoadConfigDirectory32v5{}
			err = pe.structUnpack(&loadCfgv5, fileOffset, structSize)
			loadCfg = loadCfgv5
		case 0x80:
			loadCfgv6 := ImageLoadConfigDirectory32v6{}
			err = pe.structUnpack(&loadCfgv6, fileOffset, structSize)
			loadCfg = loadCfgv6
		case 0x90:
			loadCfgv7 := ImageLoadConfigDirectory32v7{}
			err = pe.structUnpack(&loadCfgv7, fileOffset, structSize)
			loadCfg = loadCfgv7
		case 0x98:
			loadCfgv8 := ImageLoadConfigDirectory32v8{}
			err = pe.structUnpack(&loadCfgv8, fileOffset, structSize)
			loadCfg = loadCfgv8
		case 0xa0:
			loadCfgv9 := ImageLoadConfigDirectory32v9{}
			err = pe.structUnpack(&loadCfgv9, fileOffset, structSize)
			loadCfg = loadCfgv9
		case 0xa4:
			loadCfgv10 := ImageLoadConfigDirectory32v10{}
			err = pe.structUnpack(&loadCfgv10, fileOffset, structSize)
			loadCfg = loadCfgv10
		case 0xac:
			loadCfgv11 := ImageLoadConfigDirectory32v11{}
			err = pe.structUnpack(&loadCfgv11, fileOffset, structSize)
			loadCfg = loadCfgv11
		case 0xb8:
			loadCfgv12 := ImageLoadConfigDirectory32v12{}
			err = pe.structUnpack(&loadCfgv12, fileOffset, structSize)
			loadCfg = loadCfgv12
		default:
			// We use the oldest load config.
			loadCfg32 := ImageLoadConfigDirectory32v1{}
			err = pe.structUnpack(&loadCfg32, fileOffset, structSize)
			loadCfg = loadCfg32
		}
	} else {
		switch int(structSize) {
		case 0x70:
			loadCfgv2 := ImageLoadConfigDirectory64v2{}
			err = pe.structUnpack(&loadCfgv2, fileOffset, structSize)
			loadCfg = loadCfgv2
		case 0x94:
			loadCfgv3 := ImageLoadConfigDirectory64v3{}
			err = pe.structUnpack(&loadCfgv3, fileOffset, structSize)
			loadCfg = loadCfgv3
		case 0xa0:
			loadCfgv4 := ImageLoadConfigDirectory64v4{}
			err = pe.structUnpack(&loadCfgv4, fileOffset, structSize)
			loadCfg = loadCfgv4
		case 0xc0:
			loadCfgv5 := ImageLoadConfigDirectory64v5{}
			err = pe.structUnpack(&loadCfgv5, fileOffset, structSize)
			loadCfg = loadCfgv5
		case 0xd0:
			loadCfgv6 := ImageLoadConfigDirectory64v6{}
			err = pe.structUnpack(&loadCfgv6, fileOffset, structSize)
			loadCfg = loadCfgv6
		case 0xe8:
			loadCfgv7 := ImageLoadConfigDirectory64v7{}
			err = pe.structUnpack(&loadCfgv7, fileOffset, structSize)
			loadCfg = loadCfgv7
		case 0xf4:
			loadCfgv8 := ImageLoadConfigDirectory64v8{}
			err = pe.structUnpack(&loadCfgv8, fileOffset, structSize)
			loadCfg = loadCfgv8
		case 0x100:
			loadCfgv9 := ImageLoadConfigDirectory64v9{}
			err = pe.structUnpack(&loadCfgv9, fileOffset, structSize)
			loadCfg = loadCfgv9
		case 0x108:
			loadCfgv10 := ImageLoadConfigDirectory64v10{}
			err = pe.structUnpack(&loadCfgv10, fileOffset, structSize)
			loadCfg = loadCfgv10
		case 0x118:
			loadCfgv11 := ImageLoadConfigDirectory64v11{}
			err = pe.structUnpack(&loadCfgv11, fileOffset, structSize)
			loadCfg = loadCfgv11
		case 0x130:
			loadCfgv12 := ImageLoadConfigDirectory64v12{}
			err = pe.structUnpack(&loadCfgv12, fileOffset, structSize)
			loadCfg = loadCfgv12
		default:
			// We use the oldest load config.
			loadCfg64 := ImageLoadConfigDirectory64v2{}
			err = pe.structUnpack(&loadCfg64, fileOffset, structSize)
			loadCfg = loadCfg64
		}
	}

	if err != nil {
		return err
	}

	// Save the load config struct.
	pe.HasLoadCFG = true
	pe.LoadConfig.LoadCfgStruct = loadCfg

	// Retrieve SEH handlers if there are any..
	if pe.Is32 {
		handlers := pe.getSEHHandlers()
		pe.LoadConfig.SEH = handlers
	}

	// Retrieve Control Flow Guard Function Targets if there are any.
	pe.LoadConfig.GFIDS = pe.getControlFlowGuardFunctions()

	// Retrieve Control Flow Guard IAT entries if there are any.
	pe.LoadConfig.CFGIAT = pe.getControlFlowGuardIAT()

	// Retrieve Long jump target functions if there are any.
	pe.LoadConfig.CFGLongJump = pe.getLongJumpTargetTable()

	// Retrieve compiled hybrid PE metadata if there are any.
	pe.LoadConfig.CHPE = pe.getHybridPE()

	// Retrieve dynamic value relocation table if there are any.
	pe.LoadConfig.DVRT = pe.getDynamicValueRelocTable()

	// Retrieve enclave configuration if there are any.
	pe.LoadConfig.Enclave = pe.getEnclaveConfiguration()

	// Retrieve volatile metadata table if there are any.
	pe.LoadConfig.VolatileMetadata = pe.getVolatileMetadata()

	return nil
}

// PrintLoadConfigStruct will print size of each load config struct.
func PrintLoadConfigStruct() {
	fmt.Printf("ImageLoadConfigDirectory32v1 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory32v1{}))
	fmt.Printf("ImageLoadConfigDirectory32v2 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory32v2{}))
	fmt.Printf("ImageLoadConfigDirectory32v3 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory32v3{}))
	fmt.Printf("ImageLoadConfigDirectory32v4 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory32v4{}))
	fmt.Printf("ImageLoadConfigDirectory32v5 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory32v5{}))
	fmt.Printf("ImageLoadConfigDirectory32v6 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory32v6{}))
	fmt.Printf("ImageLoadConfigDirectory32v7 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory32v7{}))
	fmt.Printf("ImageLoadConfigDirectory32v8 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory32v8{}))
	fmt.Printf("ImageLoadConfigDirectory32v9 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory32v9{}))
	fmt.Printf("ImageLoadConfigDirectory32v10 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory32v10{}))
	fmt.Printf("ImageLoadConfigDirectory32v11 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory32v11{}))
	fmt.Printf("ImageLoadConfigDirectory32v12 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory32v12{}))

	fmt.Printf("ImageLoadConfigDirectory64v2 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory64v2{}))
	fmt.Printf("ImageLoadConfigDirectory64v3 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory64v3{}))
	fmt.Printf("ImageLoadConfigDirectory64v4 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory64v4{}))
	fmt.Printf("ImageLoadConfigDirectory64v5 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory64v5{}))
	fmt.Printf("ImageLoadConfigDirectory64v6 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory64v6{}))
	fmt.Printf("ImageLoadConfigDirectory64v7 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory64v7{}))
	fmt.Printf("ImageLoadConfigDirectory64v8 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory64v8{}))
	fmt.Printf("ImageLoadConfigDirectory64v9 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory64v9{}))
	fmt.Printf("ImageLoadConfigDirectory64v10 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory64v10{}))
	fmt.Printf("ImageLoadConfigDirectory64v11 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory64v11{}))
	fmt.Printf("ImageLoadConfigDirectory64v12 size : 0x%x\n", binary.Size(ImageLoadConfigDirectory64v12{}))
}

// StringifyGuardFlags returns list of strings which describes the GuardFlags.
func StringifyGuardFlags(flags uint32) []string {
	var values []string
	guardFlagMap := map[uint32]string{
		ImageGuardCfInstrumented:                 "Instrumented",
		ImageGuardCfWInstrumented:                "WriteInstrumented",
		ImageGuardCfFunctionTablePresent:         "TargetMetadata",
		ImageGuardSecurityCookieUnused:           "SecurityCookieUnused",
		ImageGuardProtectDelayloadIAT:            "DelayloadIAT",
		ImageGuardDelayloadIATInItsOwnSection:    "DelayloadIATInItsOwnSection",
		ImageGuardCfExportSuppressionInfoPresent: "ExportSuppressionInfoPresent",
		ImageGuardCfEnableExportSuppression:      "EnableExportSuppression",
		ImageGuardCfLongjumpTablePresent:         "LongjumpTablePresent",
	}

	for k, s := range guardFlagMap {
		if k&flags != 0 {
			values = append(values, s)
		}
	}
	return values
}

func (pe *File) getSEHHandlers() []uint32 {

	var handlers []uint32
	v := reflect.ValueOf(pe.LoadConfig.LoadCfgStruct)

	// SEHandlerCount is found in index 19 of the struct.
	if v.NumField() > 19 {
		SEHandlerCount := uint32(v.Field(19).Uint())
		if SEHandlerCount > 0 {
			SEHandlerTable := uint32(v.Field(18).Uint())
			imageBase := pe.NtHeader.OptionalHeader.(ImageOptionalHeader32).ImageBase
			rva := SEHandlerTable - imageBase
			for i := uint32(0); i < SEHandlerCount; i++ {
				offset := pe.GetOffsetFromRva(rva + i*4)
				handler, err := pe.ReadUint32(offset)
				if err != nil {
					return handlers
				}

				handlers = append(handlers, handler)
			}
		}
	}

	return handlers
}

func (pe *File) getControlFlowGuardFunctions() []CFGFunction {

	v := reflect.ValueOf(pe.LoadConfig.LoadCfgStruct)
	var GFIDS []CFGFunction
	var err error

	// GuardCFFunctionCount is found in index 23 of the struct.
	if v.NumField() > 23 {
		// The GFIDS table is an array of 4 + n bytes, where n is given by :
		// ((GuardFlags & IMAGE_GUARD_CF_FUNCTION_TABLE_SIZE_MASK) >>
		// IMAGE_GUARD_CF_FUNCTION_TABLE_SIZE_SHIFT).

		// This allows for extra metadata to be attached to CFG call targets in
		// the future. The only currently defined metadata is an optional 1-byte
		// extra flags field (“GFIDS flags”) that is attached to each GFIDS
		// entry if any call targets have metadata.
		GuardFlags := v.Field(24).Uint()
		n := (GuardFlags & ImageGuardCfFunctionTableSizeMask) >>
			ImageGuardCfFunctionTableSizeShift
		GuardCFFunctionCount := v.Field(23).Uint()
		if GuardCFFunctionCount > 0 {
			if pe.Is32 {
				GuardCFFunctionTable := uint32(v.Field(22).Uint())
				imageBase := pe.NtHeader.OptionalHeader.(ImageOptionalHeader32).ImageBase
				rva := GuardCFFunctionTable - imageBase
				offset := pe.GetOffsetFromRva(rva)
				for i := uint32(1); i <= uint32(GuardCFFunctionCount); i++ {
					cfgFunction := CFGFunction{}
					var cfgFlags uint8
					cfgFunction.Target, err = pe.ReadUint32(offset)
					if err != nil {
						return GFIDS
					}
					if n > 0 {
						err = pe.structUnpack(&cfgFlags, offset+4, uint32(n))
						if err != nil {
							return GFIDS
						}
						cfgFunction.Flags = &cfgFlags
						if cfgFlags == ImageGuardFlagFIDSuppressed ||
							cfgFlags == ImageGuardFlagExportSuppressed {
							exportName := pe.GetExportFunctionByRVA(cfgFunction.Target)
							cfgFunction.Description = exportName.Name
						}
					}

					GFIDS = append(GFIDS, cfgFunction)
					offset += 4 + uint32(n)
				}
			} else {
				GuardCFFunctionTable := v.Field(22).Uint()
				imageBase := pe.NtHeader.OptionalHeader.(ImageOptionalHeader64).ImageBase
				rva := uint32(GuardCFFunctionTable - imageBase)
				offset := pe.GetOffsetFromRva(rva)
				for i := uint64(1); i <= GuardCFFunctionCount; i++ {
					var cfgFlags uint8
					cfgFunction := CFGFunction{}
					cfgFunction.Target, err = pe.ReadUint32(offset)
					if err != nil {
						return GFIDS
					}
					if n > 0 {
						pe.structUnpack(&cfgFlags, offset+4, uint32(n))
						cfgFunction.Flags = &cfgFlags
						if cfgFlags == ImageGuardFlagFIDSuppressed ||
							cfgFlags == ImageGuardFlagExportSuppressed {
							exportName := pe.GetExportFunctionByRVA(cfgFunction.Target)
							cfgFunction.Description = exportName.Name
						}
					}

					GFIDS = append(GFIDS, cfgFunction)
					offset += 4 + uint32(n)
				}
			}

		}
	}
	return GFIDS
}

func (pe *File) getControlFlowGuardIAT() []CFGIATEntry {

	v := reflect.ValueOf(pe.LoadConfig.LoadCfgStruct)
	var GFGIAT []CFGIATEntry
	var err error

	// GuardAddressTakenIatEntryCount is found in index 27 of the struct.
	if v.NumField() > 27 {
		// An image that supports CFG ES includes a GuardAddressTakenIatEntryTable
		// whose count is provided by the GuardAddressTakenIatEntryCount as part
		// of its load configuration directory. This table is structurally
		// formatted the same as the GFIDS table. It uses the same GuardFlags
		// IMAGE_GUARD_CF_FUNCTION_TABLE_SIZE_MASK mechanism to encode extra
		// optional metadata bytes in the address taken IAT table, though all
		// metadata bytes must be zero for the address taken IAT table and are
		// reserved.
		GuardFlags := v.Field(24).Uint()
		n := (GuardFlags & ImageGuardCfFunctionTableSizeMask) >>
			ImageGuardCfFunctionTableSizeShift
		GuardAddressTakenIatEntryCount := v.Field(27).Uint()
		if GuardAddressTakenIatEntryCount > 0 {
			if pe.Is32 {
				GuardAddressTakenIatEntryTable := uint32(v.Field(26).Uint())
				imageBase := pe.NtHeader.OptionalHeader.(ImageOptionalHeader32).ImageBase
				rva := GuardAddressTakenIatEntryTable - imageBase
				offset := pe.GetOffsetFromRva(rva)
				for i := uint32(1); i <= uint32(GuardAddressTakenIatEntryCount); i++ {
					cfgIATEntry := CFGIATEntry{}
					cfgIATEntry.RVA, err = pe.ReadUint32(offset)
					if err != nil {
						return GFGIAT
					}
					imp, index := pe.GetImportEntryInfoByRVA(cfgIATEntry.RVA)
					if len(imp.Functions) != 0 {
						cfgIATEntry.INTValue = uint32(imp.Functions[index].OriginalThunkValue)
						cfgIATEntry.IATValue = uint32(imp.Functions[index].ThunkValue)
						cfgIATEntry.Description = imp.Name + "!" + imp.Functions[index].Name
					}
					GFGIAT = append(GFGIAT, cfgIATEntry)
					offset += 4 + uint32(n)
				}
			} else {
				GuardAddressTakenIatEntryTable := v.Field(26).Uint()
				imageBase := pe.NtHeader.OptionalHeader.(ImageOptionalHeader64).ImageBase
				rva := uint32(GuardAddressTakenIatEntryTable - imageBase)
				offset := pe.GetOffsetFromRva(rva)
				for i := uint64(1); i <= GuardAddressTakenIatEntryCount; i++ {
					cfgIATEntry := CFGIATEntry{}
					cfgIATEntry.RVA, err = pe.ReadUint32(offset)
					if err != nil {
						return GFGIAT
					}
					imp, index := pe.GetImportEntryInfoByRVA(cfgIATEntry.RVA)
					if len(imp.Functions) != 0 {
						cfgIATEntry.INTValue = uint32(imp.Functions[index].OriginalThunkValue)
						cfgIATEntry.IATValue = uint32(imp.Functions[index].ThunkValue)
						cfgIATEntry.Description = imp.Name + "!" + imp.Functions[index].Name
					}

					GFGIAT = append(GFGIAT, cfgIATEntry)
					GFGIAT = append(GFGIAT, cfgIATEntry)
					offset += 4 + uint32(n)
				}
			}

		}
	}
	return GFGIAT
}

func (pe *File) getLongJumpTargetTable() []uint32 {

	v := reflect.ValueOf(pe.LoadConfig.LoadCfgStruct)
	var longJumpTargets []uint32

	// GuardLongJumpTargetCount is found in index 29 of the struct.
	if v.NumField() > 29 {
		// The long jump table represents a sorted array of RVAs that are valid
		// long jump targets. If a long jump target module sets
		// IMAGE_GUARD_CF_LONGJUMP_TABLE_PRESENT in its GuardFlags field, then
		// all long jump targets must be enumerated in the LongJumpTargetTable.
		GuardFlags := v.Field(24).Uint()
		n := (GuardFlags & ImageGuardCfFunctionTableSizeMask) >>
			ImageGuardCfFunctionTableSizeShift
		GuardLongJumpTargetCount := v.Field(29).Uint()
		if GuardLongJumpTargetCount > 0 {
			if pe.Is32 {
				GuardLongJumpTargetTable := uint32(v.Field(28).Uint())
				imageBase := pe.NtHeader.OptionalHeader.(ImageOptionalHeader32).ImageBase
				rva := GuardLongJumpTargetTable - imageBase
				offset := pe.GetOffsetFromRva(rva)
				for i := uint32(1); i <= uint32(GuardLongJumpTargetCount); i++ {
					target, err := pe.ReadUint32(offset)
					if err != nil {
						return longJumpTargets
					}
					longJumpTargets = append(longJumpTargets, target)
					offset += 4 + uint32(n)
				}
			} else {
				GuardLongJumpTargetTable := v.Field(26).Uint()
				imageBase := pe.NtHeader.OptionalHeader.(ImageOptionalHeader64).ImageBase
				rva := uint32(GuardLongJumpTargetTable - imageBase)
				offset := pe.GetOffsetFromRva(rva)
				for i := uint64(1); i <= GuardLongJumpTargetCount; i++ {
					target, err := pe.ReadUint32(offset)
					if err != nil {
						return longJumpTargets
					}
					longJumpTargets = append(longJumpTargets, target)
					offset += 4 + uint32(n)
				}
			}

		}
	}
	return longJumpTargets
}

func (pe *File) getHybridPE() *HybridPE {
	v := reflect.ValueOf(pe.LoadConfig.LoadCfgStruct)
	hybridPE := HybridPE{}

	// CHPEMetadataPointer is found in index 31 of the struct.
	if v.NumField() <= 30 {
		return nil
	}

	CHPEMetadataPointer := v.Field(31).Uint()
	if CHPEMetadataPointer == 0 {
		return nil
	}
	var rva uint32
	if pe.Is32 {
		imageBase := pe.NtHeader.OptionalHeader.(ImageOptionalHeader32).ImageBase
		rva = uint32(CHPEMetadataPointer) - imageBase
	} else {
		imageBase := pe.NtHeader.OptionalHeader.(ImageOptionalHeader64).ImageBase
		rva = uint32(CHPEMetadataPointer - imageBase)
	}

	// As the image CHPE metadata structure changes over time,
	// we first read its version to figure out which one we have to
	// cast against.
	fileOffset := pe.GetOffsetFromRva(rva)
	version, err := pe.ReadUint32(fileOffset)
	if err != nil {
		return nil
	}

	var ImageCHPEMetaX86 interface{}

	switch version {
	case 0x1:
		ImageCHPEMetaX86v1 := ImageCHPEMetadataX86v1{}
		structSize := uint32(binary.Size(ImageCHPEMetaX86v1))
		err = pe.structUnpack(&ImageCHPEMetaX86v1, fileOffset, structSize)
		if err != nil {
			return nil
		}
		ImageCHPEMetaX86 = ImageCHPEMetaX86v1
	case 0x2:
		ImageCHPEMetaX86v2 := ImageCHPEMetadataX86v2{}
		structSize := uint32(binary.Size(ImageCHPEMetaX86v2))
		err = pe.structUnpack(&ImageCHPEMetaX86v2, fileOffset, structSize)
		if err != nil {
			return nil
		}
		ImageCHPEMetaX86 = ImageCHPEMetaX86v2
	case 0x3:
	default:
		ImageCHPEMetaX86v3 := ImageCHPEMetadataX86v3{}
		structSize := uint32(binary.Size(ImageCHPEMetaX86v3))
		err = pe.structUnpack(&ImageCHPEMetaX86v3, fileOffset, structSize)
		if err != nil {
			return nil
		}
		ImageCHPEMetaX86 = ImageCHPEMetaX86v3
	}

	hybridPE.CHPEMetadata = ImageCHPEMetaX86

	v = reflect.ValueOf(ImageCHPEMetaX86)
	CHPECodeAddressRangeOffset := uint32(v.Field(1).Uint())
	CHPECodeAddressRangeCount := int(v.Field(2).Uint())

	// Code Ranges

	/*
		typedef struct _IMAGE_CHPE_RANGE_ENTRY {
			union {
				ULONG StartOffset;
				struct {
					ULONG NativeCode : 1;
					ULONG AddressBits : 31;
				} DUMMYSTRUCTNAME;
			} DUMMYUNIONNAME;

			ULONG Length;
		} IMAGE_CHPE_RANGE_ENTRY, *PIMAGE_CHPE_RANGE_ENTRY;
	*/

	rva = CHPECodeAddressRangeOffset
	for i := 0; i < CHPECodeAddressRangeCount; i++ {

		codeRange := CodeRange{}
		fileOffset := pe.GetOffsetFromRva(rva)
		begin, err := pe.ReadUint32(fileOffset)
		if err != nil {
			break
		}

		if begin&1 == 1 {
			codeRange.Machine = 1
			begin = uint32(int(begin) & ^1)
		}
		codeRange.Begin = begin

		fileOffset += 4
		size, err := pe.ReadUint32(fileOffset)
		if err != nil {
			break
		}
		codeRange.Length = size

		hybridPE.CodeRanges = append(hybridPE.CodeRanges, codeRange)
		rva += 8
	}

	// Compiler IAT
	CompilerIATPointer := uint32(v.Field(10).Uint())
	if CompilerIATPointer != 0 {
		rva := CompilerIATPointer
		for i := 0; i < 1024; i++ {
			compilerIAT := CompilerIAT{}
			compilerIAT.RVA = rva
			fileOffset = pe.GetOffsetFromRva(rva)
			compilerIAT.Value, err = pe.ReadUint32(fileOffset)
			if err != nil {
				break
			}

			pe.LoadConfig.CHPE.CompilerIAT = append(
				pe.LoadConfig.CHPE.CompilerIAT, compilerIAT)
			rva += 4
		}
	}
	return &hybridPE
}

func (pe *File) getDynamicValueRelocTable() *DVRT {

	var structSize uint32
	var imgDynRelocSize uint32
	dvrt := DVRT{}
	imgDynRelocTable := ImageDynamicRelocationTable{}

	v := reflect.ValueOf(pe.LoadConfig.LoadCfgStruct)
	if v.NumField() <= 35 {
		return nil
	}

	DynamicValueRelocTableOffset := v.Field(34).Uint()
	DynamicValueRelocTableSection := v.Field(35).Uint()
	if DynamicValueRelocTableOffset == 0 || DynamicValueRelocTableSection == 0 {
		return nil
	}

	section := pe.getSectionByName(".reloc")
	if section == nil {
		return nil
	}

	// Get the dynamic value relocation table.
	rva := section.VirtualAddress + uint32(DynamicValueRelocTableOffset)
	offset := pe.GetOffsetFromRva(rva)
	structSize = uint32(binary.Size(imgDynRelocTable))
	err := pe.structUnpack(&imgDynRelocTable, offset, structSize)
	if err != nil {
		return nil
	}

	dvrt.ImgDynRelocTable = imgDynRelocTable
	offset += structSize

	// Get dynamic relocation entries according to version.
	switch imgDynRelocTable.Version {
	case 1:
		relocTableIt := uint32(0)
		baseBlockSize := uint32(0)

		// Iterate over our dynamic reloc table entries.
		for relocTableIt < imgDynRelocTable.Size {

			relocEntry := RelocEntry{}
			if pe.Is32 {
				imgDynReloc := ImageDynamicRelocation32{}
				imgDynRelocSize = uint32(binary.Size(imgDynReloc))
				err = pe.structUnpack(&imgDynReloc, offset, imgDynRelocSize)
				if err != nil {
					return nil
				}
				relocEntry.ImgDynReloc = imgDynReloc
				baseBlockSize = imgDynReloc.BaseRelocSize
			} else {
				imgDynReloc := ImageDynamicRelocation64{}
				imgDynRelocSize = uint32(binary.Size(imgDynReloc))
				err = pe.structUnpack(&imgDynReloc, offset, imgDynRelocSize)
				if err != nil {
					return nil
				}
				relocEntry.ImgDynReloc = imgDynReloc
				baseBlockSize = imgDynReloc.BaseRelocSize
			}
			offset += imgDynRelocSize
			relocTableIt += imgDynRelocSize

			// Iterate over reach block.
			blockIt := uint32(0)
			for blockIt < baseBlockSize-imgDynRelocSize {
				relocBlock := RelocBlock{}

				baseReloc := ImageBaseRelocation{}
				structSize = uint32(binary.Size(baseReloc))
				err = pe.structUnpack(&baseReloc, offset, structSize)
				if err != nil {
					return nil
				}

				relocBlock.ImgBaseReloc = baseReloc
				offset += structSize
				numTypeOffsets := (baseReloc.SizeOfBlock - structSize) / 2
				for i := uint32(0); i < numTypeOffsets; i++ {
					typeOffset := TypeOffset{}
					typeOffset.Value, err = pe.ReadUint16(offset)
					if err != nil {
						return nil
					}

					typeOffset.DynamicSymbolOffset = typeOffset.Value & 0xfff
					typeOffset.Type = uint8(typeOffset.Value & 0xf000 >> 12)
					offset += 2

					// Padding at the end of the block ?
					if (TypeOffset{}) == typeOffset && i+1 == numTypeOffsets {
						break
					}

					relocBlock.TypeOffsets = append(relocBlock.TypeOffsets, typeOffset)
				}

				blockIt += baseReloc.SizeOfBlock
				relocEntry.RelocBlocks = append(relocEntry.RelocBlocks, relocBlock)
			}

			dvrt.Entries = append(dvrt.Entries, relocEntry)
			relocTableIt += baseBlockSize
		}
	case 2:
		fmt.Print("Got version 2 !")
	}

	return &dvrt
}

func (pe *File) getEnclaveConfiguration() *Enclave {

	enclave := Enclave{}

	v := reflect.ValueOf(pe.LoadConfig.LoadCfgStruct)
	if v.NumField() <= 40 {
		return nil
	}

	EnclaveConfigurationPointer := v.Field(40).Uint()
	if EnclaveConfigurationPointer == 0 {
		return nil
	}

	if pe.Is32 {
		imgEnclaveCfg := ImageEnclaveConfig32{}
		imgEnclaveCfgSize := uint32(binary.Size(imgEnclaveCfg))
		imageBase := pe.NtHeader.OptionalHeader.(ImageOptionalHeader32).ImageBase
		rva := uint32(EnclaveConfigurationPointer) - imageBase
		offset := pe.GetOffsetFromRva(rva)
		err := pe.structUnpack(&imgEnclaveCfg, offset, imgEnclaveCfgSize)
		if err != nil {
			return nil
		}
		enclave.Config = imgEnclaveCfg
	} else {
		imgEnclaveCfg := ImageEnclaveConfig64{}
		imgEnclaveCfgSize := uint32(binary.Size(imgEnclaveCfg))
		imageBase := pe.NtHeader.OptionalHeader.(ImageOptionalHeader64).ImageBase
		rva := uint32(EnclaveConfigurationPointer - imageBase)
		offset := pe.GetOffsetFromRva(rva)
		err := pe.structUnpack(&imgEnclaveCfg, offset, imgEnclaveCfgSize)
		if err != nil {
			return nil
		}
		enclave.Config = imgEnclaveCfg
	}

	// Get the array of images that an enclave can import.
	val := reflect.ValueOf(enclave.Config)
	ImportListRVA := val.FieldByName("ImportList").Interface().(uint32)
	NumberOfImports := val.FieldByName("NumberOfImports").Interface().(uint32)

	offset := pe.GetOffsetFromRva(ImportListRVA)
	for i := uint32(0); i < NumberOfImports; i++ {
		imgEncImp := ImageEnclaveImport{}
		imgEncImpSize := uint32(binary.Size(imgEncImp))
		err := pe.structUnpack(&imgEncImp, offset, imgEncImpSize)
		if err != nil {
			return nil
		}

		offset += imgEncImpSize
		enclave.Imports = append(enclave.Imports, imgEncImp)
	}

	return nil
}

func (pe *File) getVolatileMetadata() *VolatileMetadata {

	volatileMeta := VolatileMetadata{}
	imgVolatileMeta := ImageVolatileMetadata{}
	rva := uint32(0)

	v := reflect.ValueOf(pe.LoadConfig.LoadCfgStruct)
	if v.NumField() <= 41 {
		return nil
	}

	VolatileMetadataPointer := v.Field(41).Uint()
	if VolatileMetadataPointer == 0 {
		return nil
	}

	if pe.Is32 {
		imageBase := pe.NtHeader.OptionalHeader.(ImageOptionalHeader32).ImageBase
		rva = uint32(VolatileMetadataPointer) - imageBase
	} else {
		imageBase := pe.NtHeader.OptionalHeader.(ImageOptionalHeader64).ImageBase
		rva = uint32(VolatileMetadataPointer - imageBase)
	}

	offset := pe.GetOffsetFromRva(rva)
	imgVolatileMetaSize := uint32(binary.Size(imgVolatileMeta))
	err := pe.structUnpack(&imgVolatileMeta, offset, imgVolatileMetaSize)
	if err != nil {
		return nil
	}
	volatileMeta.Struct = imgVolatileMeta

	if imgVolatileMeta.VolatileAccessTable != 0 &&
		imgVolatileMeta.VolatileAccessTableSize != 0 {
		offset := pe.GetOffsetFromRva(imgVolatileMeta.VolatileAccessTable)
		for i := uint32(0); i < imgVolatileMeta.VolatileAccessTableSize/4; i++ {
			accessRVA, err := pe.ReadUint32(offset)
			if err != nil {
				break
			}

			volatileMeta.AccessRVATable = append(volatileMeta.AccessRVATable, accessRVA)
			offset += 4
		}
	}

	if imgVolatileMeta.VolatileInfoRangeTable != 0 && imgVolatileMeta.VolatileAccessTableSize != 0 {
		offset := pe.GetOffsetFromRva(imgVolatileMeta.VolatileInfoRangeTable)
		rangeEntrySize := uint32(binary.Size(RangeTableEntry{}))
		for i := uint32(0); i < imgVolatileMeta.VolatileAccessTableSize/rangeEntrySize; i++ {
			entry := RangeTableEntry{}
			err := pe.structUnpack(&entry, offset, rangeEntrySize)
			if err != nil {
				break
			}

			volatileMeta.InfoRangeTable = append(volatileMeta.InfoRangeTable, entry)
			offset += rangeEntrySize
		}
	}

	return &volatileMeta
}
