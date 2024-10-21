// pkg/synth/engine.go
// 2024-10-20 : Need to finish HandleMIDI function after implementing MIDIEvent struct in pkg/midi/input.go and NoteOn, NoteOff constants in pkg/midi/constants.go

package synth

import (
	"math"
	"sync"
	"synthesizer/pkg/util"
)
// Engine struct manages global state

type Engine struct {
	Oscillators [2]*Oscillator // Two oscillators
	mutex	   sync.RWMutex      
}

// NewEngine creates a new Engine instance

func NewEngine() *Engine {
	return &Engine {
		Oscillators: [2]*Oscillator {
			NewOscillator(SineWave, 440), // Oscillator 1 Defaults to A4
			NewOscillator(SquareWave, 440), // Oscillator 2 Defaults to A4
		},
	}
}

// GenerateAudio generates audio samples for the audio output callback
func (e *Engine) GenerateAudio(buffer []float32) {
	e.mutex.RLock()	// Lock for reading (shared data)
	defer e.mutex.RUnlock() // Unlock when done

	// Iterate over buffer and fill it with audio samples
	for i := range buffer {
		sample := 0.0

		// Sum the outputs of both oscillators
		for _, osc := range e.Oscillators {
			sample += osc.GenerateSample()
		}

		// Normalize Sample to prevent clipping -
		buffer[i] = float32(sample * 0.5)
		// Note: Linear normalization is a simple way to prevent clipping, not ideal
		// Other normalization techniques can be used for better results such as: RMS normalization
		// Normalization may not be the best method to prevent clipping in all cases
		// Other solutions might be preferred such as compressing, limiting, etc.
		// For our purposes this is a very simple additive synth so linear normalization is fine.
	}
}

// HandleMIDI - process incoming MIDI events
// All of this is handled under pkg/midi 
// Also see the portmidi package for more information on MIDI handling

func (e *Engine) HandleMIDI(event MIDIEvent) { 
	e.mutex.Lock() // Lock for writing shared data
	defer e.mutex.Unlock() // Unlock when done

	// Extract MIDI event data
	messageType := event.Status & 0xF0 // Note: for portMIDI 0xF0 is a mask for the status byte

	switch messageType {
	case NoteOn:
		if event.Data2 > 0 { // NoteOn with velocity > 0
			// Update oscillator frequency (both oscillators)
			for _, osc := range e.Oscillators {
				osc.SetFrequency(frequency)
			}
		} else {
			// NoteOn with velocity 0 is equivalent to NoteOff
			// NYI: Implement NoteOff handling (if necessary)
		}
	case NoteOff:
		// NYI: Implement NoteOff handling (if necessary)
	default:
		// NYI: Implement other MIDI message types (if necessary)
	}


}