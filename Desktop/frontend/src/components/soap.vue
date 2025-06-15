<template>
    <v-container>
        <h2>Contact Search</h2>
        <v-row>
            <v-col cols="12">
                <v-text-field v-model="searchDomain" label="Search Domain" prepend-inner-icon="mdi-magnify"
                class="dashed-border" variant="plain" @keydown.enter="searchContacts"></v-text-field>
            </v-col>
        </v-row>
    </v-container>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useSnackbar } from '@/components/SnackbarProvider.vue';
import { SearchContacts } from "../../wailsjs/go/main/App";

const searchDomain = ref('');
const contactsData = ref<any[]>([]);
const snackbar = useSnackbar();

const searchContacts = async () => {
    if(searchDomain.value == null || searchDomain.value === ''){
        snackbar.Error("Please insert domain")
        return
    }

    try {
        const result = await SearchContacts(searchDomain.value);
        contactsData.value = result;
        if (result.length === 0) {
            console.log(result)
            snackbar.Info("No contacts found for this domain.");
        } else {
            console.log(result)
             snackbar.Success("Contacts found!");
        }
    } catch (error: any) {
        console.log(error)
        snackbar.Error(`Error searching contacts: ${error.message}`);
    }
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
</style>