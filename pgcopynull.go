package pgcopy2sql

type PgcopyNull struct{}

func (n PgcopyNull) Reset()                      {}
func (n PgcopyNull) WriteTo(w ValueWriter) error { return w.WriteNull() }
func (n PgcopyNull) IsNull() bool                { return true }
func (n PgcopyNull) AsValue() Value              { return n }
func (n PgcopyNull) String() string              { return "Null" }
