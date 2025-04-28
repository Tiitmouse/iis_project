using System.Text.Json.Serialization;

namespace WebsiteContactsService.Models
{
    public class Contact
    {
        public string? Domain { get; set; }
        public string? Query { get; set; }
        public List<EmailEntry>? Emails { get; set; }
        [JsonPropertyName("phone_numbers")]
        public List<PhoneEntry>? PhoneNumbers { get; set; }
        public string? Facebook { get; set; }
        public string? Instagram { get; set; }
        public string? Tiktok { get; set; }
        public string? Snapchat { get; set; }
        public string? Twitter { get; set; }
        public string? Linkedin { get; set; }
        public string? Github { get; set; }
        public string? Youtube { get; set; }
        public string? Pinterest { get; set; }
    }

    public class EmailEntry
    {
        public string? Value { get; set; }
        public List<string>? Sources { get; set; }
    }

    public class PhoneEntry
    {
        public string? Value { get; set; }
        public List<string>? Sources { get; set; }
    }

    public class ApiResponseWrapper
    {
        public string? Status { get; set; }
        [JsonPropertyName("request_id")]
        public string? RequestId { get; set; }
        public List<Contact>? Data { get; set; }
    }
}