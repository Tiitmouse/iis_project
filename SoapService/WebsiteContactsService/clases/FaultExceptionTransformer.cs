using System.ServiceModel.Channels;
using System.Xml;
using CoreWCF;
using SoapCore.Extensibility;

namespace WebsiteContactsService.clases;

public class FaultExceptionTransformer : IFaultExceptionTransformer
{
    public object TransformFaultException<TMessage>(System.ServiceModel.FaultException<TMessage> faultException)
    {
        return new { Message = faultException.Message };
    }

    public Message ProvideFault(Exception exception, MessageVersion messageVersion, Message requestMessage,
        XmlNamespaceManager xmlNamespaceManager)
    {
        var fault = Message.CreateMessage(messageVersion, "", new FaultReason(exception.Message));
        fault.Headers.Action = requestMessage.Headers.Action;
        return fault;
    }
}