# WhatsApp Status Video Splitter

A simple and efficient Go application that splits MP4 videos into 30-second segments, perfect for uploading to WhatsApp Status. This tool automatically divides your long videos into smaller parts that comply with WhatsApp's status video requirements.

## Features

- Ìæ¨ **Automatic Video Splitting**: Splits videos into 30-second segments
- Ì≥± **WhatsApp Optimized**: Perfect for WhatsApp Status uploads
- ‚ö° **Fast Processing**: Uses FFmpeg's copy codec for no-re-encoding speed
- Ì∂•Ô∏è **Cross-Platform**: Works on Windows, macOS, and Linux
- Ì≥Å **Organized Output**: Creates a dedicated folder for all video segments
- ÌæØ **User-Friendly**: Simple drag-and-drop interface

## Prerequisites

Before using this application, you need to install:

### 1. Go (Go 1.25.1 or later)
- **Windows**: Download from [golang.org](https://golang.org/dl/)
- **macOS**: Use Homebrew: `brew install go`
- **Linux**: Use your package manager or download from [golang.org](https://golang.org/dl/)

### 2. FFmpeg
FFmpeg is required for video processing. Install it and ensure it's available in your system PATH.

#### Windows:
- Download from [ffmpeg.org](https://ffmpeg.org/download.html)
- Extract and add the `bin` folder to your system PATH
- Or use Chocolatey: `choco install ffmpeg`

#### macOS:
```bash
brew install ffmpeg
```

#### Linux (Ubuntu/Debian):
```bash
sudo apt update
sudo apt install ffmpeg
```

#### Linux (CentOS/RHEL):
```bash
sudo yum install ffmpeg
# or for newer versions
sudo dnf install ffmpeg
```

## Installation

### Option 1: Download Pre-built Binary
1. Download the `video-splitter.exe` file from the releases
2. Place it in your desired directory
3. Run it directly (no installation required)

### Option 2: Build from Source
1. Clone this repository:
   ```bash
   git clone https://github.com/MeYo0o/whatsapp_video_splitter.git
   cd whatsapp_video_splitter
   ```

2. Build the application:
   ```bash
   go build -o video-splitter video-splitter.go
   ```

3. Run the executable:
   - **Windows**: `video-splitter.exe`
   - **macOS/Linux**: `./video-splitter`

## How to Use

1. **Run the Application**
   - Double-click `video-splitter.exe` (Windows) or run `./video-splitter` (macOS/Linux)

2. **Select Your Video**
   - When prompted, drag and drop your MP4 file into the terminal/command prompt
   - Press Enter to confirm

3. **Wait for Processing**
   - The application will:
     - Validate your file
     - Create an output directory
     - Calculate the number of segments needed
     - Split the video into 30-second parts

4. **Find Your Segments**
   - All video segments will be saved in a folder named `[your-video-name]_parts`
   - Each segment is named `part1.mp4`, `part2.mp4`, etc.

5. **Upload to WhatsApp**
   - Upload the individual segments to your WhatsApp Status
   - Upload them in order for the best viewing experience

## Supported Operating Systems

- ‚úÖ **Windows** (Windows 10/11)
- ‚úÖ **macOS** (macOS 10.15 or later)
- ‚úÖ **Linux** (Ubuntu, Debian, CentOS, RHEL, and other distributions)

## File Requirements

- **Input Format**: MP4 files only
- **Output Format**: MP4 files (30-second segments)
- **Codec**: Uses copy codec for fast processing without quality loss

## Example Usage

```
=== WhatsApp Status Video Splitter ===
Drag and drop your MP4 file here, then press Enter:
C:\Users\YourName\Videos\my_video.mp4

Created output directory: C:\Users\YourName\Videos\my_video_parts
Video duration: 125.50 seconds
Splitting into 5 segments...
Creating segment 1/5 (30.00 seconds)...
Creating segment 2/5 (30.00 seconds)...
Creating segment 3/5 (30.00 seconds)...
Creating segment 4/5 (30.00 seconds)...
Creating segment 5/5 (5.50 seconds)...

Processing complete!
Your video segments are in: C:\Users\YourName\Videos\my_video_parts
You can now upload these to WhatsApp Status
```

## Troubleshooting

### Common Issues

1. **"ffmpeg not found" error**
   - Make sure FFmpeg is installed and added to your system PATH
   - Restart your terminal/command prompt after installing FFmpeg

2. **"only MP4 files are supported" error**
   - Convert your video to MP4 format before using this tool
   - Use online converters or FFmpeg: `ffmpeg -i input.mov output.mp4`

3. **Permission errors**
   - Make sure you have write permissions in the video's directory
   - Try running as administrator (Windows) or with sudo (Linux/macOS)

## Technical Details

- **Language**: Go 1.25.1
- **Dependencies**: None (uses only standard library)
- **Video Processing**: FFmpeg
- **Segment Duration**: 30 seconds (hardcoded for WhatsApp optimization)
- **Processing Method**: Copy codec (no re-encoding for speed)

## Contributing

Feel free to submit issues, feature requests, or pull requests to improve this tool.

## License

This project is open source. Feel free to use, modify, and distribute as needed.

---

**Happy video splitting! Ìæ¨Ì≥±**
