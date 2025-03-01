# 🚀 startctl: Your Startup Program Commander

Welcome to **startctl**, the CLI tool that puts you in control of your Windows startup programs. Whether you're adding, removing, enabling, or disabling, startctl has got your back. It's like having a personal assistant for your startup apps, minus the coffee runs. ☕️

![Build Status](https://img.shields.io/badge/build-passing-brightgreen) ![License](https://img.shields.io/badge/license-Unlicense-blue)
[![Go Version](https://img.shields.io/badge/Go-1.23-blue.svg)](https://golang.org/)
[![Windows Support](https://img.shields.io/badge/Windows-✔-brightgreen)](https://www.microsoft.com/)
[![GitHub release](https://img.shields.io/github/v/release/neontowel/startctl)](https://github.com/neontowel/startctl/releases)
[![Issues](https://img.shields.io/github/issues/neontowel/startctl)](https://github.com/neontowel/startctl/issues)
[![Stars](https://img.shields.io/github/stars/neontowel/startctl?style=social)](https://github.com/neontowel/startctl/stargazers)
![Scoop Support](https://img.shields.io/badge/Scoop-Supported-brightgreen)
[![Go Report Card](https://goreportcard.com/badge/github.com/neontowel/startctl)](https://goreportcard.com/report/github.com/neontowel/startctl)

## 💻 Development

1. **Prerequisites**: Ensure you have Go 1.23 or later installed. You can download it from [the official Go website](https://golang.org/dl/). Additionally, make sure `go-task` is installed. You can download it from [go-task's GitHub](https://github.com/go-task/task) and follow the installation instructions for your platform.
2. **Clone the repo**: 
   ```bash
   git clone https://github.com/neontowel/startctl.git
   ```
3. **Navigate to the directory**: 
   ```bash
   cd startctl
   ```

## 🛠️ Build and run

Here's how you can use startctl to manage your startup programs:

- **Build the application**: 
  ```bash
  task build
  ```
- **Run the application in development mode**: 
  ```bash
  task run
  ```

## 📥 Installation

To install startctl using Scoop, follow these steps:

1. **Add the Scoop bucket**: 
   ```bash
   scoop bucket add neontowel https://github.com/NeonTowel/scoop-bucket
   ```

2. **Install startctl**:
   ```bash
   scoop install neontowel/startctl
   ```

## 🌟 Features

- **Add Programs**: Easily add any program to your startup list.
- **Remove Programs**: Clean up your startup by removing unnecessary programs.
- **Enable/Disable Programs**: Toggle programs on or off without deleting them.
- **List Programs**: Get a quick overview of all your startup programs.

## 🤝 Contributing

Contributions are welcome! Feel free to fork the repo and submit a pull request. Let's make startctl even better together.

## 📜 License

This project is licensed under the Unlicense. Feel free to use, modify, and distribute as you wish. No strings attached!

---

*In a world where startup programs run amok, one CLI tool stands to bring order...* 🕶️ 