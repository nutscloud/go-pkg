// Copyright 2025 nutscloud authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package syncmap

import "sync"

type SyncMap[KEY comparable, VALUE any] struct {
	m map[KEY]VALUE
	sync.RWMutex
}

func New[KEY comparable, VALUE any]() *SyncMap[KEY, VALUE] {
	return &SyncMap[KEY, VALUE]{
		m: make(map[KEY]VALUE),
	}
}

func (sm *SyncMap[KEY, VALUE]) Load(key KEY) (val VALUE, ok bool) {
	if sm == nil {
		return
	}

	sm.RLock()
	defer sm.RUnlock()

	val, ok = sm.m[key]
	return
}

func (sm *SyncMap[KEY, VALUE]) Set(key KEY, val VALUE) {
	if sm == nil {
		return
	}

	sm.Lock()
	defer sm.Unlock()

	sm.m[key] = val
}

func (sm *SyncMap[KEY, VALUE]) Delete(key KEY) {
	if sm == nil {
		return
	}

	sm.Lock()
	defer sm.Unlock()

	delete(sm.m, key)
}

func (sm *SyncMap[KEY, VALUE]) Range(f func(key KEY, val VALUE) bool) {
	if sm == nil {
		return
	}

	if f == nil {
		panic("syncmap.Range: f cannot be nil")
	}

	sm.RLock()
	defer sm.RUnlock()

	for k, v := range sm.m {
		if !f(k, v) {
			break
		}
	}
}

func (sm *SyncMap[KEY, VALUE]) Len() int {
	if sm == nil {
		return 0
	}

	sm.RLock()
	defer sm.RUnlock()

	return len(sm.m)
}

func (sm *SyncMap[KEY, VALUE]) Clear() {
	if sm == nil {
		return
	}

	sm.Lock()
	defer sm.Unlock()

	sm.m = make(map[KEY]VALUE)
}
func (sm *SyncMap[KEY, VALUE]) Keys() (keys []KEY) {
	if sm == nil {
		return keys
	}

	sm.RLock()
	defer sm.RUnlock()

	keys = make([]KEY, 0, len(sm.m))
	for key := range sm.m {
		keys = append(keys, key)
	}
	return keys
}
