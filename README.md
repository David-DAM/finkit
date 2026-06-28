# FinKit

FinKit is a powerful and lightweight Command Line Interface (CLI) tool built in Go for financial operations.

## Features

- **Real-time Currency Conversion**: Convert amounts between dozens of world currencies.
- **Efficient Caching**: Uses a file-based cache to store exchange rates and supported currencies, reducing API calls and improving response times.
- **Interactive Autocompletion**: Supports shell completion for currency codes.
- **Modern Go Architecture**: Built with Go 1.25, using clean architecture principles and the Cobra CLI framework.

## Installation

### Prerequisites

- [Go 1.25](https://go.dev/dl/) or higher.

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/DAVID-DAM/finkit.git
   cd finkit
   ```

2. Build the application:
   ```bash
   go build -o finkit main.go
   ```

3. (Optional) Move to your bin directory:
   ```bash
   mv finkit /usr/local/bin/
   ```

## Usage

The main command for currency operations is `currency`.

### Convert Currencies

Use the `convert` subcommand followed by the amount, the source currency, and the target currency.

```bash
finkit currency convert 100 EUR USD
```

**Examples:**
```bash
# Convert 100 Euros to US Dollars
finkit currency convert 100 EUR USD

# Convert 50 British Pounds to Japanese Yen
finkit currency convert 50 GBP JPY

# Enable verbose logging to see cache hits/misses
finkit currency convert 1000 USD CHF -v
```

### Help

You can always use the `--help` flag to see available commands and options.

```bash
finkit --help
finkit currency convert --help
```

## Architecture

FinKit follows a modular architecture:
- **`cmd/`**: Handles CLI command definitions and user interaction.
- **`internal/cli/currency/`**: Contains the business logic for currency operations.
- **`internal/cache/`**: Implements a file-based TTL cache.
- **`internal/bootstrap/`**: Manages dependency injection and application startup.
- **`internal/logger/`**: Configures structured logging.

## Data Provider

Currency data is provided by the [Frankfurter API](https://api.frankfurter.dev/), an open-source and reliable source for current and historical exchange rates.

## License

This project is licensed under the MIT License – see the [LICENSE](LICENSE) file for details.
