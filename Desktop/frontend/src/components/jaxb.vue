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
import { useSnackbar } from '@/components/SnackbarProvider.vue'; // Added import

const snackbar = useSnackbar(); // Added snackbar instance

// TODO add java  virus
const validateFile = async () => { 
  // @ts-ignore: Accessing 'window.go', which is injected by Wails at runtime.
  const wailsGoApp = window.go?.main?.App;

  if (wailsGoApp && typeof wailsGoApp.RunJaxbValidation === 'function') {
    try {
      // @ts-ignore: Calling a dynamically checked Wails function.
      const result = await wailsGoApp.RunJaxbValidation();
      console.log('JAXB Process Output:', result);

      if (result.includes("Unmarshalling successful")) {
        snackbar.Success("JAXB validation successful!");
      } else if (result.includes("VALIDATION EVENT")) {
        // Extracting a bit of the message for the snackbar
        const eventMsgMatch = result.match(/MESSAGE: ([^\n]+)/);
        const detail = eventMsgMatch && eventMsgMatch[1] ? eventMsgMatch[1] : "See console for details.";
        snackbar.Warning(`JAXB validation event: ${detail}`);
      } else if (result.includes("XML processing failed") || result.toLowerCase().includes("error") || result.toLowerCase().includes("exception")) {
        snackbar.Error("JAXB validation failed. See console for details.");
      } else if (result.trim() === "") {
        snackbar.Error("JAXB validation returned empty output. Check backend logs.");
      } else {
        snackbar.Info("JAXB process finished. See console for output.");
      }
    } catch (err: any) {
      console.error("Error calling Wails backend function 'RunJaxbValidation':", err);
      snackbar.Error(`JAXB validation error: ${err.message || err}`);
    }
  } else {
    console.warn(
      "Wails backend function 'RunJaxbValidation' is not available at 'window.go.main.App.RunJaxbValidation'. " +
      "Please ensure the Go function is correctly implemented and bound by Wails."
    );
    snackbar.Error("JAXB validation function not available on backend.");
  }
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
