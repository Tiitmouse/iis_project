import { FetchContacts, CreateContact, DeleteContact as WailsDeleteContact, UpdateContact as WailsUpdateContact, FetchContact as WailsFetchContactById } from '../../wailsjs/go/api/Secure'
import type { api } from "../../wailsjs/go/models";

export interface Contact {
  id: string;
  type: 'email' | 'phone' | 'social';
  value: string;
  name?: string;
  sources?: string[];
}

export const fetchContacts = async (): Promise<api.Contact[]> => {
  const rez = await FetchContacts();
  return rez; 
};

export const deleteContact = async (id: string): Promise<void> => {
  await WailsDeleteContact(id);
};

export const createContact = async (contactInfo: Omit<Contact, 'id'>): Promise<api.Contact> => {
  const contactToSend: api.Contact = {
    id: "",
    type: contactInfo.type,
    value: contactInfo.value,
    name: contactInfo.name || "",
    sources: contactInfo.sources || [],
  };
  const newContact = await CreateContact(contactToSend);
  return newContact;
};

export const updateContact = async (id: string, contact: Contact): Promise<Contact> => {
  const contactToSend: api.Contact = {
    id: contact.id,
    type: contact.type,
    value: contact.value,
    name: contact.name || "",
    sources: contact.sources || [],
  };
  const updatedContactFromWails = await WailsUpdateContact(id, contactToSend);
  return {
    id: updatedContactFromWails.id,
    type: updatedContactFromWails.type as 'email' | 'phone' | 'social',
    value: updatedContactFromWails.value,
    name: updatedContactFromWails.name,
    sources: updatedContactFromWails.sources,
  };
};

export const fetchContact = async (id: string): Promise<Contact> => {
  const contactFromWails = await WailsFetchContactById(id);
  return {
    id: contactFromWails.id,
    type: contactFromWails.type as 'email' | 'phone' | 'social',
    value: contactFromWails.value,
    name: contactFromWails.name,
    sources: contactFromWails.sources,
  };
};