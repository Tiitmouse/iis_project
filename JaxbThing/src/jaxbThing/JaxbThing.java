package jaxbThing;

import generated.Root;
import java.io.File;
import java.net.MalformedURLException;
import java.net.URL;
import java.util.logging.Level;
import java.util.logging.Logger;
import javax.xml.XMLConstants;
import javax.xml.bind.JAXBContext;
import javax.xml.bind.JAXBException;
import javax.xml.bind.Unmarshaller;
import javax.xml.validation.Schema;
import javax.xml.validation.SchemaFactory;
import org.xml.sax.SAXException;

public class JaxbThing {

    public static void main(String[] args) {
        try {
            JAXBContext jc = JAXBContext.newInstance(Root.class);

            SchemaFactory sf = SchemaFactory.newInstance(XMLConstants.W3C_XML_SCHEMA_NS_URI);
            URL schemaURL = new File("xml-resources/jaxb/Contact/getXSDschema.xsd").toURI().toURL();
            Schema schema = sf.newSchema(schemaURL);

            Unmarshaller unmarshaller = jc.createUnmarshaller();
            unmarshaller.setSchema(schema);
            unmarshaller.setEventHandler(event -> {
                if (event.getSeverity() >= javax.xml.bind.ValidationEvent.WARNING) {
                    System.out.println("\nVALIDATION EVENT:");
                    System.out.println("SEVERITY:  " + event.getSeverity());
                    System.out.println("MESSAGE:  " + event.getMessage());
                }
                return true;
            });

            File xmlFileToValidate = new File("../SoapService/WebsiteContactsService/contacts_data.xml");
            Root rootElement = (Root) unmarshaller.unmarshal(xmlFileToValidate);
            
            System.out.println("\nUnmarshalling successful. Root element data:");
            System.out.println("\tStatus: " + rootElement.getStatus());
            System.out.println("\tRequest ID: " + rootElement.getRequestId());
            System.out.println("\tDomain: " + rootElement.getData().getDomain());

        } catch (JAXBException | SAXException | MalformedURLException ex) {
            Logger.getLogger(JaxbThing.class.getName()).log(Level.SEVERE, "XML processing failed", ex);
        }
    }
}