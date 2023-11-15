Selene.git_clone_or_pull("https://github.com/exosite/lua-yaml.git", "mod/yaml")

local function read(file)
	local f = assert(io.open(file, "r")) -- Open the file in read mode

	-- Read the entire file content as a string
	local content = f:read("*all")

	f:close() -- Close the file

	return content -- Return the file content as a string
end

local yaml = require("mod.yaml.yaml")
local str = read("project.yaml")
local t = yaml.eval(str)

local expected_fields = { "bin_name" }
for _, field in ipairs(expected_fields) do
	if t[field] == nil then
		error("Missing field: " .. field .. " in project.yaml")
	end
end

return t
