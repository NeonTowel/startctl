# 🚀 startctl: Your Startup Program Commander

Welcome to **startctl**, the CLI tool that puts you in control of your Windows startup programs. Whether you're adding, removing, enabling, or disabling, startctl has got your back. It's like having a personal assistant for your startup apps, minus the coffee runs.

![Build Status](https://img.shields.io/badge/build-passing-brightgreen) ![License](https://img.shields.io/badge/license-Unlicense-blue)
[![Go Version](https://img.shields.io/badge/Go-1.23-blue.svg)](https://golang.org/)
[![Windows Support](https://img.shields.io/badge/Windows-✔-brightgreen)](https://www.microsoft.com/)
[![GitHub release](https://img.shields.io/github/v/release/neontowel/startctl)](https://github.com/neontowel/startctl/releases)
[![Issues](https://img.shields.io/github/issues/neontowel/startctl)](https://github.com/neontowel/startctl/issues)
[![Stars](https://img.shields.io/github/stars/neontowel/startctl?style=social)](https://github.com/neontowel/startctl/stargazers)


## 📥 Installation

1. Ensure you have `go-task` installed. You can download it from [go-task's GitHub](https://github.com/go-task/task) and follow the installation instructions for your platform.
2. Clone the repo: `git clone https://github.com/neontowel/startctl.git`
3. Navigate to the directory: `cd startctl`

## 🛠️ Usage

Here's how you can use startctl to manage your startup programs:

- **Build the application**: `task build`
- **Run the application in development mode**: `task run`

- **Add a program**: `startctl add <name> <path>`
- **Remove a program**: `startctl remove <name>`
- **Enable a program**: `startctl enable <name>`
- **Disable a program**: `startctl disable <name>`
- **List all programs**: `startctl list`

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

*In a world where startup programs run amok, one CLI tool stands to bring order...* 