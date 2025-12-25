@echo off
set /p tagname="Enter tag name (e.g., v0.1.0): "

echo.
echo === Cleaning up Go modules ===
go mod tidy

echo.
echo === Staging all changes ===
git add .

echo.
echo === Committing changes ===
git commit -m "chore: prepare release %tagname%"

echo.
echo === Pushing to main ===
git push origin main

echo.
echo === Creating and pushing tag %tagname% ===
git tag -a %tagname% -m "Release %tagname%"
git push origin %tagname%

echo.
echo Done! GitHub Actions will now trigger the GoReleaser pipeline.
echo Check your progress at: https://github.com/004Ongoro/gitaid/actions
pause