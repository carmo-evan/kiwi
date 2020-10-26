// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package stdkiwi

import (
	"testing"

	"github.com/sdslabs/kiwi"
)

func TestStore(t *testing.T) {
	store := NewStore()
	if store == nil {
		t.Error("NewStore returned a nil value")
	}

	invalidSchema := make(map[string]kiwi.ValueType)
	invalidSchema["foo"] = "bar"

	if _, err := NewStoreFromSchema(invalidSchema); err == nil {
		t.Error("Empty schema should have returned an error for NewStoreFromSchema")
	}
}
