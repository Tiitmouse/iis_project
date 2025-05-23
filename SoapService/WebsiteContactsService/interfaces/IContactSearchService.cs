using System.Runtime.Serialization;
using System.ServiceModel;

namespace WebsiteContactsService.Contracts
{
    [ServiceContract]
    public interface IContactSearchService
    {
        [OperationContract]
        List<ContactRecord> SearchContacts(string searchTerm);
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