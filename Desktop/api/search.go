package api

import (
	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://tempuri.org/"

// NewIContactSearchService creates an initializes a IContactSearchService.
func NewIContactSearchService(cli *soap.Client) IContactSearchService {
	return &iContactSearchService{cli}
}

// IContactSearchService was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type IContactSearchService interface {
	// SearchContacts was auto-generated from WSDL.
	SearchContacts(SearchContacts *SearchContacts) (*SearchContactsResponse, error)
}

// Char was auto-generated from WSDL.
type Char int

// Guid was auto-generated from WSDL.
type Guid string

// ArrayOfContactRecord was auto-generated from WSDL.
type ArrayOfContactRecord struct {
	ContactRecord []*ContactRecord `xml:"ContactRecord,omitempty" json:"ContactRecord,omitempty" yaml:"ContactRecord,omitempty"`
}

// ArrayOfString was auto-generated from WSDL.
type ArrayOfString struct {
	String []*string `xml:"string,omitempty" json:"string,omitempty" yaml:"string,omitempty"`
}

// ArrayOfstring was auto-generated from WSDL.
type ArrayOfstring struct {
	String []*string `xml:"string,omitempty" json:"string,omitempty" yaml:"string,omitempty"`
}

// ContactRecord was auto-generated from WSDL.
type ContactRecord struct {
	RecordType *string        `xml:"RecordType,omitempty" json:"RecordType,omitempty" yaml:"RecordType,omitempty"`
	Sources    *ArrayOfstring `xml:"Sources,omitempty" json:"Sources,omitempty" yaml:"Sources,omitempty"`
	Value      *string        `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

// SearchContacts was auto-generated from WSDL.
type SearchContacts struct {
	SearchTerm *string `xml:"searchTerm,omitempty" json:"searchTerm,omitempty" yaml:"searchTerm,omitempty"`
}

// SearchContactsResponse was auto-generated from WSDL.
type SearchContactsResponse struct {
	SearchContactsResult *ArrayOfContactRecord `xml:"SearchContactsResult,omitempty" json:"SearchContactsResult,omitempty" yaml:"SearchContactsResult,omitempty"`
}

// Operation wrapper for SearchContacts.
// OperationIContactSearchService_SearchContacts_InputMessage was
// auto-generated from WSDL.
type OperationIContactSearchService_SearchContacts_InputMessage struct {
	SearchContacts *SearchContacts `xml:"SearchContacts,omitempty" json:"SearchContacts,omitempty" yaml:"SearchContacts,omitempty"`
}

// Operation wrapper for SearchContacts.
// OperationIContactSearchService_SearchContacts_OutputMessage
// was auto-generated from WSDL.
type OperationIContactSearchService_SearchContacts_OutputMessage struct {
	SearchContactsResponse *SearchContactsResponse `xml:"SearchContactsResponse,omitempty" json:"SearchContactsResponse,omitempty" yaml:"SearchContactsResponse,omitempty"`
}

// iContactSearchService implements the IContactSearchService interface.
type iContactSearchService struct {
	cli *soap.Client
}

// SearchContacts was auto-generated from WSDL.
func (p *iContactSearchService) SearchContacts(SearchContacts *SearchContacts) (*SearchContactsResponse, error) {
	α := struct {
		OperationIContactSearchService_SearchContacts_InputMessage `xml:"tns:SearchContacts"`
	}{
		OperationIContactSearchService_SearchContacts_InputMessage{
			SearchContacts,
		},
	}

	γ := struct {
		OperationIContactSearchService_SearchContacts_OutputMessage `xml:"SearchContactsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://tempuri.org/IContactSearchService/SearchContacts", α, &γ); err != nil {
		return nil, err
	}
	return γ.SearchContactsResponse, nil
}
