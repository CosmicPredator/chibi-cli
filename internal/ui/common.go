package ui

// Simple Key Value pair used for ordered output
// since map produces unordered kv pairs while iterating
type KV struct {
	Key   string
	Value string
}