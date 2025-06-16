import { ManualSearch } from "../../wailsjs/go/main/App";
import { api } from "../../wailsjs/go/models";
import { useSnackbar } from '@/components/SnackbarProvider.vue';

const snackbar = useSnackbar();

export async function searchContactsByDomain(domain: string): Promise<api.SoapContactRecord[]> {
    console.log("IN SEARCHBYDOMAIN....................")
    if (!domain) {
        console.log("Please enter a domain to search.");
        return [];
    }

    try {
        const result = await ManualSearch(domain);
        console.log("!! RESULT: ", result);
        if (!result || result.length === 0) {
            console.log("No contacts found for this domain.");
            return [];
        } else {
            console.log("Contacts found!");
            return result;
        }
    } catch (error: any) {
        console.log("Error searching contacts:", error.message);
        return [];
    }
}

