# push_one_by_one.ps1
Write-Host "Starting sequential push..." -ForegroundColor Cyan

$status = git status --porcelain=v1 -uall
if (-not $status) {
    Write-Host "No changes to push." -ForegroundColor Green
    exit
}

$files = @()
foreach ($line in $status) {
    if ($line.Length -ge 4) {
        $path = $line.Substring(3).Trim()
        if ($path.StartsWith('"')) {
            $path = $path.Substring(1, $path.Length - 2)
        }
        $files += $path
    }
}

$branch = git branch --show-current
$total = $files.Count

for ($i = 0; $i -lt $total; $i++) {
    $file = $files[$i]
    $num = $i + 1
    Write-Host "[$num/$total] Processing: $file" -ForegroundColor Yellow
    
    git add "$file"
    if ($LASTEXITCODE -eq 0) {
        git commit -m "chore: update file $num - $file"
        if ($LASTEXITCODE -eq 0) {
            Write-Host "  Pushing..." -ForegroundColor DarkGray
            git push origin "$branch"
            if ($LASTEXITCODE -ne 0) {
                Write-Host "  Push failed for $file." -ForegroundColor Red
                break
            }
        } else {
            Write-Host "  Nothing to commit for $file." -ForegroundColor DarkYellow
        }
    } else {
        Write-Host "  Failed to stage $file." -ForegroundColor Red
    }
}

Write-Host "Process completed." -ForegroundColor Green
