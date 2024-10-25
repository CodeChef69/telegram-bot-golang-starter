# Go Telegram Bot Template

**Go Telegram Bot Template** is a lightweight boilerplate for building Telegram bots using Go. This template provides a structured foundation, enabling developers to create, deploy, and manage Telegram bots efficiently.

## Features

- **Modular Architecture**: Organized folder structure for easy navigation and maintenance.
- **Config Management**: Loads configurations from a `.env` file for better security and flexibility.
- **Basic Handlers**: Predefined command handlers to get you started quickly.
- **Easy Setup**: Get your bot up and running with minimal setup required.

## Getting Started

### Prerequisites

- Go 1.16 or higher
- A Telegram account
- Your Telegram Bot Token (create a new bot using [BotFather](https://core.telegram.org/bots#botfather))

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/CodeChef69/telegram-bot-golang-starter
2. Navigate to the project directory
   ```bash
   cd telegram-bot-golang-starter
3. Create a ```.env``` file with your Telegram Bot token:
   ```bash
   TELEGRAM_BOT_TOKEN=your-telegram-bot-token
4. Install dependencies:
   ```bash
   go mod tidy
5. Install dependencies:
   ```bash
   go run cmd/bot/main.go