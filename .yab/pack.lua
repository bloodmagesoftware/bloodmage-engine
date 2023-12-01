local pack = require("project")
local bin_name = pack.bin_name

local exe_name = require("build")

Yab.zip({
	exe_name,
	"assets/",
}, bin_name .. ".zip")

os.remove(exe_name)
