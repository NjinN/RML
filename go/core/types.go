package core

const (
	NIL			= iota
	NONE		
	ERR 		
	LIT_WORD	
	GET_WORD 	
	DATATYPE 	
	LOGIC 		
	INTEGER 	
	DECIMAL 	
	CHAR	  	
	STRING 		
	PAREN 		
	BLOCK
	PROP
	WORD
	SET_WORD
	PUT_WORD 		
	PATH 			
	OP 			
	NATIVE 		
	FUNC 		
)

func TypeStr(n int) string{
	switch n {
	case NIL:
		return "nil!"
	case NONE:
		return "none!"
	case ERR:
		return "error!"
	case LIT_WORD:
		return "lit-word!"
	case GET_WORD:
		return "get-word!"
	case DATATYPE:
		return "datatype!"
	case LOGIC:
		return "logic!"
	case INTEGER:
		return "integer!"
	case DECIMAL:
		return "decimal!"
	case CHAR:
		return "char!"
	case STRING:
		return "string!"
	case PAREN:
		return "paren!"
	case BLOCK:
		return "block!"
	case PROP:
		return "prop!"
	case PATH:
		return "path!"
	case WORD:
		return "word!"
	case SET_WORD:
		return "set-word!"
	case PUT_WORD:
		return "put-word!"
	case OP:
		return "op!"
	case NATIVE:
		return "native!"
	case FUNC:
		return "function!"
	default:
		return "404 NOT FOUND!"
	}
}
