local proto_files = Yab.find("pkg", "**.proto")

-- compile proto files
for _, proto_file in ipairs(proto_files) do
	local proto_info = Yab.fileinfo(proto_file)
	local out_file = proto_file:sub(1, -6) .. "pb.go"
	local success, out_info = pcall(Yab.fileinfo, out_file)

	if not success or proto_info.modtime > out_info.modtime then
		print("Compiling proto file: " .. proto_file)
		os.execute("protoc --go_out=. --go_opt=paths=source_relative " .. proto_file)
	else
		print("Skip compiling proto file: " .. proto_file)
	end
end
