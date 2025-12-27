# otio-aaf: AAF Adapter for OpenTimelineIO (Go)

AAF (Advanced Authoring Format) adapter for OpenTimelineIO in Go.

## Status: Python Bridge Recommended

This is a stub implementation that directs users to the Python bridge for full AAF support. AAF is an extremely complex binary format, and the Python ecosystem has mature libraries for handling it.

## What is AAF?

Advanced Authoring Format (AAF) is a professional multimedia file format designed for the interchange of digital media and metadata between platforms, systems, and applications. It is:

- Standardized by the Advanced Media Workflow Association (AMWA)
- Widely used in professional video editing and post-production
- Based on a Structured Storage container format (similar to a filesystem)
- Contains a complex object model with MobSlots, SourceClips, Sequences, and other entities

## Why Use the Python Bridge?

AAF support requires:

1. **Binary Container Parsing**: AAF uses Microsoft's Structured Storage format (compound document)
2. **Complex Object Model**: Extensive relationships between Mobs, MobSlots, SourceClips, Sequences, etc.
3. **Multiple Versions**: Support for various AAF specification versions
4. **Vendor Extensions**: Handling application-specific extensions from different NLEs

The Python ecosystem has:
- **PyAAF2**: A mature, well-tested library for AAF manipulation
- **otio-aaf-adapter**: A comprehensive adapter leveraging PyAAF2
- Years of production use and testing

## Using AAF with the Python Bridge

### Installation

Ensure you have Python with OpenTimelineIO and the AAF adapter installed:

```bash
pip install opentimelineio
pip install otio-aaf-adapter
```

### Go Usage

```go
import (
    "log"
    "github.com/mrjoshuak/gotio/adapters"
    _ "github.com/mrjoshuak/gotio/adapters" // Registers Python bridge
)

func main() {
    // Read AAF file (uses Python bridge automatically)
    timeline, err := adapters.ReadFromFile("project.aaf")
    if err != nil {
        log.Fatal(err)
    }

    // Work with timeline in Go
    for _, track := range timeline.Tracks.Children() {
        // Process tracks
    }

    // Write back to AAF (uses Python bridge)
    err = adapters.WriteToFile(timeline, "output.aaf")
    if err != nil {
        log.Fatal(err)
    }
}
```

## What This Package Provides

This package provides:

- Type definitions for potential future native implementation
- Clear documentation about AAF format complexity
- Stub implementations that return `ErrNotImplemented`
- API surface compatible with other gotio adapters

## Current Implementation

The current Go implementation includes:

### aaf.go
- Package documentation
- Error types (`ErrNotImplemented`, `DecodeError`, `EncodeError`)
- Format complexity documentation

### decoder.go
- `Decoder` type with `NewDecoder(io.Reader)` constructor
- `Decode()` method (returns `ErrNotImplemented`)
- Documentation of what full implementation would require

### encoder.go
- `Encoder` type with `NewEncoder(io.Writer)` constructor
- `Encode(*Timeline)` method (returns `ErrNotImplemented`)
- Documentation of encoding requirements

## Future Native Implementation

A native Go implementation would require:

1. **Structured Storage Parser**
   - Parse Microsoft compound document format
   - Navigate internal directory structure
   - Extract object streams

2. **AAF Object Model**
   - Implement AAF entities (Mob, MobSlot, SourceClip, etc.)
   - Handle object references and relationships
   - Support metadata and properties

3. **OTIO Mapping**
   - Map AAF entities to OpenTimelineIO objects
   - Preserve metadata, markers, and effects
   - Handle track types and compositions

4. **Version Support**
   - Support multiple AAF specification versions
   - Handle vendor-specific extensions
   - Ensure compatibility with major NLEs

This would be a significant undertaking (several thousand lines of code) and would need extensive testing against real-world AAF files from various applications.

## Supported Features (Python Bridge)

When using the Python bridge, the following AAF features are supported:

- Single and multiple video tracks
- Audio tracks and clips
- Gap/filler elements
- Markers and nested compositions
- Transitions (dissolves, wipes)
- Linear speed effects
- Metadata preservation

Not currently supported:
- Audio/video effects (beyond speed)
- Advanced speed ramping
- Color Decision Lists (CDL)
- Image sequences

## References

- [AAF Specification](https://www.amwa.tv/projects/MS-04.shtml) - Official AMWA specification
- [PyAAF2](https://github.com/markreidvfx/pyaaf2) - Python AAF library
- [otio-aaf-adapter](https://github.com/OpenTimelineIO/otio-aaf-adapter) - Python OTIO adapter
- [OpenTimelineIO](https://github.com/OpenTimelineIO/OpenTimelineIO) - Main OTIO project

## Development

This package uses a local replace directive for development:

```bash
cd /Users/joshua/ws/active/gotio/otio-aaf
go build
```

The `go.mod` includes:
```
replace github.com/mrjoshuak/gotio => ../gotio
```

## Contributing

If you're interested in implementing native AAF support in Go:

1. Start with the AAF specification and PyAAF2 source code
2. Implement the Structured Storage parser
3. Build the AAF object model
4. Create mapping to OTIO objects
5. Add comprehensive tests using real AAF files

This would be a valuable contribution to the Go OTIO ecosystem!

## License

Apache-2.0 - See LICENSE file for details.

Copyright Contributors to the OpenTimelineIO project.
