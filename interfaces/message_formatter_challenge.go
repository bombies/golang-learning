package interfaces

type Formatter interface {
	Format() string
}

type Message struct {
	message string
}

type PlainText struct {
	Message
}

type Bold struct {
	Message
}

type Code struct {
	Message
}

func (p PlainText) Format() string {
	return p.message
}

func (b Bold) Format() string {
	return "**" + b.message + "**"
}

func (c Code) Format() string {
	return "`" + c.message + "`"
}

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}
