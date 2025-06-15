using System.Xml.Linq;
using WebsiteContactsService.Contracts;
using WebsiteContactsService.Models;

namespace WebsiteContactsService.Services
{
    public class ContactSearchService : IContactSearchService
    {
        private const string XmlFilePath = "contacts_data.xml";
        private readonly RapidApiService _rapidApiService;

        public ContactSearchService(RapidApiService rapidApiService)
        {
            _rapidApiService = rapidApiService;
        }

        public List<Contact> SearchContacts(string searchTerm)
        {
            _rapidApiService.FetchAndGenerateXmlAsync(searchTerm).Wait();

            var xDoc = XDocument.Load(XmlFilePath);
            var matchingContacts = xDoc.Descendants("Contact")
                .Where(contact => string.Equals(contact.Element("Domain")?.Value, searchTerm, StringComparison.OrdinalIgnoreCase))
                .Select(contact => new Contact
                {
                    Domain = contact.Element("Domain")?.Value,
                    Query = contact.Element("Query")?.Value,
                    Emails = contact.Element("Emails")?.Elements("Email")
                        .Select(email => new EmailEntry
                        {
                            Value = email.Element("Value")?.Value,
                            Sources = email.Element("Sources")?.Elements("Source")
                                .Select(source => source.Value).ToList()
                        }).ToList(),
                    PhoneNumbers = contact.Element("PhoneNumbers")?.Elements("Phone")
                        .Select(phone => new PhoneEntry
                        {
                            Value = phone.Element("Value")?.Value,
                            Sources = phone.Element("Sources")?.Elements("Source")
                                .Select(source => source.Value).ToList()
                        }).ToList(),
                    Facebook = contact.Element("Facebook")?.Value,
                    Instagram = contact.Element("Instagram")?.Value,
                    Twitter = contact.Element("Twitter")?.Value,
                    Youtube = contact.Element("Youtube")?.Value
                })
                .ToList();

            return matchingContacts;
        }
    }
}