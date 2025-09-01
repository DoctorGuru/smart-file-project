# Smart File Organizer

Smart File Organizer is a cross-platform tool written in **Go** that automatically organizes files into folders based on rules you define in a config file.  

It can run once, continuously watch a folder, or simulate actions in **dry-run** mode.

---

## Features
- 📂 Automatically sort files (images, docs, videos, etc.)
- ⚙️ Fully customizable rules via `config.yaml`
- 🔄 Continuous folder monitoring with [fsnotify](https://github.com/fsnotify/fsnotify)
- 📝 Logging support
- ✅ CLI options: `run`, `once`, `dry`

---

## Project Structure
smart-file-organizer/
│
├── cmd/organizer/ # Main entry point
├── internal/ # Core logic
│ ├── app/ # Application runner
│ ├── service/ # Rules engine
│ ├── repository/ # File system operations
│ └── transport/ # CLI
├── pkg/ # Reusable packages
│ ├── config/ # Config loader
│ └── logger/ # Logger setup
├── configs/ # Configuration files
│ └── config.yaml
├── test/ # Unit tests
└── README.md # Documentation

---

## Installation

Clone the repo:

```bash
git clone https://github.com/DoctorGuru/smart-file-organizer.git
cd smart-file-organizer
