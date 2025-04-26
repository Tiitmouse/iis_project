<template>
  <div>
    <h2>Validate with:</h2>
    <v-radio-group v-model="validationType" inline style="color: #f0e9e9;">
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
debugger
    if (!response?.Error) {
      snackbar.Success("validation success")
    } else {
      console.log(response.Error)
      snackbar.Error(`validation failed: ${response.Error}`)
    }
  } catch (error: any) {
    console.error("Error during validation:", error);

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
  background: var(--sheets-color);
}
.btnValidate {
  background-color: var(--button-color);
  color: var(--font-color);
  font-weight: bold;
  margin-top: 20px;
}
h2{
  margin-bottom: 20px;
  color: var(--font-color);
}
.v-list-item-title {
  color: var(--font-color);
}
.v-list-item-subtitle {
  color: var(--font-color);
  }
.mdi-close-circle{
  color: var(--font-color);
}
</style>