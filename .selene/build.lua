local pack = require("pack")
local bin_name = pack.bin_name
if Selene.os_type() == "windows" then
	bin_name = bin_name .. ".exe"
end

os.execute('go build -ldflags="-s -w" -o \'' .. bin_name .. "' game/main.go")
