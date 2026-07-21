/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sandboxbackend

import (
	"errors"
	"fmt"
)

// DependencyPendingError reports that reconciliation is blocked on another
// Kubernetes object that may be created by the user or another controller later.
type DependencyPendingError struct {
	Reason  string
	Message string
	Err     error
}

func (e *DependencyPendingError) Error() string {
	if e == nil {
		return ""
	}
	if e.Err == nil {
		return e.Message
	}
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func (e *DependencyPendingError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}

func DependencyPendingFromError(err error) *DependencyPendingError {
	var dep *DependencyPendingError
	if errors.As(err, &dep) {
		return dep
	}
	return nil
}
