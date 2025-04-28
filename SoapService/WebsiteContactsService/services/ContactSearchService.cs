using System.Xml.Linq;
using System.Xml.XPath;
using WebsiteContactsService.Contracts;

namespace WebsiteContactsService.Services
{
    public class ContactSearchService : IContactSearchService
    {
        private const string XmlFilePath = "contacts_data.xml";

        public List<ContactRecord> SearchContacts(string searchTerm)
        {
            Console.WriteLine($"SOAP SearchContacts called with term: '{searchTerm}'");
            var results = new List<ContactRecord>();

            if (!File.Exists(XmlFilePath))
            {
                Console.WriteLine($"Error: XML file not found at {XmlFilePath}");
                return results;
            }

            try
            {
                XDocument xDoc = XDocument.Load(XmlFilePath);
                
                string xpathQuery = $"//Contact/*[self::Emails/Email or self::PhoneNumbers/Phone][contains(Value, '{searchTerm}')]";

                Console.WriteLine($"Executing XPath: {xpathQuery}");
                var matchingNodes = xDoc.XPathSelectElements(xpathQuery);

                foreach (var node in matchingNodes) 
                {
                    var record = new ContactRecord
                    {
                        RecordType = node.Name.LocalName,
                        Value = node.Element("Value")?.Value,
                        Sources = node.Element("Sources")?.Elements("Source")
                                        .Select(src => src.Value)
                                        .ToList() ?? new List<string>()
                    };
                    results.Add(record);
                }
                Console.WriteLine($"Found {results.Count} matches.");
            }
            catch (XPathException e)
            {
                Console.WriteLine($"XPathException occurred: {e.Message}");
                Console.WriteLine($"Stack Trace: {e.StackTrace}");
                throw;
            }
            catch (Exception e)
            {
                Console.WriteLine($"An unexpected error occurred: {e.Message}");
                Console.WriteLine($"Stack Trace: {e.StackTrace}");
                throw;
            }

            return results;
        }
    }
}