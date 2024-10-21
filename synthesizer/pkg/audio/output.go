// pkg/audio/output.go

package audio

import (
	"synthesizer/pkg/util"

	"github.com/gordonklaus/portaudio"
)

// Stream is global audio stream

var stream *portaudio.Stream

//Init te PortAudio lib

func Init() error {
	return portaudio.Initialize()
}

// Start opens stream and starts audio playback
func Start(generateAudio func([]float32)) error {
	stream, err := portaudio.OpenDefaultStream(0, 1, util.SampleRate, len(make([]float32, util.BufferSize)), func(out []float32) {
		// Call synthesizer's Generate Audio function
		generateAudio(out)
	})
	if err != nil {
		return err
	}

	// Start the stream
	err = stream.Start()
	if err != nil {
		return err
	}

	//Keep the stream running until the user closes the program
	select {} // Blocks forever

	// In a real appplication (AKA NYI) there should be a way to stop the stream gracefully
	// For example, a signal handler to stop the stream and close the audio device
	// For now, we will just block forever
}

// Close terminates PortAudio library

func Close() {
	portaudio.Terminate()
}
