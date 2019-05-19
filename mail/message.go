package mail

import "strings"

type Message struct {
	from        string
	to          []string
	subject     string
	body        string
	inline      []string
	attachments []string
}

func (m *Message) To(to string) {
	m.to = append(m.to, to)
}

func (m *Message) From(from string) {
	m.from = from
}

func (m *Message) Subject(subject string) {
	m.subject = subject
}

func (m *Message) Body(body string) {
	m.body = body
}

func (m *Message) Attach(attach string) {
	m.attachments = append(m.attachments, attach)
}

func (m *Message) Inline(inline string) {
	m.inline = append(m.inline, inline)
}

func (m *Message) getMsg() []byte {
	return []byte("To: " + strings.Join(m.to, ", ") + "\r\n" +
		"Subject: " + m.subject + "\r\n" +
		"\r\n" +
		m.body + "\r\n")
}
