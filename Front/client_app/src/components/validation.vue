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

const validationType = ref("xsd");
const uploadedFile = ref<File | null>(null);

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
    alert("Please upload a file.");
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
      alert(response.data);
    } else {
      alert("Validation failed.");
    }
  } catch (error) {
    console.error("Error during validation:", error);
    alert(`An error occurred during validation: ${error}`);
  }
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
</style>