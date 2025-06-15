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
                    return new List<Contact>();
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