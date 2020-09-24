package core
import "bytes"


type View struct {
	Tp 		string
	Attr 	*BindMap
	Raw 	interface{}
	Child 	*TokenList
}


func (v *View) ToString() string {
	var buffer bytes.Buffer
	buffer.WriteString("!view{\"")
	buffer.WriteString(v.Tp)
	
	if nil == v.Attr {
		buffer.WriteString("\" {}")
	}else{
		buffer.WriteString("\" {")
		for k, v := range v.Attr.Table {
			buffer.WriteString(k)
			buffer.WriteString(": ")
			buffer.WriteString(v.ToString())
			buffer.WriteString(" ")
		}
		if len(v.Attr.Table) > 0 {
			buffer.Bytes()[len(buffer.Bytes())-1] = '}'
		}else{
			buffer.WriteString("}")
		}
	}
	
	if nil == v.Child {
		buffer.WriteString(" []")
	}else{
		buffer.WriteString(" [")

		for _, c := range v.Child.List() {
			buffer.WriteString(c.ToString())
			buffer.WriteString(" ")
		}
		if v.Child.Len() > 0 {
			buffer.Bytes()[len(buffer.Bytes())-1] = ']'
		}else{
			buffer.WriteString("]")
	}
	}



	buffer.WriteString("}")
	return buffer.String()
}




