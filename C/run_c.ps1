
Remove-Item -r -Force ./main.exe && gcc ./main.c -o ./main

Write-Output "Compiled C code ✔️"

./main.exe
