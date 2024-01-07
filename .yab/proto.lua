-- Bloodmage Engine - Retro first person game engine
-- Copyright (C) 2024  Frank Mayer
--
-- This program is free software: you can redistribute it and/or modify
-- it under the terms of the GNU Affero General Public License as published by
-- the Free Software Foundation, either version 3 of the License, or
-- (at your option) any later version.
--
-- This program is distributed in the hope that it will be useful,
-- but WITHOUT ANY WARRANTY; without even the implied warranty of
-- MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
-- GNU Affero General Public License for more details.
--
-- You should have received a copy of the GNU Affero General Public License
-- along with this program.  If not, see <http://www.gnu.org/licenses/>.

local proto_files = Yab.find("engine", "**.proto")

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
