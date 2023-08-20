package core

import (
	"testing"
	"time"
)

func TestBlock_SetHash(t *testing.T) {
	data := "Test Data 1"
	type fields struct {
		Timestamp     int64
		Data          []byte
		PrevBlockHash []byte
		Hash          []byte
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Success Set Hash",
			fields: fields{
				Timestamp:     time.Now().Unix(),
				Data:          []byte(data),
				PrevBlockHash: []byte{},
				Hash:          []byte("123"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Block{
				Timestamp:     tt.fields.Timestamp,
				Data:          tt.fields.Data,
				PrevBlockHash: tt.fields.PrevBlockHash,
				Hash:          tt.fields.Hash,
			}
			b.SetHash()
		})
	}
}
