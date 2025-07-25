# Temperature Conversion Application

[![CI](https://github.com/muslchn/temperature-conversion/actions/workflows/ci.yml/badge.svg)](https://github.com/muslchn/temperature-conversion/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/muslchn/temperature-conversion)](https://goreportcard.com/report/github.com/muslchn/temperature-conversion)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A command-line temperature conversion application built in Go that converts Celsius temperatures to Fahrenheit and Réaumur scales, with location-based measurements and temperature classification.

## Features

- 🌡️ **Temperature Conversion**: Convert Celsius to Fahrenheit and Réaumur
- 📍 **Location Tracking**: Record the location where temperature is measured
- 🏷️ **Temperature Classification**: Automatically classify temperatures as:
  - ❄️ **Dingin (Cold)**: Below 18°C
  - 🌤️ **Hangat (Warm)**: 18°C - 25°C
  - 🔥 **Panas (Hot)**: Above 25°C
- ✅ **Input Validation**:
  - Location must be alphabetic characters and spaces only
  - Temperature must be numeric values
- 🏗️ **Clean Architecture**: Uses structs and pointer-based functions for efficient memory usage

## Installation

### Prerequisites

- Go 1.21 or higher
- Git

### Install from Source

```bash
# Clone the repository
git clone https://github.com/muslchn/temperature-conversion.git

# Navigate to project directory
cd temperature-conversion

# Build the application
go build -o temperature-converter .

# Run the application
./temperature-converter
```

### Quick Install with Go

```bash
go install github.com/muslchn/temperature-conversion@latest
```

## Usage

### Interactive Mode

Run the application and follow the prompts:

```bash
./temperature-converter
```

### Example Usage

```text
--- Konverter Suhu ---
Masukkan lokasi pengukuran suhu: Jakarta
Masukkan suhu dalam Celsius: 25

> Suhu di Jakarta adalah hangat
> Suhu di Jakarta dalam Reamur = 20.00
> Suhu di Jakarta dalam Fahrenheit = 77.00
```

### Error Handling

The application provides clear error messages for invalid inputs:

```text
--- Konverter Suhu ---
Masukkan lokasi pengukuran suhu: Madinah
Masukkan suhu dalam Celsius: dua puluh lima

> input tidak valid, hanya menerima angka
```

## Technical Details

### Architecture

The application is built using clean architecture principles:

- **Struct-based Data Management**: Uses `TemperatureData` struct to encapsulate location and temperature information
- **Pointer-based Functions**: Implements pass-by-reference for efficient memory usage
- **Separation of Concerns**: Clear separation between input validation, business logic, and output formatting

### Key Components

```go
type TemperatureData struct {
    Location string
    Celsius  float64
}
```

### Conversion Formulas

- **Fahrenheit**: `F = (C × 9/5) + 32`
- **Réaumur**: `R = C × 4/5`

### Temperature Classification Rules

| Category | Temperature Range (°C) | Description |
|----------|------------------------|-------------|
| Dingin   | < 18                   | Cold        |
| Hangat   | 18 - 25                | Warm        |
| Panas    | > 25                   | Hot         |

## Development

### Project Structure

```text
temperature-conversion/
├── .github/
│   └── workflows/
│       └── ci.yml          # GitHub Actions CI/CD pipeline
├── main.go                 # Main application code
├── go.mod                  # Go module definition
├── temperature-converter   # Compiled binary
└── README.md              # This file
```

### Building from Source

```bash
# Download dependencies
go mod download

# Build the application
go build -o temperature-converter .

# Run tests (when available)
go test ./...

# Run linting
go vet ./...
```

### Code Quality

This project maintains high code quality standards:

- **Static Analysis**: Uses `staticcheck` for advanced Go linting
- **Code Formatting**: Follows `gofmt` standards
- **Error Handling**: Implements proper error handling patterns
- **Input Validation**: Comprehensive input validation for user safety

## CI/CD

The project uses GitHub Actions for continuous integration:

- **Automated Testing**: Runs on every push and pull request
- **Multiple Go Versions**: Tested against Go 1.21+
- **Static Analysis**: Includes `go vet` and `staticcheck`
- **Build Verification**: Ensures clean builds across platforms

## Contributing

We welcome contributions! Please follow these steps:

1. **Fork** the repository
2. **Create** a feature branch (`git checkout -b feature/amazing-feature`)
3. **Commit** your changes (`git commit -m 'Add some amazing feature'`)
4. **Push** to the branch (`git push origin feature/amazing-feature`)
5. **Open** a Pull Request

### Development Guidelines

- Follow Go naming conventions
- Add tests for new features
- Ensure all CI checks pass
- Update documentation as needed

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built as part of learning Go programming fundamentals
- Focuses on practical application of structs, pointers, and input validation
- Implements clean architecture patterns for maintainable code

## Support

If you encounter any issues or have questions:

1. Check the [Issues](https://github.com/muslchn/temperature-conversion/issues) page
2. Create a new issue with detailed information
3. Follow the issue template for faster resolution

---

Made with ❤️ and Go
