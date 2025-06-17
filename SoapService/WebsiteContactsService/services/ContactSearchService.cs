using System.Xml.Linq;
using WebsiteContactsService.Contracts;
using WebsiteContactsService.Models;
using System.Xml.XPath; // Added for XPath support

namespace WebsiteContactsService.Services
{
    
    public class ContactSearchService : IContactSearchService
    {
        private const string XmlFilePath = "contacts_data.xml";
        private readonly RapidApiService _rapidApiService;
        private Contact _contactData;
        private List<Contact> _contactList;

        public ContactSearchService(RapidApiService rapidApiService)
        {
            _rapidApiService = rapidApiService;
        }

        public async Task<List<Contact>> SearchContacts(string searchTerm)
        {

            var contacts = LoadFromFile(searchTerm);
            if (contacts.Any())
            {
                Console.WriteLine($"Contact with domain '{searchTerm}' found in file. Returning existing contact.");
                foreach (var contact in contacts)
                {
                    Console.WriteLine($"Contact Domain: {contact.Domain}, results: {contact.Emails?.Count ?? 0} emails, {contact.PhoneNumbers?.Count ?? 0} phone numbers");
                }
                return contacts;
            }
            
            Console.WriteLine($"Contact with domain '{searchTerm}' not found in file. Fetching from API...");
            
            try
            {
                Console.WriteLine($"Fetching data from api for search term: {searchTerm}");
                _contactData = await _rapidApiService.FetchAndGenerateXmlAsync(searchTerm);
                _contactList = _contactData != null ? new List<Contact> { _contactData } : new List<Contact>();
                GenerateXml(_contactData);
            }
            catch (Exception ex)
            {
                Console.WriteLine($"Error fetching data from API: {ex.Message}");
                return new List<Contact>();
            }
            
            Console.WriteLine($"Contact data fetched successfully for search term: {searchTerm}");
            
            foreach (var contact in _contactList)
            {
                Console.WriteLine($"Contact Domain: {contact.Domain}, results: {contact.Emails?.Count ?? 0} emails, {contact.PhoneNumbers?.Count ?? 0} phone numbers");
            }
            
            return _contactList;
        }

        private List<Contact> LoadFromFile(string searchTerm)
        {
            if (File.Exists(XmlFilePath))
            {
                var existingDoc = XDocument.Load(XmlFilePath);
                var dataElement = existingDoc.XPathSelectElement($"/root/data[domain='{searchTerm}']");

                if (dataElement == null)
                {
                    Console.WriteLine($"No contact found with the domain '{searchTerm}' in the XML file using XPath.");
                    return new List<Contact>();
                }

                var contact = new Contact
                {
                    Domain = dataElement.Element("domain")?.Value,
                    Query = dataElement.Element("query")?.Value,
                    Emails = dataElement.Elements("emails").Select(e => new EmailEntry
                    {
                        Value = e.Element("value")?.Value,
                        Sources = e.Elements("sources").Select(s => s.Value).ToList()
                    }).ToList(),
                    PhoneNumbers = dataElement.Elements("phone_numbers").Select(p => new PhoneEntry
                    {
                        Value = p.Element("value")?.Value,
                        Sources = p.Elements("sources").Select(s => s.Value).ToList()
                    }).ToList(),
                    Facebook = dataElement.Element("facebook")?.Value,
                    Instagram = dataElement.Element("instagram")?.Value
                };

                Console.WriteLine($"Contact with domain '{searchTerm}' found in file using XPath.");
                return new List<Contact> { contact };
            }

            Console.WriteLine("XML file not found.");
            return new List<Contact>();
        }
        
        private void GenerateXml(Contact contactData)
        {
            if (contactData == null)
            {
                var emptyDoc = new XDocument(
                    new XDeclaration("1.0", "UTF-8", null),
                    new XElement("root",
                        new XElement("status", "OK"),
                        new XElement("request_id", Guid.NewGuid().ToString()),
                        new XElement("data")
                    )
                );
                emptyDoc.Save(XmlFilePath);
                Console.WriteLine($"Empty XML data file created at {XmlFilePath}");
                return;
            }

            var dataContent = new List<XObject>();
            dataContent.Add(new XElement("domain", contactData.Domain));
            dataContent.Add(new XElement("query", contactData.Query));

            if (contactData.Emails != null && contactData.Emails.Any())
            {
                dataContent.AddRange(contactData.Emails.Select(email =>
                    new XElement("emails",
                        new XElement("value", email.Value),
                        email.Sources?.Select(source => new XElement("sources", source))
                    )
                ));
            }

            if (contactData.PhoneNumbers != null && contactData.PhoneNumbers.Any())
            {
                dataContent.AddRange(contactData.PhoneNumbers.Select(phone =>
                    new XElement("phone_numbers",
                        new XElement("value", phone.Value),
                        phone.Sources?.Select(source => new XElement("sources", source))
                    )
                ));
            }
            
            if (contactData.Facebook != null) dataContent.Add(new XElement("facebook", contactData.Facebook));
            if (contactData.Instagram != null) dataContent.Add(new XElement("instagram", contactData.Instagram));
            if (contactData.Tiktok != null) dataContent.Add(new XElement("tiktok", contactData.Tiktok));
            if (contactData.Snapchat != null) dataContent.Add(new XElement("snapchat", contactData.Snapchat));
            if (contactData.Twitter != null) dataContent.Add(new XElement("twitter", contactData.Twitter));
            if (contactData.Linkedin != null) dataContent.Add(new XElement("linkedin", contactData.Linkedin));
            if (contactData.Github != null) dataContent.Add(new XElement("github", contactData.Github));
            if (contactData.Youtube != null) dataContent.Add(new XElement("youtube", contactData.Youtube));
            if (contactData.Pinterest != null) dataContent.Add(new XElement("pinterest", contactData.Pinterest));

            var xDoc = new XDocument(
                new XDeclaration("1.0", "UTF-8", null),
                new XElement("root",
                    new XElement("status", "OK"),
                    new XElement("request_id", Guid.NewGuid().ToString()),
                    new XElement("data", dataContent)
                )
            );

            xDoc.Save(XmlFilePath);
            Console.WriteLine($"XML data saved to {XmlFilePath}");
        }
    }
}