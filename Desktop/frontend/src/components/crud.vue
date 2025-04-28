<template>
  <div v-if="isLoggedIn">
    <div style="display: flex; align-items: center; justify-content: space-between;">
      <h2>Contacts</h2>
      <v-icon big color="error" @click="logout">mdi-logout</v-icon>
    </div>
    <v-data-table-server :headers="headers" :items="serverItems" :items-length="serverItems.length" :loading="loading"
      item-value="id" @update:options="loadItems" fixed-header height="500" class="coloring dashed-border">
      <template v-slot:top>
        <v-toolbar flat color="transparent">
          <v-toolbar-title>Contacts</v-toolbar-title>
          <v-divider class="mx-4" inset vertical></v-divider>
          <v-spacer></v-spacer>
          <v-btn color="primary" dark class="mb-2" @click="openNewContactDialog">New Contact</v-btn>
        </v-toolbar>
      </template>
      <template v-slot:item.actions="{ item }">
        <v-icon small class="mr-2" @click="editItem(item)">mdi-pencil</v-icon>
        <v-icon small @click="deleteItem(item)">mdi-delete</v-icon>
      </template>
      <template v-slot:bottom></template>
    </v-data-table-server>

    <v-dialog v-model="newContactDialog" max-width="500px">
      <ContactFloat @save="saveNewContact" @cancel="closeNewContactDialog" />
    </v-dialog>

    <v-dialog v-model="editContactDialog" max-width="500px">
      <ContactFloat :existingContact="selectedContact || undefined" @save="saveEditedContact"
        @cancel="closeEditContactDialog" />
    </v-dialog>
  </div>
  <Login v-else @login-success="onLoginSuccess" />
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import Login from './login.vue';
import { fetchContacts, deleteContact, createContact, updateContact, type Contact} from '@/api/contactAPI';
import { useSnackbar } from '@/components/SnackbarProvider.vue';
import ContactFloat from './contactFloat.vue';

const isLoggedIn = ref(false);
const snackbar = useSnackbar();
const serverItems = ref<Contact[]>([]);
const loading = ref(true);
const currentSortBy = ref<any[]>([]);
const router = useRouter();

const headers = ref([
  { title: 'Type', key: 'type' },
  { title: 'Value', key: 'value' },
  { title: 'Actions', key: 'actions', sortable: false },
]);

const newContactDialog = ref(false);
const editContactDialog = ref(false);
const selectedContact = ref<Contact | null>(null);

const loadItems = async ({ sortBy }: { sortBy: any }) => {
  loading.value = true;
  currentSortBy.value = sortBy;
  try {
    const allContacts = await fetchContacts();
    serverItems.value = allContacts;

    if (sortBy.length > 0) {
      const sortKey = sortBy[0].key;
      const sortOrder = sortBy[0].order;
      serverItems.value.sort((a: any, b: any) => {
        const valA = a[sortKey];
        const valB = b[sortKey];
        if (valA == null && valB == null) return 0;
        if (valA == null) return sortOrder === 'asc' ? -1 : 1;
        if (valB == null) return sortOrder === 'asc' ? 1 : -1;

        if (valA < valB) return sortOrder === 'asc' ? -1 : 1;
        if (valA > valB) return sortOrder === 'asc' ? 1 : -1;
        return 0;
      });
    }
  } catch (error: any) {
    snackbar.Error(`Failed to load contacts`);
    serverItems.value = [];
  } finally {
    loading.value = false;
  }
};

const logout = () => {
  localStorage.removeItem('accessToken');
  isLoggedIn.value = false;
  router.push('/');
  snackbar.Success('Logged out');
};

const editItem = async (item: Contact) => {
  selectedContact.value = item;
  editContactDialog.value = true;
};

const deleteItem = async (item: Contact) => {
  console.log('Deleting item:', item);
  try {
    await deleteContact(item.id);
    snackbar.Success('Contact deleted successfully');
    await loadItems({ sortBy: currentSortBy.value });
  } catch (error: any) {
    console.error('Failed to delete contact:', error);
    snackbar.Error(`Failed to delete contact: ${error}`);
  }
};

const openNewContactDialog = () => {
  newContactDialog.value = true;
};

const closeNewContactDialog = () => {
  newContactDialog.value = false;
};

const closeEditContactDialog = () => {
  editContactDialog.value = false;
  selectedContact.value = null;
};

const saveNewContact = async (contact: Contact) => {
  console.log('Saving new contact:', contact);
  try {
    await createContact(contact);
    snackbar.Success('Contact created successfully');
    closeNewContactDialog();
    await loadItems({ sortBy: currentSortBy.value });
  } catch (error: any) {
    console.error('Failed to create contact:', error);
    snackbar.Error(`Failed to create contact: ${error}`);
  }
};

const saveEditedContact = async (contact: Contact) => {
  console.log('Saving edited contact:', contact);
  try {
    await updateContact(contact.id, contact);
    snackbar.Success('Contact updated successfully');
    closeEditContactDialog();
    await loadItems({ sortBy: currentSortBy.value });
  } catch (error: any) {
    console.error('Failed to update contact:', error);
    snackbar.Error(`Failed to update contact: ${error}`);
  }
};

onMounted(() => {
  if (localStorage.getItem('accessToken')) {
    console.log('User already logged in');
    isLoggedIn.value = true;
  }
});

const onLoginSuccess = () => {
  snackbar.Success(`Login successful!`);
  console.log('Login successful!');
  isLoggedIn.value = true;
};
</script>

<style lang="css" scoped>
h2 {
  margin-bottom: 20px;
  color: var(--font-color);
}

.coloring {
  background-color: #5373b348;
  border-radius: 5px;
}

.coloring :deep(.v-data-table-fixed__header > table > thead th),
.coloring :deep(.v-table__wrapper > table > thead th) {
  border-radius: 5px;
  background-color: #3f51b5 !important;
  color: var(--font-color);
}

.dashed-border {
  border: 2px dashed #697ea885;
  padding-left: 10px;
  border-radius: 5px;
}
</style>