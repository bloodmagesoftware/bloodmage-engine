local proto_files = Yab.find("internal", "**.proto")

-- compile proto files
for _, proto_file in ipairs(proto_files) do
	print("Compiling proto file: " .. proto_file)
    os.execute("protoc --go_out=. --go_opt=paths=source_relative " .. proto_file)
end
