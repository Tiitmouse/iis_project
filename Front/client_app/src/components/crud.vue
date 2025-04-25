<template>
    <div v-if="isLoggedIn">
      <h3>Entity manipulation</h3>
      <p>its pointless i know </p>
    </div>
    <Login v-else @login-success="onLoginSuccess" />
  </template>
  
  <script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import Login from './login.vue';
  
  const isLoggedIn = ref(false);
  
  onMounted(() => {
    if (localStorage.getItem('accessToken')) {
      isLoggedIn.value = true;
    }
  });
  
  const onLoginSuccess = () => {
    isLoggedIn.value = true;
  };
  </script>