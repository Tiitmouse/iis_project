na/Documents/algebra/projekti/iis_project/Front/client_app/src/components/login.vue
<template>
  <v-sheet class="pa-12" rounded>
    <v-card class="mx-auto px-6 py-8" max-width="344">
      <v-form v-model="form" @submit.prevent="onSubmit">
        <v-text-field v-model="username" :readonly="loading" :rules="[required]" class="mb-2" label="Username"
          clearable></v-text-field>

        <v-text-field type="password" v-model="password" :readonly="loading" :rules="[required]" label="Password"
          placeholder="Enter your password" clearable></v-text-field>

        <br>

        <v-btn :disabled="!form" :loading="loading" color="success" size="large" type="submit" variant="elevated"
          block>
          Sign In
        </v-btn>
      </v-form>
    </v-card>
  </v-sheet>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { login } from '@/api/loginAPI';

const form = ref(false);
const username = ref('');
const password = ref('');
const loading = ref(false);
const emit = defineEmits(['login-success']);

async function onSubmit() {
  if (!form.value) return;

  loading.value = true;

  try {
    const { access_token, refresh_token } = await login(username.value, password.value);

    localStorage.setItem('accessToken', access_token);
    localStorage.setItem('refreshToken', refresh_token);
    emit('login-success');

  } catch (error: any) {
    console.error('Login failed:', error.message);
  } finally {
    loading.value = false;
  }
}

function required(v: any) {
  return !!v || 'Field is required';
}
</script>