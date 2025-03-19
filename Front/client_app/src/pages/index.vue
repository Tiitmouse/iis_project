<!-- eslint-disable vue/max-attributes-per-line -->
<!-- eslint-disable vue/html-self-closing -->

<template>
  <v-app>
    <v-app-bar app color="indigo" dark>
      <v-toolbar-title>My App</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn v-for="item in items" :key="item" text="true" @click="setTab(item)">
        {{ item }}
      </v-btn>
    </v-app-bar>

    <v-navigation-drawer v-model="drawer" app temporary>
      <v-list>
        <v-list-item-group v-for="item in items" :key="item">
          <v-list-item @click="setTab(item)">
            <v-list-item-content>
              <v-list-item-title>{{ item }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-navigation-drawer>

    <v-main>
      <v-container>
        <v-card class="whitebg">
          <v-card-text class="cardText">
            <component :is="currentComponent" v-if="tab !== ''"></component>
          </v-card-text>
        </v-card>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import { ref, computed, defineAsyncComponent } from 'vue'

const Validation = defineAsyncComponent(() => import('../components/validation.vue'))
const ApiConJwt = defineAsyncComponent(() => import('../components/apiConJwt.vue'))
const Weather = defineAsyncComponent(() => import('../components/weather.vue'))
const Soap = defineAsyncComponent(() => import('../components/soap.vue'))
const JAXB = defineAsyncComponent(() => import('../components/jaxb.vue'))

const tab = ref('') 
const drawer = ref(false) 

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
      return { template: '<div> Select a tab </div>' }
  }
})

const setTab = (selectedTab: string) => {
  tab.value = selectedTab
  drawer.value = false 
}
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
.v-main{
  background: white;
}
</style>