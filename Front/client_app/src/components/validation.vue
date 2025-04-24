<template>
  <div>
    <h3>Validate with:</h3>
    <v-radio-group v-model="validationType" inline>
      <v-radio label="XSD" value="xsd"></v-radio>
      <v-radio label="RelaxNG" value="rng"></v-radio>
    </v-radio-group>
  </div>
  <div>
    <v-file-upload
      clearable
      density="compact"
      title="Drag and drop .xml file"
      variant="compact"
      accept=".xml"
      @change="handleFileUpload"
    ></v-file-upload>
  </div>
  <div>
    <v-btn class="btnValidate" rounded="xs" block @click="validateFile">
      Validate
    </v-btn>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { VFileUpload } from "vuetify/labs/VFileUpload";
import { xsdValidate, rngValidate } from "@/api/validationAPI";
import { useSnackbar } from '@/components/SnackbarProvider.vue';


const validationType = ref("xsd");
const uploadedFile = ref<File | null>(null);
const snackbar = useSnackbar()


function handleFileUpload(event: Event) {
  const input = event.target as HTMLInputElement;
  if (input.files && input.files[0]) {
    uploadedFile.value = input.files[0];
  } else {
    uploadedFile.value = null;
  }
  console.log(uploadedFile.value);
}

async function validateFile() {
  if (!uploadedFile.value) {
    snackbar.Info("Please upload a file.")
    return;
  }

  try {
    let response;
    if (validationType.value === "xsd") {
      response = await xsdValidate(uploadedFile.value);
    } else if (validationType.value === "rng") {
      response = await rngValidate(uploadedFile.value);
    }

    if (response && response.data) {
      snackbar.Success("validation success")
    } else {
      snackbar.Error("validation failed")
    }
  } catch (error: any) {
    console.error("Error during validation:", error);

    if (error.response && error.response.data) {
      const errorData = error.response.data;

      if (typeof errorData === 'string' && errorData.includes('<')) {
        const errorDetails = parseErrorMessage(errorData);
        snackbar.Error(`Validation error: ${errorDetails}`);
      } else if (typeof errorData === 'object') {
        const errorMessage = errorData.message || JSON.stringify(errorData);
        snackbar.Error(`Validation error: ${errorMessage}`);
      }
      else {
        snackbar.Error(`An error occurred during validation: ${errorData}`);
      }
    } else {
      snackbar.Error(`An error occurred during validation: ${error}`);
    }
  }
}

function parseErrorMessage(errorMessage: string) {
  const lineRegex = /Line: (\d+)/i;
  const tagRegex = /Tag: ([\w:]+)/i;
  const messageRegex = /Message: (.+)/i;

  const lineMatch = errorMessage.match(lineRegex);
  const tagMatch = errorMessage.match(tagRegex);
  const messageMatch = errorMessage.match(messageRegex);

  const line = lineMatch ? lineMatch[1] : 'N/A';
  const tag = tagMatch ? tagMatch[1] : 'N/A';
  const message = messageMatch ? messageMatch[1] : 'Unknown error';

  return `Line: ${line}, Tag: ${tag}, Message: ${message}`;
}
</script>

<style>
.v-sheet {
  background: #5373b3;
}
.btnValidate {
  background-color: white;
  color: #5373b3;
  font-weight: bold;
  margin-top: 20px;
  box-shadow: 0 0 15px 2px #5373b3;
}
h3{
  margin-bottom: 20px
}
</style>