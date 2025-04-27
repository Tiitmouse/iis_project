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