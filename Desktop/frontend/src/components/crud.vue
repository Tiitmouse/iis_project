<template>
  <div v-if="isLoggedIn">
    <h2>Entity manipulation</h2>
    <v-data-table-server
      :headers="headers"
      :items="serverItems"
      :items-length="serverItems.length"
      :loading="loading"
      item-value="id"
      @update:options="loadItems"
      fixed-header  
      height="500"  
    >
      <template v-slot:top>
        <v-toolbar flat color="transparent">
          <v-toolbar-title>Contacts</v-toolbar-title>
          <v-divider class="mx-4" inset vertical></v-divider>
          <v-spacer></v-spacer>
          <v-btn color="primary" dark class="mb-2">New Contact</v-btn>
        </v-toolbar>
      </template>
       <template v-slot:item.actions="{ item }">
        <v-icon small class="mr-2" @click="editItem(item)">mdi-pencil</v-icon>
        <v-icon small @click="deleteItem(item)">mdi-delete</v-icon>
      </template>
      <template v-slot:bottom></template>
    </v-data-table-server>
  </div>
  <Login v-else @login-success="onLoginSuccess" />
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import Login from './login.vue';
import { fetchContacts, type Contact } from '@/api/contactsAPI';
import { useSnackbar } from '@/components/SnackbarProvider.vue';

const isLoggedIn = ref(false);
const snackbar = useSnackbar();

const headers = ref([
  { title: 'Type', key: 'type' },
  { title: 'Value', key: 'value' },
  { title: 'Name', key: 'name' },
  { title: 'Actions', key: 'actions', sortable: false },
]);
const serverItems = ref<Contact[]>([]);
const loading = ref(true);


const loadItems = async ({ sortBy }: { sortBy: any }) => {
  loading.value = true;
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
    console.error('Failed to load contacts:', error);
    snackbar.Error(`Failed to load contacts: ${error.message || 'Unknown error'}`);
    serverItems.value = [];
  } finally {
    loading.value = false;
  }
};

const editItem = (item: Contact) => {
  console.log('Edit item:', item);
};

const deleteItem = (item: Contact) => {
  console.log('Delete item:', item);
};


onMounted(() => {
  if (localStorage.getItem('accessToken')) {
    console.log('User already logged in');
    isLoggedIn.value = true;
  }
});

const onLoginSuccess = () => {
  console.log('Login successful!');
  isLoggedIn.value = true;
};
</script>

<style lang="css" scoped>
h2 {
  margin-bottom: 20px;
  color: var(--font-color);
}
</style>