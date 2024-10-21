// pkg//synth/waveform.go

package synth

import "math"

//WaveformType represents the type of waveform

type WaveformType int

const (
	SineWave WaveformType = iota
	SquareWave
	TriangleWave
	SawToothWave
)

// GenerateWaveForm generates a waveform sample based on type, phase
// Given the phase: 0.0 <= phase < 1.0, the following formulas are used:
// SineWave: sin(2 * Pi * phase)
// SquareWave: 1 if phase < 0.5, -1 otherwise
// TriangleWave: 4 * abs(phase - 0.5) - 1
// SawToothWave: 2 * (phase - floor(phase + 0.5))
// For resources on waveform generation:
// https://en.wikipedia.org/wiki/Waveform
// https://thewolfsound.com/sine-saw-square-triangle-pulse-basic-waveforms-in-synthesis/




func GenerateWaveForm(waveformType WaveformType, phase float64) float64 {
	switch waveformType {
	case SineWave:
		return math.Sin(2 * math.Pi * phase)
	case SquareWave:
		if phase < 0.5 {
			return 1
		}
		return -1
	case TriangleWave:
		return 4*math.Abs(phase-0.5) - 1.0
	case SawToothWave:
		return 2 * (phase - math.Floor(phase+0.5))
	default:
		return 0.0
	}
}
