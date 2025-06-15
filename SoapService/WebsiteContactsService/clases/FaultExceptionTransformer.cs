using System.ServiceModel.Channels;
using System.Xml;
using SoapCore.Extensibility;

namespace WebsiteContactsService.clases;

public class FaultExceptionTransformer : IFaultExceptionTransformer
{
    public object TransformFaultException<TMessage>(System.ServiceModel.FaultException<TMessage> faultException)
    {
        return new { Message = faultException.Message };
    }

    public Message ProvideFault(Exception exception, MessageVersion messageVersion, Message requestMessage, XmlNamespaceManager xmlNamespaceManager)
    {
            Console.WriteLine($"Fault Exception: {exception.Message}");
        return null;
    }
}