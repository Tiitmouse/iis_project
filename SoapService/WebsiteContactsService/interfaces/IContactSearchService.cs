using System.Runtime.Serialization;
using System.ServiceModel;
using WebsiteContactsService.Models;

namespace WebsiteContactsService.Contracts
{
    [ServiceContract(Namespace = "http://tempuri.org/")]
    public interface IContactSearchService
    {
        [OperationContract]
        List<Contact> SearchContacts(string searchTerm);
    }

    [DataContract]
    public class ContactRecord
    {
        [DataMember]
        public string? RecordType { get; set; }

        [DataMember]
        public string? Value { get; set; }

        [DataMember]
        public List<string>? Sources { get; set; }
    }
}