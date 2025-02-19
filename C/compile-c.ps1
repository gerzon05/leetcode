#requires -Version 5.1

param (
    [Parameter(Mandatory = $true, Position = 0)]
    [string]$SourceFile
)

# Define colors
$GREEN = "`e[0;32m"
$NC = "`e[0m"

# Get the filename without the extension
$filename = [System.IO.Path]::GetFileNameWithoutExtension($SourceFile)

# Remove the directory if it exists
if (Test-Path -Path "./$filename.exe" -PathType Container) {
    Remove-Item -Path "./$filename.exe" -Recurse -Force
}

# Compile the C code
Write-Host "Compiling $SourceFile..."
$gccOutput = gcc "$SourceFile" -o "$filename.exe" 2>&1 # Redirect stderr to stdout
$gccExitCode = $LASTEXITCODE

if ($gccExitCode -eq 0) {
    Write-Host "${GREEN}Compiled C code ✅${NC}"

    # Execute the compiled program
    Write-Host "Executing ./$filename.exe..."
    & ".\$filename.exe"

    # Remove the binary
    Remove-Item -Path "./$filename.exe" -Force
    Write-Host "${GREEN}Binary cleaned 🧹${NC}"
} else {
    Write-Host "Error al compilar:"
    Write-Host $gccOutput
}
