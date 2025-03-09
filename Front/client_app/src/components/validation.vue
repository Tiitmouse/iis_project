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
    ></v-file-upload>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { VFileUpload } from "vuetify/labs/VFileUpload"

const validationType = ref("xsd")
</script>

<style>
.v-sheet {
    background: #5373B3;
}
</style>
