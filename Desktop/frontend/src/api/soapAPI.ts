import { ManualSearch } from "../../wailsjs/go/main/App";
import { api } from "../../wailsjs/go/models";
import { useSnackbar } from '@/components/SnackbarProvider.vue';

const snackbar = useSnackbar();

// The function now returns a Promise resolving to an array of SoapContactRecord
export async function searchContactsByDomain(domain: string): Promise<api.SoapContactRecord[]> {
    if (!domain) {
        snackbar.Error("Please enter a domain to search.");
        return [];
    }

    try {
        // Call the new ManualSearch function
        const result = await ManualSearch(domain);
        if (!result || result.length === 0) {
            snackbar.Info("No contacts found for this domain.");
            return [];
        } else {
            snackbar.Success("Contacts found!");
            return result;
        }
    } catch (error: any) {
        snackbar.Error(`Error searching contacts: ${error.message || 'An unknown error occurred'}`);
        return [];
    }
}