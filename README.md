# Golang Additive Synthesizer w/GUI using Fyne

This is an initial commit so this is a work in progress we don't need 
This project is a basic additive synthesizer built using Golang with a graphical user interface (GUI). The synthesizer features:

- Two oscillators with selectable waveforms (sine, square, triangle, sawtooth)
- Three filters (e.g., low-pass, high-pass, bandpass)
- Two Low-Frequency Oscillators (LFOs) for modulation
- MIDI input support for real-time control
- A simple GUI for user interaction (using Fyne) 

## Project Hierarchy

```
synthesizer/
├── cmd/
│   └── synth/
│       └── main.go
├── pkg/
│   ├── audio/
│   │   └── output.go
│   ├── midi/
│   │   └── input.go
│   ├── synth/
│   │   ├── engine.go         // Added engine.go
│   │   ├── oscillator.go
│   │   ├── waveform.go
│   │   ├── filter.go
│   │   ├── lfo.go
│   │   ├── voice.go
│   │   └── envelope.go
│   └── util/
│       ├── constants.go
│       └── config.go
├── go.mod
├── go.sum
└── README.md
```

## Dependencies



### Audio Output Library: PortAudio Go Bindings

Provides real-time audio streaming capabilities.

Install via:

```bash
go get github.com/gordonklaus/portaudio
```

### MIDI Input Library: PortMIDI Go Bindings

Allows interaction with MIDI devices.

Install viawith GUI:

```bash
go get github.com/rakyll/portmidi
```

### GUI Library: Fyne

A cross-platform GUI toolkit for Go.

Install via:

```bash
go get fyne.io/fyne/v2
```

### All Dependencies

To install all dependencies, run:

```bash
go get github.com/gordonklaus/portaudio
go get github.com/rakyll/portmidi
go get fyne.io/fyne/v2
```

## Components

### 1. cmd/synth/main.go

Initializes the synthesizer components, sets up MIDI input, audio output, and the GUI, starts main event loop.

### 2. pkg/gui/app.go

Manages the GUI application lifecycle, including initializing the GUI window and components, running the main GUI event loop, and interfacing MIDI input with engine for real-time control

### 3. pkg/gui/components.go

Defines the GUI components (buttons, sliders, etc.) and binds the GUI controls to synthesizer parameters.

### 4. pkg/gui/handlers.go

Contains event handlers for GUI interactions. Updates synthesizer parameters based on user input.

### 5. pkg/audio/output.go

Handles audio output streaming via PortAudio.

### 6. pkg/midi/input.go

Handles MIDI input from external devices using PortMIDI.

### 7. pkg/synth/oscillator.go

Defines oscillator components and methods to update parameters like frequency and waveform in real-time.

### 8. pkg/synth/engine.go

The central synthesizer engine manages voices, global parameters, and updates parameters based on input from the GUI or MIDI.

### 9. pkg/synth/voice.go

Manages individual synthesizer voices, including the signal chain from oscillators, filters, and envelopes.

### 10. pkg/util/constants.go

Stores constant values used throughout the application.

### 11. pkg/util/config.go

Holds configuration settings, including any GUI-related settings.


## License

This project is licensed under the GPL3 License.

