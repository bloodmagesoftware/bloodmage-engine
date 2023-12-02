require("proto")
local level = Yab.args()[1]
if level == nil then
    print("No level specified")
    os.exit(1)
end
os.execute("go run ./cmd/editor/ -level=" .. level)
