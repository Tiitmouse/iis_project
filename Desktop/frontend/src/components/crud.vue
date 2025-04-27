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
    isLoggedIn.value = true;
  }
});

const onLoginSuccess = () => {
  isLoggedIn.value = true;
};
</script>

<style lang="css">
h2 {
  margin-bottom: 20px;
  color: var(--font-color);
}</style>