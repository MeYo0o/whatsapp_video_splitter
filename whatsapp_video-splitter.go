package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("=== WhatsApp Status Video Splitter ===")
	fmt.Println("Drag and drop your MP4 file here, then press Enter:")

	// Get file path from user
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSpace(filePath)
	filePath = strings.Trim(filePath, `"`) // Remove quotes from drag & drop

	// Validate input file
	if err := validateFile(filePath); err != nil {
		log.Fatalf("Validation error: %v", err)
	}

	// Create output directory
	outputDir := createOutputDir(filePath)
	fmt.Printf("Created output directory: %s\n", outputDir)

	// Get video duration
	duration, err := getVideoDuration(filePath)
	if err != nil {
		log.Fatalf("Error getting video duration: %v", err)
	}
	fmt.Printf("Video duration: %.2f seconds\n", duration)

	// Calculate segments
	segmentCount := int(duration/30) + 1
	fmt.Printf("Splitting into %d segments...\n", segmentCount)

	// Process segments
	for i := 0; i < segmentCount; i++ {
		start := i * 30
		end := start + 30
		if end > int(duration) {
			end = int(duration)
		}

		outputFile := filepath.Join(outputDir, fmt.Sprintf("part%d.mp4", i+1))
		fmt.Printf("Creating segment %d/%d (%.2f seconds)...\n", i+1, segmentCount, float64(end-start))

		if err := createVideoSegment(filePath, outputFile, start, end); err != nil {
			log.Printf("Error creating segment %d: %v", i+1, err)
		}
	}

	fmt.Println("\nProcessing complete!")
	fmt.Printf("Your video segments are in: %s\n", outputDir)
	fmt.Println("You can now upload these to WhatsApp Status")
}

func validateFile(path string) error {
	// Check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(path))
	if ext != ".mp4" {
		return fmt.Errorf("only MP4 files are supported")
	}

	// Check FFmpeg installation
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		return fmt.Errorf("ffmpeg not found. Please install ffmpeg and add it to your PATH")
	}

	return nil
}

func createOutputDir(inputPath string) string {
	base := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath))
	outputDir := filepath.Join(filepath.Dir(inputPath), base+"_parts")

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	return outputDir
}

func getVideoDuration(path string) (float64, error) {
	cmd := exec.Command(
		"ffprobe",
		"-v", "error",
		"-show_entries", "format=duration",
		"-of", "default=noprint_wrappers=1:nokey=1",
		path,
	)

	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("ffprobe error: %v", err)
	}

	duration, err := strconv.ParseFloat(strings.TrimSpace(string(output)), 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse duration: %v", err)
	}

	return duration, nil
}

func createVideoSegment(input, output string, start, end int) error {
	duration := end - start

	cmd := exec.Command(
		"ffmpeg",
		"-y", // Overwrite existing files
		"-ss", fmt.Sprintf("%d", start), // Start time
		"-i", input,                    // Input file
		"-t", fmt.Sprintf("%d", duration), // Duration
		"-c", "copy",                   // Copy codec (no re-encoding)
		"-avoid_negative_ts", "make_zero",
		"-loglevel", "error", // Only show errors
		output,
	)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg execution failed: %v", err)
	}

	return nil
}