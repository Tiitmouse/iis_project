<template>
  <v-container>
    <h2>Weather check</h2>
    <v-row>
      <v-col cols="12">
        <v-text-field v-model="searchCity" label="Search City" prepend-inner-icon="mdi-magnify" class="dashed-border" variant="plain"
          @input="searchWeather"></v-text-field>
      </v-col>
    </v-row>
    <v-row v-if="weatherData.length > 0">
      <v-col cols="12" sm="6" md="4" v-for="(item, index) in weatherData" :key="index">
        <v-card class="transparent-card">
          <v-row>
            <v-col cols="8">
              <v-card-title>{{ item.City }}</v-card-title>
              <v-card-subtitle>Temperature: {{ item.Temperature }}°C</v-card-subtitle>
            </v-col>
            <v-col cols="4">
              <v-card-text>
                <v-icon
                  v-if="item.WeatherCondition.includes('vedro') || item.WeatherCondition.includes('sunčano')">mdi-weather-sunny</v-icon>
                <v-icon v-else-if="item.WeatherCondition.includes('oblačno')">mdi-weather-cloudy</v-icon>
                <v-icon
                  v-else-if="item.WeatherCondition.includes('kiša') || item.WeatherCondition.includes('pljusak')">mdi-weather-rainy</v-icon>
                <v-icon v-else-if="item.WeatherCondition.includes('grmljavina')">mdi-weather-lightning</v-icon>
                <v-icon
                  v-else-if="item.WeatherCondition.includes('vjetar') || item.WeatherCondition.includes('lahor')">mdi-weather-windy</v-icon>
                <v-icon v-else>mdi-cloud-question</v-icon></v-card-text>
            </v-col>
          </v-row>
        </v-card>
      </v-col>
    </v-row>
    <v-row v-else-if="searchCity !== ''">
      <v-col cols="12">
        <p>No matching cities found.</p>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts" setup>
import { ref, watch } from 'vue';
import { fetchWeather } from '@/api/weatherAPI';
import { VTextField, VContainer, VRow, VCol, VCard, VCardTitle, VCardSubtitle, VCardText, VIcon } from 'vuetify/components';
import { useSnackbar } from '@/components/SnackbarProvider.vue';
import { api } from '../../wailsjs/go/models';

const searchCity = ref('');
const weatherData = ref<api.CityWeatherInfo[]>([]);
const snackbar = useSnackbar()

const searchWeather = async () => {
  if (searchCity.value.trim() !== '') {
    try {
      weatherData.value = await fetchWeather(searchCity.value);
    } catch (error) {
      console.error("Error fetching weather data:", error);
      weatherData.value = [];
      snackbar.Error("Error fetching weather data")
    }
  } else {
    weatherData.value = [];
  }
};

watch(searchCity, () => {
  if (!searchCity.value) {
    weatherData.value = [];
  }
});

</script>

<style lang="css" scoped>
.transparent-card {
  background-color: transparent !important;
  padding: 16px !important;
  color: var(--font-color);
  font-weight: bold;
  box-shadow: 0 0 15px 2px #5373b364;
}
.v-text-field {
  color: var(--font-color);
}
h2, p {
  margin-bottom: 20px;
  color: var(--font-color);
}
.dashed-border {
    border: 2px dashed #697ea885;
    background-color: #5373b321 !important;
    max-height: 60px;
    padding-left: 10px;
    border-radius: 5px;
  }
</style>