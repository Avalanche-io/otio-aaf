// SPDX-License-Identifier: Apache-2.0
// Copyright Contributors to the OpenTimelineIO project

// Package aaf provides support for Advanced Authoring Format (AAF) files.
//
// AAF (Advanced Authoring Format) is a professional multimedia file format
// designed for the interchange of digital media and metadata between platforms,
// systems, and applications. It is standardized by the Advanced Media Workflow
// Association (AMWA) and widely used in professional video editing and
// post-production workflows.
//
// # Format Complexity
//
// AAF is an extremely complex binary format with several challenging characteristics:
//
//   - Binary container format similar to a filesystem with internal directory structure
//   - Complex object model with MobSlots, SourceClips, Sequences, and other entities
//   - Extensive use of cross-references and object relationships
//   - Support for multiple versions and vendor-specific extensions
//   - Rich metadata capabilities including timecode, markers, and effects
//
// # Current Implementation Status
//
// Due to the complexity of the AAF format, this Go implementation currently provides
// only stub implementations that return ErrNotImplemented. Full AAF support requires
// extensive binary parsing and object model implementation.
//
// # Recommended Approach
//
// For production use, we recommend using the Python bridge to access the mature
// otio-aaf-adapter implementation, which leverages the PyAAF2 library for
// comprehensive AAF support. The Python bridge provides:
//
//   - Full read/write support for AAF files
//   - Support for multiple video and audio tracks
//   - Markers, transitions, and nested compositions
//   - Linear speed effects and metadata preservation
//
// To use the Python bridge:
//
//	import (
//	    "github.com/Avalanche-io/gotio/adapters"
//	    _ "github.com/Avalanche-io/gotio/adapters" // registers Python bridge
//	)
//
//	// Python bridge will automatically handle .aaf files
//	timeline, err := adapters.ReadFromFile("project.aaf")
//
// # Future Native Implementation
//
// A native Go implementation of AAF support may be developed in the future.
// This would require:
//
//   - Binary parser for the AAF container format (based on Structured Storage)
//   - Object model implementation for AAF entities
//   - Mapping between AAF objects and OpenTimelineIO objects
//   - Support for various AAF versions and extensions
//
// # References
//
//   - AAF Specification: https://www.amwa.tv/projects/MS-04.shtml
//   - PyAAF2 Library: https://github.com/markreidvfx/pyaaf2
//   - OTIO AAF Adapter: https://github.com/OpenTimelineIO/otio-aaf-adapter
package aaf

import (
	"errors"
	"fmt"
)

// ErrNotImplemented is returned when AAF functionality is not yet implemented
// in the native Go adapter. Use the Python bridge for AAF support.
var ErrNotImplemented = errors.New("aaf: native Go implementation not available, use Python bridge for AAF support")

// DecodeError represents an error that occurred during AAF decoding.
type DecodeError struct {
	Message string
}

func (e *DecodeError) Error() string {
	return fmt.Sprintf("aaf decode error: %s", e.Message)
}

// EncodeError represents an error that occurred during AAF encoding.
type EncodeError struct {
	Message string
}

func (e *EncodeError) Error() string {
	return fmt.Sprintf("aaf encode error: %s", e.Message)
}

// UnsupportedVersionError indicates an unsupported AAF version.
type UnsupportedVersionError struct {
	Version string
}

func (e *UnsupportedVersionError) Error() string {
	return fmt.Sprintf("aaf: unsupported version %s", e.Version)
}
