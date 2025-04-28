import axios from "@/plugins/axios";

export interface Contact {
  id: string;
  type: 'email' | 'phone' | 'social';
  value: string;
  name?: string;
  sources?: string[];
}

export const fetchContacts = async (): Promise<Contact[]> => {
  const response = await axios.get('/api/contacts');
  return response.data;
};

export const deleteContact = async (id: string): Promise<void> => {
  await axios.delete(`/api/contacts/${id}`);
};

export const createContact = async (contact: Omit<Contact, 'id'>): Promise<Contact> => {
  const response = await axios.post('/api/contacts', contact);
  return response.data;
};

export const updateContact = async (id: string, contact: Contact): Promise<Contact> => {
  const response = await axios.put(`/api/contacts/${id}`, contact);
  return response.data;
};

export const fetchContact = async (id: string): Promise<Contact> => {
  const response = await axios.get(`/api/contacts/${id}`);
  return response.data;
};