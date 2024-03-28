# TGBotPocket Telegram Bot

## Overview
TGBotPocket is a Telegram bot designed to save web bookmarks directly via Telegram. It allows users to quickly and easily save links to their favorite webpages, articles, and online resources for later reference, all within the Telegram interface.

## Getting Started

### Prerequisites
- Docker
- Go (for building the project without Docker)

### Installation

#### Using Docker (Recommended)
1. **Build the Docker Image**

   Use the `make` command to build the Docker image for TGBotPocket:
   ```
   make build-image
   ```
   This command compiles the Docker image with the tag `tgpocket:v0.1`.

2. **Start the Container**

   After building the image, start the Docker container with the following command:
   ```
   make start-container
   ```
   This command runs the Docker container and maps port 80 of the container to port 80 of the host, allowing you to access the bot through `localhost`.

#### Building Manually
If you prefer to build and run the bot without Docker, follow these steps:

1. **Build the Bot**

   Compile the bot executable with the following command:
   ```
   make build
   ```
   This command compiles the Go code and places the resulting binary in the `./.bin` directory.

2. **Run the Bot**

   Start the bot using the compiled binary with:
   ```
   make run
   ```
   This will run the bot executable, allowing it to start processing requests.

## Usage

After starting TGBotPocket (either via Docker or manually), you can interact with it through Telegram to save your web bookmarks. Here's how to get started:

1. **Find the Bot on Telegram**

   Search for the TGBotPocket bot in Telegram using its name.

2. **Saving Bookmarks**

   To save a bookmark, simply send a message to the bot with the URL you wish to save. The bot will confirm that the bookmark has been saved.



