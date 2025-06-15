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
            if (File.Exists(XmlFilePath))
            {
                var existingDoc = XDocument.Load(XmlFilePath);
                var existingContact = existingDoc.Root?
                    .Elements("Contact")
                    .FirstOrDefault(c => c.Element("Domain")?.Value == searchTerm);

                if (existingContact != null)
                {
                    Console.WriteLine($"Contact with domain '{searchTerm}' already exists. Skipping processing.");
                    Console.WriteLine($"Fetched data from XML: {existingContact}");
                    var contact = new Contact
                    {
                        Domain = existingContact.Element("Domain")?.Value,
                        Query = existingContact.Element("Query")?.Value,
                        Emails = existingContact.Element("Emails")?.Elements("Email").Select(e => new EmailEntry
                        {
                            Value = e.Element("Value")?.Value,
                            Sources = e.Element("Sources")?.Elements("Source").Select(s => s.Value).ToList()
                        }).ToList(),
                        PhoneNumbers = existingContact.Element("PhoneNumbers")?.Elements("Phone").Select(p => new PhoneEntry
                        {
                            Value = p.Element("Value")?.Value,
                            Sources = p.Element("Sources")?.Elements("Source").Select(s => s.Value).ToList()
                        }).ToList(),
                        Facebook = existingContact.Element("Facebook")?.Value,
                        Instagram = existingContact.Element("Instagram")?.Value,
                        Tiktok = existingContact.Element("Tiktok")?.Value,
                        Snapchat = existingContact.Element("Snapchat")?.Value,
                        Twitter = existingContact.Element("Twitter")?.Value,
                        Linkedin = existingContact.Element("Linkedin")?.Value,
                        Github = existingContact.Element("Github")?.Value,
                        Youtube = existingContact.Element("Youtube")?.Value,
                        Pinterest = existingContact.Element("Pinterest")?.Value
                    };

                    return new List<Contact> { contact };
                }
            }

            try
            {
                Console.WriteLine($"Fetching data for search term: {searchTerm}");
                _rapidApiService.FetchAndGenerateXmlAsync(searchTerm).Wait();
            }
            catch (Exception ex)
            {
                Console.WriteLine($"Error fetching data from API: {ex.Message}");
                return new List<Contact>();
            }

            return new List<Contact>
            {
                new Contact
                {
                    Domain = searchTerm,
                    Query = searchTerm
                }
            };
        }
    }
}