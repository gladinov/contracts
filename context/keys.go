package contextkeys

type ctxKey struct{}

var TraceIDKey = ctxKey{}
var EncryptedTokenKey = ctxKey{}
var ChatIDKey = ctxKey{}
