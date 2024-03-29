package imap_runner

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net"
	"net/textproto"
	"sort"
	"strconv"
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
		if strings.Contains(err.Error(), "LOGIN Invalid credentials") {
			return ErrInvalidCredentials
		}
		if strings.Contains(err.Error(), "LOGIN failed") {
			return ErrInvalidCredentials
		}
		if strings.Contains(err.Error(), "Account is blocked") || strings.Contains(err.Error(), "Login to your account via a web browser to verify your identity") {
			return ErrInvalidCredentials
		}
		if strings.Contains(err.Error(), "Please log in via your web browser") || strings.Contains(err.Error(), "https://support.google.com/mail/accounts/answer/78754") {
			return ErrInvalidCredentials
		}
		if isConnectionError(err) {
			return &net.OpError{Op: "read", Err: err}
		}

		return fmt.Errorf("error while signing in: %w", err)
	}

	return nil
}

func (i *ImapRunner) SelectMailBox(name string, readOnly bool) (mailbox *imap.MailboxStatus, err error) {
	mailbox, err = i.client.Select(name, readOnly)
	if err != nil && isConnectionError(err) {
		return mailbox, &net.OpError{Op: "read", Err: err}
	}

	return mailbox, err
}

func (i *ImapRunner) SelectInbox(readOnly bool) (mailbox *imap.MailboxStatus, err error) {
	return i.SelectMailBox("INBOX", readOnly)
}

func (i *ImapRunner) RemoveAllMessages(mailbox *imap.MailboxStatus) (err error) {
	if mailbox.Messages <= 0 {
		return err
	}

	seqSet := new(imap.SeqSet)
	seqSet.AddRange(1, mailbox.Messages)

	if err := i.client.Store(seqSet, imap.FormatFlagsOp(imap.AddFlags, true), []interface{}{imap.DeletedFlag}, nil); err != nil {
		if isConnectionError(err) {
			return &net.OpError{Op: "read", Err: err}
		}

		return fmt.Errorf("set flag error: %w", err)
	}

	if err = i.client.Expunge(nil); err != nil {
		if isConnectionError(err) {
			return &net.OpError{Op: "read", Err: err}
		}

		return fmt.Errorf("expunge error: %w", err)
	}

	return err
}

func (i *ImapRunner) RemoveMessage(mailbox *imap.MailboxStatus) (err error) {
	if mailbox.Messages <= 0 {
		return err
	}

	seqSet := new(imap.SeqSet)
	seqSet.AddRange(1, mailbox.Messages)

	if err := i.client.Store(seqSet, imap.FormatFlagsOp(imap.AddFlags, true), []interface{}{imap.DeletedFlag}, nil); err != nil {
		if isConnectionError(err) {
			return &net.OpError{Op: "read", Err: err}
		}

		return fmt.Errorf("set flag error: %w", err)
	}

	if err = i.client.Expunge(nil); err != nil {
		if isConnectionError(err) {
			return &net.OpError{Op: "read", Err: err}
		}

		return fmt.Errorf("expunge error: %w", err)
	}

	return err
}

func (i *ImapRunner) RemoveMessages(mailbox *imap.MailboxStatus, messageUids []string) (err error) {
	if mailbox.Messages <= 0 {
		return err
	}

	seqSet := new(imap.SeqSet)
	for _, uid := range messageUids {
		seqSet.Add(uid)
	}

	if err := i.client.UidStore(seqSet, imap.FormatFlagsOp(imap.AddFlags, true), []interface{}{imap.DeletedFlag}, nil); err != nil {
		if isConnectionError(err) {
			return &net.OpError{Op: "read", Err: err}
		}

		return fmt.Errorf("set flag error: %w", err)
	}

	if err = i.client.Expunge(nil); err != nil {
		if isConnectionError(err) {
			return &net.OpError{Op: "read", Err: err}
		}

		return fmt.Errorf("expunge error: %w", err)
	}

	return err
}

func (i *ImapRunner) GetAllMessages(mailbox *imap.MailboxStatus) (emails []parsemail.Email, err error) {
	return i.GetMessages(mailbox, math.MaxInt32)
}

