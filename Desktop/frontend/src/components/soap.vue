<template>
  <v-container>
    <h2>Contacts Search</h2>
    <v-row>
      <v-col cols="12">
        <v-text-field
          v-model="searchDomain"
          label="Search Domain"
          style="color: white;"
          outlined
          dense
          prepend-inner-icon="mdi-magnify"
          class="dashed-border"
          variant="plain"
          @keyup.enter="handleSearch"
        ></v-text-field>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12">
        <v-data-table
          :headers="tableHeaders"
          :items="tableRows"
          class="coloring dashed-border"
          fixed-header
          height="400"
          :items-per-page="-1"
          hide-default-footer
          style=" padding: 10px 10px 0 10px;"
        >
          <template v-slot:item.sources="{ item }">
            <ul v-if="item.sources && item.sources.length" class="source-list">
              <li v-for="(src, i) in item.sources" :key="i">
                <a :href="src" target="_blank" class="source-link">{{ src }}</a>
              </li>
            </ul>
          </template>
        </v-data-table>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts" setup>
import { ref, computed, toRaw } from 'vue';
import { searchContactsByDomain } from '@/api/soapAPI';
import { useSnackbar } from '@/components/SnackbarProvider.vue';
import { api } from '../../wailsjs/go/models';

const searchDomain = ref('');
const contactsData = ref<api.SoapContactRecord[]>([]);
const snackbar = useSnackbar();

const tableHeaders = [
  { title: 'Type', key: 'type' },
  { title: 'Value', key: 'value' },
  { title: 'Sources', key: 'sources' },
];

const tableRows = computed(() => {
  const rows: { type: string; value: string; sources?: string[] }[] = [];
  const seen = new Set<string>();
  toRaw(contactsData.value).forEach(contact => {
    (contact.Emails || []).forEach(email => {
      if (email.Value && !seen.has('email|' + email.Value)) {
        seen.add('email|' + email.Value);
        rows.push({ type: 'email', value: email.Value, sources: email.Sources });
      }
    });
    (contact.PhoneNumbers || []).forEach(phone => {
      if (phone.Value && !seen.has('phone|' + phone.Value)) {
        seen.add('phone|' + phone.Value);
        rows.push({ type: 'phone', value: phone.Value, sources: phone.Sources });
      }
    });
    [
      contact.Facebook,
      contact.Instagram,
      contact.Github,
      contact.Linkedin,
      contact.Twitter,
      contact.Youtube,
      contact.Pinterest,
      contact.Tiktok,
      contact.Snapchat
    ].forEach(social => {
      if (social && !seen.has('social|' + social)) {
        seen.add('social|' + social);
        rows.push({ type: 'social', value: social });
      }
    });
  });
  return rows;
});

const handleSearch = async () => {
    try {
        if (!searchDomain.value) {
            snackbar.Error("Please enter a domain to search.");
            contactsData.value = [];
            return;
        }
        const result = await searchContactsByDomain(searchDomain.value.toLowerCase());
        if (result && result.length > 0) {
            snackbar.Success("Contacts found!");
            contactsData.value = result;
        } else {
            snackbar.Info("No contacts found for this domain.");
            contactsData.value = [];
        }
        console.log('contactsData', contactsData.value);
        console.log('tableRows', tableRows.value);
    } catch (error: any) {
        snackbar.Error(`Error searching contacts: ${error.message || 'An unknown error occurred'}`);
        contactsData.value = [];
    }
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
h4 {
    color: var(--font-color-secondary);
    font-weight: 500;
}
.entry-item {
    border: 1px solid #7e95c785;
    border-radius: 4px;
    background-color: #5373b321;
}
.source-list {
    list-style-type: disc;
    padding-left: 25px;
    margin-top: 5px;
}
.source-link {
    color: #b3c5f1;
    text-decoration: none;
}
.source-link:hover {
    text-decoration: underline;
}
.social-link-item {
    padding: 2px 0;
}
.contacts-scrollable {
  max-height: 500px;
  overflow-y: auto;
}
</style>