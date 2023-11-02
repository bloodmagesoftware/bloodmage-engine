local bin_name = selene_os_type() == "windows" and "bloodmage_game.exe" or "bloodmage_game"

os.execute("go build -o " .. bin_name .. " game/main.go")
