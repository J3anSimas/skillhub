watch-tailwindcss:
	@powershell -ExecutionPolicy Bypass -Command "\
	if (-not (Get-Command tailwindcss -ErrorAction SilentlyContinue)) { \
		Write-Output 'Installing tailwindcss...'; \
		go install github.com/tailwindlabs/tailwindcss@latest; \
		$$env:PATH += ';' + (go env GOPATH) + '\bin'; \
	} ; \
	Write-Output 'Watching TailwindCSS...'; \
	tailwindcss -i cmd/web/styles/input.css -o cmd/web/assets/css/output.css --watch"

watch-templ:
	@powershell -ExecutionPolicy Bypass -Command "if (Get-Command templ -ErrorAction SilentlyContinue) { \
		templ generate --watch --proxy="http://localhost:8080" --cmd="go run cmd/api/main.go"; \
		Write-Output 'Watching...'; \
	} else { \
		Write-Output 'Installing templ...'; \
		go install github.com/a-h/templ/cmd/templ@latest; \
		templ generate; \
		Write-Output 'Watching...'; \
	}"

watch: watch-tailwindcss watch-templ
.PHONY: watch watch-tailwindcss watch-templ
