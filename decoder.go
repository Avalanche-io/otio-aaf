// SPDX-License-Identifier: Apache-2.0
// Copyright Contributors to the OpenTimelineIO project

package aaf

import (
	"io"

	"github.com/Avalanche-io/gotio/opentimelineio"
)

// Decoder reads AAF (Advanced Authoring Format) files and converts them to OpenTimelineIO objects.
//
// AAF is a complex binary format that requires extensive parsing of a structured storage
// container, object model deserialization, and mapping to OTIO concepts. This decoder
// currently returns ErrNotImplemented and directs users to the Python bridge for AAF support.
//
// Example usage (when implemented):
//
//	f, err := os.Open("project.aaf")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer f.Close()
//
//	decoder := aaf.NewDecoder(f)
//	timeline, err := decoder.Decode()
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// For production use, please use the Python bridge which provides full AAF support
// via the otio-aaf-adapter and PyAAF2 library.
type Decoder struct {
	r io.Reader
}

// NewDecoder creates a new AAF decoder that reads from r.
//
// The decoder is currently not implemented. Use the Python bridge for AAF support:
//
//	import "github.com/Avalanche-io/gotio/adapters"
//	timeline, err := adapters.ReadFromFile("project.aaf")
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

// Decode reads an AAF file and returns an OpenTimelineIO Timeline.
//
// This method currently returns ErrNotImplemented. AAF is a complex binary format
// that requires:
//
//   - Parsing the Structured Storage container format (similar to a filesystem)
//   - Deserializing the AAF object model (Mobs, MobSlots, SourceClips, etc.)
//   - Resolving object references and relationships
//   - Mapping AAF entities to OpenTimelineIO objects
//   - Handling multiple AAF versions and vendor extensions
//
// For full AAF support, use the Python bridge which leverages the mature
// otio-aaf-adapter implementation:
//
//	import "github.com/Avalanche-io/gotio/adapters"
//	timeline, err := adapters.ReadFromFile("project.aaf")
//
// Returns:
//   - ErrNotImplemented: Always returned in current implementation
//
// Future native implementation would need to:
//   - Parse AAF binary container (Structured Storage format)
//   - Extract and deserialize object model
//   - Map to OTIO Timeline, Track, Clip, etc.
//   - Preserve metadata, markers, and effects
func (d *Decoder) Decode() (*opentimelineio.Timeline, error) {
	return nil, ErrNotImplemented
}

// DecodeWithOptions reads an AAF file with additional options and returns an OpenTimelineIO object.
//
// Options (when implemented) may include:
//   - "simplified": bool - Simplify nested compositions
//   - "attach_markers": bool - Attach markers to clips
//   - "transcribe_linear_speed_effects": bool - Convert linear speed effects
//   - "rate": float64 - Override timeline rate
//
// This method currently returns ErrNotImplemented. Use the Python bridge for full support.
func (d *Decoder) DecodeWithOptions(options map[string]interface{}) (opentimelineio.SerializableObject, error) {
	return nil, ErrNotImplemented
}
