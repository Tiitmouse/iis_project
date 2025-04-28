<template>
  <div>
    <h2>Validate with:</h2>
    <v-radio-group v-model="validationType" inline style="color: #f0e9e9;">
      <v-radio label="XSD" value="xsd"></v-radio>
      <v-radio label="RelaxNG" value="rng"></v-radio>
    </v-radio-group>
  </div>
  <div>
    <v-file-upload clearable density="compact" title="Click to upload .xml file" variant="compact"
      accept=".xml,text/xml,application/xml" @change="handleFileUpload" :v-model="fileInputModel"
      :key="fileInputKey"></v-file-upload>
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
const fileInputModel = ref<File | null>(null)
const fileInputKey = ref(0);

function handleFileUpload(event: Event) {
  const input = event.target as HTMLInputElement;
  const files = input.files;

  if (!files || files.length === 0) {
    uploadedFile.value = null;
    fileInputModel.value = null;
    return;
  }

  const file = files[0];

  if (!file.name.endsWith(".xml")) {
    snackbar.Error("please select an XML file.");
    input.value = "";
    uploadedFile.value = null;
    fileInputModel.value = null;
    fileInputKey.value++;
    return;
  }

  uploadedFile.value = file;
  console.log(uploadedFile.value);
}

async function validateFile() {
  if (!uploadedFile.value) {
    snackbar.Info("please upload an XML file.")
    return;
  }

  try {
    let response;
    if (validationType.value === "xsd") {
      response = await xsdValidate(uploadedFile.value);
    } else if (validationType.value === "rng") {
      response = await rngValidate(uploadedFile.value);
    }
    if (!response?.Error) {
      snackbar.Success("validation success");
    } else {
      console.log(response.Error);
      snackbar.Error(`validation failed: ${response.Error}`);
    }
  } catch (error: any) {
    console.error("Error during validation:", error);
  } finally {
    uploadedFile.value = null;
    fileInputModel.value = null;
    fileInputKey.value++; // forced rerender 
  }
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

h2 {
  margin-bottom: 20px;
  color: var(--font-color);
}

.v-list-item-title {
  color: var(--font-color);
}

.v-list-item-subtitle {
  color: var(--font-color);
}

.mdi-close-circle {
  color: var(--font-color);
}
</style>