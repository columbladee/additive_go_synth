// pkg/synth/oscillator.go

package synth

import "synthesizer/pkg/util"

// Oscillator represents a single oscillator

type Oscillator struct {
	Frequency float64      // Frequency in Hz
	Waveform  WaveformType // Waveform type (sine, square, triangle, sawtooth)
	Phase     float64      // Current phase [0.0, 1.0)
	PhaseStep float64      // Phasestep per sample
}

// NewOscillator creates a new Oscillator instance - given waveform and frequency
func NewOscillator(waveform WaveformType, frequency float64) *Oscillator {
	osc := &Oscillator{
		Frequency: frequency,
		Waveform:  waveform,
	}
	osc.UpdatePhaseStep() // later
	return osc
}

// UpdatePhaseStep updates phase step based on frequences
func (o *Oscillator) UpdatePhaseStep() {
	o.PhaseStep = o.Frequency / util.SampleRate // SampleRate is a constant from util = 44100.0
}

// SetFrequency updates oscillator frequency and phase step
func (o *Oscillator) SetFrequency(freq float64) {
	o.Frequency = freq
	o.UpdatePhaseStep()
}

// GenerateSample generates the next audio sample
func (o *Oscillator) GenerateSample() float64 {
	// Get the waveform value at the current phase
	sample := GenerateWaveForm(o.Waveform, o.Phase)

	// Increment phase wrapping around at 1.0
	o.Phase += o.PhaseStep
	if o.Phase >= 1.0 {
		o.Phase -= 1.0
	}
	return sample

}
