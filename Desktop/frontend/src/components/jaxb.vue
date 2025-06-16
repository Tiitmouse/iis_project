<template>
  <div>
    <h2>Validate with JAXB</h2>
  </div>
  <div class="d-flex align-center my-2">
    <v-icon 
      v-if="showSoapStatusIcon || validationAttempted" 
      :icon="validationAttempted ? validationStatusIcon : (soapFileExists ? 'mdi-file-document-check-outline' : 'mdi-file-document-remove-outline')"
      :style="{ color: validationAttempted ? validationStatusColor : 'var(--button-color)' }"
      class="mr-2" 
      :size="72"
    ></v-icon>
    <span v-if="showSoapStatusIcon && !validationAttempted" class="soap-status-text">
      SOAP Service File Status: {{ soapFileExists ? 'Available' : 'Unavailable' }}
    </span>
    <span v-if="validationAttempted" :style="{ color: validationStatusColor }" class="validation-status-text">
      JAXB Validation: {{ validationStatusText }}
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
import { useSnackbar } from '@/components/SnackbarProvider.vue';

const snackbar = useSnackbar();

const validationAttempted = ref(false);
const validationStatusIcon = ref('mdi-file-document-outline');
const validationStatusColor = ref('var(--button-color)'); 
const validationStatusText = ref('');

const validateFile = async () => { 
  validationAttempted.value = true;
  // @ts-ignore: accessing 'window.go'injected by 1ails at runtime
  const wailsGoApp = window.go?.main?.App;

  if (wailsGoApp && typeof wailsGoApp.RunJaxbValidation === 'function') {
    try {
      // @ts-ignore:calling a dynamically checked wails function
      const result = await wailsGoApp.RunJaxbValidation();
      console.log('JAXB Process Output:', result);

      if (result.includes("Unmarshalling successful")) {
        snackbar.Success("JAXB validation successful!");
        validationStatusIcon.value = 'mdi-file-certificate-outline';
        validationStatusColor.value = 'green';
        validationStatusText.value = 'Successful';
      } else if (result.includes("VALIDATION EVENT")) {
        const eventMsgMatch = result.match(/MESSAGE: ([^\n]+)/);
        const detail = eventMsgMatch && eventMsgMatch[1] ? eventMsgMatch[1] : "See console for details.";
        snackbar.Warning(`JAXB validation event: ${detail}`);
        validationStatusIcon.value = 'mdi-file-alert-outline';
        validationStatusColor.value = 'orange';
        validationStatusText.value = 'Warning';
      } else if (result.includes("XML processing failed") || result.toLowerCase().includes("error") || result.toLowerCase().includes("exception")) {
        snackbar.Error("JAXB validation failed. See console for details.");
        validationStatusIcon.value = 'mdi-file-cancel-outline';
        validationStatusColor.value = 'red';
        validationStatusText.value = 'Failed';
      } else if (result.trim() === "") {
        snackbar.Error("JAXB validation returned empty output. Check backend logs.");
        validationStatusIcon.value = 'mdi-file-question-outline';
        validationStatusColor.value = 'grey';
        validationStatusText.value = 'Empty Output';
      } else {
        snackbar.Info("JAXB process finished. See console for output.");
        validationStatusIcon.value = 'mdi-file-eye-outline';
        validationStatusColor.value = 'blue';
        validationStatusText.value = 'Info';
      }
    } catch (err: any) {
      console.error("Error calling Wails backend function 'RunJaxbValidation':", err);
      snackbar.Error(`JAXB validation error: ${err.message || err}`);
      validationStatusIcon.value = 'mdi-file-cancel-outline';
      validationStatusColor.value = 'red';
      validationStatusText.value = 'Error';
    }
  } else {
    console.warn(
      "Wails backend function 'RunJaxbValidation' is not available at 'window.go.main.App.RunJaxbValidation'. " +
      "Please ensure the Go function is correctly implemented and bound by Wails."
    );
    snackbar.Error("JAXB validation function not available on backend.");
    validationStatusIcon.value = 'mdi-file-remove-outline';
    validationStatusColor.value = 'grey';
    validationStatusText.value = 'Unavailable';
  }
};

// status icon
const soapFileExists = ref(false);
const showSoapStatusIcon = ref(false); 

onMounted(async () => {
  // @ts-ignore: Accessing 'window.go' injected by wails at runtime
  const wailsGoApp = window.go?.main?.App;

  if (wailsGoApp && typeof wailsGoApp.CheckSoapFileExists === 'function') {
    try {
      // @ts-ignore: calling dynamically checked wails function.
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

.validation-status-text {
  font-weight: bold;
}
</style>
