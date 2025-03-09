<!-- eslint-disable vue/max-attributes-per-line -->
<!-- eslint-disable vue/html-self-closing -->
<!-- eslint-disable vue/html-indent -->
<template>
  <div>
    <H3>Validate with:</H3>
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
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { VFileUpload } from "vuetify/labs/VFileUpload"

const validationType = ref("xsd")
const uploadedFilePath = ref<string | undefined>()

function handleFileUpload(event: Event) {
  const input = event.target as HTMLInputElement
  if (input.files && input.files[0]) {
    uploadedFilePath.value = input.files[0].name
  }
  console.log(uploadedFilePath)
}
</script>

<style>
.v-sheet {
    background: #5373B3;
}
</style>
