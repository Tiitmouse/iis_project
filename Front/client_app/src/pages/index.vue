<!-- eslint-disable vue/max-attributes-per-line -->
<!-- eslint-disable vue/html-self-closing -->
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

    <v-tabs-window v-model="tab">
      <v-tabs-window-item v-for="item in items" :key="item" :value="item">
        <v-card flat class="whitebg">
          <v-card-text class="cardText">
            <component :is="currentComponent"></component>
          </v-card-text>
        </v-card>
      </v-tabs-window-item>
    </v-tabs-window>
  </v-card>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import Validation from '../components/validation.vue'
import ApiConJwt from '../components/apiConJwt.vue'
import Weather from '../components/weather.vue'
import Soap from '../components/soap.vue'
import JAXB from '../components/jaxb.vue'

const tab = ref('Menus')

const items = [
  'Validation',
  'Soap',
  'JAXB',
  'Weather',
  'API con JWT'
]

const currentComponent = computed(() => {
  switch (tab.value) {
    case 'Validation':
      return Validation
    case 'Soap':
      return Soap
    case 'JAXB':
      return JAXB
    case 'Weather':
      return Weather
    case 'API con JWT':
      return ApiConJwt
    default:
      return { template: '<div> default value </div>', props: ['text'] }
  }
})


</script>

<style>
.topCard {
  background-color: #5373b3 !important;
}
.whitebg {
  background-color: white !important;
}
.cardText{
  color: black;
}
.text-basil {
  color: white !important;
}
</style>