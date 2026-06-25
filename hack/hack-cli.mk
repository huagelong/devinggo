
# Install/Update to the latest CLI tool (cross-platform via go install).
.PHONY: cli
cli:
	@go install github.com/gogf/gf/cmd/gf/v2@latest


# Ensure CLI tool is installed (idempotent, works on Windows/Linux/macOS).
.PHONY: cli.install
cli.install: cli
