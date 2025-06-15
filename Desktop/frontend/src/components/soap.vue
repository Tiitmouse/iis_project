<template>
    <v-container>
        <h2>Contact Search (Manual Method)</h2>
        <v-row>
            <v-col cols="12">
                <v-text-field
                    v-model="searchDomain"
                    label="Search Domain"
                    prepend-inner-icon="mdi-magnify"
                    class="dashed-border"
                    variant="plain"
                    @keydown.enter="handleSearch"
                ></v-text-field>
            </v-col>
        </v-row>
        <v-row v-if="contactsData.length > 0">
            <v-col v-for="(contact, index) in contactsData" :key="index" cols="12" md="6">
                <v-card class="contact-card">
                    <v-card-title>
                        {{ contact.Value }}
                    </v-card-title>
                    <v-card-subtitle>{{ contact.RecordType }}</v-card-subtitle>
                    <v-card-text v-if="contact.Sources && contact.Sources.length > 0">
                        <strong>Sources:</strong>
                        <ul>
                            <li v-for="(source, s) in contact.Sources" :key="s">{{ source }}</li>
                        </ul>
                    </v-card-text>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { searchContactsByDomain } from '@/api/soapAPI';
import { api } from '../../wailsjs/go/models';

const searchDomain = ref('');
const contactsData = ref<api.SoapContactRecord[]>([]);

const handleSearch = async () => {
    const result = await searchContactsByDomain(searchDomain.value);
    contactsData.value = result;
};
</script>

<style lang="css" scoped>
h2 {
    margin-bottom: 20px;
    color: var(--font-color);
}
.dashed-border {
    border: 2px dashed #697ea885;
    background-color: #5373b321 !important;
    max-height: 60px;
    padding-left: 10px;
    border-radius: 5px;
}
.contact-card {
    background-color: #5373b348;
    border-radius: 5px;
    border: 2px dashed #697ea885;
    margin-top: 10px;
    padding: 10px;
    min-height: 150px;
}
</style>