GOARCH=arm64 GOOS=linux go build -ldflags "-s -w" -o app .
cp app /mnt/orangepi/root/.gam-app/maldan-gam-app-calculator-v0.0.3