<!-- eslint-disable vue/singleline-html-element-content-newline -->
<!-- eslint-disable vue/html-self-closing -->
<!-- eslint-disable vue/max-attributes-per-line -->

<template>
  <v-card class="topCard">
    <v-card-title class="text-center justify-center py-6">
      <h1 class="font-weight-bold text-h2 text-basil">
        Interoperability of information systems
      </h1>
    </v-card-title>

    <v-tabs v-model="tab" bg-color="transparent" grow>
      <v-tab v-for="item in items" :key="item" :text="item" :value="item"></v-tab>
    </v-tabs>

    <v-tabs-window v-model="tab" class="tabWindow">
      <v-tabs-window-item v-for="item in items" :key="item" :value="item">
        <v-card flat>
          <v-card-text class="cardText">{{ text }}</v-card-text>
        </v-card>
      </v-tabs-window-item>
    </v-tabs-window>
  </v-card>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const tab = ref('Menus')

const items = [
  'Validation',
  'Soap',
  'JAXB',
  'Weather',
  'API con JWT'
]
const text = 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.'
</script>

<style>
.topCard {
  background-color: #5373b3 !important;
}

.tabWindow {
  background-color: white;
}
.cardText{
  color: white;
}
.text-basil {
  color: white !important;
}
</style>