func (i *ImapRunner) GetMessages(mailbox *imap.MailboxStatus, maxCount uint32) (emails []parsemail.Email, err error) {
	if mailbox.Messages <= 0 {
		return emails, ErrInboxEmpty
	}

	min := int32(mailbox.Messages) + 1 - int32(maxCount)
	if min < 1 {
		min = 1
	}

	seqSet := imap.SeqSet{}
	seqSet.AddRange(uint32(min), mailbox.Messages)

	section := &imap.BodySectionName{}
	items := append(imap.FetchAll.Expand(), imap.FetchUid, section.FetchItem())

	messages := make(chan *imap.Message, (seqSet.Set[0].Stop-seqSet.Set[0].Start)+1)

	if err := i.client.Fetch(&seqSet, items, messages); err != nil {
		if strings.EqualFold(err.Error(), "FETCH Server error while fetching messages") {
			return emails, ErrFetchBug
		}

		if isConnectionError(err) {
			return emails, &net.OpError{Op: "read", Err: err}
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

		email.MessageID = strconv.Itoa(int(message.Uid)) // not in RFC, костыль по русски говоря
		email.Date = message.InternalDate                // mailparser bug
		emails = append(emails, email)
	}

	sort.Slice(emails, func(i, j int) bool {
		return emails[i].Date.After(emails[j].Date)
	})

	return emails, err
}

func (i *ImapRunner) GetAllMessagesRaw(mailbox *imap.MailboxStatus) (emails [][]byte, err error) {
	return i.GetMessagesRaw(mailbox, math.MaxInt32)
}

func (i *ImapRunner) GetMessagesRaw(mailbox *imap.MailboxStatus, maxCount uint32) (emails [][]byte, err error) {
	if mailbox.Messages <= 0 {
		return emails, ErrInboxEmpty
	}

	min := int32(mailbox.Messages) + 1 - int32(maxCount)
	if min < 1 {
		min = 1
	}

	seqSet := imap.SeqSet{}
	seqSet.AddRange(uint32(min), mailbox.Messages)

	section := &imap.BodySectionName{}
	items := append(imap.FetchAll.Expand(), imap.FetchUid, section.FetchItem())

	messages := make(chan *imap.Message, (seqSet.Set[0].Stop-seqSet.Set[0].Start)+1)

	if err := i.client.Fetch(&seqSet, items, messages); err != nil {
		if strings.EqualFold(err.Error(), "FETCH Server error while fetching messages") {
			return emails, ErrFetchBug
		}

		if isConnectionError(err) {
			return emails, &net.OpError{Op: "read", Err: err}
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
		if isConnectionError(err) {
			return emails, &net.OpError{Op: "read", Err: err}
		}

		return emails, fmt.Errorf("search with search criteria error: %w", err)
	}

	//

	if len(foundIds) > 0 {
		seqSet := imap.SeqSet{}
		seqSet.AddNum(foundIds...)

		section := &imap.BodySectionName{}
		items := append(imap.FetchAll.Expand(), imap.FetchUid, section.FetchItem())

		messages := make(chan *imap.Message, 10)

		if err := i.client.Fetch(&seqSet, items, messages); err != nil {
			if strings.EqualFold(err.Error(), "FETCH Server error while fetching messages") {
				return emails, ErrFetchBug
			}

			if isConnectionError(err) {
				return emails, &net.OpError{Op: "read", Err: err}
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

			email.MessageID = strconv.Itoa(int(message.Uid)) // not in RFC, костыль по русски говоря
			email.Date = message.InternalDate                // mailparser bug
			emails = append(emails, email)
		}
	}

	sort.Slice(emails, func(i, j int) bool {
		return emails[i].Date.After(emails[j].Date)
	})

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
		if isConnectionError(err) {
			return emails, &net.OpError{Op: "read", Err: err}
		}

		return emails, fmt.Errorf("search with search criteria error: %w", err)
	}

	//

	if len(foundIds) > 0 {
		seqSet := imap.SeqSet{}
		seqSet.AddNum(foundIds...)

		section := &imap.BodySectionName{}
		items := append(imap.FetchAll.Expand(), imap.FetchUid, section.FetchItem())

		messages := make(chan *imap.Message, len(foundIds))

		if err := i.client.Fetch(&seqSet, items, messages); err != nil {
			if strings.EqualFold(err.Error(), "FETCH Server error while fetching messages") {
				return emails, ErrFetchBug
			}

			if isConnectionError(err) {
				return emails, &net.OpError{Op: "read", Err: err}
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

			email.MessageID = strconv.Itoa(int(message.Uid)) // not in RFC, костыль по русски говоря
			email.Date = message.InternalDate                // mailparser bug
			emails = append(emails, email)
		}
	}

	sort.Slice(emails, func(i, j int) bool {
		return emails[i].Date.After(emails[j].Date)
	})

	return emails, err
}

func (i *ImapRunner) MoveMessages(mailboxFrom *imap.MailboxStatus, mailboxTo string, messageUids []uint32) (err error) {
	if mailboxFrom.Messages <= 0 {
		return err
	}

	seqSet := new(imap.SeqSet)
	seqSet.AddNum(messageUids...)

	if err := i.client.UidMove(seqSet, mailboxTo); err != nil {
		if isConnectionError(err) {
			return &net.OpError{Op: "read", Err: err}
		}

		return fmt.Errorf("mail move error: %w", err)
	}

	return err
}

func NewImapRunner(dialer *rule.Proxy, email, password string) (*ImapRunner, error) {
	server, err := parseServer(email)
	if err != nil {
		return nil, fmt.Errorf("imap server recognize error: %w", err)
	}

	tlsConfig := tls.Config{
		MinVersion: tls.VersionTLS12,
		MaxVersion: tls.VersionTLS12,
	}

	client, err := imapClient.DialWithDialerTLS(dialer.NextDialer(server), server, &tlsConfig)
	if err != nil {
		if isConnectionError(err) {
			return nil, &net.OpError{Op: "read", Err: err}
		}

		return nil, err
	}

	_, err = id.NewClient(client).ID(map[string]string{
		id.FieldName:    "Thunderbird",
		id.FieldVersion: "102.2.1",
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
		if isConnectionError(err) {
			return nil, &net.OpError{Op: "read", Err: err}
		}

		return nil, err
	}

	//client.SetDebug(os.Stdout)

	return &ImapRunner{client, email, password}, err
}
