na/Documents/algebra/projekti/iis_project/Front/client_app/src/components/login.vue
<template>
  <v-sheet class="pa-12" color="transparent">
    <v-card class="mx-auto px-6 py-8 dashed-border" max-width="344" elevation="0">
      <v-form v-model="form" @submit.prevent="onSubmit">
        <v-text-field v-model="username" :readonly="loading" :rules="[required]" label="Username"></v-text-field>

        <v-text-field type="password" v-model="password" :readonly="loading" :rules="[required]" label="Password"
          placeholder="Enter your password"></v-text-field>

        <br>

        <v-btn :disabled="!form" :loading="loading" class="btnLogin" size="large" type="submit" variant="elevated"
          block>
          Login
        </v-btn>
      </v-form>
    </v-card>
  </v-sheet>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { login } from '@/api/loginAPI';
import { useSnackbar } from '@/components/SnackbarProvider.vue';

const snackbar = useSnackbar();

const form = ref(false);
const username = ref('');
const password = ref('');
const loading = ref(false);
const emit = defineEmits(['login-success']);

async function onSubmit() {
  if (!form.value) return;

  loading.value = true;

  try {
     await login(username.value, password.value);
    emit('login-success');
    snackbar.Success("Login successful!");

  } catch (error: any) {
    console.error('Login failed:', error.message);
    snackbar.Error(`Login failed! ${error.message}`);
  } finally {
    loading.value = false;
  }
}

function required(v: any) {
  return !!v || 'Field is required';
}
</script>

<style lang="css">
.btnLogin {
  background-color: var(--button-color) !important;
  color: var(--font-color);
  font-weight: bold;
  margin-top: 20px;
}

.dashed-border {
  border: 2px dashed #697ea885;
  background-color: #5373b321 !important;
  border-radius: 5px;
}
</style>