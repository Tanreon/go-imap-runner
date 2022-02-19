package imap_runner

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/textproto"
	"strings"

	"github.com/DusanKasan/parsemail"
	"github.com/emersion/go-imap"
	"github.com/nadoo/glider/rule"

	id "github.com/emersion/go-imap-id"
	imapClient "github.com/emersion/go-imap/client"
)

var ErrInboxEmpty = errors.New("INBOX_IS_EMPTY")
var ErrFetchBug = errors.New("FETCH_BUG")
var ErrImapServerNotRecognized = errors.New("IMAP_SERVER_NOT_RECOGNIZED")
var ErrInvalidCredentials = errors.New("IMAP_INVALID_CREDENTIALS")

type ImapRunner struct {
	client   *imapClient.Client
	email    string
	password string
}

func (i *ImapRunner) Logout() {
	i.client.Logout()
}

func (i *ImapRunner) Login() error {
	if err := i.client.Login(i.email, i.password); err != nil {
		if strings.EqualFold(err.Error(), "LOGIN Invalid credentials") {
			return ErrInvalidCredentials
		}

		return fmt.Errorf("error while signing in: %w", err)
	}

	return nil
}

func (i *ImapRunner) SelectMailBox(name string, readOnly bool) (mailbox *imap.MailboxStatus, err error) {
	return i.client.Select(name, readOnly)
}

func (i *ImapRunner) SelectInbox(readOnly bool) (mailbox *imap.MailboxStatus, err error) {
	return i.SelectMailBox("INBOX", readOnly)
}

func (i *ImapRunner) RemoveAllMessages(mailbox *imap.MailboxStatus) (err error) {
	if mailbox.Messages > 0 {
		seqSet := new(imap.SeqSet)
		seqSet.AddRange(uint32(1), mailbox.Messages)

		if err := i.client.Store(seqSet, imap.FormatFlagsOp(imap.AddFlags, true), []interface{}{imap.DeletedFlag}, nil); err != nil {
			return fmt.Errorf("set flag error: %w", err)
		}

		if err = i.client.Expunge(nil); err != nil {
			return fmt.Errorf("expunge error: %w", err)
		}
	}

	return nil
}

func (i *ImapRunner) GetAllMessages(mailbox *imap.MailboxStatus) (emails []parsemail.Email, err error) {
	if mailbox.Messages <= 0 {
		return emails, ErrInboxEmpty
	}

	seqSet := imap.SeqSet{}
	seqSet.AddRange(uint32(1), mailbox.Messages)

	section := &imap.BodySectionName{}
	items := []imap.FetchItem{section.FetchItem()}

	messages := make(chan *imap.Message, mailbox.Messages)

	if err := i.client.Fetch(&seqSet, items, messages); err != nil {
		if strings.EqualFold(err.Error(), "FETCH Server error while fetching messages") {
			return emails, ErrFetchBug
		}

		return emails, fmt.Errorf("fetching found messages error: %w", err)
	}

	for message := range messages {
		bodyReader := message.GetBody(&imap.BodySectionName{})
		if bodyReader == nil {
			return emails, fmt.Errorf("server didn't returned data body")
		}

		email, err := parsemail.Parse(bodyReader) // returns Email struct and error
		if err != nil {
			return emails, fmt.Errorf("mail parse error: %w", err)
		}

		emails = append(emails, email)
	}

	return emails, err
}

func (i *ImapRunner) GetAllMessagesRaw(mailbox *imap.MailboxStatus) (emails [][]byte, err error) {
	if mailbox.Messages <= 0 {
		return emails, ErrInboxEmpty
	}

	seqSet := imap.SeqSet{}
	seqSet.AddRange(uint32(1), mailbox.Messages)

	section := &imap.BodySectionName{}
	items := []imap.FetchItem{section.FetchItem()}

	messages := make(chan *imap.Message, mailbox.Messages)

	if err := i.client.Fetch(&seqSet, items, messages); err != nil {
		if strings.EqualFold(err.Error(), "FETCH Server error while fetching messages") {
			return emails, ErrFetchBug
		}

		return emails, fmt.Errorf("fetching found messages error: %w", err)
	}

	for message := range messages {
		bodyReader := message.GetBody(&imap.BodySectionName{})
		if bodyReader == nil {
			return emails, fmt.Errorf("server didn't returned data body")
		}

		emailBytes, err := ioutil.ReadAll(bodyReader)
		if err != nil {
			return emails, fmt.Errorf("mail read error: %w", err)
		}

		emails = append(emails, emailBytes)
	}

	return emails, err
}

