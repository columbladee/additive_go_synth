// cmd/synth/main.go

package main

import (
	"fmt"
	"synthesizer/pkg/audio"
	"synthesizer/pkg/midi"
	"synthesizer/pkg/synth"
)

func main() {
	// Initialize Audio system

	err := audio.Init()

	if err != nil {
		fmt.Println("Error initalizing audio:", err)
		return
	}
	defer audio.Close() // Ensure audio system is closed on exit

	// Initialize MIDI system
	err = midi.Init()

	if err != nil {
		fmt.Println("Error initializing MIDI:", err)
		return
	}
	defer midi.Close() // Ensure MIDI is closed on exit

	// Create synth engine instance

	engine := synth.NewEngine()
	
	// Start Listening for MIDI input in a separate goroutine
	go midi.Listen(engine.HandleMIDI)

	// Start audio output, passing in engine's audio generation function

	err = audio.Start(engine.GenerateAudio)
	if err != nil {
		fmt.Println("Error starting audio output:", err)
		return
	}

	// Keep main function running

	select {} // Blocks forever
}
