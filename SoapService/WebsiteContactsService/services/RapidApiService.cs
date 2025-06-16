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

        public RapidApiService(IHttpClientFactory httpClientFactory, IConfiguration configuration)
        {
            _httpClientFactory = httpClientFactory;
            _configuration = configuration;
        }

        public async Task<Contact> FetchAndGenerateXmlAsync(string domainToScrape)
        {
            Console.WriteLine($"Api call for: {domainToScrape}...");
            var client = _httpClientFactory.CreateClient();
            var apiKey = _configuration["RapidApi:ApiKey"];
            var apiHost = _configuration["RapidApi:Host"];

            if (string.IsNullOrEmpty(apiKey) || string.IsNullOrEmpty(apiHost))
            {
                Console.WriteLine("Error: RapidAPI Key or Host not configured.");
                return null;
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

                    return contact;
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
            return null;
        }
    }

    public class ApiResponseWrapper
    {
        public List<Contact> Data { get; set; }
    }
}


