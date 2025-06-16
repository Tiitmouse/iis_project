<template>
  <div>
    <h2>Validate with JAXB</h2>
  </div>
  <div class="d-flex align-center my-2">
    <v-icon 
      v-if="showSoapStatusIcon"
      :icon="soapFileExists ? 'mdi-file-document-check-outline' : 'mdi-file-document-remove-outline'"
      :style="{ color: 'var(--button-color)' }"
      class="mr-2" 
      :size="72"
    ></v-icon>
    <span v-if="showSoapStatusIcon" class="soap-status-text">
      SOAP Service File Status: {{ soapFileExists ? 'Available' : 'Unavailable' }}
    </span>
  </div>
  <div>
    <v-btn class="btnValidate" rounded="xs" block @click="validateFile">
      Validate
    </v-btn>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue';

// TODO add java  virus
const validateFile = () => { 
  console.log('Validate button clicked (no validation configured)');
};

// status icon
const soapFileExists = ref(false);
const showSoapStatusIcon = ref(false); 

onMounted(async () => {
  // @ts-ignore: Accessing 'window.go', which is injected by Wails at runtime.
  const wailsGoApp = window.go?.main?.App;

  if (wailsGoApp && typeof wailsGoApp.CheckSoapFileExists === 'function') {
    try {
      // @ts-ignore: Calling a dynamically checked Wails function.
      const exists = await wailsGoApp.CheckSoapFileExists();
      soapFileExists.value = exists;
    } catch (err) {
      console.error("Error calling Wails backend function 'CheckSoapFileExists':", err);
      soapFileExists.value = false; 
    }
  } else {
    console.warn(
      "Wails backend function 'CheckSoapFileExists' is not available at 'window.go.main.App.CheckSoapFileExists'. " +
      "The icon will default to the 'file not found' state. " +
      "Please ensure the Go function is correctly implemented and bound by Wails."
    );
    soapFileExists.value = false; 
  }
  showSoapStatusIcon.value = true;
});
</script>

<style scoped>
h2 {
  margin-bottom: 20px;
  color: var(--font-color);
}

.btnValidate {
  background-color: var(--button-color);
  color: var(--font-color);
  font-weight: bold;
  margin-top: 20px;
}

.soap-status-text {
  color: var(--font-color);
}
</style>
