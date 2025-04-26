
<template>
  <v-app>
    <v-app-bar app color="indigo" dark id="drag-area" class="top-bar">
      <v-toolbar-title style="margin-left: 90px; margin-top: -26px;font-size: 1em">INTEROPERABILITY OF INFORMATION SYSTEMS</v-toolbar-title>
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
         <v-card elevation="0">
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

const StartPage = defineAsyncComponent(() => import('../components/startpage.vue'))
const Validation = defineAsyncComponent(() => import('../components/validation.vue'))
const Crud = defineAsyncComponent(() => import('../components/crud.vue'))
const Weather = defineAsyncComponent(() => import('../components/weather.vue'))
const Soap = defineAsyncComponent(() => import('../components/soap.vue'))
const JAXB = defineAsyncComponent(() => import('../components/jaxb.vue'))

const tab = ref('StartPage') 
const drawer = ref(false) 

const items = [
  'StartPage',
  'Validation',
  'Soap',
  'JAXB',
  'Weather',
  'CRUD'
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
    case 'CRUD':
      return Crud
    default:
      return StartPage
  }
})

const setTab = (selectedTab: string) => {
  tab.value = selectedTab
  drawer.value = false 
}
</script>

<style>
html, body, #app, .v-application, .v-application__wrap, .v-main, .v-card {
  background-color: transparent !important;
}
#drag-area {
  -webkit-app-region: drag;
  cursor: move; 
  --wails-draggable: drag; 
}
#drag-area * {
  -webkit-user-select: none;
  cursor: default;
}
.cardText {
  color: black;
}
.top-bar{
  height: 36px;
}
.v-app-bar .v-btn {
  margin-top: -30px;
}
</style>