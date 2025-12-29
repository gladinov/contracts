package contextkeys

type TraceIDKeyStruct struct{}
type EncryptedTokenKeyStruct struct{}
type ChatIDKeyStruct struct{}

var TraceIDKey = TraceIDKeyStruct{}
var EncryptedTokenKey = EncryptedTokenKeyStruct{}
var ChatIDKey = ChatIDKeyStruct{}
