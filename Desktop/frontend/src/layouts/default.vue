<template>
  <SnackbarProvider>
  <v-main>
    <router-view />
  </v-main>
<AppFooter />
</SnackbarProvider>

</template>

<script lang="ts" setup>
</script>

<style>
html,
body {
    overflow: hidden;
    height: 100vh;
}
body::-webkit-scrollbar {
    display: none;
}
</style>