func (i *ImapRunner) SearchSubject(phrase string, unseenOnly bool) (emails []parsemail.Email, err error) {
	searchCriteria := imap.NewSearchCriteria()
	searchCriteria.Header = textproto.MIMEHeader{"Subject": {phrase}}

	if unseenOnly {
		searchCriteria.WithoutFlags = []string{imap.SeenFlag}
	}

	foundIds, err := i.client.Search(searchCriteria)
	if err != nil {
		return emails, fmt.Errorf("search with search criteria error: %w", err)
	}

	//

	if len(foundIds) > 0 {
		seqSet := imap.SeqSet{}
		seqSet.AddNum(foundIds...)

		section := &imap.BodySectionName{}
		items := []imap.FetchItem{section.FetchItem()}

		messages := make(chan *imap.Message, 10)

		if err := i.client.Fetch(&seqSet, items, messages); err != nil {
			if strings.EqualFold(err.Error(), "FETCH Server error while fetching messages") {
				return emails, ErrFetchBug
			}

			return emails, fmt.Errorf("fetching found messages error: %w", err)
		}

		for message := range messages {
			bodyReader := message.GetBody(&imap.BodySectionName{})
			if bodyReader == nil {
				return emails, fmt.Errorf("server didn't returned message body")
			}

			email, err := parsemail.Parse(bodyReader) // returns Email struct and error
			if err != nil {
				return emails, fmt.Errorf("mail parse error: %w", err)
			}

			emails = append(emails, email)
		}
	}

	return emails, err
}

func (i *ImapRunner) SearchFrom(from string, unseenOnly bool) (emails []parsemail.Email, err error) {
	searchCriteria := imap.NewSearchCriteria()
	searchCriteria.Header = textproto.MIMEHeader{"From": {from}}

	if unseenOnly {
		searchCriteria.WithoutFlags = []string{imap.SeenFlag}
	}

	foundIds, err := i.client.Search(searchCriteria)
	if err != nil {
		return emails, fmt.Errorf("search with search criteria error: %w", err)
	}

	//

	if len(foundIds) > 0 {
		seqSet := imap.SeqSet{}
		seqSet.AddNum(foundIds...)

		section := &imap.BodySectionName{}
		items := []imap.FetchItem{section.FetchItem()}

		messages := make(chan *imap.Message, len(foundIds))

		if err := i.client.Fetch(&seqSet, items, messages); err != nil {
			if strings.EqualFold(err.Error(), "FETCH Server error while fetching messages") {
				return emails, ErrFetchBug
			}

			return emails, fmt.Errorf("fetching found messages error: %w", err)
		}

		for message := range messages {
			bodyReader := message.GetBody(&imap.BodySectionName{})
			if bodyReader == nil {
				return emails, fmt.Errorf("server didn't returned message body")
			}

			email, err := parsemail.Parse(bodyReader) // returns Email struct and error
			if err != nil {
				return emails, fmt.Errorf("mail parse error: %w", err)
			}

			emails = append(emails, email)
		}
	}

	return emails, err
}

func NewImapRunner(dialer *rule.Proxy, email, password string) (*ImapRunner, error) {
	var server string

	atPosition := strings.LastIndex(email, "@")
	if atPosition >= 0 {
		switch email[atPosition+1:] {
		case "yahoo.com":
			server = "imap.mail.yahoo.com:993"
		case "outlook.com":
			server = "outlook.office365.com:993"
		case "gmail.com":
			server = "imap.gmail.com:993"
		default:
			return nil, ErrImapServerNotRecognized
			//log.Printf("[%s] [%s] skip reason: unknown server", LOG_HEAD, imapData.login)
			//return
		}
	} else {
		return nil, ErrImapServerNotRecognized
		//log.Printf("[%s] [%s] skip reason: unknown server", LOG_HEAD, imapData.login)
		//return
	}

	tlsConfig := tls.Config{
		MinVersion: tls.VersionTLS12,
		MaxVersion: tls.VersionTLS12,
	}

	client, err := imapClient.DialWithDialerTLS(dialer.NextDialer(server), server, &tlsConfig)
	if err != nil {
		return nil, err
	}

	_, err = id.NewClient(client).ID(map[string]string{
		id.FieldName:    "Thunderbird",
		id.FieldVersion: "91.6.0",
		//[id.FieldOS] = "",
		//[id.FieldOSVersion] = "",
		//[id.FieldVendor] = "",
		//[id.FieldSupportURL] = "",
		//[id.FieldAddress] = "",
		//[id.FieldDate] = "",
		//[id.FieldCommand] = "",
		//[id.FieldArguments] = "",
		//[id.FieldEnvironment] = "",
	})
	if err != nil {
		return nil, err
	}

	//client.SetDebug(os.Stdout)

	return &ImapRunner{client, email, password}, err
}
