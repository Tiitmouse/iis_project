using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Text.Json;
using System.Threading.Tasks;
using System.Xml.Linq;
using WebsiteContactsService.Models;
using Microsoft.Extensions.Configuration;

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
                    var options = new JsonSerializerOptions { PropertyNameCaseInsensitive = true };
                    var apiResponse = JsonSerializer.Deserialize<ApiResponseWrapper>(body, options);
                    var contact = apiResponse?.Data?.FirstOrDefault();

                    if (contact != null)
                    {
                        GenerateXml(contact);
                    }
                    else
                    {
                        GenerateXml(null);
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

    public class ApiResponseWrapper
    {
        public List<Contact> Data { get; set; }
    }
}
