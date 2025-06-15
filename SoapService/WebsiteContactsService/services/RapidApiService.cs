using System.Text.Json;
using System.Xml.Linq;
using WebsiteContactsService.Models;

namespace WebsiteContactsService.Services
{
    public class RapidApiService
    {
        private readonly IHttpClientFactory _httpClientFactory;
        private readonly IConfiguration _configuration;
        private const string XmlFilePath = "contacts_data.xml";

        public RapidApiService(IHttpClientFactory httpClientFactory, IConfiguration configuration)
        {
            _httpClientFactory = httpClientFactory;
            _configuration = configuration;
        }

public async Task FetchAndGenerateXmlAsync(string domainToScrape)
{
    Console.WriteLine($"Fetching data for domain: {domainToScrape}...");
    var client = _httpClientFactory.CreateClient();
    var apiKey = _configuration["RapidApi:ApiKey"];
    var apiHost = _configuration["RapidApi:Host"]; 

    if (string.IsNullOrEmpty(apiKey) || string.IsNullOrEmpty(apiHost))
    {
        Console.WriteLine("Error: RapidAPI Key or Host not configured.");
        return;
    }
    
    var query = Uri.EscapeDataString(domainToScrape);
    var matchEmailDomain = "false";
    var externalMatching = "false";
    var apiUrl = $"https://{apiHost}/scrape-contacts?query={query}&match_email_domain={matchEmailDomain}&external_matching={externalMatching}";

    Console.WriteLine($"Constructed API URL: {apiUrl}");

    var request = new HttpRequestMessage
    {
        Method = HttpMethod.Get,
        RequestUri = new Uri(apiUrl),
        Headers =
        {
            { "x-rapidapi-key", apiKey }, 
            { "x-rapidapi-host", apiHost },
        },
    };

    try
    {
        using (var response = await client.SendAsync(request))
        {
            response.EnsureSuccessStatusCode();
            var body = await response.Content.ReadAsStringAsync();
            Console.WriteLine($"API Response: {body}"); // Log the raw API response

            var options = new JsonSerializerOptions { PropertyNameCaseInsensitive = true };
            var apiResponse = JsonSerializer.Deserialize<ApiResponseWrapper>(body, options);

            if (apiResponse?.Data != null && apiResponse.Data.Count > 0)
            {
                Console.WriteLine($"Deserialization successful ({apiResponse.Data.Count} records). Generating XML...");
                GenerateXml(apiResponse.Data);
                Console.WriteLine($"XML data saved to {XmlFilePath}");
            }
            else
            {
                Console.WriteLine("Failed to deserialize API response or data list is null/empty.");
                GenerateXml(new List<Contact>());
                Console.WriteLine($"Empty XML data file created at {XmlFilePath}");
            }
        }
    }
    catch (HttpRequestException e)
    {
        Console.WriteLine($"Error fetching data from RapidAPI ({e.StatusCode}): {e.Message}");
    }
    catch (JsonException e)
    {
        Console.WriteLine($"Error deserializing API response: {e.Message}");
    }
    catch (Exception e) 
    {
        Console.WriteLine($"An unexpected error occurred: {e.Message}");
    }
}

private void GenerateXml(List<Contact> contactDataList)
{
    var xDoc = new XDocument(
        new XElement("Contacts",
            contactDataList.Select(contactData =>
                new XElement("Contact",
                    new XElement("Domain", contactData.Domain),
                    new XElement("Query", contactData.Query),
                    new XElement("Emails",
                        contactData.Emails?.Select(e =>
                            new XElement("Email",
                                new XElement("Value", e.Value),
                                new XElement("Sources",
                                    e.Sources?.Select(s => new XElement("Source", s)) ?? Enumerable.Empty<XElement>()
                                )
                            )
                        ) ?? Enumerable.Empty<XElement>()
                    ),
                    new XElement("PhoneNumbers",
                        contactData.PhoneNumbers?.Select(p =>
                            new XElement("Phone",
                                new XElement("Value", p.Value),
                                new XElement("Sources",
                                    p.Sources?.Select(s => new XElement("Source", s)) ?? Enumerable.Empty<XElement>()
                                )
                            )
                        ) ?? Enumerable.Empty<XElement>()
                    ),
                    contactData.Facebook == null ? null : new XElement("Facebook", contactData.Facebook),
                    contactData.Instagram == null ? null : new XElement("Instagram", contactData.Instagram),
                    contactData.Tiktok == null ? null : new XElement("Tiktok", contactData.Tiktok),
                    contactData.Snapchat == null ? null : new XElement("Snapchat", contactData.Snapchat),
                    contactData.Twitter == null ? null : new XElement("Twitter", contactData.Twitter),
                    contactData.Linkedin == null ? null : new XElement("Linkedin", contactData.Linkedin),
                    contactData.Github == null ? null : new XElement("Github", contactData.Github),
                    contactData.Youtube == null ? null : new XElement("Youtube", contactData.Youtube),
                    contactData.Pinterest == null ? null : new XElement("Pinterest", contactData.Pinterest)
                )
            )
        )
    );

    xDoc.Save(XmlFilePath);
    Console.WriteLine($"XML data saved to {XmlFilePath}");
}
    }
}