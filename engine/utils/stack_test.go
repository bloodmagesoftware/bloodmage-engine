// Bloodmage Engine - Retro first person game engine
// Copyright (C) 2024  Frank Mayer
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package utils_test

import (
	"testing"

	"github.com/bloodmagesoftware/bloodmage-engine/engine/utils"
)

func TestStack(t *testing.T) {
	t.Parallel()
	s := utils.NewStack[int]()

	if s.Len() != 0 {
		t.Errorf("Stack should be empty")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Len() != 3 {
		t.Errorf("Stack should have 3 elements")
	}

	if v, b := s.Pop(); *v != 3 || !b {
		t.Errorf("Stack should return 3")
	}

	if s.Len() != 2 {
		t.Errorf("Stack should have 2 elements")
	}

	if v, b := s.Pop(); *v != 2 || !b {
		t.Errorf("Stack should return 2")
	}

	if s.Len() != 1 {
		t.Errorf("Stack should have 1 element")
	}

	if v, b := s.Pop(); *v != 1 || !b {
		t.Errorf("Stack should return 1")
	}

	if s.Len() != 0 {
		t.Errorf("Stack should be empty")
	}

	if v, b := s.Pop(); v != nil || b {
		t.Errorf("Stack should return nil and false")
	}

	if s.Len() != 0 {
		t.Errorf("Stack should be empty")
	}
}
