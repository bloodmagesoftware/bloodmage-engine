// Bloodmage Engine
// Copyright (C) 2023 Frank Mayer
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://github.com/bloodmagesoftware/bloodmage-engine/blob/main/LICENSE.md>.

package topdown

import (
	"github.com/bloodmagesoftware/bloodmage-engine/pkg/engine/core"
	"github.com/charmbracelet/log"
)

func Init() {
	f := func() {
		log.Debug("topdown window resize")
	}
	core.OnResize(&f)
}
