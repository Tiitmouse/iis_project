<template>
  <v-card>
    <v-card-title>
      <span class="text-h5">{{ isEditing ? 'Edit Contact' : 'Create New Contact' }}</span>
    </v-card-title>
    <v-card-text>
      <v-form ref="form" @submit.prevent="saveContact">
        <v-select
          v-model="contact.type"
          :items="contactTypes"
          label="Type"
          required
        ></v-select>
        <v-text-field v-model="contact.value" label="Value" required></v-text-field>
        <v-text-field v-model="contact.name" label="Name"></v-text-field>
        <v-textarea v-model="sourceInput" label="Sources (comma-separated)"></v-textarea>
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn color="red" @click="cancel">Cancel</v-btn>
      <v-btn color="white" @click="saveContact">Save</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import { defineComponent, ref, reactive, type PropType, computed } from 'vue';
import { useSnackbar } from '@/components/SnackbarProvider.vue';

export interface Contact {
  id?: string;
  type: string;
  value: string;
  name?: string;
  sources?: string[];
}

export default defineComponent({
  name: 'ContactFloat',
  props: {
    existingContact: {
      type: Object as PropType<Contact>,
      default: null,
    },
  },
  emits: ['save', 'cancel'],
  setup(props, { emit }) {
    const form = ref(null);
    const snackbar = useSnackbar();

    const contact = reactive<Contact>(props.existingContact ? { ...props.existingContact } : {
      type: '',
      value: '',
      name: '',
      sources: [],
    });

    const isEditing = computed(() => !!props.existingContact);

    const sourceInput = ref<string>('');

    if (props.existingContact && props.existingContact.sources) {
      sourceInput.value = props.existingContact.sources.join(', ');
    }

    const saveContact = async () => {
      if (!(form.value as any).validate()) {
        return;
      }

      if (!contact.type || !contact.value) {
        snackbar.Error('please insert all needed information.');
        return;
      }

      if (sourceInput.value) {
        contact.sources = sourceInput.value.split(',').map(s => s.trim());
      } else {
        contact.sources = [];
      }

      emit('save', contact);
    };

    const cancel = () => {
      emit('cancel');
    };

    const contactTypes = ref(['email', 'phone', 'social']);

    return {
      form,
      contact,
      saveContact,
      cancel,
      isEditing,
      sourceInput,
      contactTypes,
    };
  },
});
</script>

<style lang="css" scoped>
.v-card {
  background-color: #33466c !important;
  border: 2px dashed #697ea885;
}
</style>