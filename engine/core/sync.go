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

package core

import "sync"

type mainThreadFunc func() error

var (
	// mainThreadQueue contains functions to be run on the main thread
	mainThreadQueue = make([]mainThreadFunc, 0)
	// mainThreadMutex is used to lock the main thread when running queued functions
	mainThreadMutex = &sync.Mutex{}
)

// RunOnMainThread runs the given function on the main thread.
func RunOnMainThread(fn mainThreadFunc) {
	mainThreadQueue = append(mainThreadQueue, fn)
}

func runMainThreadQueue() error {
	mainThreadMutex.Lock()
	defer mainThreadMutex.Unlock()

	for _, fn := range mainThreadQueue {
		if err := fn(); err != nil {
			return err
		}
	}

	mainThreadQueue = make([]mainThreadFunc, 0)
	return nil
}
