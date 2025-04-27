<template>
  <div v-if="isLoggedIn">
    <h2>Entity manipulation</h2>
  </div>
  <Login v-else @login-success="onLoginSuccess" />
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import Login from './login.vue';

const isLoggedIn = ref(false);

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

<style lang="css">
h2 {
  margin-bottom: 20px;
  color: var(--font-color);
}</style>