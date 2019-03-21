// Copyright 2017 Istio Authors
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

package pool

import (
	"sync"
	"testing"
)

func TestWorkerPool(t *testing.T) {
	const numWorkers = 123
	const numWorkItems = 456

	parameterMismatch := false

	for i := 0; i < 2; i++ {
		gp := NewGoroutinePool(128, i == 0)
		gp.AddWorkers(numWorkers)

		wg := &sync.WaitGroup{}
		wg.Add(numWorkItems)

		for i := 0; i < numWorkItems; i++ {
			passedParam := i // capture the parameter on stack to avoid closing on the loop variable.
			gp.ScheduleWork(func(param interface{}) {
				paramI := param.(int)
				if paramI != passedParam {
					parameterMismatch = true
				}
				wg.Done()
			}, passedParam)
		}

		// wait for all the functions to have run
		wg.Wait()

		if parameterMismatch {
			t.Fatal("Passed parameter was not as expected")
		}

		// make sure the pool can be shutdown cleanly
		_ = gp.Close()
	}
}
