require("proto")

local pack = require("project")
local bin_name = pack.bin_name
if Yab.os_type() == "windows" then
	bin_name = bin_name .. ".exe"
end

os.execute('go build -ldflags="-s -w" -o \'' .. bin_name .. "' cmd/game/main.go")

return bin_name
