# Contributing Guide

## Getting Started

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Development Setup

```bash
# Clone and setup
git clone https://github.com/mkhaonguyen42-stack/jxcf-api.git
cd jxcf-api

# Install dependencies
go mod download

# Setup environment
cp .env.example .env
# Edit .env with your DeepSeek API key

# Run tests
make test
```

## Code Style

- Follow Go conventions (gofmt, golint)
- Use meaningful variable names
- Add comments for exported functions
- Write tests for new features

## Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run specific test
go test -v ./internal/service -run TestAnalyzer
```

## Commit Messages

Follow conventional commits:
- `feat:` for new features
- `fix:` for bug fixes
- `docs:` for documentation changes
- `test:` for test additions
- `chore:` for maintenance

Example: `feat: Add batch article analysis`

## Pull Request Process

1. Update documentation
2. Add/update tests
3. Ensure all tests pass
4. Request review from maintainers

## Issues

- Use GitHub Issues for bugs and feature requests
- Provide detailed reproduction steps for bugs
- Include expected vs actual behavior

## License

MIT License - see LICENSE file for details
