local bin_name = require("project").bin_name
local exe_name = require("build")

Yab.zip({
	exe_name,
	"assets/",
}, bin_name .. ".zip")

os.remove(exe_name)
