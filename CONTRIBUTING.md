# Contributing to test-repo-standard-check

Thank you for your interest in contributing! This document provides guidelines for contributing to this project.

## Code of Conduct

Be respectful, inclusive, and professional in all interactions.

## Development Workflow

### 1. Fork and Clone

```bash
# Fork the repo on GitHub, then:
git clone https://github.com/YOUR-USERNAME/test-repo-standard-check.git
cd test-repo-standard-check
```

### 2. Create a Branch

```bash
git checkout -b feature/your-feature-name
# or
git checkout -b fix/your-bug-fix
```

### 3. Make Changes

- Write clean, idiomatic Go code
- Follow the existing code style
- Add tests for new functionality
- Update documentation as needed

### 4. Test Your Changes

```bash
# Run tests
make test

# Run linter
make lint

# Format code
make fmt
```

### 5. Commit Your Changes

- Write clear, descriptive commit messages
- Reference any related issues

```bash
git add .
git commit -m "Add feature: description of your change"
```

### 6. Push and Create PR

```bash
git push origin feature/your-feature-name
```

Then create a Pull Request on GitHub with:
- Clear description of changes
- Reference to any related issues
- Screenshots or examples if applicable

## Pull Request Guidelines

- **One feature per PR** - Keep changes focused
- **Pass all checks** - Ensure CI passes (lint + test)
- **Update tests** - Add tests for new code
- **Update docs** - Keep README and docs in sync
- **Clean history** - Squash commits if needed

## Code Style

- Follow standard Go conventions
- Use `gofmt` for formatting (run `make fmt`)
- Keep functions small and focused
- Add comments for exported functions and types
- Write idiomatic Go code

## Testing

- Write tests for all new functionality
- Maintain or improve code coverage
- Use table-driven tests where appropriate
- Test edge cases and error conditions

## Questions?

Feel free to open an issue for:
- Bug reports
- Feature requests
- Questions about contributing

Thank you for contributing! 🎉
